// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: vod/business/vod_play.proto

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

type VodGetOriginalPlayInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileType      string  `protobuf:"bytes,1,opt,name=FileType,proto3" json:"FileType,omitempty"`            //返回的媒体类型(video/audio)
	Duration      float32 `protobuf:"fixed32,2,opt,name=Duration,proto3" json:"Duration,omitempty"`          //视频时长(单位：s)
	Size          float64 `protobuf:"fixed64,3,opt,name=Size,proto3" json:"Size,omitempty"`                  //视频文件大小
	Height        int32   `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`               //视频高度
	Width         int32   `protobuf:"varint,5,opt,name=Width,proto3" json:"Width,omitempty"`                 //视频宽度
	Format        string  `protobuf:"bytes,6,opt,name=Format,proto3" json:"Format,omitempty"`                //视频格式
	Codec         string  `protobuf:"bytes,7,opt,name=Codec,proto3" json:"Codec,omitempty"`                  //编码类型
	Bitrate       int32   `protobuf:"varint,8,opt,name=Bitrate,proto3" json:"Bitrate,omitempty"`             //码率(Kbps)
	Md5           string  `protobuf:"bytes,9,opt,name=Md5,proto3" json:"Md5,omitempty"`                      // hash值
	MainPlayUrl   string  `protobuf:"bytes,10,opt,name=MainPlayUrl,proto3" json:"MainPlayUrl,omitempty"`     //主播放地址
	BackupPlayUrl string  `protobuf:"bytes,11,opt,name=BackupPlayUrl,proto3" json:"BackupPlayUrl,omitempty"` //备用播放地址
}

func (x *VodGetOriginalPlayInfoResult) Reset() {
	*x = VodGetOriginalPlayInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetOriginalPlayInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetOriginalPlayInfoResult) ProtoMessage() {}

func (x *VodGetOriginalPlayInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetOriginalPlayInfoResult.ProtoReflect.Descriptor instead.
func (*VodGetOriginalPlayInfoResult) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{0}
}

func (x *VodGetOriginalPlayInfoResult) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetSize() float64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetCodec() string {
	if x != nil {
		return x.Codec
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetBitrate() int32 {
	if x != nil {
		return x.Bitrate
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetMainPlayUrl() string {
	if x != nil {
		return x.MainPlayUrl
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetBackupPlayUrl() string {
	if x != nil {
		return x.BackupPlayUrl
	}
	return ""
}

type VodPrivateDrmPlayAuthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayAuthId      string `protobuf:"bytes,1,opt,name=PlayAuthId,proto3" json:"PlayAuthId,omitempty"`           // 播放许可id
	PlayAuthContent string `protobuf:"bytes,2,opt,name=PlayAuthContent,proto3" json:"PlayAuthContent,omitempty"` // 播放许可内容
}

func (x *VodPrivateDrmPlayAuthInfo) Reset() {
	*x = VodPrivateDrmPlayAuthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodPrivateDrmPlayAuthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodPrivateDrmPlayAuthInfo) ProtoMessage() {}

func (x *VodPrivateDrmPlayAuthInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodPrivateDrmPlayAuthInfo.ProtoReflect.Descriptor instead.
func (*VodPrivateDrmPlayAuthInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{1}
}

func (x *VodPrivateDrmPlayAuthInfo) GetPlayAuthId() string {
	if x != nil {
		return x.PlayAuthId
	}
	return ""
}

func (x *VodPrivateDrmPlayAuthInfo) GetPlayAuthContent() string {
	if x != nil {
		return x.PlayAuthContent
	}
	return ""
}

type VodGetPrivateDrmPlayAuthResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayAuthInfoList []*VodPrivateDrmPlayAuthInfo `protobuf:"bytes,1,rep,name=PlayAuthInfoList,proto3" json:"PlayAuthInfoList,omitempty"` // 播放许可信息列表
}

func (x *VodGetPrivateDrmPlayAuthResult) Reset() {
	*x = VodGetPrivateDrmPlayAuthResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetPrivateDrmPlayAuthResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetPrivateDrmPlayAuthResult) ProtoMessage() {}

func (x *VodGetPrivateDrmPlayAuthResult) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetPrivateDrmPlayAuthResult.ProtoReflect.Descriptor instead.
func (*VodGetPrivateDrmPlayAuthResult) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{2}
}

func (x *VodGetPrivateDrmPlayAuthResult) GetPlayAuthInfoList() []*VodPrivateDrmPlayAuthInfo {
	if x != nil {
		return x.PlayAuthInfoList
	}
	return nil
}

type VodGetHlsDecryptionKeyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretKey string `protobuf:"bytes,1,opt,name=SecretKey,proto3" json:"SecretKey,omitempty"` // Drm密钥
}

func (x *VodGetHlsDecryptionKeyResult) Reset() {
	*x = VodGetHlsDecryptionKeyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetHlsDecryptionKeyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetHlsDecryptionKeyResult) ProtoMessage() {}

func (x *VodGetHlsDecryptionKeyResult) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetHlsDecryptionKeyResult.ProtoReflect.Descriptor instead.
func (*VodGetHlsDecryptionKeyResult) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{3}
}

func (x *VodGetHlsDecryptionKeyResult) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

type VodPlayInfoWithLiveTimeShiftScene struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreUri      string  `protobuf:"bytes,1,opt,name=StoreUri,proto3" json:"StoreUri,omitempty"`           // URI
	MainPlayUrl   string  `protobuf:"bytes,2,opt,name=MainPlayUrl,proto3" json:"MainPlayUrl,omitempty"`     // 主播放URL
	BackupPlayUrl string  `protobuf:"bytes,3,opt,name=BackupPlayUrl,proto3" json:"BackupPlayUrl,omitempty"` // 备播放URL，可能不存在
	UrlExpire     float64 `protobuf:"fixed64,4,opt,name=UrlExpire,proto3" json:"UrlExpire,omitempty"`       // 过期DDL,Unix时间戳，单位：秒
}

func (x *VodPlayInfoWithLiveTimeShiftScene) Reset() {
	*x = VodPlayInfoWithLiveTimeShiftScene{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodPlayInfoWithLiveTimeShiftScene) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodPlayInfoWithLiveTimeShiftScene) ProtoMessage() {}

func (x *VodPlayInfoWithLiveTimeShiftScene) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodPlayInfoWithLiveTimeShiftScene.ProtoReflect.Descriptor instead.
func (*VodPlayInfoWithLiveTimeShiftScene) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{4}
}

func (x *VodPlayInfoWithLiveTimeShiftScene) GetStoreUri() string {
	if x != nil {
		return x.StoreUri
	}
	return ""
}

func (x *VodPlayInfoWithLiveTimeShiftScene) GetMainPlayUrl() string {
	if x != nil {
		return x.MainPlayUrl
	}
	return ""
}

func (x *VodPlayInfoWithLiveTimeShiftScene) GetBackupPlayUrl() string {
	if x != nil {
		return x.BackupPlayUrl
	}
	return ""
}

func (x *VodPlayInfoWithLiveTimeShiftScene) GetUrlExpire() float64 {
	if x != nil {
		return x.UrlExpire
	}
	return 0
}

type VodGetPlayInfoWithLiveTimeShiftSceneResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayInfoList []*VodPlayInfoWithLiveTimeShiftScene `protobuf:"bytes,1,rep,name=PlayInfoList,proto3" json:"PlayInfoList,omitempty"` // 播放信息列表
}

func (x *VodGetPlayInfoWithLiveTimeShiftSceneResult) Reset() {
	*x = VodGetPlayInfoWithLiveTimeShiftSceneResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetPlayInfoWithLiveTimeShiftSceneResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetPlayInfoWithLiveTimeShiftSceneResult) ProtoMessage() {}

func (x *VodGetPlayInfoWithLiveTimeShiftSceneResult) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetPlayInfoWithLiveTimeShiftSceneResult.ProtoReflect.Descriptor instead.
func (*VodGetPlayInfoWithLiveTimeShiftSceneResult) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{5}
}

func (x *VodGetPlayInfoWithLiveTimeShiftSceneResult) GetPlayInfoList() []*VodPlayInfoWithLiveTimeShiftScene {
	if x != nil {
		return x.PlayInfoList
	}
	return nil
}

var File_vod_business_vod_play_proto protoreflect.FileDescriptor

var file_vod_business_vod_play_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x6f, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x42,
	0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0xba, 0x02, 0x0a, 0x1c,
	0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x6c,
	0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43,
	0x6f, 0x64, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x64, 0x35, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x64, 0x35,
	0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x55,
	0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x79,
	0x55, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x22, 0x65, 0x0a, 0x19, 0x56, 0x6f, 0x64, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x72, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x85, 0x01, 0x0a, 0x1e, 0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x44, 0x72, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x63, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x42,
	0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x72, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x1c, 0x56, 0x6f, 0x64, 0x47, 0x65,
	0x74, 0x48, 0x6c, 0x73, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x22, 0xa5, 0x01, 0x0a, 0x21, 0x56, 0x6f, 0x64, 0x50, 0x6c, 0x61,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x68, 0x69, 0x66, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x55, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x55, 0x72, 0x69, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50,
	0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x61,
	0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x72, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x55, 0x72, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x91, 0x01,
	0x0a, 0x2a, 0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x57, 0x69, 0x74, 0x68, 0x4c, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x68, 0x69, 0x66,
	0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x63, 0x0a, 0x0c,
	0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x56, 0x6f,
	0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x56, 0x6f, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x69, 0x74,
	0x68, 0x4c, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x68, 0x69, 0x66, 0x74, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x52, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0xd7, 0x01, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c,
	0x75, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x6f, 0x64, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x07, 0x56,
	0x6f, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f,
	0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x24, 0x42, 0x79, 0x74,
	0x65, 0x70, 0x6c, 0x75, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f,
	0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0xe2, 0x02, 0x27, 0x42, 0x79, 0x74, 0x65, 0x70, 0x6c, 0x75, 0x73, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_vod_business_vod_play_proto_rawDescOnce sync.Once
	file_vod_business_vod_play_proto_rawDescData = file_vod_business_vod_play_proto_rawDesc
)

func file_vod_business_vod_play_proto_rawDescGZIP() []byte {
	file_vod_business_vod_play_proto_rawDescOnce.Do(func() {
		file_vod_business_vod_play_proto_rawDescData = protoimpl.X.CompressGZIP(file_vod_business_vod_play_proto_rawDescData)
	})
	return file_vod_business_vod_play_proto_rawDescData
}

var file_vod_business_vod_play_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vod_business_vod_play_proto_goTypes = []interface{}{
	(*VodGetOriginalPlayInfoResult)(nil),               // 0: Byteplus.Vod.Models.Business.VodGetOriginalPlayInfoResult
	(*VodPrivateDrmPlayAuthInfo)(nil),                  // 1: Byteplus.Vod.Models.Business.VodPrivateDrmPlayAuthInfo
	(*VodGetPrivateDrmPlayAuthResult)(nil),             // 2: Byteplus.Vod.Models.Business.VodGetPrivateDrmPlayAuthResult
	(*VodGetHlsDecryptionKeyResult)(nil),               // 3: Byteplus.Vod.Models.Business.VodGetHlsDecryptionKeyResult
	(*VodPlayInfoWithLiveTimeShiftScene)(nil),          // 4: Byteplus.Vod.Models.Business.VodPlayInfoWithLiveTimeShiftScene
	(*VodGetPlayInfoWithLiveTimeShiftSceneResult)(nil), // 5: Byteplus.Vod.Models.Business.VodGetPlayInfoWithLiveTimeShiftSceneResult
}
var file_vod_business_vod_play_proto_depIdxs = []int32{
	1, // 0: Byteplus.Vod.Models.Business.VodGetPrivateDrmPlayAuthResult.PlayAuthInfoList:type_name -> Byteplus.Vod.Models.Business.VodPrivateDrmPlayAuthInfo
	4, // 1: Byteplus.Vod.Models.Business.VodGetPlayInfoWithLiveTimeShiftSceneResult.PlayInfoList:type_name -> Byteplus.Vod.Models.Business.VodPlayInfoWithLiveTimeShiftScene
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vod_business_vod_play_proto_init() }
func file_vod_business_vod_play_proto_init() {
	if File_vod_business_vod_play_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vod_business_vod_play_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetOriginalPlayInfoResult); i {
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
		file_vod_business_vod_play_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodPrivateDrmPlayAuthInfo); i {
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
		file_vod_business_vod_play_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetPrivateDrmPlayAuthResult); i {
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
		file_vod_business_vod_play_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetHlsDecryptionKeyResult); i {
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
		file_vod_business_vod_play_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodPlayInfoWithLiveTimeShiftScene); i {
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
		file_vod_business_vod_play_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetPlayInfoWithLiveTimeShiftSceneResult); i {
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
			RawDescriptor: file_vod_business_vod_play_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vod_business_vod_play_proto_goTypes,
		DependencyIndexes: file_vod_business_vod_play_proto_depIdxs,
		MessageInfos:      file_vod_business_vod_play_proto_msgTypes,
	}.Build()
	File_vod_business_vod_play_proto = out.File
	file_vod_business_vod_play_proto_rawDesc = nil
	file_vod_business_vod_play_proto_goTypes = nil
	file_vod_business_vod_play_proto_depIdxs = nil
}
