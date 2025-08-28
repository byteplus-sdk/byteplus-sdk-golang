package vod

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

type Vod struct {
	*base.Client
	DomainCache map[string]map[string]int
	Lock        sync.RWMutex
}

func NewInstance() *Vod {
	instance := &Vod{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(ServiceInfoMap[base.RegionApSingapore], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *Vod {
	var serviceInfo *base.ServiceInfo
	var ok bool
	if serviceInfo, ok = ServiceInfoMap[region]; !ok {
		panic("Cant find the region, please check it carefully")
	}

	instance := &Vod{
		DomainCache: make(map[string]map[string]int),
		Client:      base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionApSingapore: {
			Timeout: 60 * time.Second,
			Scheme:  "https",
			Host:    "vod.byteplusapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "vod"},
		},
		base.RegionApSouthEast1: {
			Timeout: 60 * time.Second,
			Scheme:  "https",
			Host:    "vod.byteplusapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSouthEast1, Service: "vod"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		// **********************************************************************
		// 播放
		// **********************************************************************
		"GetPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetPlayInfoWithLiveTimeShiftScene": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfoWithLiveTimeShiftScene"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetPrivateDrmPlayAuth": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPrivateDrmPlayAuth"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetHlsDecryptionKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetHlsDecryptionKey"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateHlsDecryptionKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateHlsDecryptionKey"},
				"Version": []string{"2023-07-01"},
			},
		},
		// **********************************************************************
		// 上传
		// **********************************************************************
		"UploadMediaByUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadMediaByUrl"},
				"Version": []string{"2023-01-01"},
			},
		},
		"QueryUploadTaskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryUploadTaskInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ApplyUploadInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyUploadInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CommitUploadInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUploadInfo"},
				"Version": []string{"2023-01-01"},
			},
		},

		// **********************************************************************
		// 媒资
		// **********************************************************************
		"CreatePlaylist": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePlaylist"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetPlaylists": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlaylists"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdatePlaylist": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePlaylist"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeletePlaylist": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePlaylist"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateMediaInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateMediaInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateMediaPublishStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateMediaPublishStatus"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateMediaStorageClass": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateMediaStorageClass"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetMediaInfos": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMediaInfos"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetRecommendedPoster": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecommendedPoster"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteMedia": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteMedia"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteTranscodes": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTranscodes"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetMediaList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetMediaList"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetSubtitleInfoList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSubtitleInfoList"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateSubtitleStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSubtitleStatus"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateSubtitleInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSubtitleInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateVideoClassification": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateVideoClassification"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateVideoClassification": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateVideoClassification"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteVideoClassification": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteVideoClassification"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVideoClassifications": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVideoClassifications"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListSnapshots": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSnapshots"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetFileInfos": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetFileInfos"},
				"Version": []string{"2023-07-01"},
			},
		},
		"DeleteMediaTosFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteMediaTosFile"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListFileMetaInfosByFileNames": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListFileMetaInfosByFileNames"},
				"Version": []string{"2023-07-01"},
			},
		},
		// **********************************************************************
		// 转码
		// **********************************************************************
		"StartWorkflow": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWorkflow"},
				"Version": []string{"2023-01-01"},
			},
		},
		"RetrieveTranscodeResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RetrieveTranscodeResult"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetWorkflowExecution": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWorkflowExecution"},
				"Version": []string{"2023-01-01"},
			},
		},

		// **********************************************************************
		// 视频编辑
		// **********************************************************************

		// **********************************************************************
		// 空间管理
		// **********************************************************************
		"CreateSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSpace"},
				"Version": []string{"2023-07-01"},
			},
		},
		"ListSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSpace"},
				"Version": []string{"2023-07-01"},
			},
		},
		"GetSpaceDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSpaceDetail"},
				"Version": []string{"2023-07-01"},
			},
		},
		"UpdateSpaceUploadConfig": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSpaceUploadConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodSpaceStorageData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceStorageData"},
				"Version": []string{"2023-01-01"},
			},
		},
		// **********************************************************************
		// 分发加速管理
		// **********************************************************************
		"ListDomain": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateCdnRefreshTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCdnRefreshTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateCdnPreloadTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCdnPreloadTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCdnTasks": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnTasks"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCdnAccessLog": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnAccessLog"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCdnTopAccessUrl": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnTopAccessUrl"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodDomainBandwidthData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodDomainBandwidthData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCdnUsageData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnUsageData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCdnStatusData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnStatusData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeIpInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeIpInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodDomainTrafficData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodDomainTrafficData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCdnPvData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCdnPvData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateDomainExpire": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainExpire"},
				"Version": []string{"2023-01-01"},
			},
		},
		// **********************************************************************
		// 回调管理
		// **********************************************************************
		"AddCallbackSubscription": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddCallbackSubscription"},
				"Version": []string{"2023-01-01"},
			},
		},
		"SetCallbackEvent": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetCallbackEvent"},
				"Version": []string{"2023-01-01"},
			},
		},
		// **********************************************************************
		// 计量计费
		// **********************************************************************
		"DescribeVodSpaceTranscodeData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceTranscodeData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodSnapshotData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSnapshotData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodEnhanceImageData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodEnhanceImageData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodSpaceSubtitleStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodSpaceSubtitleStatisData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodPlayedStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodPlayedStatisData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodMostPlayedStatisData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodMostPlayedStatisData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVodRealtimeMediaData": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVodRealtimeMediaData"},
				"Version": []string{"2023-01-01"},
			},
		},
		// **********************************************************************
		// 商业drm
		// **********************************************************************
		"GetDrmLicense": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDrmLicense"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetFairPlayCert": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetFairPlayCert"},
				"Version": []string{"2023-01-01"},
			},
		},
	}
)
