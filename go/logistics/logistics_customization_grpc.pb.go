// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package logistics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CustomizationClient is the client API for Customization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomizationClient interface {
	GetZones(ctx context.Context, in *GetZonesParams, opts ...grpc.CallOption) (*GetZonesResponse, error)
}

type customizationClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomizationClient(cc grpc.ClientConnInterface) CustomizationClient {
	return &customizationClient{cc}
}

func (c *customizationClient) GetZones(ctx context.Context, in *GetZonesParams, opts ...grpc.CallOption) (*GetZonesResponse, error) {
	out := new(GetZonesResponse)
	err := c.cc.Invoke(ctx, "/logistics.customization/GetZones", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomizationServer is the server API for Customization service.
// All implementations must embed UnimplementedCustomizationServer
// for forward compatibility
type CustomizationServer interface {
	GetZones(context.Context, *GetZonesParams) (*GetZonesResponse, error)
	mustEmbedUnimplementedCustomizationServer()
}

// UnimplementedCustomizationServer must be embedded to have forward compatible implementations.
type UnimplementedCustomizationServer struct {
}

func (UnimplementedCustomizationServer) GetZones(context.Context, *GetZonesParams) (*GetZonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZones not implemented")
}
func (UnimplementedCustomizationServer) mustEmbedUnimplementedCustomizationServer() {}

// UnsafeCustomizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomizationServer will
// result in compilation errors.
type UnsafeCustomizationServer interface {
	mustEmbedUnimplementedCustomizationServer()
}

func RegisterCustomizationServer(s grpc.ServiceRegistrar, srv CustomizationServer) {
	s.RegisterService(&_Customization_serviceDesc, srv)
}

func _Customization_GetZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetZonesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomizationServer).GetZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.customization/GetZones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomizationServer).GetZones(ctx, req.(*GetZonesParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.customization",
	HandlerType: (*CustomizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetZones",
			Handler:    _Customization_GetZones_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logistics_customization.proto",
}
