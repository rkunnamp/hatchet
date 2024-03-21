// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: events.proto

package contracts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventsServiceClient is the client API for EventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsServiceClient interface {
	Push(ctx context.Context, in *PushEventRequest, opts ...grpc.CallOption) (*Event, error)
	ReplaySingleEvent(ctx context.Context, in *ReplayEventRequest, opts ...grpc.CallOption) (*Event, error)
	PutLog(ctx context.Context, in *PutLogRequest, opts ...grpc.CallOption) (*PutLogResponse, error)
}

type eventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsServiceClient(cc grpc.ClientConnInterface) EventsServiceClient {
	return &eventsServiceClient{cc}
}

func (c *eventsServiceClient) Push(ctx context.Context, in *PushEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/EventsService/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) ReplaySingleEvent(ctx context.Context, in *ReplayEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/EventsService/ReplaySingleEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) PutLog(ctx context.Context, in *PutLogRequest, opts ...grpc.CallOption) (*PutLogResponse, error) {
	out := new(PutLogResponse)
	err := c.cc.Invoke(ctx, "/EventsService/PutLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServiceServer is the server API for EventsService service.
// All implementations must embed UnimplementedEventsServiceServer
// for forward compatibility
type EventsServiceServer interface {
	Push(context.Context, *PushEventRequest) (*Event, error)
	ReplaySingleEvent(context.Context, *ReplayEventRequest) (*Event, error)
	PutLog(context.Context, *PutLogRequest) (*PutLogResponse, error)
	mustEmbedUnimplementedEventsServiceServer()
}

// UnimplementedEventsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventsServiceServer struct {
}

func (UnimplementedEventsServiceServer) Push(context.Context, *PushEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedEventsServiceServer) ReplaySingleEvent(context.Context, *ReplayEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaySingleEvent not implemented")
}
func (UnimplementedEventsServiceServer) PutLog(context.Context, *PutLogRequest) (*PutLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutLog not implemented")
}
func (UnimplementedEventsServiceServer) mustEmbedUnimplementedEventsServiceServer() {}

// UnsafeEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServiceServer will
// result in compilation errors.
type UnsafeEventsServiceServer interface {
	mustEmbedUnimplementedEventsServiceServer()
}

func RegisterEventsServiceServer(s grpc.ServiceRegistrar, srv EventsServiceServer) {
	s.RegisterService(&EventsService_ServiceDesc, srv)
}

func _EventsService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventsService/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).Push(ctx, req.(*PushEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_ReplaySingleEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplayEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).ReplaySingleEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventsService/ReplaySingleEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).ReplaySingleEvent(ctx, req.(*ReplayEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_PutLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).PutLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventsService/PutLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).PutLog(ctx, req.(*PutLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventsService_ServiceDesc is the grpc.ServiceDesc for EventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EventsService",
	HandlerType: (*EventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _EventsService_Push_Handler,
		},
		{
			MethodName: "ReplaySingleEvent",
			Handler:    _EventsService_ReplaySingleEvent_Handler,
		},
		{
			MethodName: "PutLog",
			Handler:    _EventsService_PutLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}
