package lotus

import "github.com/shopspring/decimal"

type CreateCustomerRequest struct {
	CustomerId          string                 `json:"customer_id,omitempty"`
	CustomerName        string                 `json:"customer_name,omitempty"`
	DefaultCurrencyCode string                 `json:"default_currency_code,omitempty"`
	Email               string                 `json:"email,omitempty"`
	PaymentProvider     string                 `json:"payment_provider,omitempty"`
	PaymentProviderId   string                 `json:"payment_provider_id,omitempty"`
	Properties          map[string]interface{} `json:"properties,omitempty"`
	TaxRate             decimal.Decimal        `json:"tax_rate"`
	Address             *Address               `json:"address,omitempty"`
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
	Address           *Address               `json:"address,omitempty,omitempty"`
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
	err = c.get("/api/customers/", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
