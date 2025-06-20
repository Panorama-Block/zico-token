package ocr

import (
	"context"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"go.uber.org/ratelimit"

	client2 "github.com/smartcontractkit/chainlink-testing-framework/lib/client"
	"github.com/smartcontractkit/chainlink-testing-framework/seth"
	"github.com/smartcontractkit/chainlink-testing-framework/wasp"

	"github.com/smartcontractkit/chainlink/deployment/environment/nodeclient"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/testconfig/ocr"
)

// VU is a virtual user for the OCR load test
// it creates a feed and triggers new rounds
type VU struct {
	*wasp.VUControl
	rl            ratelimit.Limiter
	rate          int
	rateUnit      time.Duration
	roundNum      atomic.Int64
	seth          *seth.Client
	lta           common.Address
	bootstrapNode *nodeclient.ChainlinkK8sClient
	workerNodes   []*nodeclient.ChainlinkK8sClient
	msClient      *client2.MockserverClient //nolint:staticcheck //SA1019 no need to upgrade
	l             zerolog.Logger
	ocrInstances  []contracts.OffchainAggregator
	config        ocr.OffChainAggregatorsConfig
}

func NewVU(
	l zerolog.Logger,
	seth *seth.Client,
	config ocr.OffChainAggregatorsConfig,
	rate int,
	rateUnit time.Duration,
	lta common.Address,
	bootstrapNode *nodeclient.ChainlinkK8sClient,
	workerNodes []*nodeclient.ChainlinkK8sClient,
	msClient *client2.MockserverClient, //nolint:staticcheck //SA1019 no need to upgrade
) *VU {
	return &VU{
		VUControl:     wasp.NewVUControl(),
		rl:            ratelimit.New(rate, ratelimit.Per(rateUnit)),
		rate:          rate,
		rateUnit:      rateUnit,
		l:             l,
		seth:          seth,
		lta:           lta,
		msClient:      msClient,
		bootstrapNode: bootstrapNode,
		workerNodes:   workerNodes,
		config:        config,
	}
}

func (m *VU) Clone(_ *wasp.Generator) wasp.VirtualUser {
	return &VU{
		VUControl:     wasp.NewVUControl(),
		rl:            ratelimit.New(m.rate, ratelimit.Per(m.rateUnit)),
		rate:          m.rate,
		rateUnit:      m.rateUnit,
		l:             m.l,
		seth:          m.seth,
		lta:           m.lta,
		msClient:      m.msClient,
		bootstrapNode: m.bootstrapNode,
		workerNodes:   m.workerNodes,
		config:        m.config,
	}
}

func (m *VU) Setup(_ *wasp.Generator) error {
	ocrInstances, err := actions.SetupOCRv1Contracts(m.l, m.seth, m.config, m.lta, contracts.ChainlinkK8sClientToChainlinkNodeWithKeysAndAddress(m.workerNodes))
	if err != nil {
		return err
	}
	err = actions.CreateOCRJobs(ocrInstances, m.bootstrapNode, m.workerNodes, 5, m.msClient, strconv.FormatInt(m.seth.ChainID, 10))
	if err != nil {
		return err
	}
	m.ocrInstances = ocrInstances
	return nil
}

func (m *VU) Teardown(_ *wasp.Generator) error {
	return nil
}

func (m *VU) Call(l *wasp.Generator) {
	m.rl.Take()
	m.roundNum.Add(1)
	requestedRound := m.roundNum.Load()
	m.l.Info().
		Int64("RoundNum", requestedRound).
		Str("FeedID", m.ocrInstances[0].Address()).
		Msg("starting new round")
	err := m.ocrInstances[0].RequestNewRound()
	if err != nil {
		l.ResponsesChan <- &wasp.Response{Error: err.Error(), Failed: true}
	}
	for {
		time.Sleep(5 * time.Second)
		lr, err := m.ocrInstances[0].GetLatestRound(context.Background())
		if err != nil {
			l.ResponsesChan <- &wasp.Response{Error: err.Error(), Failed: true}
		}
		m.l.Info().Interface("LatestRound", lr).Msg("latest round")
		if lr.RoundId.Int64() >= requestedRound {
			l.ResponsesChan <- &wasp.Response{}
		}
	}
}
