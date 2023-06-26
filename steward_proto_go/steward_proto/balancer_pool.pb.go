//
// Protos for function calls to the Compound CToken adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: adaptors/balancer/balancer_pool.proto

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

// Represents call data for the Compound C Token adaptor V2, managing lending positions on Compound.
type BalancerPoolAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*BalancerPoolAdaptorV1_RevokeApproval
	//	*BalancerPoolAdaptorV1_RelayerJoinPool_
	//	*BalancerPoolAdaptorV1_RelayerExitPool_
	//	*BalancerPoolAdaptorV1_StakeBpt
	//	*BalancerPoolAdaptorV1_UnstakeBpt
	//	*BalancerPoolAdaptorV1_ClaimRewards_
	Function isBalancerPoolAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *BalancerPoolAdaptorV1) Reset() {
	*x = BalancerPoolAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerPoolAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerPoolAdaptorV1) ProtoMessage() {}

func (x *BalancerPoolAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerPoolAdaptorV1.ProtoReflect.Descriptor instead.
func (*BalancerPoolAdaptorV1) Descriptor() ([]byte, []int) {
	return file_adaptors_balancer_balancer_pool_proto_rawDescGZIP(), []int{0}
}

func (m *BalancerPoolAdaptorV1) GetFunction() isBalancerPoolAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *BalancerPoolAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*BalancerPoolAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *BalancerPoolAdaptorV1) GetRelayerJoinPool() *BalancerPoolAdaptorV1_RelayerJoinPool {
	if x, ok := x.GetFunction().(*BalancerPoolAdaptorV1_RelayerJoinPool_); ok {
		return x.RelayerJoinPool
	}
	return nil
}

func (x *BalancerPoolAdaptorV1) GetRelayerExitPool() *BalancerPoolAdaptorV1_RelayerExitPool {
	if x, ok := x.GetFunction().(*BalancerPoolAdaptorV1_RelayerExitPool_); ok {
		return x.RelayerExitPool
	}
	return nil
}

func (x *BalancerPoolAdaptorV1) GetStakeBpt() *BalancerPoolAdaptorV1_StakeBPT {
	if x, ok := x.GetFunction().(*BalancerPoolAdaptorV1_StakeBpt); ok {
		return x.StakeBpt
	}
	return nil
}

func (x *BalancerPoolAdaptorV1) GetUnstakeBpt() *BalancerPoolAdaptorV1_UnstakeBPT {
	if x, ok := x.GetFunction().(*BalancerPoolAdaptorV1_UnstakeBpt); ok {
		return x.UnstakeBpt
	}
	return nil
}

func (x *BalancerPoolAdaptorV1) GetClaimRewards() *BalancerPoolAdaptorV1_ClaimRewards {
	if x, ok := x.GetFunction().(*BalancerPoolAdaptorV1_ClaimRewards_); ok {
		return x.ClaimRewards
	}
	return nil
}

type isBalancerPoolAdaptorV1_Function interface {
	isBalancerPoolAdaptorV1_Function()
}

type BalancerPoolAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type BalancerPoolAdaptorV1_RelayerJoinPool_ struct {
	// Represents function `relayerJoinPool(ERC20[] tokensIn, uint256[] amountsIn, ERC20 btpOut, bytes[] memory callData)`
	RelayerJoinPool *BalancerPoolAdaptorV1_RelayerJoinPool `protobuf:"bytes,2,opt,name=relayer_join_pool,json=relayerJoinPool,proto3,oneof"`
}

type BalancerPoolAdaptorV1_RelayerExitPool_ struct {
	// Represents function `relayerExitPool(ERC20 bptIn, uint256 amountIn, ERC20[] memory tokensOut, bytes[] memory callData)`
	RelayerExitPool *BalancerPoolAdaptorV1_RelayerExitPool `protobuf:"bytes,3,opt,name=relayer_exit_pool,json=relayerExitPool,proto3,oneof"`
}

type BalancerPoolAdaptorV1_StakeBpt struct {
	// Represents function `stakeBPT(ERC20 _bpt, address _liquidityGauge, uint256 _amountIn)`
	StakeBpt *BalancerPoolAdaptorV1_StakeBPT `protobuf:"bytes,4,opt,name=stake_bpt,json=stakeBpt,proto3,oneof"`
}

type BalancerPoolAdaptorV1_UnstakeBpt struct {
	// Represents function `unstakeBPT(ERC20 _bpt, address _liquidityGauge, uint256 _amountOut)`
	UnstakeBpt *BalancerPoolAdaptorV1_UnstakeBPT `protobuf:"bytes,5,opt,name=unstake_bpt,json=unstakeBpt,proto3,oneof"`
}

type BalancerPoolAdaptorV1_ClaimRewards_ struct {
	// Represents function `claimRewards(address gauge)`
	ClaimRewards *BalancerPoolAdaptorV1_ClaimRewards `protobuf:"bytes,6,opt,name=claim_rewards,json=claimRewards,proto3,oneof"`
}

func (*BalancerPoolAdaptorV1_RevokeApproval) isBalancerPoolAdaptorV1_Function() {}

func (*BalancerPoolAdaptorV1_RelayerJoinPool_) isBalancerPoolAdaptorV1_Function() {}

func (*BalancerPoolAdaptorV1_RelayerExitPool_) isBalancerPoolAdaptorV1_Function() {}

func (*BalancerPoolAdaptorV1_StakeBpt) isBalancerPoolAdaptorV1_Function() {}

func (*BalancerPoolAdaptorV1_UnstakeBpt) isBalancerPoolAdaptorV1_Function() {}

func (*BalancerPoolAdaptorV1_ClaimRewards_) isBalancerPoolAdaptorV1_Function() {}

type BalancerPoolAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*BalancerPoolAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *BalancerPoolAdaptorV1Calls) Reset() {
	*x = BalancerPoolAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerPoolAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerPoolAdaptorV1Calls) ProtoMessage() {}

func (x *BalancerPoolAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerPoolAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*BalancerPoolAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_adaptors_balancer_balancer_pool_proto_rawDescGZIP(), []int{1}
}

func (x *BalancerPoolAdaptorV1Calls) GetCalls() []*BalancerPoolAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Call `BalancerRelayer` on mainnet to carry out join txs
//
// Represents function `relayerJoinPool(ERC20[] tokensIn, uint256[] amountsIn, ERC20 btpOut, bytes[] memory callData)`
type BalancerPoolAdaptorV1_RelayerJoinPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tokens to join.
	TokensIn []string `protobuf:"bytes,1,rep,name=tokens_in,json=tokensIn,proto3" json:"tokens_in,omitempty"`
	// The amounts to join
	AmountsIn []string `protobuf:"bytes,2,rep,name=amounts_in,json=amountsIn,proto3" json:"amounts_in,omitempty"`
	// The token to receive
	BtpOut string `protobuf:"bytes,3,opt,name=btp_out,json=btpOut,proto3" json:"btp_out,omitempty"`
	// The call data for the relayer
	CallData [][]byte `protobuf:"bytes,4,rep,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
}

func (x *BalancerPoolAdaptorV1_RelayerJoinPool) Reset() {
	*x = BalancerPoolAdaptorV1_RelayerJoinPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerPoolAdaptorV1_RelayerJoinPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerPoolAdaptorV1_RelayerJoinPool) ProtoMessage() {}

func (x *BalancerPoolAdaptorV1_RelayerJoinPool) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerPoolAdaptorV1_RelayerJoinPool.ProtoReflect.Descriptor instead.
func (*BalancerPoolAdaptorV1_RelayerJoinPool) Descriptor() ([]byte, []int) {
	return file_adaptors_balancer_balancer_pool_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BalancerPoolAdaptorV1_RelayerJoinPool) GetTokensIn() []string {
	if x != nil {
		return x.TokensIn
	}
	return nil
}

func (x *BalancerPoolAdaptorV1_RelayerJoinPool) GetAmountsIn() []string {
	if x != nil {
		return x.AmountsIn
	}
	return nil
}

func (x *BalancerPoolAdaptorV1_RelayerJoinPool) GetBtpOut() string {
	if x != nil {
		return x.BtpOut
	}
	return ""
}

func (x *BalancerPoolAdaptorV1_RelayerJoinPool) GetCallData() [][]byte {
	if x != nil {
		return x.CallData
	}
	return nil
}

//
// Call `BalancerRelayer` on mainnet to carry out exit txs
//
// Represents function `relayerExitPool(ERC20 bptIn, uint256 amountIn, ERC20[] memory tokensOut, bytes[] memory callData)``
type BalancerPoolAdaptorV1_RelayerExitPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token to exit
	BptIn string `protobuf:"bytes,1,opt,name=bpt_in,json=bptIn,proto3" json:"bpt_in,omitempty"`
	// The amount to exit
	AmountIn string `protobuf:"bytes,2,opt,name=amount_in,json=amountIn,proto3" json:"amount_in,omitempty"`
	// The tokens to receive
	TokensOut []string `protobuf:"bytes,3,rep,name=tokens_out,json=tokensOut,proto3" json:"tokens_out,omitempty"`
	// The call data for the relayer
	CallData [][]byte `protobuf:"bytes,4,rep,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
}

func (x *BalancerPoolAdaptorV1_RelayerExitPool) Reset() {
	*x = BalancerPoolAdaptorV1_RelayerExitPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerPoolAdaptorV1_RelayerExitPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerPoolAdaptorV1_RelayerExitPool) ProtoMessage() {}

func (x *BalancerPoolAdaptorV1_RelayerExitPool) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerPoolAdaptorV1_RelayerExitPool.ProtoReflect.Descriptor instead.
func (*BalancerPoolAdaptorV1_RelayerExitPool) Descriptor() ([]byte, []int) {
	return file_adaptors_balancer_balancer_pool_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BalancerPoolAdaptorV1_RelayerExitPool) GetBptIn() string {
	if x != nil {
		return x.BptIn
	}
	return ""
}

func (x *BalancerPoolAdaptorV1_RelayerExitPool) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

func (x *BalancerPoolAdaptorV1_RelayerExitPool) GetTokensOut() []string {
	if x != nil {
		return x.TokensOut
	}
	return nil
}

func (x *BalancerPoolAdaptorV1_RelayerExitPool) GetCallData() [][]byte {
	if x != nil {
		return x.CallData
	}
	return nil
}

//
// Stake (deposit) BPTs into respective pool gauge
//
// Represents `function stakeBPT(ERC20 _bpt, address _liquidityGauge, uint256 _amountIn)``
type BalancerPoolAdaptorV1_StakeBPT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The BPT to stake
	Bpt string `protobuf:"bytes,1,opt,name=bpt,proto3" json:"bpt,omitempty"`
	// The liquidity gauge to stake into
	LiquidityGauge string `protobuf:"bytes,2,opt,name=liquidity_gauge,json=liquidityGauge,proto3" json:"liquidity_gauge,omitempty"`
	// The amount to stake
	AmountIn string `protobuf:"bytes,3,opt,name=amount_in,json=amountIn,proto3" json:"amount_in,omitempty"`
}

func (x *BalancerPoolAdaptorV1_StakeBPT) Reset() {
	*x = BalancerPoolAdaptorV1_StakeBPT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerPoolAdaptorV1_StakeBPT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerPoolAdaptorV1_StakeBPT) ProtoMessage() {}

func (x *BalancerPoolAdaptorV1_StakeBPT) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerPoolAdaptorV1_StakeBPT.ProtoReflect.Descriptor instead.
func (*BalancerPoolAdaptorV1_StakeBPT) Descriptor() ([]byte, []int) {
	return file_adaptors_balancer_balancer_pool_proto_rawDescGZIP(), []int{0, 2}
}

func (x *BalancerPoolAdaptorV1_StakeBPT) GetBpt() string {
	if x != nil {
		return x.Bpt
	}
	return ""
}

func (x *BalancerPoolAdaptorV1_StakeBPT) GetLiquidityGauge() string {
	if x != nil {
		return x.LiquidityGauge
	}
	return ""
}

func (x *BalancerPoolAdaptorV1_StakeBPT) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

//
// Unstake (withdraw) BPT from respective pool gauge
//
// Represents `function unstakeBPT(ERC20 _bpt, address _liquidityGauge, uint256 _amountOut)``
type BalancerPoolAdaptorV1_UnstakeBPT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The BPT to unstake
	Bpt string `protobuf:"bytes,1,opt,name=bpt,proto3" json:"bpt,omitempty"`
	// The liquidity gauge to unstake from
	LiquidityGauge string `protobuf:"bytes,2,opt,name=liquidity_gauge,json=liquidityGauge,proto3" json:"liquidity_gauge,omitempty"`
	// The amount to unstake
	AmountOut string `protobuf:"bytes,3,opt,name=amount_out,json=amountOut,proto3" json:"amount_out,omitempty"`
}

func (x *BalancerPoolAdaptorV1_UnstakeBPT) Reset() {
	*x = BalancerPoolAdaptorV1_UnstakeBPT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerPoolAdaptorV1_UnstakeBPT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerPoolAdaptorV1_UnstakeBPT) ProtoMessage() {}

func (x *BalancerPoolAdaptorV1_UnstakeBPT) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerPoolAdaptorV1_UnstakeBPT.ProtoReflect.Descriptor instead.
func (*BalancerPoolAdaptorV1_UnstakeBPT) Descriptor() ([]byte, []int) {
	return file_adaptors_balancer_balancer_pool_proto_rawDescGZIP(), []int{0, 3}
}

func (x *BalancerPoolAdaptorV1_UnstakeBPT) GetBpt() string {
	if x != nil {
		return x.Bpt
	}
	return ""
}

func (x *BalancerPoolAdaptorV1_UnstakeBPT) GetLiquidityGauge() string {
	if x != nil {
		return x.LiquidityGauge
	}
	return ""
}

func (x *BalancerPoolAdaptorV1_UnstakeBPT) GetAmountOut() string {
	if x != nil {
		return x.AmountOut
	}
	return ""
}

//
// Claim rewards ($BAL) from LP position
//
// Represents `function claimRewards(address gauge)`
type BalancerPoolAdaptorV1_ClaimRewards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The gauge to claim rewards from
	Guage string `protobuf:"bytes,1,opt,name=guage,proto3" json:"guage,omitempty"`
}

func (x *BalancerPoolAdaptorV1_ClaimRewards) Reset() {
	*x = BalancerPoolAdaptorV1_ClaimRewards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancerPoolAdaptorV1_ClaimRewards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerPoolAdaptorV1_ClaimRewards) ProtoMessage() {}

func (x *BalancerPoolAdaptorV1_ClaimRewards) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_balancer_balancer_pool_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerPoolAdaptorV1_ClaimRewards.ProtoReflect.Descriptor instead.
func (*BalancerPoolAdaptorV1_ClaimRewards) Descriptor() ([]byte, []int) {
	return file_adaptors_balancer_balancer_pool_proto_rawDescGZIP(), []int{0, 4}
}

func (x *BalancerPoolAdaptorV1_ClaimRewards) GetGuage() string {
	if x != nil {
		return x.Guage
	}
	return ""
}

var File_adaptors_balancer_balancer_pool_proto protoreflect.FileDescriptor

var file_adaptors_balancer_balancer_pool_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x33, 0x1a, 0x13, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x08, 0x0a, 0x15, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72,
	0x56, 0x31, 0x12, 0x45, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x5f, 0x0a, 0x11, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x33, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a,
	0x6f, 0x69, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x5f, 0x0a, 0x11, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x33, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x45, 0x78, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x45, 0x78, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x49, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x5f, 0x62, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x42, 0x50, 0x54, 0x48, 0x00, 0x52, 0x08, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x42, 0x70, 0x74, 0x12, 0x4f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6b,
	0x65, 0x5f, 0x62, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x55,
	0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x42, 0x50, 0x54, 0x48, 0x00, 0x52, 0x0a, 0x75, 0x6e, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x42, 0x70, 0x74, 0x12, 0x55, 0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56,
	0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x48, 0x00,
	0x52, 0x0c, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x83,
	0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x49, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x74, 0x70, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x74, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x81, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x45, 0x78, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x70, 0x74, 0x5f,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x70, 0x74, 0x49, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08,
	0x63, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x62, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x6b,
	0x65, 0x42, 0x50, 0x54, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x70, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x5f, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x1a, 0x66, 0x0a, 0x0a,
	0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x42, 0x50, 0x54, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x70, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x4f, 0x75, 0x74, 0x1a, 0x24, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x1a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x43,
	0x61, 0x6c, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x33,
	0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a,
	0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adaptors_balancer_balancer_pool_proto_rawDescOnce sync.Once
	file_adaptors_balancer_balancer_pool_proto_rawDescData = file_adaptors_balancer_balancer_pool_proto_rawDesc
)

func file_adaptors_balancer_balancer_pool_proto_rawDescGZIP() []byte {
	file_adaptors_balancer_balancer_pool_proto_rawDescOnce.Do(func() {
		file_adaptors_balancer_balancer_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_adaptors_balancer_balancer_pool_proto_rawDescData)
	})
	return file_adaptors_balancer_balancer_pool_proto_rawDescData
}

var file_adaptors_balancer_balancer_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_adaptors_balancer_balancer_pool_proto_goTypes = []interface{}{
	(*BalancerPoolAdaptorV1)(nil),                 // 0: steward.v3.BalancerPoolAdaptorV1
	(*BalancerPoolAdaptorV1Calls)(nil),            // 1: steward.v3.BalancerPoolAdaptorV1Calls
	(*BalancerPoolAdaptorV1_RelayerJoinPool)(nil), // 2: steward.v3.BalancerPoolAdaptorV1.RelayerJoinPool
	(*BalancerPoolAdaptorV1_RelayerExitPool)(nil), // 3: steward.v3.BalancerPoolAdaptorV1.RelayerExitPool
	(*BalancerPoolAdaptorV1_StakeBPT)(nil),        // 4: steward.v3.BalancerPoolAdaptorV1.StakeBPT
	(*BalancerPoolAdaptorV1_UnstakeBPT)(nil),      // 5: steward.v3.BalancerPoolAdaptorV1.UnstakeBPT
	(*BalancerPoolAdaptorV1_ClaimRewards)(nil),    // 6: steward.v3.BalancerPoolAdaptorV1.ClaimRewards
	(*RevokeApproval)(nil),                        // 7: steward.v3.RevokeApproval
}
var file_adaptors_balancer_balancer_pool_proto_depIdxs = []int32{
	7, // 0: steward.v3.BalancerPoolAdaptorV1.revoke_approval:type_name -> steward.v3.RevokeApproval
	2, // 1: steward.v3.BalancerPoolAdaptorV1.relayer_join_pool:type_name -> steward.v3.BalancerPoolAdaptorV1.RelayerJoinPool
	3, // 2: steward.v3.BalancerPoolAdaptorV1.relayer_exit_pool:type_name -> steward.v3.BalancerPoolAdaptorV1.RelayerExitPool
	4, // 3: steward.v3.BalancerPoolAdaptorV1.stake_bpt:type_name -> steward.v3.BalancerPoolAdaptorV1.StakeBPT
	5, // 4: steward.v3.BalancerPoolAdaptorV1.unstake_bpt:type_name -> steward.v3.BalancerPoolAdaptorV1.UnstakeBPT
	6, // 5: steward.v3.BalancerPoolAdaptorV1.claim_rewards:type_name -> steward.v3.BalancerPoolAdaptorV1.ClaimRewards
	0, // 6: steward.v3.BalancerPoolAdaptorV1Calls.calls:type_name -> steward.v3.BalancerPoolAdaptorV1
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_adaptors_balancer_balancer_pool_proto_init() }
func file_adaptors_balancer_balancer_pool_proto_init() {
	if File_adaptors_balancer_balancer_pool_proto != nil {
		return
	}
	file_adaptors_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_adaptors_balancer_balancer_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerPoolAdaptorV1); i {
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
		file_adaptors_balancer_balancer_pool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerPoolAdaptorV1Calls); i {
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
		file_adaptors_balancer_balancer_pool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerPoolAdaptorV1_RelayerJoinPool); i {
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
		file_adaptors_balancer_balancer_pool_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerPoolAdaptorV1_RelayerExitPool); i {
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
		file_adaptors_balancer_balancer_pool_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerPoolAdaptorV1_StakeBPT); i {
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
		file_adaptors_balancer_balancer_pool_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerPoolAdaptorV1_UnstakeBPT); i {
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
		file_adaptors_balancer_balancer_pool_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancerPoolAdaptorV1_ClaimRewards); i {
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
	file_adaptors_balancer_balancer_pool_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BalancerPoolAdaptorV1_RevokeApproval)(nil),
		(*BalancerPoolAdaptorV1_RelayerJoinPool_)(nil),
		(*BalancerPoolAdaptorV1_RelayerExitPool_)(nil),
		(*BalancerPoolAdaptorV1_StakeBpt)(nil),
		(*BalancerPoolAdaptorV1_UnstakeBpt)(nil),
		(*BalancerPoolAdaptorV1_ClaimRewards_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adaptors_balancer_balancer_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adaptors_balancer_balancer_pool_proto_goTypes,
		DependencyIndexes: file_adaptors_balancer_balancer_pool_proto_depIdxs,
		MessageInfos:      file_adaptors_balancer_balancer_pool_proto_msgTypes,
	}.Build()
	File_adaptors_balancer_balancer_pool_proto = out.File
	file_adaptors_balancer_balancer_pool_proto_rawDesc = nil
	file_adaptors_balancer_balancer_pool_proto_goTypes = nil
	file_adaptors_balancer_balancer_pool_proto_depIdxs = nil
}
