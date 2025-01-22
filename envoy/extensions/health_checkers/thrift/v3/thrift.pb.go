// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.26.1
// source: envoy/extensions/health_checkers/thrift/v3/thrift.proto

package thriftv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/sefaphlvn/versioned-go-control-plane/envoy/extensions/filters/network/thrift_proxy/v3"
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

type Thrift struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the method name that will be set on each health check request dispatched to an upstream host.
	// Note that method name is case sensitive.
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// Configures the transport type to be used with the health checks. Note that
	// :ref:`AUTO_TRANSPORT<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.TransportType.AUTO_TRANSPORT>`
	// is not supported, and we don't honor :ref:`ThriftProtocolOptions<envoy_v3_api_msg_extensions.filters.network.thrift_proxy.v3.ThriftProtocolOptions>`
	// since it's possible to set to :ref:`AUTO_TRANSPORT<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.TransportType.AUTO_TRANSPORT>`.
	// [#extension-category: envoy.filters.network]
	Transport v3.TransportType `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.TransportType" json:"transport,omitempty"`
	// Configures the protocol type to be used with the health checks. Note that
	// :ref:`AUTO_PROTOCOL<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.AUTO_PROTOCOL>`
	// and :ref:`TWITTER<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.TWITTER>`
	// are not supported, and we don't honor :ref:`ThriftProtocolOptions<envoy_v3_api_msg_extensions.filters.network.thrift_proxy.v3.ThriftProtocolOptions>`
	// since it's possible to set to :ref:`AUTO_PROTOCOL<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.AUTO_PROTOCOL>`
	// or :ref:`TWITTER<envoy_v3_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.TWITTER>`.
	Protocol v3.ProtocolType `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType" json:"protocol,omitempty"`
}

func (x *Thrift) Reset() {
	*x = Thrift{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thrift) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thrift) ProtoMessage() {}

func (x *Thrift) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thrift.ProtoReflect.Descriptor instead.
func (*Thrift) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescGZIP(), []int{0}
}

func (x *Thrift) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *Thrift) GetTransport() v3.TransportType {
	if x != nil {
		return x.Transport
	}
	return v3.TransportType(0)
}

func (x *Thrift) GetProtocol() v3.ProtocolType {
	if x != nil {
		return x.Protocol
	}
	return v3.ProtocolType(0)
}

var File_envoy_extensions_health_checkers_thrift_v3_thrift_proto protoreflect.FileDescriptor

var file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDesc = []byte{
	0x0a, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x73, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x68, 0x72,
	0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x68, 0x72, 0x69,
	0x66, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x06, 0x54, 0x68, 0x72, 0x69, 0x66, 0x74, 0x12, 0x28, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68,
	0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x64, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0xad, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x73, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x54, 0x68, 0x72,
	0x69, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x73, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x68,
	0x72, 0x69, 0x66, 0x74, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescOnce sync.Once
	file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescData = file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDesc
)

func file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescGZIP() []byte {
	file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescData)
	})
	return file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDescData
}

var file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_goTypes = []interface{}{
	(*Thrift)(nil),        // 0: envoy.extensions.health_checkers.thrift.v3.Thrift
	(v3.TransportType)(0), // 1: envoy.extensions.filters.network.thrift_proxy.v3.TransportType
	(v3.ProtocolType)(0),  // 2: envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType
}
var file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.health_checkers.thrift.v3.Thrift.transport:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.TransportType
	2, // 1: envoy.extensions.health_checkers.thrift.v3.Thrift.protocol:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_init() }
func file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_init() {
	if File_envoy_extensions_health_checkers_thrift_v3_thrift_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thrift); i {
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
			RawDescriptor: file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_msgTypes,
	}.Build()
	File_envoy_extensions_health_checkers_thrift_v3_thrift_proto = out.File
	file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_rawDesc = nil
	file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_goTypes = nil
	file_envoy_extensions_health_checkers_thrift_v3_thrift_proto_depIdxs = nil
}
