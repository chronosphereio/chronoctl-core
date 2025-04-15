// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogControlRuleDropField DropField is the configuration for a drop field action.
//
// swagger:model LogControlRuleDropField
type LogControlRuleDropField struct {

	// Regular expression to match the field(s) to drop.
	FieldRegex string `json:"field_regex,omitempty"`

	// Fully specified path to the the field(s) to drop.
	// If empty, the field(s) to drop are at the root level of the log.
	ParentPath []string `json:"parent_path"`
}

// Validate validates this log control rule drop field
func (m *LogControlRuleDropField) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log control rule drop field based on context it is used
func (m *LogControlRuleDropField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogControlRuleDropField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogControlRuleDropField) UnmarshalBinary(b []byte) error {
	var res LogControlRuleDropField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
