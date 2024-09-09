// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricUsageQueryExecutionCountsByType metric usage query execution counts by type
//
// swagger:model MetricUsageQueryExecutionCountsByType
type MetricUsageQueryExecutionCountsByType struct {

	// explorer
	Explorer int32 `json:"explorer,omitempty"`

	// dashboard
	Dashboard int32 `json:"dashboard,omitempty"`

	// external
	External int32 `json:"external,omitempty"`
}

// Validate validates this metric usage query execution counts by type
func (m *MetricUsageQueryExecutionCountsByType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metric usage query execution counts by type based on context it is used
func (m *MetricUsageQueryExecutionCountsByType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricUsageQueryExecutionCountsByType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricUsageQueryExecutionCountsByType) UnmarshalBinary(b []byte) error {
	var res MetricUsageQueryExecutionCountsByType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
