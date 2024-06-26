// Copyright 2017 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.20.1
// source: client/configpb/multilog.proto

package configpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TemporalLogConfig is a set of LogShardConfig messages, whose
// time limits should be contiguous.
type TemporalLogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shard []*LogShardConfig `protobuf:"bytes,1,rep,name=shard,proto3" json:"shard,omitempty"`
}

func (x *TemporalLogConfig) Reset() {
	*x = TemporalLogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_configpb_multilog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporalLogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporalLogConfig) ProtoMessage() {}

func (x *TemporalLogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_client_configpb_multilog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporalLogConfig.ProtoReflect.Descriptor instead.
func (*TemporalLogConfig) Descriptor() ([]byte, []int) {
	return file_client_configpb_multilog_proto_rawDescGZIP(), []int{0}
}

func (x *TemporalLogConfig) GetShard() []*LogShardConfig {
	if x != nil {
		return x.Shard
	}
	return nil
}

// LogShardConfig describes the acceptable date range for a single shard of a temporal
// log.
type LogShardConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The log's public key in DER-encoded PKIX form.
	PublicKeyDer []byte `protobuf:"bytes,2,opt,name=public_key_der,json=publicKeyDer,proto3" json:"public_key_der,omitempty"`
	// not_after_start defines the start of the range of acceptable NotAfter
	// values, inclusive.
	// Leaving this unset implies no lower bound to the range.
	NotAfterStart *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=not_after_start,json=notAfterStart,proto3" json:"not_after_start,omitempty"`
	// not_after_limit defines the end of the range of acceptable NotAfter values,
	// exclusive.
	// Leaving this unset implies no upper bound to the range.
	NotAfterLimit *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=not_after_limit,json=notAfterLimit,proto3" json:"not_after_limit,omitempty"`
}

func (x *LogShardConfig) Reset() {
	*x = LogShardConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_configpb_multilog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogShardConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogShardConfig) ProtoMessage() {}

func (x *LogShardConfig) ProtoReflect() protoreflect.Message {
	mi := &file_client_configpb_multilog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogShardConfig.ProtoReflect.Descriptor instead.
func (*LogShardConfig) Descriptor() ([]byte, []int) {
	return file_client_configpb_multilog_proto_rawDescGZIP(), []int{1}
}

func (x *LogShardConfig) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *LogShardConfig) GetPublicKeyDer() []byte {
	if x != nil {
		return x.PublicKeyDer
	}
	return nil
}

func (x *LogShardConfig) GetNotAfterStart() *timestamppb.Timestamp {
	if x != nil {
		return x.NotAfterStart
	}
	return nil
}

func (x *LogShardConfig) GetNotAfterLimit() *timestamppb.Timestamp {
	if x != nil {
		return x.NotAfterLimit
	}
	return nil
}

var File_client_configpb_multilog_proto protoreflect.FileDescriptor

var file_client_configpb_multilog_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70,
	0x62, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x11, 0x54,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2e, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x53, 0x68, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0f, 0x6e,
	0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x42, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_configpb_multilog_proto_rawDescOnce sync.Once
	file_client_configpb_multilog_proto_rawDescData = file_client_configpb_multilog_proto_rawDesc
)

func file_client_configpb_multilog_proto_rawDescGZIP() []byte {
	file_client_configpb_multilog_proto_rawDescOnce.Do(func() {
		file_client_configpb_multilog_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_configpb_multilog_proto_rawDescData)
	})
	return file_client_configpb_multilog_proto_rawDescData
}

var file_client_configpb_multilog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_configpb_multilog_proto_goTypes = []interface{}{
	(*TemporalLogConfig)(nil),     // 0: configpb.TemporalLogConfig
	(*LogShardConfig)(nil),        // 1: configpb.LogShardConfig
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_client_configpb_multilog_proto_depIdxs = []int32{
	1, // 0: configpb.TemporalLogConfig.shard:type_name -> configpb.LogShardConfig
	2, // 1: configpb.LogShardConfig.not_after_start:type_name -> google.protobuf.Timestamp
	2, // 2: configpb.LogShardConfig.not_after_limit:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_configpb_multilog_proto_init() }
func file_client_configpb_multilog_proto_init() {
	if File_client_configpb_multilog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_configpb_multilog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemporalLogConfig); i {
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
		file_client_configpb_multilog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogShardConfig); i {
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
			RawDescriptor: file_client_configpb_multilog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_configpb_multilog_proto_goTypes,
		DependencyIndexes: file_client_configpb_multilog_proto_depIdxs,
		MessageInfos:      file_client_configpb_multilog_proto_msgTypes,
	}.Build()
	File_client_configpb_multilog_proto = out.File
	file_client_configpb_multilog_proto_rawDesc = nil
	file_client_configpb_multilog_proto_goTypes = nil
	file_client_configpb_multilog_proto_depIdxs = nil
}
