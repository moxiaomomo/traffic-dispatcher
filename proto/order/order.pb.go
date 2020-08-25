// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: proto/order/order.proto

package order

import (
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId     string `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	SrcGeo      string `protobuf:"bytes,3,opt,name=srcGeo,proto3" json:"srcGeo,omitempty"`
	DestGeo     string `protobuf:"bytes,4,opt,name=destGeo,proto3" json:"destGeo,omitempty"`
	CreateAt    uint64 `protobuf:"varint,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	CancelAt    uint64 `protobuf:"varint,6,opt,name=cancelAt,proto3" json:"cancelAt,omitempty"`
	FinishAt    uint64 `protobuf:"varint,7,opt,name=finishAt,proto3" json:"finishAt,omitempty"`
	CancelRole  int32  `protobuf:"varint,8,opt,name=cancelRole,proto3" json:"cancelRole,omitempty"`
	Cost        uint64 `protobuf:"varint,9,opt,name=cost,proto3" json:"cost,omitempty"`
	PassengerId string `protobuf:"bytes,10,opt,name=passengerId,proto3" json:"passengerId,omitempty"`
	DriverId    string `protobuf:"bytes,11,opt,name=driverId,proto3" json:"driverId,omitempty"`
	Status      int32  `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetSrcGeo() string {
	if x != nil {
		return x.SrcGeo
	}
	return ""
}

func (x *Order) GetDestGeo() string {
	if x != nil {
		return x.DestGeo
	}
	return ""
}

func (x *Order) GetCreateAt() uint64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Order) GetCancelAt() uint64 {
	if x != nil {
		return x.CancelAt
	}
	return 0
}

func (x *Order) GetFinishAt() uint64 {
	if x != nil {
		return x.FinishAt
	}
	return 0
}

func (x *Order) GetCancelRole() int32 {
	if x != nil {
		return x.CancelRole
	}
	return 0
}

func (x *Order) GetCost() uint64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Order) GetPassengerId() string {
	if x != nil {
		return x.PassengerId
	}
	return ""
}

func (x *Order) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *Order) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ReqCreateOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Sign  string `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *ReqCreateOrder) Reset() {
	*x = ReqCreateOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqCreateOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqCreateOrder) ProtoMessage() {}

func (x *ReqCreateOrder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqCreateOrder.ProtoReflect.Descriptor instead.
func (*ReqCreateOrder) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *ReqCreateOrder) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *ReqCreateOrder) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type RespCreateOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Order   *Order `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RespCreateOrder) Reset() {
	*x = RespCreateOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespCreateOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespCreateOrder) ProtoMessage() {}

func (x *RespCreateOrder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespCreateOrder.ProtoReflect.Descriptor instead.
func (*RespCreateOrder) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *RespCreateOrder) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespCreateOrder) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *RespCreateOrder) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_order_order_proto protoreflect.FileDescriptor

var file_proto_order_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x72, 0x63, 0x47, 0x65, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x72, 0x63, 0x47, 0x65, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x6f, 0x6c,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x42, 0x0a,
	0x0e, 0x52, 0x65, 0x71, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x22, 0x5d, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x3b, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x71, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_order_proto_rawDescOnce sync.Once
	file_proto_order_order_proto_rawDescData = file_proto_order_order_proto_rawDesc
)

func file_proto_order_order_proto_rawDescGZIP() []byte {
	file_proto_order_order_proto_rawDescOnce.Do(func() {
		file_proto_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_order_proto_rawDescData)
	})
	return file_proto_order_order_proto_rawDescData
}

var file_proto_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_order_order_proto_goTypes = []interface{}{
	(*Order)(nil),           // 0: order
	(*ReqCreateOrder)(nil),  // 1: ReqCreateOrder
	(*RespCreateOrder)(nil), // 2: RespCreateOrder
}
var file_proto_order_order_proto_depIdxs = []int32{
	0, // 0: ReqCreateOrder.order:type_name -> order
	0, // 1: RespCreateOrder.order:type_name -> order
	1, // 2: Order.CreateOrder:input_type -> ReqCreateOrder
	2, // 3: Order.CreateOrder:output_type -> RespCreateOrder
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_order_order_proto_init() }
func file_proto_order_order_proto_init() {
	if File_proto_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_proto_order_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqCreateOrder); i {
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
		file_proto_order_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespCreateOrder); i {
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
			RawDescriptor: file_proto_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_order_proto_goTypes,
		DependencyIndexes: file_proto_order_order_proto_depIdxs,
		MessageInfos:      file_proto_order_order_proto_msgTypes,
	}.Build()
	File_proto_order_order_proto = out.File
	file_proto_order_order_proto_rawDesc = nil
	file_proto_order_order_proto_goTypes = nil
	file_proto_order_order_proto_depIdxs = nil
}
