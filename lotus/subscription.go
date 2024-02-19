package lotus

import "time"

type CreateSubscriptionRequest struct {
	CustomerId                        string                              `json:"customer_id,omitempty"`
	PlanId                            string                              `json:"plan_id,omitempty"`
	AutoRenew                         bool                                `json:"auto_renew,omitempty"`
	IsNew                             bool                                `json:"is_new,omitempty"`
	StartDate                         time.Time                           `json:"start_date,omitempty"`
	EndDate                           *time.Time                          `json:"end_date,omitempty"`
	ComponentFixedChargesInitialUnits []*ComponentFixedChargesInitialUnit `json:"component_fixed_charges_initial_units,omitempty"`
	SubscriptionFilters               []*SubscriptionFilter               `json:"subscription_filters,omitempty"`
	Metadata                          map[string]interface{}              `json:"metadata,omitempty"`
}

type ListSubscriptionsRequest struct {
	CustomerId          string                `json:"customer_id,omitempty"`
	PlanId              *string               `json:"plan_id,omitempty"`
	RangeStart          *time.Time            `json:"range_start,omitempty"`
	RangeEnd            *time.Time            `json:"range_end,omitempty"`
	Status              []string              `json:"status,omitempty"`
	SubscriptionFilters []*SubscriptionFilter `json:"subscription_filters,omitempty"`
}

func (r ListSubscriptionsRequest) q() map[string]string {
	m := make(map[string]string)
	m["customer_id"] = r.CustomerId
	if r.PlanId != nil {
		m["plan_id"] = *r.PlanId
	}
	if r.RangeStart != nil {
		m["range_start"] = r.RangeStart.Format(time.RFC3339)
	}
	if r.RangeEnd != nil {
		m["range_end"] = r.RangeEnd.Format(time.RFC3339)
	}
	return m
}

type ListSubscriptionsResponse []*Subscription

type UpdateSubscriptionRequest struct {
	SubscriptionId   string                 `json:"-"`
	EndDate          *time.Time             `json:"end_date,omitempty"`
	TurnOffAutoRenew bool                   `json:"turn_off_auto_renew,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
}

type SwitchSubscriptionPlanRequest struct {
	SubscriptionId                    string                              `json:"-"`
	SwitchPlanId                      string                              `json:"switch_plan_id,omitempty"`
	InvoicingBehavior                 InvoicingBehavior                   `json:"invoicing_behavior,omitempty"`
	UsageBehavior                     SwitchPlanUsageBehavior             `json:"usage_behavior,omitempty"`
	AutoRenew                         bool                                `json:"auto_renew,omitempty"` // TODO
	ComponentFixedChargesInitialUnits []*ComponentFixedChargesInitialUnit `json:"component_fixed_charges_initial_units,omitempty"`
	Metadata                          map[string]interface{}              `json:"metadata,omitempty"` // TODO
}

type CancelSubscriptionRequest struct {
	SubscriptionId    string                      `json:"-"`
	InvoicingBehavior InvoicingBehavior           `json:"invoicing_behavior,omitempty"`
	UsageBehavior     CancellationUsageBehavior   `json:"usage_behavior,omitempty"`
	FlatFeeBehavior   CancellationFlatFeeBehavior `json:"flat_fee_behavior,omitempty"`
}

type AttachAddonRequest struct {
	SubscriptionId string `json:"-"`
	AddonId        string `json:"addon_id,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
}

type CancelAddonRequest struct {
	SubscriptionId    string                      `json:"-"`
	AddonId           string                      `json:"-"`
	InvoicingBehavior InvoicingBehavior           `json:"invoicing_behavior,omitempty"`
	UsageBehavior     CancellationUsageBehavior   `json:"usage_behavior,omitempty"`
	FlatFeeBehavior   CancellationFlatFeeBehavior `json:"flat_fee_behavior,omitempty"`
}

type Subscription struct {
	SubscriptionId                    string                              `json:"subscription_id,omitempty"`
	Customer                          *CustomerMeta                       `json:"customer,omitempty"`
	BillingPlan                       *PlanMeta                           `json:"billing_plan,omitempty"`
	AutoRenew                         bool                                `json:"auto_renew,omitempty"`
	IsNew                             bool                                `json:"is_new,omitempty"`
	StartDate                         time.Time                           `json:"start_date,omitempty"`
	EndDate                           *time.Time                          `json:"end_date,omitempty"`
	ComponentFixedChargesInitialUnits []*ComponentFixedChargesInitialUnit `json:"component_fixed_charges_initial_units,omitempty"`
	SubscriptionFilters               []*SubscriptionFilter               `json:"subscription_filters,omitempty"`
	Metadata                          map[string]interface{}              `json:"metadata,omitempty"`
}

type ComponentFixedChargesInitialUnit struct {
	MetricId string `json:"metric_id,omitempty"`
	Units    int    `json:"units,omitempty"`
}

type SubscriptionFilter struct {
	PropertyName string `json:"property_name,omitempty"`
	Value        string `json:"value,omitempty"`
}

type SubscriptionMeta struct {
	StartDate           time.Time             `json:"start_date,omitempty"`
	EndDate             time.Time             `json:"end_date,omitempty"`
	Plan                *PlanMeta             `json:"plan,omitempty"`
	SubscriptionFilters []*SubscriptionFilter `json:"subscription_filters,omitempty"`
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
