// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy trace jaeger remote sampling strategy rate limiting sampling strategy
//
// swagger:model TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy
type TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy struct {

	// Maximum number of traces to sample per second.
	MaxTracesPerSecond int32 `json:"max_traces_per_second,omitempty"`
}

// Validate validates this trace jaeger remote sampling strategy rate limiting sampling strategy
func (m *TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this trace jaeger remote sampling strategy rate limiting sampling strategy based on context it is used
func (m *TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy) UnmarshalBinary(b []byte) error {
	var res TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
