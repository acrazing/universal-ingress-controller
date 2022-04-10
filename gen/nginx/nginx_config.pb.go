// @since 2022-04-09 19:52:51
// @author acrazing <joking.young@gmail.com>

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: nginx/nginx_config.proto

package nginxpb

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

type NginxConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NginxConfig) Reset() {
	*x = NginxConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nginx_nginx_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NginxConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NginxConfig) ProtoMessage() {}

func (x *NginxConfig) ProtoReflect() protoreflect.Message {
	mi := &file_nginx_nginx_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NginxConfig.ProtoReflect.Descriptor instead.
func (*NginxConfig) Descriptor() ([]byte, []int) {
	return file_nginx_nginx_config_proto_rawDescGZIP(), []int{0}
}

var File_nginx_nginx_config_proto protoreflect.FileDescriptor

var file_nginx_nginx_config_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0x2f, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0x22, 0x0d,
	0x0a, 0x0b, 0x4e, 0x67, 0x69, 0x6e, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x44, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x72, 0x61,
	0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x2d, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0x3b, 0x6e, 0x67, 0x69, 0x6e,
	0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nginx_nginx_config_proto_rawDescOnce sync.Once
	file_nginx_nginx_config_proto_rawDescData = file_nginx_nginx_config_proto_rawDesc
)

func file_nginx_nginx_config_proto_rawDescGZIP() []byte {
	file_nginx_nginx_config_proto_rawDescOnce.Do(func() {
		file_nginx_nginx_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_nginx_nginx_config_proto_rawDescData)
	})
	return file_nginx_nginx_config_proto_rawDescData
}

var file_nginx_nginx_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nginx_nginx_config_proto_goTypes = []interface{}{
	(*NginxConfig)(nil), // 0: universal_ingress_controller.nginx.NginxConfig
}
var file_nginx_nginx_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nginx_nginx_config_proto_init() }
func file_nginx_nginx_config_proto_init() {
	if File_nginx_nginx_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nginx_nginx_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NginxConfig); i {
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
			RawDescriptor: file_nginx_nginx_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nginx_nginx_config_proto_goTypes,
		DependencyIndexes: file_nginx_nginx_config_proto_depIdxs,
		MessageInfos:      file_nginx_nginx_config_proto_msgTypes,
	}.Build()
	File_nginx_nginx_config_proto = out.File
	file_nginx_nginx_config_proto_rawDesc = nil
	file_nginx_nginx_config_proto_goTypes = nil
	file_nginx_nginx_config_proto_depIdxs = nil
}
