// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObjectDiscoveryRuleMetricRule object discovery rule metric rule
//
// swagger:model ObjectDiscoveryRuleMetricRule
type ObjectDiscoveryRuleMetricRule struct {

	// expr
	Expr string `json:"expr,omitempty"`

	// object labels
	ObjectLabels []string `json:"objectLabels"`

	// object mapping label
	ObjectMappingLabel string `json:"objectMappingLabel,omitempty"`
}

// Validate validates this object discovery rule metric rule
func (m *ObjectDiscoveryRuleMetricRule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this object discovery rule metric rule based on context it is used
func (m *ObjectDiscoveryRuleMetricRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectDiscoveryRuleMetricRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectDiscoveryRuleMetricRule) UnmarshalBinary(b []byte) error {
	var res ObjectDiscoveryRuleMetricRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
