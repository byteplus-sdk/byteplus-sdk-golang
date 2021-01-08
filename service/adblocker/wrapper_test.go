package adblocker

import (
	"fmt"
	"testing"
)

const (
	region = "region"
	testAk = "AKLTMWRiNDFiYWVmZjkyNGI1ZmE3YmIxNTQxOTdjODU1ZTE"
	testSk = "TUdFeU0ySmtZek0zWWpZek5EYzJNVGxrWlRkbU56TmhaV1JoTjJFeU1EQQ=="
)

func TestAdBlock(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	res, err := DefaultInstance.AdBlock(&AdBlockRequest{
		AppID:      3332,
		Service:    "chat",
		Parameters: "{\"uid\":123411, \"operate_time\":1609818934, \"chat_text\":\"a😊\"}",
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}
