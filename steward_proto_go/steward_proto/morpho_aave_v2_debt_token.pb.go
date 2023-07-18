//
// Protos for function calls to the Morpho Aave V2 Debt Token adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: morpho_aave_v2_debt_token.proto

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

// Represents call data for the Morpho Aave V2 Debt Token adaptor.
type MorphoAaveV2DebtTokenAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*MorphoAaveV2DebtTokenAdaptorV1_RevokeApproval
	//	*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho_
	//	*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt_
	Function isMorphoAaveV2DebtTokenAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *MorphoAaveV2DebtTokenAdaptorV1) Reset() {
	*x = MorphoAaveV2DebtTokenAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2DebtTokenAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2DebtTokenAdaptorV1) ProtoMessage() {}

func (x *MorphoAaveV2DebtTokenAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2DebtTokenAdaptorV1.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2DebtTokenAdaptorV1) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_debt_token_proto_rawDescGZIP(), []int{0}
}

func (m *MorphoAaveV2DebtTokenAdaptorV1) GetFunction() isMorphoAaveV2DebtTokenAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *MorphoAaveV2DebtTokenAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*MorphoAaveV2DebtTokenAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *MorphoAaveV2DebtTokenAdaptorV1) GetBorrowFromAaveV2Morpho() *MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho {
	if x, ok := x.GetFunction().(*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho_); ok {
		return x.BorrowFromAaveV2Morpho
	}
	return nil
}

func (x *MorphoAaveV2DebtTokenAdaptorV1) GetRepayAaveV2MorphoDebt() *MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt {
	if x, ok := x.GetFunction().(*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt_); ok {
		return x.RepayAaveV2MorphoDebt
	}
	return nil
}

type isMorphoAaveV2DebtTokenAdaptorV1_Function interface {
	isMorphoAaveV2DebtTokenAdaptorV1_Function()
}

type MorphoAaveV2DebtTokenAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho_ struct {
	// Represents function `borrowFromAaveV2Morpho(address aToken, uint256 amountToBorrow)`
	BorrowFromAaveV2Morpho *MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho `protobuf:"bytes,2,opt,name=borrow_from_aave_v2_morpho,json=borrowFromAaveV2Morpho,proto3,oneof"`
}

type MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt_ struct {
	// Represents function `repayAaveV2MorphoDebt(IAaveToken aToken, uint256 amountToRepay)`
	RepayAaveV2MorphoDebt *MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt `protobuf:"bytes,3,opt,name=repay_aave_v2_morpho_debt,json=repayAaveV2MorphoDebt,proto3,oneof"`
}

func (*MorphoAaveV2DebtTokenAdaptorV1_RevokeApproval) isMorphoAaveV2DebtTokenAdaptorV1_Function() {}

func (*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho_) isMorphoAaveV2DebtTokenAdaptorV1_Function() {
}

func (*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt_) isMorphoAaveV2DebtTokenAdaptorV1_Function() {
}

type MorphoAaveV2DebtTokenAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*MorphoAaveV2DebtTokenAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *MorphoAaveV2DebtTokenAdaptorV1Calls) Reset() {
	*x = MorphoAaveV2DebtTokenAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2DebtTokenAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2DebtTokenAdaptorV1Calls) ProtoMessage() {}

func (x *MorphoAaveV2DebtTokenAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2DebtTokenAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2DebtTokenAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_debt_token_proto_rawDescGZIP(), []int{1}
}

func (x *MorphoAaveV2DebtTokenAdaptorV1Calls) GetCalls() []*MorphoAaveV2DebtTokenAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows strategists to borrow assets from Aave.
//
// Represents function `borrowFromAaveV2Morpho(address aToken, uint256 amountToBorrow)`
type MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the Aave V2 aToken to borrow.
	AToken string `protobuf:"bytes,1,opt,name=a_token,json=aToken,proto3" json:"a_token,omitempty"`
	// The amount of the asset to borrow.
	AmountToBorrow string `protobuf:"bytes,2,opt,name=amount_to_borrow,json=amountToBorrow,proto3" json:"amount_to_borrow,omitempty"`
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho) Reset() {
	*x = MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho) ProtoMessage() {}

func (x *MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_debt_token_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho) GetAToken() string {
	if x != nil {
		return x.AToken
	}
	return ""
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho) GetAmountToBorrow() string {
	if x != nil {
		return x.AmountToBorrow
	}
	return ""
}

//
// Allows strategists to repay loan debt on Aave.
//
// Represents function `repayAaveV2MorphoDebt(IAaveToken aToken, uint256 amountToRepay)`
type MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the Aave V2 aToken to repay.
	AToken string `protobuf:"bytes,1,opt,name=a_token,json=aToken,proto3" json:"a_token,omitempty"`
	// The amount of the asset to repay.
	AmountToRepay string `protobuf:"bytes,2,opt,name=amount_to_repay,json=amountToRepay,proto3" json:"amount_to_repay,omitempty"`
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt) Reset() {
	*x = MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt) ProtoMessage() {}

func (x *MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt) ProtoReflect() protoreflect.Message {
	mi := &file_morpho_aave_v2_debt_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt.ProtoReflect.Descriptor instead.
func (*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt) Descriptor() ([]byte, []int) {
	return file_morpho_aave_v2_debt_token_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt) GetAToken() string {
	if x != nil {
		return x.AToken
	}
	return ""
}

func (x *MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt) GetAmountToRepay() string {
	if x != nil {
		return x.AmountToRepay
	}
	return ""
}

var File_morpho_aave_v2_debt_token_proto protoreflect.FileDescriptor

var file_morpho_aave_v2_debt_token_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32,
	0x5f, 0x64, 0x65, 0x62, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x1a, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x04, 0x0a, 0x1e, 0x4d, 0x6f,
	0x72, 0x70, 0x68, 0x6f, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x45, 0x0a, 0x0f,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x34, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x12, 0x7f, 0x0a, 0x1a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32, 0x5f, 0x6d, 0x6f, 0x72, 0x70, 0x68,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x41, 0x61, 0x76, 0x65, 0x56,
	0x32, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61,
	0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x48, 0x00, 0x52, 0x16, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f,
	0x72, 0x70, 0x68, 0x6f, 0x12, 0x7c, 0x0a, 0x19, 0x72, 0x65, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x61,
	0x76, 0x65, 0x5f, 0x76, 0x32, 0x5f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x5f, 0x64, 0x65, 0x62,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x41, 0x61, 0x76, 0x65, 0x56,
	0x32, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x61, 0x79, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d,
	0x6f, 0x72, 0x70, 0x68, 0x6f, 0x44, 0x65, 0x62, 0x74, 0x48, 0x00, 0x52, 0x15, 0x72, 0x65, 0x70,
	0x61, 0x79, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x44, 0x65,
	0x62, 0x74, 0x1a, 0x5b, 0x0a, 0x16, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d,
	0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x74, 0x6f, 0x5f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x1a,
	0x58, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x61, 0x79, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x4d, 0x6f,
	0x72, 0x70, 0x68, 0x6f, 0x44, 0x65, 0x62, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x70, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x61, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x23, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x41,
	0x61, 0x76, 0x65, 0x56, 0x32, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x05,
	0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x6f, 0x41,
	0x61, 0x76, 0x65, 0x56, 0x32, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10,
	0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_morpho_aave_v2_debt_token_proto_rawDescOnce sync.Once
	file_morpho_aave_v2_debt_token_proto_rawDescData = file_morpho_aave_v2_debt_token_proto_rawDesc
)

func file_morpho_aave_v2_debt_token_proto_rawDescGZIP() []byte {
	file_morpho_aave_v2_debt_token_proto_rawDescOnce.Do(func() {
		file_morpho_aave_v2_debt_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_morpho_aave_v2_debt_token_proto_rawDescData)
	})
	return file_morpho_aave_v2_debt_token_proto_rawDescData
}

var file_morpho_aave_v2_debt_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_morpho_aave_v2_debt_token_proto_goTypes = []interface{}{
	(*MorphoAaveV2DebtTokenAdaptorV1)(nil),                        // 0: steward.v4.MorphoAaveV2DebtTokenAdaptorV1
	(*MorphoAaveV2DebtTokenAdaptorV1Calls)(nil),                   // 1: steward.v4.MorphoAaveV2DebtTokenAdaptorV1Calls
	(*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho)(nil), // 2: steward.v4.MorphoAaveV2DebtTokenAdaptorV1.BorrowFromAaveV2Morpho
	(*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt)(nil),  // 3: steward.v4.MorphoAaveV2DebtTokenAdaptorV1.RepayAaveV2MorphoDebt
	(*RevokeApproval)(nil),                                        // 4: steward.v4.RevokeApproval
}
var file_morpho_aave_v2_debt_token_proto_depIdxs = []int32{
	4, // 0: steward.v4.MorphoAaveV2DebtTokenAdaptorV1.revoke_approval:type_name -> steward.v4.RevokeApproval
	2, // 1: steward.v4.MorphoAaveV2DebtTokenAdaptorV1.borrow_from_aave_v2_morpho:type_name -> steward.v4.MorphoAaveV2DebtTokenAdaptorV1.BorrowFromAaveV2Morpho
	3, // 2: steward.v4.MorphoAaveV2DebtTokenAdaptorV1.repay_aave_v2_morpho_debt:type_name -> steward.v4.MorphoAaveV2DebtTokenAdaptorV1.RepayAaveV2MorphoDebt
	0, // 3: steward.v4.MorphoAaveV2DebtTokenAdaptorV1Calls.calls:type_name -> steward.v4.MorphoAaveV2DebtTokenAdaptorV1
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_morpho_aave_v2_debt_token_proto_init() }
func file_morpho_aave_v2_debt_token_proto_init() {
	if File_morpho_aave_v2_debt_token_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_morpho_aave_v2_debt_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2DebtTokenAdaptorV1); i {
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
		file_morpho_aave_v2_debt_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2DebtTokenAdaptorV1Calls); i {
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
		file_morpho_aave_v2_debt_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho); i {
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
		file_morpho_aave_v2_debt_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt); i {
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
	file_morpho_aave_v2_debt_token_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MorphoAaveV2DebtTokenAdaptorV1_RevokeApproval)(nil),
		(*MorphoAaveV2DebtTokenAdaptorV1_BorrowFromAaveV2Morpho_)(nil),
		(*MorphoAaveV2DebtTokenAdaptorV1_RepayAaveV2MorphoDebt_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_morpho_aave_v2_debt_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_morpho_aave_v2_debt_token_proto_goTypes,
		DependencyIndexes: file_morpho_aave_v2_debt_token_proto_depIdxs,
		MessageInfos:      file_morpho_aave_v2_debt_token_proto_msgTypes,
	}.Build()
	File_morpho_aave_v2_debt_token_proto = out.File
	file_morpho_aave_v2_debt_token_proto_rawDesc = nil
	file_morpho_aave_v2_debt_token_proto_goTypes = nil
	file_morpho_aave_v2_debt_token_proto_depIdxs = nil
}
