package live_v20200801

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

const (
	ServiceName    = "live"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-north-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "open.byteplusapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-north-1",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"ListStorageSpace": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListStorageSpace"},
				"Version": []string{"2020-08-01"},
			},
		},
		"ListStorageSpaceDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListStorageSpaceDetail"},
				"Version": []string{"2020-08-01"},
			},
		},
	}
)
