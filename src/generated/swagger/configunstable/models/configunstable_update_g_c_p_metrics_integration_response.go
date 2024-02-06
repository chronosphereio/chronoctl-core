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

// ConfigunstableUpdateGCPMetricsIntegrationResponse configunstable update g c p metrics integration response
//
// swagger:model configunstableUpdateGCPMetricsIntegrationResponse
type ConfigunstableUpdateGCPMetricsIntegrationResponse struct {

	// gcp metrics integration
	GcpMetricsIntegration *ConfigunstableGCPMetricsIntegration `json:"gcp_metrics_integration,omitempty"`
}

// Validate validates this configunstable update g c p metrics integration response
func (m *ConfigunstableUpdateGCPMetricsIntegrationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGcpMetricsIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateGCPMetricsIntegrationResponse) validateGcpMetricsIntegration(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpMetricsIntegration) { // not required
		return nil
	}

	if m.GcpMetricsIntegration != nil {
		if err := m.GcpMetricsIntegration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_metrics_integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_metrics_integration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable update g c p metrics integration response based on the context it is used
func (m *ConfigunstableUpdateGCPMetricsIntegrationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGcpMetricsIntegration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateGCPMetricsIntegrationResponse) contextValidateGcpMetricsIntegration(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpMetricsIntegration != nil {

		if swag.IsZero(m.GcpMetricsIntegration) { // not required
			return nil
		}

		if err := m.GcpMetricsIntegration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp_metrics_integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp_metrics_integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableUpdateGCPMetricsIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableUpdateGCPMetricsIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableUpdateGCPMetricsIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
