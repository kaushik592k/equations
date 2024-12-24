// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: equation.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChemicalBalancer_BalanceEquation_FullMethodName = "/protobuf.ChemicalBalancer/BalanceEquation"
)

// ChemicalBalancerClient is the client API for ChemicalBalancer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChemicalBalancerClient interface {
	BalanceEquation(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type chemicalBalancerClient struct {
	cc grpc.ClientConnInterface
}

func NewChemicalBalancerClient(cc grpc.ClientConnInterface) ChemicalBalancerClient {
	return &chemicalBalancerClient{cc}
}

func (c *chemicalBalancerClient) BalanceEquation(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, ChemicalBalancer_BalanceEquation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChemicalBalancerServer is the server API for ChemicalBalancer service.
// All implementations must embed UnimplementedChemicalBalancerServer
// for forward compatibility.
type ChemicalBalancerServer interface {
	BalanceEquation(context.Context, *BalanceRequest) (*BalanceResponse, error)
	mustEmbedUnimplementedChemicalBalancerServer()
}

// UnimplementedChemicalBalancerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChemicalBalancerServer struct{}

func (UnimplementedChemicalBalancerServer) BalanceEquation(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceEquation not implemented")
}
func (UnimplementedChemicalBalancerServer) mustEmbedUnimplementedChemicalBalancerServer() {}
func (UnimplementedChemicalBalancerServer) testEmbeddedByValue()                          {}

// UnsafeChemicalBalancerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChemicalBalancerServer will
// result in compilation errors.
type UnsafeChemicalBalancerServer interface {
	mustEmbedUnimplementedChemicalBalancerServer()
}

func RegisterChemicalBalancerServer(s grpc.ServiceRegistrar, srv ChemicalBalancerServer) {
	// If the following call pancis, it indicates UnimplementedChemicalBalancerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChemicalBalancer_ServiceDesc, srv)
}

func _ChemicalBalancer_BalanceEquation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChemicalBalancerServer).BalanceEquation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChemicalBalancer_BalanceEquation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChemicalBalancerServer).BalanceEquation(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChemicalBalancer_ServiceDesc is the grpc.ServiceDesc for ChemicalBalancer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChemicalBalancer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.ChemicalBalancer",
	HandlerType: (*ChemicalBalancerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BalanceEquation",
			Handler:    _ChemicalBalancer_BalanceEquation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "equation.proto",
}
