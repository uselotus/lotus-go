package lotus

import (
	"github.com/shopspring/decimal"
	"time"
)

type ListPlansRequest struct {
	Duration Duration `json:"duration"`
	// TODO: add more filters
}

type ListPlansResponse []*Plan

type GetPlanRequest struct {
	PlanId string
}

type Plan struct {
	PlanId              string          `json:"plan_id,omitempty"`
	PlanName            string          `json:"plan_name,omitempty"`
	PlanDuration        string          `json:"plan_duration,omitempty"`
	PlanDescription     string          `json:"plan_description,omitempty"`
	ExternalLinks       []*ExternalLink `json:"external_links,omitempty"`
	NumVersions         int             `json:"num_versions,omitempty"`
	ActiveVersion       int             `json:"active_version,omitempty"`
	ActiveSubscriptions int             `json:"active_subscriptions,omitempty"`
	Tags                []string        `json:"tags,omitempty"`
	Versions            []*Version      `json:"versions,omitempty"`
	ParentPlan          *Plan           `json:"parent_plan,omitempty"`
	TargetCustomer      *CustomerMeta   `json:"target_customer,omitempty"`
	DisplayVersion      *Version        `json:"display_version,omitempty"`
	Status              string          `json:"status,omitempty"`
}

type PlanMeta struct {
	PlanId    string `json:"plan_id,omitempty"`
	PlanName  string `json:"plan_name,omitempty"`
	Version   int    `json:"version,omitempty"`
	VersionId string `json:"version_id,omitempty"`
}

type ExternalLink struct {
	Source         string `json:"source,omitempty"`
	ExternalPlanId string `json:"external_plan_id,omitempty"`
}

type Version struct {
	Version               int                 `json:"version,omitempty"`
	PlanName              string              `json:"plan_name,omitempty"`
	LocalizedName         string              `json:"localized_name,omitempty"`
	ActiveFrom            time.Time           `json:"active_from,omitempty"`
	ActiveTo              time.Time           `json:"active_to,omitempty"`
	CreatedOn             time.Time           `json:"created_on,omitempty"`
	Description           string              `json:"description,omitempty"`
	Currency              *Currency           `json:"currency,omitempty"`
	Features              []*FeatureMeta      `json:"features,omitempty"`
	RecurringCharges      []*RecurringCharge  `json:"recurring_charges,omitempty"`
	Components            []*PricingComponent `json:"components,omitempty"`
	PriceAdjustment       *PriceAdjustment    `json:"price_adjustment,omitempty"`
	TargetCustomers       []*CustomerMeta     `json:"target_customers,omitempty"`
	UsageBillingFrequency string              `json:"usage_billing_frequency,omitempty"`
	FlatFeeBillingType    string              `json:"flat_fee_billing_type,omitempty"`
	FlatRate              decimal.Decimal     `json:"flat_rate,omitempty"`
	Status                string              `json:"status,omitempty"`
}

type Currency struct {
	Code   string `json:"code,omitempty"`
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

type PriceAdjustment struct {
	PriceAdjustmentName        string          `json:"price_adjustment_name,omitempty"`
	PriceAdjustmentDescription string          `json:"price_adjustment_description,omitempty"`
	PriceAdjustmentType        string          `json:"price_adjustment_type,omitempty"`
	PriceAdjustmentAmount      decimal.Decimal `json:"price_adjustment_amount,omitempty"`
}

type PricingTier struct {
	Type                string          `json:"type,omitempty"`
	RangeStart          decimal.Decimal `json:"range_start,omitempty"`
	RangeEnd            decimal.Decimal `json:"range_end,omitempty"`
	CostPerBatch        decimal.Decimal `json:"cost_per_batch,omitempty"`
	MetricUnitsPerBatch decimal.Decimal `json:"metric_units_per_batch,omitempty"`
	BatchRoundingType   string          `json:"batch_rounding_type,omitempty"`
}

type PrepaidCharge struct {
	Units          int    `json:"units,omitempty"`
	ChargeBehavior string `json:"charge_behavior,omitempty"`
}

type Filter struct {
	PropertyName    string `json:"property_name,omitempty"`
	Operator        string `json:"operator,omitempty"`
	ComparisonValue int    `json:"comparison_value,omitempty"`
}

type PricingComponent struct {
	BillableMetric         *BillableMetric `json:"billable_metric,omitempty"`
	Tiers                  []*PricingTier  `json:"tiers,omitempty"`
	PricingUnit            *Currency       `json:"pricing_unit,omitempty"`
	InvoicingIntervalUnit  string          `json:"invoicing_interval_unit,omitempty"`
	InvoicingIntervalCount int             `json:"invoicing_interval_count,omitempty"`
	ResetIntervalUnit      string          `json:"reset_interval_unit,omitempty"`
	ResetIntervalCount     int             `json:"reset_interval_count,omitempty"`
	PrepaidCharge          *PrepaidCharge  `json:"prepaid_charge,omitempty"`
}

type RecurringCharge struct {
	Name                   string          `json:"name,omitempty"`
	ChargeTiming           string          `json:"charge_timing,omitempty"`
	ChargeBehavior         string          `json:"charge_behavior,omitempty"`
	Amount                 decimal.Decimal `json:"amount,omitempty"`
	PricingUnit            *Currency       `json:"pricing_unit,omitempty"`
	InvoicingIntervalUnit  string          `json:"invoicing_interval_unit,omitempty"`
	InvoicingIntervalCount int             `json:"invoicing_interval_count,omitempty"`
	ResetIntervalUnit      string          `json:"reset_interval_unit,omitempty"`
	ResetIntervalCount     int             `json:"reset_interval_count,omitempty"`
}

// ListPlans retrieves an array of plan objects
// See: https://docs.uselotus.io/api-reference/plans/list-plans
func (c *Client) ListPlans(req ListPlansRequest) (resp ListPlansResponse, err error) {
	resp = make(ListPlansResponse, 0)
	err = c.get("/api/plans/", nil, resp)
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
