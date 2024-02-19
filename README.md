# Lotus Golang Library

[![MIT License](https://img.shields.io/badge/License-MIT-red.svg?style=flat)](https://opensource.org/licenses/MIT)
Official Lotus Golang library to capture and send events to any Lotus instance (self-hosted or cloud).

NOTE: This client library is still under development and primary interfaces may change in the future.

## Installation

```bash
go get -u github.com/uselotus/lotus-go
```

In your app, import the lotus library and set your api key **before** making any calls.

```go
package main

import (
	"fmt"
	"github.com/uselotus/lotus-go/lotus"
)

func main() {
	c := lotus.NewClient("<host>", "<key>")
	resp, err := c.Ping()
	fmt.Println(resp, err)
}
```

The Lotus client enables keepalive and retry (when status code >= 500) by default. You can find your key in the
`/settings` page in Lotus. To debug HTTP requests, you can set debug mode:

```go
c := lotus.NewClient("<host>", "<key>").WithDebug(true)
```

You can also set HTTP timeout (defaults to 5 seconds) as well:

```go
c := lotus.NewClient("<host>", "<key>").WithTimeout(time.Second * 5)
```

## Making calls

Please refer to the [Lotus documentation](https://docs.uselotus.io/docs/api/) for more information. Here are some
examples:

### Examples

**Creating New Customer**

```go
package main

import (
	"fmt"
	"github.com/uselotus/lotus-go/lotus"
)

func main() {
	c := lotus.NewClient("<host>", "<key>")
	req := lotus.CreateCustomerRequest{
		CustomerId:          "<customer>",
		CustomerName:        "<customer>",
		DefaultCurrencyCode: "USD",
		Email:               "<customer>@sample.com",
	}
	resp, err := c.CreateCustomer(req)
	fmt.Println(resp, err)
}

```

**Listing Plans**

```go
package main

import (
	"fmt"
	"github.com/uselotus/lotus-go/lotus"
)

func main() {
	c := lotus.NewClient("<host>", "<key>")
	req := lotus.ListPlansRequest{
		Duration: lotus.DurationMonthly,
	}
	resp, err := c.ListPlans(req)
	fmt.Println(resp, err)
}
```

**Creating New Subscription**

```go
package main

import (
	"fmt"
	"github.com/uselotus/lotus-go/lotus"
	"time"
)

func main() {
	c := lotus.NewClient("<host>", "<key>")
	req := lotus.CreateSubscriptionRequest{
		CustomerId: "<customer>",
		PlanId:     "<plan>",
		AutoRenew:  true,
		StartDate:  time.Now(),
		Metadata: map[string]interface{}{
			"CreatedBy": "SDK",
		},
	}
	resp, err := c.CreateSubscription(req)
	fmt.Println(resp, err)
}

```

**Tracking Events**

```go
package main

import (
	"fmt"
	"github.com/uselotus/lotus-go/lotus"
	"time"
)

func main() {
	c := lotus.NewClient("<host>", "<key>")
	req := lotus.TrackEventsRequest{
		Batch: []*lotus.Event{
			{
				CustomerId:    "<customer>",
				EventName:     "<event>",
				IdempotencyId: "event_foo",
				TimeCreated:   time.Now(),
			},
			{
				CustomerId:    "<customer>",
				EventName:     "<event>",
				IdempotencyId: "event_bar",
				TimeCreated:   time.Now(),
			},
		},
	}
	resp, err := c.TrackEvents(req)
	fmt.Println(resp, err)
}

```

## Error Handling

We have declared `lotus.Error` type to represent Lotus API errors with some helper functions:

- `IsLotusError(err error) bool`: Checks whether given error is one of Lotus defined types
- `IsNotFound(err error) bool`: Checks whether error is not found error
- `IsDuplicated(err error) bool`: Checks whether error is duplicate error
- `IsTimeout(err error) bool`: Checks whether error is timeout error

```go
package lotus

type Error struct {
	Title            string            `json:"title"`
	Type             string            `json:"type"`
	Detail           string            `json:"detail"`
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
}

```

## Currently Supported Methods

1. Tracking:
    - [x] Track Event
2. Customers
    - [x] List Customers
    - [x] Get Customer
    - [x] Create Customer (Server issue found)
3. Credits
    - [ ] List Credits
    - [ ] Create Credit
    - [ ] Update Credit
    - [ ] Void Credit
4. Subscriptions
    - [x] List Subscriptions
    - [x] Create Subscription
    - [x] Cancel Subscription
    - [x] Switch a Subscription's plan (Server issue found)
    - [x] Update Subscription
5. Access Management
    - [x] Get Feature Access
    - [x] Get Metric Access
6. Plans
    - [x] List Plans (Missing some query filters)
    - [x] Get Plan
7. Add-ons
    - [x] Attach Add-on
    - [x] Cancel Add-on
