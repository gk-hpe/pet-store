// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package petstorepb

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

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreServiceClient interface {
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*AddProductResponse, error)
	ListProduct(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, "/petstorepb.StoreService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) ListProduct(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error) {
	out := new(ItemResponse)
	err := c.cc.Invoke(ctx, "/petstorepb.StoreService/ListProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
// All implementations must embed UnimplementedStoreServiceServer
// for forward compatibility
type StoreServiceServer interface {
	AddProduct(context.Context, *Product) (*AddProductResponse, error)
	ListProduct(context.Context, *ItemRequest) (*ItemResponse, error)
	mustEmbedUnimplementedStoreServiceServer()
}

// UnimplementedStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (UnimplementedStoreServiceServer) AddProduct(context.Context, *Product) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedStoreServiceServer) ListProduct(context.Context, *ItemRequest) (*ItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedStoreServiceServer) mustEmbedUnimplementedStoreServiceServer() {}

// UnsafeStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServiceServer will
// result in compilation errors.
type UnsafeStoreServiceServer interface {
	mustEmbedUnimplementedStoreServiceServer()
}

func RegisterStoreServiceServer(s grpc.ServiceRegistrar, srv StoreServiceServer) {
	s.RegisterService(&StoreService_ServiceDesc, srv)
}

func _StoreService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstorepb.StoreService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstorepb.StoreService/ListProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).ListProduct(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreService_ServiceDesc is the grpc.ServiceDesc for StoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "petstorepb.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _StoreService_AddProduct_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _StoreService_ListProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petstorepb/petstore.proto",
}
