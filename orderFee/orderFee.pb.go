// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.14.0
// source: orderFee/orderFee.proto

package orderFee

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

type OrderFeeDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StaffId uint64 `protobuf:"varint,3,opt,name=staffId,proto3" json:"staffId,omitempty"`
	OrderId uint64 `protobuf:"varint,4,opt,name=orderId,proto3" json:"orderId,omitempty"`
	FeeType uint32 `protobuf:"varint,5,opt,name=feeType,proto3" json:"feeType,omitempty"`
	Status  uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Amount  string `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *OrderFeeDto) Reset() {
	*x = OrderFeeDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderFee_orderFee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFeeDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFeeDto) ProtoMessage() {}

func (x *OrderFeeDto) ProtoReflect() protoreflect.Message {
	mi := &file_orderFee_orderFee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFeeDto.ProtoReflect.Descriptor instead.
func (*OrderFeeDto) Descriptor() ([]byte, []int) {
	return file_orderFee_orderFee_proto_rawDescGZIP(), []int{0}
}

func (x *OrderFeeDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderFeeDto) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *OrderFeeDto) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderFeeDto) GetFeeType() uint32 {
	if x != nil {
		return x.FeeType
	}
	return 0
}

func (x *OrderFeeDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OrderFeeDto) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type FeeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeeList []*OrderFeeDto `protobuf:"bytes,1,rep,name=feeList,proto3" json:"feeList,omitempty"`
}

func (x *FeeList) Reset() {
	*x = FeeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderFee_orderFee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeList) ProtoMessage() {}

func (x *FeeList) ProtoReflect() protoreflect.Message {
	mi := &file_orderFee_orderFee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeList.ProtoReflect.Descriptor instead.
func (*FeeList) Descriptor() ([]byte, []int) {
	return file_orderFee_orderFee_proto_rawDescGZIP(), []int{1}
}

func (x *FeeList) GetFeeList() []*OrderFeeDto {
	if x != nil {
		return x.FeeList
	}
	return nil
}

type FeeCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	FeeType uint32 `protobuf:"varint,5,opt,name=feeType,proto3" json:"feeType,omitempty"`
}

func (x *FeeCond) Reset() {
	*x = FeeCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderFee_orderFee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeCond) ProtoMessage() {}

func (x *FeeCond) ProtoReflect() protoreflect.Message {
	mi := &file_orderFee_orderFee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeCond.ProtoReflect.Descriptor instead.
func (*FeeCond) Descriptor() ([]byte, []int) {
	return file_orderFee_orderFee_proto_rawDescGZIP(), []int{2}
}

func (x *FeeCond) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *FeeCond) GetFeeType() uint32 {
	if x != nil {
		return x.FeeType
	}
	return 0
}

var File_orderFee_orderFee_proto protoreflect.FileDescriptor

var file_orderFee_orderFee_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b,
	0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x07,
	0x46, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x44, 0x74, 0x6f, 0x52,
	0x07, 0x66, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x43,
	0x6f, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x66, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32, 0xe1, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x65, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x44,
	0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x65, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x43, 0x6f, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x2e,
	0x46, 0x65, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64,
	0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderFee_orderFee_proto_rawDescOnce sync.Once
	file_orderFee_orderFee_proto_rawDescData = file_orderFee_orderFee_proto_rawDesc
)

func file_orderFee_orderFee_proto_rawDescGZIP() []byte {
	file_orderFee_orderFee_proto_rawDescOnce.Do(func() {
		file_orderFee_orderFee_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderFee_orderFee_proto_rawDescData)
	})
	return file_orderFee_orderFee_proto_rawDescData
}

var file_orderFee_orderFee_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orderFee_orderFee_proto_goTypes = []interface{}{
	(*OrderFeeDto)(nil),     // 0: orderFee.OrderFeeDto
	(*FeeList)(nil),         // 1: orderFee.FeeList
	(*FeeCond)(nil),         // 2: orderFee.FeeCond
	(*common.Response)(nil), // 3: common.Response
}
var file_orderFee_orderFee_proto_depIdxs = []int32{
	0, // 0: orderFee.FeeList.feeList:type_name -> orderFee.OrderFeeDto
	0, // 1: orderFee.OrderFee.Add:input_type -> orderFee.OrderFeeDto
	0, // 2: orderFee.OrderFee.Delete:input_type -> orderFee.OrderFeeDto
	0, // 3: orderFee.OrderFee.Update:input_type -> orderFee.OrderFeeDto
	2, // 4: orderFee.OrderFee.GetFeeListByCond:input_type -> orderFee.FeeCond
	3, // 5: orderFee.OrderFee.Add:output_type -> common.Response
	3, // 6: orderFee.OrderFee.Delete:output_type -> common.Response
	3, // 7: orderFee.OrderFee.Update:output_type -> common.Response
	3, // 8: orderFee.OrderFee.GetFeeListByCond:output_type -> common.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orderFee_orderFee_proto_init() }
func file_orderFee_orderFee_proto_init() {
	if File_orderFee_orderFee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderFee_orderFee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFeeDto); i {
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
		file_orderFee_orderFee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeList); i {
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
		file_orderFee_orderFee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeCond); i {
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
			RawDescriptor: file_orderFee_orderFee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderFee_orderFee_proto_goTypes,
		DependencyIndexes: file_orderFee_orderFee_proto_depIdxs,
		MessageInfos:      file_orderFee_orderFee_proto_msgTypes,
	}.Build()
	File_orderFee_orderFee_proto = out.File
	file_orderFee_orderFee_proto_rawDesc = nil
	file_orderFee_orderFee_proto_goTypes = nil
	file_orderFee_orderFee_proto_depIdxs = nil
}
