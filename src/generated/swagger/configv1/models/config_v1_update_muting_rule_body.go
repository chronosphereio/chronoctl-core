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

// ConfigV1UpdateMutingRuleBody config v1 update muting rule body
//
// swagger:model ConfigV1UpdateMutingRuleBody
type ConfigV1UpdateMutingRuleBody struct {

	// muting rule
	MutingRule *Configv1MutingRule `json:"muting_rule,omitempty"`

	// If true, the MutingRule will be created if it does not already exist, identified by slug. If false, an error will be returned if the MutingRule does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the MutingRule isn't created or updated, and no response MutingRule will be returned. The response will return an error if the given MutingRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update muting rule body
func (m *ConfigV1UpdateMutingRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMutingRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateMutingRuleBody) validateMutingRule(formats strfmt.Registry) error {
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

// ContextValidate validate this config v1 update muting rule body based on the context it is used
func (m *ConfigV1UpdateMutingRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMutingRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateMutingRuleBody) contextValidateMutingRule(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigV1UpdateMutingRuleBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateMutingRuleBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateMutingRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
