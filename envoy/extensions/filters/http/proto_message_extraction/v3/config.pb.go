// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.2
// source: envoy/extensions/filters/http/proto_message_extraction/v3/config.proto

package proto_message_extractionv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/sefaphlvn/versioned-go-control-plane/envoy/config/core/v3"
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

type ProtoMessageExtractionConfig_ExtractMode int32

const (
	ProtoMessageExtractionConfig_ExtractMode_UNSPECIFIED ProtoMessageExtractionConfig_ExtractMode = 0
	// The filter will extract the first and the last message for
	// for streaming cases, containing
	// client-side streaming, server-side streaming or bi-directional streaming.
	ProtoMessageExtractionConfig_FIRST_AND_LAST ProtoMessageExtractionConfig_ExtractMode = 1
)

// Enum value maps for ProtoMessageExtractionConfig_ExtractMode.
var (
	ProtoMessageExtractionConfig_ExtractMode_name = map[int32]string{
		0: "ExtractMode_UNSPECIFIED",
		1: "FIRST_AND_LAST",
	}
	ProtoMessageExtractionConfig_ExtractMode_value = map[string]int32{
		"ExtractMode_UNSPECIFIED": 0,
		"FIRST_AND_LAST":          1,
	}
)

func (x ProtoMessageExtractionConfig_ExtractMode) Enum() *ProtoMessageExtractionConfig_ExtractMode {
	p := new(ProtoMessageExtractionConfig_ExtractMode)
	*p = x
	return p
}

func (x ProtoMessageExtractionConfig_ExtractMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoMessageExtractionConfig_ExtractMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_enumTypes[0].Descriptor()
}

func (ProtoMessageExtractionConfig_ExtractMode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_enumTypes[0]
}

func (x ProtoMessageExtractionConfig_ExtractMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtoMessageExtractionConfig_ExtractMode.Descriptor instead.
func (ProtoMessageExtractionConfig_ExtractMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescGZIP(), []int{0, 0}
}

type MethodExtraction_ExtractDirective int32

const (
	MethodExtraction_ExtractDirective_UNSPECIFIED MethodExtraction_ExtractDirective = 0
	// The value of this field will be extracted.
	MethodExtraction_EXTRACT MethodExtraction_ExtractDirective = 1
	// It should be only annotated on Message type fields so if the field isn't
	// empty, an empty Struct will be extracted.
	MethodExtraction_EXTRACT_REDACT MethodExtraction_ExtractDirective = 2
)

// Enum value maps for MethodExtraction_ExtractDirective.
var (
	MethodExtraction_ExtractDirective_name = map[int32]string{
		0: "ExtractDirective_UNSPECIFIED",
		1: "EXTRACT",
		2: "EXTRACT_REDACT",
	}
	MethodExtraction_ExtractDirective_value = map[string]int32{
		"ExtractDirective_UNSPECIFIED": 0,
		"EXTRACT":                      1,
		"EXTRACT_REDACT":               2,
	}
)

func (x MethodExtraction_ExtractDirective) Enum() *MethodExtraction_ExtractDirective {
	p := new(MethodExtraction_ExtractDirective)
	*p = x
	return p
}

func (x MethodExtraction_ExtractDirective) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodExtraction_ExtractDirective) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_enumTypes[1].Descriptor()
}

func (MethodExtraction_ExtractDirective) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_enumTypes[1]
}

func (x MethodExtraction_ExtractDirective) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodExtraction_ExtractDirective.Descriptor instead.
func (MethodExtraction_ExtractDirective) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescGZIP(), []int{1, 0}
}

type ProtoMessageExtractionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The proto descriptor set binary for the gRPC services.
	//
	// Types that are assignable to DescriptorSet:
	//
	//	*ProtoMessageExtractionConfig_DataSource
	//	*ProtoMessageExtractionConfig_ProtoDescriptorTypedMetadata
	DescriptorSet isProtoMessageExtractionConfig_DescriptorSet `protobuf_oneof:"descriptor_set"`
	Mode          ProtoMessageExtractionConfig_ExtractMode     `protobuf:"varint,3,opt,name=mode,proto3,enum=envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig_ExtractMode" json:"mode,omitempty"`
	// Specify the message extraction info.
	// The key is the fully qualified gRPC method name.
	// “${package}.${Service}.${Method}“, like
	// “endpoints.examples.bookstore.BookStore.GetShelf“
	//
	// The value is the message extraction information for individual gRPC
	// methods.
	ExtractionByMethod map[string]*MethodExtraction `protobuf:"bytes,4,rep,name=extraction_by_method,json=extractionByMethod,proto3" json:"extraction_by_method,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProtoMessageExtractionConfig) Reset() {
	*x = ProtoMessageExtractionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoMessageExtractionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMessageExtractionConfig) ProtoMessage() {}

func (x *ProtoMessageExtractionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMessageExtractionConfig.ProtoReflect.Descriptor instead.
func (*ProtoMessageExtractionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescGZIP(), []int{0}
}

func (m *ProtoMessageExtractionConfig) GetDescriptorSet() isProtoMessageExtractionConfig_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (x *ProtoMessageExtractionConfig) GetDataSource() *v3.DataSource {
	if x, ok := x.GetDescriptorSet().(*ProtoMessageExtractionConfig_DataSource); ok {
		return x.DataSource
	}
	return nil
}

func (x *ProtoMessageExtractionConfig) GetProtoDescriptorTypedMetadata() string {
	if x, ok := x.GetDescriptorSet().(*ProtoMessageExtractionConfig_ProtoDescriptorTypedMetadata); ok {
		return x.ProtoDescriptorTypedMetadata
	}
	return ""
}

func (x *ProtoMessageExtractionConfig) GetMode() ProtoMessageExtractionConfig_ExtractMode {
	if x != nil {
		return x.Mode
	}
	return ProtoMessageExtractionConfig_ExtractMode_UNSPECIFIED
}

func (x *ProtoMessageExtractionConfig) GetExtractionByMethod() map[string]*MethodExtraction {
	if x != nil {
		return x.ExtractionByMethod
	}
	return nil
}

type isProtoMessageExtractionConfig_DescriptorSet interface {
	isProtoMessageExtractionConfig_DescriptorSet()
}

type ProtoMessageExtractionConfig_DataSource struct {
	// It could be passed by a local file through “Datasource.filename“ or
	// embedded in the “Datasource.inline_bytes“.
	DataSource *v3.DataSource `protobuf:"bytes,1,opt,name=data_source,json=dataSource,proto3,oneof"`
}

type ProtoMessageExtractionConfig_ProtoDescriptorTypedMetadata struct {
	// Unimplemented, the key of proto descriptor TypedMetadata.
	// Among filters depending on the proto descriptor, we can have a
	// TypedMetadata for proto descriptors, so that these filters can share one
	// copy of proto descriptor in memory.
	ProtoDescriptorTypedMetadata string `protobuf:"bytes,2,opt,name=proto_descriptor_typed_metadata,json=protoDescriptorTypedMetadata,proto3,oneof"`
}

func (*ProtoMessageExtractionConfig_DataSource) isProtoMessageExtractionConfig_DescriptorSet() {}

func (*ProtoMessageExtractionConfig_ProtoDescriptorTypedMetadata) isProtoMessageExtractionConfig_DescriptorSet() {
}

// This message can be used to support per route config approach later even
// though the Istio doesn't support that so far.
type MethodExtraction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mapping of field path to its ExtractDirective for request messages
	RequestExtractionByField map[string]MethodExtraction_ExtractDirective `protobuf:"bytes,2,rep,name=request_extraction_by_field,json=requestExtractionByField,proto3" json:"request_extraction_by_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction_ExtractDirective"`
	// The mapping of field path to its ExtractDirective for response messages
	ResponseExtractionByField map[string]MethodExtraction_ExtractDirective `protobuf:"bytes,3,rep,name=response_extraction_by_field,json=responseExtractionByField,proto3" json:"response_extraction_by_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction_ExtractDirective"`
}

func (x *MethodExtraction) Reset() {
	*x = MethodExtraction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodExtraction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodExtraction) ProtoMessage() {}

func (x *MethodExtraction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodExtraction.ProtoReflect.Descriptor instead.
func (*MethodExtraction) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescGZIP(), []int{1}
}

func (x *MethodExtraction) GetRequestExtractionByField() map[string]MethodExtraction_ExtractDirective {
	if x != nil {
		return x.RequestExtractionByField
	}
	return nil
}

func (x *MethodExtraction) GetResponseExtractionByField() map[string]MethodExtraction_ExtractDirective {
	if x != nil {
		return x.ResponseExtractionByField
	}
	return nil
}

var File_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDesc = []byte{
	0x0a, 0x46, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x05, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x1f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x77, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0xa1, 0x01, 0x0a,
	0x14, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x1a, 0x92, 0x01, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x61,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4c,
	0x41, 0x53, 0x54, 0x10, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x22, 0x9b, 0x06, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa8, 0x01, 0x0a,
	0x1b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0xab, 0x01, 0x0a, 0x1c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x6a,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x19, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0xa9, 0x01, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x72, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0xaa, 0x01, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x72, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x55,
	0x0a, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x54, 0x52, 0x41, 0x43, 0x54, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x58, 0x54, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x44,
	0x41, 0x43, 0x54, 0x10, 0x02, 0x42, 0xe5, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02,
	0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x47, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x33, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x7b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x33, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescData = file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDesc
)

func file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDescData
}

var file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_goTypes = []interface{}{
	(ProtoMessageExtractionConfig_ExtractMode)(0), // 0: envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.ExtractMode
	(MethodExtraction_ExtractDirective)(0),        // 1: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.ExtractDirective
	(*ProtoMessageExtractionConfig)(nil),          // 2: envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig
	(*MethodExtraction)(nil),                      // 3: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction
	nil,                                           // 4: envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.ExtractionByMethodEntry
	nil,                                           // 5: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.RequestExtractionByFieldEntry
	nil,                                           // 6: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.ResponseExtractionByFieldEntry
	(*v3.DataSource)(nil),                         // 7: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_depIdxs = []int32{
	7, // 0: envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.data_source:type_name -> envoy.config.core.v3.DataSource
	0, // 1: envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.mode:type_name -> envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.ExtractMode
	4, // 2: envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.extraction_by_method:type_name -> envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.ExtractionByMethodEntry
	5, // 3: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.request_extraction_by_field:type_name -> envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.RequestExtractionByFieldEntry
	6, // 4: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.response_extraction_by_field:type_name -> envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.ResponseExtractionByFieldEntry
	3, // 5: envoy.extensions.filters.http.proto_message_extraction.v3.ProtoMessageExtractionConfig.ExtractionByMethodEntry.value:type_name -> envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction
	1, // 6: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.RequestExtractionByFieldEntry.value:type_name -> envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.ExtractDirective
	1, // 7: envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.ResponseExtractionByFieldEntry.value:type_name -> envoy.extensions.filters.http.proto_message_extraction.v3.MethodExtraction.ExtractDirective
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_init() }
func file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_init() {
	if File_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoMessageExtractionConfig); i {
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
		file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodExtraction); i {
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
	file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProtoMessageExtractionConfig_DataSource)(nil),
		(*ProtoMessageExtractionConfig_ProtoDescriptorTypedMetadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto = out.File
	file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_rawDesc = nil
	file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_goTypes = nil
	file_envoy_extensions_filters_http_proto_message_extraction_v3_config_proto_depIdxs = nil
}
