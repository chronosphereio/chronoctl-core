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

// TraceBehaviorBaselineBehaviorOptionsFastSampleOptions trace behavior baseline behavior options fast sample options
//
// swagger:model TraceBehaviorBaselineBehaviorOptionsFastSampleOptions
type TraceBehaviorBaselineBehaviorOptionsFastSampleOptions struct {

	// Duration in seconds under which traces are sampled
	// according to the given sample rate.
	MaxDurationSeconds float64 `json:"max_duration_seconds,omitempty"`

	// Sample rate for traces under the given duration.
	SampleRate float64 `json:"sample_rate,omitempty"`

	// sampling type
	SamplingType TraceBehaviorBaselineBehaviorOptionsSamplingType `json:"sampling_type,omitempty"`

	// Whether or not to use these options.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this trace behavior baseline behavior options fast sample options
func (m *TraceBehaviorBaselineBehaviorOptionsFastSampleOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamplingType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptionsFastSampleOptions) validateSamplingType(formats strfmt.Registry) error {
	if swag.IsZero(m.SamplingType) { // not required
		return nil
	}

	if err := m.SamplingType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sampling_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sampling_type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this trace behavior baseline behavior options fast sample options based on the context it is used
func (m *TraceBehaviorBaselineBehaviorOptionsFastSampleOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSamplingType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptionsFastSampleOptions) contextValidateSamplingType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SamplingType) { // not required
		return nil
	}

	if err := m.SamplingType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sampling_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sampling_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraceBehaviorBaselineBehaviorOptionsFastSampleOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceBehaviorBaselineBehaviorOptionsFastSampleOptions) UnmarshalBinary(b []byte) error {
	var res TraceBehaviorBaselineBehaviorOptionsFastSampleOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
