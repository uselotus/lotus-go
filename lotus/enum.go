package lotus

type Duration string

const (
	DurationMonthly   Duration = "monthly"
	DurationQuarterly Duration = "quarterly"
	DurationYearly    Duration = "yearly"
)

// InvoicingBehavior defines the invoicing behavior to use when replacing the plan or canceling a subscription.
// Invoice now will invoice the customer for the prorated difference of the old plan and the new plan,
// whereas add_to_next_invoice will wait until the end of the subscription to do the calculation.
type InvoicingBehavior string

const (
	InvoiceNow       InvoicingBehavior = "invoice_now"
	AddToNextInvoice InvoicingBehavior = "add_to_next_invoice"
)

// SwitchPlanUsageBehavior defines the usage behavior to use when replacing the plan.
// Transfer to new subscription will transfer the usage from the old subscription to the new subscription,
// whereas keep_separate will reset the usage to 0 for the new subscription, while keeping the old usage on the old subscription and charging for that appropriately at the end of the month.
type SwitchPlanUsageBehavior string

const (
	TransferToNewSubscription SwitchPlanUsageBehavior = "transfer_to_new_subscription"
	KeepSeparate              SwitchPlanUsageBehavior = "keep_separate"
)

// CancellationFlatFeeBehavior defines when canceling a subscription, the behavior used to calculate the flat fee.
// If null or not provided, the charge's default behavior will be used according to the subscription's start and end dates.
// If charge_full, the full flat fee will be charged, regardless of the duration of the subscription.
// If it is refund, the flat fee will not be charged. If charge_prorated, the prorated flat fee will be charged.
type CancellationFlatFeeBehavior string

const (
	Refund         CancellationFlatFeeBehavior = "refund"
	ChargeProrated CancellationFlatFeeBehavior = "charge_prorated"
	ChargeFull     CancellationFlatFeeBehavior = "charge_full"
)

// CancellationUsageBehavior defines the usage behavior to use when cancelling a subscription.
// If bill_full, current usage will be billed on the invoice.
// If bill_none, current un-billed usage will be dropped from the invoice.
// Defaults to bill_full.
type CancellationUsageBehavior string

const (
	BillFull CancellationUsageBehavior = "bill_full"
	BillNone CancellationUsageBehavior = "bill_none"
)
