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

// ConfigunstableReadConsumptionBudgetResponse configunstable read consumption budget response
//
// swagger:model configunstableReadConsumptionBudgetResponse
type ConfigunstableReadConsumptionBudgetResponse struct {

	// consumption budget
	ConsumptionBudget *ConfigunstableConsumptionBudget `json:"consumption_budget,omitempty"`
}

// Validate validates this configunstable read consumption budget response
func (m *ConfigunstableReadConsumptionBudgetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumptionBudget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadConsumptionBudgetResponse) validateConsumptionBudget(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionBudget) { // not required
		return nil
	}

	if m.ConsumptionBudget != nil {
		if err := m.ConsumptionBudget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumption_budget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumption_budget")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable read consumption budget response based on the context it is used
func (m *ConfigunstableReadConsumptionBudgetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsumptionBudget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadConsumptionBudgetResponse) contextValidateConsumptionBudget(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsumptionBudget != nil {

		if swag.IsZero(m.ConsumptionBudget) { // not required
			return nil
		}

		if err := m.ConsumptionBudget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumption_budget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumption_budget")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableReadConsumptionBudgetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableReadConsumptionBudgetResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableReadConsumptionBudgetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
