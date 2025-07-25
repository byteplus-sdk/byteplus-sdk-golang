package sms

import (
	"net/http"
	"net/url"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20200101 = "2020-01-01"
	ServiceVersion20210101 = "2021-01-01"
	ServiceVersion20210111 = "2021-01-11"
	ServiceName            = "volcSMS"
)

var (
	ServiceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Timeout: 5 * time.Second,
			Host:    "sms.volcengineapi.com",
			Scheme:  "https",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "sms.byteplusapi.com",
			Scheme:  "https",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"SendSms": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendSms"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"SendBatchSms": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendBatchSms"},
				"Version": []string{ServiceVersion20210101},
			},
		},
		"SendSmsVerifyCode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendSmsVerifyCode"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"CheckSmsVerifyCode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CheckSmsVerifyCode"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"Conversion": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Conversion"},
				"Version": []string{ServiceVersion20200101},
			},
		},
		"GetSmsTemplateAndOrderList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSmsTemplateAndOrderList"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"ApplySmsTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplySmsTemplate"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"DeleteSmsTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSmsTemplate"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetSubAccountList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSubAccountList"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"GetSubAccountDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSubAccountDetail"},
				"Version": []string{ServiceVersion20210111},
			},
		},
		"InsertSubAccount": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"InsertSubAccount"},
				"Version": []string{ServiceVersion20210111},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type SMS struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *SMS {
	instance := &SMS{}
	instance.Client = base.NewClient(ServiceInfo[DefaultRegion], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

func NewInstanceI18n(region string) *SMS {
	instance := &SMS{}
	instance.Client = base.NewClient(ServiceInfo[base.RegionApSingapore], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = region
	return instance
}

func NewInstanceI18nWithProxy(region string, proxy *url.URL) *SMS {
	serviceInfo := ServiceInfo[base.RegionApSingapore]
	serviceInfo.Proxy = proxy

	instance := &SMS{}
	instance.Client = base.NewClient(serviceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = region
	return instance
}

// GetServiceInfo interface
func (p *SMS) GetServiceInfo(region string) *base.ServiceInfo {
	if serviceInfo, ok := ServiceInfo[region]; ok {
		return serviceInfo
	}
	return nil
}

// GetAPIInfo interface
func (p *SMS) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (s *SMS) SetRegion(region string) {
	if serviceInfo := s.GetServiceInfo(region); serviceInfo != nil {
		serviceInfo.Credentials.Region = region
	}
}

// SetHost .
func (s *SMS) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *SMS) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}
