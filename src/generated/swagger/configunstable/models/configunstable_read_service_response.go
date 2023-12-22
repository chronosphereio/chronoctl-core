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

// ConfigunstableReadServiceResponse configunstable read service response
//
// swagger:model configunstableReadServiceResponse
type ConfigunstableReadServiceResponse struct {

	// service
	Service *ConfigunstableService `json:"service,omitempty"`
}

// Validate validates this configunstable read service response
func (m *ConfigunstableReadServiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadServiceResponse) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable read service response based on the context it is used
func (m *ConfigunstableReadServiceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadServiceResponse) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {

		if swag.IsZero(m.Service) { // not required
			return nil
		}

		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableReadServiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableReadServiceResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableReadServiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
