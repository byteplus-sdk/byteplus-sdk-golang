// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: byteplus/vod/business/vod_project.proto

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

type VodListProjectsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectList []*VodProjectInfo `protobuf:"bytes,1,rep,name=ProjectList,proto3" json:"ProjectList,omitempty"` // 项目列表
}

func (x *VodListProjectsResult) Reset() {
	*x = VodListProjectsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_byteplus_vod_business_vod_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodListProjectsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodListProjectsResult) ProtoMessage() {}

func (x *VodListProjectsResult) ProtoReflect() protoreflect.Message {
	mi := &file_byteplus_vod_business_vod_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodListProjectsResult.ProtoReflect.Descriptor instead.
func (*VodListProjectsResult) Descriptor() ([]byte, []int) {
	return file_byteplus_vod_business_vod_project_proto_rawDescGZIP(), []int{0}
}

func (x *VodListProjectsResult) GetProjectList() []*VodProjectInfo {
	if x != nil {
		return x.ProjectList
	}
	return nil
}

type VodProjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=ProjectName,proto3" json:"ProjectName,omitempty"` // 项目名称
	DisplayName string `protobuf:"bytes,2,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"` // 项目展示名称
}

func (x *VodProjectInfo) Reset() {
	*x = VodProjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_byteplus_vod_business_vod_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodProjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodProjectInfo) ProtoMessage() {}

func (x *VodProjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_byteplus_vod_business_vod_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodProjectInfo.ProtoReflect.Descriptor instead.
func (*VodProjectInfo) Descriptor() ([]byte, []int) {
	return file_byteplus_vod_business_vod_project_proto_rawDescGZIP(), []int{1}
}

func (x *VodProjectInfo) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *VodProjectInfo) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_byteplus_vod_business_vod_project_proto protoreflect.FileDescriptor

var file_byteplus_vod_business_vod_project_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x42, 0x79, 0x74, 0x65, 0x70,
	0x6c, 0x75, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x67, 0x0a, 0x15, 0x56, 0x6f, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x4e, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73,
	0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x54, 0x0a, 0x0e, 0x56, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xda, 0x01, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x6f, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x42, 0x0a, 0x56, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x01,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x74,
	0x65, 0x70, 0x6c, 0x75, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c,
	0x75, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2,
	0x02, 0x00, 0xca, 0x02, 0x24, 0x42, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x27, 0x42, 0x79, 0x74, 0x65,
	0x70, 0x6c, 0x75, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_byteplus_vod_business_vod_project_proto_rawDescOnce sync.Once
	file_byteplus_vod_business_vod_project_proto_rawDescData = file_byteplus_vod_business_vod_project_proto_rawDesc
)

func file_byteplus_vod_business_vod_project_proto_rawDescGZIP() []byte {
	file_byteplus_vod_business_vod_project_proto_rawDescOnce.Do(func() {
		file_byteplus_vod_business_vod_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_byteplus_vod_business_vod_project_proto_rawDescData)
	})
	return file_byteplus_vod_business_vod_project_proto_rawDescData
}

var file_byteplus_vod_business_vod_project_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_byteplus_vod_business_vod_project_proto_goTypes = []interface{}{
	(*VodListProjectsResult)(nil), // 0: Byteplus.Vod.Models.Business.VodListProjectsResult
	(*VodProjectInfo)(nil),        // 1: Byteplus.Vod.Models.Business.VodProjectInfo
}
var file_byteplus_vod_business_vod_project_proto_depIdxs = []int32{
	1, // 0: Byteplus.Vod.Models.Business.VodListProjectsResult.ProjectList:type_name -> Byteplus.Vod.Models.Business.VodProjectInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_byteplus_vod_business_vod_project_proto_init() }
func file_byteplus_vod_business_vod_project_proto_init() {
	if File_byteplus_vod_business_vod_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_byteplus_vod_business_vod_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodListProjectsResult); i {
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
		file_byteplus_vod_business_vod_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodProjectInfo); i {
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
			RawDescriptor: file_byteplus_vod_business_vod_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_byteplus_vod_business_vod_project_proto_goTypes,
		DependencyIndexes: file_byteplus_vod_business_vod_project_proto_depIdxs,
		MessageInfos:      file_byteplus_vod_business_vod_project_proto_msgTypes,
	}.Build()
	File_byteplus_vod_business_vod_project_proto = out.File
	file_byteplus_vod_business_vod_project_proto_rawDesc = nil
	file_byteplus_vod_business_vod_project_proto_goTypes = nil
	file_byteplus_vod_business_vod_project_proto_depIdxs = nil
}
