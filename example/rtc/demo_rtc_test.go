// Description: Demonstrates how to use the OpenAPI authentication SDK
// When using, please refer to the package github.com/byteplus-sdk/byteplus-sdk-golang/service/rtc or copy it to your own project to append the required API
//
// The service/rtc directory contains three files:
// 1. config.go defines API properties, introduces signature packages and initializes services
// 2. model.go stores the definition of API request parameters and return parameters
// 3. The method body of the wrappper.go API is mainly divided into two categories: GET and POST
package main

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/rtc"
)

const (
	// ak/sk references: https://docs.byteplus.com/byteplus-rtc/reference/69828
	appId  = ""
	testAk = ""
	testSk = ""
)

var (
	roomId     = "Your_roomId"
	busineseId = "Your_BusinessId"
	taskId     = appId + "_" + roomId + "_xxx"

	s *rtc.RTC
)

func TestMain(t *testing.M) {
	// init rtc service, Just initialize once
	s = rtc.NewInstance()
	s.Client.SetAccessKey(testAk)
	s.Client.SetSecretKey(testSk)
	s.Client.SetHost(rtc.ServiceHost)

	t.Run()
}

func TestGetRecordTask(t *testing.T) {
	query := url.Values{}
	query.Set("AppId", appId)
	query.Set("RoomId", roomId)
	query.Set("TaskId", taskId)

	res, _, err := rtc.GetRecordTask(s, query)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if res.ResponseMetadata.Error != nil {
		fmt.Printf("response err:%v\n", res.ResponseMetadata.Error)
		return
	}
	fmt.Printf("metaData: %+v\n", res.ResponseMetadata)
	fmt.Printf("result: %+v\n", res.Result)
}

func TestStartRecord(t *testing.T) {
	req := rtc.StartRecordRequest{
		AppId:      appId,
		BusinessId: busineseId,
		RoomId:     roomId,
		TaskId:     taskId,
		Encode: &rtc.Encode{
			VideoWidth:   1920,
			VideoHeight:  1080,
			VideoFps:     15,
			VideoBitrate: 4000,
		},
		FileFormatConfig: &rtc.FileFormatConfig{
			FileFormat: []string{"MP4"},
		},
		StorageConfig: rtc.StorageConfig{
			Type: 1,
			VodConfig: &rtc.VodConfig{
				AccountId: "Your_Volc_AccountId",
				Space:     "Your_Space",
			},
		},
	}
	res, _, err := rtc.StartRecord(s, &req)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	if res.ResponseMetadata.Error != nil {
		fmt.Printf("response err:%v\n", res.ResponseMetadata.Error)
		return
	}
	fmt.Printf("metaData: %+v\n", res.ResponseMetadata)
	fmt.Printf("result: %+v\n", res.Result)
}
