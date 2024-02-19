package lotus

import (
	"github.com/shopspring/decimal"
)

type GetFeatureAccessRequest struct {
	CustomerId string `json:"customer_id,omitempty"`
	FeatureId  string `json:"feature_id,omitempty"`
}

func (r GetFeatureAccessRequest) q() map[string]string {
	return map[string]string{
		"customer_id": r.CustomerId,
		"feature_id":  r.FeatureId,
	}
}

type GetMetricAccessRequest struct {
	CustomerId string `json:"customer_id,omitempty"`
	MetricId   string `json:"metric_id,omitempty"`
}

func (r GetMetricAccessRequest) q() map[string]string {
	return map[string]string{
		"customer_id": r.CustomerId,
		"metric_id":   r.MetricId,
	}
}

type FeatureEntitlement struct {
	Access                bool                         `json:"access"`
	Customer              *CustomerMeta                `json:"customer,omitempty"`
	Feature               *FeatureMeta                 `json:"feature,omitempty"`
	AccessPerSubscription []*FeatureSubscriptionAccess `json:"access_per_subscription,omitempty"`
}

type FeatureSubscriptionAccess struct {
	Access       bool              `json:"access"`
	Subscription *SubscriptionMeta `json:"subscription,omitempty"`
}

type MetricEntitlement struct {
	Access                bool                        `json:"access"`
	Customer              *CustomerMeta               `json:"customer,omitempty"`
	Metric                *MetricMeta                 `json:"metric,omitempty"`
	AccessPerSubscription []*MetricSubscriptionAccess `json:"access_per_subscription,omitempty"`
}

type MetricSubscriptionAccess struct {
	MetricFreeLimit  decimal.Decimal   `json:"metric_free_limit"`
	MetricTotalLimit decimal.Decimal   `json:"metric_total_limit"`
	MetricUsage      decimal.Decimal   `json:"metric_usage"`
	Subscription     *SubscriptionMeta `json:"subscription,omitempty"`
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
