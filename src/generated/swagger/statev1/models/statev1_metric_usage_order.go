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

// Statev1MetricUsageOrder statev1 metric usage order
//
// swagger:model statev1MetricUsageOrder
type Statev1MetricUsageOrder struct {

	// ascending
	Ascending bool `json:"ascending,omitempty"`

	// by
	By MetricUsageOrderBy `json:"by,omitempty"`
}

// Validate validates this statev1 metric usage order
func (m *Statev1MetricUsageOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Statev1MetricUsageOrder) validateBy(formats strfmt.Registry) error {
	if swag.IsZero(m.By) { // not required
		return nil
	}

	if err := m.By.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("by")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("by")
		}
		return err
	}

	return nil
}

// ContextValidate validate this statev1 metric usage order based on the context it is used
func (m *Statev1MetricUsageOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Statev1MetricUsageOrder) contextValidateBy(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.By) { // not required
		return nil
	}

	if err := m.By.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("by")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("by")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Statev1MetricUsageOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Statev1MetricUsageOrder) UnmarshalBinary(b []byte) error {
	var res Statev1MetricUsageOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
