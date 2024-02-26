package lotus

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type PingResponse struct {
	OrganizationId string `json:"organization_id"`
}

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

type GetFeatureAccessRequest struct {
	CustomerId          string                // Required
	FeatureId           string                // Required
	SubscriptionFilters []*SubscriptionFilter // Optional
}

func (r GetFeatureAccessRequest) Strings() map[string]string {
	return map[string]string{
		"customer_id": r.CustomerId,
		"feature_id":  r.FeatureId,
	}
}

func (r GetFeatureAccessRequest) StringArrays() map[string][]string {
	m := make(map[string][]string)
	if len(r.SubscriptionFilters) > 0 {
		tmp := make([]string, 0)
		for i := range r.SubscriptionFilters {
			byt, err := json.Marshal(r.SubscriptionFilters[i])
			if err == nil {
				tmp = append(tmp, string(byt))
			}
		}
		m["subscription_filters"] = tmp
	}
	return m
}

type GetMetricAccessRequest struct {
	CustomerId          string                // Required
	MetricId            string                // Required
	SubscriptionFilters []*SubscriptionFilter // Optional
}

func (r GetMetricAccessRequest) Strings() map[string]string {
	return map[string]string{
		"customer_id": r.CustomerId,
		"metric_id":   r.MetricId,
	}
}

func (r GetMetricAccessRequest) StringArrays() map[string][]string {
	m := make(map[string][]string)
	if len(r.SubscriptionFilters) > 0 {
		tmp := make([]string, 0)
		for i := range r.SubscriptionFilters {
			byt, err := json.Marshal(r.SubscriptionFilters[i])
			if err == nil {
				tmp = append(tmp, string(byt))
			}
		}
		m["subscription_filters"] = tmp
	}
	return m
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
	IdsNotFound []string                   `json:"ids_not_found"`
	Status      VerifyEventIngestionStatus `json:"status"`
}

type ListPlansRequest struct {
	Duration            *PlanDuration       // Optional
	IncludeTags         []string            // Optional
	ExcludeTags         []string            // Optional
	IncludeTagsAll      []string            // Optional
	VersionCurrencyCode *string             // Optional
	VersionCustomType   *PlanVersionType    // Optional
	VersionStatus       []PlanVersionStatus // Optional
}

func (r ListPlansRequest) Strings() map[string]string {
	m := make(map[string]string)
	if r.Duration != nil {
		m["plan_duration"] = string(*r.Duration)
	}
	if r.VersionCurrencyCode != nil {
		m["version_currency_code"] = *r.VersionCurrencyCode
	}
	if r.VersionCustomType != nil {
		m["version_custom_type"] = string(*r.VersionCustomType)
	}
	return m
}

func (r ListPlansRequest) StringArrays() map[string][]string {
	m := make(map[string][]string)
	if len(r.IncludeTags) > 0 {
		m["include_tags"] = r.IncludeTags
	}
	if len(r.ExcludeTags) > 0 {
		m["exclude_tags"] = r.ExcludeTags
	}
	if len(r.IncludeTagsAll) > 0 {
		m["include_tags_all"] = r.IncludeTagsAll
	}
	if len(r.VersionStatus) > 0 {
		tmp := make([]string, 0)
		for i := range r.VersionStatus {
			tmp = append(tmp, string(r.VersionStatus[i]))
		}
		m["version_status"] = tmp
	}
	return m
}

type ListPlansResponse []*Plan

type GetPlanRequest struct {
	PlanId              string              // Required
	VersionCurrencyCode *string             // Optional
	VersionCustomType   *PlanVersionType    // Optional
	VersionStatus       []PlanVersionStatus // Optional
}

func (r GetPlanRequest) Strings() map[string]string {
	m := make(map[string]string)
	if r.VersionCurrencyCode != nil {
		m["version_currency_code"] = *r.VersionCurrencyCode
	}
	if r.VersionCustomType != nil {
		m["version_custom_type"] = string(*r.VersionCustomType)
	}
	return m
}

func (r GetPlanRequest) StringArrays() map[string][]string {
	m := make(map[string][]string)
	if len(r.VersionStatus) > 0 {
		tmp := make([]string, 0)
		for i := range r.VersionStatus {
			tmp = append(tmp, string(r.VersionStatus[i]))
		}
		m["version_status"] = tmp
	}
	return m
}

type CreateSubscriptionRequest struct {
	CustomerId                        string                              `json:"customer_id,omitempty"`
	PlanId                            string                              `json:"plan_id,omitempty"`
	AutoRenew                         bool                                `json:"auto_renew,omitempty"`
	IsNew                             bool                                `json:"is_new,omitempty"`
	StartDate                         time.Time                           `json:"start_date,omitempty"`
	EndDate                           *time.Time                          `json:"end_date,omitempty"`
	ComponentFixedChargesInitialUnits []*ComponentFixedChargesInitialUnit `json:"component_fixed_charges_initial_units,omitempty"`
	SubscriptionFilters               []*SubscriptionFilter               `json:"subscription_filters,omitempty"` // Add filter key, value pairs that define which events will be applied to this plan subscription
	Metadata                          map[string]interface{}              `json:"metadata,omitempty"`
}

type ListSubscriptionsRequest struct {
	CustomerId          string                // Required
	PlanId              *string               // Optional
	RangeStart          *time.Time            // Optional
	RangeEnd            *time.Time            // Optional
	Status              []SubscriptionStatus  // Optional
	SubscriptionFilters []*SubscriptionFilter // Optional
}

func (r ListSubscriptionsRequest) Strings() map[string]string {
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

func (r ListSubscriptionsRequest) StringArrays() map[string][]string {
	m := make(map[string][]string)
	if len(r.Status) > 0 {
		tmp := make([]string, 0)
		for i := range r.Status {
			tmp = append(tmp, string(r.Status[i]))
		}
		m["status"] = tmp
	}
	if len(r.SubscriptionFilters) > 0 {
		tmp := make([]string, 0)
		for i := range r.SubscriptionFilters {
			byt, err := json.Marshal(r.SubscriptionFilters[i])
			if err == nil {
				tmp = append(tmp, string(byt))
			}
		}
		m["subscription_filters"] = tmp
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
	UsageBehavior                     UsageBehavior                       `json:"usage_behavior,omitempty"`
	ComponentFixedChargesInitialUnits []*ComponentFixedChargesInitialUnit `json:"component_fixed_charges_initial_units,omitempty"`
	AutoRenew                         bool                                `json:"auto_renew,omitempty"` // TODO: add support server side
	Metadata                          map[string]interface{}              `json:"metadata,omitempty"`   // TODO: add support server side
}

type CancelSubscriptionRequest struct {
	SubscriptionId       string               `json:"-"`
	InvoicingBehavior    InvoicingBehavior    `json:"invoicing_behavior,omitempty"`
	UsageBillingBehavior UsageBillingBehavior `json:"usage_behavior,omitempty"` // Usage billing behavior, distinct from usage behavior
	FlatFeeBehavior      FlatFeeBehavior      `json:"flat_fee_behavior,omitempty"`
}

type AttachAddonRequest struct {
	SubscriptionId string `json:"-"`
	AddonId        string `json:"addon_id,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
}

type CancelAddonRequest struct {
	SubscriptionId       string               `json:"-"`
	AddonId              string               `json:"-"`
	InvoicingBehavior    InvoicingBehavior    `json:"invoicing_behavior,omitempty"`
	UsageBillingBehavior UsageBillingBehavior `json:"usage_behavior,omitempty"` // Usage billing behavior, distinct from usage behavior
	FlatFeeBehavior      FlatFeeBehavior      `json:"flat_fee_behavior,omitempty"`
}

type ListCreditsRequest struct {
	CustomerId      string         // Required
	CurrencyCode    *string        // Optional
	EffectiveAfter  *time.Time     // Optional
	EffectiveBefore *time.Time     // Optional
	ExpiresAfter    *time.Time     // Optional
	ExpiresBefore   *time.Time     // Optional
	IssuedAfter     *time.Time     // Optional
	IssuedBefore    *time.Time     // Optional
	Status          []CreditStatus // Optional
}

func (r ListCreditsRequest) Strings() map[string]string {
	m := make(map[string]string)
	m["customer_id"] = r.CustomerId
	if r.CurrencyCode != nil {
		m["currency_code"] = *r.CurrencyCode
	}
	if r.EffectiveAfter != nil {
		m["effective_after"] = r.EffectiveAfter.Format(time.RFC3339)
	}
	if r.EffectiveBefore != nil {
		m["effective_before"] = r.EffectiveBefore.Format(time.RFC3339)
	}
	if r.ExpiresAfter != nil {
		m["expires_after"] = r.ExpiresAfter.Format(time.RFC3339)
	}
	if r.ExpiresBefore != nil {
		m["expires_before"] = r.ExpiresBefore.Format(time.RFC3339)
	}
	if r.IssuedAfter != nil {
		m["issued_after"] = r.IssuedAfter.Format(time.RFC3339)
	}
	if r.IssuedBefore != nil {
		m["issued_before"] = r.IssuedBefore.Format(time.RFC3339)
	}
	return m
}

func (r ListCreditsRequest) StringArrays() map[string][]string {
	m := make(map[string][]string)
	if len(r.Status) > 0 {
		tmp := make([]string, 0)
		for i := range r.Status {
			tmp = append(tmp, string(r.Status[i]))
		}
		m["status"] = tmp
	}
	return m
}

type ListCreditsResponse []*Credit

type CreateCreditRequest struct {
	Amount                 decimal.Decimal  `json:"amount,omitempty"`
	CurrencyCode           string           `json:"currency_code,omitempty"`
	CustomerId             string           `json:"customer_id,omitempty"`
	AmountPaid             *decimal.Decimal `json:"amount_paid,omitempty"`
	AmountPaidCurrencyCode *string          `json:"amount_paid_currency_code,omitempty"`
	EffectiveAt            *time.Time       `json:"effective_at,omitempty"`
	ExpiresAt              *time.Time       `json:"expires_at,omitempty"`
	Description            *string          `json:"description,omitempty"`
}

type VoidCreditRequest struct {
	CreditId string `json:"-"`
}

type UpdateCreditRequest struct {
	CreditId    string     `json:"-"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	Description *string    `json:"description,omitempty"`
}

type ListInvoicesRequest struct {
	CustomerId    *string                // Optional
	PaymentStatus []InvoicePaymentStatus // Optional
}

func (r ListInvoicesRequest) Strings() map[string]string {
	m := make(map[string]string)
	if r.CustomerId != nil {
		m["customer_id"] = *r.CustomerId
	}
	return m
}

func (r ListInvoicesRequest) StringArrays() map[string][]string {
	m := make(map[string][]string)
	if len(r.PaymentStatus) > 0 {
		tmp := make([]string, 0)
		for i := range r.PaymentStatus {
			tmp = append(tmp, string(r.PaymentStatus[i]))
		}
		m["payment_status"] = tmp
	}
	return m
}

type ListInvoicesResponse []*Invoice

type GetInvoiceRequest struct {
	InvoiceId string
}

type GetInvoicePDFUrlRequest struct {
	InvoiceId string
}

type GetInvoicePDFUrlResponse struct {
	Url string `json:"url"`
}
