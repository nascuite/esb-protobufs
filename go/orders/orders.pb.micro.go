// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/orders.proto

package orders

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Offline service

func NewOfflineEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Offline service

type OfflineService interface {
	ByClient(ctx context.Context, in *ParamsOfflineByClient, opts ...client.CallOption) (*ResponseOfflineByClient, error)
}

type offlineService struct {
	c    client.Client
	name string
}

func NewOfflineService(name string, c client.Client) OfflineService {
	return &offlineService{
		c:    c,
		name: name,
	}
}

func (c *offlineService) ByClient(ctx context.Context, in *ParamsOfflineByClient, opts ...client.CallOption) (*ResponseOfflineByClient, error) {
	req := c.c.NewRequest(c.name, "Offline.ByClient", in)
	out := new(ResponseOfflineByClient)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Offline service

type OfflineHandler interface {
	ByClient(context.Context, *ParamsOfflineByClient, *ResponseOfflineByClient) error
}

func RegisterOfflineHandler(s server.Server, hdlr OfflineHandler, opts ...server.HandlerOption) error {
	type offline interface {
		ByClient(ctx context.Context, in *ParamsOfflineByClient, out *ResponseOfflineByClient) error
	}
	type Offline struct {
		offline
	}
	h := &offlineHandler{hdlr}
	return s.Handle(s.NewHandler(&Offline{h}, opts...))
}

type offlineHandler struct {
	OfflineHandler
}

func (h *offlineHandler) ByClient(ctx context.Context, in *ParamsOfflineByClient, out *ResponseOfflineByClient) error {
	return h.OfflineHandler.ByClient(ctx, in, out)
}

// Api Endpoints for Online service

func NewOnlineEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Online service

type OnlineService interface {
	ByClient(ctx context.Context, in *ParamsOnlineByClient, opts ...client.CallOption) (*ResponseOnlineByClient, error)
}

type onlineService struct {
	c    client.Client
	name string
}

func NewOnlineService(name string, c client.Client) OnlineService {
	return &onlineService{
		c:    c,
		name: name,
	}
}

func (c *onlineService) ByClient(ctx context.Context, in *ParamsOnlineByClient, opts ...client.CallOption) (*ResponseOnlineByClient, error) {
	req := c.c.NewRequest(c.name, "Online.ByClient", in)
	out := new(ResponseOnlineByClient)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Online service

type OnlineHandler interface {
	ByClient(context.Context, *ParamsOnlineByClient, *ResponseOnlineByClient) error
}

func RegisterOnlineHandler(s server.Server, hdlr OnlineHandler, opts ...server.HandlerOption) error {
	type online interface {
		ByClient(ctx context.Context, in *ParamsOnlineByClient, out *ResponseOnlineByClient) error
	}
	type Online struct {
		online
	}
	h := &onlineHandler{hdlr}
	return s.Handle(s.NewHandler(&Online{h}, opts...))
}

type onlineHandler struct {
	OnlineHandler
}

func (h *onlineHandler) ByClient(ctx context.Context, in *ParamsOnlineByClient, out *ResponseOnlineByClient) error {
	return h.OnlineHandler.ByClient(ctx, in, out)
}
