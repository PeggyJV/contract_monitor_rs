//
// Protos for function calls to the Aave Debt Token adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: adaptors/aave/debt_token.proto

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

// Represents call data for the Aave Debt Token adaptor, used for borrowing and repaying debt on Aave.
type AaveDebtTokenAdaptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*AaveDebtTokenAdaptor_Swap
	//	*AaveDebtTokenAdaptor_OracleSwap
	//	*AaveDebtTokenAdaptor_RevokeApproval
	//	*AaveDebtTokenAdaptor_BorrowFromAave_
	//	*AaveDebtTokenAdaptor_RepayAaveDebt_
	//	*AaveDebtTokenAdaptor_SwapAndRepay_
	Function isAaveDebtTokenAdaptor_Function `protobuf_oneof:"function"`
}

func (x *AaveDebtTokenAdaptor) Reset() {
	*x = AaveDebtTokenAdaptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_debt_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveDebtTokenAdaptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveDebtTokenAdaptor) ProtoMessage() {}

func (x *AaveDebtTokenAdaptor) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_debt_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveDebtTokenAdaptor.ProtoReflect.Descriptor instead.
func (*AaveDebtTokenAdaptor) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_debt_token_proto_rawDescGZIP(), []int{0}
}

func (m *AaveDebtTokenAdaptor) GetFunction() isAaveDebtTokenAdaptor_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *AaveDebtTokenAdaptor) GetSwap() *Swap {
	if x, ok := x.GetFunction().(*AaveDebtTokenAdaptor_Swap); ok {
		return x.Swap
	}
	return nil
}

func (x *AaveDebtTokenAdaptor) GetOracleSwap() *OracleSwap {
	if x, ok := x.GetFunction().(*AaveDebtTokenAdaptor_OracleSwap); ok {
		return x.OracleSwap
	}
	return nil
}

func (x *AaveDebtTokenAdaptor) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*AaveDebtTokenAdaptor_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *AaveDebtTokenAdaptor) GetBorrowFromAave() *AaveDebtTokenAdaptor_BorrowFromAave {
	if x, ok := x.GetFunction().(*AaveDebtTokenAdaptor_BorrowFromAave_); ok {
		return x.BorrowFromAave
	}
	return nil
}

func (x *AaveDebtTokenAdaptor) GetRepayAaveDebt() *AaveDebtTokenAdaptor_RepayAaveDebt {
	if x, ok := x.GetFunction().(*AaveDebtTokenAdaptor_RepayAaveDebt_); ok {
		return x.RepayAaveDebt
	}
	return nil
}

func (x *AaveDebtTokenAdaptor) GetSwapAndRepay() *AaveDebtTokenAdaptor_SwapAndRepay {
	if x, ok := x.GetFunction().(*AaveDebtTokenAdaptor_SwapAndRepay_); ok {
		return x.SwapAndRepay
	}
	return nil
}

type isAaveDebtTokenAdaptor_Function interface {
	isAaveDebtTokenAdaptor_Function()
}

type AaveDebtTokenAdaptor_Swap struct {
	// Represents function `swap(ERC20 assetIn, ERC20 assetOut, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params)`
	Swap *Swap `protobuf:"bytes,1,opt,name=swap,proto3,oneof"`
}

type AaveDebtTokenAdaptor_OracleSwap struct {
	// Represents function `oracleSwap(ERC20 assetIn, ERC20 assetOut, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params, uint64 slippage)`
	OracleSwap *OracleSwap `protobuf:"bytes,2,opt,name=oracle_swap,json=oracleSwap,proto3,oneof"`
}

type AaveDebtTokenAdaptor_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,3,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type AaveDebtTokenAdaptor_BorrowFromAave_ struct {
	// Represents function `borrowFromAave(ERC20 debtTokenToBorrow, uint256 amountToBorrow)`
	BorrowFromAave *AaveDebtTokenAdaptor_BorrowFromAave `protobuf:"bytes,4,opt,name=borrow_from_aave,json=borrowFromAave,proto3,oneof"`
}

type AaveDebtTokenAdaptor_RepayAaveDebt_ struct {
	// Represents function `repayAaveDebt(ERC20 tokenToRepay, uint256 amountToRepay)`
	RepayAaveDebt *AaveDebtTokenAdaptor_RepayAaveDebt `protobuf:"bytes,5,opt,name=repay_aave_debt,json=repayAaveDebt,proto3,oneof"`
}

type AaveDebtTokenAdaptor_SwapAndRepay_ struct {
	// Represents function `swapAndRepay(ERC20 tokenIn, ERC20 tokenToRepay, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params)`
	SwapAndRepay *AaveDebtTokenAdaptor_SwapAndRepay `protobuf:"bytes,6,opt,name=swap_and_repay,json=swapAndRepay,proto3,oneof"`
}

func (*AaveDebtTokenAdaptor_Swap) isAaveDebtTokenAdaptor_Function() {}

func (*AaveDebtTokenAdaptor_OracleSwap) isAaveDebtTokenAdaptor_Function() {}

func (*AaveDebtTokenAdaptor_RevokeApproval) isAaveDebtTokenAdaptor_Function() {}

func (*AaveDebtTokenAdaptor_BorrowFromAave_) isAaveDebtTokenAdaptor_Function() {}

func (*AaveDebtTokenAdaptor_RepayAaveDebt_) isAaveDebtTokenAdaptor_Function() {}

func (*AaveDebtTokenAdaptor_SwapAndRepay_) isAaveDebtTokenAdaptor_Function() {}

type AaveDebtTokenAdaptorCalls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*AaveDebtTokenAdaptor `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *AaveDebtTokenAdaptorCalls) Reset() {
	*x = AaveDebtTokenAdaptorCalls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_debt_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveDebtTokenAdaptorCalls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveDebtTokenAdaptorCalls) ProtoMessage() {}

func (x *AaveDebtTokenAdaptorCalls) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_debt_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveDebtTokenAdaptorCalls.ProtoReflect.Descriptor instead.
func (*AaveDebtTokenAdaptorCalls) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_debt_token_proto_rawDescGZIP(), []int{1}
}

func (x *AaveDebtTokenAdaptorCalls) GetCalls() []*AaveDebtTokenAdaptor {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows strategists to borrow assets from Aave.
//
// Represents function `depositToAave(ERC20 tokenToDeposit, uint256 amountToDeposit)`
type AaveDebtTokenAdaptor_BorrowFromAave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the ERC20 token to borrow
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The amount to borrow
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AaveDebtTokenAdaptor_BorrowFromAave) Reset() {
	*x = AaveDebtTokenAdaptor_BorrowFromAave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_debt_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveDebtTokenAdaptor_BorrowFromAave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveDebtTokenAdaptor_BorrowFromAave) ProtoMessage() {}

func (x *AaveDebtTokenAdaptor_BorrowFromAave) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_debt_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveDebtTokenAdaptor_BorrowFromAave.ProtoReflect.Descriptor instead.
func (*AaveDebtTokenAdaptor_BorrowFromAave) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_debt_token_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AaveDebtTokenAdaptor_BorrowFromAave) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AaveDebtTokenAdaptor_BorrowFromAave) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

//
// Allows strategists to repay loan debt on Aave.
//
// Represents function `repayAaveDebt(ERC20 tokenToRepay, uint256 amountToRepay)`
type AaveDebtTokenAdaptor_RepayAaveDebt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the ERC20 token to repay
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The amount to repay
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AaveDebtTokenAdaptor_RepayAaveDebt) Reset() {
	*x = AaveDebtTokenAdaptor_RepayAaveDebt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_debt_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveDebtTokenAdaptor_RepayAaveDebt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveDebtTokenAdaptor_RepayAaveDebt) ProtoMessage() {}

func (x *AaveDebtTokenAdaptor_RepayAaveDebt) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_debt_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveDebtTokenAdaptor_RepayAaveDebt.ProtoReflect.Descriptor instead.
func (*AaveDebtTokenAdaptor_RepayAaveDebt) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_debt_token_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AaveDebtTokenAdaptor_RepayAaveDebt) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AaveDebtTokenAdaptor_RepayAaveDebt) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

//
// Allows strategists to swap assets and repay loans in one call.
//
// Represents function `swapAndRepay(ERC20 tokenIn, ERC20 tokenToRepay, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params)`
type AaveDebtTokenAdaptor_SwapAndRepay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the token to swap from
	TokenIn string `protobuf:"bytes,1,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// The address of the token to swap to and repay with
	TokenToRepay string `protobuf:"bytes,2,opt,name=token_to_repay,json=tokenToRepay,proto3" json:"token_to_repay,omitempty"`
	// The amount to swap
	AmountIn string `protobuf:"bytes,3,opt,name=amount_in,json=amountIn,proto3" json:"amount_in,omitempty"`
	// The exchange to make the swap on
	Exchange Exchange `protobuf:"varint,4,opt,name=exchange,proto3,enum=steward.v2.Exchange" json:"exchange,omitempty"`
	// The parameters for the swap
	Params *SwapParams `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) Reset() {
	*x = AaveDebtTokenAdaptor_SwapAndRepay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_aave_debt_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveDebtTokenAdaptor_SwapAndRepay) ProtoMessage() {}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_aave_debt_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveDebtTokenAdaptor_SwapAndRepay.ProtoReflect.Descriptor instead.
func (*AaveDebtTokenAdaptor_SwapAndRepay) Descriptor() ([]byte, []int) {
	return file_adaptors_aave_debt_token_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) GetTokenIn() string {
	if x != nil {
		return x.TokenIn
	}
	return ""
}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) GetTokenToRepay() string {
	if x != nil {
		return x.TokenToRepay
	}
	return ""
}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) GetExchange() Exchange {
	if x != nil {
		return x.Exchange
	}
	return Exchange_EXCHANGE_UNSPECIFIED
}

func (x *AaveDebtTokenAdaptor_SwapAndRepay) GetParams() *SwapParams {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_adaptors_aave_debt_token_proto protoreflect.FileDescriptor

var file_adaptors_aave_debt_token_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x61, 0x76, 0x65, 0x2f,
	0x64, 0x65, 0x62, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x1a, 0x13, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaa, 0x06, 0x0a, 0x14, 0x41, 0x61, 0x76, 0x65, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x77, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x48, 0x00, 0x52, 0x04, 0x73, 0x77, 0x61, 0x70,
	0x12, 0x39, 0x0a, 0x0b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x77, 0x61, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x32, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x48, 0x00, 0x52,
	0x0a, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x12, 0x45, 0x0a, 0x0f, 0x72,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x12, 0x5b, 0x0a, 0x10, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x44, 0x65,
	0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x42,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61, 0x76, 0x65, 0x48, 0x00, 0x52,
	0x0e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61, 0x76, 0x65, 0x12,
	0x58, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x64, 0x65,
	0x62, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x61, 0x79,
	0x41, 0x61, 0x76, 0x65, 0x44, 0x65, 0x62, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x41, 0x61, 0x76, 0x65, 0x44, 0x65, 0x62, 0x74, 0x12, 0x55, 0x0a, 0x0e, 0x73, 0x77, 0x61,
	0x70, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x61, 0x76, 0x65, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x61, 0x79,
	0x48, 0x00, 0x52, 0x0c, 0x73, 0x77, 0x61, 0x70, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x61, 0x79,
	0x1a, 0x3e, 0x0a, 0x0e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x61,
	0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x3d, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x61, 0x79, 0x41, 0x61, 0x76, 0x65, 0x44, 0x65, 0x62,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0xce, 0x01, 0x0a, 0x0c, 0x53, 0x77, 0x61, 0x70, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x61, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x70, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x61,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x30,
	0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x77,
	0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x19,
	0x41, 0x61, 0x76, 0x65, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x6f, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x63, 0x61, 0x6c,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x44, 0x65, 0x62, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c,
	0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adaptors_aave_debt_token_proto_rawDescOnce sync.Once
	file_adaptors_aave_debt_token_proto_rawDescData = file_adaptors_aave_debt_token_proto_rawDesc
)

func file_adaptors_aave_debt_token_proto_rawDescGZIP() []byte {
	file_adaptors_aave_debt_token_proto_rawDescOnce.Do(func() {
		file_adaptors_aave_debt_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_adaptors_aave_debt_token_proto_rawDescData)
	})
	return file_adaptors_aave_debt_token_proto_rawDescData
}

var file_adaptors_aave_debt_token_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_adaptors_aave_debt_token_proto_goTypes = []interface{}{
	(*AaveDebtTokenAdaptor)(nil),                // 0: steward.v2.AaveDebtTokenAdaptor
	(*AaveDebtTokenAdaptorCalls)(nil),           // 1: steward.v2.AaveDebtTokenAdaptorCalls
	(*AaveDebtTokenAdaptor_BorrowFromAave)(nil), // 2: steward.v2.AaveDebtTokenAdaptor.BorrowFromAave
	(*AaveDebtTokenAdaptor_RepayAaveDebt)(nil),  // 3: steward.v2.AaveDebtTokenAdaptor.RepayAaveDebt
	(*AaveDebtTokenAdaptor_SwapAndRepay)(nil),   // 4: steward.v2.AaveDebtTokenAdaptor.SwapAndRepay
	(*Swap)(nil),           // 5: steward.v2.Swap
	(*OracleSwap)(nil),     // 6: steward.v2.OracleSwap
	(*RevokeApproval)(nil), // 7: steward.v2.RevokeApproval
	(Exchange)(0),          // 8: steward.v2.Exchange
	(*SwapParams)(nil),     // 9: steward.v2.SwapParams
}
var file_adaptors_aave_debt_token_proto_depIdxs = []int32{
	5, // 0: steward.v2.AaveDebtTokenAdaptor.swap:type_name -> steward.v2.Swap
	6, // 1: steward.v2.AaveDebtTokenAdaptor.oracle_swap:type_name -> steward.v2.OracleSwap
	7, // 2: steward.v2.AaveDebtTokenAdaptor.revoke_approval:type_name -> steward.v2.RevokeApproval
	2, // 3: steward.v2.AaveDebtTokenAdaptor.borrow_from_aave:type_name -> steward.v2.AaveDebtTokenAdaptor.BorrowFromAave
	3, // 4: steward.v2.AaveDebtTokenAdaptor.repay_aave_debt:type_name -> steward.v2.AaveDebtTokenAdaptor.RepayAaveDebt
	4, // 5: steward.v2.AaveDebtTokenAdaptor.swap_and_repay:type_name -> steward.v2.AaveDebtTokenAdaptor.SwapAndRepay
	0, // 6: steward.v2.AaveDebtTokenAdaptorCalls.calls:type_name -> steward.v2.AaveDebtTokenAdaptor
	8, // 7: steward.v2.AaveDebtTokenAdaptor.SwapAndRepay.exchange:type_name -> steward.v2.Exchange
	9, // 8: steward.v2.AaveDebtTokenAdaptor.SwapAndRepay.params:type_name -> steward.v2.SwapParams
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_adaptors_aave_debt_token_proto_init() }
func file_adaptors_aave_debt_token_proto_init() {
	if File_adaptors_aave_debt_token_proto != nil {
		return
	}
	file_adaptors_base_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_adaptors_aave_debt_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveDebtTokenAdaptor); i {
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
		file_adaptors_aave_debt_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveDebtTokenAdaptorCalls); i {
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
		file_adaptors_aave_debt_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveDebtTokenAdaptor_BorrowFromAave); i {
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
		file_adaptors_aave_debt_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveDebtTokenAdaptor_RepayAaveDebt); i {
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
		file_adaptors_aave_debt_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveDebtTokenAdaptor_SwapAndRepay); i {
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
	file_adaptors_aave_debt_token_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AaveDebtTokenAdaptor_Swap)(nil),
		(*AaveDebtTokenAdaptor_OracleSwap)(nil),
		(*AaveDebtTokenAdaptor_RevokeApproval)(nil),
		(*AaveDebtTokenAdaptor_BorrowFromAave_)(nil),
		(*AaveDebtTokenAdaptor_RepayAaveDebt_)(nil),
		(*AaveDebtTokenAdaptor_SwapAndRepay_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adaptors_aave_debt_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adaptors_aave_debt_token_proto_goTypes,
		DependencyIndexes: file_adaptors_aave_debt_token_proto_depIdxs,
		MessageInfos:      file_adaptors_aave_debt_token_proto_msgTypes,
	}.Build()
	File_adaptors_aave_debt_token_proto = out.File
	file_adaptors_aave_debt_token_proto_rawDesc = nil
	file_adaptors_aave_debt_token_proto_goTypes = nil
	file_adaptors_aave_debt_token_proto_depIdxs = nil
}
