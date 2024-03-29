// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: order.proto

package orderproto

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

const (
	OrderProto_CreateOrder_FullMethodName = "/OrderProto/CreateOrder"
	OrderProto_GetOrders_FullMethodName   = "/OrderProto/GetOrders"
)

// OrderProtoClient is the client API for OrderProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderProtoClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
}

type orderProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderProtoClient(cc grpc.ClientConnInterface) OrderProtoClient {
	return &orderProtoClient{cc}
}

func (c *orderProtoClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderProto_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderProtoClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, OrderProto_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderProtoServer is the server API for OrderProto service.
// All implementations must embed UnimplementedOrderProtoServer
// for forward compatibility
type OrderProtoServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	mustEmbedUnimplementedOrderProtoServer()
}

// UnimplementedOrderProtoServer must be embedded to have forward compatible implementations.
type UnimplementedOrderProtoServer struct {
}

func (UnimplementedOrderProtoServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderProtoServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderProtoServer) mustEmbedUnimplementedOrderProtoServer() {}

// UnsafeOrderProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderProtoServer will
// result in compilation errors.
type UnsafeOrderProtoServer interface {
	mustEmbedUnimplementedOrderProtoServer()
}

func RegisterOrderProtoServer(s grpc.ServiceRegistrar, srv OrderProtoServer) {
	s.RegisterService(&OrderProto_ServiceDesc, srv)
}

func _OrderProto_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProtoServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProto_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProtoServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderProto_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderProtoServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderProto_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderProtoServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderProto_ServiceDesc is the grpc.ServiceDesc for OrderProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderProto",
	HandlerType: (*OrderProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderProto_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrderProto_GetOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
