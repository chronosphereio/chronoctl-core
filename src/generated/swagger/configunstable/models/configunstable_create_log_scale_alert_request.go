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

// ConfigunstableCreateLogScaleAlertRequest configunstable create log scale alert request
//
// swagger:model configunstableCreateLogScaleAlertRequest
type ConfigunstableCreateLogScaleAlertRequest struct {

	// log scale alert
	LogScaleAlert *ConfigunstableLogScaleAlert `json:"log_scale_alert,omitempty"`

	// If true, the LogScaleAlert will not be created, and no response LogScaleAlert will be returned. The response will return an error if the given LogScaleAlert is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this configunstable create log scale alert request
func (m *ConfigunstableCreateLogScaleAlertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogScaleAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLogScaleAlertRequest) validateLogScaleAlert(formats strfmt.Registry) error {
	if swag.IsZero(m.LogScaleAlert) { // not required
		return nil
	}

	if m.LogScaleAlert != nil {
		if err := m.LogScaleAlert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_scale_alert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_scale_alert")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create log scale alert request based on the context it is used
func (m *ConfigunstableCreateLogScaleAlertRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogScaleAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLogScaleAlertRequest) contextValidateLogScaleAlert(ctx context.Context, formats strfmt.Registry) error {

	if m.LogScaleAlert != nil {

		if swag.IsZero(m.LogScaleAlert) { // not required
			return nil
		}

		if err := m.LogScaleAlert.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_scale_alert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_scale_alert")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateLogScaleAlertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateLogScaleAlertRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateLogScaleAlertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
