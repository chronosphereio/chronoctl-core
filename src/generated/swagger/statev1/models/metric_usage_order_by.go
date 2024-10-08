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

// MetricUsageOrderBy metric usage order by
//
// swagger:model MetricUsageOrderBy
type MetricUsageOrderBy string

func NewMetricUsageOrderBy(value MetricUsageOrderBy) *MetricUsageOrderBy {
	return &value
}

// Pointer returns a pointer to a freshly-allocated MetricUsageOrderBy.
func (m MetricUsageOrderBy) Pointer() *MetricUsageOrderBy {
	return &m
}

const (

	// MetricUsageOrderByVALUABLE captures enum value "VALUABLE"
	MetricUsageOrderByVALUABLE MetricUsageOrderBy = "VALUABLE"

	// MetricUsageOrderByDPPS captures enum value "DPPS"
	MetricUsageOrderByDPPS MetricUsageOrderBy = "DPPS"

	// MetricUsageOrderByUTILITY captures enum value "UTILITY"
	MetricUsageOrderByUTILITY MetricUsageOrderBy = "UTILITY"

	// MetricUsageOrderByREFERENCES captures enum value "REFERENCES"
	MetricUsageOrderByREFERENCES MetricUsageOrderBy = "REFERENCES"

	// MetricUsageOrderByEXECUTIONS captures enum value "EXECUTIONS"
	MetricUsageOrderByEXECUTIONS MetricUsageOrderBy = "EXECUTIONS"

	// MetricUsageOrderByUNIQUEVALUES captures enum value "UNIQUE_VALUES"
	MetricUsageOrderByUNIQUEVALUES MetricUsageOrderBy = "UNIQUE_VALUES"

	// MetricUsageOrderByUNIQUEUSERS captures enum value "UNIQUE_USERS"
	MetricUsageOrderByUNIQUEUSERS MetricUsageOrderBy = "UNIQUE_USERS"
)

// for schema
var metricUsageOrderByEnum []interface{}

func init() {
	var res []MetricUsageOrderBy
	if err := json.Unmarshal([]byte(`["VALUABLE","DPPS","UTILITY","REFERENCES","EXECUTIONS","UNIQUE_VALUES","UNIQUE_USERS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricUsageOrderByEnum = append(metricUsageOrderByEnum, v)
	}
}

func (m MetricUsageOrderBy) validateMetricUsageOrderByEnum(path, location string, value MetricUsageOrderBy) error {
	if err := validate.EnumCase(path, location, value, metricUsageOrderByEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this metric usage order by
func (m MetricUsageOrderBy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMetricUsageOrderByEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this metric usage order by based on context it is used
func (m MetricUsageOrderBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
