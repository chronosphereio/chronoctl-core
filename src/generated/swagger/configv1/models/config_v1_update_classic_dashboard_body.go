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

// ConfigV1UpdateClassicDashboardBody config v1 update classic dashboard body
//
// swagger:model ConfigV1UpdateClassicDashboardBody
type ConfigV1UpdateClassicDashboardBody struct {

	// classic dashboard
	ClassicDashboard *Configv1GrafanaDashboard `json:"classic_dashboard,omitempty"`

	// If true, the GrafanaDashboard will be created if it does not already exist, identified by slug. If false, an error will be returned if the GrafanaDashboard does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the GrafanaDashboard will not be created nor updated, and no response GrafanaDashboard will be returned. The response will return an error if the given GrafanaDashboard is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update classic dashboard body
func (m *ConfigV1UpdateClassicDashboardBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassicDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateClassicDashboardBody) validateClassicDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.ClassicDashboard) { // not required
		return nil
	}

	if m.ClassicDashboard != nil {
		if err := m.ClassicDashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classic_dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("classic_dashboard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update classic dashboard body based on the context it is used
func (m *ConfigV1UpdateClassicDashboardBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClassicDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateClassicDashboardBody) contextValidateClassicDashboard(ctx context.Context, formats strfmt.Registry) error {

	if m.ClassicDashboard != nil {

		if swag.IsZero(m.ClassicDashboard) { // not required
			return nil
		}

		if err := m.ClassicDashboard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classic_dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("classic_dashboard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateClassicDashboardBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateClassicDashboardBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateClassicDashboardBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
