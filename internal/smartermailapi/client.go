package smartermailapi

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type Client struct {
	host         string
	username     string
	password     string
	refreshToken string
	accessToken  string
}

func NewClient(host string, username string, password string) *Client {
	return &Client{
		host:     host,
		username: username,
		password: password,
	}
}

func NewClientFromRefreshToken(host string, refreshToken string) *Client {
	return &Client{
		refreshToken: refreshToken,
	}
}

func (c *Client) defaultPostRequest(path string) *fasthttp.Request {
	req := c.defaultGetRequest(path)
	req.Header.SetMethod("POST")

	return req
}

func (c *Client) defaultGetRequest(path string) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("%s/%s", c.host, path))
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")

	return req
}

func (c *Client) includeAccessToken(req *fasthttp.Request) *fasthttp.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return req
}
