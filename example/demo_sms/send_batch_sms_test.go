package demo_sms

import (
	"encoding/json"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/sms"
)

type sendBatchSmsTemplateParam struct {
	Code string `json:"code"`
}

/*
send content with different parameters to one/multiple phoneNumbers
*/
func TestSmsBatchSendI18n(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)

	c := sendBatchSmsTemplateParam{Code: "111"}
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
