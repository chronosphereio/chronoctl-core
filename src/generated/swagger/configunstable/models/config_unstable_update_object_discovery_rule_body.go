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

// ConfigUnstableUpdateObjectDiscoveryRuleBody config unstable update object discovery rule body
//
// swagger:model ConfigUnstableUpdateObjectDiscoveryRuleBody
type ConfigUnstableUpdateObjectDiscoveryRuleBody struct {

	// object discovery rule
	ObjectDiscoveryRule *ConfigunstableObjectDiscoveryRule `json:"object_discovery_rule,omitempty"`

	// If true, the ObjectDiscoveryRule will be created if it does not already exist, identified by slug. If false, an error will be returned if the ObjectDiscoveryRule does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`
}

// Validate validates this config unstable update object discovery rule body
func (m *ConfigUnstableUpdateObjectDiscoveryRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectDiscoveryRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigUnstableUpdateObjectDiscoveryRuleBody) validateObjectDiscoveryRule(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectDiscoveryRule) { // not required
		return nil
	}

	if m.ObjectDiscoveryRule != nil {
		if err := m.ObjectDiscoveryRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_discovery_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object_discovery_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config unstable update object discovery rule body based on the context it is used
func (m *ConfigUnstableUpdateObjectDiscoveryRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectDiscoveryRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigUnstableUpdateObjectDiscoveryRuleBody) contextValidateObjectDiscoveryRule(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectDiscoveryRule != nil {

		if swag.IsZero(m.ObjectDiscoveryRule) { // not required
			return nil
		}

		if err := m.ObjectDiscoveryRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_discovery_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object_discovery_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigUnstableUpdateObjectDiscoveryRuleBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigUnstableUpdateObjectDiscoveryRuleBody) UnmarshalBinary(b []byte) error {
	var res ConfigUnstableUpdateObjectDiscoveryRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}