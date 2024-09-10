// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigunstableLabel Label is a key, value pair used for filtering.
//
// swagger:model configunstableLabel
type ConfigunstableLabel struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this configunstable label
func (m *ConfigunstableLabel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this configunstable label based on context it is used
func (m *ConfigunstableLabel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableLabel) UnmarshalBinary(b []byte) error {
	var res ConfigunstableLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
