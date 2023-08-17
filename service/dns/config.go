package dns

import (
	"net/http"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

var (
	DefaultRegion  string
	ServiceVersion string
	ServiceName    string
	ServiceHost    string
	Timeout        int
	ServiceInfo    *base.ServiceInfo
)

func init() {
	ServiceName = "dns"
	ServiceHost = "open.byteplusapi.com"
	DefaultRegion = "ap-singapore-1"
	ServiceVersion = "2018-08-01"
	Timeout = 15

	ServiceInfo = &base.ServiceInfo{
		Timeout: time.Duration(Timeout) * time.Second,
		Host:    ServiceHost,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Scheme: "http",
		Credentials: base.Credentials{
			Service: ServiceName,
			Region:  DefaultRegion,
		},
	}
}
