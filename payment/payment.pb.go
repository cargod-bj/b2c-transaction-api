// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: payment/payment.proto

package payment

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
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

type PaymentDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint64   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	OrderId     uint64   `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	StaffId     uint64   `protobuf:"varint,4,opt,name=staffId,proto3" json:"staffId,omitempty"`
	Method      uint32   `protobuf:"varint,5,opt,name=method,proto3" json:"method,omitempty"`
	Amount      string   `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Status      uint32   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime  int64    `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  int64    `protobuf:"varint,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	RecordType  uint32   `protobuf:"varint,10,opt,name=recordType,proto3" json:"recordType,omitempty"`
	PaymentType uint32   `protobuf:"varint,11,opt,name=paymentType,proto3" json:"paymentType,omitempty"`
	Remark      string   `protobuf:"bytes,12,opt,name=remark,proto3" json:"remark,omitempty"`
	Images      []string `protobuf:"bytes,13,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *PaymentDto) Reset() {
	*x = PaymentDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentDto) ProtoMessage() {}

func (x *PaymentDto) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentDto.ProtoReflect.Descriptor instead.
func (*PaymentDto) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentDto) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PaymentDto) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *PaymentDto) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *PaymentDto) GetMethod() uint32 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *PaymentDto) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *PaymentDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PaymentDto) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *PaymentDto) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *PaymentDto) GetRecordType() uint32 {
	if x != nil {
		return x.RecordType
	}
	return 0
}

func (x *PaymentDto) GetPaymentType() uint32 {
	if x != nil {
		return x.PaymentType
	}
	return 0
}

func (x *PaymentDto) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *PaymentDto) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type PaymentCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    uint64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	RecordType uint32 `protobuf:"varint,2,opt,name=recordType,proto3" json:"recordType,omitempty"`
	Status     uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PaymentCond) Reset() {
	*x = PaymentCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCond) ProtoMessage() {}

func (x *PaymentCond) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCond.ProtoReflect.Descriptor instead.
func (*PaymentCond) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentCond) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *PaymentCond) GetRecordType() uint32 {
	if x != nil {
		return x.RecordType
	}
	return 0
}

func (x *PaymentCond) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PaymentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentList []*PaymentDto `protobuf:"bytes,1,rep,name=paymentList,proto3" json:"paymentList,omitempty"`
}

func (x *PaymentList) Reset() {
	*x = PaymentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentList) ProtoMessage() {}

func (x *PaymentList) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentList.ProtoReflect.Descriptor instead.
func (*PaymentList) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentList) GetPaymentList() []*PaymentDto {
	if x != nil {
		return x.PaymentList
	}
	return nil
}

var File_payment_payment_proto protoreflect.FileDescriptor

var file_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x0a,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x5f, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x44, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x35, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x8b, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_payment_payment_proto_rawDescOnce sync.Once
	file_payment_payment_proto_rawDescData = file_payment_payment_proto_rawDesc
)

func file_payment_payment_proto_rawDescGZIP() []byte {
	file_payment_payment_proto_rawDescOnce.Do(func() {
		file_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_payment_proto_rawDescData)
	})
	return file_payment_payment_proto_rawDescData
}

var file_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payment_payment_proto_goTypes = []interface{}{
	(*PaymentDto)(nil),      // 0: payment.PaymentDto
	(*PaymentCond)(nil),     // 1: payment.PaymentCond
	(*PaymentList)(nil),     // 2: payment.PaymentList
	(*common.Page)(nil),     // 3: common.Page
	(*common.Response)(nil), // 4: common.Response
}
var file_payment_payment_proto_depIdxs = []int32{
	0, // 0: payment.PaymentList.paymentList:type_name -> payment.PaymentDto
	0, // 1: payment.Payment.Add:input_type -> payment.PaymentDto
	0, // 2: payment.Payment.Delete:input_type -> payment.PaymentDto
	0, // 3: payment.Payment.Update:input_type -> payment.PaymentDto
	3, // 4: payment.Payment.List:input_type -> common.Page
	1, // 5: payment.Payment.GetPaymentListByCond:input_type -> payment.PaymentCond
	4, // 6: payment.Payment.Add:output_type -> common.Response
	4, // 7: payment.Payment.Delete:output_type -> common.Response
	4, // 8: payment.Payment.Update:output_type -> common.Response
	4, // 9: payment.Payment.List:output_type -> common.Response
	4, // 10: payment.Payment.GetPaymentListByCond:output_type -> common.Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_payment_proto_init() }
func file_payment_payment_proto_init() {
	if File_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentDto); i {
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
		file_payment_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCond); i {
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
		file_payment_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentList); i {
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
			RawDescriptor: file_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_payment_proto_goTypes,
		DependencyIndexes: file_payment_payment_proto_depIdxs,
		MessageInfos:      file_payment_payment_proto_msgTypes,
	}.Build()
	File_payment_payment_proto = out.File
	file_payment_payment_proto_rawDesc = nil
	file_payment_payment_proto_goTypes = nil
	file_payment_payment_proto_depIdxs = nil
}
