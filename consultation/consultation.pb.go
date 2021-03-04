// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.14.0
// source: consultation/consultation.proto

package consultation

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

type ConsultationRecordDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId             uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	CarId              uint64 `protobuf:"varint,3,opt,name=carId,proto3" json:"carId,omitempty"`
	StoreId            uint64 `protobuf:"varint,4,opt,name=storeId,proto3" json:"storeId,omitempty"`
	StaffId            uint64 `protobuf:"varint,5,opt,name=staffId,proto3" json:"staffId,omitempty"`
	ConsultationNo     string `protobuf:"bytes,6,opt,name=ConsultationNo,proto3" json:"ConsultationNo,omitempty"`
	ConsultationStatus uint32 `protobuf:"varint,7,opt,name=ConsultationStatus,proto3" json:"ConsultationStatus,omitempty"`
	Remark             string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"` //备注
	Status             string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime         uint64 `protobuf:"varint,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime         uint64 `protobuf:"varint,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *ConsultationRecordDTO) Reset() {
	*x = ConsultationRecordDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consultation_consultation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultationRecordDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationRecordDTO) ProtoMessage() {}

func (x *ConsultationRecordDTO) ProtoReflect() protoreflect.Message {
	mi := &file_consultation_consultation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultationRecordDTO.ProtoReflect.Descriptor instead.
func (*ConsultationRecordDTO) Descriptor() ([]byte, []int) {
	return file_consultation_consultation_proto_rawDescGZIP(), []int{0}
}

func (x *ConsultationRecordDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConsultationRecordDTO) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ConsultationRecordDTO) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *ConsultationRecordDTO) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ConsultationRecordDTO) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *ConsultationRecordDTO) GetConsultationNo() string {
	if x != nil {
		return x.ConsultationNo
	}
	return ""
}

func (x *ConsultationRecordDTO) GetConsultationStatus() uint32 {
	if x != nil {
		return x.ConsultationStatus
	}
	return 0
}

func (x *ConsultationRecordDTO) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ConsultationRecordDTO) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ConsultationRecordDTO) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ConsultationRecordDTO) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type ConsultationRecordList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsultationRecord []*ConsultationRecordDTO `protobuf:"bytes,1,rep,name=consultationRecord,proto3" json:"consultationRecord,omitempty"`
}

func (x *ConsultationRecordList) Reset() {
	*x = ConsultationRecordList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consultation_consultation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultationRecordList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationRecordList) ProtoMessage() {}

func (x *ConsultationRecordList) ProtoReflect() protoreflect.Message {
	mi := &file_consultation_consultation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultationRecordList.ProtoReflect.Descriptor instead.
func (*ConsultationRecordList) Descriptor() ([]byte, []int) {
	return file_consultation_consultation_proto_rawDescGZIP(), []int{1}
}

func (x *ConsultationRecordList) GetConsultationRecord() []*ConsultationRecordDTO {
	if x != nil {
		return x.ConsultationRecord
	}
	return nil
}

type EffectiveConsultationCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CarId   uint64 `protobuf:"varint,2,opt,name=carId,proto3" json:"carId,omitempty"`
	StaffId uint64 `protobuf:"varint,3,opt,name=staffId,proto3" json:"staffId,omitempty"`
}

func (x *EffectiveConsultationCond) Reset() {
	*x = EffectiveConsultationCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consultation_consultation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectiveConsultationCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveConsultationCond) ProtoMessage() {}

func (x *EffectiveConsultationCond) ProtoReflect() protoreflect.Message {
	mi := &file_consultation_consultation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveConsultationCond.ProtoReflect.Descriptor instead.
func (*EffectiveConsultationCond) Descriptor() ([]byte, []int) {
	return file_consultation_consultation_proto_rawDescGZIP(), []int{2}
}

func (x *EffectiveConsultationCond) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *EffectiveConsultationCond) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *EffectiveConsultationCond) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

type NsgConsultationCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *NsgConsultationCond) Reset() {
	*x = NsgConsultationCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consultation_consultation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NsgConsultationCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NsgConsultationCond) ProtoMessage() {}

func (x *NsgConsultationCond) ProtoReflect() protoreflect.Message {
	mi := &file_consultation_consultation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NsgConsultationCond.ProtoReflect.Descriptor instead.
func (*NsgConsultationCond) Descriptor() ([]byte, []int) {
	return file_consultation_consultation_proto_rawDescGZIP(), []int{3}
}

func (x *NsgConsultationCond) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type BatchUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ids                []uint64 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	ConsultationStatus uint32   `protobuf:"varint,3,opt,name=ConsultationStatus,proto3" json:"ConsultationStatus,omitempty"`
	Remark             string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"` //备注
	Status             string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BatchUpdateRequest) Reset() {
	*x = BatchUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consultation_consultation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpdateRequest) ProtoMessage() {}

func (x *BatchUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consultation_consultation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpdateRequest.ProtoReflect.Descriptor instead.
func (*BatchUpdateRequest) Descriptor() ([]byte, []int) {
	return file_consultation_consultation_proto_rawDescGZIP(), []int{4}
}

func (x *BatchUpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BatchUpdateRequest) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *BatchUpdateRequest) GetConsultationStatus() uint32 {
	if x != nil {
		return x.ConsultationStatus
	}
	return 0
}

func (x *BatchUpdateRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *BatchUpdateRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_consultation_consultation_proto protoreflect.FileDescriptor

var file_consultation_consultation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67,
	0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x15, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x12, 0x2e,
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6d,
	0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x63, 0x0a,
	0x19, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x49, 0x64, 0x22, 0x2d, 0x0a, 0x13, 0x4e, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xac, 0x03, 0x0a, 0x0c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x57, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x65, 0x64, 0x53, 0x65, 0x6e, 0x74, 0x32,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4e, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62,
	0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consultation_consultation_proto_rawDescOnce sync.Once
	file_consultation_consultation_proto_rawDescData = file_consultation_consultation_proto_rawDesc
)

func file_consultation_consultation_proto_rawDescGZIP() []byte {
	file_consultation_consultation_proto_rawDescOnce.Do(func() {
		file_consultation_consultation_proto_rawDescData = protoimpl.X.CompressGZIP(file_consultation_consultation_proto_rawDescData)
	})
	return file_consultation_consultation_proto_rawDescData
}

var file_consultation_consultation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_consultation_consultation_proto_goTypes = []interface{}{
	(*ConsultationRecordDTO)(nil),     // 0: consultation.ConsultationRecordDTO
	(*ConsultationRecordList)(nil),    // 1: consultation.ConsultationRecordList
	(*EffectiveConsultationCond)(nil), // 2: consultation.EffectiveConsultationCond
	(*NsgConsultationCond)(nil),       // 3: consultation.NsgConsultationCond
	(*BatchUpdateRequest)(nil),        // 4: consultation.BatchUpdateRequest
	(*common.Response)(nil),           // 5: common.Response
}
var file_consultation_consultation_proto_depIdxs = []int32{
	0, // 0: consultation.ConsultationRecordList.consultationRecord:type_name -> consultation.ConsultationRecordDTO
	0, // 1: consultation.Consultation.AddConsultation:input_type -> consultation.ConsultationRecordDTO
	2, // 2: consultation.Consultation.GetEffectiveConsultation:input_type -> consultation.EffectiveConsultationCond
	3, // 3: consultation.Consultation.GetNeedSent2GoogleConsultation:input_type -> consultation.NsgConsultationCond
	0, // 4: consultation.Consultation.UpdateConsultation:input_type -> consultation.ConsultationRecordDTO
	4, // 5: consultation.Consultation.BatchUpdateConsultation:input_type -> consultation.BatchUpdateRequest
	5, // 6: consultation.Consultation.AddConsultation:output_type -> common.Response
	5, // 7: consultation.Consultation.GetEffectiveConsultation:output_type -> common.Response
	5, // 8: consultation.Consultation.GetNeedSent2GoogleConsultation:output_type -> common.Response
	5, // 9: consultation.Consultation.UpdateConsultation:output_type -> common.Response
	5, // 10: consultation.Consultation.BatchUpdateConsultation:output_type -> common.Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_consultation_consultation_proto_init() }
func file_consultation_consultation_proto_init() {
	if File_consultation_consultation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consultation_consultation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultationRecordDTO); i {
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
		file_consultation_consultation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultationRecordList); i {
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
		file_consultation_consultation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectiveConsultationCond); i {
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
		file_consultation_consultation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NsgConsultationCond); i {
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
		file_consultation_consultation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpdateRequest); i {
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
			RawDescriptor: file_consultation_consultation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consultation_consultation_proto_goTypes,
		DependencyIndexes: file_consultation_consultation_proto_depIdxs,
		MessageInfos:      file_consultation_consultation_proto_msgTypes,
	}.Build()
	File_consultation_consultation_proto = out.File
	file_consultation_consultation_proto_rawDesc = nil
	file_consultation_consultation_proto_goTypes = nil
	file_consultation_consultation_proto_depIdxs = nil
}
