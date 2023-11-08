package billing

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"net/http"
	"net/url"
)

var (
	ActionListBillDetail      = "ListBillDetail"
	ActionListSplitBillDetail = "ListSplitBillDetail"
)

var ApiInfoList = map[string]*base.ApiInfo{
	ActionListBillDetail: {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{ActionListBillDetail},
			"Version": []string{ServiceVersion},
		},
	},
	ActionListSplitBillDetail: {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{ActionListSplitBillDetail},
			"Version": []string{ServiceVersion},
		},
	},
}
