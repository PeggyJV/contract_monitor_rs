// Protos for calling functions on the 0x adapter.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: adaptors/0x/zero_x.proto

package steward_proto

import (
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

type ZeroXAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*ZeroXAdaptorV1_RevokeApproval
	//	*ZeroXAdaptorV1_SwapWith_0X
	Function isZeroXAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *ZeroXAdaptorV1) Reset() {
	*x = ZeroXAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_0x_zero_x_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZeroXAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZeroXAdaptorV1) ProtoMessage() {}

func (x *ZeroXAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_0x_zero_x_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZeroXAdaptorV1.ProtoReflect.Descriptor instead.
func (*ZeroXAdaptorV1) Descriptor() ([]byte, []int) {
	return file_adaptors_0x_zero_x_proto_rawDescGZIP(), []int{0}
}

func (m *ZeroXAdaptorV1) GetFunction() isZeroXAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *ZeroXAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*ZeroXAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *ZeroXAdaptorV1) GetSwapWith_0X() *ZeroXAdaptorV1_SwapWith0X {
	if x, ok := x.GetFunction().(*ZeroXAdaptorV1_SwapWith_0X); ok {
		return x.SwapWith_0X
	}
	return nil
}

type isZeroXAdaptorV1_Function interface {
	isZeroXAdaptorV1_Function()
}

type ZeroXAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type ZeroXAdaptorV1_SwapWith_0X struct {
	// Represents function `swapWith0x(ERC20 tokenIn, ERC20 tokenOut, uint256 amount, bytes memory swapCallData)`
	SwapWith_0X *ZeroXAdaptorV1_SwapWith0X `protobuf:"bytes,2,opt,name=swap_with_0x,json=swapWith0x,proto3,oneof"`
}

func (*ZeroXAdaptorV1_RevokeApproval) isZeroXAdaptorV1_Function() {}

func (*ZeroXAdaptorV1_SwapWith_0X) isZeroXAdaptorV1_Function() {}

type ZeroXAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls *ZeroXAdaptorV1 `protobuf:"bytes,1,opt,name=calls,proto3" json:"calls,omitempty"`
}

func (x *ZeroXAdaptorV1Calls) Reset() {
	*x = ZeroXAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_0x_zero_x_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZeroXAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZeroXAdaptorV1Calls) ProtoMessage() {}

func (x *ZeroXAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_0x_zero_x_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZeroXAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*ZeroXAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_adaptors_0x_zero_x_proto_rawDescGZIP(), []int{1}
}

func (x *ZeroXAdaptorV1Calls) GetCalls() *ZeroXAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows strategists to make ERC20 swaps using 0x.
//
// Represents function `swapWith0x(ERC20 tokenIn, ERC20 tokenOut, uint256 amount, bytes memory swapCallData)`
type ZeroXAdaptorV1_SwapWith0X struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenIn      string `protobuf:"bytes,1,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	TokenOut     string `protobuf:"bytes,2,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
	Amount       string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	SwapCallData []byte `protobuf:"bytes,4,opt,name=swap_call_data,json=swapCallData,proto3" json:"swap_call_data,omitempty"`
}

func (x *ZeroXAdaptorV1_SwapWith0X) Reset() {
	*x = ZeroXAdaptorV1_SwapWith0X{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_0x_zero_x_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZeroXAdaptorV1_SwapWith0X) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZeroXAdaptorV1_SwapWith0X) ProtoMessage() {}

func (x *ZeroXAdaptorV1_SwapWith0X) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_0x_zero_x_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZeroXAdaptorV1_SwapWith0X.ProtoReflect.Descriptor instead.
func (*ZeroXAdaptorV1_SwapWith0X) Descriptor() ([]byte, []int) {
	return file_adaptors_0x_zero_x_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ZeroXAdaptorV1_SwapWith0X) GetTokenIn() string {
	if x != nil {
		return x.TokenIn
	}
	return ""
}

func (x *ZeroXAdaptorV1_SwapWith0X) GetTokenOut() string {
	if x != nil {
		return x.TokenOut
	}
	return ""
}

func (x *ZeroXAdaptorV1_SwapWith0X) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ZeroXAdaptorV1_SwapWith0X) GetSwapCallData() []byte {
	if x != nil {
		return x.SwapCallData
	}
	return nil
}

var File_adaptors_0x_zero_x_proto protoreflect.FileDescriptor

var file_adaptors_0x_zero_x_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x30, 0x78, 0x2f, 0x7a, 0x65,
	0x72, 0x6f, 0x5f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x1a, 0x13, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0e,
	0x5a, 0x65, 0x72, 0x6f, 0x58, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x45,
	0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x49, 0x0a, 0x0c, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x30, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x5a, 0x65, 0x72, 0x6f, 0x58, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68,
	0x30, 0x78, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x30, 0x78,
	0x1a, 0x82, 0x01, 0x0a, 0x0a, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x30, 0x78, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x0e, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x77, 0x61, 0x70, 0x43, 0x61, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x47, 0x0a, 0x13, 0x5a, 0x65, 0x72, 0x6f, 0x58, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x33, 0x2e, 0x5a, 0x65, 0x72, 0x6f, 0x58, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adaptors_0x_zero_x_proto_rawDescOnce sync.Once
	file_adaptors_0x_zero_x_proto_rawDescData = file_adaptors_0x_zero_x_proto_rawDesc
)

func file_adaptors_0x_zero_x_proto_rawDescGZIP() []byte {
	file_adaptors_0x_zero_x_proto_rawDescOnce.Do(func() {
		file_adaptors_0x_zero_x_proto_rawDescData = protoimpl.X.CompressGZIP(file_adaptors_0x_zero_x_proto_rawDescData)
	})
	return file_adaptors_0x_zero_x_proto_rawDescData
}

var file_adaptors_0x_zero_x_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_adaptors_0x_zero_x_proto_goTypes = []interface{}{
	(*ZeroXAdaptorV1)(nil),            // 0: steward.v3.ZeroXAdaptorV1
	(*ZeroXAdaptorV1Calls)(nil),       // 1: steward.v3.ZeroXAdaptorV1Calls
	(*ZeroXAdaptorV1_SwapWith0X)(nil), // 2: steward.v3.ZeroXAdaptorV1.SwapWith0x
	(*RevokeApproval)(nil),            // 3: steward.v3.RevokeApproval
}
var file_adaptors_0x_zero_x_proto_depIdxs = []int32{
	3, // 0: steward.v3.ZeroXAdaptorV1.revoke_approval:type_name -> steward.v3.RevokeApproval
	2, // 1: steward.v3.ZeroXAdaptorV1.swap_with_0x:type_name -> steward.v3.ZeroXAdaptorV1.SwapWith0x
	0, // 2: steward.v3.ZeroXAdaptorV1Calls.calls:type_name -> steward.v3.ZeroXAdaptorV1
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_adaptors_0x_zero_x_proto_init() }
func file_adaptors_0x_zero_x_proto_init() {
	if File_adaptors_0x_zero_x_proto != nil {
		return
	}
	file_adaptors_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_adaptors_0x_zero_x_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZeroXAdaptorV1); i {
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
		file_adaptors_0x_zero_x_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZeroXAdaptorV1Calls); i {
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
		file_adaptors_0x_zero_x_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZeroXAdaptorV1_SwapWith0X); i {
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
	file_adaptors_0x_zero_x_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ZeroXAdaptorV1_RevokeApproval)(nil),
		(*ZeroXAdaptorV1_SwapWith_0X)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adaptors_0x_zero_x_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adaptors_0x_zero_x_proto_goTypes,
		DependencyIndexes: file_adaptors_0x_zero_x_proto_depIdxs,
		MessageInfos:      file_adaptors_0x_zero_x_proto_msgTypes,
	}.Build()
	File_adaptors_0x_zero_x_proto = out.File
	file_adaptors_0x_zero_x_proto_rawDesc = nil
	file_adaptors_0x_zero_x_proto_goTypes = nil
	file_adaptors_0x_zero_x_proto_depIdxs = nil
}
