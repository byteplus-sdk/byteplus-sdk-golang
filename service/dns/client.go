// Code Generated by gadget/xsdk, DO NOT EDIT

package dns

import (
	"context"
	"net/http"
)

type Client struct {
	caller *BytePlusCaller
}

type Caller interface {
	Do(*http.Request) (*http.Response, error)
}

func NewClient(c *BytePlusCaller) *Client {
	return &Client{caller: c}
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return c.caller.Do(req)
}

func (c *Client) SetAccessKey(ak string) {
	c.caller.SetAccessKey(ak)
}

func (c *Client) SetSecretKey(sk string) {
	c.caller.SetSecretKey(sk)
}
