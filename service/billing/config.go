package billing

import (
	"net/http"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

const (
	DefaultRegion  = base.RegionApSingapore
	ServiceHost    = "open.byteplusapi.com"
	ServiceVersion = "2022-01-01"
	ServiceName    = "billing"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Host:    ServiceHost,
		Timeout: 1 * time.Minute,
		Header: http.Header{
			"Accept":       []string{"application/json"},
			"Content-Type": []string{"application/json"},
		},
	}
)

type Billing struct {
	Client *base.Client
}

func NewInstance() *Billing {
	instance := new(Billing)
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// DefaultInstance Package level default instance
var DefaultInstance = NewInstance()

func (b *Billing) GetServiceInfo() *base.ServiceInfo {
	return b.Client.ServiceInfo
}

func (b *Billing) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

func (b *Billing) SetRegion(region string) {
	b.Client.ServiceInfo.Credentials.Region = region
}

func (b *Billing) SetHost(host string) {
	b.Client.ServiceInfo.Host = host
}

func (b *Billing) SetSchema(schema string) {
	b.Client.ServiceInfo.Scheme = schema
}

func (b *Billing) SetHeader(k string, v []string) {
	b.Client.ServiceInfo.Header[k] = v
}
