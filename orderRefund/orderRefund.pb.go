// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: orderRefund/orderRefund.proto

package orderRefund

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

type OrderRefundDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId        uint64                  `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Remark         string                  `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	CustomerAmount string                  `protobuf:"bytes,5,opt,name=customerAmount,proto3" json:"customerAmount,omitempty"`
	BankAmount     string                  `protobuf:"bytes,6,opt,name=bankAmount,proto3" json:"bankAmount,omitempty"`
	Status         uint32                  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Reason         string                  `protobuf:"bytes,8,opt,name=reason,proto3" json:"reason,omitempty"`
	StaffId        uint64                  `protobuf:"varint,9,opt,name=staffId,proto3" json:"staffId,omitempty"`
	StoreId        uint64                  `protobuf:"varint,10,opt,name=storeId,proto3" json:"storeId,omitempty"`
	FileDto        []*fileResource.FileDTO `protobuf:"bytes,11,rep,name=fileDto,proto3" json:"fileDto,omitempty"`
}

func (x *OrderRefundDto) Reset() {
	*x = OrderRefundDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderRefund_orderRefund_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRefundDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRefundDto) ProtoMessage() {}

func (x *OrderRefundDto) ProtoReflect() protoreflect.Message {
	mi := &file_orderRefund_orderRefund_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRefundDto.ProtoReflect.Descriptor instead.
func (*OrderRefundDto) Descriptor() ([]byte, []int) {
	return file_orderRefund_orderRefund_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRefundDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderRefundDto) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderRefundDto) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *OrderRefundDto) GetCustomerAmount() string {
	if x != nil {
		return x.CustomerAmount
	}
	return ""
}

func (x *OrderRefundDto) GetBankAmount() string {
	if x != nil {
		return x.BankAmount
	}
	return ""
}

func (x *OrderRefundDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OrderRefundDto) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *OrderRefundDto) GetStaffId() uint64 {
	if x != nil {
		return x.StaffId
	}
	return 0
}

func (x *OrderRefundDto) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *OrderRefundDto) GetFileDto() []*fileResource.FileDTO {
	if x != nil {
		return x.FileDto
	}
	return nil
}

type RefundCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status  uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RefundCond) Reset() {
	*x = RefundCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderRefund_orderRefund_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundCond) ProtoMessage() {}

func (x *RefundCond) ProtoReflect() protoreflect.Message {
	mi := &file_orderRefund_orderRefund_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundCond.ProtoReflect.Descriptor instead.
func (*RefundCond) Descriptor() ([]byte, []int) {
	return file_orderRefund_orderRefund_proto_rawDescGZIP(), []int{1}
}

func (x *RefundCond) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *RefundCond) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_orderRefund_orderRefund_proto protoreflect.FileDescriptor

var file_orderRefund_orderRefund_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x1a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d,
	0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x0e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x74, 0x6f, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x54,
	0x4f, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0a, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x80, 0x02, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64,
	0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67,
	0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderRefund_orderRefund_proto_rawDescOnce sync.Once
	file_orderRefund_orderRefund_proto_rawDescData = file_orderRefund_orderRefund_proto_rawDesc
)

func file_orderRefund_orderRefund_proto_rawDescGZIP() []byte {
	file_orderRefund_orderRefund_proto_rawDescOnce.Do(func() {
		file_orderRefund_orderRefund_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderRefund_orderRefund_proto_rawDescData)
	})
	return file_orderRefund_orderRefund_proto_rawDescData
}

var file_orderRefund_orderRefund_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orderRefund_orderRefund_proto_goTypes = []interface{}{
	(*OrderRefundDto)(nil),       // 0: orderRefund.OrderRefundDto
	(*RefundCond)(nil),           // 1: orderRefund.RefundCond
	(*fileResource.FileDTO)(nil), // 2: fileResource.FileDTO
	(*common.Response)(nil),      // 3: common.Response
}
var file_orderRefund_orderRefund_proto_depIdxs = []int32{
	2, // 0: orderRefund.OrderRefundDto.fileDto:type_name -> fileResource.FileDTO
	0, // 1: orderRefund.OrderRefund.Add:input_type -> orderRefund.OrderRefundDto
	0, // 2: orderRefund.OrderRefund.Delete:input_type -> orderRefund.OrderRefundDto
	0, // 3: orderRefund.OrderRefund.Update:input_type -> orderRefund.OrderRefundDto
	1, // 4: orderRefund.OrderRefund.GetOrderRefundByCond:input_type -> orderRefund.RefundCond
	3, // 5: orderRefund.OrderRefund.Add:output_type -> common.Response
	3, // 6: orderRefund.OrderRefund.Delete:output_type -> common.Response
	3, // 7: orderRefund.OrderRefund.Update:output_type -> common.Response
	3, // 8: orderRefund.OrderRefund.GetOrderRefundByCond:output_type -> common.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orderRefund_orderRefund_proto_init() }
func file_orderRefund_orderRefund_proto_init() {
	if File_orderRefund_orderRefund_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderRefund_orderRefund_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRefundDto); i {
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
		file_orderRefund_orderRefund_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundCond); i {
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
			RawDescriptor: file_orderRefund_orderRefund_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderRefund_orderRefund_proto_goTypes,
		DependencyIndexes: file_orderRefund_orderRefund_proto_depIdxs,
		MessageInfos:      file_orderRefund_orderRefund_proto_msgTypes,
	}.Build()
	File_orderRefund_orderRefund_proto = out.File
	file_orderRefund_orderRefund_proto_rawDesc = nil
	file_orderRefund_orderRefund_proto_goTypes = nil
	file_orderRefund_orderRefund_proto_depIdxs = nil
}
