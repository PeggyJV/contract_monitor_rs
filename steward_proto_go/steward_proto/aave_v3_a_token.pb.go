//
// Protos for function calls to the Aave V3 AToken adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: adaptors/aave/aave_v3_a_token.proto

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

// Represents call data for the Aave AToken adaptor, used to manage lending positions on Aave
type AaveV3ATokenAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*AaveV3ATokenAdaptorV1_RevokeApproval
	//	*AaveV3ATokenAdaptorV1_DepositToAave_
	//	*AaveV3ATokenAdaptorV1_WithdrawFromAave_
	//	*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral_
	//	*AaveV3ATokenAdaptorV1_ChangeEmode
	Function isAaveV3ATokenAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *AaveV3ATokenAdaptorV1) Reset() {
	*x = AaveV3ATokenAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV3ATokenAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV3ATokenAdaptorV1) ProtoMessage() {}

func (x *AaveV3ATokenAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveV3ATokenAdaptorV1.ProtoReflect.Descriptor instead.
func (*AaveV3ATokenAdaptorV1) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_aave_v3_a_token_proto_rawDescGZIP(), []int{0}
}

func (m *AaveV3ATokenAdaptorV1) GetFunction() isAaveV3ATokenAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *AaveV3ATokenAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*AaveV3ATokenAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *AaveV3ATokenAdaptorV1) GetDepositToAave() *AaveV3ATokenAdaptorV1_DepositToAave {
	if x, ok := x.GetFunction().(*AaveV3ATokenAdaptorV1_DepositToAave_); ok {
		return x.DepositToAave
	}
	return nil
}

func (x *AaveV3ATokenAdaptorV1) GetWithdrawFromAave() *AaveV3ATokenAdaptorV1_WithdrawFromAave {
	if x, ok := x.GetFunction().(*AaveV3ATokenAdaptorV1_WithdrawFromAave_); ok {
		return x.WithdrawFromAave
	}
	return nil
}

func (x *AaveV3ATokenAdaptorV1) GetAdjustIsolationModeAssetAsCollateral() *AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral {
	if x, ok := x.GetFunction().(*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral_); ok {
		return x.AdjustIsolationModeAssetAsCollateral
	}
	return nil
}

func (x *AaveV3ATokenAdaptorV1) GetChangeEmode() *AaveV3ATokenAdaptorV1_ChangeEMode {
	if x, ok := x.GetFunction().(*AaveV3ATokenAdaptorV1_ChangeEmode); ok {
		return x.ChangeEmode
	}
	return nil
}

type isAaveV3ATokenAdaptorV1_Function interface {
	isAaveV3ATokenAdaptorV1_Function()
}

type AaveV3ATokenAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,3,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type AaveV3ATokenAdaptorV1_DepositToAave_ struct {
	// Represents function `depositToAave(ERC20 tokenToDeposit, uint256 amountToDeposit)`
	DepositToAave *AaveV3ATokenAdaptorV1_DepositToAave `protobuf:"bytes,4,opt,name=deposit_to_aave,json=depositToAave,proto3,oneof"`
}

type AaveV3ATokenAdaptorV1_WithdrawFromAave_ struct {
	// Represents function `withdrawFromAave(ERC20 tokenToWithdraw, uint256 amountToWithdraw)`
	WithdrawFromAave *AaveV3ATokenAdaptorV1_WithdrawFromAave `protobuf:"bytes,5,opt,name=withdraw_from_aave,json=withdrawFromAave,proto3,oneof"`
}

type AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral_ struct {
	// Represents function `adjustIsolationModeAssetAsCollateral(ERC20 asset, bool useAsCollateral)`
	AdjustIsolationModeAssetAsCollateral *AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral `protobuf:"bytes,6,opt,name=adjust_isolation_mode_asset_as_collateral,json=adjustIsolationModeAssetAsCollateral,proto3,oneof"`
}

type AaveV3ATokenAdaptorV1_ChangeEmode struct {
	// Represents function `changeEMode(uint8 categoryId)`
	ChangeEmode *AaveV3ATokenAdaptorV1_ChangeEMode `protobuf:"bytes,7,opt,name=change_emode,json=changeEmode,proto3,oneof"`
}

func (*AaveV3ATokenAdaptorV1_RevokeApproval) isAaveV3ATokenAdaptorV1_Function() {}

func (*AaveV3ATokenAdaptorV1_DepositToAave_) isAaveV3ATokenAdaptorV1_Function() {}

func (*AaveV3ATokenAdaptorV1_WithdrawFromAave_) isAaveV3ATokenAdaptorV1_Function() {}

func (*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral_) isAaveV3ATokenAdaptorV1_Function() {
}

func (*AaveV3ATokenAdaptorV1_ChangeEmode) isAaveV3ATokenAdaptorV1_Function() {}

type AaveV3ATokenAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*AaveV3ATokenAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *AaveV3ATokenAdaptorV1Calls) Reset() {
	*x = AaveV3ATokenAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV3ATokenAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV3ATokenAdaptorV1Calls) ProtoMessage() {}

func (x *AaveV3ATokenAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveV3ATokenAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*AaveV3ATokenAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_aave_v3_a_token_proto_rawDescGZIP(), []int{1}
}

func (x *AaveV3ATokenAdaptorV1Calls) GetCalls() []*AaveV3ATokenAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows strategists to lend assets on Aave.
//
// Represents function `depositToAave(ERC20 tokenToDeposit, uint256 amountToDeposit)`
type AaveV3ATokenAdaptorV1_DepositToAave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the ERC20 token to deposit
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The amount to deposit
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AaveV3ATokenAdaptorV1_DepositToAave) Reset() {
	*x = AaveV3ATokenAdaptorV1_DepositToAave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV3ATokenAdaptorV1_DepositToAave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV3ATokenAdaptorV1_DepositToAave) ProtoMessage() {}

func (x *AaveV3ATokenAdaptorV1_DepositToAave) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveV3ATokenAdaptorV1_DepositToAave.ProtoReflect.Descriptor instead.
func (*AaveV3ATokenAdaptorV1_DepositToAave) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_aave_v3_a_token_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AaveV3ATokenAdaptorV1_DepositToAave) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AaveV3ATokenAdaptorV1_DepositToAave) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

//
// Allows strategists to withdraw assets from Aave.
//
// Represents function `withdrawFromAave(ERC20 tokenToWithdraw, uint256 amountToWithdraw)`
type AaveV3ATokenAdaptorV1_WithdrawFromAave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the ERC20 token to withdraw
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The amount to withdraw
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AaveV3ATokenAdaptorV1_WithdrawFromAave) Reset() {
	*x = AaveV3ATokenAdaptorV1_WithdrawFromAave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV3ATokenAdaptorV1_WithdrawFromAave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV3ATokenAdaptorV1_WithdrawFromAave) ProtoMessage() {}

func (x *AaveV3ATokenAdaptorV1_WithdrawFromAave) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveV3ATokenAdaptorV1_WithdrawFromAave.ProtoReflect.Descriptor instead.
func (*AaveV3ATokenAdaptorV1_WithdrawFromAave) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_aave_v3_a_token_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AaveV3ATokenAdaptorV1_WithdrawFromAave) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AaveV3ATokenAdaptorV1_WithdrawFromAave) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

//
// Allows strategists to adjust an asset's isolation mode.
//
// Represents function `adjustIsolationModeAssetAsCollateral(ERC20 asset, bool useAsCollateral)`
type AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the ERC20 token
	Asset string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// Whether to use the asset as collateral
	UseAsCollateral bool `protobuf:"varint,2,opt,name=use_as_collateral,json=useAsCollateral,proto3" json:"use_as_collateral,omitempty"`
}

func (x *AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral) Reset() {
	*x = AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral) ProtoMessage() {}

func (x *AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral.ProtoReflect.Descriptor instead.
func (*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_aave_v3_a_token_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral) GetUseAsCollateral() bool {
	if x != nil {
		return x.UseAsCollateral
	}
	return false
}

//
// Allows strategists to enter different EModes.
//
// Represents function `changeEMode(uint8 categoryId)`
type AaveV3ATokenAdaptorV1_ChangeEMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The category ID
	CategoryId uint32 `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *AaveV3ATokenAdaptorV1_ChangeEMode) Reset() {
	*x = AaveV3ATokenAdaptorV1_ChangeEMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV3ATokenAdaptorV1_ChangeEMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV3ATokenAdaptorV1_ChangeEMode) ProtoMessage() {}

func (x *AaveV3ATokenAdaptorV1_ChangeEMode) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_aave_v3_a_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveV3ATokenAdaptorV1_ChangeEMode.ProtoReflect.Descriptor instead.
func (*AaveV3ATokenAdaptorV1_ChangeEMode) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_aave_v3_a_token_proto_rawDescGZIP(), []int{0, 3}
}

func (x *AaveV3ATokenAdaptorV1_ChangeEMode) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

var File_adaptors_aave_aave_v3_a_token_proto protoreflect.FileDescriptor

var file_adaptors_aave_aave_v3_a_token_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x61, 0x76, 0x65, 0x2f,
	0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x33, 0x5f, 0x61, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x33, 0x1a, 0x13, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x06, 0x0a, 0x15, 0x41, 0x61, 0x76, 0x65, 0x56,
	0x33, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31,
	0x12, 0x45, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x59, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x61,
	0x76, 0x65, 0x56, 0x33, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x61, 0x76,
	0x65, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x61,
	0x76, 0x65, 0x12, 0x62, 0x0a, 0x12, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x61, 0x76, 0x65,
	0x56, 0x33, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56,
	0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61,
	0x76, 0x65, 0x48, 0x00, 0x52, 0x10, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x72,
	0x6f, 0x6d, 0x41, 0x61, 0x76, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x29, 0x61, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x5f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74,
	0x65, 0x72, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x73, 0x74, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x56, 0x33, 0x41, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x41, 0x64,
	0x6a, 0x75, 0x73, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x41, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x24, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x73, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x41, 0x73,
	0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x52, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x61,
	0x76, 0x65, 0x56, 0x33, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x4d, 0x6f, 0x64, 0x65, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6d, 0x6f, 0x64, 0x65, 0x1a, 0x3d,
	0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x61, 0x76, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x40, 0x0a,
	0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61, 0x76,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x68, 0x0a, 0x24, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x41, 0x73, 0x43, 0x6f, 0x6c,
	0x6c, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x75, 0x73, 0x65, 0x5f, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x41, 0x73, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x1a, 0x2e, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x45, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x1a, 0x41, 0x61, 0x76, 0x65, 0x56, 0x33, 0x41,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x43, 0x61,
	0x6c, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x61, 0x76, 0x65, 0x56, 0x33, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x6f, 0x72, 0x56, 0x31, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e,
	0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adaptors_aave_aave_v3_a_token_proto_rawDescOnce sync.Once
	file_adaptors_aave_aave_v3_a_token_proto_rawDescData = file_adaptors_aave_aave_v3_a_token_proto_rawDesc
)

func file_adaptors_aave_aave_v3_a_token_proto_rawDescGZIP() []byte {
	file_adaptors_aave_aave_v3_a_token_proto_rawDescOnce.Do(func() {
		file_adaptors_aave_aave_v3_a_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_adaptors_aave_aave_v3_a_token_proto_rawDescData)
	})
	return file_adaptors_aave_aave_v3_a_token_proto_rawDescData
}

var file_adaptors_aave_aave_v3_a_token_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_adaptors_aave_aave_v3_a_token_proto_goTypes = []interface{}{
	(*AaveV3ATokenAdaptorV1)(nil),                                      // 0: steward.v3.AaveV3ATokenAdaptorV1
	(*AaveV3ATokenAdaptorV1Calls)(nil),                                 // 1: steward.v3.AaveV3ATokenAdaptorV1Calls
	(*AaveV3ATokenAdaptorV1_DepositToAave)(nil),                        // 2: steward.v3.AaveV3ATokenAdaptorV1.DepositToAave
	(*AaveV3ATokenAdaptorV1_WithdrawFromAave)(nil),                     // 3: steward.v3.AaveV3ATokenAdaptorV1.WithdrawFromAave
	(*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral)(nil), // 4: steward.v3.AaveV3ATokenAdaptorV1.AdjustIsolationModeAssetAsCollateral
	(*AaveV3ATokenAdaptorV1_ChangeEMode)(nil),                          // 5: steward.v3.AaveV3ATokenAdaptorV1.ChangeEMode
	(*RevokeApproval)(nil),                                             // 6: steward.v3.RevokeApproval
}
var file_adaptors_aave_aave_v3_a_token_proto_depIdxs = []int32{
	6, // 0: steward.v3.AaveV3ATokenAdaptorV1.revoke_approval:type_name -> steward.v3.RevokeApproval
	2, // 1: steward.v3.AaveV3ATokenAdaptorV1.deposit_to_aave:type_name -> steward.v3.AaveV3ATokenAdaptorV1.DepositToAave
	3, // 2: steward.v3.AaveV3ATokenAdaptorV1.withdraw_from_aave:type_name -> steward.v3.AaveV3ATokenAdaptorV1.WithdrawFromAave
	4, // 3: steward.v3.AaveV3ATokenAdaptorV1.adjust_isolation_mode_asset_as_collateral:type_name -> steward.v3.AaveV3ATokenAdaptorV1.AdjustIsolationModeAssetAsCollateral
	5, // 4: steward.v3.AaveV3ATokenAdaptorV1.change_emode:type_name -> steward.v3.AaveV3ATokenAdaptorV1.ChangeEMode
	0, // 5: steward.v3.AaveV3ATokenAdaptorV1Calls.calls:type_name -> steward.v3.AaveV3ATokenAdaptorV1
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_adaptors_aave_aave_v3_a_token_proto_init() }
func file_adaptors_aave_aave_v3_a_token_proto_init() {
	if File_adaptors_aave_aave_v3_a_token_proto != nil {
		return
	}
	file_adaptors_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_adaptors_aave_aave_v3_a_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveV3ATokenAdaptorV1); i {
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
		file_adaptors_aave_aave_v3_a_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveV3ATokenAdaptorV1Calls); i {
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
		file_adaptors_aave_aave_v3_a_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveV3ATokenAdaptorV1_DepositToAave); i {
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
		file_adaptors_aave_aave_v3_a_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveV3ATokenAdaptorV1_WithdrawFromAave); i {
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
		file_adaptors_aave_aave_v3_a_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral); i {
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
		file_adaptors_aave_aave_v3_a_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveV3ATokenAdaptorV1_ChangeEMode); i {
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
	file_adaptors_aave_aave_v3_a_token_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AaveV3ATokenAdaptorV1_RevokeApproval)(nil),
		(*AaveV3ATokenAdaptorV1_DepositToAave_)(nil),
		(*AaveV3ATokenAdaptorV1_WithdrawFromAave_)(nil),
		(*AaveV3ATokenAdaptorV1_AdjustIsolationModeAssetAsCollateral_)(nil),
		(*AaveV3ATokenAdaptorV1_ChangeEmode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adaptors_aave_aave_v3_a_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adaptors_aave_aave_v3_a_token_proto_goTypes,
		DependencyIndexes: file_adaptors_aave_aave_v3_a_token_proto_depIdxs,
		MessageInfos:      file_adaptors_aave_aave_v3_a_token_proto_msgTypes,
	}.Build()
	File_adaptors_aave_aave_v3_a_token_proto = out.File
	file_adaptors_aave_aave_v3_a_token_proto_rawDesc = nil
	file_adaptors_aave_aave_v3_a_token_proto_goTypes = nil
	file_adaptors_aave_aave_v3_a_token_proto_depIdxs = nil
}
