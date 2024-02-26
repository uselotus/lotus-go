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

// Lotus API
// ---------

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

// CreateCustomer creates a customer
// See: https://docs.uselotus.io/api-reference/customers/create-customer
func (c *Client) CreateCustomer(req CreateCustomerRequest) (resp *Customer, err error) {
	resp = new(Customer)
	err = c.post("/api/customers/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetCustomer gets the customer specified by customer id
// See: https://docs.uselotus.io/api-reference/customers/retrieve-customer
func (c *Client) GetCustomer(req GetCustomerRequest) (resp *Customer, err error) {
	resp = new(Customer)
	err = c.get("/api/customers/"+req.CustomerId, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListCustomers retrieves an array of customer objects
// See: https://docs.uselotus.io/api-reference/customers/list-customers
func (c *Client) ListCustomers() (resp ListCustomersResponse, err error) {
	resp = make(ListCustomersResponse, 0)
	err = c.get("/api/customers/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFeatureAccess checks customer feature access
// See: https://docs.uselotus.io/api-reference/access/feature-access
func (c *Client) GetFeatureAccess(req GetFeatureAccessRequest) (resp *FeatureEntitlement, err error) {
	resp = new(FeatureEntitlement)
	err = c.get("/api/feature_access/", req.q(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetMetricAccess checks customer metric access
// See https://docs.uselotus.io/api-reference/access/metric-access
func (c *Client) GetMetricAccess(req GetMetricAccessRequest) (resp *MetricEntitlement, err error) {
	resp = new(MetricEntitlement)
	err = c.get("/api/metric_access/", req.q(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TrackEvents ingests events into Lotus
// See: https://docs.uselotus.io/api-reference/events/tracking
func (c *Client) TrackEvents(req TrackEventsRequest) (resp *TrackEventsResponse, err error) {
	resp = new(TrackEventsResponse)
	err = c.post("/api/track/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VerifyEventIngestion verifies if events with specific idempotency events have been ingested into Lotus
// See: https://docs.uselotus.io/api-reference/events/verify-ingestion
func (c *Client) VerifyEventIngestion(req VerifyEventIngestionRequest) (resp *VerifyEventIngestionResponse, err error) {
	resp = new(VerifyEventIngestionResponse)
	err = c.post("/api/verify_idems_received/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListPlans retrieves an array of plan objects
// See: https://docs.uselotus.io/api-reference/plans/list-plans
func (c *Client) ListPlans(req ListPlansRequest) (resp ListPlansResponse, err error) {
	resp = make(ListPlansResponse, 0)
	err = c.get("/api/plans/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetPlan gets the plan specified by plan id
// See: https://docs.uselotus.io/api-reference/plans/retrieve-plan
func (c *Client) GetPlan(req GetPlanRequest) (resp *Plan, err error) {
	resp = new(Plan)
	err = c.get("/api/plans/"+req.PlanId, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateSubscription adds a plan to a customer’s subscription
// See: https://docs.uselotus.io/api-reference/subscriptions/create-subscription
func (c *Client) CreateSubscription(req CreateSubscriptionRequest) (resp *Subscription, err error) {
	resp = new(Subscription)
	err = c.post("/api/subscriptions/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListSubscriptions retrieves an array of subscription objects
// See: https://docs.uselotus.io/api-reference/subscriptions/list-subscriptions
func (c *Client) ListSubscriptions(req ListSubscriptionsRequest) (resp ListSubscriptionsResponse, err error) {
	resp = make(ListSubscriptionsResponse, 0)
	err = c.get("/api/subscriptions/", req.q(), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateSubscription cancels auto-renew or changes the end date of a subscription
// See: https://docs.uselotus.io/api-reference/subscriptions/change-subscription
func (c *Client) UpdateSubscription(req UpdateSubscriptionRequest) (resp *Subscription, err error) {
	resp = new(Subscription)
	err = c.post("/api/subscriptions/"+req.SubscriptionId+"/update/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SwitchSubscriptionPlan upgrades or downgrades the subscription’s plan
// See: https://docs.uselotus.io/api-reference/subscriptions/switch-subscription-plan
func (c *Client) SwitchSubscriptionPlan(req SwitchSubscriptionPlanRequest) (resp *Subscription, err error) {
	resp = new(Subscription)
	err = c.post("/api/subscriptions/"+req.SubscriptionId+"/switch_plan/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelSubscription cancels a subscription
// See: https://docs.uselotus.io/api-reference/subscriptions/cancel-subscription
func (c *Client) CancelSubscription(req CancelSubscriptionRequest) (resp *Subscription, err error) {
	resp = new(Subscription)
	err = c.post("/api/subscriptions/"+req.SubscriptionId+"/cancel/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AttachAddon adds an addon to a customer’s subscription
// See: https://docs.uselotus.io/api-reference/subscriptions/attach-addon
func (c *Client) AttachAddon(req AttachAddonRequest) (resp *Subscription, err error) {
	resp = new(Subscription)
	err = c.post("/api/subscriptions/"+req.SubscriptionId+"/addons/attach/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelAddon cancels an addon in the customer’s subscription
// See: https://docs.uselotus.io/api-reference/subscriptions/attach-addon
func (c *Client) CancelAddon(req CancelAddonRequest) (resp *Subscription, err error) {
	resp = new(Subscription)
	err = c.post("/api/subscriptions/"+req.SubscriptionId+"/addons/"+req.AddonId+"/cancel/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
