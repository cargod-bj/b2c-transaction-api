// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: order/order.proto

package order

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
	_ "github.com/cargod-bj/b2c-transaction-api/orderCheckList"
	_ "github.com/cargod-bj/b2c-transaction-api/orderFee"
	_ "github.com/cargod-bj/b2c-transaction-api/orderFlowLog"
	_ "github.com/cargod-bj/b2c-transaction-api/orderRefund"
	_ "github.com/cargod-bj/b2c-transaction-api/payment"
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

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	//新增订单，返回data.nil
	Add(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*Response, error)
	//根据id删除订单，这里是假删除，设置status为2，返回data.nil
	Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error)
	//更新订单信息，返回data.nil
	Update(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*Response, error)
	//查询订单列表，返回订单列表数据
	GetList(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*Response, error)
	//查询订单详情
	GetDetail(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*Response, error)
	//取消订单
	CancelOrder(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*Response, error)
	//保存快递信息
	SaveDeliveryInfo(ctx context.Context, in *DeliveryInfo, opts ...client.CallOption) (*Response, error)
	//更新订单状态
	UpdateOrderStatus(ctx context.Context, in *UpdateInfo, opts ...client.CallOption) (*common.Response, error)
	//更新订单费用信息
	UpdateCost(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*common.Response, error)
	// change car
	ChangeCar(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*common.Response, error)
	//查询车辆是否已经创建订单
	CheckCarInvalidInOrder(ctx context.Context, in *CarIds, opts ...client.CallOption) (*common.Response, error)
	//查询要发短信的订单
	QueryOrderForSendSms(ctx context.Context, in *DeliveryInfo, opts ...client.CallOption) (*common.Response, error)
	//查询线上预定订单并且pic为空的订单
	GetNoAssignCustomerByOrder(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error)
	//查询订单列表，返回订单列表数据
	GetListWithPay(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error)
	//给没有分配PIC订单，分配PIC
	AssignOrderPIC(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error)
	CancelOnlineOrder(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error)
	GetPayInfo(ctx context.Context, in *QueryPayDTO, opts ...client.CallOption) (*common.Response, error)
	CheckCarInvalid(ctx context.Context, in *CarCheckDTO, opts ...client.CallOption) (*common.Response, error)
	//根据carid查询订单
	Get(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error)
	GetByOrderNo(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Add(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.Add", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Update(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetList(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.GetList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetDetail(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.GetDetail", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CancelOrder(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.CancelOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) SaveDeliveryInfo(ctx context.Context, in *DeliveryInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.SaveDeliveryInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrderStatus(ctx context.Context, in *UpdateInfo, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateOrderStatus", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateCost(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateCost", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) ChangeCar(ctx context.Context, in *OrderDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.ChangeCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CheckCarInvalidInOrder(ctx context.Context, in *CarIds, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.CheckCarInvalidInOrder", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) QueryOrderForSendSms(ctx context.Context, in *DeliveryInfo, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.QueryOrderForSendSms", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetNoAssignCustomerByOrder(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.GetNoAssignCustomerByOrder", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetListWithPay(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.GetListWithPay", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AssignOrderPIC(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.AssignOrderPIC", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CancelOnlineOrder(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.CancelOnlineOrder", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetPayInfo(ctx context.Context, in *QueryPayDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.GetPayInfo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CheckCarInvalid(ctx context.Context, in *CarCheckDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.CheckCarInvalid", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Get(ctx context.Context, in *OrderCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetByOrderNo(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Order.GetByOrderNo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	//新增订单，返回data.nil
	Add(context.Context, *OrderDTO, *Response) error
	//根据id删除订单，这里是假删除，设置status为2，返回data.nil
	Delete(context.Context, *DeleteId, *Response) error
	//更新订单信息，返回data.nil
	Update(context.Context, *OrderDTO, *Response) error
	//查询订单列表，返回订单列表数据
	GetList(context.Context, *OrderCondition, *Response) error
	//查询订单详情
	GetDetail(context.Context, *OrderCondition, *Response) error
	//取消订单
	CancelOrder(context.Context, *OrderCondition, *Response) error
	//保存快递信息
	SaveDeliveryInfo(context.Context, *DeliveryInfo, *Response) error
	//更新订单状态
	UpdateOrderStatus(context.Context, *UpdateInfo, *common.Response) error
	//更新订单费用信息
	UpdateCost(context.Context, *OrderDTO, *common.Response) error
	// change car
	ChangeCar(context.Context, *OrderDTO, *common.Response) error
	//查询车辆是否已经创建订单
	CheckCarInvalidInOrder(context.Context, *CarIds, *common.Response) error
	//查询要发短信的订单
	QueryOrderForSendSms(context.Context, *DeliveryInfo, *common.Response) error
	//查询线上预定订单并且pic为空的订单
	GetNoAssignCustomerByOrder(context.Context, *OrderCondition, *common.Response) error
	//查询订单列表，返回订单列表数据
	GetListWithPay(context.Context, *OrderCondition, *common.Response) error
	//给没有分配PIC订单，分配PIC
	AssignOrderPIC(context.Context, *OrderCondition, *common.Response) error
	CancelOnlineOrder(context.Context, *OrderCondition, *common.Response) error
	GetPayInfo(context.Context, *QueryPayDTO, *common.Response) error
	CheckCarInvalid(context.Context, *CarCheckDTO, *common.Response) error
	//根据carid查询订单
	Get(context.Context, *OrderCondition, *common.Response) error
	GetByOrderNo(context.Context, *common.IdDto, *common.Response) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		Add(ctx context.Context, in *OrderDTO, out *Response) error
		Delete(ctx context.Context, in *DeleteId, out *Response) error
		Update(ctx context.Context, in *OrderDTO, out *Response) error
		GetList(ctx context.Context, in *OrderCondition, out *Response) error
		GetDetail(ctx context.Context, in *OrderCondition, out *Response) error
		CancelOrder(ctx context.Context, in *OrderCondition, out *Response) error
		SaveDeliveryInfo(ctx context.Context, in *DeliveryInfo, out *Response) error
		UpdateOrderStatus(ctx context.Context, in *UpdateInfo, out *common.Response) error
		UpdateCost(ctx context.Context, in *OrderDTO, out *common.Response) error
		ChangeCar(ctx context.Context, in *OrderDTO, out *common.Response) error
		CheckCarInvalidInOrder(ctx context.Context, in *CarIds, out *common.Response) error
		QueryOrderForSendSms(ctx context.Context, in *DeliveryInfo, out *common.Response) error
		GetNoAssignCustomerByOrder(ctx context.Context, in *OrderCondition, out *common.Response) error
		GetListWithPay(ctx context.Context, in *OrderCondition, out *common.Response) error
		AssignOrderPIC(ctx context.Context, in *OrderCondition, out *common.Response) error
		CancelOnlineOrder(ctx context.Context, in *OrderCondition, out *common.Response) error
		GetPayInfo(ctx context.Context, in *QueryPayDTO, out *common.Response) error
		CheckCarInvalid(ctx context.Context, in *CarCheckDTO, out *common.Response) error
		Get(ctx context.Context, in *OrderCondition, out *common.Response) error
		GetByOrderNo(ctx context.Context, in *common.IdDto, out *common.Response) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) Add(ctx context.Context, in *OrderDTO, out *Response) error {
	return h.OrderHandler.Add(ctx, in, out)
}

func (h *orderHandler) Delete(ctx context.Context, in *DeleteId, out *Response) error {
	return h.OrderHandler.Delete(ctx, in, out)
}

func (h *orderHandler) Update(ctx context.Context, in *OrderDTO, out *Response) error {
	return h.OrderHandler.Update(ctx, in, out)
}

func (h *orderHandler) GetList(ctx context.Context, in *OrderCondition, out *Response) error {
	return h.OrderHandler.GetList(ctx, in, out)
}

func (h *orderHandler) GetDetail(ctx context.Context, in *OrderCondition, out *Response) error {
	return h.OrderHandler.GetDetail(ctx, in, out)
}

func (h *orderHandler) CancelOrder(ctx context.Context, in *OrderCondition, out *Response) error {
	return h.OrderHandler.CancelOrder(ctx, in, out)
}

func (h *orderHandler) SaveDeliveryInfo(ctx context.Context, in *DeliveryInfo, out *Response) error {
	return h.OrderHandler.SaveDeliveryInfo(ctx, in, out)
}

func (h *orderHandler) UpdateOrderStatus(ctx context.Context, in *UpdateInfo, out *common.Response) error {
	return h.OrderHandler.UpdateOrderStatus(ctx, in, out)
}

func (h *orderHandler) UpdateCost(ctx context.Context, in *OrderDTO, out *common.Response) error {
	return h.OrderHandler.UpdateCost(ctx, in, out)
}

func (h *orderHandler) ChangeCar(ctx context.Context, in *OrderDTO, out *common.Response) error {
	return h.OrderHandler.ChangeCar(ctx, in, out)
}

func (h *orderHandler) CheckCarInvalidInOrder(ctx context.Context, in *CarIds, out *common.Response) error {
	return h.OrderHandler.CheckCarInvalidInOrder(ctx, in, out)
}

func (h *orderHandler) QueryOrderForSendSms(ctx context.Context, in *DeliveryInfo, out *common.Response) error {
	return h.OrderHandler.QueryOrderForSendSms(ctx, in, out)
}

func (h *orderHandler) GetNoAssignCustomerByOrder(ctx context.Context, in *OrderCondition, out *common.Response) error {
	return h.OrderHandler.GetNoAssignCustomerByOrder(ctx, in, out)
}

func (h *orderHandler) GetListWithPay(ctx context.Context, in *OrderCondition, out *common.Response) error {
	return h.OrderHandler.GetListWithPay(ctx, in, out)
}

func (h *orderHandler) AssignOrderPIC(ctx context.Context, in *OrderCondition, out *common.Response) error {
	return h.OrderHandler.AssignOrderPIC(ctx, in, out)
}

func (h *orderHandler) CancelOnlineOrder(ctx context.Context, in *OrderCondition, out *common.Response) error {
	return h.OrderHandler.CancelOnlineOrder(ctx, in, out)
}

func (h *orderHandler) GetPayInfo(ctx context.Context, in *QueryPayDTO, out *common.Response) error {
	return h.OrderHandler.GetPayInfo(ctx, in, out)
}

func (h *orderHandler) CheckCarInvalid(ctx context.Context, in *CarCheckDTO, out *common.Response) error {
	return h.OrderHandler.CheckCarInvalid(ctx, in, out)
}

func (h *orderHandler) Get(ctx context.Context, in *OrderCondition, out *common.Response) error {
	return h.OrderHandler.Get(ctx, in, out)
}

func (h *orderHandler) GetByOrderNo(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.OrderHandler.GetByOrderNo(ctx, in, out)
}
