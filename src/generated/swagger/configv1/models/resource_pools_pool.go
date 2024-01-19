// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourcePoolsPool resource pools pool
//
// swagger:model ResourcePoolsPool
type ResourcePoolsPool struct {

	// Required name of the pool. Must be unique.
	Name string `json:"name,omitempty"`

	// allocation
	Allocation *ResourcePoolsAllocation `json:"allocation,omitempty"`

	// Required filters which define which metrics map to this pool, where any
	// metric which matches at least one filter will map to the pool.
	Filters []*Configv1LabelFilter `json:"filters"`

	// priorities
	Priorities *ResourcePoolsPriorities `json:"priorities,omitempty"`
}

// Validate validates this resource pools pool
func (m *ResourcePoolsPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
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

func (m *ResourcePoolsPool) validateAllocation(formats strfmt.Registry) error {
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

func (m *ResourcePoolsPool) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcePoolsPool) validatePriorities(formats strfmt.Registry) error {
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

// ContextValidate validate this resource pools pool based on the context it is used
func (m *ResourcePoolsPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
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

func (m *ResourcePoolsPool) contextValidateAllocation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ResourcePoolsPool) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {

			if swag.IsZero(m.Filters[i]) { // not required
				return nil
			}

			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcePoolsPool) contextValidatePriorities(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ResourcePoolsPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePoolsPool) UnmarshalBinary(b []byte) error {
	var res ResourcePoolsPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}