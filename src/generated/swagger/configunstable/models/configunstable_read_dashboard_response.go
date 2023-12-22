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

// ConfigunstableReadDashboardResponse configunstable read dashboard response
//
// swagger:model configunstableReadDashboardResponse
type ConfigunstableReadDashboardResponse struct {

	// dashboard
	Dashboard *ConfigunstableDashboard `json:"dashboard,omitempty"`
}

// Validate validates this configunstable read dashboard response
func (m *ConfigunstableReadDashboardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadDashboardResponse) validateDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboard) { // not required
		return nil
	}

	if m.Dashboard != nil {
		if err := m.Dashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dashboard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable read dashboard response based on the context it is used
func (m *ConfigunstableReadDashboardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadDashboardResponse) contextValidateDashboard(ctx context.Context, formats strfmt.Registry) error {

	if m.Dashboard != nil {

		if swag.IsZero(m.Dashboard) { // not required
			return nil
		}

		if err := m.Dashboard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dashboard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableReadDashboardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableReadDashboardResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableReadDashboardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
