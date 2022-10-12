package demo_sms

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/sms"
)

func TestGetSmsTemplateAndOrderList(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &sms.GetSmsTemplateAndOrderListRequest{
		SubAccount: "smsAccount",
		PageIndex:  1,
		PageSize:   10,
	}
	result, statusCode, err := i18nInstance.GetSmsTemplateAndOrderList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsTemplate(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &sms.ApplySmsTemplateRequest{
		SubAccount:  "smsAccount",
		ChannelType: sms.SmsChannelTypeI18nMKT,
		Content:     "We're offering our xxxx community 20% off with code THANKYOU. It's our way of showing our appreciation to you for standing by us in this time. Shop Now: https://webhook.site/edd2b39a-6c8d-4161-a310-36a470c840d4",
		Desc:        "Test SDK",
		Name:        "template_name",
		ShortUrlConfig: &sms.ShortUrlConfig{
			IsEnabled:          sms.EnableStatusEnabled,
			IsNeedClickDetails: sms.EnableStatusNotEnabled,
		},
	}
	result, statusCode, err := i18nInstance.ApplySmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteSmsTemplate(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &sms.DeleteSmsTemplateRequest{
		SubAccount: "smsAccount",
		Id:         "idOfTemplateToDelete",
		IsOrder:    true, // whether it is a order
	}
	result, statusCode, err := i18nInstance.DeleteSmsTemplate(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
