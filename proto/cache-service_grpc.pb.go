// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/cache-service.proto

package gen

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

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheServiceClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/service.CacheService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/service.CacheService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
// All implementations must embed UnimplementedCacheServiceServer
// for forward compatibility
type CacheServiceServer interface {
	Set(context.Context, *SetRequest) (*ServerResponse, error)
	Get(context.Context, *GetRequest) (*ServerResponse, error)
	mustEmbedUnimplementedCacheServiceServer()
}

// UnimplementedCacheServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (UnimplementedCacheServiceServer) Set(context.Context, *SetRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacheServiceServer) Get(context.Context, *GetRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacheServiceServer) mustEmbedUnimplementedCacheServiceServer() {}

// UnsafeCacheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServiceServer will
// result in compilation errors.
type UnsafeCacheServiceServer interface {
	mustEmbedUnimplementedCacheServiceServer()
}

func RegisterCacheServiceServer(s grpc.ServiceRegistrar, srv CacheServiceServer) {
	s.RegisterService(&CacheService_ServiceDesc, srv)
}

func _CacheService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CacheService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CacheService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheService_ServiceDesc is the grpc.ServiceDesc for CacheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _CacheService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cache-service.proto",
}
