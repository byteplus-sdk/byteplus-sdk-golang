package livesaas

import (
	"net/http"
	"net/url"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

const (
	DefaultRegion          = base.RegionApSingapore
	ServiceVersion20200601 = "2020-06-01"
	ServiceName            = "livesaas"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.byteplusapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"CreateActivityAPIV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateActivityAPIV2"},
				"Version": []string{ServiceVersion20200601},
			},
		},
		"GetTemporaryLoginTokenAPI": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTemporaryLoginTokenAPI"},
				"Version": []string{ServiceVersion20200601},
			},
		},
		"GetSDKTokenAPI": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSDKTokenAPI"},
				"Version": []string{ServiceVersion20200601},
			},
		},
		"ListActivityAPI": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListActivityAPI"},
				"Version": []string{ServiceVersion20200601},
			},
		},
		"UpdateActivityBasicConfigAPI": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateActivityBasicConfigAPI"},
				"Version": []string{ServiceVersion20200601},
			},
		},
	}
)

// DefaultInstance
var DefaultInstance = NewInstance()

// LIVE .
type LIVESAAS struct {
	Client *base.Client
}

// NewInstance create a instance
func NewInstance() *LIVESAAS {
	instance := &LIVESAAS{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *LIVESAAS) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *LIVESAAS) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *LIVESAAS) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *LIVESAAS) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *LIVESAAS) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
