// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScheduleDayTimeRange schedule day time range
//
// swagger:model ScheduleDayTimeRange
type ScheduleDayTimeRange struct {

	// Start time in the in format `"<hour>:<minute>"`. For example, `"15:30"`.
	StartHhMm string `json:"start_hh_mm,omitempty"`

	// End time in the in format `"<hour>:<minute>"`. For example, `"15:30"`.
	EndHhMm string `json:"end_hh_mm,omitempty"`
}

// Validate validates this schedule day time range
func (m *ScheduleDayTimeRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schedule day time range based on context it is used
func (m *ScheduleDayTimeRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleDayTimeRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleDayTimeRange) UnmarshalBinary(b []byte) error {
	var res ScheduleDayTimeRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
