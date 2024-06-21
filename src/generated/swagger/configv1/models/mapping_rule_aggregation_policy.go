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

// MappingRuleAggregationPolicy mapping rule aggregation policy
//
// swagger:model MappingRuleAggregationPolicy
type MappingRuleAggregationPolicy struct {

	// aggregation
	Aggregation Configv1AggregationType `json:"aggregation,omitempty"`

	// storage policy
	StoragePolicy *Configv1MappingRuleStoragePolicy `json:"storage_policy,omitempty"`

	// Interval between aggregated data points, equivalent to the resolution
	// field in storage policy. If set, then the storage_policy field can't be
	// set.
	Interval string `json:"interval,omitempty"`

	// Deprecated: This field is no longer supported.
	DropTimestamp bool `json:"drop_timestamp,omitempty"`
}

// Validate validates this mapping rule aggregation policy
func (m *MappingRuleAggregationPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MappingRuleAggregationPolicy) validateAggregation(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregation) { // not required
		return nil
	}

	if err := m.Aggregation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aggregation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("aggregation")
		}
		return err
	}

	return nil
}

func (m *MappingRuleAggregationPolicy) validateStoragePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.StoragePolicy) { // not required
		return nil
	}

	if m.StoragePolicy != nil {
		if err := m.StoragePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mapping rule aggregation policy based on the context it is used
func (m *MappingRuleAggregationPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MappingRuleAggregationPolicy) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Aggregation) { // not required
		return nil
	}

	if err := m.Aggregation.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aggregation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("aggregation")
		}
		return err
	}

	return nil
}

func (m *MappingRuleAggregationPolicy) contextValidateStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.StoragePolicy != nil {

		if swag.IsZero(m.StoragePolicy) { // not required
			return nil
		}

		if err := m.StoragePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MappingRuleAggregationPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MappingRuleAggregationPolicy) UnmarshalBinary(b []byte) error {
	var res MappingRuleAggregationPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
