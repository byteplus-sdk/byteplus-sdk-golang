package upload

import (
	"encoding/json"
	"fmt"
	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/business"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/upload/functions"
	"testing"
)

func TestVod_CommitUploadInfo(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	space := "your space name"
	session := "apply return session"

	funcs := make([]business.VodUploadFunction, 0)

	snapShotFunc := functions.SnapshotFunc(business.VodUploadFunctionInput{SnapshotTime: 1.3})
	getMetaFunc := functions.GetMetaFunc()

	funcs = append(funcs, snapShotFunc)
	funcs = append(funcs, getMetaFunc)

	fbts, err := json.Marshal(funcs)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("\n%s", fbts)

	commitRequest := &request.VodCommitUploadInfoRequest{
		SpaceName:    space,
		SessionKey:   session,
		CallbackArgs: "my callback",
		Functions:    string(fbts),
	}
	resp, _, err := instance.CommitUploadInfo(commitRequest)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	if resp.GetResponseMetadata().GetError() != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}

	fmt.Println(resp.GetResult().GetData().GetVid())
	bts, _ := json.Marshal(resp)
	fmt.Printf("\nresp = %s", bts)
}
