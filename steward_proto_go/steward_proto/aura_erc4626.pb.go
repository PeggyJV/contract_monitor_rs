//
// Protos for function calls to the Aura ERC4626 adaptor.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: aura_erc4626.proto

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

// Represents call data for the Aura ERC4626 adaptor V1
type AuraERC4626AdaptorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*AuraERC4626AdaptorV1_RevokeApproval
	//	*AuraERC4626AdaptorV1_GetRewards_
	Function isAuraERC4626AdaptorV1_Function `protobuf_oneof:"function"`
}

func (x *AuraERC4626AdaptorV1) Reset() {
	*x = AuraERC4626AdaptorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aura_erc4626_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuraERC4626AdaptorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuraERC4626AdaptorV1) ProtoMessage() {}

func (x *AuraERC4626AdaptorV1) ProtoReflect() protoreflect.Message {
	mi := &file_aura_erc4626_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuraERC4626AdaptorV1.ProtoReflect.Descriptor instead.
func (*AuraERC4626AdaptorV1) Descriptor() ([]byte, []int) {
	return file_aura_erc4626_proto_rawDescGZIP(), []int{0}
}

func (m *AuraERC4626AdaptorV1) GetFunction() isAuraERC4626AdaptorV1_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *AuraERC4626AdaptorV1) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*AuraERC4626AdaptorV1_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *AuraERC4626AdaptorV1) GetGetRewards() *AuraERC4626AdaptorV1_GetRewards {
	if x, ok := x.GetFunction().(*AuraERC4626AdaptorV1_GetRewards_); ok {
		return x.GetRewards
	}
	return nil
}

type isAuraERC4626AdaptorV1_Function interface {
	isAuraERC4626AdaptorV1_Function()
}

type AuraERC4626AdaptorV1_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,1,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type AuraERC4626AdaptorV1_GetRewards_ struct {
	// Represents function `getRewards(IBaseRewardPool _auraPool, bool _claimExtras)`
	GetRewards *AuraERC4626AdaptorV1_GetRewards `protobuf:"bytes,2,opt,name=get_rewards,json=getRewards,proto3,oneof"`
}

func (*AuraERC4626AdaptorV1_RevokeApproval) isAuraERC4626AdaptorV1_Function() {}

func (*AuraERC4626AdaptorV1_GetRewards_) isAuraERC4626AdaptorV1_Function() {}

type AuraERC4626AdaptorV1Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*AuraERC4626AdaptorV1 `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *AuraERC4626AdaptorV1Calls) Reset() {
	*x = AuraERC4626AdaptorV1Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aura_erc4626_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuraERC4626AdaptorV1Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuraERC4626AdaptorV1Calls) ProtoMessage() {}

func (x *AuraERC4626AdaptorV1Calls) ProtoReflect() protoreflect.Message {
	mi := &file_aura_erc4626_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuraERC4626AdaptorV1Calls.ProtoReflect.Descriptor instead.
func (*AuraERC4626AdaptorV1Calls) Descriptor() ([]byte, []int) {
	return file_aura_erc4626_proto_rawDescGZIP(), []int{1}
}

func (x *AuraERC4626AdaptorV1Calls) GetCalls() []*AuraERC4626AdaptorV1 {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows strategist to get rewards for an Aura pool.
//
// Represents function `getRewards(IBaseRewardPool _auraPool, bool _claimExtras)`
type AuraERC4626AdaptorV1_GetRewards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the Aura pool to get rewards for
	AuraPool string `protobuf:"bytes,1,opt,name=aura_pool,json=auraPool,proto3" json:"aura_pool,omitempty"`
	// Whether to claim extra rewards associated with the pool
	ClaimExtras bool `protobuf:"varint,2,opt,name=claim_extras,json=claimExtras,proto3" json:"claim_extras,omitempty"`
}

func (x *AuraERC4626AdaptorV1_GetRewards) Reset() {
	*x = AuraERC4626AdaptorV1_GetRewards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aura_erc4626_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuraERC4626AdaptorV1_GetRewards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuraERC4626AdaptorV1_GetRewards) ProtoMessage() {}

func (x *AuraERC4626AdaptorV1_GetRewards) ProtoReflect() protoreflect.Message {
	mi := &file_aura_erc4626_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuraERC4626AdaptorV1_GetRewards.ProtoReflect.Descriptor instead.
func (*AuraERC4626AdaptorV1_GetRewards) Descriptor() ([]byte, []int) {
	return file_aura_erc4626_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AuraERC4626AdaptorV1_GetRewards) GetAuraPool() string {
	if x != nil {
		return x.AuraPool
	}
	return ""
}

func (x *AuraERC4626AdaptorV1_GetRewards) GetClaimExtras() bool {
	if x != nil {
		return x.ClaimExtras
	}
	return false
}

var File_aura_erc4626_proto protoreflect.FileDescriptor

var file_aura_erc4626_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x72, 0x61, 0x5f, 0x65, 0x72, 0x63, 0x34, 0x36, 0x32, 0x36, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34,
	0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a,
	0x14, 0x41, 0x75, 0x72, 0x61, 0x45, 0x52, 0x43, 0x34, 0x36, 0x32, 0x36, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x45, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x4e, 0x0a, 0x0b,
	0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e, 0x41,
	0x75, 0x72, 0x61, 0x45, 0x52, 0x43, 0x34, 0x36, 0x32, 0x36, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x56, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x48, 0x00,
	0x52, 0x0a, 0x67, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x4c, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x72, 0x61, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x72, 0x61, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x45, 0x78, 0x74, 0x72, 0x61, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x19, 0x41, 0x75, 0x72, 0x61, 0x45, 0x52,
	0x43, 0x34, 0x36, 0x32, 0x36, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x43, 0x61,
	0x6c, 0x6c, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x34, 0x2e,
	0x41, 0x75, 0x72, 0x61, 0x45, 0x52, 0x43, 0x34, 0x36, 0x32, 0x36, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x6f, 0x72, 0x56, 0x31, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f,
	0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aura_erc4626_proto_rawDescOnce sync.Once
	file_aura_erc4626_proto_rawDescData = file_aura_erc4626_proto_rawDesc
)

func file_aura_erc4626_proto_rawDescGZIP() []byte {
	file_aura_erc4626_proto_rawDescOnce.Do(func() {
		file_aura_erc4626_proto_rawDescData = protoimpl.X.CompressGZIP(file_aura_erc4626_proto_rawDescData)
	})
	return file_aura_erc4626_proto_rawDescData
}

var file_aura_erc4626_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aura_erc4626_proto_goTypes = []interface{}{
	(*AuraERC4626AdaptorV1)(nil),            // 0: steward.v4.AuraERC4626AdaptorV1
	(*AuraERC4626AdaptorV1Calls)(nil),       // 1: steward.v4.AuraERC4626AdaptorV1Calls
	(*AuraERC4626AdaptorV1_GetRewards)(nil), // 2: steward.v4.AuraERC4626AdaptorV1.GetRewards
	(*RevokeApproval)(nil),                  // 3: steward.v4.RevokeApproval
}
var file_aura_erc4626_proto_depIdxs = []int32{
	3, // 0: steward.v4.AuraERC4626AdaptorV1.revoke_approval:type_name -> steward.v4.RevokeApproval
	2, // 1: steward.v4.AuraERC4626AdaptorV1.get_rewards:type_name -> steward.v4.AuraERC4626AdaptorV1.GetRewards
	0, // 2: steward.v4.AuraERC4626AdaptorV1Calls.calls:type_name -> steward.v4.AuraERC4626AdaptorV1
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aura_erc4626_proto_init() }
func file_aura_erc4626_proto_init() {
	if File_aura_erc4626_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aura_erc4626_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuraERC4626AdaptorV1); i {
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
		file_aura_erc4626_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuraERC4626AdaptorV1Calls); i {
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
		file_aura_erc4626_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuraERC4626AdaptorV1_GetRewards); i {
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
	file_aura_erc4626_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AuraERC4626AdaptorV1_RevokeApproval)(nil),
		(*AuraERC4626AdaptorV1_GetRewards_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aura_erc4626_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aura_erc4626_proto_goTypes,
		DependencyIndexes: file_aura_erc4626_proto_depIdxs,
		MessageInfos:      file_aura_erc4626_proto_msgTypes,
	}.Build()
	File_aura_erc4626_proto = out.File
	file_aura_erc4626_proto_rawDesc = nil
	file_aura_erc4626_proto_goTypes = nil
	file_aura_erc4626_proto_depIdxs = nil
}
