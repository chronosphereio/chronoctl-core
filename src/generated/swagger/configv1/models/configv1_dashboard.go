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

// Configv1Dashboard configv1 dashboard
//
// swagger:model configv1Dashboard
type Configv1Dashboard struct {

	// Unique identifier of the Dashboard. If slug is not provided, one will be generated based of the name field. Cannot be modified after the Dashboard is created.
	Slug string `json:"slug,omitempty"`

	// Required name of the Dashboard. May be modified after the Dashboard is created.
	Name string `json:"name,omitempty"`

	// Timestamp of when the Dashboard was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the Dashboard was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Optional slug of the collection the dashboard belongs to.
	CollectionSlug string `json:"collection_slug,omitempty"`

	// collection
	Collection *Configv1CollectionReference `json:"collection,omitempty"`

	// Required raw JSON of the dashboard.
	DashboardJSON string `json:"dashboard_json,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`
}

// Validate validates this configv1 dashboard
func (m *Configv1Dashboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1Dashboard) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1Dashboard) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1Dashboard) validateCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.Collection) { // not required
		return nil
	}

	if m.Collection != nil {
		if err := m.Collection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 dashboard based on the context it is used
func (m *Configv1Dashboard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1Dashboard) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1Dashboard) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1Dashboard) contextValidateCollection(ctx context.Context, formats strfmt.Registry) error {

	if m.Collection != nil {

		if swag.IsZero(m.Collection) { // not required
			return nil
		}

		if err := m.Collection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1Dashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1Dashboard) UnmarshalBinary(b []byte) error {
	var res Configv1Dashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
