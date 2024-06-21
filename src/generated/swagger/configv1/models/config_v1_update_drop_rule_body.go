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

// ConfigV1UpdateDropRuleBody config v1 update drop rule body
//
// swagger:model ConfigV1UpdateDropRuleBody
type ConfigV1UpdateDropRuleBody struct {

	// drop rule
	DropRule *Configv1DropRule `json:"drop_rule,omitempty"`

	// If true, the DropRule will be created if it does not already exist, identified by slug. If false, an error will be returned if the DropRule does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the DropRule will not be created nor updated, and no response DropRule will be returned. The response will return an error if the given DropRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update drop rule body
func (m *ConfigV1UpdateDropRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDropRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateDropRuleBody) validateDropRule(formats strfmt.Registry) error {
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

// ContextValidate validate this config v1 update drop rule body based on the context it is used
func (m *ConfigV1UpdateDropRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDropRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateDropRuleBody) contextValidateDropRule(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigV1UpdateDropRuleBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateDropRuleBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateDropRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
