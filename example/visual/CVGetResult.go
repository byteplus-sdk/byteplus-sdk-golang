package main

import (
	"encoding/json"
	"fmt"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/visual"
)

func main() {
	testAk := ""
	testSk := ""

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//request body. Please refer to the request body in interface documentation
	reqBody := map[string]interface{}{
		"req_key": "",
		"task_id": "",
	}

	resp, status, err := visual.DefaultInstance.CVGetResult(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))

	// example：transfer resp to response structure
	// if status != 200 {
	// 	fmt.Println("request err, response is:", string(b))
	// 	return
	// }
	// respData := new(model.ResponseStructxxx)
	// if err = json.Unmarshal(b, respData); err != nil {
	// 	fmt.Println("unmarshal err:", err)
	// 	return
	// }
	// fmt.Println(respData.Data.Status)
}
