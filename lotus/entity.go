package lotus

import (
	"github.com/shopspring/decimal"
	"time"
)

type Customer struct {
	CustomerId        string                 `json:"customer_id"`
	Email             string                 `json:"email"`
	CustomerName      string                 `json:"customer_name"`
	DefaultCurrency   *Currency              `json:"default_currency"`
	TotalAmountDue    decimal.Decimal        `json:"total_amount_due"`
	TaxRate           decimal.Decimal        `json:"tax_rate"`
	Timezone          string                 `json:"timezone"`
	HasPaymentMethod  bool                   `json:"has_payment_method"`
	PaymentProvider   string                 `json:"payment_provider"`
	PaymentProviderId string                 `json:"payment_provider_id"`
	BillingAddress    *Address               `json:"billing_address"`
	ShippingAddress   *Address               `json:"shipping_address"`
	Subscriptions     []*Subscription        `json:"subscriptions"`
	Invoices          []interface{}          `json:"invoices"`      // TODO
	Integrations      interface{}            `json:"integrations"`  // TODO
	TaxProviders      []interface{}          `json:"tax_providers"` // TODO
	Properties        map[string]interface{} `json:"properties"`
}

type CustomerMeta struct {
	CustomerId   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Email        string `json:"email"`
}

type CustomerAddressMeta struct {
	CustomerMeta
	BillingAddress *Address `json:"billing_address"`
}

type SellerAddressMeta struct {
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
	Address *Address `json:"address"`
}

type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}

type FeatureEntitlement struct {
	Access                bool                         `json:"access"`
	Customer              *CustomerMeta                `json:"customer"`
	Feature               *FeatureMeta                 `json:"feature"`
	AccessPerSubscription []*FeatureSubscriptionAccess `json:"access_per_subscription"`
}

type FeatureSubscriptionAccess struct {
	Access       bool              `json:"access"`
	Subscription *SubscriptionMeta `json:"subscription"`
}

type MetricEntitlement struct {
	Access                bool                        `json:"access"`
	Customer              *CustomerMeta               `json:"customer"`
	Metric                *MetricMeta                 `json:"metric"`
	AccessPerSubscription []*MetricSubscriptionAccess `json:"access_per_subscription"`
}

type MetricSubscriptionAccess struct {
	MetricFreeLimit  decimal.Decimal   `json:"metric_free_limit"`
	MetricTotalLimit decimal.Decimal   `json:"metric_total_limit"`
	MetricUsage      decimal.Decimal   `json:"metric_usage"`
	Subscription     *SubscriptionMeta `json:"subscription"`
}

type Event struct {
	CustomerId    string                 `json:"customer_id"`
	EventName     string                 `json:"event_name"`
	IdempotencyId string                 `json:"idempotency_id"`
	TimeCreated   time.Time              `json:"time_created"`
	Properties    map[string]interface{} `json:"properties"`
}

type BillableMetric struct {
	MetricId           string    `json:"metric_id"`
	EventName          string    `json:"event_name"`
	PropertyName       string    `json:"property_name"`
	AggregationType    string    `json:"aggregation_type"`
	Granularity        string    `json:"granularity"`
	EventType          string    `json:"event_type"`
	MetricType         string    `json:"metric_type"`
	MetricName         string    `json:"metric_name"`
	NumericFilters     []*Filter `json:"numeric_filters"`
	CategoricalFilters []*Filter `json:"categorical_filters"`
	IsCostMetric       bool      `json:"is_cost_metric"`
	CustomSql          string    `json:"custom_sql"`
	Proration          string    `json:"proration"`
}

type MetricMeta struct {
	MetricId   string `json:"metric_id"`
	EventName  string `json:"event_name"`
	MetricName string `json:"metric_name"`
}

type FeatureMeta struct {
	FeatureId          string `json:"feature_id"`
	FeatureName        string `json:"feature_name"`
	FeatureDescription string `json:"feature_description"`
}

type Plan struct {
	PlanId              string          `json:"plan_id"`
	PlanName            string          `json:"plan_name"`
	PlanDuration        string          `json:"plan_duration"`
	PlanDescription     string          `json:"plan_description"`
	ExternalLinks       []*ExternalLink `json:"external_links"`
	NumVersions         int             `json:"num_versions"`
	ActiveVersion       int             `json:"active_version"`
	ActiveSubscriptions int             `json:"active_subscriptions"`
	Tags                []string        `json:"tags"`
	Versions            []*Version      `json:"versions"`
	ParentPlan          *Plan           `json:"parent_plan"`
	TargetCustomer      *CustomerMeta   `json:"target_customer"`
	DisplayVersion      *Version        `json:"display_version"`
}

type PlanMeta struct {
	PlanId    string `json:"plan_id"`
	PlanName  string `json:"plan_name"`
	Version   int    `json:"version"`
	VersionId string `json:"version_id"`
}

type ExternalLink struct {
	Source         string `json:"source"`
	ExternalPlanId string `json:"external_plan_id"`
}

type Version struct {
	Version               int                 `json:"version"`
	PlanName              string              `json:"plan_name"`
	LocalizedName         string              `json:"localized_name"`
	ActiveFrom            time.Time           `json:"active_from"`
	ActiveTo              time.Time           `json:"active_to"`
	CreatedOn             time.Time           `json:"created_on"`
	Description           string              `json:"description"`
	Currency              *Currency           `json:"currency"`
	Features              []*FeatureMeta      `json:"features"`
	RecurringCharges      []*RecurringCharge  `json:"recurring_charges"`
	Components            []*PricingComponent `json:"components"`
	PriceAdjustment       *PriceAdjustment    `json:"price_adjustment"`
	TargetCustomers       []*CustomerMeta     `json:"target_customers"`
	UsageBillingFrequency string              `json:"usage_billing_frequency"`
	FlatFeeBillingType    string              `json:"flat_fee_billing_type"`
	FlatRate              decimal.Decimal     `json:"flat_rate"`
	Status                PlanVersionStatus   `json:"status"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type PriceAdjustment struct {
	PriceAdjustmentName        string              `json:"price_adjustment_name"`
	PriceAdjustmentDescription string              `json:"price_adjustment_description"`
	PriceAdjustmentType        PriceAdjustmentType `json:"price_adjustment_type"`
	PriceAdjustmentAmount      decimal.Decimal     `json:"price_adjustment_amount"`
}

type PricingTier struct {
	Type                string          `json:"type"`
	RangeStart          decimal.Decimal `json:"range_start"`
	RangeEnd            decimal.Decimal `json:"range_end"`
	CostPerBatch        decimal.Decimal `json:"cost_per_batch"`
	MetricUnitsPerBatch decimal.Decimal `json:"metric_units_per_batch"`
	BatchRoundingType   string          `json:"batch_rounding_type"`
}

type PrepaidCharge struct {
	Units          int    `json:"units"`
	ChargeBehavior string `json:"charge_behavior"`
}

type Filter struct {
	PropertyName    string `json:"property_name"`
	Operator        string `json:"operator"`
	ComparisonValue int    `json:"comparison_value"`
}

type PricingComponent struct {
	BillableMetric         *BillableMetric `json:"billable_metric"`
	Tiers                  []*PricingTier  `json:"tiers"`
	PricingUnit            *Currency       `json:"pricing_unit"`
	InvoicingIntervalUnit  string          `json:"invoicing_interval_unit"`
	InvoicingIntervalCount int             `json:"invoicing_interval_count"`
	ResetIntervalUnit      string          `json:"reset_interval_unit"`
	ResetIntervalCount     int             `json:"reset_interval_count"`
	PrepaidCharge          *PrepaidCharge  `json:"prepaid_charge"`
}

type RecurringCharge struct {
	Name                   string          `json:"name"`
	ChargeTiming           string          `json:"charge_timing"`
	ChargeBehavior         string          `json:"charge_behavior"`
	Amount                 decimal.Decimal `json:"amount"`
	PricingUnit            *Currency       `json:"pricing_unit"`
	InvoicingIntervalUnit  string          `json:"invoicing_interval_unit"`
	InvoicingIntervalCount int             `json:"invoicing_interval_count"`
	ResetIntervalUnit      string          `json:"reset_interval_unit"`
	ResetIntervalCount     int             `json:"reset_interval_count"`
}

type AddonMeta struct {
	AddonId          string    `json:"addon_id"`
	AddonName        string    `json:"addon_name"`
	AddonType        AddonType `json:"addon_type"`
	BillingFrequency string    `json:"billing_frequency"`
}

type Addon struct {
	AddonSubscriptionId string                 `json:"addon_subscription_id"`
	Addon               *AddonMeta             `json:"addon"`
	Customer            *CustomerMeta          `json:"customer"`
	Parent              *Subscription          `json:"parent"`
	StartDate           time.Time              `json:"start_date"`
	EndDate             time.Time              `json:"end_date"`
	AutoRenew           bool                   `json:"auto_renew"`
	FullyBilled         bool                   `json:"fully_billed"`
	Metadata            map[string]interface{} `json:"metadata"`
}

type Subscription struct {
	SubscriptionId                    string                              `json:"subscription_id"`
	Customer                          *CustomerMeta                       `json:"customer"`
	BillingPlan                       *PlanMeta                           `json:"billing_plan"`
	AutoRenew                         bool                                `json:"auto_renew"`
	IsNew                             bool                                `json:"is_new"`
	StartDate                         time.Time                           `json:"start_date"`
	EndDate                           *time.Time                          `json:"end_date"`
	ComponentFixedChargesInitialUnits []*ComponentFixedChargesInitialUnit `json:"component_fixed_charges_initial_units"`
	SubscriptionFilters               []*SubscriptionFilter               `json:"subscription_filters"`
	Metadata                          map[string]interface{}              `json:"metadata"`
}

type ComponentFixedChargesInitialUnit struct {
	MetricId string `json:"metric_id"`
	Units    int    `json:"units"`
}

type SubscriptionFilter struct {
	PropertyName string `json:"property_name"`
	Value        string `json:"value"`
}

type SubscriptionMeta struct {
	StartDate           time.Time             `json:"start_date"`
	EndDate             time.Time             `json:"end_date"`
	Plan                *PlanMeta             `json:"plan"`
	SubscriptionFilters []*SubscriptionFilter `json:"subscription_filters"`
}

type Credit struct {
	CreditId           string            `json:"credit_id"`
	Amount             decimal.Decimal   `json:"amount"`
	AmountPaid         decimal.Decimal   `json:"amount_paid"`
	AmountRemaining    decimal.Decimal   `json:"amount_remaining"`
	AmountPaidCurrency *Currency         `json:"amount_paid_currency"`
	Currency           *Currency         `json:"currency"`
	Customer           *CustomerMeta     `json:"customer"`
	DrawDowns          []*CreditDrawDown `json:"drawdowns"`
	EffectiveAt        time.Time         `json:"effective_at"`
	ExpiresAt          time.Time         `json:"expires_at"`
	Description        string            `json:"description"`
	Status             CreditStatus      `json:"status"`
}

// CreditDrawDown refers to the record consuming particular amount from specified credit
type CreditDrawDown struct {
	CreditId    string          `json:"credit_id"`
	Amount      decimal.Decimal `json:"amount"`
	AppliedAt   time.Time       `json:"applied_at"`
	Description string          `json:"description"`
}

type InvoiceItem struct {
	Name                string                 `json:"name"`
	Plan                *PlanMeta              `json:"plan"`
	Quantity            decimal.Decimal        `json:"quantity"`
	Base                decimal.Decimal        `json:"base"`
	Amount              decimal.Decimal        `json:"amount"`
	Adjustments         []*InvoiceAdjustment   `json:"adjustments"`
	Subtotal            decimal.Decimal        `json:"subtotal"`
	BillingType         string                 `json:"billing_type"`
	SubscriptionFilters []*SubscriptionFilter  `json:"subscription_filters"`
	StartDate           time.Time              `json:"start_date"`
	EndDate             time.Time              `json:"end_date"`
	Metadata            map[string]interface{} `json:"metadata"`
}

type InvoiceAdjustment struct {
	Account        string              `json:"account"`
	AdjustmentType PriceAdjustmentType `json:"adjustment_type"`
	Amount         decimal.Decimal     `json:"amount"`
}

type Invoice struct {
	InvoiceId                string               `json:"invoice_id"`
	InvoiceNumber            string               `json:"invoice_number"`
	Customer                 *CustomerAddressMeta `json:"customer"`
	Currency                 *Currency            `json:"currency"`
	Amount                   decimal.Decimal      `json:"amount"`
	CostDue                  decimal.Decimal      `json:"cost_due"`
	IssueDate                time.Time            `json:"issue_date"`
	DueDate                  time.Time            `json:"due_date"`
	EndDate                  string               `json:"end_date"`   // e.g. 2024-02-03
	StartDate                string               `json:"start_date"` // e.g. 2024-02-03
	ExternalPaymentObjId     string               `json:"external_payment_obj_id"`
	ExternalPaymentObjStatus string               `json:"external_payment_obj_status"`
	ExternalPaymentObjType   PaymentProcessors    `json:"external_payment_obj_type"`
	InvoicePDF               string               `json:"invoice_pdf"`
	LineItems                []*InvoiceItem       `json:"line_items"`
	Seller                   *SellerAddressMeta   `json:"seller"`
	PaymentStatus            InvoiceStatus        `json:"payment_status"`
}
