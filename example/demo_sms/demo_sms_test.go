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

func TestSmsSendI18n(t *testing.T) {
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
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSmsBatchSendI18n(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)

	c := Code{Code: "111"}
	cj, _ := json.Marshal(c)
	messages := make([]*sms.SmsBatchMessages, 0)
	messages = append(messages, &sms.SmsBatchMessages{
		TemplateParam: string(cj),
		PhoneNumber:   "+65xxxxxxxx",
	})

	req := &sms.SmsBatchRequest{
		SmsAccount: "smsAccount",
		From:       "BytePlus",
		TemplateID: "ST_xxx",
		Tag:        "tag",
		Messages:   messages,
	}
	result, statusCode, err := i18nInstance.BatchSend(req)
	t.Logf("result = %+v\n", result)
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSmsConversion(t *testing.T) {
	client := sms.NewInstanceI18n("ap-singapore-1")
	client.Client.SetAccessKey(testAk)
	client.Client.SetSecretKey(testSk)
	req := &sms.ConversionRequest{
		MessageIDs: []string{"MessageID"},
	}
	result, statusCode, err := client.Conversion(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSmsSendVerifyCode(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)

	req := &sms.SmsVerifyCodeRequest{
		SmsAccount:  "smsAccount",
		From:        "BytePlus",
		TemplateID:  "ST_xxx",
		PhoneNumber: "+65xxxxxxxx",
		Scene:       "myscene",
		ExpireTime:  1800,
		TryCount:    3,
		CodeType:    6,
		Tag:         "tag",
	}
	result, statusCode, err := i18nInstance.SendVerifyCode(req)
	t.Logf("result = %+v\n", result)
	if result != nil && result.ResponseMetadata.Error == nil && result.Result != nil {
		t.Logf("MessageID = %+v\n", result.Result.MessageID)
	}
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSmsCheckVerifyCode(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)

	req := &sms.CheckSmsVerifyCodeRequest{
		SmsAccount:  "smsAccount",
		PhoneNumber: "+65xxxxxxxx",
		Scene:       "myscene",
		Code:        "111",
	}
	result, statusCode, err := i18nInstance.CheckVerifyCode(req)
	/**
	Result = "0" // OTP correct
	Result = "1" // OTP error
	Result = "2" // OTP expired
	*/
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
