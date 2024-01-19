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

// Configv1UpdateCollectionResponse configv1 update collection response
//
// swagger:model configv1UpdateCollectionResponse
type Configv1UpdateCollectionResponse struct {

	// collection
	Collection *Configv1Collection `json:"collection,omitempty"`
}

// Validate validates this configv1 update collection response
func (m *Configv1UpdateCollectionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateCollectionResponse) validateCollection(formats strfmt.Registry) error {
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

// ContextValidate validate this configv1 update collection response based on the context it is used
func (m *Configv1UpdateCollectionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateCollectionResponse) contextValidateCollection(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Configv1UpdateCollectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1UpdateCollectionResponse) UnmarshalBinary(b []byte) error {
	var res Configv1UpdateCollectionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}