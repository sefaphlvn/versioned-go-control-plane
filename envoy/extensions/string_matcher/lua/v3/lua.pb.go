// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.26.1
// source: envoy/extensions/string_matcher/lua/v3/lua.proto

package luav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/sefaphlvn/versioned-go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Lua struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Lua code that Envoy will execute
	SourceCode *v3.DataSource `protobuf:"bytes,1,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
}

func (x *Lua) Reset() {
	*x = Lua{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_string_matcher_lua_v3_lua_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lua) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lua) ProtoMessage() {}

func (x *Lua) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_string_matcher_lua_v3_lua_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lua.ProtoReflect.Descriptor instead.
func (*Lua) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescGZIP(), []int{0}
}

func (x *Lua) GetSourceCode() *v3.DataSource {
	if x != nil {
		return x.SourceCode
	}
	return nil
}

var File_envoy_extensions_string_matcher_lua_v3_lua_proto protoreflect.FileDescriptor

var file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x6c, 0x75, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x75, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x6c, 0x75, 0x61, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x03, 0x4c, 0x75, 0x61, 0x12, 0x4b, 0x0a, 0x0b, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x9f, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x0a, 0x34, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x6c, 0x75, 0x61, 0x2e, 0x76, 0x33, 0x42, 0x08, 0x4c, 0x75, 0x61, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6c, 0x75, 0x61,
	0x2f, 0x76, 0x33, 0x3b, 0x6c, 0x75, 0x61, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescOnce sync.Once
	file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescData = file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDesc
)

func file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescGZIP() []byte {
	file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescData)
	})
	return file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDescData
}

var file_envoy_extensions_string_matcher_lua_v3_lua_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_string_matcher_lua_v3_lua_proto_goTypes = []interface{}{
	(*Lua)(nil),           // 0: envoy.extensions.string_matcher.lua.v3.Lua
	(*v3.DataSource)(nil), // 1: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_string_matcher_lua_v3_lua_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.string_matcher.lua.v3.Lua.source_code:type_name -> envoy.config.core.v3.DataSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_string_matcher_lua_v3_lua_proto_init() }
func file_envoy_extensions_string_matcher_lua_v3_lua_proto_init() {
	if File_envoy_extensions_string_matcher_lua_v3_lua_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_string_matcher_lua_v3_lua_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lua); i {
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
			RawDescriptor: file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_string_matcher_lua_v3_lua_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_string_matcher_lua_v3_lua_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_string_matcher_lua_v3_lua_proto_msgTypes,
	}.Build()
	File_envoy_extensions_string_matcher_lua_v3_lua_proto = out.File
	file_envoy_extensions_string_matcher_lua_v3_lua_proto_rawDesc = nil
	file_envoy_extensions_string_matcher_lua_v3_lua_proto_goTypes = nil
	file_envoy_extensions_string_matcher_lua_v3_lua_proto_depIdxs = nil
}
