// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Configv1DropRule configv1 drop rule
//
// swagger:model configv1DropRule
type Configv1DropRule struct {

	// Unique identifier of the DropRule. If slug is not provided, one will be generated based of the name field. Cannot be modified after the DropRule is created.
	Slug string `json:"slug,omitempty"`

	// Required name of the DropRule. May be modified after the DropRule is created.
	Name string `json:"name,omitempty"`

	// Timestamp of when the DropRule was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the DropRule was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// mode
	Mode Configv1DropRuleMode `json:"mode,omitempty"`

	// Series that match this filter are dropped.
	Filters []*Configv1LabelFilter `json:"filters"`

	// conditional rate based drop
	ConditionalRateBasedDrop *DropRuleConditionalRateBasedDrop `json:"conditional_rate_based_drop,omitempty"`

	// value based drop
	ValueBasedDrop *DropRuleValueBasedDrop `json:"value_based_drop,omitempty"`

	// Drops datapoints if datapoint values are NaN.
	DropNanValue bool `json:"drop_nan_value,omitempty"`
}

// Validate validates this configv1 drop rule
func (m *Configv1DropRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditionalRateBasedDrop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueBasedDrop(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1DropRule) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1DropRule) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1DropRule) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *Configv1DropRule) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1DropRule) validateConditionalRateBasedDrop(formats strfmt.Registry) error {
	if swag.IsZero(m.ConditionalRateBasedDrop) { // not required
		return nil
	}

	if m.ConditionalRateBasedDrop != nil {
		if err := m.ConditionalRateBasedDrop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditional_rate_based_drop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditional_rate_based_drop")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1DropRule) validateValueBasedDrop(formats strfmt.Registry) error {
	if swag.IsZero(m.ValueBasedDrop) { // not required
		return nil
	}

	if m.ValueBasedDrop != nil {
		if err := m.ValueBasedDrop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value_based_drop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value_based_drop")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 drop rule based on the context it is used
func (m *Configv1DropRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditionalRateBasedDrop(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValueBasedDrop(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1DropRule) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1DropRule) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1DropRule) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *Configv1DropRule) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {

			if swag.IsZero(m.Filters[i]) { // not required
				return nil
			}

			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1DropRule) contextValidateConditionalRateBasedDrop(ctx context.Context, formats strfmt.Registry) error {

	if m.ConditionalRateBasedDrop != nil {

		if swag.IsZero(m.ConditionalRateBasedDrop) { // not required
			return nil
		}

		if err := m.ConditionalRateBasedDrop.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditional_rate_based_drop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditional_rate_based_drop")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1DropRule) contextValidateValueBasedDrop(ctx context.Context, formats strfmt.Registry) error {

	if m.ValueBasedDrop != nil {

		if swag.IsZero(m.ValueBasedDrop) { // not required
			return nil
		}

		if err := m.ValueBasedDrop.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value_based_drop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value_based_drop")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1DropRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1DropRule) UnmarshalBinary(b []byte) error {
	var res Configv1DropRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
