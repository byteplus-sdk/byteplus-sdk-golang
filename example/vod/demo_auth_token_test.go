package vod

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

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

func TestVod_GetUploadAuthWithExpiredTime(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetAccessKey("")
	instance.SetSecretKey("")

	// 默认过期时间为1小时
	ret, _ := instance.GetUploadAuth()
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))

	ret2, _ := instance.GetUploadAuthWithExpiredTime(3 * time.Hour)
	b2, _ := json.Marshal(ret2)
	fmt.Println(string(b2))
}
func TestVod_GetSha1HlsDrmAuthToken(t *testing.T) {
	instance := vod.NewInstance()
	//instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})
	// or set ak and ak as follow
	instance.SetAccessKey("")
	instance.SetSecretKey("")
	expireDuration := int64(6000000) //change to your expire duration (s), no default duration
	hlsToken, _ := instance.CreateSha1HlsDrmAuthToken(expireDuration)
	fmt.Println(hlsToken)
}

func TestVod_GetPrivateDrmAuthToken(t *testing.T) {
	instance := vod.NewInstance()

	//instance.SetCredential(base.Credentials{
	// AccessKeyID:     "your ak",
	// SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	instance.SetAccessKey("your ak")
	instance.SetSecretKey("your sk")

	query := &request.VodGetPrivateDrmPlayAuthRequest{
		Vid:         "your vid",
		DrmType:     "your drmType",
		PlayAuthIds: "your playAuthIds",
		UnionInfo:   "your unionInfo",
	}
	tokenExpireTime := 6000000 // Token Expire Duration（s）
	newToken, _ := instance.GetPrivateDrmAuthToken(query, tokenExpireTime)
	fmt.Println(newToken)
}
