// Code generated by protoc-gen-byteplus-sdk
// source: VodPlayService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"
)

func Test_GetPlayInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetPlayInfoRequest{
		Vid:                "your Vid",
		Format:             "your Format",
		Codec:              "your Codec",
		Definition:         "your Definition",
		FileType:           "your FileType",
		LogoType:           "your LogoType",
		Base64:             "your Base64",
		Ssl:                "your Ssl",
		NeedThumbs:         "your NeedThumbs",
		NeedBarrageMask:    "your NeedBarrageMask",
		CdnType:            "your CdnType",
		UnionInfo:          "your UnionInfo",
		HDRDefinition:      "your HDRDefinition",
		PlayScene:          "your PlayScene",
		DrmExpireTimestamp: "your DrmExpireTimestamp",
		Quality:            "your Quality",
		PlayConfig:         "your PlayConfig",
	}

	resp, status, err := instance.GetPlayInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetPlayInfoWithLiveTimeShiftScene(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetPlayInfoWithLiveTimeShiftSceneRequest{
		StoreUris:             "your StoreUris",
		SpaceName:             "your SpaceName",
		Ssl:                   "your Ssl",
		ExpireTimestamp:       "your ExpireTimestamp",
		NeedComposeBucketName: "your NeedComposeBucketName",
	}

	resp, status, err := instance.GetPlayInfoWithLiveTimeShiftScene(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
