// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderRefund/orderRefund.proto

package orderRefund

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
	_ "github.com/cargod-bj/b2c-transaction-api/orderCheckList"
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

// Api Endpoints for OrderRefund service

func NewOrderRefundEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderRefund service

type OrderRefundService interface {
	Add(ctx context.Context, in *OrderRefundDto, opts ...client.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *OrderRefundDto, opts ...client.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *OrderRefundDto, opts ...client.CallOption) (*common.Response, error)
}

type orderRefundService struct {
	c    client.Client
	name string
}

func NewOrderRefundService(name string, c client.Client) OrderRefundService {
	return &orderRefundService{
		c:    c,
		name: name,
	}
}

func (c *orderRefundService) Add(ctx context.Context, in *OrderRefundDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "OrderRefund.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderRefundService) Delete(ctx context.Context, in *OrderRefundDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "OrderRefund.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderRefundService) Update(ctx context.Context, in *OrderRefundDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "OrderRefund.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderRefund service

type OrderRefundHandler interface {
	Add(context.Context, *OrderRefundDto, *common.Response) error
	Delete(context.Context, *OrderRefundDto, *common.Response) error
	Update(context.Context, *OrderRefundDto, *common.Response) error
}

func RegisterOrderRefundHandler(s server.Server, hdlr OrderRefundHandler, opts ...server.HandlerOption) error {
	type orderRefund interface {
		Add(ctx context.Context, in *OrderRefundDto, out *common.Response) error
		Delete(ctx context.Context, in *OrderRefundDto, out *common.Response) error
		Update(ctx context.Context, in *OrderRefundDto, out *common.Response) error
	}
	type OrderRefund struct {
		orderRefund
	}
	h := &orderRefundHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderRefund{h}, opts...))
}

type orderRefundHandler struct {
	OrderRefundHandler
}

func (h *orderRefundHandler) Add(ctx context.Context, in *OrderRefundDto, out *common.Response) error {
	return h.OrderRefundHandler.Add(ctx, in, out)
}

func (h *orderRefundHandler) Delete(ctx context.Context, in *OrderRefundDto, out *common.Response) error {
	return h.OrderRefundHandler.Delete(ctx, in, out)
}

func (h *orderRefundHandler) Update(ctx context.Context, in *OrderRefundDto, out *common.Response) error {
	return h.OrderRefundHandler.Update(ctx, in, out)
}
