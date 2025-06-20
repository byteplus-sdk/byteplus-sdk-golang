// Code generated by protoc-gen-byteplus-sdk
// source: VodMediaService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"
)

func Test_UpdateMediaInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateMediaInfoRequest{
		Vid:              "your Vid",
		PosterUri:        nil,
		Title:            nil,
		Description:      nil,
		Tags:             nil,
		ClassificationId: nil,
	}

	resp, status, err := instance.UpdateMediaInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateMediaPublishStatus(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateMediaPublishStatusRequest{
		Vid:    "your Vid",
		Status: "your Status",
	}

	resp, status, err := instance.UpdateMediaPublishStatus(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateMediaStorageClass(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateMediaStorageClassRequest{
		Vids:         "your Vids",
		StorageClass: "your StorageClass",
		CallbackArgs: "your CallbackArgs",
		FileIds:      "your FileIds",
	}

	resp, status, err := instance.UpdateMediaStorageClass(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetMediaInfos(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetMediaInfosRequest{
		Vids: "your Vids",
	}

	resp, status, err := instance.GetMediaInfos(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetRecommendedPoster(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetRecommendedPosterRequest{
		Vids: "your Vids",
	}

	resp, status, err := instance.GetRecommendedPoster(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteMedia(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteMediaRequest{
		Vids:         "your Vids",
		CallbackArgs: "your CallbackArgs",
	}

	resp, status, err := instance.DeleteMedia(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteTranscodes(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteTranscodesRequest{
		Vid:          "your Vid",
		FileIds:      "your FileIds",
		CallbackArgs: "your CallbackArgs",
	}

	resp, status, err := instance.DeleteTranscodes(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetMediaList(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetMediaListRequest{
		SpaceName:         "your SpaceName",
		Vid:               "your Vid",
		Status:            "your Status",
		Order:             "your Order",
		Tags:              "your Tags",
		StartTime:         "your StartTime",
		EndTime:           "your EndTime",
		Offset:            "your Offset",
		PageSize:          "your PageSize",
		ClassificationIds: "your ClassificationIds",
		TosStorageClasses: "your TosStorageClasses",
		VodUploadSources:  "your VodUploadSources",
	}

	resp, status, err := instance.GetMediaList(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetSubtitleInfoList(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetSubtitleInfoListRequest{
		Vid:         "your Vid",
		FileIds:     "your FileIds",
		Languages:   "your Languages",
		Formats:     "your Formats",
		LanguageIds: "your LanguageIds",
		SubtitleIds: "your SubtitleIds",
		Status:      "your Status",
		Title:       "your Title",
		Tag:         "your Tag",
		Offset:      "your Offset",
		PageSize:    "your PageSize",
		Ssl:         "your Ssl",
	}

	resp, status, err := instance.GetSubtitleInfoList(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSubtitleStatus(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSubtitleStatusRequest{
		Vid:       "your Vid",
		FileIds:   "your FileIds",
		Languages: "your Languages",
		Formats:   "your Formats",
		Status:    "your Status",
	}

	resp, status, err := instance.UpdateSubtitleStatus(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSubtitleInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSubtitleInfoRequest{
		Vid:      "your Vid",
		FileId:   "your FileId",
		Language: "your Language",
		Format:   "your Format",
		Title:    nil,
		Tag:      nil,
	}

	resp, status, err := instance.UpdateSubtitleInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreatePlaylist(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCreatePlaylistRequest{
		SpaceName:  "your SpaceName",
		Name:       "your Name",
		Format:     "your Format",
		Codec:      "your Codec",
		Definition: "your Definition",
		Vids:       "your Vids",
		StartTime:  "your StartTime",
		EndTime:    "your EndTime",
		Cycles:     "your Cycles",
	}

	resp, status, err := instance.CreatePlaylist(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetPlaylists(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetPlaylistsRequest{
		SpaceName: "your SpaceName",
		Ids:       "your Ids",
		Name:      "your Name",
		Limit:     0,
		Offset:    0,
	}

	resp, status, err := instance.GetPlaylists(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdatePlaylist(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdatePlaylistRequest{
		SpaceName:  "your SpaceName",
		Id:         "your Id",
		Name:       nil,
		Format:     nil,
		Codec:      nil,
		Definition: nil,
		Vids:       nil,
		StartTime:  nil,
		EndTime:    nil,
		Cycles:     nil,
	}

	resp, status, err := instance.UpdatePlaylist(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeletePlaylist(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeletePlaylistRequest{
		SpaceName: "your SpaceName",
		Id:        "your Id",
	}

	resp, status, err := instance.DeletePlaylist(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreateVideoClassification(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCreateVideoClassificationRequest{
		SpaceName:      "your SpaceName",
		Level:          0,
		ParentId:       0,
		Classification: "your Classification",
	}

	resp, status, err := instance.CreateVideoClassification(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateVideoClassification(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateVideoClassificationRequest{
		SpaceName:        "your SpaceName",
		ClassificationId: 0,
		Classification:   "your Classification",
	}

	resp, status, err := instance.UpdateVideoClassification(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteVideoClassification(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteVideoClassificationRequest{
		SpaceName:        "your SpaceName",
		ClassificationId: 0,
	}

	resp, status, err := instance.DeleteVideoClassification(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListVideoClassifications(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListVideoClassificationsRequest{
		SpaceName:        "your SpaceName",
		ClassificationId: 0,
	}

	resp, status, err := instance.ListVideoClassifications(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListSnapshots(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListSnapshotsRequest{
		Vid: "your Vid",
	}

	resp, status, err := instance.ListSnapshots(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
