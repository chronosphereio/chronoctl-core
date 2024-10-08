// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricUsageReferenceCountsByType metric usage reference counts by type
//
// swagger:model MetricUsageReferenceCountsByType
type MetricUsageReferenceCountsByType struct {

	// dashboards
	Dashboards int32 `json:"dashboards,omitempty"`

	// monitors
	Monitors int32 `json:"monitors,omitempty"`

	// recording rules
	RecordingRules int32 `json:"recording_rules,omitempty"`

	// drop rules
	DropRules int32 `json:"drop_rules,omitempty"`

	// aggregation rules
	AggregationRules int32 `json:"aggregation_rules,omitempty"`
}

// Validate validates this metric usage reference counts by type
func (m *MetricUsageReferenceCountsByType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metric usage reference counts by type based on context it is used
func (m *MetricUsageReferenceCountsByType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricUsageReferenceCountsByType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricUsageReferenceCountsByType) UnmarshalBinary(b []byte) error {
	var res MetricUsageReferenceCountsByType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
