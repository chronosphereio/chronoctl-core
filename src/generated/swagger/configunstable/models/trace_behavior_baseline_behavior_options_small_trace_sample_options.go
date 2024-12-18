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

// TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions trace behavior baseline behavior options small trace sample options
//
// swagger:model TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions
type TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions struct {

	// For N = number of spans in the trace, if N <= span_count_threshold, the trace is sampled according to the
	// given sample rate.
	SpanCountThreshold int64 `json:"span_count_threshold,omitempty"`

	// Sample rate.
	SampleRate float64 `json:"sample_rate,omitempty"`

	// sampling type
	SamplingType TraceBehaviorBaselineBehaviorOptionsSamplingType `json:"sampling_type,omitempty"`

	// Whether or not to use these options.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this trace behavior baseline behavior options small trace sample options
func (m *TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamplingType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions) validateSamplingType(formats strfmt.Registry) error {
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

// ContextValidate validate this trace behavior baseline behavior options small trace sample options based on the context it is used
func (m *TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSamplingType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions) contextValidateSamplingType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions) UnmarshalBinary(b []byte) error {
	var res TraceBehaviorBaselineBehaviorOptionsSmallTraceSampleOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
