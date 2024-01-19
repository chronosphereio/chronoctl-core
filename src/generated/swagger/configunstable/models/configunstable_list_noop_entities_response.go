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
)

// ConfigunstableListNoopEntitiesResponse configunstable list noop entities response
//
// swagger:model configunstableListNoopEntitiesResponse
type ConfigunstableListNoopEntitiesResponse struct {

	// noop entities
	NoopEntities []*ConfigunstableNoopEntity `json:"noop_entities"`

	// page
	Page *Configv1PageResult `json:"page,omitempty"`
}

// Validate validates this configunstable list noop entities response
func (m *ConfigunstableListNoopEntitiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNoopEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableListNoopEntitiesResponse) validateNoopEntities(formats strfmt.Registry) error {
	if swag.IsZero(m.NoopEntities) { // not required
		return nil
	}

	for i := 0; i < len(m.NoopEntities); i++ {
		if swag.IsZero(m.NoopEntities[i]) { // not required
			continue
		}

		if m.NoopEntities[i] != nil {
			if err := m.NoopEntities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("noop_entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("noop_entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListNoopEntitiesResponse) validatePage(formats strfmt.Registry) error {
	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if m.Page != nil {
		if err := m.Page.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable list noop entities response based on the context it is used
func (m *ConfigunstableListNoopEntitiesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNoopEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableListNoopEntitiesResponse) contextValidateNoopEntities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NoopEntities); i++ {

		if m.NoopEntities[i] != nil {

			if swag.IsZero(m.NoopEntities[i]) { // not required
				return nil
			}

			if err := m.NoopEntities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("noop_entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("noop_entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListNoopEntitiesResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	if m.Page != nil {

		if swag.IsZero(m.Page) { // not required
			return nil
		}

		if err := m.Page.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableListNoopEntitiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableListNoopEntitiesResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableListNoopEntitiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}