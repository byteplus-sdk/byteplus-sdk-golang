package demo_sms

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/sms"
)

/*
SendVerifyCode cooperates with CheckVerifyCode, to provide verify capability.
*/
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
