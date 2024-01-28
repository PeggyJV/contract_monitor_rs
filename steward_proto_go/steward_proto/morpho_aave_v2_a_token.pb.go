//
// Protos for function calls to the Morpho Aave V2 AToken adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: morpho_aave_v2_a_token.proto

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

// Represents call data for the Morpho Aave V2 AToken adaptor.
type MorphoAaveV2ATokenAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*MorphoAaveV2ATokenAdaptorV1_RevokeApproval
	//	*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho_
	//	*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho_
	//	*MorphoAaveV2ATokenAdaptorV1_Claim
	Function isMorphoAaveV2ATokenAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *MorphoAaveV2ATokenAdaptorV1) Reset() {
	*x = MorphoAaveV2ATokenAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_a_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2ATokenAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2ATokenAdaptorV1) ProtoMessage() {}

func (x *MorphoAaveV2ATokenAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_a_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2ATokenAdaptorV1.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2ATokenAdaptorV1) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_a_token_proto_rawDescGZIP(), []int{0}
}

func (m *MorphoAaveV2ATokenAdaptorV1) GetFunction() isMorphoAaveV2ATokenAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *MorphoAaveV2ATokenAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*MorphoAaveV2ATokenAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *MorphoAaveV2ATokenAdaptorV1) GetDepositToAaveV2Morpho() *MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho {
	if x, ok := x.GetFunction().(*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho_); ok {
		return x.DepositToAaveV2Morpho
	}
	return nil
}

func (x *MorphoAaveV2ATokenAdaptorV1) GetWithdrawFromAaveV2Morpho() *MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho {
	if x, ok := x.GetFunction().(*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho_); ok {
		return x.WithdrawFromAaveV2Morpho
	}
	return nil
}

func (x *MorphoAaveV2ATokenAdaptorV1) GetClaim() *Claim {
	if x, ok := x.GetFunction().(*MorphoAaveV2ATokenAdaptorV1_Claim); ok {
		return x.Claim
	}
	return nil
}

type isMorphoAaveV2ATokenAdaptorV1_Function interface {
	isMorphoAaveV2ATokenAdaptorV1_Function()
}

type MorphoAaveV2ATokenAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho_ struct {
	// Represents function `depositToAaveV2Morpho(IAaveToken aToken, uint256 amountToDeposit)`
	DepositToAaveV2Morpho *MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho `protobuf:"bytes,2,opt,name=deposit_to_aave_v2_morpho,json=depositToAaveV2Morpho,proto3,oneof"`
}

type MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho_ struct {
	// Represents function `withdrawFromAaveV2Morpho(IAaveToken aToken, uint256 amountToWithdraw)`
	WithdrawFromAaveV2Morpho *MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho `protobuf:"bytes,3,opt,name=withdraw_from_aave_v2_morpho,json=withdrawFromAaveV2Morpho,proto3,oneof"`
}

type MorphoAaveV2ATokenAdaptorV1_Claim struct {
	// Represents function `claim(uint256 claimable, bytes32[] memory proof)`
	Claim *Claim `protobuf:"bytes,4,opt,name=claim,proto3,oneof"`
}

func (*MorphoAaveV2ATokenAdaptorV1_RevokeApproval) isMorphoAaveV2ATokenAdaptorV1_Function() {}

func (*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho_) isMorphoAaveV2ATokenAdaptorV1_Function() {}

func (*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho_) isMorphoAaveV2ATokenAdaptorV1_Function() {
}

func (*MorphoAaveV2ATokenAdaptorV1_Claim) isMorphoAaveV2ATokenAdaptorV1_Function() {}

type MorphoAaveV2ATokenAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*MorphoAaveV2ATokenAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *MorphoAaveV2ATokenAdaptorV1Calls) Reset() {
	*x = MorphoAaveV2ATokenAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_a_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2ATokenAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2ATokenAdaptorV1Calls) ProtoMessage() {}

func (x *MorphoAaveV2ATokenAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_a_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2ATokenAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2ATokenAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_a_token_proto_rawDescGZIP(), []int{1}
}

func (x *MorphoAaveV2ATokenAdaptorV1Calls) GetCalls() []*MorphoAaveV2ATokenAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows strategists to lend assets on Morpho.
//
// Represents function `depositToAaveV2Morpho(IAaveToken aToken, uint256 amountToDeposit)`
type MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the Aave V2 aToken to deposit to.
	AToken string `protobuf:"bytes,1,opt,name=a_token,json=aToken,proto3" json:"a_token,omitempty"`
	// The amount of the asset to deposit.
	AmountToDeposit string `protobuf:"bytes,2,opt,name=amount_to_deposit,json=amountToDeposit,proto3" json:"amount_to_deposit,omitempty"`
}

func (x *MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho) Reset() {
	*x = MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_a_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho) ProtoMessage() {}

func (x *MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_a_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_a_token_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho) GetAToken() string {
	if x != nil {
		return x.AToken
	}
	return ""
}

func (x *MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho) GetAmountToDeposit() string {
	if x != nil {
		return x.AmountToDeposit
	}
	return ""
}

//
// Allows strategists to withdraw assets from Morpho.
//
// Represents function `withdrawFromAaveV2Morpho(IAaveToken aToken, uint256 amountToWithdraw)`
type MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the Aave V2 aToken to withdraw from.
	AToken string `protobuf:"bytes,1,opt,name=a_token,json=aToken,proto3" json:"a_token,omitempty"`
	// The amount of the asset to withdraw.
	AmountToWithdraw string `protobuf:"bytes,2,opt,name=amount_to_withdraw,json=amountToWithdraw,proto3" json:"amount_to_withdraw,omitempty"`
}

func (x *MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho) Reset() {
	*x = MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_a_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho) ProtoMessage() {}

func (x *MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_a_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_a_token_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho) GetAToken() string {
	if x != nil {
		return x.AToken
	}
	return ""
}

func (x *MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho) GetAmountToWithdraw() string {
	if x != nil {
		return x.AmountToWithdraw
	}
	return ""
}

var File_morpho_aave_v2_a_token_proto protoreflect.FileDescriptor

var file_morpho_aave_v2_a_token_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32,
	0x5f, 0x61, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x04, 0x0a, 0x1b, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x41, 0x61,
	0x76, 0x65, 0x56, 0x32, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x12, 0x45, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x79, 0x0a, 0x19, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32,
	0x5f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x6f, 0x72, 0x70, 0x68,
	0x6f, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f,
	0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x48, 0x00, 0x52, 0x15,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d,
	0x6f, 0x72, 0x70, 0x68, 0x6f, 0x12, 0x82, 0x01, 0x0a, 0x1c, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32, 0x5f,
	0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f,
	0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x72,
	0x6f, 0x6d, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x48, 0x00,
	0x52, 0x18, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61,
	0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x1a, 0x5c, 0x0a, 0x15, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x54, 0x6f, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x1a, 0x61, 0x0a, 0x18, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46,
	0x72, 0x6f, 0x6d, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x61, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x20, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x41, 0x61, 0x76, 0x65,
	0x56, 0x32, 0x41, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56,
	0x31, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x34, 0x2e, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x41,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x52, 0x05,
	0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_morpho_aave_v2_a_token_proto_rawDescOnce sync.Once
	file_morpho_aave_v2_a_token_proto_rawDescData = file_morpho_aave_v2_a_token_proto_rawDesc
)

func file_morpho_aave_v2_a_token_proto_rawDescGZIP() []byte {
	file_morpho_aave_v2_a_token_proto_rawDescOnce.Do(func() {
		file_morpho_aave_v2_a_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_morpho_aave_v2_a_token_proto_rawDescData)
	})
	return file_morpho_aave_v2_a_token_proto_rawDescData
}

var file_morpho_aave_v2_a_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_morpho_aave_v2_a_token_proto_goTypes = []interface{}{
	(*MorphoAaveV2ATokenAdaptorV1)(nil),                          // 0: steward.v4.MorphoAaveV2ATokenAdaptorV1
	(*MorphoAaveV2ATokenAdaptorV1Calls)(nil),                     // 1: steward.v4.MorphoAaveV2ATokenAdaptorV1Calls
	(*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho)(nil),    // 2: steward.v4.MorphoAaveV2ATokenAdaptorV1.DepositToAaveV2Morpho
	(*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho)(nil), // 3: steward.v4.MorphoAaveV2ATokenAdaptorV1.WithdrawFromAaveV2Morpho
	(*RevokeApproval)(nil),                                       // 4: steward.v4.RevokeApproval
	(*Claim)(nil),                                                // 5: steward.v4.Claim
}
var file_morpho_aave_v2_a_token_proto_depIdxs = []int32{
	4, // 0: steward.v4.MorphoAaveV2ATokenAdaptorV1.revoke_approval:type_name -> steward.v4.RevokeApproval
	2, // 1: steward.v4.MorphoAaveV2ATokenAdaptorV1.deposit_to_aave_v2_morpho:type_name -> steward.v4.MorphoAaveV2ATokenAdaptorV1.DepositToAaveV2Morpho
	3, // 2: steward.v4.MorphoAaveV2ATokenAdaptorV1.withdraw_from_aave_v2_morpho:type_name -> steward.v4.MorphoAaveV2ATokenAdaptorV1.WithdrawFromAaveV2Morpho
	5, // 3: steward.v4.MorphoAaveV2ATokenAdaptorV1.claim:type_name -> steward.v4.Claim
	0, // 4: steward.v4.MorphoAaveV2ATokenAdaptorV1Calls.calls:type_name -> steward.v4.MorphoAaveV2ATokenAdaptorV1
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_morpho_aave_v2_a_token_proto_init() }
func file_morpho_aave_v2_a_token_proto_init() {
	if File_morpho_aave_v2_a_token_proto != nil {
		return
	}
	file_base_proto_init()
	file_morpho_reward_handler_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_morpho_aave_v2_a_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2ATokenAdaptorV1); i {
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
		file_morpho_aave_v2_a_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2ATokenAdaptorV1Calls); i {
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
		file_morpho_aave_v2_a_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho); i {
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
		file_morpho_aave_v2_a_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho); i {
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
	file_morpho_aave_v2_a_token_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MorphoAaveV2ATokenAdaptorV1_RevokeApproval)(nil),
		(*MorphoAaveV2ATokenAdaptorV1_DepositToAaveV2Morpho_)(nil),
		(*MorphoAaveV2ATokenAdaptorV1_WithdrawFromAaveV2Morpho_)(nil),
		(*MorphoAaveV2ATokenAdaptorV1_Claim)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_morpho_aave_v2_a_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_morpho_aave_v2_a_token_proto_goTypes,
		DependencyIndexes: file_morpho_aave_v2_a_token_proto_depIdxs,
		MessageInfos:      file_morpho_aave_v2_a_token_proto_msgTypes,
	}.Build()
	File_morpho_aave_v2_a_token_proto = out.File
	file_morpho_aave_v2_a_token_proto_rawDesc = nil
	file_morpho_aave_v2_a_token_proto_goTypes = nil
	file_morpho_aave_v2_a_token_proto_depIdxs = nil
}
