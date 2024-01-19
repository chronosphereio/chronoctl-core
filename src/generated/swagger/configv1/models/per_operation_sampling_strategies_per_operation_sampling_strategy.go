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

// PerOperationSamplingStrategiesPerOperationSamplingStrategy per operation sampling strategies per operation sampling strategy
//
// swagger:model PerOperationSamplingStrategiesPerOperationSamplingStrategy
type PerOperationSamplingStrategiesPerOperationSamplingStrategy struct {

	// The operation to which this specific strategy should apply.
	Operation string `json:"operation,omitempty"`

	// probabilistic sampling strategy
	ProbabilisticSamplingStrategy *TraceJaegerRemoteSamplingStrategyProbabilisticStrategy `json:"probabilistic_sampling_strategy,omitempty"`
}

// Validate validates this per operation sampling strategies per operation sampling strategy
func (m *PerOperationSamplingStrategiesPerOperationSamplingStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProbabilisticSamplingStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerOperationSamplingStrategiesPerOperationSamplingStrategy) validateProbabilisticSamplingStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.ProbabilisticSamplingStrategy) { // not required
		return nil
	}

	if m.ProbabilisticSamplingStrategy != nil {
		if err := m.ProbabilisticSamplingStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("probabilistic_sampling_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("probabilistic_sampling_strategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this per operation sampling strategies per operation sampling strategy based on the context it is used
func (m *PerOperationSamplingStrategiesPerOperationSamplingStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProbabilisticSamplingStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerOperationSamplingStrategiesPerOperationSamplingStrategy) contextValidateProbabilisticSamplingStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.ProbabilisticSamplingStrategy != nil {

		if swag.IsZero(m.ProbabilisticSamplingStrategy) { // not required
			return nil
		}

		if err := m.ProbabilisticSamplingStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("probabilistic_sampling_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("probabilistic_sampling_strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerOperationSamplingStrategiesPerOperationSamplingStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerOperationSamplingStrategiesPerOperationSamplingStrategy) UnmarshalBinary(b []byte) error {
	var res PerOperationSamplingStrategiesPerOperationSamplingStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}