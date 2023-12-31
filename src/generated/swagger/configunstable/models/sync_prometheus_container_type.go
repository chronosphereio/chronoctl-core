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

// SyncPrometheusContainerType TODO: confirm naming with InfoModel team
//
// swagger:model SyncPrometheusContainerType
type SyncPrometheusContainerType string

func NewSyncPrometheusContainerType(value SyncPrometheusContainerType) *SyncPrometheusContainerType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SyncPrometheusContainerType.
func (m SyncPrometheusContainerType) Pointer() *SyncPrometheusContainerType {
	return &m
}

const (

	// SyncPrometheusContainerTypeCOLLECTIONS captures enum value "COLLECTIONS"
	SyncPrometheusContainerTypeCOLLECTIONS SyncPrometheusContainerType = "COLLECTIONS"

	// SyncPrometheusContainerTypeBUCKETS captures enum value "BUCKETS"
	SyncPrometheusContainerTypeBUCKETS SyncPrometheusContainerType = "BUCKETS"
)

// for schema
var syncPrometheusContainerTypeEnum []interface{}

func init() {
	var res []SyncPrometheusContainerType
	if err := json.Unmarshal([]byte(`["COLLECTIONS","BUCKETS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncPrometheusContainerTypeEnum = append(syncPrometheusContainerTypeEnum, v)
	}
}

func (m SyncPrometheusContainerType) validateSyncPrometheusContainerTypeEnum(path, location string, value SyncPrometheusContainerType) error {
	if err := validate.EnumCase(path, location, value, syncPrometheusContainerTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sync prometheus container type
func (m SyncPrometheusContainerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSyncPrometheusContainerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sync prometheus container type based on context it is used
func (m SyncPrometheusContainerType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
