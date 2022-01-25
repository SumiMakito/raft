// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// TransportClient is the client API for Transport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportClient interface {
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	InstallSnapshot(ctx context.Context, in *InstallSnapshotRequest, opts ...grpc.CallOption) (*InstallSnapshotResponse, error)
	ApplyLog(ctx context.Context, in *ApplyLogRequest, opts ...grpc.CallOption) (*ApplyLogResponse, error)
}

type transportClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportClient(cc grpc.ClientConnInterface) TransportClient {
	return &transportClient{cc}
}

func (c *transportClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := c.cc.Invoke(ctx, "/pb.Transport/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/pb.Transport/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportClient) InstallSnapshot(ctx context.Context, in *InstallSnapshotRequest, opts ...grpc.CallOption) (*InstallSnapshotResponse, error) {
	out := new(InstallSnapshotResponse)
	err := c.cc.Invoke(ctx, "/pb.Transport/InstallSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportClient) ApplyLog(ctx context.Context, in *ApplyLogRequest, opts ...grpc.CallOption) (*ApplyLogResponse, error) {
	out := new(ApplyLogResponse)
	err := c.cc.Invoke(ctx, "/pb.Transport/ApplyLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServer is the server API for Transport service.
// All implementations must embed UnimplementedTransportServer
// for forward compatibility
type TransportServer interface {
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	InstallSnapshot(context.Context, *InstallSnapshotRequest) (*InstallSnapshotResponse, error)
	ApplyLog(context.Context, *ApplyLogRequest) (*ApplyLogResponse, error)
	mustEmbedUnimplementedTransportServer()
}

// UnimplementedTransportServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServer struct {
}

func (UnimplementedTransportServer) AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (UnimplementedTransportServer) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedTransportServer) InstallSnapshot(context.Context, *InstallSnapshotRequest) (*InstallSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallSnapshot not implemented")
}
func (UnimplementedTransportServer) ApplyLog(context.Context, *ApplyLogRequest) (*ApplyLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyLog not implemented")
}
func (UnimplementedTransportServer) mustEmbedUnimplementedTransportServer() {}

// UnsafeTransportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServer will
// result in compilation errors.
type UnsafeTransportServer interface {
	mustEmbedUnimplementedTransportServer()
}

func RegisterTransportServer(s grpc.ServiceRegistrar, srv TransportServer) {
	s.RegisterService(&Transport_ServiceDesc, srv)
}

func _Transport_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transport/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transport_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transport/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transport_InstallSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServer).InstallSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transport/InstallSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServer).InstallSnapshot(ctx, req.(*InstallSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transport_ApplyLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServer).ApplyLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transport/ApplyLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServer).ApplyLog(ctx, req.(*ApplyLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transport_ServiceDesc is the grpc.ServiceDesc for Transport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Transport",
	HandlerType: (*TransportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntries",
			Handler:    _Transport_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _Transport_RequestVote_Handler,
		},
		{
			MethodName: "InstallSnapshot",
			Handler:    _Transport_InstallSnapshot_Handler,
		},
		{
			MethodName: "ApplyLog",
			Handler:    _Transport_ApplyLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}
