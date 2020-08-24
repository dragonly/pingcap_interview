// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: pkg/cluster/api.proto

package cluster

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TopNInBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataBlock *DataBlock `protobuf:"bytes,1,opt,name=data_block,json=dataBlock,proto3" json:"data_block,omitempty"`
	KeyRange  *KeyRange  `protobuf:"bytes,2,opt,name=key_range,json=keyRange,proto3" json:"key_range,omitempty"`
}

func (x *TopNInBlockRequest) Reset() {
	*x = TopNInBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cluster_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNInBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNInBlockRequest) ProtoMessage() {}

func (x *TopNInBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNInBlockRequest.ProtoReflect.Descriptor instead.
func (*TopNInBlockRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cluster_api_proto_rawDescGZIP(), []int{0}
}

func (x *TopNInBlockRequest) GetDataBlock() *DataBlock {
	if x != nil {
		return x.DataBlock
	}
	return nil
}

func (x *TopNInBlockRequest) GetKeyRange() *KeyRange {
	if x != nil {
		return x.KeyRange
	}
	return nil
}

type KeyRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxKey int64 `protobuf:"varint,1,opt,name=max_key,json=maxKey,proto3" json:"max_key,omitempty"`
	MinKey int64 `protobuf:"varint,2,opt,name=min_key,json=minKey,proto3" json:"min_key,omitempty"`
}

func (x *KeyRange) Reset() {
	*x = KeyRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cluster_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRange) ProtoMessage() {}

func (x *KeyRange) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRange.ProtoReflect.Descriptor instead.
func (*KeyRange) Descriptor() ([]byte, []int) {
	return file_pkg_cluster_api_proto_rawDescGZIP(), []int{1}
}

func (x *KeyRange) GetMaxKey() int64 {
	if x != nil {
		return x.MaxKey
	}
	return 0
}

func (x *KeyRange) GetMinKey() int64 {
	if x != nil {
		return x.MinKey
	}
	return 0
}

type DataBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	BlockIndex int64  `protobuf:"varint,2,opt,name=block_index,json=blockIndex,proto3" json:"block_index,omitempty"`
}

func (x *DataBlock) Reset() {
	*x = DataBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cluster_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBlock) ProtoMessage() {}

func (x *DataBlock) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBlock.ProtoReflect.Descriptor instead.
func (*DataBlock) Descriptor() ([]byte, []int) {
	return file_pkg_cluster_api_proto_rawDescGZIP(), []int{2}
}

func (x *DataBlock) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DataBlock) GetBlockIndex() int64 {
	if x != nil {
		return x.BlockIndex
	}
	return 0
}

type TopNInBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *TopNInBlockResponse) Reset() {
	*x = TopNInBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cluster_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNInBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNInBlockResponse) ProtoMessage() {}

func (x *TopNInBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNInBlockResponse.ProtoReflect.Descriptor instead.
func (*TopNInBlockResponse) Descriptor() ([]byte, []int) {
	return file_pkg_cluster_api_proto_rawDescGZIP(), []int{3}
}

func (x *TopNInBlockResponse) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cluster_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_pkg_cluster_api_proto_rawDescGZIP(), []int{4}
}

func (x *Record) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *Record) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_cluster_api_proto protoreflect.FileDescriptor

var file_pkg_cluster_api_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x22,
	0x75, 0x0a, 0x12, 0x54, 0x6f, 0x70, 0x4e, 0x49, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2d, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x69, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x69,
	0x6e, 0x4b, 0x65, 0x79, 0x22, 0x48, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x3f,
	0x0a, 0x13, 0x54, 0x6f, 0x70, 0x4e, 0x49, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0x2e, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x50, 0x0a, 0x04, 0x54, 0x6f, 0x70, 0x4e, 0x12, 0x48, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x4e, 0x49,
	0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x54, 0x6f, 0x70, 0x4e, 0x49, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x70, 0x4e,
	0x49, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x6c, 0x79, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x70,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_cluster_api_proto_rawDescOnce sync.Once
	file_pkg_cluster_api_proto_rawDescData = file_pkg_cluster_api_proto_rawDesc
)

func file_pkg_cluster_api_proto_rawDescGZIP() []byte {
	file_pkg_cluster_api_proto_rawDescOnce.Do(func() {
		file_pkg_cluster_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_cluster_api_proto_rawDescData)
	})
	return file_pkg_cluster_api_proto_rawDescData
}

var file_pkg_cluster_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_cluster_api_proto_goTypes = []interface{}{
	(*TopNInBlockRequest)(nil),  // 0: mapper.TopNInBlockRequest
	(*KeyRange)(nil),            // 1: mapper.KeyRange
	(*DataBlock)(nil),           // 2: mapper.DataBlock
	(*TopNInBlockResponse)(nil), // 3: mapper.TopNInBlockResponse
	(*Record)(nil),              // 4: mapper.Record
}
var file_pkg_cluster_api_proto_depIdxs = []int32{
	2, // 0: mapper.TopNInBlockRequest.data_block:type_name -> mapper.DataBlock
	1, // 1: mapper.TopNInBlockRequest.key_range:type_name -> mapper.KeyRange
	4, // 2: mapper.TopNInBlockResponse.records:type_name -> mapper.Record
	0, // 3: mapper.TopN.TopNInBlock:input_type -> mapper.TopNInBlockRequest
	3, // 4: mapper.TopN.TopNInBlock:output_type -> mapper.TopNInBlockResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_cluster_api_proto_init() }
func file_pkg_cluster_api_proto_init() {
	if File_pkg_cluster_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_cluster_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNInBlockRequest); i {
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
		file_pkg_cluster_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRange); i {
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
		file_pkg_cluster_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBlock); i {
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
		file_pkg_cluster_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNInBlockResponse); i {
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
		file_pkg_cluster_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_cluster_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_cluster_api_proto_goTypes,
		DependencyIndexes: file_pkg_cluster_api_proto_depIdxs,
		MessageInfos:      file_pkg_cluster_api_proto_msgTypes,
	}.Build()
	File_pkg_cluster_api_proto = out.File
	file_pkg_cluster_api_proto_rawDesc = nil
	file_pkg_cluster_api_proto_goTypes = nil
	file_pkg_cluster_api_proto_depIdxs = nil
}