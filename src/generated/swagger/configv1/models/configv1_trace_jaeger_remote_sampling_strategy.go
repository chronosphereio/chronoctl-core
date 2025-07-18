// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Configv1TraceJaegerRemoteSamplingStrategy configv1 trace jaeger remote sampling strategy
//
// swagger:model configv1TraceJaegerRemoteSamplingStrategy
type Configv1TraceJaegerRemoteSamplingStrategy struct {

	// Unique identifier of the TraceJaegerRemoteSamplingStrategy. If a `slug` isn't provided, one will be generated based of the `name` field. You can't modify this field after the TraceJaegerRemoteSamplingStrategy is created.
	Slug string `json:"slug,omitempty"`

	// Required. Name of the TraceJaegerRemoteSamplingStrategy. You can modify this value after the TraceJaegerRemoteSamplingStrategy is created.
	Name string `json:"name,omitempty"`

	// Timestamp of when the TraceJaegerRemoteSamplingStrategy was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the TraceJaegerRemoteSamplingStrategy was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// The name of the service this sampling strategy applies to. This must match the slug and name fields.
	ServiceName string `json:"service_name,omitempty"`

	// applied strategy
	AppliedStrategy *TraceJaegerRemoteSamplingStrategyAppliedStrategy `json:"applied_strategy,omitempty"`
}

// Validate validates this configv1 trace jaeger remote sampling strategy
func (m *Configv1TraceJaegerRemoteSamplingStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppliedStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1TraceJaegerRemoteSamplingStrategy) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceJaegerRemoteSamplingStrategy) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceJaegerRemoteSamplingStrategy) validateAppliedStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.AppliedStrategy) { // not required
		return nil
	}

	if m.AppliedStrategy != nil {
		if err := m.AppliedStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applied_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applied_strategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 trace jaeger remote sampling strategy based on the context it is used
func (m *Configv1TraceJaegerRemoteSamplingStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppliedStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1TraceJaegerRemoteSamplingStrategy) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceJaegerRemoteSamplingStrategy) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1TraceJaegerRemoteSamplingStrategy) contextValidateAppliedStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.AppliedStrategy != nil {

		if swag.IsZero(m.AppliedStrategy) { // not required
			return nil
		}

		if err := m.AppliedStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applied_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applied_strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1TraceJaegerRemoteSamplingStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1TraceJaegerRemoteSamplingStrategy) UnmarshalBinary(b []byte) error {
	var res Configv1TraceJaegerRemoteSamplingStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
