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

// ConfigunstableTraceBehavior configunstable trace behavior
//
// swagger:model configunstableTraceBehavior
type ConfigunstableTraceBehavior struct {

	// Required name of the TraceBehavior. May be modified after the TraceBehavior is created.
	Name string `json:"name,omitempty"`

	// Unique identifier of the TraceBehavior. If slug is not provided, one will be generated based of the name field. Cannot be modified after the TraceBehavior is created.
	Slug string `json:"slug,omitempty"`

	// Timestamp of when the TraceBehavior was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the TraceBehavior was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// type
	Type TraceBehaviorTraceBehaviorType `json:"type,omitempty"`

	// baseline behavior options
	BaselineBehaviorOptions *TraceBehaviorBaselineBehaviorOptions `json:"baseline_behavior_options,omitempty"`
}

// Validate validates this configunstable trace behavior
func (m *ConfigunstableTraceBehavior) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaselineBehaviorOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableTraceBehavior) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableTraceBehavior) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableTraceBehavior) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *ConfigunstableTraceBehavior) validateBaselineBehaviorOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.BaselineBehaviorOptions) { // not required
		return nil
	}

	if m.BaselineBehaviorOptions != nil {
		if err := m.BaselineBehaviorOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseline_behavior_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseline_behavior_options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable trace behavior based on the context it is used
func (m *ConfigunstableTraceBehavior) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaselineBehaviorOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableTraceBehavior) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableTraceBehavior) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableTraceBehavior) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *ConfigunstableTraceBehavior) contextValidateBaselineBehaviorOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.BaselineBehaviorOptions != nil {

		if swag.IsZero(m.BaselineBehaviorOptions) { // not required
			return nil
		}

		if err := m.BaselineBehaviorOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseline_behavior_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseline_behavior_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableTraceBehavior) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableTraceBehavior) UnmarshalBinary(b []byte) error {
	var res ConfigunstableTraceBehavior
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}