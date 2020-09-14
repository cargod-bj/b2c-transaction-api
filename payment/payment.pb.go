// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.9.0
// source: payment/payment.proto

package payment

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	fileResource "github.com/cargod-bj/b2c-transaction-api/fileResource"
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

	Id          uint64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	OrderId     uint64                  `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	StaffId     uint64                  `protobuf:"varint,4,opt,name=staffId,proto3" json:"staffId,omitempty"`
	Method      uint32                  `protobuf:"varint,5,opt,name=method,proto3" json:"method,omitempty"`
	Amount      string                  `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Status      uint32                  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime  int64                   `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  int64                   `protobuf:"varint,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	RecordType  uint32                  `protobuf:"varint,10,opt,name=recordType,proto3" json:"recordType,omitempty"`
	PaymentType uint32                  `protobuf:"varint,11,opt,name=paymentType,proto3" json:"paymentType,omitempty"`
	Remark      string                  `protobuf:"bytes,12,opt,name=remark,proto3" json:"remark,omitempty"`
	FileDto     []*fileResource.FileDTO `protobuf:"bytes,13,rep,name=fileDto,proto3" json:"fileDto,omitempty"`
	CouponNo    string                  `protobuf:"bytes,14,opt,name=couponNo,proto3" json:"couponNo,omitempty"`
	RejectNote  string                  `protobuf:"bytes,16,opt,name=rejectNote,proto3" json:"rejectNote,omitempty"`
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

func (x *PaymentDto) GetFileDto() []*fileResource.FileDTO {
	if x != nil {
		return x.FileDto
	}
	return nil
}

func (x *PaymentDto) GetCouponNo() string {
	if x != nil {
		return x.CouponNo
	}
	return ""
}

func (x *PaymentDto) GetRejectNote() string {
	if x != nil {
		return x.RejectNote
	}
	return ""
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

type OrderId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *OrderId) Reset() {
	*x = OrderId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderId) ProtoMessage() {}

func (x *OrderId) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OrderId.ProtoReflect.Descriptor instead.
func (*OrderId) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *OrderId) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type PaymentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentList []*PaymentDto `protobuf:"bytes,1,rep,name=paymentList,proto3" json:"paymentList,omitempty"`
	TotalAmount string        `protobuf:"bytes,2,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
}

func (x *PaymentList) Reset() {
	*x = PaymentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentList) ProtoMessage() {}

func (x *PaymentList) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[3]
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
	return file_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentList) GetPaymentList() []*PaymentDto {
	if x != nil {
		return x.PaymentList
	}
	return nil
}

func (x *PaymentList) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

type PaymentCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PaymentCheck) Reset() {
	*x = PaymentCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCheck) ProtoMessage() {}

func (x *PaymentCheck) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCheck.ProtoReflect.Descriptor instead.
func (*PaymentCheck) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *PaymentCheck) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ApprovalCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId    uint64 `protobuf:"varint,1,opt,name=storeId,proto3" json:"storeId,omitempty"`
	OrderNo    string `protobuf:"bytes,2,opt,name=orderNo,proto3" json:"orderNo,omitempty"`
	Status     uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	LicenseNum string `protobuf:"bytes,4,opt,name=licenseNum,proto3" json:"licenseNum,omitempty"`
	Method     uint32 `protobuf:"varint,5,opt,name=method,proto3" json:"method,omitempty"`
	UserId     uint64 `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
	Page       int32  `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32  `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	FromTime   string `protobuf:"bytes,9,opt,name=fromTime,proto3" json:"fromTime,omitempty"`
	ToTime     string `protobuf:"bytes,10,opt,name=toTime,proto3" json:"toTime,omitempty"`
}

func (x *ApprovalCond) Reset() {
	*x = ApprovalCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalCond) ProtoMessage() {}

func (x *ApprovalCond) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalCond.ProtoReflect.Descriptor instead.
func (*ApprovalCond) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{5}
}

func (x *ApprovalCond) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ApprovalCond) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *ApprovalCond) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ApprovalCond) GetLicenseNum() string {
	if x != nil {
		return x.LicenseNum
	}
	return ""
}

func (x *ApprovalCond) GetMethod() uint32 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *ApprovalCond) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ApprovalCond) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ApprovalCond) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ApprovalCond) GetFromTime() string {
	if x != nil {
		return x.FromTime
	}
	return ""
}

func (x *ApprovalCond) GetToTime() string {
	if x != nil {
		return x.ToTime
	}
	return ""
}

var File_payment_payment_proto protoreflect.FileDescriptor

var file_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x03, 0x0a,
	0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x74, 0x6f, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x54,
	0x4f, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x6f, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x5f, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x23, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x0b,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x88, 0x02, 0x0a,
	0x0c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x82, 0x04, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d,
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
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e,
	0x64, 0x12, 0x15, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f,
	0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_payment_payment_proto_goTypes = []interface{}{
	(*PaymentDto)(nil),           // 0: payment.PaymentDto
	(*PaymentCond)(nil),          // 1: payment.PaymentCond
	(*OrderId)(nil),              // 2: payment.OrderId
	(*PaymentList)(nil),          // 3: payment.PaymentList
	(*PaymentCheck)(nil),         // 4: payment.PaymentCheck
	(*ApprovalCond)(nil),         // 5: payment.ApprovalCond
	(*fileResource.FileDTO)(nil), // 6: fileResource.FileDTO
	(*common.Page)(nil),          // 7: common.Page
	(*common.Response)(nil),      // 8: common.Response
}
var file_payment_payment_proto_depIdxs = []int32{
	6,  // 0: payment.PaymentDto.fileDto:type_name -> fileResource.FileDTO
	0,  // 1: payment.PaymentList.paymentList:type_name -> payment.PaymentDto
	0,  // 2: payment.Payment.Add:input_type -> payment.PaymentDto
	0,  // 3: payment.Payment.Delete:input_type -> payment.PaymentDto
	0,  // 4: payment.Payment.Update:input_type -> payment.PaymentDto
	7,  // 5: payment.Payment.List:input_type -> common.Page
	1,  // 6: payment.Payment.GetPaymentListByCond:input_type -> payment.PaymentCond
	0,  // 7: payment.Payment.AddPaymentImages:input_type -> payment.PaymentDto
	2,  // 8: payment.Payment.CheckOrderPaymentStatus:input_type -> payment.OrderId
	0,  // 9: payment.Payment.UpdateStatus:input_type -> payment.PaymentDto
	5,  // 10: payment.Payment.GetPaymentsByCond:input_type -> payment.ApprovalCond
	8,  // 11: payment.Payment.Add:output_type -> common.Response
	8,  // 12: payment.Payment.Delete:output_type -> common.Response
	8,  // 13: payment.Payment.Update:output_type -> common.Response
	8,  // 14: payment.Payment.List:output_type -> common.Response
	8,  // 15: payment.Payment.GetPaymentListByCond:output_type -> common.Response
	8,  // 16: payment.Payment.AddPaymentImages:output_type -> common.Response
	8,  // 17: payment.Payment.CheckOrderPaymentStatus:output_type -> common.Response
	8,  // 18: payment.Payment.UpdateStatus:output_type -> common.Response
	8,  // 19: payment.Payment.GetPaymentsByCond:output_type -> common.Response
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
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
			switch v := v.(*OrderId); i {
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
		file_payment_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_payment_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCheck); i {
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
		file_payment_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalCond); i {
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
			NumMessages:   6,
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
