package visual

import (
	"encoding/json"
	"errors"
	"net/url"
)

func (p *Visual) commonHandler(api string, form url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) commonJsonHandler(api string, jsonStr string, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Json(api, nil, jsonStr)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) commonQueryHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

// CVProcess comman sync interface
// bodyMap: Please refer to the request body in interface documentation
func (p *Visual) CVProcess(bodyMap interface{}) (map[string]interface{}, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := make(map[string]interface{})
	statusCode, err := p.commonJsonHandler("CVProcess", string(reqByte), &resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// CVSubmitTask Common async submit task interface
// bodyMap: Please refer to the request body in interface documentation
func (p *Visual) CVSubmitTask(bodyMap interface{}) (map[string]interface{}, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := make(map[string]interface{})
	statusCode, err := p.commonJsonHandler("CVSubmitTask", string(reqByte), &resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// CVGetResult common async get result interface
// bodyMap: Please refer to the request body in interface documentation
func (p *Visual) CVGetResult(bodyMap interface{}) (map[string]interface{}, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := make(map[string]interface{})
	statusCode, err := p.commonJsonHandler("CVGetResult", string(reqByte), &resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// CVCancelTask common async cancel task interface
// bodyMap: Please refer to the request body in interface documentation
func (p *Visual) CVCancelTask(bodyMap interface{}) (map[string]interface{}, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := make(map[string]interface{})
	statusCode, err := p.commonJsonHandler("CVCancelTask", string(reqByte), &resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// CVSync2AsyncSubmitTask common sync to async submit task interface
// bodyMap: Please refer to the request body in interface documentation
func (p *Visual) CVSync2AsyncSubmitTask(bodyMap interface{}) (map[string]interface{}, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := make(map[string]interface{})
	statusCode, err := p.commonJsonHandler("CVSync2AsyncSubmitTask", string(reqByte), &resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// CVSync2AsyncGetResult common sync to async get result interface
// bodyMap: Please refer to the request body in interface documentation
func (p *Visual) CVSync2AsyncGetResult(bodyMap interface{}) (map[string]interface{}, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := make(map[string]interface{})
	statusCode, err := p.commonJsonHandler("CVSync2AsyncGetResult", string(reqByte), &resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
