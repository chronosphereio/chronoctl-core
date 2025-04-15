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

// Configv1SLO configv1 SLO
//
// swagger:model configv1SLO
type Configv1SLO struct {

	// Unique identifier of the SLO. If slug is not provided, one will be generated based of the name field. Cannot be modified after the SLO is created.
	Slug string `json:"slug,omitempty"`

	// Required name of the SLO. May be modified after the SLO is created.
	Name string `json:"name,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Timestamp of when the SLO was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the SLO was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// collection ref
	CollectionRef *Configv1CollectionReference `json:"collection_ref,omitempty"`

	// Optional notification policy to explicitly apply to the generated monitors.
	// If this is not set then the team this SLO will belong to must have a
	// default notification policy
	NotificationPolicySlug string `json:"notification_policy_slug,omitempty"`

	// signal grouping
	SignalGrouping *MonitorSignalGrouping `json:"signal_grouping,omitempty"`

	// Labels are visible in notifications generated by this SLO,
	// and can be used to route alerts with notification overrides.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are visible in notifications generated by this SLO
	// They can be be templated with labels from notifications.
	Annotations map[string]string `json:"annotations,omitempty"`

	// sli
	Sli *Configv1SLI `json:"sli,omitempty"`

	// definition
	Definition *SLODefinition `json:"definition,omitempty"`
}

// Validate validates this configv1 SLO
func (m *Configv1SLO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectionRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignalGrouping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSli(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1SLO) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1SLO) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1SLO) validateCollectionRef(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectionRef) { // not required
		return nil
	}

	if m.CollectionRef != nil {
		if err := m.CollectionRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collection_ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collection_ref")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1SLO) validateSignalGrouping(formats strfmt.Registry) error {
	if swag.IsZero(m.SignalGrouping) { // not required
		return nil
	}

	if m.SignalGrouping != nil {
		if err := m.SignalGrouping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signal_grouping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signal_grouping")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1SLO) validateSli(formats strfmt.Registry) error {
	if swag.IsZero(m.Sli) { // not required
		return nil
	}

	if m.Sli != nil {
		if err := m.Sli.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sli")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sli")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1SLO) validateDefinition(formats strfmt.Registry) error {
	if swag.IsZero(m.Definition) { // not required
		return nil
	}

	if m.Definition != nil {
		if err := m.Definition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 SLO based on the context it is used
func (m *Configv1SLO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectionRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignalGrouping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSli(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1SLO) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1SLO) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1SLO) contextValidateCollectionRef(ctx context.Context, formats strfmt.Registry) error {

	if m.CollectionRef != nil {

		if swag.IsZero(m.CollectionRef) { // not required
			return nil
		}

		if err := m.CollectionRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collection_ref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collection_ref")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1SLO) contextValidateSignalGrouping(ctx context.Context, formats strfmt.Registry) error {

	if m.SignalGrouping != nil {

		if swag.IsZero(m.SignalGrouping) { // not required
			return nil
		}

		if err := m.SignalGrouping.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signal_grouping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signal_grouping")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1SLO) contextValidateSli(ctx context.Context, formats strfmt.Registry) error {

	if m.Sli != nil {

		if swag.IsZero(m.Sli) { // not required
			return nil
		}

		if err := m.Sli.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sli")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sli")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1SLO) contextValidateDefinition(ctx context.Context, formats strfmt.Registry) error {

	if m.Definition != nil {

		if swag.IsZero(m.Definition) { // not required
			return nil
		}

		if err := m.Definition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1SLO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1SLO) UnmarshalBinary(b []byte) error {
	var res Configv1SLO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
