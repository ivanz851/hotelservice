// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: hotel.proto

package hotel

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
	HotelService_CreateHotel_FullMethodName = "/hotel.HotelService/CreateHotel"
	HotelService_GetHotel_FullMethodName    = "/hotel.HotelService/GetHotel"
	HotelService_GetHotels_FullMethodName   = "/hotel.HotelService/GetHotels"
)

// HotelServiceClient is the client API for HotelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelServiceClient interface {
	CreateHotel(ctx context.Context, in *CreateHotelRequest, opts ...grpc.CallOption) (*CreateHotelResponse, error)
	GetHotel(ctx context.Context, in *GetHotelRequest, opts ...grpc.CallOption) (*GetHotelResponse, error)
	GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (*GetHotelsResponse, error)
}

type hotelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelServiceClient(cc grpc.ClientConnInterface) HotelServiceClient {
	return &hotelServiceClient{cc}
}

func (c *hotelServiceClient) CreateHotel(ctx context.Context, in *CreateHotelRequest, opts ...grpc.CallOption) (*CreateHotelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateHotelResponse)
	err := c.cc.Invoke(ctx, HotelService_CreateHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetHotel(ctx context.Context, in *GetHotelRequest, opts ...grpc.CallOption) (*GetHotelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHotelResponse)
	err := c.cc.Invoke(ctx, HotelService_GetHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (*GetHotelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHotelsResponse)
	err := c.cc.Invoke(ctx, HotelService_GetHotels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServiceServer is the server API for HotelService service.
// All implementations must embed UnimplementedHotelServiceServer
// for forward compatibility.
type HotelServiceServer interface {
	CreateHotel(context.Context, *CreateHotelRequest) (*CreateHotelResponse, error)
	GetHotel(context.Context, *GetHotelRequest) (*GetHotelResponse, error)
	GetHotels(context.Context, *GetHotelsRequest) (*GetHotelsResponse, error)
	mustEmbedUnimplementedHotelServiceServer()
}

// UnimplementedHotelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHotelServiceServer struct{}

func (UnimplementedHotelServiceServer) CreateHotel(context.Context, *CreateHotelRequest) (*CreateHotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetHotel(context.Context, *GetHotelRequest) (*GetHotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetHotels(context.Context, *GetHotelsRequest) (*GetHotelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotels not implemented")
}
func (UnimplementedHotelServiceServer) mustEmbedUnimplementedHotelServiceServer() {}
func (UnimplementedHotelServiceServer) testEmbeddedByValue()                      {}

// UnsafeHotelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServiceServer will
// result in compilation errors.
type UnsafeHotelServiceServer interface {
	mustEmbedUnimplementedHotelServiceServer()
}

func RegisterHotelServiceServer(s grpc.ServiceRegistrar, srv HotelServiceServer) {
	// If the following call pancis, it indicates UnimplementedHotelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HotelService_ServiceDesc, srv)
}

func _HotelService_CreateHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CreateHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_CreateHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CreateHotel(ctx, req.(*CreateHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotel(ctx, req.(*GetHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetHotels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetHotels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotels(ctx, req.(*GetHotelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelService_ServiceDesc is the grpc.ServiceDesc for HotelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hotel.HotelService",
	HandlerType: (*HotelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHotel",
			Handler:    _HotelService_CreateHotel_Handler,
		},
		{
			MethodName: "GetHotel",
			Handler:    _HotelService_GetHotel_Handler,
		},
		{
			MethodName: "GetHotels",
			Handler:    _HotelService_GetHotels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotel.proto",
}
