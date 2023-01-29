package rtc

import (
	"net/http"
	"net/url"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

// RTC ... use base client
type RTC struct {
	Client *base.Client
}

// NewInstance ...
func NewInstance() *RTC {
	instance := &RTC{}
	instance.Client = base.NewClient(ServiceInfo, DefaultApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

const (
	DefaultRegion          = base.RegionApSingapore
	ServiceVersion20201201 = "2020-12-01"
	ServiceVersion20220601 = "2022-06-01"
	ServiceName            = "rtc"
	//Please replace it with the address in the official website document
	ServiceHost = "open.byteplusapi.com"

	// action name
	ActionStartRecord   = "StartRecord"
	ActionGetRecordTask = "GetRecordTask"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    ServiceHost,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	DefaultApiInfoList = map[string]*base.ApiInfo{
		ActionStartRecord: {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionStartRecord},
				"Version": []string{ServiceVersion20220601},
			},
		},
		ActionGetRecordTask: {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{ActionGetRecordTask},
				"Version": []string{ServiceVersion20220601},
			},
		},

		//ActionExample, add new action
		/*
			"ActionExample": {
				Method: http.MethodGet,
				Path:   "/",
				Query: url.Values{
					"Action":  []string{"ActionExample"},
					"Version": []string{ServiceVersion20201201},
				},
			},
		*/
	}
)
