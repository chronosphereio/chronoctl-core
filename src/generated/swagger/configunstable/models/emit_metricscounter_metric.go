// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmitMetricscounterMetric emit metricscounter metric
//
// swagger:model EmitMetricscounterMetric
type EmitMetricscounterMetric struct {

	// Label field is the label field to count.
	LabelField string `json:"label_field,omitempty"`
}

// Validate validates this emit metricscounter metric
func (m *EmitMetricscounterMetric) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this emit metricscounter metric based on context it is used
func (m *EmitMetricscounterMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmitMetricscounterMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmitMetricscounterMetric) UnmarshalBinary(b []byte) error {
	var res EmitMetricscounterMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}