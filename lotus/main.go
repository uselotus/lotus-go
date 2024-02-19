package lotus

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	host   string
	key    string
	client *resty.Client
}

// NewClient initializes a Lotus client with keep-alive & retry enabled
// See https://docs.uselotus.io/api-reference/api-overview#how-to-use-your-api-key for how to create an API key
func NewClient(host, key string) *Client {
	tran := &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     60 * time.Second,
	}
	client := resty.New().
		SetTransport(tran).
		SetTimeout(time.Second * 5).
		SetRetryCount(2).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return r.StatusCode() >= 500
		})
	return &Client{
		host:   host,
		key:    key,
		client: client,
	}
}

// WithDebug enables request debugging
func (c *Client) WithDebug(d bool) *Client {
	c.client.SetDebug(d)
	return c
}

// WithTimeout sets HTTP client timeout
func (c *Client) WithTimeout(t time.Duration) *Client {
	if t.Milliseconds() > 0 {
		c.client.SetTimeout(t)
	}
	return c
}

// WithTransport sets HTTP client transport
func (c *Client) WithTransport(tran *http.Transport) *Client {
	if tran != nil {
		c.client.SetTransport(tran)
	}
	return c
}

// baseUrl returns base url
func (c *Client) baseUrl() string {
	if !strings.HasPrefix(c.host, "http") {
		return "https://" + c.host
	}
	return c.host
}

// auth returns authorization key
func (c *Client) auth() string {
	return c.key
}

// post sends a POST request with body req, response will be stored in res
func (c *Client) post(url string, q map[string]string, req interface{}, res interface{}) error {
	resp, err := c.client.R().
		SetHeader("X-API-Key", c.auth()).
		SetQueryParams(q).
		SetBody(req).
		SetResult(res).
		Post(c.baseUrl() + url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		var e Error
		if json.Unmarshal(resp.Body(), &e) == nil {
			return e
		}
		return respError(resp)
	}
	return nil
}

// get sends a GET request with query q, response will be stored in res
func (c *Client) get(url string, q map[string]string, res interface{}) error {
	resp, err := c.client.R().
		SetHeader("X-API-Key", c.auth()).
		SetQueryParams(q).
		SetResult(res).
		Get(c.baseUrl() + url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		var e Error
		if json.Unmarshal(resp.Body(), &e) == nil {
			return e
		}
		return respError(resp)
	}
	return nil
}

// respError returns corresponding error
func respError(resp *resty.Response) error {
	if resp.String() == "" {
		return fmt.Errorf("server status code: %v, nil body", resp.StatusCode())
	}
	return fmt.Errorf("server status code: %v, %v", resp.StatusCode(), resp.String())
}

type PingResponse struct {
	OrganizationId string `json:"organization_id"`
}

// Ping pings Lotus API to check if the API key is valid
// See: https://docs.uselotus.io/api-reference/api-overview
func (c *Client) Ping() (resp *PingResponse, err error) {
	resp = new(PingResponse)
	err = c.get("/api/ping/", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
