// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.14.0
// source: contactRecord/contactRecord.proto

package contactRecord

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

type ContactRecordDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	StaffId     uint64 `protobuf:"varint,3,opt,name=staffId,proto3" json:"staffId,omitempty"`
	StoreId     uint64 `protobuf:"varint,4,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Method      uint32 `protobuf:"varint,5,opt,name=method,proto3" json:"method,omitempty"`
	Content     string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`          //备注
	ContactTime uint64 `protobuf:"varint,7,opt,name=contactTime,proto3" json:"contactTime,omitempty"` //沟通时间
	CreateTime  uint64 `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  uint64 `protobuf:"varint,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *ContactRecordDTO) Reset() {
	*x = ContactRecordDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactRecord_contactRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRecordDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRecordDTO) ProtoMessage() {}

func (x *ContactRecordDTO) ProtoReflect() protoreflect.Message {
	mi := &file_contactRecord_contactRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRecordDTO.ProtoReflect.Descriptor instead.
func (*ContactRecordDTO) Descriptor() ([]byte, []int) {
	return file_contactRecord_contactRecord_proto_rawDescGZIP(), []int{0}
}

func (x *ContactRecordDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ContactRecordDTO) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ContactRecordDTO) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *ContactRecordDTO) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ContactRecordDTO) GetMethod() uint32 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *ContactRecordDTO) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ContactRecordDTO) GetContactTime() uint64 {
	if x != nil {
		return x.ContactTime
	}
	return 0
}

func (x *ContactRecordDTO) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ContactRecordDTO) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type ContactRecordList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactRecord []*ContactRecordDTO `protobuf:"bytes,5,rep,name=contactRecord,proto3" json:"contactRecord,omitempty"`
}

func (x *ContactRecordList) Reset() {
	*x = ContactRecordList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactRecord_contactRecord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRecordList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRecordList) ProtoMessage() {}

func (x *ContactRecordList) ProtoReflect() protoreflect.Message {
	mi := &file_contactRecord_contactRecord_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRecordList.ProtoReflect.Descriptor instead.
func (*ContactRecordList) Descriptor() ([]byte, []int) {
	return file_contactRecord_contactRecord_proto_rawDescGZIP(), []int{1}
}

func (x *ContactRecordList) GetContactRecord() []*ContactRecordDTO {
	if x != nil {
		return x.ContactRecord
	}
	return nil
}

type ContactRecordCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo   uint32   `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize uint32   `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	UserId   uint64   `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	StaffId  uint64   `protobuf:"varint,5,opt,name=staffId,proto3" json:"staffId,omitempty"`
	StoreId  uint64   `protobuf:"varint,8,opt,name=storeId,proto3" json:"storeId,omitempty"`
	StoreIds []uint64 `protobuf:"varint,9,rep,packed,name=storeIds,proto3" json:"storeIds,omitempty"`
}

func (x *ContactRecordCondition) Reset() {
	*x = ContactRecordCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactRecord_contactRecord_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRecordCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRecordCondition) ProtoMessage() {}

func (x *ContactRecordCondition) ProtoReflect() protoreflect.Message {
	mi := &file_contactRecord_contactRecord_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRecordCondition.ProtoReflect.Descriptor instead.
func (*ContactRecordCondition) Descriptor() ([]byte, []int) {
	return file_contactRecord_contactRecord_proto_rawDescGZIP(), []int{2}
}

func (x *ContactRecordCondition) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *ContactRecordCondition) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ContactRecordCondition) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ContactRecordCondition) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *ContactRecordCondition) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ContactRecordCondition) GetStoreIds() []uint64 {
	if x != nil {
		return x.StoreIds
	}
	return nil
}

type LatestContactCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []uint64 `protobuf:"varint,1,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	StoreId uint64   `protobuf:"varint,2,opt,name=storeId,proto3" json:"storeId,omitempty"`
}

func (x *LatestContactCond) Reset() {
	*x = LatestContactCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactRecord_contactRecord_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatestContactCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatestContactCond) ProtoMessage() {}

func (x *LatestContactCond) ProtoReflect() protoreflect.Message {
	mi := &file_contactRecord_contactRecord_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatestContactCond.ProtoReflect.Descriptor instead.
func (*LatestContactCond) Descriptor() ([]byte, []int) {
	return file_contactRecord_contactRecord_proto_rawDescGZIP(), []int{3}
}

func (x *LatestContactCond) GetUserIds() []uint64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *LatestContactCond) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

var File_contactRecord_contactRecord_proto protoreflect.FileDescriptor

var file_contactRecord_contactRecord_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44,
	0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0xb4,
	0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x73, 0x22, 0x47, 0x0a, 0x11, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x32, 0xdf,
	0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x3a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_contactRecord_contactRecord_proto_rawDescOnce sync.Once
	file_contactRecord_contactRecord_proto_rawDescData = file_contactRecord_contactRecord_proto_rawDesc
)

func file_contactRecord_contactRecord_proto_rawDescGZIP() []byte {
	file_contactRecord_contactRecord_proto_rawDescOnce.Do(func() {
		file_contactRecord_contactRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_contactRecord_contactRecord_proto_rawDescData)
	})
	return file_contactRecord_contactRecord_proto_rawDescData
}

var file_contactRecord_contactRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contactRecord_contactRecord_proto_goTypes = []interface{}{
	(*ContactRecordDTO)(nil),       // 0: contactRecord.ContactRecordDTO
	(*ContactRecordList)(nil),      // 1: contactRecord.ContactRecordList
	(*ContactRecordCondition)(nil), // 2: contactRecord.ContactRecordCondition
	(*LatestContactCond)(nil),      // 3: contactRecord.LatestContactCond
	(*common.Response)(nil),        // 4: common.Response
}
var file_contactRecord_contactRecord_proto_depIdxs = []int32{
	0, // 0: contactRecord.ContactRecordList.contactRecord:type_name -> contactRecord.ContactRecordDTO
	0, // 1: contactRecord.ContactRecord.Add:input_type -> contactRecord.ContactRecordDTO
	2, // 2: contactRecord.ContactRecord.GetList:input_type -> contactRecord.ContactRecordCondition
	3, // 3: contactRecord.ContactRecord.GetLatestContactList:input_type -> contactRecord.LatestContactCond
	4, // 4: contactRecord.ContactRecord.Add:output_type -> common.Response
	4, // 5: contactRecord.ContactRecord.GetList:output_type -> common.Response
	4, // 6: contactRecord.ContactRecord.GetLatestContactList:output_type -> common.Response
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contactRecord_contactRecord_proto_init() }
func file_contactRecord_contactRecord_proto_init() {
	if File_contactRecord_contactRecord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contactRecord_contactRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRecordDTO); i {
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
		file_contactRecord_contactRecord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRecordList); i {
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
		file_contactRecord_contactRecord_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRecordCondition); i {
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
		file_contactRecord_contactRecord_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatestContactCond); i {
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
			RawDescriptor: file_contactRecord_contactRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contactRecord_contactRecord_proto_goTypes,
		DependencyIndexes: file_contactRecord_contactRecord_proto_depIdxs,
		MessageInfos:      file_contactRecord_contactRecord_proto_msgTypes,
	}.Build()
	File_contactRecord_contactRecord_proto = out.File
	file_contactRecord_contactRecord_proto_rawDesc = nil
	file_contactRecord_contactRecord_proto_goTypes = nil
	file_contactRecord_contactRecord_proto_depIdxs = nil
}
