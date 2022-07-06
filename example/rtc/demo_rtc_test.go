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

func TestListRoomInformation(t *testing.T) {
	query := url.Values{}
	query.Set("AppId", appId)
	query.Set("StartTime", "2022-04-22T12:00:00+08:00")
	query.Set("EndTime", "2022-05-22T12:59:00+08:00")

	res, _, err := rtc.ListRoomInformation(rtcClient, query)
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
		StartTime: "2022-05-17T00:00:00+08:00",
		EndTime:   "2022-06-18T00:00:00+08:00",
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
