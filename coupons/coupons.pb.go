// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: coupons/coupons.proto

package coupons

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupons_coupons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_coupons_coupons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_coupons_coupons_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CouponDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CouponNo   string `protobuf:"bytes,2,opt,name=CouponNo,proto3" json:"CouponNo,omitempty"`
	Title      string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Icon       string `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	CouponDesc string `protobuf:"bytes,5,opt,name=CouponDesc,proto3" json:"CouponDesc,omitempty"`
	Amount     string `protobuf:"bytes,6,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Type       uint32 `protobuf:"varint,7,opt,name=Type,proto3" json:"Type,omitempty"`
	StartTime  int64  `protobuf:"varint,8,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime    int64  `protobuf:"varint,9,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	UserId     uint64 `protobuf:"varint,10,opt,name=UserId,proto3" json:"UserId,omitempty"`
	OrderId    uint64 `protobuf:"varint,11,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	PaymentId  uint64 `protobuf:"varint,12,opt,name=PaymentId,proto3" json:"PaymentId,omitempty"`
	UseStatus  uint32 `protobuf:"varint,13,opt,name=UseStatus,proto3" json:"UseStatus,omitempty"`
	Status     uint32 `protobuf:"varint,14,opt,name=Status,proto3" json:"Status,omitempty"`
	CreateTime int64  `protobuf:"varint,15,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,16,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *CouponDto) Reset() {
	*x = CouponDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupons_coupons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponDto) ProtoMessage() {}

func (x *CouponDto) ProtoReflect() protoreflect.Message {
	mi := &file_coupons_coupons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponDto.ProtoReflect.Descriptor instead.
func (*CouponDto) Descriptor() ([]byte, []int) {
	return file_coupons_coupons_proto_rawDescGZIP(), []int{1}
}

func (x *CouponDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CouponDto) GetCouponNo() string {
	if x != nil {
		return x.CouponNo
	}
	return ""
}

func (x *CouponDto) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CouponDto) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CouponDto) GetCouponDesc() string {
	if x != nil {
		return x.CouponDesc
	}
	return ""
}

func (x *CouponDto) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CouponDto) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CouponDto) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *CouponDto) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *CouponDto) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CouponDto) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *CouponDto) GetPaymentId() uint64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *CouponDto) GetUseStatus() uint32 {
	if x != nil {
		return x.UseStatus
	}
	return 0
}

func (x *CouponDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CouponDto) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CouponDto) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type PagedList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo   uint32     `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize uint32     `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Total    uint32     `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	List     []*any.Any `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PagedList) Reset() {
	*x = PagedList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupons_coupons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedList) ProtoMessage() {}

func (x *PagedList) ProtoReflect() protoreflect.Message {
	mi := &file_coupons_coupons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedList.ProtoReflect.Descriptor instead.
func (*PagedList) Descriptor() ([]byte, []int) {
	return file_coupons_coupons_proto_rawDescGZIP(), []int{2}
}

func (x *PagedList) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *PagedList) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PagedList) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagedList) GetList() []*any.Any {
	if x != nil {
		return x.List
	}
	return nil
}

var File_coupons_coupons_proto protoreflect.FileDescriptor

var file_coupons_coupons_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62,
	0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xab, 0x03, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x55, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xfd, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x73, 0x12, 0x2d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x74, 0x6f,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_coupons_coupons_proto_rawDescOnce sync.Once
	file_coupons_coupons_proto_rawDescData = file_coupons_coupons_proto_rawDesc
)

func file_coupons_coupons_proto_rawDescGZIP() []byte {
	file_coupons_coupons_proto_rawDescOnce.Do(func() {
		file_coupons_coupons_proto_rawDescData = protoimpl.X.CompressGZIP(file_coupons_coupons_proto_rawDescData)
	})
	return file_coupons_coupons_proto_rawDescData
}

var file_coupons_coupons_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coupons_coupons_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: coupons.User
	(*CouponDto)(nil),       // 1: coupons.CouponDto
	(*PagedList)(nil),       // 2: coupons.PagedList
	(*any.Any)(nil),         // 3: google.protobuf.Any
	(*common.Page)(nil),     // 4: common.Page
	(*common.Response)(nil), // 5: common.Response
}
var file_coupons_coupons_proto_depIdxs = []int32{
	3, // 0: coupons.PagedList.list:type_name -> google.protobuf.Any
	1, // 1: coupons.Coupons.Add:input_type -> coupons.CouponDto
	1, // 2: coupons.Coupons.Delete:input_type -> coupons.CouponDto
	1, // 3: coupons.Coupons.Update:input_type -> coupons.CouponDto
	4, // 4: coupons.Coupons.List:input_type -> common.Page
	0, // 5: coupons.Coupons.FindCouponByUser:input_type -> coupons.User
	5, // 6: coupons.Coupons.Add:output_type -> common.Response
	5, // 7: coupons.Coupons.Delete:output_type -> common.Response
	5, // 8: coupons.Coupons.Update:output_type -> common.Response
	5, // 9: coupons.Coupons.List:output_type -> common.Response
	5, // 10: coupons.Coupons.FindCouponByUser:output_type -> common.Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coupons_coupons_proto_init() }
func file_coupons_coupons_proto_init() {
	if File_coupons_coupons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coupons_coupons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coupons_coupons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coupons_coupons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coupons_coupons_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coupons_coupons_proto_goTypes,
		DependencyIndexes: file_coupons_coupons_proto_depIdxs,
		MessageInfos:      file_coupons_coupons_proto_msgTypes,
	}.Build()
	File_coupons_coupons_proto = out.File
	file_coupons_coupons_proto_rawDesc = nil
	file_coupons_coupons_proto_goTypes = nil
	file_coupons_coupons_proto_depIdxs = nil
}
