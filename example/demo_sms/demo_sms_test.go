package demo_sms

import (
	"encoding/json"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/sms"
)

const (
	testAk = "testAK"
	testSk = "testSk"
)

type Code struct {
	Code string `json:"code"`
}

func TestSMS_sendI18n(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	c := Code{Code: "111"}
	cj, _ := json.Marshal(c)
	req := &sms.SmsRequest{
		SmsAccount:    "smsAccount",
		From:          "BytePlus",
		TemplateID:    "ST_xxx",
		TemplateParam: string(cj),
		PhoneNumbers:  "+65xxxxxxxx",
		Tag:           "tag",
	}
	result, statusCode, err := i18nInstance.Send(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSMS_conversion(t *testing.T) {
	client := sms.NewInstanceI18n("ap-singapore-1")
	client.Client.SetAccessKey(testAk)
	client.Client.SetSecretKey(testSk)
	req := &sms.ConversionRequest{
		MessageIDs: []string{"test_msg_id"},
	}
	result, statusCode, err := client.Conversion(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
