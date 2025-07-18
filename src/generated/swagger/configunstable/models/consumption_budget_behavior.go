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

// ConsumptionBudgetBehavior consumption budget behavior
//
// swagger:model ConsumptionBudgetBehavior
type ConsumptionBudgetBehavior struct {

	// action
	Action BehaviorAction `json:"action,omitempty"`

	// threshold type
	ThresholdType BehaviorThresholdType `json:"threshold_type,omitempty"`

	// instant rate threshold
	InstantRateThreshold *BehaviorInstantRateThreshold `json:"instant_rate_threshold,omitempty"`
}

// Validate validates this consumption budget behavior
func (m *ConsumptionBudgetBehavior) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholdType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstantRateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsumptionBudgetBehavior) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *ConsumptionBudgetBehavior) validateThresholdType(formats strfmt.Registry) error {
	if swag.IsZero(m.ThresholdType) { // not required
		return nil
	}

	if err := m.ThresholdType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("threshold_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("threshold_type")
		}
		return err
	}

	return nil
}

func (m *ConsumptionBudgetBehavior) validateInstantRateThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.InstantRateThreshold) { // not required
		return nil
	}

	if m.InstantRateThreshold != nil {
		if err := m.InstantRateThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instant_rate_threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instant_rate_threshold")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consumption budget behavior based on the context it is used
func (m *ConsumptionBudgetBehavior) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThresholdType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstantRateThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsumptionBudgetBehavior) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *ConsumptionBudgetBehavior) contextValidateThresholdType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ThresholdType) { // not required
		return nil
	}

	if err := m.ThresholdType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("threshold_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("threshold_type")
		}
		return err
	}

	return nil
}

func (m *ConsumptionBudgetBehavior) contextValidateInstantRateThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.InstantRateThreshold != nil {

		if swag.IsZero(m.InstantRateThreshold) { // not required
			return nil
		}

		if err := m.InstantRateThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instant_rate_threshold")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instant_rate_threshold")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsumptionBudgetBehavior) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsumptionBudgetBehavior) UnmarshalBinary(b []byte) error {
	var res ConsumptionBudgetBehavior
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
