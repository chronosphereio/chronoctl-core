// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GroupByKeyGroupByKeyType  - SERVICE: Group by service.
//   - OPERATION: Group by operation.
//   - TAG: Group by span tag.
//
// swagger:model GroupByKeyGroupByKeyType
type GroupByKeyGroupByKeyType string

func NewGroupByKeyGroupByKeyType(value GroupByKeyGroupByKeyType) *GroupByKeyGroupByKeyType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GroupByKeyGroupByKeyType.
func (m GroupByKeyGroupByKeyType) Pointer() *GroupByKeyGroupByKeyType {
	return &m
}

const (

	// GroupByKeyGroupByKeyTypeSERVICE captures enum value "SERVICE"
	GroupByKeyGroupByKeyTypeSERVICE GroupByKeyGroupByKeyType = "SERVICE"

	// GroupByKeyGroupByKeyTypeOPERATION captures enum value "OPERATION"
	GroupByKeyGroupByKeyTypeOPERATION GroupByKeyGroupByKeyType = "OPERATION"

	// GroupByKeyGroupByKeyTypeTAG captures enum value "TAG"
	GroupByKeyGroupByKeyTypeTAG GroupByKeyGroupByKeyType = "TAG"
)

// for schema
var groupByKeyGroupByKeyTypeEnum []interface{}

func init() {
	var res []GroupByKeyGroupByKeyType
	if err := json.Unmarshal([]byte(`["SERVICE","OPERATION","TAG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupByKeyGroupByKeyTypeEnum = append(groupByKeyGroupByKeyTypeEnum, v)
	}
}

func (m GroupByKeyGroupByKeyType) validateGroupByKeyGroupByKeyTypeEnum(path, location string, value GroupByKeyGroupByKeyType) error {
	if err := validate.EnumCase(path, location, value, groupByKeyGroupByKeyTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this group by key group by key type
func (m GroupByKeyGroupByKeyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGroupByKeyGroupByKeyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this group by key group by key type based on context it is used
func (m GroupByKeyGroupByKeyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
