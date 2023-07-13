package vod

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"net/http"
	"net/url"
	"sync"
	"time"
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
			Host:    "vod.byteplusapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "vod"},
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

		// **********************************************************************
		// 上传
		// **********************************************************************
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
			Timeout: 8 * time.Second,
		},

		// **********************************************************************
		// 媒资
		// **********************************************************************

		// **********************************************************************
		// 转码
		// **********************************************************************

		// **********************************************************************
		// 视频编辑
		// **********************************************************************

		// **********************************************************************
		// 空间管理
		// **********************************************************************

		"ListSpace": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSpace"},
				"Version": []string{"2023-01-01"},
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

		// **********************************************************************
		// 回调管理
		// **********************************************************************

		// **********************************************************************
		// 计量计费
		// **********************************************************************

	}
)
