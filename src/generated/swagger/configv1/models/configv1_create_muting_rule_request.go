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

// Configv1CreateMutingRuleRequest configv1 create muting rule request
//
// swagger:model configv1CreateMutingRuleRequest
type Configv1CreateMutingRuleRequest struct {

	// muting rule
	MutingRule *Configv1MutingRule `json:"muting_rule,omitempty"`

	// If true, the MutingRule will not be created, and no response MutingRule will be returned. The response will return an error if the given MutingRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this configv1 create muting rule request
func (m *Configv1CreateMutingRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMutingRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateMutingRuleRequest) validateMutingRule(formats strfmt.Registry) error {
	if swag.IsZero(m.MutingRule) { // not required
		return nil
	}

	if m.MutingRule != nil {
		if err := m.MutingRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("muting_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("muting_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create muting rule request based on the context it is used
func (m *Configv1CreateMutingRuleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMutingRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateMutingRuleRequest) contextValidateMutingRule(ctx context.Context, formats strfmt.Registry) error {

	if m.MutingRule != nil {

		if swag.IsZero(m.MutingRule) { // not required
			return nil
		}

		if err := m.MutingRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("muting_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("muting_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateMutingRuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateMutingRuleRequest) UnmarshalBinary(b []byte) error {
	var res Configv1CreateMutingRuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
