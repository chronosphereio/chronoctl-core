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

// LogAllocationConfigDefaultDataset Configuration for default dataset.
//
// swagger:model LogAllocationConfigDefaultDataset
type LogAllocationConfigDefaultDataset struct {

	// allocation
	Allocation *LogAllocationConfigAllocation `json:"allocation,omitempty"`

	// priorities
	Priorities *LogAllocationConfigHighLowPriorities `json:"priorities,omitempty"`
}

// Validate validates this log allocation config default dataset
func (m *LogAllocationConfigDefaultDataset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriorities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogAllocationConfigDefaultDataset) validateAllocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Allocation) { // not required
		return nil
	}

	if m.Allocation != nil {
		if err := m.Allocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocation")
			}
			return err
		}
	}

	return nil
}

func (m *LogAllocationConfigDefaultDataset) validatePriorities(formats strfmt.Registry) error {
	if swag.IsZero(m.Priorities) { // not required
		return nil
	}

	if m.Priorities != nil {
		if err := m.Priorities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priorities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priorities")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log allocation config default dataset based on the context it is used
func (m *LogAllocationConfigDefaultDataset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriorities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogAllocationConfigDefaultDataset) contextValidateAllocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Allocation != nil {

		if swag.IsZero(m.Allocation) { // not required
			return nil
		}

		if err := m.Allocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocation")
			}
			return err
		}
	}

	return nil
}

func (m *LogAllocationConfigDefaultDataset) contextValidatePriorities(ctx context.Context, formats strfmt.Registry) error {

	if m.Priorities != nil {

		if swag.IsZero(m.Priorities) { // not required
			return nil
		}

		if err := m.Priorities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priorities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priorities")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogAllocationConfigDefaultDataset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogAllocationConfigDefaultDataset) UnmarshalBinary(b []byte) error {
	var res LogAllocationConfigDefaultDataset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
