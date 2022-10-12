package demo_sms

import (
	"encoding/json"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/sms"
)

type sendSmsTemplateParam struct {
	Code string `json:"code"`
}

/*
send same content to one/multiple phoneNumbers
*/
func TestSmsSendI18n(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	c := sendSmsTemplateParam{Code: "111"}
	cj, _ := json.Marshal(c)
	req := &sms.SmsRequest{
		SmsAccount:    "smsAccount",
		From:          "BytePlus",
		TemplateID:    "ST_xxx",
		TemplateParam: string(cj),
		PhoneNumbers:  "+65xxxxxxxx,+65xxxxxxxx", // one or multiple
		Tag:           "tag",
	}
	result, statusCode, err := i18nInstance.Send(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)

	if result == nil {
		t.Logf("result is nil")
		return
	}
	if result.ResponseMetadata.Error != nil {
		t.Logf("result.ResponseMetadata.Error = %+v\n", result.ResponseMetadata.Error)
		return
	}
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
}
