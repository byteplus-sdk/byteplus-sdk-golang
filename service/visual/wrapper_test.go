package visual

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	testAk = ""
	testSk = "=="
)

func TestVisual_CVProcess(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// use example
	reqBody := map[string]interface{}{
		"req_key":   "realman_avatar_object_detection_cv",
		"image_url": "",
	}

	resp, status, err := DefaultInstance.CVProcess(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVSubmitTask(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// use example
	reqBody := map[string]interface{}{
		"req_key":   "realman_avatar_picture_omni15_cv",
		"image_url": "",
		"audio_url": "",
	}

	resp, status, err := DefaultInstance.CVSubmitTask(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVGetResult(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// use example
	reqBody := map[string]interface{}{
		"req_key": "realman_avatar_picture_omni15_cv",
		"task_id": "",
	}

	resp, status, err := DefaultInstance.CVGetResult(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVCancelTask(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// use example
	reqBody := map[string]interface{}{
		"req_key": "realman_avatar_picture_omni15_cv",
		"task_id": "",
	}

	resp, status, err := DefaultInstance.CVCancelTask(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVSync2AsyncSubmitTask(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// use example
	reqBody := map[string]interface{}{
		"req_key": "",
		// "binary_data_base64": []string{""},
		"image_urls":      []string{""},
		"positive_prompt": "beautiful",
	}

	resp, status, err := DefaultInstance.CVSync2AsyncSubmitTask(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_CVSync2AsyncGetResult(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	// use example
	reqBody := map[string]interface{}{
		"req_key": "",
		"task_id": "",
	}

	resp, status, err := DefaultInstance.CVSync2AsyncGetResult(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
