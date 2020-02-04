// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/logistics.proto

package logistics

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Logistics service

type LogisticsService interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*PingResponse, error)
	UPSDuration(ctx context.Context, in *GetParamsUPS, opts ...client.CallOption) (*GetResponseUPS, error)
	BoxberryDuration(ctx context.Context, in *GetParamsBoxberry, opts ...client.CallOption) (*GetResponseBoxberry, error)
	BoxberryDurationByCity(ctx context.Context, in *GetParamsBoxberryDurationByCity, opts ...client.CallOption) (*GetResponseBoxberryDurationByCity, error)
	BoxberryPointInformation(ctx context.Context, in *GetParamsPointInformation, opts ...client.CallOption) (*GetResponsePointInformation, error)
	BoxberryOrderStatus(ctx context.Context, in *GetParamsOrderStatus, opts ...client.CallOption) (*GetResponseOrderStatus, error)
}

type logisticsService struct {
	c    client.Client
	name string
}

func NewLogisticsService(name string, c client.Client) LogisticsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "logistics"
	}
	return &logisticsService{
		c:    c,
		name: name,
	}
}

func (c *logisticsService) Ping(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*PingResponse, error) {
	req := c.c.NewRequest(c.name, "Logistics.Ping", in)
	out := new(PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) UPSDuration(ctx context.Context, in *GetParamsUPS, opts ...client.CallOption) (*GetResponseUPS, error) {
	req := c.c.NewRequest(c.name, "Logistics.UPSDuration", in)
	out := new(GetResponseUPS)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) BoxberryDuration(ctx context.Context, in *GetParamsBoxberry, opts ...client.CallOption) (*GetResponseBoxberry, error) {
	req := c.c.NewRequest(c.name, "Logistics.BoxberryDuration", in)
	out := new(GetResponseBoxberry)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) BoxberryDurationByCity(ctx context.Context, in *GetParamsBoxberryDurationByCity, opts ...client.CallOption) (*GetResponseBoxberryDurationByCity, error) {
	req := c.c.NewRequest(c.name, "Logistics.BoxberryDurationByCity", in)
	out := new(GetResponseBoxberryDurationByCity)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) BoxberryPointInformation(ctx context.Context, in *GetParamsPointInformation, opts ...client.CallOption) (*GetResponsePointInformation, error) {
	req := c.c.NewRequest(c.name, "Logistics.BoxberryPointInformation", in)
	out := new(GetResponsePointInformation)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticsService) BoxberryOrderStatus(ctx context.Context, in *GetParamsOrderStatus, opts ...client.CallOption) (*GetResponseOrderStatus, error) {
	req := c.c.NewRequest(c.name, "Logistics.BoxberryOrderStatus", in)
	out := new(GetResponseOrderStatus)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logistics service

type LogisticsHandler interface {
	Ping(context.Context, *empty.Empty, *PingResponse) error
	UPSDuration(context.Context, *GetParamsUPS, *GetResponseUPS) error
	BoxberryDuration(context.Context, *GetParamsBoxberry, *GetResponseBoxberry) error
	BoxberryDurationByCity(context.Context, *GetParamsBoxberryDurationByCity, *GetResponseBoxberryDurationByCity) error
	BoxberryPointInformation(context.Context, *GetParamsPointInformation, *GetResponsePointInformation) error
	BoxberryOrderStatus(context.Context, *GetParamsOrderStatus, *GetResponseOrderStatus) error
}

func RegisterLogisticsHandler(s server.Server, hdlr LogisticsHandler, opts ...server.HandlerOption) error {
	type logistics interface {
		Ping(ctx context.Context, in *empty.Empty, out *PingResponse) error
		UPSDuration(ctx context.Context, in *GetParamsUPS, out *GetResponseUPS) error
		BoxberryDuration(ctx context.Context, in *GetParamsBoxberry, out *GetResponseBoxberry) error
		BoxberryDurationByCity(ctx context.Context, in *GetParamsBoxberryDurationByCity, out *GetResponseBoxberryDurationByCity) error
		BoxberryPointInformation(ctx context.Context, in *GetParamsPointInformation, out *GetResponsePointInformation) error
		BoxberryOrderStatus(ctx context.Context, in *GetParamsOrderStatus, out *GetResponseOrderStatus) error
	}
	type Logistics struct {
		logistics
	}
	h := &logisticsHandler{hdlr}
	return s.Handle(s.NewHandler(&Logistics{h}, opts...))
}

type logisticsHandler struct {
	LogisticsHandler
}

func (h *logisticsHandler) Ping(ctx context.Context, in *empty.Empty, out *PingResponse) error {
	return h.LogisticsHandler.Ping(ctx, in, out)
}

func (h *logisticsHandler) UPSDuration(ctx context.Context, in *GetParamsUPS, out *GetResponseUPS) error {
	return h.LogisticsHandler.UPSDuration(ctx, in, out)
}

func (h *logisticsHandler) BoxberryDuration(ctx context.Context, in *GetParamsBoxberry, out *GetResponseBoxberry) error {
	return h.LogisticsHandler.BoxberryDuration(ctx, in, out)
}

func (h *logisticsHandler) BoxberryDurationByCity(ctx context.Context, in *GetParamsBoxberryDurationByCity, out *GetResponseBoxberryDurationByCity) error {
	return h.LogisticsHandler.BoxberryDurationByCity(ctx, in, out)
}

func (h *logisticsHandler) BoxberryPointInformation(ctx context.Context, in *GetParamsPointInformation, out *GetResponsePointInformation) error {
	return h.LogisticsHandler.BoxberryPointInformation(ctx, in, out)
}

func (h *logisticsHandler) BoxberryOrderStatus(ctx context.Context, in *GetParamsOrderStatus, out *GetResponseOrderStatus) error {
	return h.LogisticsHandler.BoxberryOrderStatus(ctx, in, out)
}
