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

// ConfigV1UpdateTraceMetricsRuleBody config v1 update trace metrics rule body
//
// swagger:model ConfigV1UpdateTraceMetricsRuleBody
type ConfigV1UpdateTraceMetricsRuleBody struct {

	// trace metrics rule
	TraceMetricsRule *Configv1TraceMetricsRule `json:"trace_metrics_rule,omitempty"`

	// If true, the TraceMetricsRule will be created if it does not already exist, identified by slug. If false, an error will be returned if the TraceMetricsRule does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the TraceMetricsRule will not be created nor updated, and no response TraceMetricsRule will be returned. The response will return an error if the given TraceMetricsRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update trace metrics rule body
func (m *ConfigV1UpdateTraceMetricsRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceMetricsRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateTraceMetricsRuleBody) validateTraceMetricsRule(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceMetricsRule) { // not required
		return nil
	}

	if m.TraceMetricsRule != nil {
		if err := m.TraceMetricsRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_metrics_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_metrics_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update trace metrics rule body based on the context it is used
func (m *ConfigV1UpdateTraceMetricsRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceMetricsRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateTraceMetricsRuleBody) contextValidateTraceMetricsRule(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceMetricsRule != nil {

		if swag.IsZero(m.TraceMetricsRule) { // not required
			return nil
		}

		if err := m.TraceMetricsRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_metrics_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_metrics_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateTraceMetricsRuleBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateTraceMetricsRuleBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateTraceMetricsRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}