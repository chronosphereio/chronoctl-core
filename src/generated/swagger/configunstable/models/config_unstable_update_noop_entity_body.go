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

// ConfigUnstableUpdateNoopEntityBody config unstable update noop entity body
//
// swagger:model ConfigUnstableUpdateNoopEntityBody
type ConfigUnstableUpdateNoopEntityBody struct {

	// noop entity
	NoopEntity *ConfigunstableNoopEntity `json:"noop_entity,omitempty"`

	// If true, the NoopEntity will be created if it does not already exist, identified by slug. If false, an error will be returned if the NoopEntity does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`
}

// Validate validates this config unstable update noop entity body
func (m *ConfigUnstableUpdateNoopEntityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNoopEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigUnstableUpdateNoopEntityBody) validateNoopEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.NoopEntity) { // not required
		return nil
	}

	if m.NoopEntity != nil {
		if err := m.NoopEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("noop_entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("noop_entity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config unstable update noop entity body based on the context it is used
func (m *ConfigUnstableUpdateNoopEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNoopEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigUnstableUpdateNoopEntityBody) contextValidateNoopEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.NoopEntity != nil {

		if swag.IsZero(m.NoopEntity) { // not required
			return nil
		}

		if err := m.NoopEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("noop_entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("noop_entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigUnstableUpdateNoopEntityBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigUnstableUpdateNoopEntityBody) UnmarshalBinary(b []byte) error {
	var res ConfigUnstableUpdateNoopEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
