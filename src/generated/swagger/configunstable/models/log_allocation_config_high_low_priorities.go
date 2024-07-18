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

// LogAllocationConfigHighLowPriorities HighLowPriorities defines explicit high and low priority match criteria to define which logs
// should be dropped first (low) and dropped last (high). Anything not matched by either set of rules
// is considered default priority and thus dropped after low but before high.
//
// swagger:model LogAllocationConfigHighLowPriorities
type LogAllocationConfigHighLowPriorities struct {

	// A list of search filters defining which logs are considered high priority in this dataset.
	// The filters are ORed together so only one given filter needs to match.
	HighPriorityFilters []*Configv1LogSearchFilter `json:"high_priority_filters"`

	// A list of search filters defining which logs are considered low priority in this dataset.
	// The filters are ORed together so only one given filter needs to match.
	LowPriorityFilters []*Configv1LogSearchFilter `json:"low_priority_filters"`
}

// Validate validates this log allocation config high low priorities
func (m *LogAllocationConfigHighLowPriorities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHighPriorityFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLowPriorityFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogAllocationConfigHighLowPriorities) validateHighPriorityFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.HighPriorityFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.HighPriorityFilters); i++ {
		if swag.IsZero(m.HighPriorityFilters[i]) { // not required
			continue
		}

		if m.HighPriorityFilters[i] != nil {
			if err := m.HighPriorityFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("high_priority_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("high_priority_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogAllocationConfigHighLowPriorities) validateLowPriorityFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.LowPriorityFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.LowPriorityFilters); i++ {
		if swag.IsZero(m.LowPriorityFilters[i]) { // not required
			continue
		}

		if m.LowPriorityFilters[i] != nil {
			if err := m.LowPriorityFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("low_priority_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("low_priority_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this log allocation config high low priorities based on the context it is used
func (m *LogAllocationConfigHighLowPriorities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHighPriorityFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLowPriorityFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogAllocationConfigHighLowPriorities) contextValidateHighPriorityFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HighPriorityFilters); i++ {

		if m.HighPriorityFilters[i] != nil {

			if swag.IsZero(m.HighPriorityFilters[i]) { // not required
				return nil
			}

			if err := m.HighPriorityFilters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("high_priority_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("high_priority_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogAllocationConfigHighLowPriorities) contextValidateLowPriorityFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LowPriorityFilters); i++ {

		if m.LowPriorityFilters[i] != nil {

			if swag.IsZero(m.LowPriorityFilters[i]) { // not required
				return nil
			}

			if err := m.LowPriorityFilters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("low_priority_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("low_priority_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogAllocationConfigHighLowPriorities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogAllocationConfigHighLowPriorities) UnmarshalBinary(b []byte) error {
	var res LogAllocationConfigHighLowPriorities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
