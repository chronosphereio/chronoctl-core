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

// ConfigV1UpdateLogScaleAlertBody config v1 update log scale alert body
//
// swagger:model ConfigV1UpdateLogScaleAlertBody
type ConfigV1UpdateLogScaleAlertBody struct {

	// log scale alert
	LogScaleAlert *Configv1LogScaleAlert `json:"log_scale_alert,omitempty"`

	// If true, the LogScaleAlert will be created if it does not already exist, identified by slug. If false, an error will be returned if the LogScaleAlert does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the LogScaleAlert isn't created or updated, and no response LogScaleAlert will be returned. The response will return an error if the given LogScaleAlert is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update log scale alert body
func (m *ConfigV1UpdateLogScaleAlertBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogScaleAlert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateLogScaleAlertBody) validateLogScaleAlert(formats strfmt.Registry) error {
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

// ContextValidate validate this config v1 update log scale alert body based on the context it is used
func (m *ConfigV1UpdateLogScaleAlertBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogScaleAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateLogScaleAlertBody) contextValidateLogScaleAlert(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigV1UpdateLogScaleAlertBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateLogScaleAlertBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateLogScaleAlertBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
