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

// Configv1CreateLogScaleAlertResponse configv1 create log scale alert response
//
// swagger:model configv1CreateLogScaleAlertResponse
type Configv1CreateLogScaleAlertResponse struct {

	// log scale alert
	LogScaleAlert *Configv1LogScaleAlert `json:"log_scale_alert,omitempty"`
}

// Validate validates this configv1 create log scale alert response
func (m *Configv1CreateLogScaleAlertResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogScaleAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateLogScaleAlertResponse) validateLogScaleAlert(formats strfmt.Registry) error {
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

// ContextValidate validate this configv1 create log scale alert response based on the context it is used
func (m *Configv1CreateLogScaleAlertResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogScaleAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateLogScaleAlertResponse) contextValidateLogScaleAlert(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Configv1CreateLogScaleAlertResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateLogScaleAlertResponse) UnmarshalBinary(b []byte) error {
	var res Configv1CreateLogScaleAlertResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
