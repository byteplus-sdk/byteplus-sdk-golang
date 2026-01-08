package visual

import (
	"net/http"
	"net/url"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

const (
	DefaultRegion          = "ap-singapore-1"
	ServiceVersion20200826 = "2020-08-26"
	ServiceName            = "cv"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 30 * time.Second,
		Host:    "cv.byteplusapi.com",
		Header:  http.Header{},
	}

	ApiInfoList = map[string]*base.ApiInfo{

		"CVProcess": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CVProcess"},
				"Version": []string{"2024-06-06"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CVSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CVSubmitTask"},
				"Version": []string{"2024-06-06"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CVGetResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CVGetResult"},
				"Version": []string{"2024-06-06"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CVCancelTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CVCancelTask"},
				"Version": []string{"2024-06-06"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CVSync2AsyncSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CVSync2AsyncSubmitTask"},
				"Version": []string{"2024-06-06"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CVSync2AsyncGetResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CVSync2AsyncGetResult"},
				"Version": []string{"2024-06-06"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
	}
)

// DefaultInstance Default Instance
var DefaultInstance = NewInstance()

// IAM .
type Visual struct {
	Client *base.Client
}

// NewInstance Create an instance
func NewInstance() *Visual {
	instance := &Visual{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	instance.Client.ServiceInfo.Scheme = "https"
	return instance
}

// GetServiceInfo interface
func (p *Visual) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

// GetAPIInfo interface
func (p *Visual) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion
func (p *Visual) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Visual) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *Visual) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
