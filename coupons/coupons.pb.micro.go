// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: coupons/coupons.proto

package coupons

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
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

// Api Endpoints for Coupons service

func NewCouponsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Coupons service

type CouponsService interface {
	Add(ctx context.Context, in *CouponDto, opts ...client.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *CouponDto, opts ...client.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *CouponDto, opts ...client.CallOption) (*common.Response, error)
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
	FindCouponByUser(ctx context.Context, in *User, opts ...client.CallOption) (*common.Response, error)
	FindCouponByNo(ctx context.Context, in *CouponNo, opts ...client.CallOption) (*common.Response, error)
}

type couponsService struct {
	c    client.Client
	name string
}

func NewCouponsService(name string, c client.Client) CouponsService {
	return &couponsService{
		c:    c,
		name: name,
	}
}

func (c *couponsService) Add(ctx context.Context, in *CouponDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Coupons.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponsService) Delete(ctx context.Context, in *CouponDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Coupons.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponsService) Update(ctx context.Context, in *CouponDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Coupons.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponsService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Coupons.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponsService) FindCouponByUser(ctx context.Context, in *User, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Coupons.FindCouponByUser", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponsService) FindCouponByNo(ctx context.Context, in *CouponNo, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Coupons.FindCouponByNo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Coupons service

type CouponsHandler interface {
	Add(context.Context, *CouponDto, *common.Response) error
	Delete(context.Context, *CouponDto, *common.Response) error
	Update(context.Context, *CouponDto, *common.Response) error
	List(context.Context, *common.Page, *common.Response) error
	FindCouponByUser(context.Context, *User, *common.Response) error
	FindCouponByNo(context.Context, *CouponNo, *common.Response) error
}

func RegisterCouponsHandler(s server.Server, hdlr CouponsHandler, opts ...server.HandlerOption) error {
	type coupons interface {
		Add(ctx context.Context, in *CouponDto, out *common.Response) error
		Delete(ctx context.Context, in *CouponDto, out *common.Response) error
		Update(ctx context.Context, in *CouponDto, out *common.Response) error
		List(ctx context.Context, in *common.Page, out *common.Response) error
		FindCouponByUser(ctx context.Context, in *User, out *common.Response) error
		FindCouponByNo(ctx context.Context, in *CouponNo, out *common.Response) error
	}
	type Coupons struct {
		coupons
	}
	h := &couponsHandler{hdlr}
	return s.Handle(s.NewHandler(&Coupons{h}, opts...))
}

type couponsHandler struct {
	CouponsHandler
}

func (h *couponsHandler) Add(ctx context.Context, in *CouponDto, out *common.Response) error {
	return h.CouponsHandler.Add(ctx, in, out)
}

func (h *couponsHandler) Delete(ctx context.Context, in *CouponDto, out *common.Response) error {
	return h.CouponsHandler.Delete(ctx, in, out)
}

func (h *couponsHandler) Update(ctx context.Context, in *CouponDto, out *common.Response) error {
	return h.CouponsHandler.Update(ctx, in, out)
}

func (h *couponsHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.CouponsHandler.List(ctx, in, out)
}

func (h *couponsHandler) FindCouponByUser(ctx context.Context, in *User, out *common.Response) error {
	return h.CouponsHandler.FindCouponByUser(ctx, in, out)
}

func (h *couponsHandler) FindCouponByNo(ctx context.Context, in *CouponNo, out *common.Response) error {
	return h.CouponsHandler.FindCouponByNo(ctx, in, out)
}
