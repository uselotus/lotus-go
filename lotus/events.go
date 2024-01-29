package lotus

import (
	"time"
)

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

// TrackEvents ingests events into Lotus
// See: https://docs.uselotus.io/api-reference/events/tracking
func (c *Client) TrackEvents(req TrackEventsRequest) (resp *TrackEventsResponse, err error) {
	resp = new(TrackEventsResponse)
	err = c.post("/api/track/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VerifyEventIngestion verifies if events with specific idempotency events have been ingested into Lotus
// See: https://docs.uselotus.io/api-reference/events/verify-ingestion
func (c *Client) VerifyEventIngestion(req VerifyEventIngestionRequest) (resp *VerifyEventIngestionResponse, err error) {
	resp = new(VerifyEventIngestionResponse)
	err = c.post("/api/verify_idems_received/", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
