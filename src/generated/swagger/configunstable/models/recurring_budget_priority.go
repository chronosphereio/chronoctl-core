// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecurringBudgetPriority recurring budget priority
//
// swagger:model RecurringBudgetPriority
type RecurringBudgetPriority struct {

	// dataset_slug is the required slug of the priority's dataset. Priority
	// datasets are implicitly AND'd with the budget dataset. Must match the
	// global budget's telemetry type (e.g. logs, metrics etc).
	DatasetSlug string `json:"dataset_slug,omitempty"`

	// priority is the required priority of the dataset, where priority=1 is
	// the highest priority, and priority=64 is the lowest priority.
	Priority int32 `json:"priority,omitempty"`
}

// Validate validates this recurring budget priority
func (m *RecurringBudgetPriority) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this recurring budget priority based on context it is used
func (m *RecurringBudgetPriority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecurringBudgetPriority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecurringBudgetPriority) UnmarshalBinary(b []byte) error {
	var res RecurringBudgetPriority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
