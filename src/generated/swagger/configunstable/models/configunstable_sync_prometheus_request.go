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

// ConfigunstableSyncPrometheusRequest configunstable sync prometheus request
//
// swagger:model configunstableSyncPrometheusRequest
type ConfigunstableSyncPrometheusRequest struct {

	// rules yaml
	RulesYaml string `json:"rules_yaml,omitempty"`

	// rules yaml gzip
	// Format: byte
	RulesYamlGzip strfmt.Base64 `json:"rules_yaml_gzip,omitempty"`

	// alertmanager yaml
	AlertmanagerYaml string `json:"alertmanager_yaml,omitempty"`

	// alertmanager yaml gzip
	// Format: byte
	AlertmanagerYamlGzip strfmt.Base64 `json:"alertmanager_yaml_gzip,omitempty"`

	// chrono config
	ChronoConfig *SyncPrometheusChronoConfig `json:"chrono_config,omitempty"`

	// dry run
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this configunstable sync prometheus request
func (m *ConfigunstableSyncPrometheusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChronoConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableSyncPrometheusRequest) validateChronoConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ChronoConfig) { // not required
		return nil
	}

	if m.ChronoConfig != nil {
		if err := m.ChronoConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chrono_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chrono_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable sync prometheus request based on the context it is used
func (m *ConfigunstableSyncPrometheusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChronoConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableSyncPrometheusRequest) contextValidateChronoConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ChronoConfig != nil {

		if swag.IsZero(m.ChronoConfig) { // not required
			return nil
		}

		if err := m.ChronoConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chrono_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chrono_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableSyncPrometheusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableSyncPrometheusRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableSyncPrometheusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}