//
// This is Steward's Strategy Provider API for the AaveV2StablecoinCellar.sol Cellar contract which can be found in this repo:
// https://github.com/PeggyJV/cellar-contracts
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
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

//
// Represents a function call to the Aave V2 Stablecoin cellar
type AaveV2Stablecoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The function you wish to execute on the target cellar
	//
	// Types that are assignable to Function:
	//	*AaveV2Stablecoin_AccrueFees_
	//	*AaveV2Stablecoin_ClaimAndUnstake_
	//	*AaveV2Stablecoin_EnterPosition_
	//	*AaveV2Stablecoin_Rebalance_
	//	*AaveV2Stablecoin_Reinvest_
	//	*AaveV2Stablecoin_Sweep_
	//	*AaveV2Stablecoin_TransferFees_
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

func (x *AaveV2Stablecoin) GetAccrueFees() *AaveV2Stablecoin_AccrueFees {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_AccrueFees_); ok {
		return x.AccrueFees
	}
	return nil
}

func (x *AaveV2Stablecoin) GetClaimAndUnstake() *AaveV2Stablecoin_ClaimAndUnstake {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_ClaimAndUnstake_); ok {
		return x.ClaimAndUnstake
	}
	return nil
}

func (x *AaveV2Stablecoin) GetEnterPosition() *AaveV2Stablecoin_EnterPosition {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_EnterPosition_); ok {
		return x.EnterPosition
	}
	return nil
}

func (x *AaveV2Stablecoin) GetRebalance() *AaveV2Stablecoin_Rebalance {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_Rebalance_); ok {
		return x.Rebalance
	}
	return nil
}

func (x *AaveV2Stablecoin) GetReinvest() *AaveV2Stablecoin_Reinvest {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_Reinvest_); ok {
		return x.Reinvest
	}
	return nil
}

func (x *AaveV2Stablecoin) GetSweep() *AaveV2Stablecoin_Sweep {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_Sweep_); ok {
		return x.Sweep
	}
	return nil
}

func (x *AaveV2Stablecoin) GetTransferFees() *AaveV2Stablecoin_TransferFees {
	if x, ok := x.GetFunction().(*AaveV2Stablecoin_TransferFees_); ok {
		return x.TransferFees
	}
	return nil
}

type isAaveV2Stablecoin_Function interface {
	isAaveV2Stablecoin_Function()
}

type AaveV2Stablecoin_AccrueFees_ struct {
	AccrueFees *AaveV2Stablecoin_AccrueFees `protobuf:"bytes,1,opt,name=accrue_fees,json=accrueFees,proto3,oneof"`
}

type AaveV2Stablecoin_ClaimAndUnstake_ struct {
	ClaimAndUnstake *AaveV2Stablecoin_ClaimAndUnstake `protobuf:"bytes,2,opt,name=claim_and_unstake,json=claimAndUnstake,proto3,oneof"`
}

type AaveV2Stablecoin_EnterPosition_ struct {
	EnterPosition *AaveV2Stablecoin_EnterPosition `protobuf:"bytes,3,opt,name=enter_position,json=enterPosition,proto3,oneof"`
}

type AaveV2Stablecoin_Rebalance_ struct {
	Rebalance *AaveV2Stablecoin_Rebalance `protobuf:"bytes,4,opt,name=rebalance,proto3,oneof"`
}

type AaveV2Stablecoin_Reinvest_ struct {
	Reinvest *AaveV2Stablecoin_Reinvest `protobuf:"bytes,5,opt,name=reinvest,proto3,oneof"`
}

type AaveV2Stablecoin_Sweep_ struct {
	Sweep *AaveV2Stablecoin_Sweep `protobuf:"bytes,6,opt,name=sweep,proto3,oneof"`
}

type AaveV2Stablecoin_TransferFees_ struct {
	TransferFees *AaveV2Stablecoin_TransferFees `protobuf:"bytes,7,opt,name=transfer_fees,json=transferFees,proto3,oneof"`
}

func (*AaveV2Stablecoin_AccrueFees_) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_ClaimAndUnstake_) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_EnterPosition_) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_Rebalance_) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_Reinvest_) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_Sweep_) isAaveV2Stablecoin_Function() {}

func (*AaveV2Stablecoin_TransferFees_) isAaveV2Stablecoin_Function() {}

//
// Take platform fees and performance fees off of cellar's active assets.
// Represents function accruePlatformFees()
type AaveV2Stablecoin_AccrueFees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AaveV2Stablecoin_AccrueFees) Reset() {
	*x = AaveV2Stablecoin_AccrueFees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_AccrueFees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_AccrueFees) ProtoMessage() {}

func (x *AaveV2Stablecoin_AccrueFees) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_AccrueFees.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_AccrueFees) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 0}
}

//
// Claim rewards from Aave and begin cooldown period to unstake them.
// Represents function claimAndUnstake()
type AaveV2Stablecoin_ClaimAndUnstake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AaveV2Stablecoin_ClaimAndUnstake) Reset() {
	*x = AaveV2Stablecoin_ClaimAndUnstake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_ClaimAndUnstake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_ClaimAndUnstake) ProtoMessage() {}

func (x *AaveV2Stablecoin_ClaimAndUnstake) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_ClaimAndUnstake.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_ClaimAndUnstake) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 1}
}

//
// Enters inactive assets into the current Aave stablecoin position.
// Represents function enterPosition()
type AaveV2Stablecoin_EnterPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AaveV2Stablecoin_EnterPosition) Reset() {
	*x = AaveV2Stablecoin_EnterPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_EnterPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_EnterPosition) ProtoMessage() {}

func (x *AaveV2Stablecoin_EnterPosition) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_EnterPosition.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_EnterPosition) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 2}
}

//
// Rebalances current assets into a new asset position.
// Represents function rebalance(address newLendingToken, uint256 minNewLendingTokenAmount).
// This function is based on the Curve Pool Registry exchange_multiple() function:
// https://github.com/curvefi/curve-pool-registry/blob/16a8664952cf61d7fed06acca79ad5ac696f4b20/contracts/Swaps.vy#L461-L489
type AaveV2Stablecoin_Rebalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of up to 9 addresses (4 swaps) representing a token swap route, where each triplet of addresses is a single swap route (ex. in token address, pool address, out token address)
	Route []string `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty"`
	// An array of up to 4 swap params. Attempting more than four swaps will fail.
	SwapParams []*AaveV2Stablecoin_Rebalance_SwapParams `protobuf:"bytes,2,rep,name=swap_params,json=swapParams,proto3" json:"swap_params,omitempty"`
	// Minimum acceptable assets to be received from the swap (slippage parameter)
	MinAssetsOut uint64 `protobuf:"varint,3,opt,name=min_assets_out,json=minAssetsOut,proto3" json:"min_assets_out,omitempty"`
}

func (x *AaveV2Stablecoin_Rebalance) Reset() {
	*x = AaveV2Stablecoin_Rebalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_Rebalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_Rebalance) ProtoMessage() {}

func (x *AaveV2Stablecoin_Rebalance) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_Rebalance.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_Rebalance) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 3}
}

func (x *AaveV2Stablecoin_Rebalance) GetRoute() []string {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *AaveV2Stablecoin_Rebalance) GetSwapParams() []*AaveV2Stablecoin_Rebalance_SwapParams {
	if x != nil {
		return x.SwapParams
	}
	return nil
}

func (x *AaveV2Stablecoin_Rebalance) GetMinAssetsOut() uint64 {
	if x != nil {
		return x.MinAssetsOut
	}
	return 0
}

//
// Reinvest rewards back into cellar's current position. Must be called within 2 day unstake period 10 days after `claimAndUnstake` was run.
// Represents function reinvest(uint256 minAssetsOut)
type AaveV2Stablecoin_Reinvest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum acceptable assets to be received from the swap (slippage parameter)
	MinAssetsOut uint64 `protobuf:"varint,1,opt,name=min_assets_out,json=minAssetsOut,proto3" json:"min_assets_out,omitempty"`
}

func (x *AaveV2Stablecoin_Reinvest) Reset() {
	*x = AaveV2Stablecoin_Reinvest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_Reinvest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_Reinvest) ProtoMessage() {}

func (x *AaveV2Stablecoin_Reinvest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_Reinvest.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_Reinvest) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 4}
}

func (x *AaveV2Stablecoin_Reinvest) GetMinAssetsOut() uint64 {
	if x != nil {
		return x.MinAssetsOut
	}
	return 0
}

//
// Sweep tokens sent here that are not managed by the cellar. This may be used in case the wrong tokens are accidentally sent to this contract.
// Represents function sweep(address).
type AaveV2Stablecoin_Sweep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the token to be transferred out of the Cellar
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The address to which the tokens should be transferred
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *AaveV2Stablecoin_Sweep) Reset() {
	*x = AaveV2Stablecoin_Sweep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_Sweep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_Sweep) ProtoMessage() {}

func (x *AaveV2Stablecoin_Sweep) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_Sweep.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_Sweep) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 5}
}

func (x *AaveV2Stablecoin_Sweep) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AaveV2Stablecoin_Sweep) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

//
// Transfer accrued fees to the Sommelier Chain to distribute.
// Represents function transferFees()
type AaveV2Stablecoin_TransferFees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AaveV2Stablecoin_TransferFees) Reset() {
	*x = AaveV2Stablecoin_TransferFees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_TransferFees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_TransferFees) ProtoMessage() {}

func (x *AaveV2Stablecoin_TransferFees) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_TransferFees.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_TransferFees) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 6}
}

//
// Represents parameters for a single swap. Each swap needs the indeces in Rebalance.route of the in/out token addresses and the swap type. See the Curve contract linked above for more detail.
type AaveV2Stablecoin_Rebalance_SwapParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index in the `route` array of the swap's input token address
	InIndex uint64 `protobuf:"varint,1,opt,name=in_index,json=inIndex,proto3" json:"in_index,omitempty"`
	// Index in the `route` array of the swap's output token address
	OutIndex uint64 `protobuf:"varint,2,opt,name=out_index,json=outIndex,proto3" json:"out_index,omitempty"`
	// 1 - stableswap `exchange`
	// 2 - stableswap `exchange_underlying`
	// 3 - cryptoswap `exchange`
	// 4 - cryptoswap `exchange_underlying`
	// 5 - Polygon factory metapools `exchange_underlying`
	// See the Curve Pool Registry exchange_multiple() function for more information.
	SwapType uint64 `protobuf:"varint,3,opt,name=swap_type,json=swapType,proto3" json:"swap_type,omitempty"`
}

func (x *AaveV2Stablecoin_Rebalance_SwapParams) Reset() {
	*x = AaveV2Stablecoin_Rebalance_SwapParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aave_v2_stablecoin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AaveV2Stablecoin_Rebalance_SwapParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AaveV2Stablecoin_Rebalance_SwapParams) ProtoMessage() {}

func (x *AaveV2Stablecoin_Rebalance_SwapParams) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AaveV2Stablecoin_Rebalance_SwapParams.ProtoReflect.Descriptor instead.
func (*AaveV2Stablecoin_Rebalance_SwapParams) Descriptor() ([]byte, []int) {
	return file_aave_v2_stablecoin_proto_rawDescGZIP(), []int{0, 3, 0}
}

func (x *AaveV2Stablecoin_Rebalance_SwapParams) GetInIndex() uint64 {
	if x != nil {
		return x.InIndex
	}
	return 0
}

func (x *AaveV2Stablecoin_Rebalance_SwapParams) GetOutIndex() uint64 {
	if x != nil {
		return x.OutIndex
	}
	return 0
}

func (x *AaveV2Stablecoin_Rebalance_SwapParams) GetSwapType() uint64 {
	if x != nil {
		return x.SwapType
	}
	return 0
}

var File_aave_v2_stablecoin_proto protoreflect.FileDescriptor

var file_aave_v2_stablecoin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x61, 0x76, 0x65, 0x5f, 0x76, 0x32, 0x5f, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xda, 0x07, 0x0a, 0x10, 0x41, 0x61, 0x76, 0x65, 0x56,
	0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x4a, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x72, 0x75, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61,
	0x76, 0x65, 0x56, 0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x41,
	0x63, 0x63, 0x72, 0x75, 0x65, 0x46, 0x65, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63, 0x63,
	0x72, 0x75, 0x65, 0x46, 0x65, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x11, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x41, 0x6e, 0x64, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65,
	0x48, 0x00, 0x52, 0x0f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x41, 0x6e, 0x64, 0x55, 0x6e, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x53,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x09, 0x72, 0x65, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x53,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x08, 0x72, 0x65, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x52, 0x65, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x73, 0x77, 0x65, 0x65, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x53, 0x77, 0x65, 0x65, 0x70, 0x48, 0x00, 0x52, 0x05, 0x73, 0x77, 0x65, 0x65,
	0x70, 0x12, 0x50, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x66, 0x65,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x76, 0x65, 0x56, 0x32, 0x53, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46,
	0x65, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46,
	0x65, 0x65, 0x73, 0x1a, 0x0c, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x72, 0x75, 0x65, 0x46, 0x65, 0x65,
	0x73, 0x1a, 0x11, 0x0a, 0x0f, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x41, 0x6e, 0x64, 0x55, 0x6e, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x1a, 0x0f, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xfe, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x73, 0x77, 0x61,
	0x70, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x61, 0x76, 0x65,
	0x56, 0x32, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x0a, 0x73, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a,
	0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x4f, 0x75, 0x74, 0x1a, 0x61, 0x0a, 0x0a, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x69, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x6f, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x77, 0x61,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x77,
	0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x69, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x5f, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x1a, 0x2d, 0x0a, 0x05, 0x53, 0x77, 0x65, 0x65,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x1a, 0x0e, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x65, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_aave_v2_stablecoin_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_aave_v2_stablecoin_proto_goTypes = []interface{}{
	(*AaveV2Stablecoin)(nil),                      // 0: steward.v1.AaveV2Stablecoin
	(*AaveV2Stablecoin_AccrueFees)(nil),           // 1: steward.v1.AaveV2Stablecoin.AccrueFees
	(*AaveV2Stablecoin_ClaimAndUnstake)(nil),      // 2: steward.v1.AaveV2Stablecoin.ClaimAndUnstake
	(*AaveV2Stablecoin_EnterPosition)(nil),        // 3: steward.v1.AaveV2Stablecoin.EnterPosition
	(*AaveV2Stablecoin_Rebalance)(nil),            // 4: steward.v1.AaveV2Stablecoin.Rebalance
	(*AaveV2Stablecoin_Reinvest)(nil),             // 5: steward.v1.AaveV2Stablecoin.Reinvest
	(*AaveV2Stablecoin_Sweep)(nil),                // 6: steward.v1.AaveV2Stablecoin.Sweep
	(*AaveV2Stablecoin_TransferFees)(nil),         // 7: steward.v1.AaveV2Stablecoin.TransferFees
	(*AaveV2Stablecoin_Rebalance_SwapParams)(nil), // 8: steward.v1.AaveV2Stablecoin.Rebalance.SwapParams
}
var file_aave_v2_stablecoin_proto_depIdxs = []int32{
	1, // 0: steward.v1.AaveV2Stablecoin.accrue_fees:type_name -> steward.v1.AaveV2Stablecoin.AccrueFees
	2, // 1: steward.v1.AaveV2Stablecoin.claim_and_unstake:type_name -> steward.v1.AaveV2Stablecoin.ClaimAndUnstake
	3, // 2: steward.v1.AaveV2Stablecoin.enter_position:type_name -> steward.v1.AaveV2Stablecoin.EnterPosition
	4, // 3: steward.v1.AaveV2Stablecoin.rebalance:type_name -> steward.v1.AaveV2Stablecoin.Rebalance
	5, // 4: steward.v1.AaveV2Stablecoin.reinvest:type_name -> steward.v1.AaveV2Stablecoin.Reinvest
	6, // 5: steward.v1.AaveV2Stablecoin.sweep:type_name -> steward.v1.AaveV2Stablecoin.Sweep
	7, // 6: steward.v1.AaveV2Stablecoin.transfer_fees:type_name -> steward.v1.AaveV2Stablecoin.TransferFees
	8, // 7: steward.v1.AaveV2Stablecoin.Rebalance.swap_params:type_name -> steward.v1.AaveV2Stablecoin.Rebalance.SwapParams
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
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
			switch v := v.(*AaveV2Stablecoin_AccrueFees); i {
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
			switch v := v.(*AaveV2Stablecoin_ClaimAndUnstake); i {
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
			switch v := v.(*AaveV2Stablecoin_EnterPosition); i {
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
			switch v := v.(*AaveV2Stablecoin_Rebalance); i {
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
			switch v := v.(*AaveV2Stablecoin_Reinvest); i {
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
			switch v := v.(*AaveV2Stablecoin_Sweep); i {
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
			switch v := v.(*AaveV2Stablecoin_TransferFees); i {
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
			switch v := v.(*AaveV2Stablecoin_Rebalance_SwapParams); i {
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
		(*AaveV2Stablecoin_AccrueFees_)(nil),
		(*AaveV2Stablecoin_ClaimAndUnstake_)(nil),
		(*AaveV2Stablecoin_EnterPosition_)(nil),
		(*AaveV2Stablecoin_Rebalance_)(nil),
		(*AaveV2Stablecoin_Reinvest_)(nil),
		(*AaveV2Stablecoin_Sweep_)(nil),
		(*AaveV2Stablecoin_TransferFees_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aave_v2_stablecoin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
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
