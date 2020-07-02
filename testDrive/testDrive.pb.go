// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.7.1
// source: testDrive/testDrive.proto

package testDrive

import (
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

type TestDriveDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	AppointmentId uint64 `protobuf:"varint,3,opt,name=appointmentId,proto3" json:"appointmentId,omitempty"` //预约
	StaffId       uint64 `protobuf:"varint,4,opt,name=staffId,proto3" json:"staffId,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"` //备注
	CarId         uint64 `protobuf:"varint,6,opt,name=carId,proto3" json:"carId,omitempty"`
	StoreId       uint64 `protobuf:"varint,7,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Status        uint32 `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Appointment   string `protobuf:"bytes,9,opt,name=appointment,proto3" json:"appointment,omitempty"` //预约时间
	CreateTime    uint64 `protobuf:"varint,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime    uint64 `protobuf:"varint,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *TestDriveDTO) Reset() {
	*x = TestDriveDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testDrive_testDrive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestDriveDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestDriveDTO) ProtoMessage() {}

func (x *TestDriveDTO) ProtoReflect() protoreflect.Message {
	mi := &file_testDrive_testDrive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestDriveDTO.ProtoReflect.Descriptor instead.
func (*TestDriveDTO) Descriptor() ([]byte, []int) {
	return file_testDrive_testDrive_proto_rawDescGZIP(), []int{0}
}

func (x *TestDriveDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestDriveDTO) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TestDriveDTO) GetAppointmentId() uint64 {
	if x != nil {
		return x.AppointmentId
	}
	return 0
}

func (x *TestDriveDTO) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *TestDriveDTO) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TestDriveDTO) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *TestDriveDTO) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *TestDriveDTO) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TestDriveDTO) GetAppointment() string {
	if x != nil {
		return x.Appointment
	}
	return ""
}

func (x *TestDriveDTO) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *TestDriveDTO) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type TestDriveList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo    uint32          `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize  uint32          `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	TestDrive []*TestDriveDTO `protobuf:"bytes,3,rep,name=testDrive,proto3" json:"testDrive,omitempty"`
	Count     uint32          `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *TestDriveList) Reset() {
	*x = TestDriveList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testDrive_testDrive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestDriveList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestDriveList) ProtoMessage() {}

func (x *TestDriveList) ProtoReflect() protoreflect.Message {
	mi := &file_testDrive_testDrive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestDriveList.ProtoReflect.Descriptor instead.
func (*TestDriveList) Descriptor() ([]byte, []int) {
	return file_testDrive_testDrive_proto_rawDescGZIP(), []int{1}
}

func (x *TestDriveList) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *TestDriveList) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TestDriveList) GetTestDrive() []*TestDriveDTO {
	if x != nil {
		return x.TestDrive
	}
	return nil
}

func (x *TestDriveList) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TestDriveCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	UserId uint64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *TestDriveCondition) Reset() {
	*x = TestDriveCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testDrive_testDrive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestDriveCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestDriveCondition) ProtoMessage() {}

func (x *TestDriveCondition) ProtoReflect() protoreflect.Message {
	mi := &file_testDrive_testDrive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestDriveCondition.ProtoReflect.Descriptor instead.
func (*TestDriveCondition) Descriptor() ([]byte, []int) {
	return file_testDrive_testDrive_proto_rawDescGZIP(), []int{2}
}

func (x *TestDriveCondition) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *TestDriveCondition) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TestDriveCondition) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Code string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testDrive_testDrive_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_testDrive_testDrive_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_testDrive_testDrive_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AddTestDriveData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddTestDriveData) Reset() {
	*x = AddTestDriveData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testDrive_testDrive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTestDriveData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTestDriveData) ProtoMessage() {}

func (x *AddTestDriveData) ProtoReflect() protoreflect.Message {
	mi := &file_testDrive_testDrive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTestDriveData.ProtoReflect.Descriptor instead.
func (*AddTestDriveData) Descriptor() ([]byte, []int) {
	return file_testDrive_testDrive_proto_rawDescGZIP(), []int{4}
}

func (x *AddTestDriveData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteId) Reset() {
	*x = DeleteId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testDrive_testDrive_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteId) ProtoMessage() {}

func (x *DeleteId) ProtoReflect() protoreflect.Message {
	mi := &file_testDrive_testDrive_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteId.ProtoReflect.Descriptor instead.
func (*DeleteId) Descriptor() ([]byte, []int) {
	return file_testDrive_testDrive_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteId) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_testDrive_testDrive_proto protoreflect.FileDescriptor

var file_testDrive_testDrive_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x73,
	0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x44,
	0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x35, 0x0a, 0x09,
	0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x44, 0x54, 0x4f, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x12, 0x54, 0x65, 0x73,
	0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x5a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x22, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x54, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x32, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xf3, 0x01, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x12, 0x35, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x44, 0x54, 0x4f, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x44, 0x54, 0x4f, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64,
	0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testDrive_testDrive_proto_rawDescOnce sync.Once
	file_testDrive_testDrive_proto_rawDescData = file_testDrive_testDrive_proto_rawDesc
)

func file_testDrive_testDrive_proto_rawDescGZIP() []byte {
	file_testDrive_testDrive_proto_rawDescOnce.Do(func() {
		file_testDrive_testDrive_proto_rawDescData = protoimpl.X.CompressGZIP(file_testDrive_testDrive_proto_rawDescData)
	})
	return file_testDrive_testDrive_proto_rawDescData
}

var file_testDrive_testDrive_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_testDrive_testDrive_proto_goTypes = []interface{}{
	(*TestDriveDTO)(nil),       // 0: testDrive.TestDriveDTO
	(*TestDriveList)(nil),      // 1: testDrive.TestDriveList
	(*TestDriveCondition)(nil), // 2: testDrive.TestDriveCondition
	(*Response)(nil),           // 3: testDrive.Response
	(*AddTestDriveData)(nil),   // 4: testDrive.AddTestDriveData
	(*DeleteId)(nil),           // 5: testDrive.DeleteId
	(*any.Any)(nil),            // 6: google.protobuf.Any
}
var file_testDrive_testDrive_proto_depIdxs = []int32{
	0, // 0: testDrive.TestDriveList.testDrive:type_name -> testDrive.TestDriveDTO
	6, // 1: testDrive.Response.data:type_name -> google.protobuf.Any
	0, // 2: testDrive.TestDrive.Add:input_type -> testDrive.TestDriveDTO
	5, // 3: testDrive.TestDrive.Delete:input_type -> testDrive.DeleteId
	0, // 4: testDrive.TestDrive.Update:input_type -> testDrive.TestDriveDTO
	2, // 5: testDrive.TestDrive.GetList:input_type -> testDrive.TestDriveCondition
	3, // 6: testDrive.TestDrive.Add:output_type -> testDrive.Response
	3, // 7: testDrive.TestDrive.Delete:output_type -> testDrive.Response
	3, // 8: testDrive.TestDrive.Update:output_type -> testDrive.Response
	3, // 9: testDrive.TestDrive.GetList:output_type -> testDrive.Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_testDrive_testDrive_proto_init() }
func file_testDrive_testDrive_proto_init() {
	if File_testDrive_testDrive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testDrive_testDrive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestDriveDTO); i {
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
		file_testDrive_testDrive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestDriveList); i {
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
		file_testDrive_testDrive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestDriveCondition); i {
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
		file_testDrive_testDrive_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_testDrive_testDrive_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTestDriveData); i {
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
		file_testDrive_testDrive_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteId); i {
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
			RawDescriptor: file_testDrive_testDrive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testDrive_testDrive_proto_goTypes,
		DependencyIndexes: file_testDrive_testDrive_proto_depIdxs,
		MessageInfos:      file_testDrive_testDrive_proto_msgTypes,
	}.Build()
	File_testDrive_testDrive_proto = out.File
	file_testDrive_testDrive_proto_rawDesc = nil
	file_testDrive_testDrive_proto_goTypes = nil
	file_testDrive_testDrive_proto_depIdxs = nil
}
