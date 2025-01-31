// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package logistics

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

// TransportCompaniesClient is the client API for TransportCompanies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportCompaniesClient interface {
	Create(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompanyId, error)
	Get(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*TransportCompany, error)
	List(ctx context.Context, in *ListTransportCompanyRequest, opts ...grpc.CallOption) (*ListTransportCompanyResponse, error)
	Update(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error)
	Upsert(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error)
	Delete(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type transportCompaniesClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportCompaniesClient(cc grpc.ClientConnInterface) TransportCompaniesClient {
	return &transportCompaniesClient{cc}
}

func (c *transportCompaniesClient) Create(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompanyId, error) {
	out := new(TransportCompanyId)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanies/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompaniesClient) Get(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*TransportCompany, error) {
	out := new(TransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanies/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompaniesClient) List(ctx context.Context, in *ListTransportCompanyRequest, opts ...grpc.CallOption) (*ListTransportCompanyResponse, error) {
	out := new(ListTransportCompanyResponse)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanies/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompaniesClient) Update(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error) {
	out := new(TransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanies/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompaniesClient) Upsert(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error) {
	out := new(TransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanies/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompaniesClient) Delete(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanies/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportCompaniesServer is the server API for TransportCompanies service.
// All implementations should embed UnimplementedTransportCompaniesServer
// for forward compatibility
type TransportCompaniesServer interface {
	Create(context.Context, *TransportCompany) (*TransportCompanyId, error)
	Get(context.Context, *TransportCompanyId) (*TransportCompany, error)
	List(context.Context, *ListTransportCompanyRequest) (*ListTransportCompanyResponse, error)
	Update(context.Context, *TransportCompany) (*TransportCompany, error)
	Upsert(context.Context, *TransportCompany) (*TransportCompany, error)
	Delete(context.Context, *TransportCompanyId) (*emptypb.Empty, error)
}

// UnimplementedTransportCompaniesServer should be embedded to have forward compatible implementations.
type UnimplementedTransportCompaniesServer struct {
}

func (UnimplementedTransportCompaniesServer) Create(context.Context, *TransportCompany) (*TransportCompanyId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTransportCompaniesServer) Get(context.Context, *TransportCompanyId) (*TransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTransportCompaniesServer) List(context.Context, *ListTransportCompanyRequest) (*ListTransportCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTransportCompaniesServer) Update(context.Context, *TransportCompany) (*TransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTransportCompaniesServer) Upsert(context.Context, *TransportCompany) (*TransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedTransportCompaniesServer) Delete(context.Context, *TransportCompanyId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeTransportCompaniesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportCompaniesServer will
// result in compilation errors.
type UnsafeTransportCompaniesServer interface {
	mustEmbedUnimplementedTransportCompaniesServer()
}

func RegisterTransportCompaniesServer(s grpc.ServiceRegistrar, srv TransportCompaniesServer) {
	s.RegisterService(&TransportCompanies_ServiceDesc, srv)
}

func _TransportCompanies_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompaniesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanies/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompaniesServer).Create(ctx, req.(*TransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanies_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompaniesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanies/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompaniesServer).Get(ctx, req.(*TransportCompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanies_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransportCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompaniesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanies/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompaniesServer).List(ctx, req.(*ListTransportCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanies_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompaniesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanies/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompaniesServer).Update(ctx, req.(*TransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanies_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompaniesServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanies/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompaniesServer).Upsert(ctx, req.(*TransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanies_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompaniesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanies/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompaniesServer).Delete(ctx, req.(*TransportCompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportCompanies_ServiceDesc is the grpc.ServiceDesc for TransportCompanies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportCompanies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.TransportCompanies",
	HandlerType: (*TransportCompaniesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TransportCompanies_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TransportCompanies_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TransportCompanies_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransportCompanies_Update_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _TransportCompanies_Upsert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TransportCompanies_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logistics/transport_company.proto",
}
