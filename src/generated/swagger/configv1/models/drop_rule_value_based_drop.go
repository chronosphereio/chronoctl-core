// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DropRuleValueBasedDrop drop rule value based drop
//
// swagger:model DropRuleValueBasedDrop
type DropRuleValueBasedDrop struct {

	// Enables dropping metrics based on a set value.
	Enabled bool `json:"enabled,omitempty"`

	// The target data point value at which to drop metrics.
	TargetDropValue float64 `json:"target_drop_value,omitempty"`
}

// Validate validates this drop rule value based drop
func (m *DropRuleValueBasedDrop) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this drop rule value based drop based on context it is used
func (m *DropRuleValueBasedDrop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DropRuleValueBasedDrop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DropRuleValueBasedDrop) UnmarshalBinary(b []byte) error {
	var res DropRuleValueBasedDrop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
