// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: appointment/appointment.proto

package appointment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

// Api Endpoints for Appointment service

func NewAppointmentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Appointment service

type AppointmentService interface {
	Add(ctx context.Context, in *AppointmentDTO, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *AppointmentDTO, opts ...client.CallOption) (*Response, error)
	GetList(ctx context.Context, in *AppointmentCondition, opts ...client.CallOption) (*Response, error)
}

type appointmentService struct {
	c    client.Client
	name string
}

func NewAppointmentService(name string, c client.Client) AppointmentService {
	return &appointmentService{
		c:    c,
		name: name,
	}
}

func (c *appointmentService) Add(ctx context.Context, in *AppointmentDTO, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.Add", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) Update(ctx context.Context, in *AppointmentDTO, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) GetList(ctx context.Context, in *AppointmentCondition, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.GetList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Appointment service

type AppointmentHandler interface {
	Add(context.Context, *AppointmentDTO, *Response) error
	Delete(context.Context, *DeleteId, *Response) error
	Update(context.Context, *AppointmentDTO, *Response) error
	GetList(context.Context, *AppointmentCondition, *Response) error
}

func RegisterAppointmentHandler(s server.Server, hdlr AppointmentHandler, opts ...server.HandlerOption) error {
	type appointment interface {
		Add(ctx context.Context, in *AppointmentDTO, out *Response) error
		Delete(ctx context.Context, in *DeleteId, out *Response) error
		Update(ctx context.Context, in *AppointmentDTO, out *Response) error
		GetList(ctx context.Context, in *AppointmentCondition, out *Response) error
	}
	type Appointment struct {
		appointment
	}
	h := &appointmentHandler{hdlr}
	return s.Handle(s.NewHandler(&Appointment{h}, opts...))
}

type appointmentHandler struct {
	AppointmentHandler
}

func (h *appointmentHandler) Add(ctx context.Context, in *AppointmentDTO, out *Response) error {
	return h.AppointmentHandler.Add(ctx, in, out)
}

func (h *appointmentHandler) Delete(ctx context.Context, in *DeleteId, out *Response) error {
	return h.AppointmentHandler.Delete(ctx, in, out)
}

func (h *appointmentHandler) Update(ctx context.Context, in *AppointmentDTO, out *Response) error {
	return h.AppointmentHandler.Update(ctx, in, out)
}

func (h *appointmentHandler) GetList(ctx context.Context, in *AppointmentCondition, out *Response) error {
	return h.AppointmentHandler.GetList(ctx, in, out)
}
