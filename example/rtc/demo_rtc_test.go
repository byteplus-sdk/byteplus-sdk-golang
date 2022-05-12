package main

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/rtc"
)

const (
	appId   = ""
	Your_AK = ""
	Your_SK = ""
)

var rtcClient *rtc.RTC

func TestMain(t *testing.M) {
	//init rtc service
	rtcClient = rtc.NewInstance()
	rtcClient.Client.SetAccessKey(Your_AK)
	rtcClient.Client.SetSecretKey(Your_SK)
	//Please replace it with the address in the official website document
	rtcClient.Client.SetHost(rtc.ServiceHost)

	t.Run()
}

func TestListRooms(t *testing.T) {
	query := url.Values{}
	query.Set("AppId", appId)
	//query.Set("RoomId", "zdl_room_20210818")
	res, _, err := rtc.ListRooms(rtcClient, query)
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

func TestListIndicators(t *testing.T) {
	req := rtc.ListIndicatorsRequest{
		AppId:     appId,
		StartTime: "2021-08-17T00:00:00+08:00",
		EndTime:   "2021-08-18T00:00:00+08:00",
		Indicator: "NetworkTransDelay",
	}
	//req.OS = "android"
	//req.Network = "wifi"
	res, _, err := rtc.ListIndicators(rtcClient, &req)
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
