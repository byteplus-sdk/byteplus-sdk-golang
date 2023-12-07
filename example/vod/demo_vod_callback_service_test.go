// Code generated by protoc-gen-byteplus-sdk
// source: VodCallbackService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"
)

func Test_AddCallbackSubscription(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodAddCallbackSubscriptionRequest{
		SpaceName:   "your SpaceName",
		Url:         "your Url",
		ContentType: "your ContentType",
	}

	resp, status, err := instance.AddCallbackSubscription(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_SetCallbackEvent(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodSetCallbackEventRequest{
		SpaceName:   "your SpaceName",
		Events:      "your Events",
		AuthEnabled: "your AuthEnabled",
		PrivateKey:  "your PrivateKey",
	}

	resp, status, err := instance.SetCallbackEvent(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}