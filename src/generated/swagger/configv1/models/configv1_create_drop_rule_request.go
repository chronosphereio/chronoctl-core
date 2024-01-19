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

// Configv1CreateDropRuleRequest configv1 create drop rule request
//
// swagger:model configv1CreateDropRuleRequest
type Configv1CreateDropRuleRequest struct {

	// drop rule
	DropRule *Configv1DropRule `json:"drop_rule,omitempty"`

	// If true, the DropRule will not be created, and no response DropRule will be returned. The response will return an error if the given DropRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this configv1 create drop rule request
func (m *Configv1CreateDropRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDropRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateDropRuleRequest) validateDropRule(formats strfmt.Registry) error {
	if swag.IsZero(m.DropRule) { // not required
		return nil
	}

	if m.DropRule != nil {
		if err := m.DropRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("drop_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("drop_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create drop rule request based on the context it is used
func (m *Configv1CreateDropRuleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDropRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateDropRuleRequest) contextValidateDropRule(ctx context.Context, formats strfmt.Registry) error {

	if m.DropRule != nil {

		if swag.IsZero(m.DropRule) { // not required
			return nil
		}

		if err := m.DropRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("drop_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("drop_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateDropRuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateDropRuleRequest) UnmarshalBinary(b []byte) error {
	var res Configv1CreateDropRuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}