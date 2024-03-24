// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: metadata/metadatapb/rpc.proto

package metadatapb

import (
	storepb "github.com/thanos-io/thanos/pkg/store/storepb"
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

type MetricMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric                  string                          `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	Limit                   int32                           `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	PartialResponseStrategy storepb.PartialResponseStrategy `protobuf:"varint,3,opt,name=partial_response_strategy,json=partialResponseStrategy,proto3,enum=thanos.PartialResponseStrategy" json:"partial_response_strategy,omitempty"`
}

func (x *MetricMetadataRequest) Reset() {
	*x = MetricMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadatapb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMetadataRequest) ProtoMessage() {}

func (x *MetricMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadatapb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMetadataRequest.ProtoReflect.Descriptor instead.
func (*MetricMetadataRequest) Descriptor() ([]byte, []int) {
	return file_metadata_metadatapb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *MetricMetadataRequest) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *MetricMetadataRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *MetricMetadataRequest) GetPartialResponseStrategy() storepb.PartialResponseStrategy {
	if x != nil {
		return x.PartialResponseStrategy
	}
	return storepb.PartialResponseStrategy(0)
}

type MetricMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*MetricMetadataResponse_Metadata
	//	*MetricMetadataResponse_Warning
	Result isMetricMetadataResponse_Result `protobuf_oneof:"result"`
}

func (x *MetricMetadataResponse) Reset() {
	*x = MetricMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadatapb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMetadataResponse) ProtoMessage() {}

func (x *MetricMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadatapb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMetadataResponse.ProtoReflect.Descriptor instead.
func (*MetricMetadataResponse) Descriptor() ([]byte, []int) {
	return file_metadata_metadatapb_rpc_proto_rawDescGZIP(), []int{1}
}

func (m *MetricMetadataResponse) GetResult() isMetricMetadataResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *MetricMetadataResponse) GetMetadata() *MetricMetadata {
	if x, ok := x.GetResult().(*MetricMetadataResponse_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *MetricMetadataResponse) GetWarning() string {
	if x, ok := x.GetResult().(*MetricMetadataResponse_Warning); ok {
		return x.Warning
	}
	return ""
}

type isMetricMetadataResponse_Result interface {
	isMetricMetadataResponse_Result()
}

type MetricMetadataResponse_Metadata struct {
	// / A collection of metric metadata entries.
	Metadata *MetricMetadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type MetricMetadataResponse_Warning struct {
	// / warning is considered an information piece in place of series for warning purposes.
	// / It is used to warn metadata API users about suspicious cases or partial response (if enabled).
	Warning string `protobuf:"bytes,2,opt,name=warning,proto3,oneof"`
}

func (*MetricMetadataResponse_Metadata) isMetricMetadataResponse_Result() {}

func (*MetricMetadataResponse_Warning) isMetricMetadataResponse_Result() {}

type MetricMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]*MetricMetadataEntry `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetricMetadata) Reset() {
	*x = MetricMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadatapb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMetadata) ProtoMessage() {}

func (x *MetricMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadatapb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMetadata.ProtoReflect.Descriptor instead.
func (*MetricMetadata) Descriptor() ([]byte, []int) {
	return file_metadata_metadatapb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *MetricMetadata) GetMetadata() map[string]*MetricMetadataEntry {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MetricMetadataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metas []*Meta `protobuf:"bytes,2,rep,name=metas,proto3" json:"metas,omitempty"`
}

func (x *MetricMetadataEntry) Reset() {
	*x = MetricMetadataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadatapb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricMetadataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMetadataEntry) ProtoMessage() {}

func (x *MetricMetadataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadatapb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMetadataEntry.ProtoReflect.Descriptor instead.
func (*MetricMetadataEntry) Descriptor() ([]byte, []int) {
	return file_metadata_metadatapb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *MetricMetadataEntry) GetMetas() []*Meta {
	if x != nil {
		return x.Metas
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Help string `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
	Unit string `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadatapb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadatapb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_metadata_metadatapb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *Meta) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Meta) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *Meta) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

var File_metadata_metadatapb_rpc_proto protoreflect.FileDescriptor

var file_metadata_metadatapb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x1a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x5b, 0x0a, 0x19, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x17,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x74, 0x0a, 0x16, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xac, 0x01,
	0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x58, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x39, 0x0a, 0x13,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x22, 0x42, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x32, 0x5d, 0x0a, 0x08, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x51, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x74, 0x68, 0x61, 0x6e,
	0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2d,
	0x69, 0x6f, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_metadatapb_rpc_proto_rawDescOnce sync.Once
	file_metadata_metadatapb_rpc_proto_rawDescData = file_metadata_metadatapb_rpc_proto_rawDesc
)

func file_metadata_metadatapb_rpc_proto_rawDescGZIP() []byte {
	file_metadata_metadatapb_rpc_proto_rawDescOnce.Do(func() {
		file_metadata_metadatapb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_metadatapb_rpc_proto_rawDescData)
	})
	return file_metadata_metadatapb_rpc_proto_rawDescData
}

var file_metadata_metadatapb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metadata_metadatapb_rpc_proto_goTypes = []interface{}{
	(*MetricMetadataRequest)(nil),        // 0: thanos.MetricMetadataRequest
	(*MetricMetadataResponse)(nil),       // 1: thanos.MetricMetadataResponse
	(*MetricMetadata)(nil),               // 2: thanos.MetricMetadata
	(*MetricMetadataEntry)(nil),          // 3: thanos.MetricMetadataEntry
	(*Meta)(nil),                         // 4: thanos.Meta
	nil,                                  // 5: thanos.MetricMetadata.MetadataEntry
	(storepb.PartialResponseStrategy)(0), // 6: thanos.PartialResponseStrategy
}
var file_metadata_metadatapb_rpc_proto_depIdxs = []int32{
	6, // 0: thanos.MetricMetadataRequest.partial_response_strategy:type_name -> thanos.PartialResponseStrategy
	2, // 1: thanos.MetricMetadataResponse.metadata:type_name -> thanos.MetricMetadata
	5, // 2: thanos.MetricMetadata.metadata:type_name -> thanos.MetricMetadata.MetadataEntry
	4, // 3: thanos.MetricMetadataEntry.metas:type_name -> thanos.Meta
	3, // 4: thanos.MetricMetadata.MetadataEntry.value:type_name -> thanos.MetricMetadataEntry
	0, // 5: thanos.Metadata.MetricMetadata:input_type -> thanos.MetricMetadataRequest
	1, // 6: thanos.Metadata.MetricMetadata:output_type -> thanos.MetricMetadataResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_metadata_metadatapb_rpc_proto_init() }
func file_metadata_metadatapb_rpc_proto_init() {
	if File_metadata_metadatapb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_metadatapb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricMetadataRequest); i {
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
		file_metadata_metadatapb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricMetadataResponse); i {
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
		file_metadata_metadatapb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricMetadata); i {
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
		file_metadata_metadatapb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricMetadataEntry); i {
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
		file_metadata_metadatapb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
	file_metadata_metadatapb_rpc_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MetricMetadataResponse_Metadata)(nil),
		(*MetricMetadataResponse_Warning)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metadata_metadatapb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metadata_metadatapb_rpc_proto_goTypes,
		DependencyIndexes: file_metadata_metadatapb_rpc_proto_depIdxs,
		MessageInfos:      file_metadata_metadatapb_rpc_proto_msgTypes,
	}.Build()
	File_metadata_metadatapb_rpc_proto = out.File
	file_metadata_metadatapb_rpc_proto_rawDesc = nil
	file_metadata_metadatapb_rpc_proto_goTypes = nil
	file_metadata_metadatapb_rpc_proto_depIdxs = nil
}
