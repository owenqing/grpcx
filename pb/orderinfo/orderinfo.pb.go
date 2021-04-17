// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/orderinfo.proto

package orderinfo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OrderInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId *int64 `protobuf:"varint,1,req,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (x *OrderInfoReq) Reset() {
	*x = OrderInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orderinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoReq) ProtoMessage() {}

func (x *OrderInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orderinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoReq.ProtoReflect.Descriptor instead.
func (*OrderInfoReq) Descriptor() ([]byte, []int) {
	return file_proto_orderinfo_proto_rawDescGZIP(), []int{0}
}

func (x *OrderInfoReq) GetOrderId() int64 {
	if x != nil && x.OrderId != nil {
		return *x.OrderId
	}
	return 0
}

type OrderInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId *int64  `protobuf:"varint,1,req,name=order_id,json=orderId" json:"order_id,omitempty"`
	Price   *string `protobuf:"bytes,2,req,name=price" json:"price,omitempty"`
	Desc    *string `protobuf:"bytes,3,req,name=desc" json:"desc,omitempty"`
	UserId  *int64  `protobuf:"varint,4,req,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (x *OrderInfoRsp) Reset() {
	*x = OrderInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orderinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoRsp) ProtoMessage() {}

func (x *OrderInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orderinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoRsp.ProtoReflect.Descriptor instead.
func (*OrderInfoRsp) Descriptor() ([]byte, []int) {
	return file_proto_orderinfo_proto_rawDescGZIP(), []int{1}
}

func (x *OrderInfoRsp) GetOrderId() int64 {
	if x != nil && x.OrderId != nil {
		return *x.OrderId
	}
	return 0
}

func (x *OrderInfoRsp) GetPrice() string {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return ""
}

func (x *OrderInfoRsp) GetDesc() string {
	if x != nil && x.Desc != nil {
		return *x.Desc
	}
	return ""
}

func (x *OrderInfoRsp) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

// 订单
type GetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllReq) Reset() {
	*x = GetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orderinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReq) ProtoMessage() {}

func (x *GetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orderinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReq.ProtoReflect.Descriptor instead.
func (*GetAllReq) Descriptor() ([]byte, []int) {
	return file_proto_orderinfo_proto_rawDescGZIP(), []int{2}
}

var File_proto_orderinfo_proto protoreflect.FileDescriptor

var file_proto_orderinfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x6c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x0b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x32, 0x67, 0x0a,
	0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12,
	0x25, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x73, 0x70, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x66, 0x6f,
}

var (
	file_proto_orderinfo_proto_rawDescOnce sync.Once
	file_proto_orderinfo_proto_rawDescData = file_proto_orderinfo_proto_rawDesc
)

func file_proto_orderinfo_proto_rawDescGZIP() []byte {
	file_proto_orderinfo_proto_rawDescOnce.Do(func() {
		file_proto_orderinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_orderinfo_proto_rawDescData)
	})
	return file_proto_orderinfo_proto_rawDescData
}

var file_proto_orderinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_orderinfo_proto_goTypes = []interface{}{
	(*OrderInfoReq)(nil), // 0: OrderInfoReq
	(*OrderInfoRsp)(nil), // 1: OrderInfoRsp
	(*GetAllReq)(nil),    // 2: GetAllReq
}
var file_proto_orderinfo_proto_depIdxs = []int32{
	0, // 0: OrderInfoService.GetOrderInfo:input_type -> OrderInfoReq
	2, // 1: OrderInfoService.GetAll:input_type -> GetAllReq
	1, // 2: OrderInfoService.GetOrderInfo:output_type -> OrderInfoRsp
	1, // 3: OrderInfoService.GetAll:output_type -> OrderInfoRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_orderinfo_proto_init() }
func file_proto_orderinfo_proto_init() {
	if File_proto_orderinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_orderinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoReq); i {
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
		file_proto_orderinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoRsp); i {
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
		file_proto_orderinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReq); i {
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
			RawDescriptor: file_proto_orderinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_orderinfo_proto_goTypes,
		DependencyIndexes: file_proto_orderinfo_proto_depIdxs,
		MessageInfos:      file_proto_orderinfo_proto_msgTypes,
	}.Build()
	File_proto_orderinfo_proto = out.File
	file_proto_orderinfo_proto_rawDesc = nil
	file_proto_orderinfo_proto_goTypes = nil
	file_proto_orderinfo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderInfoServiceClient is the client API for OrderInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderInfoServiceClient interface {
	GetOrderInfo(ctx context.Context, in *OrderInfoReq, opts ...grpc.CallOption) (*OrderInfoRsp, error)
	// 服务端流模式
	GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (OrderInfoService_GetAllClient, error)
}

type orderInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderInfoServiceClient(cc grpc.ClientConnInterface) OrderInfoServiceClient {
	return &orderInfoServiceClient{cc}
}

func (c *orderInfoServiceClient) GetOrderInfo(ctx context.Context, in *OrderInfoReq, opts ...grpc.CallOption) (*OrderInfoRsp, error) {
	out := new(OrderInfoRsp)
	err := c.cc.Invoke(ctx, "/OrderInfoService/GetOrderInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderInfoServiceClient) GetAll(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (OrderInfoService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderInfoService_serviceDesc.Streams[0], "/OrderInfoService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderInfoServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderInfoService_GetAllClient interface {
	Recv() (*OrderInfoRsp, error)
	grpc.ClientStream
}

type orderInfoServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *orderInfoServiceGetAllClient) Recv() (*OrderInfoRsp, error) {
	m := new(OrderInfoRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderInfoServiceServer is the server API for OrderInfoService service.
type OrderInfoServiceServer interface {
	GetOrderInfo(context.Context, *OrderInfoReq) (*OrderInfoRsp, error)
	// 服务端流模式
	GetAll(*GetAllReq, OrderInfoService_GetAllServer) error
}

// UnimplementedOrderInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderInfoServiceServer struct {
}

func (*UnimplementedOrderInfoServiceServer) GetOrderInfo(context.Context, *OrderInfoReq) (*OrderInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderInfo not implemented")
}
func (*UnimplementedOrderInfoServiceServer) GetAll(*GetAllReq, OrderInfoService_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterOrderInfoServiceServer(s *grpc.Server, srv OrderInfoServiceServer) {
	s.RegisterService(&_OrderInfoService_serviceDesc, srv)
}

func _OrderInfoService_GetOrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderInfoServiceServer).GetOrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderInfoService/GetOrderInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderInfoServiceServer).GetOrderInfo(ctx, req.(*OrderInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderInfoService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderInfoServiceServer).GetAll(m, &orderInfoServiceGetAllServer{stream})
}

type OrderInfoService_GetAllServer interface {
	Send(*OrderInfoRsp) error
	grpc.ServerStream
}

type orderInfoServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *orderInfoServiceGetAllServer) Send(m *OrderInfoRsp) error {
	return x.ServerStream.SendMsg(m)
}

var _OrderInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OrderInfoService",
	HandlerType: (*OrderInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderInfo",
			Handler:    _OrderInfoService_GetOrderInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _OrderInfoService_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/orderinfo.proto",
}
