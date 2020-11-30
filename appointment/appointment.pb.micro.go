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
	AssignAppointment(ctx context.Context, in *AssignDto, opts ...client.CallOption) (*Response, error)
	//个人中心查询预约
	GetUserList(ctx context.Context, in *AppointmentCon, opts ...client.CallOption) (*Response, error)
	//单个车查询预约
	Get(ctx context.Context, in *AppointCon, opts ...client.CallOption) (*Response, error)
	//客户管理页面查询
	GetListForManage(ctx context.Context, in *ManageCond, opts ...client.CallOption) (*Response, error)
	//appointment_list界面接口
	GetAppointmentList(ctx context.Context, in *AppointmentCondition, opts ...client.CallOption) (*Response, error)
	GetListForSMS(ctx context.Context, in *SmsCond, opts ...client.CallOption) (*Response, error)
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

func (c *appointmentService) AssignAppointment(ctx context.Context, in *AssignDto, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.AssignAppointment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) GetUserList(ctx context.Context, in *AppointmentCon, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.GetUserList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) Get(ctx context.Context, in *AppointCon, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) GetListForManage(ctx context.Context, in *ManageCond, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.GetListForManage", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) GetAppointmentList(ctx context.Context, in *AppointmentCondition, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.GetAppointmentList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentService) GetListForSMS(ctx context.Context, in *SmsCond, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appointment.GetListForSMS", in)
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
	AssignAppointment(context.Context, *AssignDto, *Response) error
	//个人中心查询预约
	GetUserList(context.Context, *AppointmentCon, *Response) error
	//单个车查询预约
	Get(context.Context, *AppointCon, *Response) error
	//客户管理页面查询
	GetListForManage(context.Context, *ManageCond, *Response) error
	//appointment_list界面接口
	GetAppointmentList(context.Context, *AppointmentCondition, *Response) error
	GetListForSMS(context.Context, *SmsCond, *Response) error
}

func RegisterAppointmentHandler(s server.Server, hdlr AppointmentHandler, opts ...server.HandlerOption) error {
	type appointment interface {
		Add(ctx context.Context, in *AppointmentDTO, out *Response) error
		Delete(ctx context.Context, in *DeleteId, out *Response) error
		Update(ctx context.Context, in *AppointmentDTO, out *Response) error
		GetList(ctx context.Context, in *AppointmentCondition, out *Response) error
		AssignAppointment(ctx context.Context, in *AssignDto, out *Response) error
		GetUserList(ctx context.Context, in *AppointmentCon, out *Response) error
		Get(ctx context.Context, in *AppointCon, out *Response) error
		GetListForManage(ctx context.Context, in *ManageCond, out *Response) error
		GetAppointmentList(ctx context.Context, in *AppointmentCondition, out *Response) error
		GetListForSMS(ctx context.Context, in *SmsCond, out *Response) error
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

func (h *appointmentHandler) AssignAppointment(ctx context.Context, in *AssignDto, out *Response) error {
	return h.AppointmentHandler.AssignAppointment(ctx, in, out)
}

func (h *appointmentHandler) GetUserList(ctx context.Context, in *AppointmentCon, out *Response) error {
	return h.AppointmentHandler.GetUserList(ctx, in, out)
}

func (h *appointmentHandler) Get(ctx context.Context, in *AppointCon, out *Response) error {
	return h.AppointmentHandler.Get(ctx, in, out)
}

func (h *appointmentHandler) GetListForManage(ctx context.Context, in *ManageCond, out *Response) error {
	return h.AppointmentHandler.GetListForManage(ctx, in, out)
}

func (h *appointmentHandler) GetAppointmentList(ctx context.Context, in *AppointmentCondition, out *Response) error {
	return h.AppointmentHandler.GetAppointmentList(ctx, in, out)
}

func (h *appointmentHandler) GetListForSMS(ctx context.Context, in *SmsCond, out *Response) error {
	return h.AppointmentHandler.GetListForSMS(ctx, in, out)
}
