package dns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/pkg/errors"
)

func NewDefaultServiceInfo() *base.Client {
	return base.NewClient(ServiceInfo, nil)
}

type BytePlusCaller struct {
	BytePlus *base.Client
	Query    map[string]string
	Headers  map[string]string
}

// GetServiceInfo interface
func (p *BytePlusCaller) GetServiceInfo() *base.ServiceInfo {
	return p.BytePlus.ServiceInfo
}

func (p *BytePlusCaller) SetService(service string) {
	p.BytePlus.ServiceInfo.Credentials.Service = service
}

func (p *BytePlusCaller) SetRegion(region string) {
	p.BytePlus.ServiceInfo.Credentials.Region = region
}

func (p *BytePlusCaller) SetHost(host string) {
	p.BytePlus.ServiceInfo.Host = host
}

func (p *BytePlusCaller) SetSchema(schema string) {
	p.BytePlus.ServiceInfo.Scheme = schema
}

func (p *BytePlusCaller) SetAccessKey(ak string) {
	p.BytePlus.SetAccessKey(ak)
}

func (p *BytePlusCaller) SetSecretKey(sk string) {
	p.BytePlus.SetSecretKey(sk)
}

func (p *BytePlusCaller) AddQuery(k, v string) {
	p.Query[k] = v
}

func (p *BytePlusCaller) SetHeader(k, v string) {
	p.Headers[k] = v
}

func NewBytePlusCaller() *BytePlusCaller {
	return &BytePlusCaller{
		BytePlus: NewDefaultServiceInfo(),
		Query:    make(map[string]string, 10),
		Headers:  make(map[string]string, 10),
	}
}

func InitDNSBytePlusClient() *Client {
	vc := NewBytePlusCaller()
	vc.SetService(ServiceName)
	vc.SetRegion(DefaultRegion)
	return NewClient(vc)
}

func (c *BytePlusCaller) Do(r *http.Request) (*http.Response, error) {
	r.URL.Host = c.BytePlus.ServiceInfo.Host
	r.URL.Scheme = c.BytePlus.ServiceInfo.Scheme
	r.Host = c.BytePlus.ServiceInfo.Host

	r.Header.Add("Content-Type", "application/json")
	if len(c.Headers) > 0 {
		for k, v := range c.Headers {
			// do not cover request header
			if len(r.Header.Get(k)) > 0 {
				continue
			}
			r.Header.Add(k, v)
		}
	}

	q := r.URL.Query()
	q.Add("Version", ServiceVersion)
	for k, v := range c.Query {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()

	r = c.BytePlus.ServiceInfo.Credentials.Sign(r)

	ctx, cancel := context.WithTimeout(r.Context(), c.BytePlus.ServiceInfo.Timeout)
	defer cancel()
	r = r.WithContext(ctx)

	resp, err := c.BytePlus.Client.Do(r)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	var payload TopResponse
	err = json.NewDecoder(resp.Body).Decode(&payload)
	resp.Body.Close()

	newResp := &http.Response{}

	if err != nil {
		return newResp, errors.WithStack(err)
	}

	if payload.ResponseMetadata.Error.Code != "" {
		return newResp, NewTOPError(&payload.ResponseMetadata)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return newResp, errors.Wrap(errors.New("http error"), strconv.Itoa(resp.StatusCode))
	}

	str, err := json.Marshal(payload.Result)
	if err != nil {
		return newResp, errors.WithStack(err)
	}

	newResp = &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(str))}

	return newResp, errors.WithStack(err)
}

type TOPError struct {
	Code      string `form:"Code" json:"Code"`
	CodeN     int64  `form:"CodeN" json:"CodeN"`
	Message   string `form:"Message" json:"Message"`
	RequestID string `form:"RequestId" json:"RequestId"`
}

func NewTOPError(respMeta *TopRespMeta) *TOPError {
	return &TOPError{
		Code:      respMeta.Error.Code,
		CodeN:     respMeta.Error.CodeN,
		Message:   respMeta.Error.Message,
		RequestID: respMeta.RequestID,
	}
}

func (err *TOPError) GetCode() int64 {
	return err.CodeN
}

func (err *TOPError) Error() string {
	return fmt.Sprintf(
		"TOP DNS response failed with code %d: %s, %s, requestID: %s",
		err.CodeN, err.Code, err.Message, err.RequestID,
	)
}

type TopResponse struct {
	ResponseMetadata TopRespMeta     `form:"ResponseMetadata" json:"ResponseMetadata"`
	Result           json.RawMessage `form:"Result" json:"Result"`
}

type TopRespMeta struct {
	Action    string       `form:"Action" json:"Action"`
	Error     TopRespError `form:"Error" json:"Error"`
	Region    string       `form:"Region" json:"Region"`
	RequestID string       `form:"RequestId" json:"RequestId"`
	Service   string       `form:"Service" json:"Service"`
	Version   string       `form:"Version" json:"Version"`
}

type TopRespError struct {
	Code    string `form:"Code" json:"Code"`
	CodeN   int64  `form:"CodeN" json:"CodeN"`
	Message string `form:"Message" json:"Message"`
}
