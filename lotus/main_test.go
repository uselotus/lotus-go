package lotus

import (
	"fmt"
	"github.com/biter777/countries"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

type TestEntity struct {
	c              *Client
	customerId     string
	planId         string
	plan2Id        string
	featureId      string
	subscriptionId string
	eventName      string
	metricId       string
}

var test TestEntity

func newEventProperty() map[string]interface{} {
	return map[string]interface{}{
		"duration": 20,
	}
}

func setup(t *testing.T) {
	test.c = NewClient(os.Getenv("LOTUS_SDK_TEST_HOST"), os.Getenv("LOTUS_SDK_TEST_KEY")).WithDebug(true)
	test.customerId = fmt.Sprintf("test_c%v", time.Now().Unix())
	test.planId = os.Getenv("LOTUS_SDK_TEST_PLAN")
	test.plan2Id = os.Getenv("LOTUS_SDK_TEST_PLAN2")
	test.featureId = os.Getenv("LOTUS_SDK_TEST_FEATURE")
	test.eventName = os.Getenv("LOTUS_SDK_TEST_EVENT")
	test.metricId = os.Getenv("LOTUS_SDK_TEST_METRIC")
}

func TestClient(t *testing.T) {
	setup(t)

	// Customer tests
	// -----------------------------------------------------------------
	t.Run("Create customer with basic info", func(t *testing.T) {
		email := test.customerId + "@sample.com"
		currency := countries.US.Currency().Alpha()
		req := CreateCustomerRequest{
			CustomerId:          &test.customerId,
			CustomerName:        &test.customerId,
			Email:               &email,
			DefaultCurrencyCode: &currency,
		}
		resp, err := test.c.CreateCustomer(req)
		assert.Nil(t, err, "CreateCustomer")
		assert.Equal(t, test.customerId, resp.CustomerId, "CreateCustomer")
		assert.NotEmpty(t, resp.CustomerName, "CreateCustomer")
		assert.NotEmpty(t, resp.Email, "CreateCustomer")
		assert.NotEmpty(t, resp.DefaultCurrency.Code, "CreateCustomer")
	})

	t.Run("Create customer with detailed info", func(t *testing.T) {
		// Fake addresses for test
		addr1 := Address{
			City:       "Paris",
			Country:    countries.France.Alpha2(),
			Line1:      "10 Rue de la Paix",
			Line2:      "Apartment 5B",
			PostalCode: "75002",
			State:      "Île-de-France",
		}
		addr2 := Address{
			City:       "Berlin",
			Country:    countries.Germany.Alpha2(),
			Line1:      "Friedrichstraße 176-179",
			Line2:      "Floor 3, Office 301",
			PostalCode: "10117",
			State:      "Berlin",
		}
		id := test.customerId + "_detailed"
		email := test.customerId + "_detailed@sample.com"
		currency := countries.France.Currency().Alpha()
		req := CreateCustomerRequest{
			CustomerId:          &id,
			CustomerName:        &id,
			Email:               &email,
			DefaultCurrencyCode: &currency,
			Properties: map[string]interface{}{
				"remark": "Remark content",
			},
			TaxRate:         decimal.NewFromFloat(0.15),
			BillingAddress:  &addr1,
			ShippingAddress: &addr2,
		}
		resp, err := test.c.CreateCustomer(req)
		assert.Nil(t, err, "CreateCustomer")
		assert.Equal(t, test.customerId+"_detailed", resp.CustomerId, "CreateCustomer")
		assert.Equal(t, test.customerId+"_detailed", resp.CustomerName, "CreateCustomer")
		assert.Contains(t, resp.Email, test.customerId+"_detailed", "CreateCustomer")
		assert.Equal(t, "EUR", resp.DefaultCurrency.Code, "CreateCustomer")
		assert.True(t, resp.TaxRate.Equals(decimal.NewFromFloat(0.15)), "CreateCustomer")
		assert.Equal(t, &addr1, resp.BillingAddress, "CreateCustomer")
		assert.Equal(t, &addr2, resp.ShippingAddress, "CreateCustomer")
	})

	t.Run("Create customer with invalid info", func(t *testing.T) {
		currency := countries.US.Currency().Alpha()
		req := CreateCustomerRequest{
			CustomerId:          &test.customerId,
			CustomerName:        &test.customerId,
			DefaultCurrencyCode: &currency,
		}
		_, err := test.c.CreateCustomer(req)
		assert.NotNil(t, err, "CreateCustomer")
		assert.True(t, IsLotusError(err), "CreateCustomer")
	})

	t.Run("Create duplicated customer", func(t *testing.T) {
		email := test.customerId + "@sample.com"
		currency := countries.US.Currency().Alpha()
		req := CreateCustomerRequest{
			CustomerId:          &test.customerId,
			CustomerName:        &test.customerId,
			Email:               &email,
			DefaultCurrencyCode: &currency,
		}
		resp, err := test.c.CreateCustomer(req)
		assert.NotNilf(t, err, "CreateCustomer")
		assert.True(t, IsLotusError(err), "CreateCustomer")
		assert.True(t, IsDuplicated(err), "CreateCustomer")
		assert.Nil(t, resp, "CreateCustomer")
	})

	t.Run("Get created customer", func(t *testing.T) {
		req := GetCustomerRequest{
			CustomerId: test.customerId,
		}
		resp, err := test.c.GetCustomer(req)
		assert.Nil(t, err, "GetCustomer")
		assert.Equal(t, test.customerId, resp.CustomerId, "GetCustomer")
		assert.Equal(t, test.customerId, resp.CustomerName, "GetCustomer")
		assert.Equal(t, test.customerId+"@sample.com", resp.Email, "GetCustomer")
	})

	t.Run("List customers", func(t *testing.T) {
		resp, err := test.c.ListCustomers()
		assert.Nil(t, err, "ListCustomers")
		assert.NotEmpty(t, len(resp), "ListCustomers")
	})

	t.Run("Get customer that does not exist", func(t *testing.T) {
		req := GetCustomerRequest{
			CustomerId: fmt.Sprintf("%v", time.Now().Unix()),
		}
		_, err := test.c.GetCustomer(req)
		assert.NotNil(t, err, "GetCustomer")
		assert.True(t, IsNotFound(err), "GetCustomer")
	})

	// Plan tests
	// -----------------------------------------------------------------
	t.Run("Get plan", func(t *testing.T) {
		resp, err := test.c.GetPlan(GetPlanRequest{
			PlanId: test.planId,
		})
		assert.Nil(t, err, "GetPlan")
		assert.Equal(t, test.planId, resp.PlanId, "GetPlan")
	})

	t.Run("Get plan that does not exist", func(t *testing.T) {
		_, err := test.c.GetPlan(GetPlanRequest{
			PlanId: "plan_00000000000000000000000000000000",
		})
		assert.NotNil(t, err, "GetPlan")
		assert.True(t, IsNotFound(err), "GetPlan")
	})

	t.Run("List plans", func(t *testing.T) {
		resp, err := test.c.ListPlans(ListPlansRequest{
			Duration: DurationMonthly,
		})
		assert.Nil(t, err, "ListPlans")
		assert.NotEmpty(t, len(resp), "ListPlans")
	})

	// Subscription tests (1)
	// -----------------------------------------------------------------
	t.Run("Create subscription", func(t *testing.T) {
		req := CreateSubscriptionRequest{
			CustomerId: test.customerId,
			PlanId:     test.planId,
			AutoRenew:  true,
			StartDate:  time.Now().AddDate(0, 0, -1),
			Metadata: map[string]interface{}{
				"CreatedBy": "SDK Test",
			},
		}
		resp, err := test.c.CreateSubscription(req)
		assert.Nil(t, err, "CreateSubscription")
		assert.NotEmpty(t, resp.SubscriptionId, "CreateSubscription")
		assert.NotEmpty(t, resp.Metadata, "CreateSubscription")
		assert.NotEmpty(t, resp.StartDate, "CreateSubscription")
		assert.NotEmpty(t, resp.EndDate, "CreateSubscription")
		assert.True(t, resp.AutoRenew, "CreateSubscription")
		assert.Equal(t, test.customerId, resp.Customer.CustomerId, "CreateSubscription")
		assert.Equal(t, test.planId, resp.BillingPlan.PlanId, "CreateSubscription")
		assert.Contains(t, resp.Metadata, "CreatedBy", "CreateSubscription")

		test.subscriptionId = resp.SubscriptionId
	})

	t.Run("Get customer with subscription", func(t *testing.T) {
		req := GetCustomerRequest{
			CustomerId: test.customerId,
		}
		resp, err := test.c.GetCustomer(req)
		assert.Nil(t, err, "GetCustomer")
		assert.Equal(t, 1, len(resp.Subscriptions), "GetCustomer")
		assert.NotEmpty(t, resp.Subscriptions[0].SubscriptionId, "GetCustomer")
		assert.Equal(t, test.customerId, resp.Subscriptions[0].Customer.CustomerId, "GetCustomer")
		assert.Equal(t, test.planId, resp.Subscriptions[0].BillingPlan.PlanId, "GetCustomer")
	})

	t.Run("Create subscription for a plan does not exist", func(t *testing.T) {
		req := CreateSubscriptionRequest{
			CustomerId: test.customerId,
			PlanId:     "plan_00000000000000000000000000000000",
			AutoRenew:  true,
			StartDate:  time.Now().AddDate(0, 0, -1),
			Metadata: map[string]interface{}{
				"CreatedBy": "SDK Test",
			},
		}
		_, err := test.c.CreateSubscription(req)
		assert.NotNil(t, err, "CreateSubscription")
		assert.True(t, IsNotFound(err), "CreateSubscription")
	})

	t.Run("Create subscription for a customer does not exist", func(t *testing.T) {
		req := CreateSubscriptionRequest{
			CustomerId: fmt.Sprintf("%v", time.Now().Unix()),
			PlanId:     test.planId,
			AutoRenew:  true,
			StartDate:  time.Now().AddDate(0, 0, -1),
			Metadata: map[string]interface{}{
				"CreatedBy": "SDK Test",
			},
		}
		_, err := test.c.CreateSubscription(req)
		assert.NotNil(t, err, "CreateSubscription")
		assert.True(t, IsNotFound(err), "CreateSubscription")
	})

	t.Run("List subscriptions", func(t *testing.T) {
		start := time.Now().AddDate(0, -1, 0)
		end := time.Now().AddDate(0, 0, 1)
		req := ListSubscriptionsRequest{
			CustomerId: test.customerId,
			PlanId:     &test.planId,
			RangeStart: &start,
			RangeEnd:   &end,
		}
		resp, err := test.c.ListSubscriptions(req)
		assert.Nil(t, err, "ListSubscriptions")
		if assert.NotEmpty(t, len(resp), "ListSubscriptions") {
			assert.Equal(t, test.customerId, resp[0].Customer.CustomerId, "ListSubscriptions")
			assert.NotEmpty(t, resp[0].BillingPlan.PlanId, "ListSubscriptions")
		}
	})

	// Event tests
	// -----------------------------------------------------------------
	t.Run("Ingest events", func(t *testing.T) {
		req := TrackEventsRequest{
			Batch: []*Event{
				{
					CustomerId:    test.customerId,
					EventName:     test.eventName,
					IdempotencyId: "event_foo",
					TimeCreated:   time.Now(),
					Properties:    newEventProperty(),
				},
				{
					CustomerId:    test.customerId,
					EventName:     test.eventName,
					IdempotencyId: "event_bar",
					TimeCreated:   time.Now(),
					Properties:    newEventProperty(),
				},
			},
		}
		resp, err := test.c.TrackEvents(req)
		assert.Nil(t, err, "TrackEvents")
		assert.Equal(t, "all", resp.Success, "TrackEvents")
		assert.Equal(t, 0, len(resp.FailedEvents), "TrackEvents")
	})

	t.Run("Verify event ingestion", func(t *testing.T) {
		time.Sleep(time.Second * 2)
		req := VerifyEventIngestionRequest{
			CustomerId: test.customerId,
			IdempotencyIds: []string{
				"event_foo",
				"event_bar",
			},
			NumberDaysLookBack: 7,
		}
		resp, err := test.c.VerifyEventIngestion(req)
		assert.Nil(t, err, "VerifyEventIngestion")
		assert.Equal(t, "success", resp.Status, "VerifyEventIngestion")
		assert.Equal(t, 0, len(resp.IdsNotFound), "VerifyEventIngestion")
	})

	// Entitlement tests
	// -----------------------------------------------------------------
	t.Run("Check feature access", func(t *testing.T) {
		req := GetFeatureAccessRequest{
			CustomerId: test.customerId,
			FeatureId:  test.featureId,
		}
		resp, err := test.c.GetFeatureAccess(req)
		assert.Nil(t, err, "GetFeatureAccess")
		assert.True(t, resp.Access, "GetFeatureAccess")
		assert.Equal(t, test.customerId, resp.Customer.CustomerId, "GetFeatureAccess")
		assert.Equal(t, test.featureId, resp.Feature.FeatureId, "GetFeatureAccess")
		assert.NotEmpty(t, len(resp.AccessPerSubscription), "GetFeatureAccess")

		subAccess := false
		for _, sa := range resp.AccessPerSubscription {
			if sa.Access {
				subAccess = true
				assert.Equal(t, test.planId, sa.Subscription.Plan.PlanId, "GetFeatureAccess")
			}
		}
		assert.True(t, subAccess, "GetFeatureAccess")
	})

	t.Run("Check feature access that does not exist", func(t *testing.T) {
		req := GetFeatureAccessRequest{
			CustomerId: test.customerId,
			FeatureId:  "feature_00000000000000000000000000000000",
		}
		_, err := test.c.GetFeatureAccess(req)
		assert.NotNil(t, err, "GetFeatureAccess")
		assert.True(t, IsNotFound(err), "GetFeatureAccess")
	})

	t.Run("Check feature access for a customer does not exist", func(t *testing.T) {
		req := GetFeatureAccessRequest{
			CustomerId: fmt.Sprintf("%v", time.Now().Unix()),
			FeatureId:  test.featureId,
		}
		_, err := test.c.GetFeatureAccess(req)
		assert.NotNil(t, err, "GetFeatureAccess")
		assert.True(t, IsNotFound(err), "GetFeatureAccess")
	})

	t.Run("Check metric access", func(t *testing.T) {
		req := GetMetricAccessRequest{
			CustomerId: test.customerId,
			MetricId:   test.metricId,
		}
		resp, err := test.c.GetMetricAccess(req)
		assert.Nil(t, err, "GetMetricAccess")
		assert.True(t, resp.Access, "GetMetricAccess")
		assert.Equal(t, test.customerId, resp.Customer.CustomerId, "GetMetricAccess")
		assert.NotEmpty(t, len(resp.AccessPerSubscription), "GetMetricAccess")

		for _, sa := range resp.AccessPerSubscription {
			assert.Equal(t, test.planId, sa.Subscription.Plan.PlanId, "GetMetricAccess")
		}
	})

	t.Run("Check metric access that does not exist", func(t *testing.T) {
		req := GetMetricAccessRequest{
			CustomerId: test.customerId,
			MetricId:   "metric_00000000000000000000000000000000",
		}
		_, err := test.c.GetMetricAccess(req)
		assert.NotNil(t, err, "GetMetricAccess")
		assert.True(t, IsNotFound(err), "GetMetricAccess")
	})

	t.Run("Check metric access for a customer does not exist", func(t *testing.T) {
		req := GetMetricAccessRequest{
			CustomerId: fmt.Sprintf("%v", time.Now().Unix()),
			MetricId:   test.metricId,
		}
		_, err := test.c.GetMetricAccess(req)
		assert.NotNil(t, err, "GetMetricAccess")
		assert.True(t, IsNotFound(err), "GetMetricAccess")
	})

	// Subscription tests (2)
	// -----------------------------------------------------------------
	t.Run("Update subscription", func(t *testing.T) {
		end := time.Now().AddDate(0, 2, 0)
		req := UpdateSubscriptionRequest{
			SubscriptionId:   test.subscriptionId,
			TurnOffAutoRenew: true,
			EndDate:          &end,
			Metadata: map[string]interface{}{
				"CreatedBy": "SDK Test",
				"UpdatedBy": "SDK Test",
			},
		}
		resp, err := test.c.UpdateSubscription(req)
		assert.Nil(t, err, "UpdateSubscription")
		assert.NotEmpty(t, resp.SubscriptionId, "UpdateSubscription")
		assert.NotEmpty(t, resp.Metadata, "UpdateSubscription")
		assert.NotEmpty(t, resp.StartDate, "UpdateSubscription")
		assert.NotEmpty(t, resp.EndDate, "UpdateSubscription")
		assert.False(t, resp.AutoRenew, "UpdateSubscription")
		assert.Equal(t, test.subscriptionId, resp.SubscriptionId, "UpdateSubscription")
		assert.Equal(t, test.customerId, resp.Customer.CustomerId, "UpdateSubscription")
		assert.Equal(t, test.planId, resp.BillingPlan.PlanId, "UpdateSubscription")
		assert.Contains(t, resp.Metadata, "CreatedBy", "UpdateSubscription")
		assert.Contains(t, resp.Metadata, "UpdatedBy", "UpdateSubscription")
	})

	t.Run("Switch subscription plan", func(t *testing.T) {
		req := SwitchSubscriptionPlanRequest{
			SubscriptionId:    test.subscriptionId,
			SwitchPlanId:      test.plan2Id,
			InvoicingBehavior: AddToNextInvoice,
			UsageBehavior:     TransferToNewSubscription,
			AutoRenew:         true,
			Metadata: map[string]interface{}{
				"CreatedBy": "SDK Test",
				"UpdatedBy": "SDK Test",
			},
		}
		resp, err := test.c.SwitchSubscriptionPlan(req)
		assert.Nil(t, err, "SwitchSubscriptionPlan")
		assert.NotEmpty(t, resp.SubscriptionId, "SwitchSubscriptionPlan")
		assert.NotEmpty(t, resp.Metadata, "SwitchSubscriptionPlan")
		assert.NotEmpty(t, resp.StartDate, "SwitchSubscriptionPlan")
		assert.NotEmpty(t, resp.EndDate, "SwitchSubscriptionPlan")
		assert.False(t, resp.AutoRenew, "SwitchSubscriptionPlan")
		assert.Equal(t, test.subscriptionId, resp.SubscriptionId, "SwitchSubscriptionPlan")
		assert.Equal(t, test.customerId, resp.Customer.CustomerId, "SwitchSubscriptionPlan")
		assert.Equal(t, test.plan2Id, resp.BillingPlan.PlanId, "SwitchSubscriptionPlan")
		assert.Contains(t, resp.Metadata, "CreatedBy", "SwitchSubscriptionPlan")
		assert.Contains(t, resp.Metadata, "UpdatedBy", "SwitchSubscriptionPlan")
	})

	t.Run("Cancel the subscription", func(t *testing.T) {
		req := CancelSubscriptionRequest{
			SubscriptionId:    test.subscriptionId,
			InvoicingBehavior: InvoiceNow,
			UsageBehavior:     BillFull,
			FlatFeeBehavior:   ChargeProrated,
		}
		resp, err := test.c.CancelSubscription(req)
		assert.Nil(t, err, "CancelSubscription")
		assert.NotEmpty(t, resp.SubscriptionId, "CancelSubscription")
	})

	// Network error tests
	// -----------------------------------------------------------------
	t.Run("Ping", func(t *testing.T) {
		resp, err := test.c.Ping()
		assert.Nil(t, err, "Ping")
		assert.NotEmpty(t, resp.OrganizationId, "Ping")
	})

	t.Run("Timeout error", func(t *testing.T) {
		c := NewClient("https://httpbin.org", "").WithTimeout(time.Second).WithDebug(true)
		err := c.post("/delay/2", nil, nil, nil)
		assert.True(t, IsTimeout(err), "Timeout error")
	})

	t.Run("5xx error", func(t *testing.T) {
		c := NewClient("https://httpbin.org", "").WithDebug(true)
		err := c.post("/status/502", nil, nil, nil)
		assert.Contains(t, err.Error(), "server status code: 502", "5xx error")
	})
}
