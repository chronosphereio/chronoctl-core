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

// ConfigunstableUpdateTraceTailSamplingRulesRequest configunstable update trace tail sampling rules request
//
// swagger:model configunstableUpdateTraceTailSamplingRulesRequest
type ConfigunstableUpdateTraceTailSamplingRulesRequest struct {

	// trace tail sampling rules
	TraceTailSamplingRules *ConfigunstableTraceTailSamplingRules `json:"trace_tail_sampling_rules,omitempty"`

	// create if missing
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// dry run
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this configunstable update trace tail sampling rules request
func (m *ConfigunstableUpdateTraceTailSamplingRulesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceTailSamplingRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceTailSamplingRulesRequest) validateTraceTailSamplingRules(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceTailSamplingRules) { // not required
		return nil
	}

	if m.TraceTailSamplingRules != nil {
		if err := m.TraceTailSamplingRules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_tail_sampling_rules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_tail_sampling_rules")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable update trace tail sampling rules request based on the context it is used
func (m *ConfigunstableUpdateTraceTailSamplingRulesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceTailSamplingRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceTailSamplingRulesRequest) contextValidateTraceTailSamplingRules(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceTailSamplingRules != nil {

		if swag.IsZero(m.TraceTailSamplingRules) { // not required
			return nil
		}

		if err := m.TraceTailSamplingRules.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_tail_sampling_rules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_tail_sampling_rules")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceTailSamplingRulesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceTailSamplingRulesRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableUpdateTraceTailSamplingRulesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
