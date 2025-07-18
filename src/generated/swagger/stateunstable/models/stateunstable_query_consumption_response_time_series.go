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

// StateunstableQueryConsumptionResponseTimeSeries stateunstable query consumption response time series
//
// swagger:model stateunstableQueryConsumptionResponseTimeSeries
type StateunstableQueryConsumptionResponseTimeSeries struct {

	// datapoints
	Datapoints []*StateunstableConsumptionDatapoint `json:"datapoints"`
}

// Validate validates this stateunstable query consumption response time series
func (m *StateunstableQueryConsumptionResponseTimeSeries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatapoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateunstableQueryConsumptionResponseTimeSeries) validateDatapoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Datapoints) { // not required
		return nil
	}

	for i := 0; i < len(m.Datapoints); i++ {
		if swag.IsZero(m.Datapoints[i]) { // not required
			continue
		}

		if m.Datapoints[i] != nil {
			if err := m.Datapoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datapoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datapoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stateunstable query consumption response time series based on the context it is used
func (m *StateunstableQueryConsumptionResponseTimeSeries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatapoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateunstableQueryConsumptionResponseTimeSeries) contextValidateDatapoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Datapoints); i++ {

		if m.Datapoints[i] != nil {

			if swag.IsZero(m.Datapoints[i]) { // not required
				return nil
			}

			if err := m.Datapoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datapoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datapoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateunstableQueryConsumptionResponseTimeSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateunstableQueryConsumptionResponseTimeSeries) UnmarshalBinary(b []byte) error {
	var res StateunstableQueryConsumptionResponseTimeSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
