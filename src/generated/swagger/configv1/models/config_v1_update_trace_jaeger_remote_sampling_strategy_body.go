// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody config v1 update trace jaeger remote sampling strategy body
//
// swagger:model ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody
type ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody struct {

	// trace jaeger remote sampling strategy
	TraceJaegerRemoteSamplingStrategy *Configv1TraceJaegerRemoteSamplingStrategy `json:"trace_jaeger_remote_sampling_strategy,omitempty"`

	// If true, the TraceJaegerRemoteSamplingStrategy will be created if it does not already exist, identified by slug. If false, an error will be returned if the TraceJaegerRemoteSamplingStrategy does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the TraceJaegerRemoteSamplingStrategy will not be created nor updated, and no response TraceJaegerRemoteSamplingStrategy will be returned. The response will return an error if the given TraceJaegerRemoteSamplingStrategy is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update trace jaeger remote sampling strategy body
func (m *ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceJaegerRemoteSamplingStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody) validateTraceJaegerRemoteSamplingStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceJaegerRemoteSamplingStrategy) { // not required
		return nil
	}

	if m.TraceJaegerRemoteSamplingStrategy != nil {
		if err := m.TraceJaegerRemoteSamplingStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_jaeger_remote_sampling_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_jaeger_remote_sampling_strategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update trace jaeger remote sampling strategy body based on the context it is used
func (m *ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceJaegerRemoteSamplingStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody) contextValidateTraceJaegerRemoteSamplingStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceJaegerRemoteSamplingStrategy != nil {

		if swag.IsZero(m.TraceJaegerRemoteSamplingStrategy) { // not required
			return nil
		}

		if err := m.TraceJaegerRemoteSamplingStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_jaeger_remote_sampling_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_jaeger_remote_sampling_strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateTraceJaegerRemoteSamplingStrategyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
