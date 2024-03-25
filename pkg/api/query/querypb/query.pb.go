// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: api/query/querypb/query.proto

package querypb

import (
	storepb "github.com/oodle-ai/thanos/pkg/store/storepb"
	prompb "github.com/oodle-ai/thanos/pkg/store/storepb/prompb"
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

type EngineType int32

const (
	EngineType_default    EngineType = 0
	EngineType_prometheus EngineType = 1
	EngineType_thanos     EngineType = 2
)

// Enum value maps for EngineType.
var (
	EngineType_name = map[int32]string{
		0: "default",
		1: "prometheus",
		2: "thanos",
	}
	EngineType_value = map[string]int32{
		"default":    0,
		"prometheus": 1,
		"thanos":     2,
	}
)

func (x EngineType) Enum() *EngineType {
	p := new(EngineType)
	*p = x
	return p
}

func (x EngineType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EngineType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_query_querypb_query_proto_enumTypes[0].Descriptor()
}

func (EngineType) Type() protoreflect.EnumType {
	return &file_api_query_querypb_query_proto_enumTypes[0]
}

func (x EngineType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EngineType.Descriptor instead.
func (EngineType) EnumDescriptor() ([]byte, []int) {
	return file_api_query_querypb_query_proto_rawDescGZIP(), []int{0}
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query                 string             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	TimeSeconds           int64              `protobuf:"varint,2,opt,name=time_seconds,json=timeSeconds,proto3" json:"time_seconds,omitempty"`
	TimeoutSeconds        int64              `protobuf:"varint,3,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	MaxResolutionSeconds  int64              `protobuf:"varint,4,opt,name=max_resolution_seconds,json=maxResolutionSeconds,proto3" json:"max_resolution_seconds,omitempty"`
	ReplicaLabels         []string           `protobuf:"bytes,5,rep,name=replica_labels,json=replicaLabels,proto3" json:"replica_labels,omitempty"`
	StoreMatchers         []*StoreMatchers   `protobuf:"bytes,6,rep,name=storeMatchers,proto3" json:"storeMatchers,omitempty"`
	EnableDedup           bool               `protobuf:"varint,7,opt,name=enableDedup,proto3" json:"enableDedup,omitempty"`
	EnablePartialResponse bool               `protobuf:"varint,8,opt,name=enablePartialResponse,proto3" json:"enablePartialResponse,omitempty"`
	SkipChunks            bool               `protobuf:"varint,10,opt,name=skipChunks,proto3" json:"skipChunks,omitempty"`
	ShardInfo             *storepb.ShardInfo `protobuf:"bytes,11,opt,name=shard_info,json=shardInfo,proto3" json:"shard_info,omitempty"`
	LookbackDeltaSeconds  int64              `protobuf:"varint,12,opt,name=lookback_delta_seconds,json=lookbackDeltaSeconds,proto3" json:"lookback_delta_seconds,omitempty"`
	Engine                EngineType         `protobuf:"varint,13,opt,name=engine,proto3,enum=thanos.EngineType" json:"engine,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_query_querypb_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_query_querypb_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_api_query_querypb_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryRequest) GetTimeSeconds() int64 {
	if x != nil {
		return x.TimeSeconds
	}
	return 0
}

func (x *QueryRequest) GetTimeoutSeconds() int64 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

func (x *QueryRequest) GetMaxResolutionSeconds() int64 {
	if x != nil {
		return x.MaxResolutionSeconds
	}
	return 0
}

func (x *QueryRequest) GetReplicaLabels() []string {
	if x != nil {
		return x.ReplicaLabels
	}
	return nil
}

func (x *QueryRequest) GetStoreMatchers() []*StoreMatchers {
	if x != nil {
		return x.StoreMatchers
	}
	return nil
}

func (x *QueryRequest) GetEnableDedup() bool {
	if x != nil {
		return x.EnableDedup
	}
	return false
}

func (x *QueryRequest) GetEnablePartialResponse() bool {
	if x != nil {
		return x.EnablePartialResponse
	}
	return false
}

func (x *QueryRequest) GetSkipChunks() bool {
	if x != nil {
		return x.SkipChunks
	}
	return false
}

func (x *QueryRequest) GetShardInfo() *storepb.ShardInfo {
	if x != nil {
		return x.ShardInfo
	}
	return nil
}

func (x *QueryRequest) GetLookbackDeltaSeconds() int64 {
	if x != nil {
		return x.LookbackDeltaSeconds
	}
	return 0
}

func (x *QueryRequest) GetEngine() EngineType {
	if x != nil {
		return x.Engine
	}
	return EngineType_default
}

type StoreMatchers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelMatchers []*storepb.LabelMatcher `protobuf:"bytes,1,rep,name=labelMatchers,proto3" json:"labelMatchers,omitempty"`
}

func (x *StoreMatchers) Reset() {
	*x = StoreMatchers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_query_querypb_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreMatchers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreMatchers) ProtoMessage() {}

func (x *StoreMatchers) ProtoReflect() protoreflect.Message {
	mi := &file_api_query_querypb_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreMatchers.ProtoReflect.Descriptor instead.
func (*StoreMatchers) Descriptor() ([]byte, []int) {
	return file_api_query_querypb_query_proto_rawDescGZIP(), []int{1}
}

func (x *StoreMatchers) GetLabelMatchers() []*storepb.LabelMatcher {
	if x != nil {
		return x.LabelMatchers
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*QueryResponse_Warnings
	//	*QueryResponse_Timeseries
	Result isQueryResponse_Result `protobuf_oneof:"result"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_query_querypb_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_query_querypb_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_api_query_querypb_query_proto_rawDescGZIP(), []int{2}
}

func (m *QueryResponse) GetResult() isQueryResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *QueryResponse) GetWarnings() string {
	if x, ok := x.GetResult().(*QueryResponse_Warnings); ok {
		return x.Warnings
	}
	return ""
}

func (x *QueryResponse) GetTimeseries() *prompb.TimeSeries {
	if x, ok := x.GetResult().(*QueryResponse_Timeseries); ok {
		return x.Timeseries
	}
	return nil
}

type isQueryResponse_Result interface {
	isQueryResponse_Result()
}

type QueryResponse_Warnings struct {
	// / warnings are additional messages coming from the PromQL engine.
	Warnings string `protobuf:"bytes,1,opt,name=warnings,proto3,oneof"`
}

type QueryResponse_Timeseries struct {
	// / timeseries is one series from the result of the executed query.
	Timeseries *prompb.TimeSeries `protobuf:"bytes,2,opt,name=timeseries,proto3,oneof"`
}

func (*QueryResponse_Warnings) isQueryResponse_Result() {}

func (*QueryResponse_Timeseries) isQueryResponse_Result() {}

type QueryRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query                 string             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	StartTimeSeconds      int64              `protobuf:"varint,2,opt,name=start_time_seconds,json=startTimeSeconds,proto3" json:"start_time_seconds,omitempty"`
	EndTimeSeconds        int64              `protobuf:"varint,3,opt,name=end_time_seconds,json=endTimeSeconds,proto3" json:"end_time_seconds,omitempty"`
	IntervalSeconds       int64              `protobuf:"varint,4,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	TimeoutSeconds        int64              `protobuf:"varint,5,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	MaxResolutionSeconds  int64              `protobuf:"varint,6,opt,name=max_resolution_seconds,json=maxResolutionSeconds,proto3" json:"max_resolution_seconds,omitempty"`
	ReplicaLabels         []string           `protobuf:"bytes,7,rep,name=replica_labels,json=replicaLabels,proto3" json:"replica_labels,omitempty"`
	StoreMatchers         []*StoreMatchers   `protobuf:"bytes,8,rep,name=storeMatchers,proto3" json:"storeMatchers,omitempty"`
	EnableDedup           bool               `protobuf:"varint,9,opt,name=enableDedup,proto3" json:"enableDedup,omitempty"`
	EnablePartialResponse bool               `protobuf:"varint,10,opt,name=enablePartialResponse,proto3" json:"enablePartialResponse,omitempty"`
	SkipChunks            bool               `protobuf:"varint,12,opt,name=skipChunks,proto3" json:"skipChunks,omitempty"`
	ShardInfo             *storepb.ShardInfo `protobuf:"bytes,13,opt,name=shard_info,json=shardInfo,proto3" json:"shard_info,omitempty"`
	LookbackDeltaSeconds  int64              `protobuf:"varint,14,opt,name=lookback_delta_seconds,json=lookbackDeltaSeconds,proto3" json:"lookback_delta_seconds,omitempty"`
	Engine                EngineType         `protobuf:"varint,15,opt,name=engine,proto3,enum=thanos.EngineType" json:"engine,omitempty"`
}

func (x *QueryRangeRequest) Reset() {
	*x = QueryRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_query_querypb_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRangeRequest) ProtoMessage() {}

func (x *QueryRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_query_querypb_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRangeRequest.ProtoReflect.Descriptor instead.
func (*QueryRangeRequest) Descriptor() ([]byte, []int) {
	return file_api_query_querypb_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRangeRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryRangeRequest) GetStartTimeSeconds() int64 {
	if x != nil {
		return x.StartTimeSeconds
	}
	return 0
}

func (x *QueryRangeRequest) GetEndTimeSeconds() int64 {
	if x != nil {
		return x.EndTimeSeconds
	}
	return 0
}

func (x *QueryRangeRequest) GetIntervalSeconds() int64 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

func (x *QueryRangeRequest) GetTimeoutSeconds() int64 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

func (x *QueryRangeRequest) GetMaxResolutionSeconds() int64 {
	if x != nil {
		return x.MaxResolutionSeconds
	}
	return 0
}

func (x *QueryRangeRequest) GetReplicaLabels() []string {
	if x != nil {
		return x.ReplicaLabels
	}
	return nil
}

func (x *QueryRangeRequest) GetStoreMatchers() []*StoreMatchers {
	if x != nil {
		return x.StoreMatchers
	}
	return nil
}

func (x *QueryRangeRequest) GetEnableDedup() bool {
	if x != nil {
		return x.EnableDedup
	}
	return false
}

func (x *QueryRangeRequest) GetEnablePartialResponse() bool {
	if x != nil {
		return x.EnablePartialResponse
	}
	return false
}

func (x *QueryRangeRequest) GetSkipChunks() bool {
	if x != nil {
		return x.SkipChunks
	}
	return false
}

func (x *QueryRangeRequest) GetShardInfo() *storepb.ShardInfo {
	if x != nil {
		return x.ShardInfo
	}
	return nil
}

func (x *QueryRangeRequest) GetLookbackDeltaSeconds() int64 {
	if x != nil {
		return x.LookbackDeltaSeconds
	}
	return 0
}

func (x *QueryRangeRequest) GetEngine() EngineType {
	if x != nil {
		return x.Engine
	}
	return EngineType_default
}

type QueryRangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*QueryRangeResponse_Warnings
	//	*QueryRangeResponse_Timeseries
	Result isQueryRangeResponse_Result `protobuf_oneof:"result"`
}

func (x *QueryRangeResponse) Reset() {
	*x = QueryRangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_query_querypb_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRangeResponse) ProtoMessage() {}

func (x *QueryRangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_query_querypb_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRangeResponse.ProtoReflect.Descriptor instead.
func (*QueryRangeResponse) Descriptor() ([]byte, []int) {
	return file_api_query_querypb_query_proto_rawDescGZIP(), []int{4}
}

func (m *QueryRangeResponse) GetResult() isQueryRangeResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *QueryRangeResponse) GetWarnings() string {
	if x, ok := x.GetResult().(*QueryRangeResponse_Warnings); ok {
		return x.Warnings
	}
	return ""
}

func (x *QueryRangeResponse) GetTimeseries() *prompb.TimeSeries {
	if x, ok := x.GetResult().(*QueryRangeResponse_Timeseries); ok {
		return x.Timeseries
	}
	return nil
}

type isQueryRangeResponse_Result interface {
	isQueryRangeResponse_Result()
}

type QueryRangeResponse_Warnings struct {
	// / warnings are additional messages coming from the PromQL engine.
	Warnings string `protobuf:"bytes,1,opt,name=warnings,proto3,oneof"`
}

type QueryRangeResponse_Timeseries struct {
	// / timeseries is one series from the result of the executed query.
	Timeseries *prompb.TimeSeries `protobuf:"bytes,2,opt,name=timeseries,proto3,oneof"`
}

func (*QueryRangeResponse_Warnings) isQueryRangeResponse_Result() {}

func (*QueryRangeResponse_Timeseries) isQueryRangeResponse_Result() {}

var File_api_query_querypb_query_proto protoreflect.FileDescriptor

var file_api_query_querypb_query_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x1a, 0x19, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04,
	0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x34, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x14, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3b, 0x0a,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x0d, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x64, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x64, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x15,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x6f, 0x6f, 0x6b, 0x62, 0x61, 0x63, 0x6b,
	0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6c, 0x6f, 0x6f, 0x6b, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x22, 0x4b, 0x0a, 0x0d,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a,
	0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0d, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x76, 0x0a, 0x0d, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x77, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x81, 0x05, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2c, 0x0a,
	0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x61, 0x78,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x64,
	0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x65, 0x64, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6b, 0x69, 0x70, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a,
	0x16, 0x6c, 0x6f, 0x6f, 0x6b, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6c,
	0x6f, 0x6f, 0x6b, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4a,
	0x04, 0x08, 0x0b, 0x10, 0x0c, 0x22, 0x7b, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x77,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2a, 0x35, 0x0a, 0x0a, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x10, 0x02, 0x32, 0x86, 0x01, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x2e, 0x74,
	0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x74, 0x68, 0x61, 0x6e,
	0x6f, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f,
	0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_query_querypb_query_proto_rawDescOnce sync.Once
	file_api_query_querypb_query_proto_rawDescData = file_api_query_querypb_query_proto_rawDesc
)

func file_api_query_querypb_query_proto_rawDescGZIP() []byte {
	file_api_query_querypb_query_proto_rawDescOnce.Do(func() {
		file_api_query_querypb_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_query_querypb_query_proto_rawDescData)
	})
	return file_api_query_querypb_query_proto_rawDescData
}

var file_api_query_querypb_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_query_querypb_query_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_query_querypb_query_proto_goTypes = []interface{}{
	(EngineType)(0),              // 0: thanos.EngineType
	(*QueryRequest)(nil),         // 1: thanos.QueryRequest
	(*StoreMatchers)(nil),        // 2: thanos.StoreMatchers
	(*QueryResponse)(nil),        // 3: thanos.QueryResponse
	(*QueryRangeRequest)(nil),    // 4: thanos.QueryRangeRequest
	(*QueryRangeResponse)(nil),   // 5: thanos.QueryRangeResponse
	(*storepb.ShardInfo)(nil),    // 6: thanos.ShardInfo
	(*storepb.LabelMatcher)(nil), // 7: thanos.LabelMatcher
	(*prompb.TimeSeries)(nil),    // 8: prometheus_copy.TimeSeries
}
var file_api_query_querypb_query_proto_depIdxs = []int32{
	2,  // 0: thanos.QueryRequest.storeMatchers:type_name -> thanos.StoreMatchers
	6,  // 1: thanos.QueryRequest.shard_info:type_name -> thanos.ShardInfo
	0,  // 2: thanos.QueryRequest.engine:type_name -> thanos.EngineType
	7,  // 3: thanos.StoreMatchers.labelMatchers:type_name -> thanos.LabelMatcher
	8,  // 4: thanos.QueryResponse.timeseries:type_name -> prometheus_copy.TimeSeries
	2,  // 5: thanos.QueryRangeRequest.storeMatchers:type_name -> thanos.StoreMatchers
	6,  // 6: thanos.QueryRangeRequest.shard_info:type_name -> thanos.ShardInfo
	0,  // 7: thanos.QueryRangeRequest.engine:type_name -> thanos.EngineType
	8,  // 8: thanos.QueryRangeResponse.timeseries:type_name -> prometheus_copy.TimeSeries
	1,  // 9: thanos.Query.Query:input_type -> thanos.QueryRequest
	4,  // 10: thanos.Query.QueryRange:input_type -> thanos.QueryRangeRequest
	3,  // 11: thanos.Query.Query:output_type -> thanos.QueryResponse
	5,  // 12: thanos.Query.QueryRange:output_type -> thanos.QueryRangeResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_query_querypb_query_proto_init() }
func file_api_query_querypb_query_proto_init() {
	if File_api_query_querypb_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_query_querypb_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_api_query_querypb_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreMatchers); i {
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
		file_api_query_querypb_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_api_query_querypb_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRangeRequest); i {
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
		file_api_query_querypb_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRangeResponse); i {
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
	file_api_query_querypb_query_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*QueryResponse_Warnings)(nil),
		(*QueryResponse_Timeseries)(nil),
	}
	file_api_query_querypb_query_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*QueryRangeResponse_Warnings)(nil),
		(*QueryRangeResponse_Timeseries)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_query_querypb_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_query_querypb_query_proto_goTypes,
		DependencyIndexes: file_api_query_querypb_query_proto_depIdxs,
		EnumInfos:         file_api_query_querypb_query_proto_enumTypes,
		MessageInfos:      file_api_query_querypb_query_proto_msgTypes,
	}.Build()
	File_api_query_querypb_query_proto = out.File
	file_api_query_querypb_query_proto_rawDesc = nil
	file_api_query_querypb_query_proto_goTypes = nil
	file_api_query_querypb_query_proto_depIdxs = nil
}
