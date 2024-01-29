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
