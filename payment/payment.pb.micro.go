// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: payment/payment.proto

package payment

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
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

// Api Endpoints for Payment service

func NewPaymentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Payment service

type PaymentService interface {
	Add(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error)
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
	//根据订单号获取支付信息 （orderID必填字段）
	GetPaymentListByCond(ctx context.Context, in *PaymentCond, opts ...client.CallOption) (*common.Response, error)
	//新增照片信息
	UploadImages(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) Add(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Payment.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) Delete(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Payment.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) Update(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Payment.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Payment.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) GetPaymentListByCond(ctx context.Context, in *PaymentCond, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Payment.GetPaymentListByCond", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) UploadImages(ctx context.Context, in *PaymentDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Payment.UploadImages", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Payment service

type PaymentHandler interface {
	Add(context.Context, *PaymentDto, *common.Response) error
	Delete(context.Context, *PaymentDto, *common.Response) error
	Update(context.Context, *PaymentDto, *common.Response) error
	List(context.Context, *common.Page, *common.Response) error
	//根据订单号获取支付信息 （orderID必填字段）
	GetPaymentListByCond(context.Context, *PaymentCond, *common.Response) error
	//新增照片信息
	UploadImages(context.Context, *PaymentDto, *common.Response) error
}

func RegisterPaymentHandler(s server.Server, hdlr PaymentHandler, opts ...server.HandlerOption) error {
	type payment interface {
		Add(ctx context.Context, in *PaymentDto, out *common.Response) error
		Delete(ctx context.Context, in *PaymentDto, out *common.Response) error
		Update(ctx context.Context, in *PaymentDto, out *common.Response) error
		List(ctx context.Context, in *common.Page, out *common.Response) error
		GetPaymentListByCond(ctx context.Context, in *PaymentCond, out *common.Response) error
		UploadImages(ctx context.Context, in *PaymentDto, out *common.Response) error
	}
	type Payment struct {
		payment
	}
	h := &paymentHandler{hdlr}
	return s.Handle(s.NewHandler(&Payment{h}, opts...))
}

type paymentHandler struct {
	PaymentHandler
}

func (h *paymentHandler) Add(ctx context.Context, in *PaymentDto, out *common.Response) error {
	return h.PaymentHandler.Add(ctx, in, out)
}

func (h *paymentHandler) Delete(ctx context.Context, in *PaymentDto, out *common.Response) error {
	return h.PaymentHandler.Delete(ctx, in, out)
}

func (h *paymentHandler) Update(ctx context.Context, in *PaymentDto, out *common.Response) error {
	return h.PaymentHandler.Update(ctx, in, out)
}

func (h *paymentHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.PaymentHandler.List(ctx, in, out)
}

func (h *paymentHandler) GetPaymentListByCond(ctx context.Context, in *PaymentCond, out *common.Response) error {
	return h.PaymentHandler.GetPaymentListByCond(ctx, in, out)
}

func (h *paymentHandler) UploadImages(ctx context.Context, in *PaymentDto, out *common.Response) error {
	return h.PaymentHandler.UploadImages(ctx, in, out)
}
