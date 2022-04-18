// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: aave_v2_stablecoin.proto

package integration_tests

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

type AaveV2Stablecoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*AaveV2Stablecoin_EnterStrategy
	//	*AaveV2Stablecoin_ReinvestAmount
	//	*AaveV2Stablecoin_Reinvest
	//	*AaveV2Stablecoin_ClaimAndUnstakeAmount
	//	*AaveV2Stablecoin_ClaimAndUnstake
	//	*AaveV2Stablecoin_Rebalance
	//	*AaveV2Stablecoin_AccruePlatformFees
	//	*AaveV2Stablecoin_TransferFees
	//	*AaveV2Stablecoin_SetInputToken
	//	*AaveV2Stablecoin_RemoveLiquidityRestriction
	//	*AaveV2Stablecoin_Sweep
	Function isAaveV2Stablecoin_Function `protobuf_oneof:"function"`
}

func (x *AaveV2Stablecoin) Reset() {
	*x = AaveV2Stablecoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin) ProtoMessage() {}

func (x *AaveV2Stablecoin) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AaveV2Stablecoin.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0}
}

func (m *AaveV2Stablecoin) GetFunction() isAaveV2Stablecoin_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *AaveV2Stablecoin) GetEnterStrategy() *EnterStrategy {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_EnterStrategy); ok {
		return x.EnterStrategy
	}
	return nil
}

func (x *AaveV2Stablecoin) GetReinvestAmount() *ReinvestWithAmount {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_ReinvestAmount); ok {
		return x.ReinvestAmount
	}
	return nil
}

func (x *AaveV2Stablecoin) GetReinvest() *Reinvest {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_Reinvest); ok {
		return x.Reinvest
	}
	return nil
}

func (x *AaveV2Stablecoin) GetClaimAndUnstakeAmount() *ClaimAndUnstakeWithAmount {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_ClaimAndUnstakeAmount); ok {
		return x.ClaimAndUnstakeAmount
	}
	return nil
}

func (x *AaveV2Stablecoin) GetClaimAndUnstake() *ClaimAndUnstake {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_ClaimAndUnstake); ok {
		return x.ClaimAndUnstake
	}
	return nil
}

func (x *AaveV2Stablecoin) GetRebalance() *Rebalance {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_Rebalance); ok {
		return x.Rebalance
	}
	return nil
}

func (x *AaveV2Stablecoin) GetAccruePlatformFees() *AccruePlatformFees {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_AccruePlatformFees); ok {
		return x.AccruePlatformFees
	}
	return nil
}

func (x *AaveV2Stablecoin) GetTransferFees() *TransferFees {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_TransferFees); ok {
		return x.TransferFees
	}
	return nil
}

func (x *AaveV2Stablecoin) GetSetInputToken() *SetInputToken {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_SetInputToken); ok {
		return x.SetInputToken
	}
	return nil
}

func (x *AaveV2Stablecoin) GetRemoveLiquidityRestriction() *RemoveLiquidityRestriction {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_RemoveLiquidityRestriction); ok {
		return x.RemoveLiquidityRestriction
	}
	return nil
}

func (x *AaveV2Stablecoin) GetSweep() *Sweep {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_Sweep); ok {
		return x.Sweep
	}
	return nil
}

type isAaveV2Stablecoin_Function interface {
	isAaveV2Stablecoin_Function()
}

type AaveV2Stablecoin_EnterStrategy struct {
	EnterStrategy *EnterStrategy `protobuf:"bytes,1,opt,name=enter_strategy,json=enterStrategy,proto3,oneof"`
}

type AaveV2Stablecoin_ReinvestAmount struct {
	ReinvestAmount *ReinvestWithAmount `protobuf:"bytes,2,opt,name=reinvest_amount,json=reinvestAmount,proto3,oneof"`
}

type AaveV2Stablecoin_Reinvest struct {
	Reinvest *Reinvest `protobuf:"bytes,3,opt,name=reinvest,proto3,oneof"`
}

type AaveV2Stablecoin_ClaimAndUnstakeAmount struct {
	ClaimAndUnstakeAmount *ClaimAndUnstakeWithAmount `protobuf:"bytes,4,opt,name=claim_and_unstake_amount,json=claimAndUnstakeAmount,proto3,oneof"`
}

type AaveV2Stablecoin_ClaimAndUnstake struct {
	ClaimAndUnstake *ClaimAndUnstake `protobuf:"bytes,5,opt,name=claim_and_unstake,json=claimAndUnstake,proto3,oneof"`
}

type AaveV2Stablecoin_Rebalance struct {
	Rebalance *Rebalance `protobuf:"bytes,6,opt,name=rebalance,proto3,oneof"`
}

type AaveV2Stablecoin_AccruePlatformFees struct {
	AccruePlatformFees *AccruePlatformFees `protobuf:"bytes,7,opt,name=accrue_platform_fees,json=accruePlatformFees,proto3,oneof"`
}

type AaveV2Stablecoin_TransferFees struct {
	TransferFees *TransferFees `protobuf:"bytes,8,opt,name=transfer_fees,json=transferFees,proto3,oneof"`
}

type AaveV2Stablecoin_SetInputToken struct {
	SetInputToken *SetInputToken `protobuf:"bytes,9,opt,name=set_input_token,json=setInputToken,proto3,oneof"`
}

type AaveV2Stablecoin_RemoveLiquidityRestriction struct {
	RemoveLiquidityRestriction *RemoveLiquidityRestriction `protobuf:"bytes,10,opt,name=remove_liquidity_restriction,json=removeLiquidityRestriction,proto3,oneof"`
}

type AaveV2Stablecoin_Sweep struct {
	Sweep *Sweep `protobuf:"bytes,11,opt,name=sweep,proto3,oneof"`
}

func (*AaveV2Stablecoin_EnterStrategy) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_ReinvestAmount) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_Reinvest) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_ClaimAndUnstakeAmount) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_ClaimAndUnstake) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_Rebalance) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_AccruePlatformFees) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_TransferFees) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_SetInputToken) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_RemoveLiquidityRestriction) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_Sweep) isAaveV2Stablecoin_Function() {}

// function enterStrategy() external;
type EnterStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnterStrategy) Reset() {
	*x = EnterStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterStrategy) ProtoMessage() {}

func (x *EnterStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterStrategy.ProtoReflect.Descriptor instead.
func (*EnterStrategy) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{1}
}

// function reinvest(uint256 amount, uint256 minAssetsOut) external;
type ReinvestWithAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount       uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	MinAssetsOut uint64 `protobuf:"varint,2,opt,name=min_assets_out,json=minAssetsOut,proto3" json:"min_assets_out,omitempty"`
}

func (x *ReinvestWithAmount) Reset() {
	*x = ReinvestWithAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReinvestWithAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReinvestWithAmount) ProtoMessage() {}

func (x *ReinvestWithAmount) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReinvestWithAmount.ProtoReflect.Descriptor instead.
func (*ReinvestWithAmount) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{2}
}

func (x *ReinvestWithAmount) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReinvestWithAmount) GetMinAssetsOut() uint64 {
	if x != nil {
		return x.MinAssetsOut
	}
	return 0
}

// function reinvest(uint256 minAssetsOut) external;
type Reinvest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinAssetsOut uint64 `protobuf:"varint,1,opt,name=min_assets_out,json=minAssetsOut,proto3" json:"min_assets_out,omitempty"`
}

func (x *Reinvest) Reset() {
	*x = Reinvest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reinvest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reinvest) ProtoMessage() {}

func (x *Reinvest) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reinvest.ProtoReflect.Descriptor instead.
func (*Reinvest) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{3}
}

func (x *Reinvest) GetMinAssetsOut() uint64 {
	if x != nil {
		return x.MinAssetsOut
	}
	return 0
}

// function claimAndUnstake(uint256 amount) external returns (uint256 claimed);
type ClaimAndUnstakeWithAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ClaimAndUnstakeWithAmount) Reset() {
	*x = ClaimAndUnstakeWithAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimAndUnstakeWithAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimAndUnstakeWithAmount) ProtoMessage() {}

func (x *ClaimAndUnstakeWithAmount) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimAndUnstakeWithAmount.ProtoReflect.Descriptor instead.
func (*ClaimAndUnstakeWithAmount) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{4}
}

func (x *ClaimAndUnstakeWithAmount) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// function claimAndUnstake() external returns (uint256);
type ClaimAndUnstake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClaimAndUnstake) Reset() {
	*x = ClaimAndUnstake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimAndUnstake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimAndUnstake) ProtoMessage() {}

func (x *ClaimAndUnstake) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimAndUnstake.ProtoReflect.Descriptor instead.
func (*ClaimAndUnstake) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{5}
}

// function rebalance(address newLendingToken, uint256 minNewLendingTokenAmount) external;
type Rebalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address                  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MinNewLendingTokenAmount uint64 `protobuf:"varint,2,opt,name=min_new_lending_token_amount,json=minNewLendingTokenAmount,proto3" json:"min_new_lending_token_amount,omitempty"`
}

func (x *Rebalance) Reset() {
	*x = Rebalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rebalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rebalance) ProtoMessage() {}

func (x *Rebalance) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rebalance.ProtoReflect.Descriptor instead.
func (*Rebalance) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{6}
}

func (x *Rebalance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Rebalance) GetMinNewLendingTokenAmount() uint64 {
	if x != nil {
		return x.MinNewLendingTokenAmount
	}
	return 0
}

// function accruePlatformFees() external;
type AccruePlatformFees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccruePlatformFees) Reset() {
	*x = AccruePlatformFees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccruePlatformFees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccruePlatformFees) ProtoMessage() {}

func (x *AccruePlatformFees) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccruePlatformFees.ProtoReflect.Descriptor instead.
func (*AccruePlatformFees) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{7}
}

// function transferFees() external;
type TransferFees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransferFees) Reset() {
	*x = TransferFees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFees) ProtoMessage() {}

func (x *TransferFees) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFees.ProtoReflect.Descriptor instead.
func (*TransferFees) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{8}
}

// function setInputToken(address token, bool isApproved) external;
type SetInputToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	IsApproved bool   `protobuf:"varint,2,opt,name=is_approved,json=isApproved,proto3" json:"is_approved,omitempty"`
}

func (x *SetInputToken) Reset() {
	*x = SetInputToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetInputToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetInputToken) ProtoMessage() {}

func (x *SetInputToken) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetInputToken.ProtoReflect.Descriptor instead.
func (*SetInputToken) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{9}
}

func (x *SetInputToken) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SetInputToken) GetIsApproved() bool {
	if x != nil {
		return x.IsApproved
	}
	return false
}

// function removeLiquidityRestriction() external;
type RemoveLiquidityRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveLiquidityRestriction) Reset() {
	*x = RemoveLiquidityRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveLiquidityRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveLiquidityRestriction) ProtoMessage() {}

func (x *RemoveLiquidityRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveLiquidityRestriction.ProtoReflect.Descriptor instead.
func (*RemoveLiquidityRestriction) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{10}
}

// function sweep(address token) external;
type Sweep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Sweep) Reset() {
	*x = Sweep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sweep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sweep) ProtoMessage() {}

func (x *Sweep) ProtoReflect() protoreflect.Message {
	mi := &file_aave_v2_stablecoin_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sweep.ProtoReflect.Descriptor instead.
func (*Sweep) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{11}
}

func (x *Sweep) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_aave_v2_stablecoin_proto protoreflect.FileDescriptor

var file_aave_v2_stablecoin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32, 0x5f, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xb6, 0x06, 0x0a, 0x10, 0x41, 0x61, 0x76, 0x65, 0x56,
	0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x42, 0x0a, 0x0e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00,
	0x52, 0x0d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x49, 0x0a, 0x0f, 0x72, 0x65, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x69, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65,
	0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x69, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x12, 0x60,
	0x0a, 0x18, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x41, 0x6e, 0x64, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x15, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x41, 0x6e, 0x64, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x49, 0x0a, 0x11, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x75, 0x6e,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x41, 0x6e,
	0x64, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x41, 0x6e, 0x64, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x72,
	0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x61, 0x63, 0x63, 0x72, 0x75, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x72, 0x75, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x65, 0x65, 0x73,
	0x48, 0x00, 0x52, 0x12, 0x61, 0x63, 0x63, 0x72, 0x75, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x46, 0x65, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x65, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x65, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x73,
	0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x6a, 0x0a, 0x1c,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1a, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x77, 0x65, 0x65,
	0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x65, 0x65, 0x70, 0x48, 0x00, 0x52, 0x05, 0x73, 0x77,
	0x65, 0x65, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x0f, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x22, 0x52, 0x0a, 0x12, 0x52, 0x65, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x6f, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x4f, 0x75, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x6f,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x33, 0x0a, 0x19, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x41,
	0x6e, 0x64, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x57, 0x69, 0x74, 0x68, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x41, 0x6e, 0x64, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x22, 0x65,
	0x0a, 0x09, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x1c, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x65, 0x77,
	0x5f, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x18, 0x6d, 0x69, 0x6e,
	0x4e, 0x65, 0x77, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x72, 0x75, 0x65, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x65, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x65, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x0d, 0x53,
	0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x05, 0x53, 0x77, 0x65, 0x65, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_aave_v2_stablecoin_proto_rawDescOnce sync.Once
	file_aave_v2_stablecoin_proto_rawDescData = file_aave_v2_stablecoin_proto_rawDesc
)

func file_aave_v2_stablecoin_proto_rawDescGZIP() []byte {
	file_aave_v2_stablecoin_proto_rawDescOnce.Do(func() {
		file_aave_v2_stablecoin_proto_rawDescData = protoimpl.X.CompressGZIP(file_aave_v2_stablecoin_proto_rawDescData)
	})
	return file_aave_v2_stablecoin_proto_rawDescData
}

var file_aave_v2_stablecoin_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_aave_v2_stablecoin_proto_goTypes = []interface{}{
	(*AaveV2Stablecoin)(nil),           // 0: steward.v1.AaveV2Stablecoin
	(*EnterStrategy)(nil),              // 1: steward.v1.EnterStrategy
	(*ReinvestWithAmount)(nil),         // 2: steward.v1.ReinvestWithAmount
	(*Reinvest)(nil),                   // 3: steward.v1.Reinvest
	(*ClaimAndUnstakeWithAmount)(nil),  // 4: steward.v1.ClaimAndUnstakeWithAmount
	(*ClaimAndUnstake)(nil),            // 5: steward.v1.ClaimAndUnstake
	(*Rebalance)(nil),                  // 6: steward.v1.Rebalance
	(*AccruePlatformFees)(nil),         // 7: steward.v1.AccruePlatformFees
	(*TransferFees)(nil),               // 8: steward.v1.TransferFees
	(*SetInputToken)(nil),              // 9: steward.v1.SetInputToken
	(*RemoveLiquidityRestriction)(nil), // 10: steward.v1.RemoveLiquidityRestriction
	(*Sweep)(nil),                      // 11: steward.v1.Sweep
}
var file_aave_v2_stablecoin_proto_depIdxs = []int32{
	1,  // 0: steward.v1.AaveV2Stablecoin.enter_strategy:type_name -> steward.v1.EnterStrategy
	2,  // 1: steward.v1.AaveV2Stablecoin.reinvest_amount:type_name -> steward.v1.ReinvestWithAmount
	3,  // 2: steward.v1.AaveV2Stablecoin.reinvest:type_name -> steward.v1.Reinvest
	4,  // 3: steward.v1.AaveV2Stablecoin.claim_and_unstake_amount:type_name -> steward.v1.ClaimAndUnstakeWithAmount
	5,  // 4: steward.v1.AaveV2Stablecoin.claim_and_unstake:type_name -> steward.v1.ClaimAndUnstake
	6,  // 5: steward.v1.AaveV2Stablecoin.rebalance:type_name -> steward.v1.Rebalance
	7,  // 6: steward.v1.AaveV2Stablecoin.accrue_platform_fees:type_name -> steward.v1.AccruePlatformFees
	8,  // 7: steward.v1.AaveV2Stablecoin.transfer_fees:type_name -> steward.v1.TransferFees
	9,  // 8: steward.v1.AaveV2Stablecoin.set_input_token:type_name -> steward.v1.SetInputToken
	10, // 9: steward.v1.AaveV2Stablecoin.remove_liquidity_restriction:type_name -> steward.v1.RemoveLiquidityRestriction
	11, // 10: steward.v1.AaveV2Stablecoin.sweep:type_name -> steward.v1.Sweep
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_aave_v2_stablecoin_proto_init() }
func file_aave_v2_stablecoin_proto_init() {
	if File_aave_v2_stablecoin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aave_v2_stablecoin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AaveV2Stablecoin); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterStrategy); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReinvestWithAmount); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reinvest); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimAndUnstakeWithAmount); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimAndUnstake); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rebalance); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccruePlatformFees); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFees); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetInputToken); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveLiquidityRestriction); i {
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
		file_aave_v2_stablecoin_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sweep); i {
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
	file_aave_v2_stablecoin_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AaveV2Stablecoin_EnterStrategy)(nil),
		(*AaveV2Stablecoin_ReinvestAmount)(nil),
		(*AaveV2Stablecoin_Reinvest)(nil),
		(*AaveV2Stablecoin_ClaimAndUnstakeAmount)(nil),
		(*AaveV2Stablecoin_ClaimAndUnstake)(nil),
		(*AaveV2Stablecoin_Rebalance)(nil),
		(*AaveV2Stablecoin_AccruePlatformFees)(nil),
		(*AaveV2Stablecoin_TransferFees)(nil),
		(*AaveV2Stablecoin_SetInputToken)(nil),
		(*AaveV2Stablecoin_RemoveLiquidityRestriction)(nil),
		(*AaveV2Stablecoin_Sweep)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aave_v2_stablecoin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aave_v2_stablecoin_proto_goTypes,
		DependencyIndexes: file_aave_v2_stablecoin_proto_depIdxs,
		MessageInfos:      file_aave_v2_stablecoin_proto_msgTypes,
	}.Build()
	File_aave_v2_stablecoin_proto = out.File
	file_aave_v2_stablecoin_proto_rawDesc = nil
	file_aave_v2_stablecoin_proto_goTypes = nil
	file_aave_v2_stablecoin_proto_depIdxs = nil
}