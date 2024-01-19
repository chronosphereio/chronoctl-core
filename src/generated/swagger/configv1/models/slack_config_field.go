// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SlackConfigField slack config field
//
// swagger:model SlackConfigField
type SlackConfigField struct {

	// title
	Title string `json:"title,omitempty"`

	// value
	Value string `json:"value,omitempty"`

	// short
	Short bool `json:"short,omitempty"`
}

// Validate validates this slack config field
func (m *SlackConfigField) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this slack config field based on context it is used
func (m *SlackConfigField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SlackConfigField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlackConfigField) UnmarshalBinary(b []byte) error {
	var res SlackConfigField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}