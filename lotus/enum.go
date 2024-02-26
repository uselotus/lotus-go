package lotus

// See: https://github.com/uselotus/lotus/blob/main/backend/metering_billing/utils/enums/enums.py

// VerifyEventIngestionStatus defines verify event ingestion status
type VerifyEventIngestionStatus string

const (
	VerifyEventIngestionStatusSuccess VerifyEventIngestionStatus = "success"
)

// CreditStatus defines credit status
type CreditStatus string

const (
	CreditStatusActive   CreditStatus = "active"
	CreditStatusInactive CreditStatus = "inactive"
)

// AddonType defines addon types
type AddonType string

const (
	AddonTypeFlat       AddonType = "flat"
	AddonTypeUsageBased AddonType = "usage_based"
)

// AddonBillingFrequency defines addon billing frequencies
type AddonBillingFrequency string

const (
	AddonBillingFrequencyOneTime   AddonBillingFrequency = "one_time"
	AddonBillingFrequencyRecurring AddonBillingFrequency = "recurring"
)

// InvoiceStatus defines invoice status
type InvoiceStatus string

const (
	InvoiceStatusDraft  InvoiceStatus = "draft"
	InvoiceStatusVoided InvoiceStatus = "voided"
	InvoiceStatusPaid   InvoiceStatus = "paid"
	InvoiceStatusUnpaid InvoiceStatus = "unpaid"
)

// InvoicePaymentStatus defines invoice payment status
type InvoicePaymentStatus string

const (
	InvoicePaymentStatusPaid   InvoicePaymentStatus = "paid"
	InvoicePaymentStatusUnpaid InvoicePaymentStatus = "unpaid"
)

// PriceTierType defines price tier types
type PriceTierType string

const (
	PriceTierTypeFlat    PriceTierType = "flat"
	PriceTierTypePerUnit PriceTierType = "per_unit"
	PriceTierTypeFree    PriceTierType = "free"
)

// BatchRoundingType defines batch rounding types
type BatchRoundingType string

const (
	BatchRoundingTypeRoundUp      BatchRoundingType = "round_up"
	BatchRoundingTypeRoundDown    BatchRoundingType = "round_down"
	BatchRoundingTypeRoundNearest BatchRoundingType = "round_nearest"
	BatchRoundingTypeNoRounding   BatchRoundingType = "no_rounding"
)

// MetricAggregation defines metric aggregation methods
type MetricAggregation string

const (
	MetricAggregationCount   MetricAggregation = "count"
	MetricAggregationSum     MetricAggregation = "sum"
	MetricAggregationMax     MetricAggregation = "max"
	MetricAggregationUnique  MetricAggregation = "unique"
	MetricAggregationLatest  MetricAggregation = "latest"
	MetricAggregationAverage MetricAggregation = "average"
)

// PriceAdjustmentType defines price adjustment types
type PriceAdjustmentType string

const (
	PriceAdjustmentTypePercentage    PriceAdjustmentType = "percentage"
	PriceAdjustmentTypeFixed         PriceAdjustmentType = "fixed"
	PriceAdjustmentTypePriceOverride PriceAdjustmentType = "price_override"
)

// PaymentProcessors defines payment processor types
type PaymentProcessors string

const (
	PaymentProcessorsStripe    PaymentProcessors = "stripe"
	PaymentProcessorsBraintree PaymentProcessors = "braintree"
)

// MetricType defines metric types
type MetricType string

const (
	MetricTypeCounter MetricType = "counter"
	MetricTypeRate    MetricType = "rate"
	MetricTypeCustom  MetricType = "custom"
	MetricTypeGauge   MetricType = "gauge"
)

// CustomerBalanceAdjustmentStatus defines customer balance adjustment statuses
type CustomerBalanceAdjustmentStatus CreditStatus // Deprecated, use CreditStatus

// MetricGranularity defines metric granularity options
type MetricGranularity string

const (
	MetricGranularitySecond  MetricGranularity = "seconds"
	MetricGranularityMinute  MetricGranularity = "minutes"
	MetricGranularityHour    MetricGranularity = "hours"
	MetricGranularityDay     MetricGranularity = "days"
	MetricGranularityMonth   MetricGranularity = "months"
	MetricGranularityQuarter MetricGranularity = "quarters"
	MetricGranularityYear    MetricGranularity = "years"
	MetricGranularityTotal   MetricGranularity = "total"
)

// EventType defines event types
type EventType string

const (
	EventTypeDelta EventType = "delta"
	EventTypeTotal EventType = "total"
)

// PlanDuration defines plan duration options
type PlanDuration string

const (
	PlanDurationMonthly   PlanDuration = "monthly"
	PlanDurationQuarterly PlanDuration = "quarterly"
	PlanDurationYearly    PlanDuration = "yearly"
)

// UsageBillingFrequency defines usage billing frequency options
type UsageBillingFrequency string

const (
	UsageBillingFrequencyMonthly     UsageBillingFrequency = "monthly"
	UsageBillingFrequencyQuarterly   UsageBillingFrequency = "quarterly"
	UsageBillingFrequencyEndOfPeriod UsageBillingFrequency = "end_of_period"
)

// ComponentResetFrequency defines component reset frequency options
type ComponentResetFrequency string

const (
	ComponentResetFrequencyWeekly    ComponentResetFrequency = "weekly"
	ComponentResetFrequencyMonthly   ComponentResetFrequency = "monthly"
	ComponentResetFrequencyQuarterly ComponentResetFrequency = "quarterly"
	ComponentResetFrequencyNone      ComponentResetFrequency = "none"
)

// InvoiceChargeTimingType defines invoice charge timing types
type InvoiceChargeTimingType string

const (
	InvoiceChargeTimingTypeInArrears    InvoiceChargeTimingType = "in_arrears"
	InvoiceChargeTimingTypeIntermediate InvoiceChargeTimingType = "intermediate"
	InvoiceChargeTimingTypeInAdvance    InvoiceChargeTimingType = "in_advance"
	InvoiceChargeTimingTypeOneTime      InvoiceChargeTimingType = "one_time"
)

// UsageCalcGranularity defines usage calculation granularity options
type UsageCalcGranularity string

const (
	UsageCalcGranularityDaily UsageCalcGranularity = "day"
	UsageCalcGranularityTotal UsageCalcGranularity = "total"
)

// NumericFilterOperators defines numeric filter operators
type NumericFilterOperators string

const (
	NumericFilterOperatorsGTE NumericFilterOperators = "gte"
	NumericFilterOperatorsGT  NumericFilterOperators = "gt"
	NumericFilterOperatorsEQ  NumericFilterOperators = "eq"
	NumericFilterOperatorsLT  NumericFilterOperators = "lt"
	NumericFilterOperatorsLTE NumericFilterOperators = "lte"
)

// CategoricalFilterOperators defines categorical filter operators
type CategoricalFilterOperators string

const (
	CategoricalFilterOperatorsIsIn    CategoricalFilterOperators = "isin"
	CategoricalFilterOperatorsIsNotIn CategoricalFilterOperators = "isnotin"
)

// SubscriptionStatus defines subscription statuses
type SubscriptionStatus string

const (
	SubscriptionStatusActive     SubscriptionStatus = "active"
	SubscriptionStatusEnded      SubscriptionStatus = "ended"
	SubscriptionStatusNotStarted SubscriptionStatus = "not_started"
)

// PlanCustomType defines plan custom types
type PlanCustomType string

const (
	PlanCustomTypeCustomOnly PlanCustomType = "custom_only"
	PlanCustomTypePublicOnly PlanCustomType = "public_only"
	PlanCustomTypeAll        PlanCustomType = "all"
)

// PlanVersionStatus defines plan version statuses
type PlanVersionStatus string

const (
	PlanVersionStatusActive        PlanVersionStatus = "active"
	PlanVersionStatusRetiring      PlanVersionStatus = "retiring"
	PlanVersionStatusGrandfathered PlanVersionStatus = "grandfathered"
	PlanVersionStatusDeleted       PlanVersionStatus = "deleted"
	PlanVersionStatusInactive      PlanVersionStatus = "inactive"
	PlanVersionStatusNotStarted    PlanVersionStatus = "not_started"
)

// PlanStatus defines plan statuses
type PlanStatus string

const (
	PlanStatusActive       PlanStatus = "active"
	PlanStatusDeleted      PlanStatus = "deleted"
	PlanStatusExperimental PlanStatus = "experimental"
)

// BacktestKPI defines backtest key performance indicators
type BacktestKPI string

const (
	BacktestKPITotalRevenue BacktestKPI = "total_revenue"
)

// AnalysisKPI defines analysis key performance indicators
type AnalysisKPI string

const (
	AnalysisKPITotalRevenue   AnalysisKPI = "total_revenue"
	AnalysisKPIAverageRevenue AnalysisKPI = "average_revenue"
	AnalysisKPINewRevenue     AnalysisKPI = "new_revenue"
	AnalysisKPITotalCost      AnalysisKPI = "total_cost"
	AnalysisKPIProfit         AnalysisKPI = "profit"
	AnalysisKPIChurn          AnalysisKPI = "churn"
)

// ExperimentStatus defines experiment statuses
type ExperimentStatus string

const (
	ExperimentStatusRunning   ExperimentStatus = "running"
	ExperimentStatusCompleted ExperimentStatus = "completed"
	ExperimentStatusFailed    ExperimentStatus = "failed"
)

// ProductStatus defines product statuses
type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusArchived ProductStatus = "archived"
)

// MetricStatus defines metric statuses
type MetricStatus string

const (
	MetricStatusActive   MetricStatus = "active"
	MetricStatusArchived MetricStatus = "archived"
)

// MakePlanVersionActiveType defines types for making plan version active
type MakePlanVersionActiveType string

const (
	MakePlanVersionActiveTypeReplaceOnRenewal MakePlanVersionActiveType = "replace_on_renewal"
	MakePlanVersionActiveTypeGrandfather      MakePlanVersionActiveType = "grandfather"
)

// OrganizationStatus defines organization statuses
type OrganizationStatus string

const (
	OrganizationStatusActive  OrganizationStatus = "Active"
	OrganizationStatusInvited OrganizationStatus = "Invited"
)

// WebhookTriggerEvents defines webhook trigger events
type WebhookTriggerEvents string

const (
	WebhookTriggerEventsCustomerCreated       WebhookTriggerEvents = "customer.created"
	WebhookTriggerEventsInvoiceCreated        WebhookTriggerEvents = "invoice.created"
	WebhookTriggerEventsInvoicePaid           WebhookTriggerEvents = "invoice.paid"
	WebhookTriggerEventsInvoicePastDue        WebhookTriggerEvents = "invoice.past_due"
	WebhookTriggerEventsSubscriptionCreated   WebhookTriggerEvents = "subscription.created"
	WebhookTriggerEventsUsageAlertTriggered   WebhookTriggerEvents = "usage_alert.triggered"
	WebhookTriggerEventsSubscriptionCancelled WebhookTriggerEvents = "subscription.cancelled"
	WebhookTriggerEventsSubscriptionRenewed   WebhookTriggerEvents = "subscription.renewed"
)

// FlatFeeBehavior defines flat fee behavior options (when canceling a subscription)
// If null or not provided, the charge's default behavior will be used according to the subscription's start and end dates.
// If charge_full, the full flat fee will be charged, regardless of the duration of the subscription.
// If it is refund, the flat fee will not be charged. If charge_prorated, the prorated flat fee will be charged.
type FlatFeeBehavior string

const (
	FlatFeeBehaviorRefund         FlatFeeBehavior = "refund"
	FlatFeeBehaviorChargeProrated FlatFeeBehavior = "charge_prorated"
	FlatFeeBehaviorChargeFull     FlatFeeBehavior = "charge_full"
)

// UsageBehavior defines usage behavior options (when switching subscription plan)
// Transfer to new subscription will transfer the usage from the old subscription to the new subscription,
// whereas keep_separate will reset the usage to 0 for the new subscription, while keeping the old usage on the old subscription and charging for that appropriately at the end of the month.
type UsageBehavior string

const (
	UsageBehaviorTransferToNewSubscription UsageBehavior = "transfer_to_new_subscription"
	UsageBehaviorKeepSeparate              UsageBehavior = "keep_separate"
)

// InvoicingBehavior defines the invoicing behavior to use when replacing the plan or canceling a subscription.
// Invoice now will invoice the customer for the prorated difference of the old plan and the new plan,
// whereas add_to_next_invoice will wait until the end of the subscription to do the calculation.
type InvoicingBehavior string

const (
	InvoicingBehaviorInvoiceNow       InvoicingBehavior = "invoice_now"
	InvoicingBehaviorAddToNextInvoice InvoicingBehavior = "add_to_next_invoice"
)

// ChargeableItemType defines chargeable item types
type ChargeableItemType string

const (
	ChargeableItemTypeUsageCharge        ChargeableItemType = "usage_charge"
	ChargeableItemTypePrepaidUsageCharge ChargeableItemType = "prepaid_usage_charge"
	ChargeableItemTypeRecurringCharge    ChargeableItemType = "recurring_charge"
	ChargeableItemTypeOneTimeCharge      ChargeableItemType = "one_time_charge"
	ChargeableItemTypePlanAdjustment     ChargeableItemType = "plan_adjustment"
	ChargeableItemTypeCustomerAdjustment ChargeableItemType = "customer_adjustment"
	ChargeableItemTypeTax                ChargeableItemType = "tax"
)

// AccountsReceivableTransactionTypes defines types of accounts receivable transactions
type AccountsReceivableTransactionTypes int

const (
	AccountsReceivableTransactionTypesInvoice    AccountsReceivableTransactionTypes = 1
	AccountsReceivableTransactionTypesReceipt    AccountsReceivableTransactionTypes = 2
	AccountsReceivableTransactionTypesAdjustment AccountsReceivableTransactionTypes = 3
	AccountsReceivableTransactionTypesReversal   AccountsReceivableTransactionTypes = 4
)

// UsageBillingBehavior defines usage billing behaviors (when cancelling a subscription)
// If bill_full, current usage will be billed on the invoice.
// If bill_none, current un-billed usage will be dropped from the invoice.
// Defaults to bill_full.
type UsageBillingBehavior string

const (
	UsageBillingBehaviorBillFull UsageBillingBehavior = "bill_full"
	UsageBillingBehaviorBillNone UsageBillingBehavior = "bill_none"
)

// OrganizationSettingNames defines organization setting names
type OrganizationSettingNames string

const (
	OrganizationSettingNamesGenerateCustomerInStripeAfterLotus    OrganizationSettingNames = "generate_customer_after_creating_in_lotus"
	OrganizationSettingNamesGenerateCustomerInBraintreeAfterLotus OrganizationSettingNames = "gen_cust_in_braintree_after_lotus"
	OrganizationSettingNamesPaymentGracePeriod                    OrganizationSettingNames = "payment_grace_period"
	OrganizationSettingNamesCrmCustomerSource                     OrganizationSettingNames = "crm_customer_source"
)

// TagGroup defines tag groups
type TagGroup string

const (
	TagGroupPlan TagGroup = "plan"
)

// OrganizationSettingGroups defines organization setting groups
type OrganizationSettingGroups string

const (
	OrganizationSettingGroupsStripe    OrganizationSettingGroups = "stripe"
	OrganizationSettingGroupsBraintree OrganizationSettingGroups = "braintree"
	OrganizationSettingGroupsBilling   OrganizationSettingGroups = "billing"
	OrganizationSettingGroupsCRM       OrganizationSettingGroups = "crm"
)

// TaxProvider defines tax providers
type TaxProvider int

const (
	TaxProviderTaxjar   TaxProvider = 1
	TaxProviderLotus    TaxProvider = 2
	TaxProviderNetsuite TaxProvider = 3
)

// PlanVersionType defines filters to versions that have this custom type.
// If you choose custom_only, you will only see versions that have target customers.
// If you choose public_only, you will only see versions that do not have target customers.
type PlanVersionType string

const (
	PlanVersionTypeCustomOnly PlanVersionType = "custom_only"
	PlanVersionTypePublicOnly PlanVersionType = "public_only"
	PlanVersionTypeAll        PlanVersionType = "all"
)
