//
// Protos for function calls to the Morpho Blue Supply adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: morpho_blue_supply.proto

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

// Represents call data for the Morpho Blue Supply adaptor.
type MorphoBlueSupplyAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//
	//	*MorphoBlueSupplyAdaptorV1_RevokeApproval
	//	*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue_
	//	*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue_
	Function isMorphoBlueSupplyAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *MorphoBlueSupplyAdaptorV1) Reset() {
	*x = MorphoBlueSupplyAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_blue_supply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoBlueSupplyAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoBlueSupplyAdaptorV1) ProtoMessage() {}

func (x *MorphoBlueSupplyAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_blue_supply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoBlueSupplyAdaptorV1.ProtoReflect.Descriptor instead.
func (*MorphoBlueSupplyAdaptorV1) Descriptor() ([]byte, []int) {
	return file_morpho_blue_supply_proto_rawDescGZIP(), []int{0}
}

func (m *MorphoBlueSupplyAdaptorV1) GetFunction() isMorphoBlueSupplyAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *MorphoBlueSupplyAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*MorphoBlueSupplyAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *MorphoBlueSupplyAdaptorV1) GetLendToMorphoBlue() *MorphoBlueSupplyAdaptorV1_LendToMorphoBlue {
	if x, ok := x.GetFunction().(*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue_); ok {
		return x.LendToMorphoBlue
	}
	return nil
}

func (x *MorphoBlueSupplyAdaptorV1) GetWithdrawFromMorphoBlue() *MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue {
	if x, ok := x.GetFunction().(*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue_); ok {
		return x.WithdrawFromMorphoBlue
	}
	return nil
}

type isMorphoBlueSupplyAdaptorV1_Function interface {
	isMorphoBlueSupplyAdaptorV1_Function()
}

type MorphoBlueSupplyAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type MorphoBlueSupplyAdaptorV1_LendToMorphoBlue_ struct {
	// Represents function `lendToMorphoBlue(MarketParams memory _market, uint256 _assets)`
	LendToMorphoBlue *MorphoBlueSupplyAdaptorV1_LendToMorphoBlue `protobuf:"bytes,2,opt,name=lend_to_morpho_blue,json=lendToMorphoBlue,proto3,oneof"`
}

type MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue_ struct {
	// Represents function `withdrawFromMorphoBlue(MarketParams memory _market, uint256 _assets)`
	WithdrawFromMorphoBlue *MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue `protobuf:"bytes,3,opt,name=withdraw_from_morpho_blue,json=withdrawFromMorphoBlue,proto3,oneof"`
}

func (*MorphoBlueSupplyAdaptorV1_RevokeApproval) isMorphoBlueSupplyAdaptorV1_Function() {}

func (*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue_) isMorphoBlueSupplyAdaptorV1_Function() {}

func (*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue_) isMorphoBlueSupplyAdaptorV1_Function() {}

type MorphoBlueSupplyAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*MorphoBlueSupplyAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *MorphoBlueSupplyAdaptorV1Calls) Reset() {
	*x = MorphoBlueSupplyAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_blue_supply_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoBlueSupplyAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoBlueSupplyAdaptorV1Calls) ProtoMessage() {}

func (x *MorphoBlueSupplyAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_blue_supply_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoBlueSupplyAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*MorphoBlueSupplyAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_morpho_blue_supply_proto_rawDescGZIP(), []int{1}
}

func (x *MorphoBlueSupplyAdaptorV1Calls) GetCalls() []*MorphoBlueSupplyAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

// Allows strategists to lend a specific amount for an asset on Morpho Blue
//
// Represents function `lendToMorphoBlue(MarketParams memory _market, uint256 _assets)`
type MorphoBlueSupplyAdaptorV1_LendToMorphoBlue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of a Morpho Blue Market
	Market *MarketParams `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	// The amount of the loan token to lend
	Assets string `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
}

func (x *MorphoBlueSupplyAdaptorV1_LendToMorphoBlue) Reset() {
	*x = MorphoBlueSupplyAdaptorV1_LendToMorphoBlue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_blue_supply_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoBlueSupplyAdaptorV1_LendToMorphoBlue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue) ProtoMessage() {}

func (x *MorphoBlueSupplyAdaptorV1_LendToMorphoBlue) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_blue_supply_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoBlueSupplyAdaptorV1_LendToMorphoBlue.ProtoReflect.Descriptor instead.
func (*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue) Descriptor() ([]byte, []int) {
	return file_morpho_blue_supply_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MorphoBlueSupplyAdaptorV1_LendToMorphoBlue) GetMarket() *MarketParams {
	if x != nil {
		return x.Market
	}
	return nil
}

func (x *MorphoBlueSupplyAdaptorV1_LendToMorphoBlue) GetAssets() string {
	if x != nil {
		return x.Assets
	}
	return ""
}

// Allows strategists to withdraw the underlying asset plus interest
//
// Represents function `withdrawFromMorphoBlue(MarketParams memory _market, uint256 _assets)`
type MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of a Morpho Blue Market
	Market *MarketParams `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	// The amount of the loan token to lend
	Assets string `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
}

func (x *MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue) Reset() {
	*x = MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_blue_supply_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue) ProtoMessage() {}

func (x *MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_blue_supply_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue.ProtoReflect.Descriptor instead.
func (*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue) Descriptor() ([]byte, []int) {
	return file_morpho_blue_supply_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue) GetMarket() *MarketParams {
	if x != nil {
		return x.Market
	}
	return nil
}

func (x *MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue) GetAssets() string {
	if x != nil {
		return x.Assets
	}
	return ""
}

var File_morpho_blue_supply_proto protoreflect.FileDescriptor

var file_morpho_blue_supply_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x5f, 0x62, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x94, 0x04, 0x0a, 0x19, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x45,
	0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x67, 0x0a, 0x13, 0x6c, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f,
	0x5f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x5f, 0x62, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e,
	0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x4c, 0x65, 0x6e, 0x64, 0x54, 0x6f,
	0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x10, 0x6c, 0x65,
	0x6e, 0x64, 0x54, 0x6f, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x12, 0x79,
	0x0a, 0x19, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x5f, 0x62, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3c, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d,
	0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x48,
	0x00, 0x52, 0x16, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x4d,
	0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x1a, 0x5c, 0x0a, 0x10, 0x4c, 0x65, 0x6e,
	0x64, 0x54, 0x6f, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x1a, 0x62, 0x0a, 0x16, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75,
	0x65, 0x12, 0x30, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x1e, 0x4d, 0x6f, 0x72, 0x70, 0x68,
	0x6f, 0x42, 0x6c, 0x75, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x6f, 0x72, 0x56, 0x31, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x63, 0x61, 0x6c,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x42, 0x6c, 0x75, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x52,
	0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_morpho_blue_supply_proto_rawDescOnce sync.Once
	file_morpho_blue_supply_proto_rawDescData = file_morpho_blue_supply_proto_rawDesc
)

func file_morpho_blue_supply_proto_rawDescGZIP() []byte {
	file_morpho_blue_supply_proto_rawDescOnce.Do(func() {
		file_morpho_blue_supply_proto_rawDescData = protoimpl.X.CompressGZIP(file_morpho_blue_supply_proto_rawDescData)
	})
	return file_morpho_blue_supply_proto_rawDescData
}

var file_morpho_blue_supply_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_morpho_blue_supply_proto_goTypes = []interface{}{
	(*MorphoBlueSupplyAdaptorV1)(nil),                        // 0: steward.v4.MorphoBlueSupplyAdaptorV1
	(*MorphoBlueSupplyAdaptorV1Calls)(nil),                   // 1: steward.v4.MorphoBlueSupplyAdaptorV1Calls
	(*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue)(nil),       // 2: steward.v4.MorphoBlueSupplyAdaptorV1.LendToMorphoBlue
	(*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue)(nil), // 3: steward.v4.MorphoBlueSupplyAdaptorV1.WithdrawFromMorphoBlue
	(*RevokeApproval)(nil),                                   // 4: steward.v4.RevokeApproval
	(*MarketParams)(nil),                                     // 5: steward.v4.MarketParams
}
var file_morpho_blue_supply_proto_depIdxs = []int32{
	4, // 0: steward.v4.MorphoBlueSupplyAdaptorV1.revoke_approval:type_name -> steward.v4.RevokeApproval
	2, // 1: steward.v4.MorphoBlueSupplyAdaptorV1.lend_to_morpho_blue:type_name -> steward.v4.MorphoBlueSupplyAdaptorV1.LendToMorphoBlue
	3, // 2: steward.v4.MorphoBlueSupplyAdaptorV1.withdraw_from_morpho_blue:type_name -> steward.v4.MorphoBlueSupplyAdaptorV1.WithdrawFromMorphoBlue
	0, // 3: steward.v4.MorphoBlueSupplyAdaptorV1Calls.calls:type_name -> steward.v4.MorphoBlueSupplyAdaptorV1
	5, // 4: steward.v4.MorphoBlueSupplyAdaptorV1.LendToMorphoBlue.market:type_name -> steward.v4.MarketParams
	5, // 5: steward.v4.MorphoBlueSupplyAdaptorV1.WithdrawFromMorphoBlue.market:type_name -> steward.v4.MarketParams
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_morpho_blue_supply_proto_init() }
func file_morpho_blue_supply_proto_init() {
	if File_morpho_blue_supply_proto != nil {
		return
	}
	file_base_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_morpho_blue_supply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoBlueSupplyAdaptorV1); i {
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
		file_morpho_blue_supply_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoBlueSupplyAdaptorV1Calls); i {
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
		file_morpho_blue_supply_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue); i {
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
		file_morpho_blue_supply_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue); i {
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
	file_morpho_blue_supply_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MorphoBlueSupplyAdaptorV1_RevokeApproval)(nil),
		(*MorphoBlueSupplyAdaptorV1_LendToMorphoBlue_)(nil),
		(*MorphoBlueSupplyAdaptorV1_WithdrawFromMorphoBlue_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_morpho_blue_supply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_morpho_blue_supply_proto_goTypes,
		DependencyIndexes: file_morpho_blue_supply_proto_depIdxs,
		MessageInfos:      file_morpho_blue_supply_proto_msgTypes,
	}.Build()
	File_morpho_blue_supply_proto = out.File
	file_morpho_blue_supply_proto_rawDesc = nil
	file_morpho_blue_supply_proto_goTypes = nil
	file_morpho_blue_supply_proto_depIdxs = nil
}
