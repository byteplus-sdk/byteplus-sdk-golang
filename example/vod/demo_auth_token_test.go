package vod

import (
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"
)

func TestVod_GetSubtitleAuthToken(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.byteplus/config
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

func TestVod_GetPlayAuthToken(t *testing.T) {
	vid := "your vid"
	tokenExpireTime := 600 // Token Expire Duration（s）
	instance := vod.NewInstance()

	//instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	//instance.SetAccessKey("")
	//instance.SetSecretKey("")

	query := &request.VodGetPlayInfoRequest{
		Vid:             vid,
		Format:          "mp4",
		Definition:      "360p",
		FileType:        "video",
		LogoType:        "",
		Ssl:             "1",
		NeedThumbs:      "0",
		NeedBarrageMask: "0",
		CdnType:         "0",
	}
	newToken, _ := instance.GetPlayAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}
