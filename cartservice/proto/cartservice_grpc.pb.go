// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.2
// source: proto/cartservice.proto

package microshopping

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CartService_AddItem_FullMethodName   = "/microshopping.CartService/AddItem"
	CartService_GetCart_FullMethodName   = "/microshopping.CartService/GetCart"
	CartService_EmptyCart_FullMethodName = "/microshopping.CartService/EmptyCart"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*Empty, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*Cart, error)
	EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*Empty, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, CartService_AddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*Cart, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Cart)
	err := c.cc.Invoke(ctx, CartService_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, CartService_EmptyCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations should embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	AddItem(context.Context, *AddItemRequest) (*Empty, error)
	GetCart(context.Context, *GetCartRequest) (*Cart, error)
	EmptyCart(context.Context, *EmptyCartRequest) (*Empty, error)
}

// UnimplementedCartServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) AddItem(context.Context, *AddItemRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedCartServiceServer) GetCart(context.Context, *GetCartRequest) (*Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) EmptyCart(context.Context, *EmptyCartRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCart not implemented")
}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_AddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_EmptyCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).EmptyCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_EmptyCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).EmptyCart(ctx, req.(*EmptyCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microshopping.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _CartService_AddItem_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "EmptyCart",
			Handler:    _CartService_EmptyCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cartservice.proto",
}

const (
	ProductCatalogService_ListProducts_FullMethodName   = "/microshopping.ProductCatalogService/ListProducts"
	ProductCatalogService_GetProduct_FullMethodName     = "/microshopping.ProductCatalogService/GetProduct"
	ProductCatalogService_SearchProducts_FullMethodName = "/microshopping.ProductCatalogService/SearchProducts"
)

// ProductCatalogServiceClient is the client API for ProductCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCatalogServiceClient interface {
	ListProducts(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
}

type productCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogServiceClient(cc grpc.ClientConnInterface) ProductCatalogServiceClient {
	return &productCatalogServiceClient{cc}
}

func (c *productCatalogServiceClient) ListProducts(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, ProductCatalogService_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductCatalogService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, ProductCatalogService_SearchProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServiceServer is the server API for ProductCatalogService service.
// All implementations should embed UnimplementedProductCatalogServiceServer
// for forward compatibility
type ProductCatalogServiceServer interface {
	ListProducts(context.Context, *PageRequest) (*ListProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
}

// UnimplementedProductCatalogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductCatalogServiceServer struct {
}

func (UnimplementedProductCatalogServiceServer) ListProducts(context.Context, *PageRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductCatalogServiceServer) GetProduct(context.Context, *GetProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}

// UnsafeProductCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCatalogServiceServer will
// result in compilation errors.
type UnsafeProductCatalogServiceServer interface {
	mustEmbedUnimplementedProductCatalogServiceServer()
}

func RegisterProductCatalogServiceServer(s grpc.ServiceRegistrar, srv ProductCatalogServiceServer) {
	s.RegisterService(&ProductCatalogService_ServiceDesc, srv)
}

func _ProductCatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_SearchProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCatalogService_ServiceDesc is the grpc.ServiceDesc for ProductCatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microshopping.ProductCatalogService",
	HandlerType: (*ProductCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _ProductCatalogService_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductCatalogService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductCatalogService_SearchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cartservice.proto",
}
