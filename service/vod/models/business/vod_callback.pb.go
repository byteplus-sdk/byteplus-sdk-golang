// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: byteplus/vod/business/vod_callback.proto

package business

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VodCallbackAuthType int32

const (
	VodCallbackAuthType_VodCallbackAuthTypeDisabled VodCallbackAuthType = 0 // 回调鉴权关闭
	VodCallbackAuthType_VodCallbackAuthTypeEnabled  VodCallbackAuthType = 1 // 回调鉴权开启
)

// Enum value maps for VodCallbackAuthType.
var (
	VodCallbackAuthType_name = map[int32]string{
		0: "VodCallbackAuthTypeDisabled",
		1: "VodCallbackAuthTypeEnabled",
	}
	VodCallbackAuthType_value = map[string]int32{
		"VodCallbackAuthTypeDisabled": 0,
		"VodCallbackAuthTypeEnabled":  1,
	}
)

func (x VodCallbackAuthType) Enum() *VodCallbackAuthType {
	p := new(VodCallbackAuthType)
	*p = x
	return p
}

func (x VodCallbackAuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VodCallbackAuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_byteplus_vod_business_vod_callback_proto_enumTypes[0].Descriptor()
}

func (VodCallbackAuthType) Type() protoreflect.EnumType {
	return &file_byteplus_vod_business_vod_callback_proto_enumTypes[0]
}

func (x VodCallbackAuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VodCallbackAuthType.Descriptor instead.
func (VodCallbackAuthType) EnumDescriptor() ([]byte, []int) {
	return file_byteplus_vod_business_vod_callback_proto_rawDescGZIP(), []int{0}
}

var File_byteplus_vod_business_vod_callback_proto protoreflect.FileDescriptor

var file_byteplus_vod_business_vod_callback_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x42, 0x79, 0x74, 0x65,
	0x70, 0x6c, 0x75, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2a, 0x56, 0x0a, 0x13, 0x56, 0x6f, 0x64, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x1b, 0x56, 0x6f, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x1e, 0x0a, 0x1a, 0x56, 0x6f, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x01,
	0x42, 0xdb, 0x01, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x6f, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x56, 0x6f,
	0x64, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x24,
	0x42, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x27, 0x42, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_byteplus_vod_business_vod_callback_proto_rawDescOnce sync.Once
	file_byteplus_vod_business_vod_callback_proto_rawDescData = file_byteplus_vod_business_vod_callback_proto_rawDesc
)

func file_byteplus_vod_business_vod_callback_proto_rawDescGZIP() []byte {
	file_byteplus_vod_business_vod_callback_proto_rawDescOnce.Do(func() {
		file_byteplus_vod_business_vod_callback_proto_rawDescData = protoimpl.X.CompressGZIP(file_byteplus_vod_business_vod_callback_proto_rawDescData)
	})
	return file_byteplus_vod_business_vod_callback_proto_rawDescData
}

var file_byteplus_vod_business_vod_callback_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_byteplus_vod_business_vod_callback_proto_goTypes = []interface{}{
	(VodCallbackAuthType)(0), // 0: Byteplus.Vod.Models.Business.VodCallbackAuthType
}
var file_byteplus_vod_business_vod_callback_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_byteplus_vod_business_vod_callback_proto_init() }
func file_byteplus_vod_business_vod_callback_proto_init() {
	if File_byteplus_vod_business_vod_callback_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_byteplus_vod_business_vod_callback_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_byteplus_vod_business_vod_callback_proto_goTypes,
		DependencyIndexes: file_byteplus_vod_business_vod_callback_proto_depIdxs,
		EnumInfos:         file_byteplus_vod_business_vod_callback_proto_enumTypes,
	}.Build()
	File_byteplus_vod_business_vod_callback_proto = out.File
	file_byteplus_vod_business_vod_callback_proto_rawDesc = nil
	file_byteplus_vod_business_vod_callback_proto_goTypes = nil
	file_byteplus_vod_business_vod_callback_proto_depIdxs = nil
}
