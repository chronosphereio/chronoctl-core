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

// Configv1CreateRollupRuleRequest configv1 create rollup rule request
//
// swagger:model configv1CreateRollupRuleRequest
type Configv1CreateRollupRuleRequest struct {

	// rollup rule
	RollupRule *Configv1RollupRule `json:"rollup_rule,omitempty"`

	// If true, the RollupRule will not be created, and no response RollupRule will be returned. The response will return an error if the given RollupRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this configv1 create rollup rule request
func (m *Configv1CreateRollupRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRollupRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateRollupRuleRequest) validateRollupRule(formats strfmt.Registry) error {
	if swag.IsZero(m.RollupRule) { // not required
		return nil
	}

	if m.RollupRule != nil {
		if err := m.RollupRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollup_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollup_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create rollup rule request based on the context it is used
func (m *Configv1CreateRollupRuleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRollupRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateRollupRuleRequest) contextValidateRollupRule(ctx context.Context, formats strfmt.Registry) error {

	if m.RollupRule != nil {

		if swag.IsZero(m.RollupRule) { // not required
			return nil
		}

		if err := m.RollupRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollup_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollup_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateRollupRuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateRollupRuleRequest) UnmarshalBinary(b []byte) error {
	var res Configv1CreateRollupRuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
