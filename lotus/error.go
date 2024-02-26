package lotus

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// Error defines Lotus API response error
// See: https://docs.uselotus.io/errors/error-responses
// https://github.com/uselotus/lotus/blob/main/backend/metering_billing/exceptions/exceptions.py
type Error struct {
	Title            string            `json:"title"`
	Type             string            `json:"type"`
	Detail           string            `json:"detail"`
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
}

type ValidationError struct {
	Code   string `json:"code"`
	Attr   string `json:"attr"`
	Detail string `json:"detail"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v (detail: %v)", e.Title, e.Detail)
}

// IsLotusError checks whether given error is one of Lotus defined types
func IsLotusError(err error) bool {
	var e Error
	ok := errors.As(err, &e)
	return ok
}

// IsNotFound checks whether error is not found error
func IsNotFound(err error) bool {
	var e Error
	ok := errors.As(err, &e)
	return ok && (strings.ToLower(e.Title) == "not_found" ||
		strings.ToLower(e.Title) == "resource_not_found" ||
		strings.ToLower(e.Title) == "does_not_exist")
}

// IsDuplicated checks whether error is duplicate error.
// Possible related resources: webhook, customer, metric, user, organization and so on.
func IsDuplicated(err error) bool {
	var e Error
	ok := errors.As(err, &e)
	return ok && strings.ToLower(e.Title) == "duplicate_resource"
}

// IsInvalidState checks whether error is invalid state error
// Possible related resources: plan, subscription
func IsInvalidState(err error) bool {
	var e Error
	ok := errors.As(err, &e)
	return ok && strings.ToLower(e.Title) == "invalid_state"
}

// IsTimeout checks whether error is timeout error
func IsTimeout(err error) bool {
	var e *url.Error
	ok := errors.As(err, &e)
	return ok && e.Timeout()
}
