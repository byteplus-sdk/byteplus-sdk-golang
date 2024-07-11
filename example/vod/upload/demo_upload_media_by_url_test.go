package upload

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/business"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"

	"testing"
)

func TestVod_UploadMediaByUrl(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	urlSets := make([]*business.VodUrlUploadURLSet, 0)
	urlSet := &business.VodUrlUploadURLSet{
		SourceUrl:     "your source url",    // 源视频 URL
		CallbackArgs:  "your callback args", // 透传信息，业务希望透传的字段可以写入，返回和回调中会返回此字段，非必须字段
		FileName:      "",                   // 设置文件名，无格式长度限制，用户可自定义,目前文件名不支持空格、+ 字符,如果要使用此字段，请联系技术支持配置白名单，非必须字段
		FileExtension: ".mp4",               // 设置文件后缀，以 . 开头，不超过8位，非必须字段
	}
	urlSets = append(urlSets, urlSet)

	query := &request.VodUrlUploadRequest{
		SpaceName: "your SpaceName",
		URLSets:   urlSets,
	}
	resp, status, err := instance.UploadMediaByUrl(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
