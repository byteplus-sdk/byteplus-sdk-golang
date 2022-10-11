package sms

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

func TestGetSubAccountList(t *testing.T) {
	i18nInstance := NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &GetSubAccountListRequest{
		SubAccount: "smsAccount",
		PageIndex:  1,
		PageSize:   10,
	}
	result, statusCode, err := i18nInstance.GetSubAccountList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestGetSubAccountDetail(t *testing.T) {
	i18nInstance := NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &GetSubAccountDetailRequest{
		SubAccount: "smsAccount",
	}
	result, statusCode, err := i18nInstance.GetSubAccountDetail(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
