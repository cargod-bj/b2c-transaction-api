// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: appointment/appointment.proto

package appointment

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

type AppointmentDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	LeadId      uint64 `protobuf:"varint,3,opt,name=leadId,proto3" json:"leadId,omitempty"`
	StaffId     uint64 `protobuf:"varint,4,opt,name=staffId,proto3" json:"staffId,omitempty"`
	Type        uint32 `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	CarId       uint64 `protobuf:"varint,6,opt,name=carId,proto3" json:"carId,omitempty"`
	StoreId     uint64 `protobuf:"varint,7,opt,name=storeId,proto3" json:"storeId,omitempty"`
	Status      uint32 `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Appointment string `protobuf:"bytes,9,opt,name=appointment,proto3" json:"appointment,omitempty"` //预约时间
	CreateTime  uint64 `protobuf:"varint,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  uint64 `protobuf:"varint,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *AppointmentDTO) Reset() {
	*x = AppointmentDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_appointment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppointmentDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppointmentDTO) ProtoMessage() {}

func (x *AppointmentDTO) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppointmentDTO.ProtoReflect.Descriptor instead.
func (*AppointmentDTO) Descriptor() ([]byte, []int) {
	return file_appointment_appointment_proto_rawDescGZIP(), []int{0}
}

func (x *AppointmentDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppointmentDTO) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AppointmentDTO) GetLeadId() uint64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

func (x *AppointmentDTO) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *AppointmentDTO) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AppointmentDTO) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *AppointmentDTO) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *AppointmentDTO) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AppointmentDTO) GetAppointment() string {
	if x != nil {
		return x.Appointment
	}
	return ""
}

func (x *AppointmentDTO) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *AppointmentDTO) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type AppointmentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo           uint32            `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize         uint32            `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Appointment      []*AppointmentDTO `protobuf:"bytes,3,rep,name=appointment,proto3" json:"appointment,omitempty"`
	Count            uint32            `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	UnCompletedCount uint32            `protobuf:"varint,5,opt,name=unCompletedCount,proto3" json:"unCompletedCount,omitempty"`
	CompletedCount   uint32            `protobuf:"varint,6,opt,name=CompletedCount,proto3" json:"CompletedCount,omitempty"`
}

func (x *AppointmentList) Reset() {
	*x = AppointmentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_appointment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppointmentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppointmentList) ProtoMessage() {}

func (x *AppointmentList) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppointmentList.ProtoReflect.Descriptor instead.
func (*AppointmentList) Descriptor() ([]byte, []int) {
	return file_appointment_appointment_proto_rawDescGZIP(), []int{1}
}

func (x *AppointmentList) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *AppointmentList) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppointmentList) GetAppointment() []*AppointmentDTO {
	if x != nil {
		return x.Appointment
	}
	return nil
}

func (x *AppointmentList) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AppointmentList) GetUnCompletedCount() uint32 {
	if x != nil {
		return x.UnCompletedCount
	}
	return 0
}

func (x *AppointmentList) GetCompletedCount() uint32 {
	if x != nil {
		return x.CompletedCount
	}
	return 0
}

type AppointmentCon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo    uint32 `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize  uint32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	UserId    uint64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Completed uint32 `protobuf:"varint,4,opt,name=Completed,proto3" json:"Completed,omitempty"` //0表示未完成的，1表示完成的
}

func (x *AppointmentCon) Reset() {
	*x = AppointmentCon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_appointment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppointmentCon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppointmentCon) ProtoMessage() {}

func (x *AppointmentCon) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppointmentCon.ProtoReflect.Descriptor instead.
func (*AppointmentCon) Descriptor() ([]byte, []int) {
	return file_appointment_appointment_proto_rawDescGZIP(), []int{2}
}

func (x *AppointmentCon) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *AppointmentCon) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppointmentCon) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AppointmentCon) GetCompleted() uint32 {
	if x != nil {
		return x.Completed
	}
	return 0
}

type AppointmentCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo      uint32   `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize    uint32   `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	UserId      []uint64 `protobuf:"varint,3,rep,packed,name=UserId,proto3" json:"UserId,omitempty"`
	CarId       uint64   `protobuf:"varint,7,opt,name=CarId,proto3" json:"CarId,omitempty"`
	StaffId     []uint64 `protobuf:"varint,6,rep,packed,name=StaffId,proto3" json:"StaffId,omitempty"`
	StoreId     uint64   `protobuf:"varint,9,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	Appointment string   `protobuf:"bytes,11,opt,name=Appointment,proto3" json:"Appointment,omitempty"` //预约时间
	Status      []uint32 `protobuf:"varint,12,rep,packed,name=Status,proto3" json:"Status,omitempty"`
}

func (x *AppointmentCondition) Reset() {
	*x = AppointmentCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_appointment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppointmentCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppointmentCondition) ProtoMessage() {}

func (x *AppointmentCondition) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppointmentCondition.ProtoReflect.Descriptor instead.
func (*AppointmentCondition) Descriptor() ([]byte, []int) {
	return file_appointment_appointment_proto_rawDescGZIP(), []int{3}
}

func (x *AppointmentCondition) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *AppointmentCondition) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppointmentCondition) GetUserId() []uint64 {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *AppointmentCondition) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *AppointmentCondition) GetStaffId() []uint64 {
	if x != nil {
		return x.StaffId
	}
	return nil
}

func (x *AppointmentCondition) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *AppointmentCondition) GetAppointment() string {
	if x != nil {
		return x.Appointment
	}
	return ""
}

func (x *AppointmentCondition) GetStatus() []uint32 {
	if x != nil {
		return x.Status
	}
	return nil
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
		mi := &file_appointment_appointment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[4]
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
	return file_appointment_appointment_proto_rawDescGZIP(), []int{4}
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

type AddAppointmentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddAppointmentData) Reset() {
	*x = AddAppointmentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_appointment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAppointmentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppointmentData) ProtoMessage() {}

func (x *AddAppointmentData) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppointmentData.ProtoReflect.Descriptor instead.
func (*AddAppointmentData) Descriptor() ([]byte, []int) {
	return file_appointment_appointment_proto_rawDescGZIP(), []int{5}
}

func (x *AddAppointmentData) GetId() uint64 {
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
		mi := &file_appointment_appointment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteId) ProtoMessage() {}

func (x *DeleteId) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[6]
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
	return file_appointment_appointment_proto_rawDescGZIP(), []int{6}
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

type AssignDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids     []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	StaffId uint64   `protobuf:"varint,2,opt,name=StaffId,proto3" json:"StaffId,omitempty"`
}

func (x *AssignDto) Reset() {
	*x = AssignDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_appointment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignDto) ProtoMessage() {}

func (x *AssignDto) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_appointment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignDto.ProtoReflect.Descriptor instead.
func (*AssignDto) Descriptor() ([]byte, []int) {
	return file_appointment_appointment_proto_rawDescGZIP(), []int{7}
}

func (x *AssignDto) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AssignDto) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

var File_appointment_appointment_proto protoreflect.FileDescriptor

var file_appointment_appointment_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x0b, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x10, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x75, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22,
	0xe6, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x08, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37,
	0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x32, 0x96, 0x03, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1b,
	0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x1a, 0x15, 0x2e, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x54, 0x4f, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x15, 0x2e, 0x61,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44,
	0x74, 0x6f, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appointment_appointment_proto_rawDescOnce sync.Once
	file_appointment_appointment_proto_rawDescData = file_appointment_appointment_proto_rawDesc
)

func file_appointment_appointment_proto_rawDescGZIP() []byte {
	file_appointment_appointment_proto_rawDescOnce.Do(func() {
		file_appointment_appointment_proto_rawDescData = protoimpl.X.CompressGZIP(file_appointment_appointment_proto_rawDescData)
	})
	return file_appointment_appointment_proto_rawDescData
}

var file_appointment_appointment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_appointment_appointment_proto_goTypes = []interface{}{
	(*AppointmentDTO)(nil),       // 0: appointment.AppointmentDTO
	(*AppointmentList)(nil),      // 1: appointment.AppointmentList
	(*AppointmentCon)(nil),       // 2: appointment.AppointmentCon
	(*AppointmentCondition)(nil), // 3: appointment.AppointmentCondition
	(*Response)(nil),             // 4: appointment.Response
	(*AddAppointmentData)(nil),   // 5: appointment.AddAppointmentData
	(*DeleteId)(nil),             // 6: appointment.DeleteId
	(*AssignDto)(nil),            // 7: appointment.AssignDto
	(*any.Any)(nil),              // 8: google.protobuf.Any
}
var file_appointment_appointment_proto_depIdxs = []int32{
	0, // 0: appointment.AppointmentList.appointment:type_name -> appointment.AppointmentDTO
	8, // 1: appointment.Response.data:type_name -> google.protobuf.Any
	0, // 2: appointment.Appointment.Add:input_type -> appointment.AppointmentDTO
	6, // 3: appointment.Appointment.Delete:input_type -> appointment.DeleteId
	0, // 4: appointment.Appointment.Update:input_type -> appointment.AppointmentDTO
	3, // 5: appointment.Appointment.GetList:input_type -> appointment.AppointmentCondition
	7, // 6: appointment.Appointment.AssignAppointment:input_type -> appointment.AssignDto
	2, // 7: appointment.Appointment.GetUserList:input_type -> appointment.AppointmentCon
	4, // 8: appointment.Appointment.Add:output_type -> appointment.Response
	4, // 9: appointment.Appointment.Delete:output_type -> appointment.Response
	4, // 10: appointment.Appointment.Update:output_type -> appointment.Response
	4, // 11: appointment.Appointment.GetList:output_type -> appointment.Response
	4, // 12: appointment.Appointment.AssignAppointment:output_type -> appointment.Response
	4, // 13: appointment.Appointment.GetUserList:output_type -> appointment.Response
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_appointment_appointment_proto_init() }
func file_appointment_appointment_proto_init() {
	if File_appointment_appointment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appointment_appointment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppointmentDTO); i {
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
		file_appointment_appointment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppointmentList); i {
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
		file_appointment_appointment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppointmentCon); i {
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
		file_appointment_appointment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppointmentCondition); i {
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
		file_appointment_appointment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_appointment_appointment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAppointmentData); i {
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
		file_appointment_appointment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_appointment_appointment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignDto); i {
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
			RawDescriptor: file_appointment_appointment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appointment_appointment_proto_goTypes,
		DependencyIndexes: file_appointment_appointment_proto_depIdxs,
		MessageInfos:      file_appointment_appointment_proto_msgTypes,
	}.Build()
	File_appointment_appointment_proto = out.File
	file_appointment_appointment_proto_rawDesc = nil
	file_appointment_appointment_proto_goTypes = nil
	file_appointment_appointment_proto_depIdxs = nil
}
