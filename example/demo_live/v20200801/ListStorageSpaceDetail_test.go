package live_v20200801_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/live/v20200801"
)

func Test_ListStorageSpaceDetail(t *testing.T) {
	instance := live_v20200801.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &live_v20200801.ListStorageSpaceDetailBody{}

	resp, err := instance.ListStorageSpaceDetail(param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
