package livesaas

import (
	"encoding/json"
	"net/url"
)

func (p *LIVESAAS) commonHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *LIVESAAS) commonHandlerJson(api string, query url.Values, resp interface{}, body string) (int, error) {
	respBody, statusCode, err := p.Client.Json(api, query, body)
	if err != nil {
		return statusCode, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *LIVESAAS) CreateActivityAPIV2(query url.Values, body string) (*CommonResponseResp, int, error) {
	resp := new(CommonResponseResp)
	statesCode, err := p.commonHandlerJson("CreateActivityAPIV2", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVESAAS) GetTemporaryLoginTokenAPI(query url.Values) (*GetTemporaryLoginTokenResponse, int, error) {
	resp := new(GetTemporaryLoginTokenResponse)
	statusCode, err := p.commonHandler("GetTemporaryLoginTokenAPI", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVESAAS) GetSDKTokenAPI(query url.Values, body string) (*GetSDKTokenResponse, int, error) {
	resp := new(GetSDKTokenResponse)
	statesCode, err := p.commonHandlerJson("GetSDKTokenAPI", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVESAAS) ListActivityAPI(query url.Values, body string) (*ListActivityAPIResponse, int, error) {
	resp := new(ListActivityAPIResponse)
	statesCode, err := p.commonHandlerJson("ListActivityAPI", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVESAAS) UpdateActivityBasicConfigAPI(query url.Values, body string) (*CommonResponseResp, int, error) {
	resp := new(CommonResponseResp)
	statesCode, err := p.commonHandlerJson("UpdateActivityBasicConfigAPI", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVESAAS) ListActivityDetailInfoAPI(query url.Values) (*ListActivityDetailInfoAPIResponse, int, error) {
	resp := new(ListActivityDetailInfoAPIResponse)
	statesCode, err := p.commonHandler("ListActivityDetailInfoAPI", query, resp)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVESAAS) GetBusinessAccountInfoAPI(query url.Values) (*GetBusinessAccountInfoAPIResponse, int, error) {
	resp := new(GetBusinessAccountInfoAPIResponse)
	statesCode, err := p.commonHandler("GetBusinessAccountInfoAPI", query, resp)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVESAAS) ListActivityFeedInfosAPI(query url.Values) (*ListActivityFeedInfosAPIResponse, int, error) {
	resp := new(ListActivityFeedInfosAPIResponse)
	statesCode, err := p.commonHandler("ListActivityFeedInfosAPI", query, resp)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVESAAS) GetStreamsAPI(query url.Values) (*GetStreamsResponse, int, error) {
	resp := new(GetStreamsResponse)
	statesCode, err := p.commonHandler("GetStreamsAPI", query, resp)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}
