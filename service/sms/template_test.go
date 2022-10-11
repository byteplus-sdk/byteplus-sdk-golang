package sms

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

func TestGetSmsTemplateList(t *testing.T) {
	i18nInstance := NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &GetSmsTemplateAndOrderListRequest{
		SubAccount: "smsAccount",
	}
	result, statusCode, err := i18nInstance.GetSmsTemplateAndOrderList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsTemplate(t *testing.T) {
	i18nInstance := NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &ApplySmsTemplateRequest{
		SubAccount:  "smsAccount",
		ChannelType: SmsChannelTypeI18nMKT,
		Content:     "We're offering our xxxx community 20% off with code THANKYOU. It's our way of showing our appreciation to you for standing by us in this time. Shop Now: https://webhook.site/edd2b39a-6c8d-4161-a310-36a470c840d4",
		Desc:        "Test SDK",
		Name:        "template_name",
		ShortUrlConfig: &ShortUrlConfig{
			IsEnabled:          EnableStatusEnabled,
			IsNeedClickDetails: EnableStatusNotEnabled,
		},
	}
	result, statusCode, err := i18nInstance.ApplySmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteSmsTemplate(t *testing.T) {
	i18nInstance := NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &DeleteSmsTemplateRequest{
		SubAccount: "smsAccount",
		Id:         "idOfTemplateToDelete",
		IsOrder:    true, // 是否是审核工单
	}
	result, statusCode, err := i18nInstance.DeleteSmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
