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

// Configv1TraceTailSamplingRules Root object containing all tail sampling rules (for a tenant).
//
// swagger:model configv1TraceTailSamplingRules
type Configv1TraceTailSamplingRules struct {

	// Set in API responses.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Set in API responses.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Optional. A list of rules, evaluated in order until a match is found,
	// and the sample rate applied, or else the default sample rate is applied.
	Rules []*Configv1TraceTailSamplingRule `json:"rules"`

	// default sample rate
	DefaultSampleRate *Configv1DefaultSampleRate `json:"default_sample_rate,omitempty"`
}

// Validate validates this configv1 trace tail sampling rules
func (m *Configv1TraceTailSamplingRules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultSampleRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1TraceTailSamplingRules) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceTailSamplingRules) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceTailSamplingRules) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1TraceTailSamplingRules) validateDefaultSampleRate(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultSampleRate) { // not required
		return nil
	}

	if m.DefaultSampleRate != nil {
		if err := m.DefaultSampleRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_sample_rate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_sample_rate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 trace tail sampling rules based on the context it is used
func (m *Configv1TraceTailSamplingRules) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultSampleRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1TraceTailSamplingRules) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceTailSamplingRules) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceTailSamplingRules) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {

			if swag.IsZero(m.Rules[i]) { // not required
				return nil
			}

			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1TraceTailSamplingRules) contextValidateDefaultSampleRate(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultSampleRate != nil {

		if swag.IsZero(m.DefaultSampleRate) { // not required
			return nil
		}

		if err := m.DefaultSampleRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_sample_rate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_sample_rate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1TraceTailSamplingRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1TraceTailSamplingRules) UnmarshalBinary(b []byte) error {
	var res Configv1TraceTailSamplingRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
