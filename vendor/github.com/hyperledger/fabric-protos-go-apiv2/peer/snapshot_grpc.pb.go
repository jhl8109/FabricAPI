// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: peer/snapshot.proto

package peer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SnapshotClient is the client API for Snapshot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotClient interface {
	// Generate a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Generate(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Cancel a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Cancel(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Query pending snapshots query. SignedSnapshotRequest contains marshalled bytes for SnaphostQuery
	QueryPendings(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*QueryPendingSnapshotsResponse, error)
}

type snapshotClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotClient(cc grpc.ClientConnInterface) SnapshotClient {
	return &snapshotClient{cc}
}

func (c *snapshotClient) Generate(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protos.Snapshot/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotClient) Cancel(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protos.Snapshot/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotClient) QueryPendings(ctx context.Context, in *SignedSnapshotRequest, opts ...grpc.CallOption) (*QueryPendingSnapshotsResponse, error) {
	out := new(QueryPendingSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/protos.Snapshot/QueryPendings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotServer is the server API for Snapshot service.
// All implementations should embed UnimplementedSnapshotServer
// for forward compatibility
type SnapshotServer interface {
	// Generate a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Generate(context.Context, *SignedSnapshotRequest) (*emptypb.Empty, error)
	// Cancel a snapshot reqeust. SignedSnapshotRequest contains marshalled bytes for SnaphostRequest
	Cancel(context.Context, *SignedSnapshotRequest) (*emptypb.Empty, error)
	// Query pending snapshots query. SignedSnapshotRequest contains marshalled bytes for SnaphostQuery
	QueryPendings(context.Context, *SignedSnapshotRequest) (*QueryPendingSnapshotsResponse, error)
}

// UnimplementedSnapshotServer should be embedded to have forward compatible implementations.
type UnimplementedSnapshotServer struct {
}

func (UnimplementedSnapshotServer) Generate(context.Context, *SignedSnapshotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedSnapshotServer) Cancel(context.Context, *SignedSnapshotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedSnapshotServer) QueryPendings(context.Context, *SignedSnapshotRequest) (*QueryPendingSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPendings not implemented")
}

// UnsafeSnapshotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotServer will
// result in compilation errors.
type UnsafeSnapshotServer interface {
	mustEmbedUnimplementedSnapshotServer()
}

func RegisterSnapshotServer(s grpc.ServiceRegistrar, srv SnapshotServer) {
	s.RegisterService(&Snapshot_ServiceDesc, srv)
}

func _Snapshot_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Snapshot/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServer).Generate(ctx, req.(*SignedSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshot_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Snapshot/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServer).Cancel(ctx, req.(*SignedSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshot_QueryPendings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServer).QueryPendings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Snapshot/QueryPendings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServer).QueryPendings(ctx, req.(*SignedSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Snapshot_ServiceDesc is the grpc.ServiceDesc for Snapshot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Snapshot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Snapshot",
	HandlerType: (*SnapshotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Snapshot_Generate_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Snapshot_Cancel_Handler,
		},
		{
			MethodName: "QueryPendings",
			Handler:    _Snapshot_QueryPendings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/snapshot.proto",
}
