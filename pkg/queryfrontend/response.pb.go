// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: queryfrontend/response.proto

package queryfrontend

import (
	labelpb "github.com/thanos-io/thanos/pkg/store/labelpb"
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

type ThanosLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string            `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data        []string          `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
	ErrorType   string            `protobuf:"bytes,3,opt,name=ErrorType,proto3" json:"ErrorType,omitempty"`
	Error       string            `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
	RespHeaders []*ResponseHeader `protobuf:"bytes,5,rep,name=RespHeaders,proto3" json:"RespHeaders,omitempty"`
}

func (x *ThanosLabelsResponse) Reset() {
	*x = ThanosLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queryfrontend_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThanosLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThanosLabelsResponse) ProtoMessage() {}

func (x *ThanosLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queryfrontend_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThanosLabelsResponse.ProtoReflect.Descriptor instead.
func (*ThanosLabelsResponse) Descriptor() ([]byte, []int) {
	return file_queryfrontend_response_proto_rawDescGZIP(), []int{0}
}

func (x *ThanosLabelsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ThanosLabelsResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ThanosLabelsResponse) GetErrorType() string {
	if x != nil {
		return x.ErrorType
	}
	return ""
}

func (x *ThanosLabelsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ThanosLabelsResponse) GetRespHeaders() []*ResponseHeader {
	if x != nil {
		return x.RespHeaders
	}
	return nil
}

type ThanosSeriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string               `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data        []*labelpb.ZLabelSet `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
	ErrorType   string               `protobuf:"bytes,3,opt,name=ErrorType,proto3" json:"ErrorType,omitempty"`
	Error       string               `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
	RespHeaders []*ResponseHeader    `protobuf:"bytes,5,rep,name=RespHeaders,proto3" json:"RespHeaders,omitempty"`
}

func (x *ThanosSeriesResponse) Reset() {
	*x = ThanosSeriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queryfrontend_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThanosSeriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThanosSeriesResponse) ProtoMessage() {}

func (x *ThanosSeriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queryfrontend_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThanosSeriesResponse.ProtoReflect.Descriptor instead.
func (*ThanosSeriesResponse) Descriptor() ([]byte, []int) {
	return file_queryfrontend_response_proto_rawDescGZIP(), []int{1}
}

func (x *ThanosSeriesResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ThanosSeriesResponse) GetData() []*labelpb.ZLabelSet {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ThanosSeriesResponse) GetErrorType() string {
	if x != nil {
		return x.ErrorType
	}
	return ""
}

func (x *ThanosSeriesResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ThanosSeriesResponse) GetRespHeaders() []*ResponseHeader {
	if x != nil {
		return x.RespHeaders
	}
	return nil
}

type ResponseHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=Values,proto3" json:"Values,omitempty"`
}

func (x *ResponseHeader) Reset() {
	*x = ResponseHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queryfrontend_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHeader) ProtoMessage() {}

func (x *ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_queryfrontend_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseHeader.ProtoReflect.Descriptor instead.
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return file_queryfrontend_response_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseHeader) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_queryfrontend_response_proto protoreflect.FileDescriptor

var file_queryfrontend_response_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x1a, 0x19, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x14, 0x54, 0x68, 0x61,
	0x6e, 0x6f, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x3f, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x14, 0x54, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2e, 0x5a, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x53, 0x65, 0x74, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3f,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6e,
	0x6f, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queryfrontend_response_proto_rawDescOnce sync.Once
	file_queryfrontend_response_proto_rawDescData = file_queryfrontend_response_proto_rawDesc
)

func file_queryfrontend_response_proto_rawDescGZIP() []byte {
	file_queryfrontend_response_proto_rawDescOnce.Do(func() {
		file_queryfrontend_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_queryfrontend_response_proto_rawDescData)
	})
	return file_queryfrontend_response_proto_rawDescData
}

var file_queryfrontend_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_queryfrontend_response_proto_goTypes = []interface{}{
	(*ThanosLabelsResponse)(nil), // 0: queryfrontend.ThanosLabelsResponse
	(*ThanosSeriesResponse)(nil), // 1: queryfrontend.ThanosSeriesResponse
	(*ResponseHeader)(nil),       // 2: queryfrontend.ResponseHeader
	(*labelpb.ZLabelSet)(nil),    // 3: thanos.ZLabelSet
}
var file_queryfrontend_response_proto_depIdxs = []int32{
	2, // 0: queryfrontend.ThanosLabelsResponse.RespHeaders:type_name -> queryfrontend.ResponseHeader
	3, // 1: queryfrontend.ThanosSeriesResponse.Data:type_name -> thanos.ZLabelSet
	2, // 2: queryfrontend.ThanosSeriesResponse.RespHeaders:type_name -> queryfrontend.ResponseHeader
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_queryfrontend_response_proto_init() }
func file_queryfrontend_response_proto_init() {
	if File_queryfrontend_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queryfrontend_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThanosLabelsResponse); i {
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
		file_queryfrontend_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThanosSeriesResponse); i {
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
		file_queryfrontend_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseHeader); i {
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
			RawDescriptor: file_queryfrontend_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queryfrontend_response_proto_goTypes,
		DependencyIndexes: file_queryfrontend_response_proto_depIdxs,
		MessageInfos:      file_queryfrontend_response_proto_msgTypes,
	}.Build()
	File_queryfrontend_response_proto = out.File
	file_queryfrontend_response_proto_rawDesc = nil
	file_queryfrontend_response_proto_goTypes = nil
	file_queryfrontend_response_proto_depIdxs = nil
}
