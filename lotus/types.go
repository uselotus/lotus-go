package lotus

import (
	"github.com/shopspring/decimal"
	"time"
)

type CreateCustomerRequest struct {
	CustomerId          *string                `json:"customer_id,omitempty"`
	CustomerName        *string                `json:"customer_name,omitempty"`
	DefaultCurrencyCode *string                `json:"default_currency_code,omitempty"`
	Email               *string                `json:"email,omitempty"`
	PaymentProvider     *string                `json:"payment_provider,omitempty"`
	PaymentProviderId   *string                `json:"payment_provider_id,omitempty"`
	Properties          map[string]interface{} `json:"properties,omitempty"`
	TaxRate             decimal.Decimal        `json:"tax_rate"`
	Address             *Address               `json:"address,omitempty"` // Deprecated, use billing address instead
	BillingAddress      *Address               `json:"billing_address,omitempty"`
	ShippingAddress     *Address               `json:"shipping_address,omitempty"`
}

type GetCustomerRequest struct {
	CustomerId string
}

type ListCustomersResponse []*Customer

type Customer struct {
	CustomerId        string                 `json:"customer_id,omitempty"`
	Email             string                 `json:"email,omitempty"`
	CustomerName      string                 `json:"customer_name,omitempty"`
	DefaultCurrency   *Currency              `json:"default_currency,omitempty"`
	TotalAmountDue    decimal.Decimal        `json:"total_amount_due"`
	TaxRate           decimal.Decimal        `json:"tax_rate"`
	Timezone          string                 `json:"timezone,omitempty"`
	PaymentProvider   string                 `json:"payment_provider,omitempty"`
	PaymentProviderId string                 `json:"payment_provider_id,omitempty"`
	Address           *Address               `json:"address,omitempty,omitempty"` // Deprecated, use billing address instead
	BillingAddress    *Address               `json:"billing_address,omitempty"`
	HasPaymentMethod  bool                   `json:"has_payment_method,omitempty"`
	ShippingAddress   *Address               `json:"shipping_address,omitempty"`
	Subscriptions     []*Subscription        `json:"subscriptions,omitempty"`
	Invoices          []interface{}          `json:"invoices,omitempty"`      // TODO
	Integrations      interface{}            `json:"integrations,omitempty"`  // TODO
	TaxProviders      []interface{}          `json:"tax_providers,omitempty"` // TODO
	Properties        map[string]interface{} `json:"properties,omitempty"`
}

type CustomerMeta struct {
	CustomerId   string `json:"customer_id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	Email        string `json:"email,omitempty"`
}

type Address struct {
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	Line1      string `json:"line1,omitempty"`
	Line2      string `json:"line2,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	State      string `json:"state,omitempty"`
}

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

type TrackEventsRequest struct {
	Batch []*Event `json:"batch"`
}

type TrackEventsResponse struct {
	Success      string                 `json:"success"`
	FailedEvents map[string]interface{} `json:"failed_events"`
}

type VerifyEventIngestionRequest struct {
	CustomerId         string   `json:"cust_id,omitempty"`
	IdempotencyIds     []string `json:"idempotency_ids"`
	NumberDaysLookBack int      `json:"number_days_lookback"`
}

type VerifyEventIngestionResponse struct {
	IdsNotFound []string `json:"ids_not_found"`
	Status      string   `json:"status"`
}

type Event struct {
	CustomerId    string                 `json:"customer_id,omitempty"`
	EventName     string                 `json:"event_name,omitempty"`
	IdempotencyId string                 `json:"idempotency_id,omitempty"`
	TimeCreated   time.Time              `json:"time_created,omitempty"`
	Properties    map[string]interface{} `json:"properties,omitempty"`
}

type BillableMetric struct {
	MetricId           string    `json:"metric_id,omitempty"`
	EventName          string    `json:"event_name,omitempty"`
	PropertyName       string    `json:"property_name,omitempty"`
	AggregationType    string    `json:"aggregation_type,omitempty"`
	Granularity        string    `json:"granularity,omitempty"`
	EventType          string    `json:"event_type,omitempty"`
	MetricType         string    `json:"metric_type,omitempty"`
	MetricName         string    `json:"metric_name,omitempty"`
	NumericFilters     []*Filter `json:"numeric_filters,omitempty"`
	CategoricalFilters []*Filter `json:"categorical_filters,omitempty"`
	IsCostMetric       bool      `json:"is_cost_metric,omitempty"`
	CustomSql          string    `json:"custom_sql,omitempty"`
	Proration          string    `json:"proration,omitempty"`
}

type MetricMeta struct {
	MetricId   string `json:"metric_id,omitempty"`
	EventName  string `json:"event_name,omitempty"`
	MetricName string `json:"metric_name,omitempty"`
}

type FeatureMeta struct {
	FeatureId          string `json:"feature_id,omitempty"`
	FeatureName        string `json:"feature_name,omitempty"`
	FeatureDescription string `json:"feature_description,omitempty"`
}

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

type PingResponse struct {
	OrganizationId string `json:"organization_id"`
}
