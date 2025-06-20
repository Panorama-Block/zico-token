package remote

import (
	"context"
	"errors"
	"sync"
	"time"

	commoncap "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/aggregation"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/messagecache"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
)

// TriggerSubscriber is a shim for remote trigger capabilities.
// It translatesd between capability API calls and network messages.
// Its responsibilities are:
//  1. Periodically refresh all registrations for remote triggers.
//  2. Collect trigger events from remote nodes and aggregate responses via a customizable aggregator.
//
// TriggerSubscriber communicates with corresponding TriggerReceivers on remote nodes.
type triggerSubscriber struct {
	config              *commoncap.RemoteTriggerConfig
	capInfo             commoncap.CapabilityInfo
	capDonInfo          commoncap.DON
	capDonMembers       map[p2ptypes.PeerID]struct{}
	localDonInfo        commoncap.DON
	dispatcher          types.Dispatcher
	aggregator          types.Aggregator
	messageCache        *messagecache.MessageCache[triggerEventKey, p2ptypes.PeerID]
	registeredWorkflows map[string]*subRegState
	mu                  sync.RWMutex // protects registeredWorkflows and messageCache
	stopCh              services.StopChan
	wg                  sync.WaitGroup
	lggr                logger.Logger
}

type triggerEventKey struct {
	triggerEventID string
	workflowID     string
}

type subRegState struct {
	callback   chan commoncap.TriggerResponse
	rawRequest []byte
}

type TriggerSubscriber interface {
	commoncap.TriggerCapability
	Receive(ctx context.Context, msg *types.MessageBody)
}

var _ commoncap.TriggerCapability = &triggerSubscriber{}
var _ types.Receiver = &triggerSubscriber{}
var _ services.Service = &triggerSubscriber{}

// TODO makes this configurable with a default
const (
	defaultSendChannelBufferSize = 1000
	maxBatchedWorkflowIDs        = 1000
)

func NewTriggerSubscriber(config *commoncap.RemoteTriggerConfig, capInfo commoncap.CapabilityInfo, capDonInfo commoncap.DON, localDonInfo commoncap.DON, dispatcher types.Dispatcher, aggregator types.Aggregator, lggr logger.Logger) *triggerSubscriber {
	if aggregator == nil {
		lggr.Warnw("no aggregator provided, using default MODE aggregator", "capabilityId", capInfo.ID)
		aggregator = aggregation.NewDefaultModeAggregator(uint32(capDonInfo.F + 1))
	}
	if config == nil {
		lggr.Info("no config provided, using default values")
		config = &commoncap.RemoteTriggerConfig{}
	}
	config.ApplyDefaults()
	capDonMembers := make(map[p2ptypes.PeerID]struct{})
	for _, member := range capDonInfo.Members {
		capDonMembers[member] = struct{}{}
	}
	return &triggerSubscriber{
		config:              config,
		capInfo:             capInfo,
		capDonInfo:          capDonInfo,
		capDonMembers:       capDonMembers,
		localDonInfo:        localDonInfo,
		dispatcher:          dispatcher,
		aggregator:          aggregator,
		messageCache:        messagecache.NewMessageCache[triggerEventKey, p2ptypes.PeerID](),
		registeredWorkflows: make(map[string]*subRegState),
		stopCh:              make(services.StopChan),
		lggr:                logger.Named(lggr, "TriggerSubscriber"),
	}
}

func (s *triggerSubscriber) Start(ctx context.Context) error {
	s.wg.Add(2)
	go s.registrationLoop()
	go s.eventCleanupLoop()
	s.lggr.Info("TriggerSubscriber started")
	return nil
}

func (s *triggerSubscriber) Info(ctx context.Context) (commoncap.CapabilityInfo, error) {
	return s.capInfo, nil
}

func (s *triggerSubscriber) RegisterTrigger(ctx context.Context, request commoncap.TriggerRegistrationRequest) (<-chan commoncap.TriggerResponse, error) {
	rawRequest, err := pb.MarshalTriggerRegistrationRequest(request)
	if err != nil {
		return nil, err
	}
	if request.Metadata.WorkflowID == "" {
		return nil, errors.New("empty workflowID")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lggr.Infow("RegisterTrigger called", "capabilityId", s.capInfo.ID, "donId", s.capDonInfo.ID, "workflowID", request.Metadata.WorkflowID)
	regState, ok := s.registeredWorkflows[request.Metadata.WorkflowID]
	if !ok {
		regState = &subRegState{
			callback:   make(chan commoncap.TriggerResponse, defaultSendChannelBufferSize),
			rawRequest: rawRequest,
		}
		s.registeredWorkflows[request.Metadata.WorkflowID] = regState
	} else {
		regState.rawRequest = rawRequest
		s.lggr.Warnw("RegisterTrigger re-registering trigger", "capabilityId", s.capInfo.ID, "donId", s.capDonInfo.ID, "workflowID", request.Metadata.WorkflowID)
	}

	return regState.callback, nil
}

func (s *triggerSubscriber) registrationLoop() {
	defer s.wg.Done()
	ticker := time.NewTicker(s.config.RegistrationRefresh)
	defer ticker.Stop()
	for {
		select {
		case <-s.stopCh:
			return
		case <-ticker.C:
			s.mu.RLock()
			s.lggr.Infow("register trigger for remote capability", "capabilityId", s.capInfo.ID, "donId", s.capDonInfo.ID, "nMembers", len(s.capDonInfo.Members), "nWorkflows", len(s.registeredWorkflows))
			if len(s.registeredWorkflows) == 0 {
				s.lggr.Infow("no workflows to register")
			}
			for _, registration := range s.registeredWorkflows {
				// NOTE: send to all by default, introduce different strategies later (KS-76)
				for _, peerID := range s.capDonInfo.Members {
					m := &types.MessageBody{
						CapabilityId:    s.capInfo.ID,
						CapabilityDonId: s.capDonInfo.ID,
						CallerDonId:     s.localDonInfo.ID,
						Method:          types.MethodRegisterTrigger,
						Payload:         registration.rawRequest,
					}
					err := s.dispatcher.Send(peerID, m)
					if err != nil {
						s.lggr.Errorw("failed to send message", "capabilityId", s.capInfo.ID, "donId", s.capDonInfo.ID, "peerId", peerID, "err", err)
					}
				}
			}
			s.mu.RUnlock()
		}
	}
}

func (s *triggerSubscriber) UnregisterTrigger(ctx context.Context, request commoncap.TriggerRegistrationRequest) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	state := s.registeredWorkflows[request.Metadata.WorkflowID]
	if state != nil && state.callback != nil {
		close(state.callback)
	}
	delete(s.registeredWorkflows, request.Metadata.WorkflowID)
	// Registrations will quickly expire on all remote nodes.
	// Alternatively, we could send UnregisterTrigger messages right away.
	return nil
}

func (s *triggerSubscriber) Receive(_ context.Context, msg *types.MessageBody) {
	sender, err := ToPeerID(msg.Sender)
	if err != nil {
		s.lggr.Errorw("failed to convert message sender to PeerID", "err", err)
		return
	}

	if _, found := s.capDonMembers[sender]; !found {
		s.lggr.Errorw("received message from unexpected node", "capabilityId", s.capInfo.ID, "sender", sender)
		return
	}
	if msg.Method == types.MethodTriggerEvent {
		meta := msg.GetTriggerEventMetadata()
		if meta == nil {
			s.lggr.Errorw("received message with invalid trigger metadata", "capabilityId", s.capInfo.ID, "sender", sender)
			return
		}
		if len(meta.WorkflowIds) > maxBatchedWorkflowIDs {
			s.lggr.Errorw("received message with too many workflow IDs - truncating", "capabilityId", s.capInfo.ID, "nWorkflows", len(meta.WorkflowIds), "sender", sender)
			meta.WorkflowIds = meta.WorkflowIds[:maxBatchedWorkflowIDs]
		}
		for _, workflowID := range meta.WorkflowIds {
			s.mu.RLock()
			registration, found := s.registeredWorkflows[workflowID]
			s.mu.RUnlock()
			if !found {
				s.lggr.Errorw("received message for unregistered workflow", "capabilityId", s.capInfo.ID, "workflowID", SanitizeLogString(workflowID), "sender", sender)
				continue
			}
			key := triggerEventKey{
				triggerEventID: meta.TriggerEventId,
				workflowID:     workflowID,
			}
			nowMs := time.Now().UnixMilli()
			s.mu.Lock()
			creationTs := s.messageCache.Insert(key, sender, nowMs, msg.Payload)
			ready, payloads := s.messageCache.Ready(key, s.config.MinResponsesToAggregate, nowMs-s.config.MessageExpiry.Milliseconds(), true)
			s.mu.Unlock()
			s.lggr.Debugw("trigger event received", "triggerEventId", meta.TriggerEventId, "capabilityId", s.capInfo.ID, "workflowId", workflowID, "sender", sender, "ready", ready, "nowTs", nowMs, "creationTs", creationTs, "minResponsesToAggregate", s.config.MinResponsesToAggregate)
			if ready {
				aggregatedResponse, err := s.aggregator.Aggregate(meta.TriggerEventId, payloads)
				if err != nil {
					s.lggr.Errorw("failed to aggregate responses", "triggerEventID", meta.TriggerEventId, "capabilityId", s.capInfo.ID, "workflowId", workflowID, "err", err)
					continue
				}
				s.lggr.Infow("remote trigger event aggregated", "triggerEventID", meta.TriggerEventId, "capabilityId", s.capInfo.ID, "workflowId", workflowID)
				registration.callback <- aggregatedResponse
			}
		}
	} else {
		s.lggr.Errorw("received trigger event with unknown method", "method", SanitizeLogString(msg.Method), "sender", sender)
	}
}

func (s *triggerSubscriber) eventCleanupLoop() {
	defer s.wg.Done()
	ticker := time.NewTicker(s.config.MessageExpiry)
	defer ticker.Stop()
	for {
		select {
		case <-s.stopCh:
			return
		case <-ticker.C:
			s.mu.Lock()
			s.messageCache.DeleteOlderThan(time.Now().UnixMilli() - s.config.MessageExpiry.Milliseconds())
			s.mu.Unlock()
		}
	}
}

func (s *triggerSubscriber) Close() error {
	close(s.stopCh)
	s.wg.Wait()
	s.lggr.Info("TriggerSubscriber closed")
	return nil
}

func (s *triggerSubscriber) Ready() error {
	return nil
}

func (s *triggerSubscriber) HealthReport() map[string]error {
	return nil
}

func (s *triggerSubscriber) Name() string {
	return s.lggr.Name()
}
