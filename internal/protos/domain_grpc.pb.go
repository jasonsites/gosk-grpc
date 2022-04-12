// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// DomainClient is the client API for Domain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainClient interface {
	Action(ctx context.Context, in *DomainRequest, opts ...grpc.CallOption) (*DomainResponse, error)
}

type domainClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainClient(cc grpc.ClientConnInterface) DomainClient {
	return &domainClient{cc}
}

func (c *domainClient) Action(ctx context.Context, in *DomainRequest, opts ...grpc.CallOption) (*DomainResponse, error) {
	out := new(DomainResponse)
	err := c.cc.Invoke(ctx, "/domain.Domain/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainServer is the server API for Domain service.
// All implementations must embed UnimplementedDomainServer
// for forward compatibility
type DomainServer interface {
	Action(context.Context, *DomainRequest) (*DomainResponse, error)
	mustEmbedUnimplementedDomainServer()
}

// UnimplementedDomainServer must be embedded to have forward compatible implementations.
type UnimplementedDomainServer struct {
}

func (UnimplementedDomainServer) Action(context.Context, *DomainRequest) (*DomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedDomainServer) mustEmbedUnimplementedDomainServer() {}

// UnsafeDomainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainServer will
// result in compilation errors.
type UnsafeDomainServer interface {
	mustEmbedUnimplementedDomainServer()
}

func RegisterDomainServer(s grpc.ServiceRegistrar, srv DomainServer) {
	s.RegisterService(&Domain_ServiceDesc, srv)
}

func _Domain_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.Domain/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).Action(ctx, req.(*DomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Domain_ServiceDesc is the grpc.ServiceDesc for Domain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Domain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "domain.Domain",
	HandlerType: (*DomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Action",
			Handler:    _Domain_Action_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain.proto",
}
