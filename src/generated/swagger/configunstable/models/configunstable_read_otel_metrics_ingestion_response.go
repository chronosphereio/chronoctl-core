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

// ConfigunstableReadOtelMetricsIngestionResponse configunstable read otel metrics ingestion response
//
// swagger:model configunstableReadOtelMetricsIngestionResponse
type ConfigunstableReadOtelMetricsIngestionResponse struct {

	// otel metrics ingestion
	OtelMetricsIngestion *ConfigunstableOtelMetricsIngestion `json:"otel_metrics_ingestion,omitempty"`
}

// Validate validates this configunstable read otel metrics ingestion response
func (m *ConfigunstableReadOtelMetricsIngestionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtelMetricsIngestion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadOtelMetricsIngestionResponse) validateOtelMetricsIngestion(formats strfmt.Registry) error {
	if swag.IsZero(m.OtelMetricsIngestion) { // not required
		return nil
	}

	if m.OtelMetricsIngestion != nil {
		if err := m.OtelMetricsIngestion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otel_metrics_ingestion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("otel_metrics_ingestion")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable read otel metrics ingestion response based on the context it is used
func (m *ConfigunstableReadOtelMetricsIngestionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOtelMetricsIngestion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadOtelMetricsIngestionResponse) contextValidateOtelMetricsIngestion(ctx context.Context, formats strfmt.Registry) error {

	if m.OtelMetricsIngestion != nil {

		if swag.IsZero(m.OtelMetricsIngestion) { // not required
			return nil
		}

		if err := m.OtelMetricsIngestion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otel_metrics_ingestion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("otel_metrics_ingestion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableReadOtelMetricsIngestionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableReadOtelMetricsIngestionResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableReadOtelMetricsIngestionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
