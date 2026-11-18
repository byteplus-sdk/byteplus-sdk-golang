package demo_livesaas

import (
	"encoding/json"
	"net/url"
	"strconv"
	"testing"
	"time"

	livesaas "github.com/byteplus-sdk/byteplus-sdk-golang/service/livesaas"
)

const (
	testAk = ""
	testSk = ""
)

const (
	account_id      = 0
	accountUserName = ""
	activity_id     = ""
	activity_name   = ""
)

func TestLIVESAAS_CreateActivityAPIV2(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{
		"AccountId":       account_id,
		"AccountUserName": accountUserName,
		"ColorThemeIndex": "light",
		"ConfigVersion":   1,
		"LiveTime":        time.Now().Unix(),
		"LiveZone":        480,
		"Name":            activity_name,
	}
	// Serialize the request parameters
	body, _ := json.Marshal(bodyMap)
	// Print the request parameters
	t.Logf(string(body))
	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.CreateActivityAPIV2(nil, string(body))

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_GetTemporaryLoginTokenAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)

	query := url.Values{}
	query.Set("ActivityId", activity_id)

	// Print the request parameters
	t.Logf("ActivityId = %+v", query.Get("ActivityId"))

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.GetTemporaryLoginTokenAPI(query)

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_GetSDKTokenAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	activityId, _ := strconv.ParseInt(activity_id, 10, 64)
	bodyMap := map[string]interface{}{
		"ActivityId": activityId,
		"Mode":       1,
	}
	body, _ := json.Marshal(bodyMap)
	// Print the request parameters
	t.Logf(string(body))

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.GetSDKTokenAPI(nil, string(body))

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_ListActivityAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	bodyMap := map[string]interface{}{}
	body, _ := json.Marshal(bodyMap)
	// Print the request parameters
	t.Logf(string(body))

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.ListActivityAPI(nil, string(body))

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_UpdateActivityBasicConfigAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	activityId, _ := strconv.ParseInt(activity_id, 10, 64)
	bodyMap := map[string]interface{}{
		"ActivityId":      activityId,
		"Name":            activity_name,
		"ColorThemeIndex": "dark",
	}
	body, _ := json.Marshal(bodyMap)
	// Print the request parameters
	t.Logf(string(body))

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.UpdateActivityBasicConfigAPI(nil, string(body))

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_ListActivityDetailInfoAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	livesaas.DefaultInstance.Client.SetTimeout(time.Second * 10)

	// Print the request parameters
	query := url.Values{}
	query.Set("PageNumber", "1")
	query.Set("PageSize", "5")

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.ListActivityDetailInfoAPI(query)

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_GetBusinessAccountInfoAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	livesaas.DefaultInstance.Client.SetTimeout(time.Second * 10)

	// Print the request parameters
	query := url.Values{}
	query.Set("ActivityId", activity_id)

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.GetBusinessAccountInfoAPI(query)

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_ListActivityFeedInfosAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	livesaas.DefaultInstance.Client.SetTimeout(time.Second * 10)

	// Print the request parameters
	query := url.Values{}
	query.Set("PageNumber", "1")
	query.Set("PageSize", "5")

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.ListActivityFeedInfosAPI(query)

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}

func TestLIVESAAS_GetStreamsAPI(t *testing.T) {
	livesaas.DefaultInstance.Client.SetAccessKey(testAk)
	livesaas.DefaultInstance.Client.SetSecretKey(testSk)
	livesaas.DefaultInstance.Client.SetTimeout(time.Second * 10)

	// Print the request parameters
	query := url.Values{}
	query.Set("ActivityId", activity_id)

	// Call the API operation to create a live-streaming activity with the serialized request parameters
	resp, statusCode, err := livesaas.DefaultInstance.GetStreamsAPI(query)

	// If an error occurs while calling the API operation, print the error
	if err != nil {
		t.Logf("error occur %v", err)
	}
	// Serialize the response data
	res, _ := json.Marshal(resp)
	// Print the response data
	t.Logf("statusCode = %+v  msgInfo = %+v \n", statusCode, string(res))
}
