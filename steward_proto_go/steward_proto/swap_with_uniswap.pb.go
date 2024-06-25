// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: swap_with_uniswap.proto

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

// Represents call data for the Uniswap V3 adaptor
type SwapWithUniswapAdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//
	//	*SwapWithUniswapAdaptorV1_RevokeApproval
	//	*SwapWithUniswapAdaptorV1_SwapWithUniV2_
	//	*SwapWithUniswapAdaptorV1_SwapWithUniV3_
	Function isSwapWithUniswapAdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *SwapWithUniswapAdaptorV1) Reset() {
	*x = SwapWithUniswapAdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swap_with_uniswap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapWithUniswapAdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapWithUniswapAdaptorV1) ProtoMessage() {}

func (x *SwapWithUniswapAdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_swap_with_uniswap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapWithUniswapAdaptorV1.ProtoReflect.Descriptor instead.
func (*SwapWithUniswapAdaptorV1) Descriptor() ([]byte, []int) {
	return file_swap_with_uniswap_proto_rawDescGZIP(), []int{0}
}

func (m *SwapWithUniswapAdaptorV1) GetFunction() isSwapWithUniswapAdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *SwapWithUniswapAdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*SwapWithUniswapAdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *SwapWithUniswapAdaptorV1) GetSwapWithUniV2() *SwapWithUniswapAdaptorV1_SwapWithUniV2 {
	if x, ok := x.GetFunction().(*SwapWithUniswapAdaptorV1_SwapWithUniV2_); ok {
		return x.SwapWithUniV2
	}
	return nil
}

func (x *SwapWithUniswapAdaptorV1) GetSwapWithUniV3() *SwapWithUniswapAdaptorV1_SwapWithUniV3 {
	if x, ok := x.GetFunction().(*SwapWithUniswapAdaptorV1_SwapWithUniV3_); ok {
		return x.SwapWithUniV3
	}
	return nil
}

type isSwapWithUniswapAdaptorV1_Function interface {
	isSwapWithUniswapAdaptorV1_Function()
}

type SwapWithUniswapAdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type SwapWithUniswapAdaptorV1_SwapWithUniV2_ struct {
	// Represents function `swapWithUniV2(address[] path, uint256 amount, uint256 amountOutMin)`
	SwapWithUniV2 *SwapWithUniswapAdaptorV1_SwapWithUniV2 `protobuf:"bytes,2,opt,name=swap_with_uni_v2,json=swapWithUniV2,proto3,oneof"`
}

type SwapWithUniswapAdaptorV1_SwapWithUniV3_ struct {
	// Represents function `swapWithUniV3(address[] path, uint24[] poolFees, uint256 amount, uint256 amountOutMin)`
	SwapWithUniV3 *SwapWithUniswapAdaptorV1_SwapWithUniV3 `protobuf:"bytes,3,opt,name=swap_with_uni_v3,json=swapWithUniV3,proto3,oneof"`
}

func (*SwapWithUniswapAdaptorV1_RevokeApproval) isSwapWithUniswapAdaptorV1_Function() {}

func (*SwapWithUniswapAdaptorV1_SwapWithUniV2_) isSwapWithUniswapAdaptorV1_Function() {}

func (*SwapWithUniswapAdaptorV1_SwapWithUniV3_) isSwapWithUniswapAdaptorV1_Function() {}

type SwapWithUniswapAdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*SwapWithUniswapAdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *SwapWithUniswapAdaptorV1Calls) Reset() {
	*x = SwapWithUniswapAdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swap_with_uniswap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapWithUniswapAdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapWithUniswapAdaptorV1Calls) ProtoMessage() {}

func (x *SwapWithUniswapAdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_swap_with_uniswap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapWithUniswapAdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*SwapWithUniswapAdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_swap_with_uniswap_proto_rawDescGZIP(), []int{1}
}

func (x *SwapWithUniswapAdaptorV1Calls) GetCalls() []*SwapWithUniswapAdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

// Perform a swap using Uniswap V2.
//
// Represents function `swapWithUniV2(address[] path, uint256 amount, uint256 amountOutMin)`
type SwapWithUniswapAdaptorV1_SwapWithUniV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path         []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Amount       string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountOutMin string   `protobuf:"bytes,3,opt,name=amount_out_min,json=amountOutMin,proto3" json:"amount_out_min,omitempty"`
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV2) Reset() {
	*x = SwapWithUniswapAdaptorV1_SwapWithUniV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swap_with_uniswap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapWithUniswapAdaptorV1_SwapWithUniV2) ProtoMessage() {}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV2) ProtoReflect() protoreflect.Message {
	mi := &file_swap_with_uniswap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapWithUniswapAdaptorV1_SwapWithUniV2.ProtoReflect.Descriptor instead.
func (*SwapWithUniswapAdaptorV1_SwapWithUniV2) Descriptor() ([]byte, []int) {
	return file_swap_with_uniswap_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV2) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV2) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV2) GetAmountOutMin() string {
	if x != nil {
		return x.AmountOutMin
	}
	return ""
}

// Perform a swap using Uniswap V3.
//
// Represents function `Represents function `swapWithUniV3(address[] path, uint24[] poolFees, uint256 amount, uint256 amountOutMin)“
type SwapWithUniswapAdaptorV1_SwapWithUniV3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path         []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	PoolFees     []uint32 `protobuf:"varint,2,rep,packed,name=pool_fees,json=poolFees,proto3" json:"pool_fees,omitempty"`
	Amount       string   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountOutMin string   `protobuf:"bytes,4,opt,name=amount_out_min,json=amountOutMin,proto3" json:"amount_out_min,omitempty"`
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV3) Reset() {
	*x = SwapWithUniswapAdaptorV1_SwapWithUniV3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swap_with_uniswap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapWithUniswapAdaptorV1_SwapWithUniV3) ProtoMessage() {}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV3) ProtoReflect() protoreflect.Message {
	mi := &file_swap_with_uniswap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapWithUniswapAdaptorV1_SwapWithUniV3.ProtoReflect.Descriptor instead.
func (*SwapWithUniswapAdaptorV1_SwapWithUniV3) Descriptor() ([]byte, []int) {
	return file_swap_with_uniswap_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV3) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV3) GetPoolFees() []uint32 {
	if x != nil {
		return x.PoolFees
	}
	return nil
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV3) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *SwapWithUniswapAdaptorV1_SwapWithUniV3) GetAmountOutMin() string {
	if x != nil {
		return x.AmountOutMin
	}
	return ""
}

var File_swap_with_uniswap_proto protoreflect.FileDescriptor

var file_swap_with_uniswap_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75, 0x6e, 0x69, 0x73,
	0x77, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x34, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8e, 0x04, 0x0a, 0x18, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e,
	0x69, 0x73, 0x77, 0x61, 0x70, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x45,
	0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x5d, 0x0a, 0x10, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x75, 0x6e, 0x69, 0x5f, 0x76, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x53, 0x77, 0x61,
	0x70, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x6f, 0x72, 0x56, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e,
	0x69, 0x56, 0x32, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55,
	0x6e, 0x69, 0x56, 0x32, 0x12, 0x5d, 0x0a, 0x10, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x5f, 0x75, 0x6e, 0x69, 0x5f, 0x76, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x53, 0x77, 0x61, 0x70,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x6f, 0x72, 0x56, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x69,
	0x56, 0x33, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e,
	0x69, 0x56, 0x33, 0x1a, 0x61, 0x0a, 0x0d, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55,
	0x6e, 0x69, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x4f, 0x75, 0x74, 0x4d, 0x69, 0x6e, 0x1a, 0x7e, 0x0a, 0x0d, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69,
	0x74, 0x68, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x6f, 0x6f, 0x6c, 0x46, 0x65, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x4f, 0x75, 0x74, 0x4d, 0x69, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x1d, 0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e,
	0x69, 0x73, 0x77, 0x61, 0x70, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x43, 0x61,
	0x6c, 0x6c, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e,
	0x53, 0x77, 0x61, 0x70, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42,
	0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_swap_with_uniswap_proto_rawDescOnce sync.Once
	file_swap_with_uniswap_proto_rawDescData = file_swap_with_uniswap_proto_rawDesc
)

func file_swap_with_uniswap_proto_rawDescGZIP() []byte {
	file_swap_with_uniswap_proto_rawDescOnce.Do(func() {
		file_swap_with_uniswap_proto_rawDescData = protoimpl.X.CompressGZIP(file_swap_with_uniswap_proto_rawDescData)
	})
	return file_swap_with_uniswap_proto_rawDescData
}

var file_swap_with_uniswap_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_swap_with_uniswap_proto_goTypes = []interface{}{
	(*SwapWithUniswapAdaptorV1)(nil),               // 0: steward.v4.SwapWithUniswapAdaptorV1
	(*SwapWithUniswapAdaptorV1Calls)(nil),          // 1: steward.v4.SwapWithUniswapAdaptorV1Calls
	(*SwapWithUniswapAdaptorV1_SwapWithUniV2)(nil), // 2: steward.v4.SwapWithUniswapAdaptorV1.SwapWithUniV2
	(*SwapWithUniswapAdaptorV1_SwapWithUniV3)(nil), // 3: steward.v4.SwapWithUniswapAdaptorV1.SwapWithUniV3
	(*RevokeApproval)(nil),                         // 4: steward.v4.RevokeApproval
}
var file_swap_with_uniswap_proto_depIdxs = []int32{
	4, // 0: steward.v4.SwapWithUniswapAdaptorV1.revoke_approval:type_name -> steward.v4.RevokeApproval
	2, // 1: steward.v4.SwapWithUniswapAdaptorV1.swap_with_uni_v2:type_name -> steward.v4.SwapWithUniswapAdaptorV1.SwapWithUniV2
	3, // 2: steward.v4.SwapWithUniswapAdaptorV1.swap_with_uni_v3:type_name -> steward.v4.SwapWithUniswapAdaptorV1.SwapWithUniV3
	0, // 3: steward.v4.SwapWithUniswapAdaptorV1Calls.calls:type_name -> steward.v4.SwapWithUniswapAdaptorV1
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_swap_with_uniswap_proto_init() }
func file_swap_with_uniswap_proto_init() {
	if File_swap_with_uniswap_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_swap_with_uniswap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapWithUniswapAdaptorV1); i {
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
		file_swap_with_uniswap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapWithUniswapAdaptorV1Calls); i {
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
		file_swap_with_uniswap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapWithUniswapAdaptorV1_SwapWithUniV2); i {
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
		file_swap_with_uniswap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapWithUniswapAdaptorV1_SwapWithUniV3); i {
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
	file_swap_with_uniswap_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SwapWithUniswapAdaptorV1_RevokeApproval)(nil),
		(*SwapWithUniswapAdaptorV1_SwapWithUniV2_)(nil),
		(*SwapWithUniswapAdaptorV1_SwapWithUniV3_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_swap_with_uniswap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_swap_with_uniswap_proto_goTypes,
		DependencyIndexes: file_swap_with_uniswap_proto_depIdxs,
		MessageInfos:      file_swap_with_uniswap_proto_msgTypes,
	}.Build()
	File_swap_with_uniswap_proto = out.File
	file_swap_with_uniswap_proto_rawDesc = nil
	file_swap_with_uniswap_proto_goTypes = nil
	file_swap_with_uniswap_proto_depIdxs = nil
}
