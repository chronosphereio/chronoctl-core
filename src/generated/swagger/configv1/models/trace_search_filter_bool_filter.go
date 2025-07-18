// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TraceSearchFilterBoolFilter trace search filter bool filter
//
// swagger:model TraceSearchFilterBoolFilter
type TraceSearchFilterBoolFilter struct {

	// The value of the filter compared to the target trace or span field.
	Value bool `json:"value,omitempty"`
}

// Validate validates this trace search filter bool filter
func (m *TraceSearchFilterBoolFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this trace search filter bool filter based on context it is used
func (m *TraceSearchFilterBoolFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TraceSearchFilterBoolFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceSearchFilterBoolFilter) UnmarshalBinary(b []byte) error {
	var res TraceSearchFilterBoolFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
