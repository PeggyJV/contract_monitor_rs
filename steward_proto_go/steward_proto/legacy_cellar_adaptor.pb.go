// Protos for depositing/withdrawing from Legacy Cellars

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: legacy_cellar_adaptor.proto

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

type LegacyCellarAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//
	//	*LegacyCellarAdaptorV1_RevokeApproval
	//	*LegacyCellarAdaptorV1_DepositToCellar_
	//	*LegacyCellarAdaptorV1_WithdrawFromCellar_
	Function isLegacyCellarAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *LegacyCellarAdaptorV1) Reset() {
	*x = LegacyCellarAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_cellar_adaptor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyCellarAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyCellarAdaptorV1) ProtoMessage() {}

func (x *LegacyCellarAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_cellar_adaptor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyCellarAdaptorV1.ProtoReflect.Descriptor instead.
func (*LegacyCellarAdaptorV1) Descriptor() ([]byte, []int) {
	return file_legacy_cellar_adaptor_proto_rawDescGZIP(), []int{0}
}

func (m *LegacyCellarAdaptorV1) GetFunction() isLegacyCellarAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *LegacyCellarAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*LegacyCellarAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *LegacyCellarAdaptorV1) GetDepositToCellar() *LegacyCellarAdaptorV1_DepositToCellar {
	if x, ok := x.GetFunction().(*LegacyCellarAdaptorV1_DepositToCellar_); ok {
		return x.DepositToCellar
	}
	return nil
}

func (x *LegacyCellarAdaptorV1) GetWithdrawFromCellar() *LegacyCellarAdaptorV1_WithdrawFromCellar {
	if x, ok := x.GetFunction().(*LegacyCellarAdaptorV1_WithdrawFromCellar_); ok {
		return x.WithdrawFromCellar
	}
	return nil
}

type isLegacyCellarAdaptorV1_Function interface {
	isLegacyCellarAdaptorV1_Function()
}

type LegacyCellarAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type LegacyCellarAdaptorV1_DepositToCellar_ struct {
	// Represents function `depositToCellar(Cellar cellar, uint256 assets, address oracle)`
	DepositToCellar *LegacyCellarAdaptorV1_DepositToCellar `protobuf:"bytes,2,opt,name=depositToCellar,proto3,oneof"`
}

type LegacyCellarAdaptorV1_WithdrawFromCellar_ struct {
	// Represents function `withdrawFromCellar(Cellar cellar, uint256 assets, address oracle)`
	WithdrawFromCellar *LegacyCellarAdaptorV1_WithdrawFromCellar `protobuf:"bytes,3,opt,name=withdrawFromCellar,proto3,oneof"`
}

func (*LegacyCellarAdaptorV1_RevokeApproval) isLegacyCellarAdaptorV1_Function() {}

func (*LegacyCellarAdaptorV1_DepositToCellar_) isLegacyCellarAdaptorV1_Function() {}

func (*LegacyCellarAdaptorV1_WithdrawFromCellar_) isLegacyCellarAdaptorV1_Function() {}

type LegacyCellarAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*LegacyCellarAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *LegacyCellarAdaptorV1Calls) Reset() {
	*x = LegacyCellarAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_cellar_adaptor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyCellarAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyCellarAdaptorV1Calls) ProtoMessage() {}

func (x *LegacyCellarAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_cellar_adaptor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyCellarAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*LegacyCellarAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_legacy_cellar_adaptor_proto_rawDescGZIP(), []int{1}
}

func (x *LegacyCellarAdaptorV1Calls) GetCalls() []*LegacyCellarAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

// Allows strategists to deposit into Cellar positions.
//
// Represents function `depositToCellar(Cellar cellar, uint256 assets, address oracle)`
type LegacyCellarAdaptorV1_DepositToCellar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cellar string `protobuf:"bytes,1,opt,name=cellar,proto3" json:"cellar,omitempty"`
	Assets string `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
	Oracle string `protobuf:"bytes,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
}

func (x *LegacyCellarAdaptorV1_DepositToCellar) Reset() {
	*x = LegacyCellarAdaptorV1_DepositToCellar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_cellar_adaptor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyCellarAdaptorV1_DepositToCellar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyCellarAdaptorV1_DepositToCellar) ProtoMessage() {}

func (x *LegacyCellarAdaptorV1_DepositToCellar) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_cellar_adaptor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyCellarAdaptorV1_DepositToCellar.ProtoReflect.Descriptor instead.
func (*LegacyCellarAdaptorV1_DepositToCellar) Descriptor() ([]byte, []int) {
	return file_legacy_cellar_adaptor_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LegacyCellarAdaptorV1_DepositToCellar) GetCellar() string {
	if x != nil {
		return x.Cellar
	}
	return ""
}

func (x *LegacyCellarAdaptorV1_DepositToCellar) GetAssets() string {
	if x != nil {
		return x.Assets
	}
	return ""
}

func (x *LegacyCellarAdaptorV1_DepositToCellar) GetOracle() string {
	if x != nil {
		return x.Oracle
	}
	return ""
}

// Allows strategists to withdraw from Cellar positions.
//
// Represents function `withdrawFromCellar(Cellar cellar, uint256 assets, address oracle)`
type LegacyCellarAdaptorV1_WithdrawFromCellar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cellar string `protobuf:"bytes,1,opt,name=cellar,proto3" json:"cellar,omitempty"`
	Assets string `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
	Oracle string `protobuf:"bytes,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
}

func (x *LegacyCellarAdaptorV1_WithdrawFromCellar) Reset() {
	*x = LegacyCellarAdaptorV1_WithdrawFromCellar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_cellar_adaptor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LegacyCellarAdaptorV1_WithdrawFromCellar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LegacyCellarAdaptorV1_WithdrawFromCellar) ProtoMessage() {}

func (x *LegacyCellarAdaptorV1_WithdrawFromCellar) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_cellar_adaptor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LegacyCellarAdaptorV1_WithdrawFromCellar.ProtoReflect.Descriptor instead.
func (*LegacyCellarAdaptorV1_WithdrawFromCellar) Descriptor() ([]byte, []int) {
	return file_legacy_cellar_adaptor_proto_rawDescGZIP(), []int{0, 1}
}

func (x *LegacyCellarAdaptorV1_WithdrawFromCellar) GetCellar() string {
	if x != nil {
		return x.Cellar
	}
	return ""
}

func (x *LegacyCellarAdaptorV1_WithdrawFromCellar) GetAssets() string {
	if x != nil {
		return x.Assets
	}
	return ""
}

func (x *LegacyCellarAdaptorV1_WithdrawFromCellar) GetOracle() string {
	if x != nil {
		return x.Oracle
	}
	return ""
}

var File_legacy_cellar_adaptor_proto protoreflect.FileDescriptor

var file_legacy_cellar_adaptor_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x5f,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x03, 0x0a, 0x15, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x12,
	0x45, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x5d, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x54, 0x6f, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72,
	0x56, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x43, 0x65, 0x6c, 0x6c,
	0x61, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x43,
	0x65, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x66, 0x0a, 0x12, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x6f, 0x72, 0x56, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x72, 0x6f,
	0x6d, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x48, 0x00, 0x52, 0x12, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x1a, 0x59, 0x0a,
	0x0f, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x1a, 0x5c, 0x0a, 0x12, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x1a, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x65, 0x6c, 0x6c,
	0x61, 0x72, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x43, 0x61, 0x6c, 0x6c, 0x73,
	0x12, 0x37, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x43, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72,
	0x56, 0x31, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_legacy_cellar_adaptor_proto_rawDescOnce sync.Once
	file_legacy_cellar_adaptor_proto_rawDescData = file_legacy_cellar_adaptor_proto_rawDesc
)

func file_legacy_cellar_adaptor_proto_rawDescGZIP() []byte {
	file_legacy_cellar_adaptor_proto_rawDescOnce.Do(func() {
		file_legacy_cellar_adaptor_proto_rawDescData = protoimpl.X.CompressGZIP(file_legacy_cellar_adaptor_proto_rawDescData)
	})
	return file_legacy_cellar_adaptor_proto_rawDescData
}

var file_legacy_cellar_adaptor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_legacy_cellar_adaptor_proto_goTypes = []interface{}{
	(*LegacyCellarAdaptorV1)(nil),                    // 0: steward.v4.LegacyCellarAdaptorV1
	(*LegacyCellarAdaptorV1Calls)(nil),               // 1: steward.v4.LegacyCellarAdaptorV1Calls
	(*LegacyCellarAdaptorV1_DepositToCellar)(nil),    // 2: steward.v4.LegacyCellarAdaptorV1.DepositToCellar
	(*LegacyCellarAdaptorV1_WithdrawFromCellar)(nil), // 3: steward.v4.LegacyCellarAdaptorV1.WithdrawFromCellar
	(*RevokeApproval)(nil),                           // 4: steward.v4.RevokeApproval
}
var file_legacy_cellar_adaptor_proto_depIdxs = []int32{
	4, // 0: steward.v4.LegacyCellarAdaptorV1.revoke_approval:type_name -> steward.v4.RevokeApproval
	2, // 1: steward.v4.LegacyCellarAdaptorV1.depositToCellar:type_name -> steward.v4.LegacyCellarAdaptorV1.DepositToCellar
	3, // 2: steward.v4.LegacyCellarAdaptorV1.withdrawFromCellar:type_name -> steward.v4.LegacyCellarAdaptorV1.WithdrawFromCellar
	0, // 3: steward.v4.LegacyCellarAdaptorV1Calls.calls:type_name -> steward.v4.LegacyCellarAdaptorV1
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_legacy_cellar_adaptor_proto_init() }
func file_legacy_cellar_adaptor_proto_init() {
	if File_legacy_cellar_adaptor_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_legacy_cellar_adaptor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyCellarAdaptorV1); i {
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
		file_legacy_cellar_adaptor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyCellarAdaptorV1Calls); i {
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
		file_legacy_cellar_adaptor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyCellarAdaptorV1_DepositToCellar); i {
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
		file_legacy_cellar_adaptor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LegacyCellarAdaptorV1_WithdrawFromCellar); i {
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
	file_legacy_cellar_adaptor_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LegacyCellarAdaptorV1_RevokeApproval)(nil),
		(*LegacyCellarAdaptorV1_DepositToCellar_)(nil),
		(*LegacyCellarAdaptorV1_WithdrawFromCellar_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_legacy_cellar_adaptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_legacy_cellar_adaptor_proto_goTypes,
		DependencyIndexes: file_legacy_cellar_adaptor_proto_depIdxs,
		MessageInfos:      file_legacy_cellar_adaptor_proto_msgTypes,
	}.Build()
	File_legacy_cellar_adaptor_proto = out.File
	file_legacy_cellar_adaptor_proto_rawDesc = nil
	file_legacy_cellar_adaptor_proto_goTypes = nil
	file_legacy_cellar_adaptor_proto_depIdxs = nil
}
