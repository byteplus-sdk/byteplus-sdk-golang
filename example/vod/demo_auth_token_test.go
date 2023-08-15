package vod

import (
	"fmt"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"
	"testing"
)

func TestVod_GetSubtitleAuthToken(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})
	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")

	// Media Info
	query := &request.VodGetSubtitleInfoListRequest{
		Vid: "your search vid",
	}
	expireSeconds := 100 // expire time duration: (s)
	token, err := instance.GetSubtitleAuthToken(query, expireSeconds)
	fmt.Println("token ===> ", token)
	fmt.Println("err ===> ", err)
}
