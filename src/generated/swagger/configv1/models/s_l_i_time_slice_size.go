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

// SLITimeSliceSize Defines the supported time slice sizes for TimeSlice SLIs.
//
// swagger:model SLITimeSliceSize
type SLITimeSliceSize string

func NewSLITimeSliceSize(value SLITimeSliceSize) *SLITimeSliceSize {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SLITimeSliceSize.
func (m SLITimeSliceSize) Pointer() *SLITimeSliceSize {
	return &m
}

const (

	// SLITimeSliceSizeTIMESLICESIZEONEMINUTE captures enum value "TIMESLICE_SIZE_ONE_MINUTE"
	SLITimeSliceSizeTIMESLICESIZEONEMINUTE SLITimeSliceSize = "TIMESLICE_SIZE_ONE_MINUTE"

	// SLITimeSliceSizeTIMESLICESIZEFIVEMINUTES captures enum value "TIMESLICE_SIZE_FIVE_MINUTES"
	SLITimeSliceSizeTIMESLICESIZEFIVEMINUTES SLITimeSliceSize = "TIMESLICE_SIZE_FIVE_MINUTES"
)

// for schema
var sLITimeSliceSizeEnum []interface{}

func init() {
	var res []SLITimeSliceSize
	if err := json.Unmarshal([]byte(`["TIMESLICE_SIZE_ONE_MINUTE","TIMESLICE_SIZE_FIVE_MINUTES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sLITimeSliceSizeEnum = append(sLITimeSliceSizeEnum, v)
	}
}

func (m SLITimeSliceSize) validateSLITimeSliceSizeEnum(path, location string, value SLITimeSliceSize) error {
	if err := validate.EnumCase(path, location, value, sLITimeSliceSizeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this s l i time slice size
func (m SLITimeSliceSize) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSLITimeSliceSizeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this s l i time slice size based on context it is used
func (m SLITimeSliceSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
