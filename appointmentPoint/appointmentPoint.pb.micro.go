// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: appointmentPoint/appointmentPoint.proto

package appointmentPoint

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
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

// Api Endpoints for AppointmentPoint service

func NewAppointmentPointEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppointmentPoint service

type AppointmentPointService interface {
	Add(ctx context.Context, in *AppointmentPointDTO, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error)
}

type appointmentPointService struct {
	c    client.Client
	name string
}

func NewAppointmentPointService(name string, c client.Client) AppointmentPointService {
	return &appointmentPointService{
		c:    c,
		name: name,
	}
}

func (c *appointmentPointService) Add(ctx context.Context, in *AppointmentPointDTO, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppointmentPoint.Add", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentPointService) Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppointmentPoint.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppointmentPoint service

type AppointmentPointHandler interface {
	Add(context.Context, *AppointmentPointDTO, *Response) error
	Delete(context.Context, *DeleteId, *Response) error
}

func RegisterAppointmentPointHandler(s server.Server, hdlr AppointmentPointHandler, opts ...server.HandlerOption) error {
	type appointmentPoint interface {
		Add(ctx context.Context, in *AppointmentPointDTO, out *Response) error
		Delete(ctx context.Context, in *DeleteId, out *Response) error
	}
	type AppointmentPoint struct {
		appointmentPoint
	}
	h := &appointmentPointHandler{hdlr}
	return s.Handle(s.NewHandler(&AppointmentPoint{h}, opts...))
}

type appointmentPointHandler struct {
	AppointmentPointHandler
}

func (h *appointmentPointHandler) Add(ctx context.Context, in *AppointmentPointDTO, out *Response) error {
	return h.AppointmentPointHandler.Add(ctx, in, out)
}

func (h *appointmentPointHandler) Delete(ctx context.Context, in *DeleteId, out *Response) error {
	return h.AppointmentPointHandler.Delete(ctx, in, out)
}
