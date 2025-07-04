// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: proxy.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MockCapability_List_FullMethodName                   = "/mockcap.MockCapability/List"
	MockCapability_GetTriggerSubscribers_FullMethodName  = "/mockcap.MockCapability/GetTriggerSubscribers"
	MockCapability_CreateCapability_FullMethodName       = "/mockcap.MockCapability/CreateCapability"
	MockCapability_SendTriggerEvent_FullMethodName       = "/mockcap.MockCapability/SendTriggerEvent"
	MockCapability_RegisterTrigger_FullMethodName        = "/mockcap.MockCapability/RegisterTrigger"
	MockCapability_UnregisterTrigger_FullMethodName      = "/mockcap.MockCapability/UnregisterTrigger"
	MockCapability_HookExecutables_FullMethodName        = "/mockcap.MockCapability/HookExecutables"
	MockCapability_RegisterToWorkflow_FullMethodName     = "/mockcap.MockCapability/RegisterToWorkflow"
	MockCapability_UnregisterFromWorkflow_FullMethodName = "/mockcap.MockCapability/UnregisterFromWorkflow"
	MockCapability_Execute_FullMethodName                = "/mockcap.MockCapability/Execute"
)

// MockCapabilityClient is the client API for MockCapability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MockCapabilityClient interface {
	// Retrieve information about all capabilities available on the node
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Retrieve unique workflow IDs of subscribers for a specific trigger
	GetTriggerSubscribers(ctx context.Context, in *GetTriggerSubscribersRequest, opts ...grpc.CallOption) (*GetTriggerSubscribersResponse, error)
	// Create a mock capability and register it with the node
	CreateCapability(ctx context.Context, in *CapabilityInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Send data through a mock trigger
	SendTriggerEvent(ctx context.Context, in *SendTriggerEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Subscribe to a trigger (includes all triggers, not limited to mock triggers),
	// creates a stream that send trigger events
	RegisterTrigger(ctx context.Context, in *TriggerRegistrationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TriggerResponse], error)
	// Unsubscribe from a trigger (includes all triggers, not limited to mock triggers)
	UnregisterTrigger(ctx context.Context, in *TriggerRegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Establish a bidirectional streaming service. When Execute is called, it streams requests and allows streaming responses back
	HookExecutables(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ExecutableResponse, ExecutableRequest], error)
	// Subscribe to a workflow (includes all executable capabilities, not limited to mocks)
	RegisterToWorkflow(ctx context.Context, in *RegisterToWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Unsubscribe from a workflow (includes all executable capabilities, not limited to mocks)
	UnregisterFromWorkflow(ctx context.Context, in *UnregisterFromWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Invoke the Execute method on an executable capability (includes all executable capabilities, not limited to mocks)
	Execute(ctx context.Context, in *ExecutableRequest, opts ...grpc.CallOption) (*CapabilityResponse, error)
}

type mockCapabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewMockCapabilityClient(cc grpc.ClientConnInterface) MockCapabilityClient {
	return &mockCapabilityClient{cc}
}

func (c *mockCapabilityClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, MockCapability_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockCapabilityClient) GetTriggerSubscribers(ctx context.Context, in *GetTriggerSubscribersRequest, opts ...grpc.CallOption) (*GetTriggerSubscribersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTriggerSubscribersResponse)
	err := c.cc.Invoke(ctx, MockCapability_GetTriggerSubscribers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockCapabilityClient) CreateCapability(ctx context.Context, in *CapabilityInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockCapability_CreateCapability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockCapabilityClient) SendTriggerEvent(ctx context.Context, in *SendTriggerEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockCapability_SendTriggerEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockCapabilityClient) RegisterTrigger(ctx context.Context, in *TriggerRegistrationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TriggerResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MockCapability_ServiceDesc.Streams[0], MockCapability_RegisterTrigger_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TriggerRegistrationRequest, TriggerResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MockCapability_RegisterTriggerClient = grpc.ServerStreamingClient[TriggerResponse]

func (c *mockCapabilityClient) UnregisterTrigger(ctx context.Context, in *TriggerRegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockCapability_UnregisterTrigger_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockCapabilityClient) HookExecutables(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ExecutableResponse, ExecutableRequest], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MockCapability_ServiceDesc.Streams[1], MockCapability_HookExecutables_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ExecutableResponse, ExecutableRequest]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MockCapability_HookExecutablesClient = grpc.BidiStreamingClient[ExecutableResponse, ExecutableRequest]

func (c *mockCapabilityClient) RegisterToWorkflow(ctx context.Context, in *RegisterToWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockCapability_RegisterToWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockCapabilityClient) UnregisterFromWorkflow(ctx context.Context, in *UnregisterFromWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MockCapability_UnregisterFromWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockCapabilityClient) Execute(ctx context.Context, in *ExecutableRequest, opts ...grpc.CallOption) (*CapabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CapabilityResponse)
	err := c.cc.Invoke(ctx, MockCapability_Execute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MockCapabilityServer is the server API for MockCapability service.
// All implementations must embed UnimplementedMockCapabilityServer
// for forward compatibility.
type MockCapabilityServer interface {
	// Retrieve information about all capabilities available on the node
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Retrieve unique workflow IDs of subscribers for a specific trigger
	GetTriggerSubscribers(context.Context, *GetTriggerSubscribersRequest) (*GetTriggerSubscribersResponse, error)
	// Create a mock capability and register it with the node
	CreateCapability(context.Context, *CapabilityInfo) (*emptypb.Empty, error)
	// Send data through a mock trigger
	SendTriggerEvent(context.Context, *SendTriggerEventRequest) (*emptypb.Empty, error)
	// Subscribe to a trigger (includes all triggers, not limited to mock triggers),
	// creates a stream that send trigger events
	RegisterTrigger(*TriggerRegistrationRequest, grpc.ServerStreamingServer[TriggerResponse]) error
	// Unsubscribe from a trigger (includes all triggers, not limited to mock triggers)
	UnregisterTrigger(context.Context, *TriggerRegistrationRequest) (*emptypb.Empty, error)
	// Establish a bidirectional streaming service. When Execute is called, it streams requests and allows streaming responses back
	HookExecutables(grpc.BidiStreamingServer[ExecutableResponse, ExecutableRequest]) error
	// Subscribe to a workflow (includes all executable capabilities, not limited to mocks)
	RegisterToWorkflow(context.Context, *RegisterToWorkflowRequest) (*emptypb.Empty, error)
	// Unsubscribe from a workflow (includes all executable capabilities, not limited to mocks)
	UnregisterFromWorkflow(context.Context, *UnregisterFromWorkflowRequest) (*emptypb.Empty, error)
	// Invoke the Execute method on an executable capability (includes all executable capabilities, not limited to mocks)
	Execute(context.Context, *ExecutableRequest) (*CapabilityResponse, error)
	mustEmbedUnimplementedMockCapabilityServer()
}

// UnimplementedMockCapabilityServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMockCapabilityServer struct{}

func (UnimplementedMockCapabilityServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMockCapabilityServer) GetTriggerSubscribers(context.Context, *GetTriggerSubscribersRequest) (*GetTriggerSubscribersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTriggerSubscribers not implemented")
}
func (UnimplementedMockCapabilityServer) CreateCapability(context.Context, *CapabilityInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCapability not implemented")
}
func (UnimplementedMockCapabilityServer) SendTriggerEvent(context.Context, *SendTriggerEventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTriggerEvent not implemented")
}
func (UnimplementedMockCapabilityServer) RegisterTrigger(*TriggerRegistrationRequest, grpc.ServerStreamingServer[TriggerResponse]) error {
	return status.Errorf(codes.Unimplemented, "method RegisterTrigger not implemented")
}
func (UnimplementedMockCapabilityServer) UnregisterTrigger(context.Context, *TriggerRegistrationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterTrigger not implemented")
}
func (UnimplementedMockCapabilityServer) HookExecutables(grpc.BidiStreamingServer[ExecutableResponse, ExecutableRequest]) error {
	return status.Errorf(codes.Unimplemented, "method HookExecutables not implemented")
}
func (UnimplementedMockCapabilityServer) RegisterToWorkflow(context.Context, *RegisterToWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterToWorkflow not implemented")
}
func (UnimplementedMockCapabilityServer) UnregisterFromWorkflow(context.Context, *UnregisterFromWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterFromWorkflow not implemented")
}
func (UnimplementedMockCapabilityServer) Execute(context.Context, *ExecutableRequest) (*CapabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedMockCapabilityServer) mustEmbedUnimplementedMockCapabilityServer() {}
func (UnimplementedMockCapabilityServer) testEmbeddedByValue()                        {}

// UnsafeMockCapabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MockCapabilityServer will
// result in compilation errors.
type UnsafeMockCapabilityServer interface {
	mustEmbedUnimplementedMockCapabilityServer()
}

func RegisterMockCapabilityServer(s grpc.ServiceRegistrar, srv MockCapabilityServer) {
	// If the following call pancis, it indicates UnimplementedMockCapabilityServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MockCapability_ServiceDesc, srv)
}

func _MockCapability_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockCapability_GetTriggerSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTriggerSubscribersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).GetTriggerSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_GetTriggerSubscribers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).GetTriggerSubscribers(ctx, req.(*GetTriggerSubscribersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockCapability_CreateCapability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilityInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).CreateCapability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_CreateCapability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).CreateCapability(ctx, req.(*CapabilityInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockCapability_SendTriggerEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTriggerEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).SendTriggerEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_SendTriggerEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).SendTriggerEvent(ctx, req.(*SendTriggerEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockCapability_RegisterTrigger_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TriggerRegistrationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MockCapabilityServer).RegisterTrigger(m, &grpc.GenericServerStream[TriggerRegistrationRequest, TriggerResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MockCapability_RegisterTriggerServer = grpc.ServerStreamingServer[TriggerResponse]

func _MockCapability_UnregisterTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).UnregisterTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_UnregisterTrigger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).UnregisterTrigger(ctx, req.(*TriggerRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockCapability_HookExecutables_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MockCapabilityServer).HookExecutables(&grpc.GenericServerStream[ExecutableResponse, ExecutableRequest]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MockCapability_HookExecutablesServer = grpc.BidiStreamingServer[ExecutableResponse, ExecutableRequest]

func _MockCapability_RegisterToWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterToWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).RegisterToWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_RegisterToWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).RegisterToWorkflow(ctx, req.(*RegisterToWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockCapability_UnregisterFromWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterFromWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).UnregisterFromWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_UnregisterFromWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).UnregisterFromWorkflow(ctx, req.(*UnregisterFromWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MockCapability_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockCapabilityServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MockCapability_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockCapabilityServer).Execute(ctx, req.(*ExecutableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MockCapability_ServiceDesc is the grpc.ServiceDesc for MockCapability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MockCapability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockcap.MockCapability",
	HandlerType: (*MockCapabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _MockCapability_List_Handler,
		},
		{
			MethodName: "GetTriggerSubscribers",
			Handler:    _MockCapability_GetTriggerSubscribers_Handler,
		},
		{
			MethodName: "CreateCapability",
			Handler:    _MockCapability_CreateCapability_Handler,
		},
		{
			MethodName: "SendTriggerEvent",
			Handler:    _MockCapability_SendTriggerEvent_Handler,
		},
		{
			MethodName: "UnregisterTrigger",
			Handler:    _MockCapability_UnregisterTrigger_Handler,
		},
		{
			MethodName: "RegisterToWorkflow",
			Handler:    _MockCapability_RegisterToWorkflow_Handler,
		},
		{
			MethodName: "UnregisterFromWorkflow",
			Handler:    _MockCapability_UnregisterFromWorkflow_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _MockCapability_Execute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterTrigger",
			Handler:       _MockCapability_RegisterTrigger_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HookExecutables",
			Handler:       _MockCapability_HookExecutables_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proxy.proto",
}
