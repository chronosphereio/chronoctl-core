// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatasetDatasetConfiguration dataset dataset configuration
//
// swagger:model DatasetDatasetConfiguration
type DatasetDatasetConfiguration struct {

	// type
	Type DatasetDatasetType `json:"type,omitempty"`

	// trace dataset
	TraceDataset *ConfigunstableTraceDataset `json:"trace_dataset,omitempty"`

	// log dataset
	LogDataset *ConfigunstableLogDataset `json:"log_dataset,omitempty"`
}

// Validate validates this dataset dataset configuration
func (m *DatasetDatasetConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceDataset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogDataset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasetDatasetConfiguration) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *DatasetDatasetConfiguration) validateTraceDataset(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceDataset) { // not required
		return nil
	}

	if m.TraceDataset != nil {
		if err := m.TraceDataset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_dataset")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetDatasetConfiguration) validateLogDataset(formats strfmt.Registry) error {
	if swag.IsZero(m.LogDataset) { // not required
		return nil
	}

	if m.LogDataset != nil {
		if err := m.LogDataset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_dataset")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dataset dataset configuration based on the context it is used
func (m *DatasetDatasetConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTraceDataset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogDataset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasetDatasetConfiguration) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *DatasetDatasetConfiguration) contextValidateTraceDataset(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceDataset != nil {

		if swag.IsZero(m.TraceDataset) { // not required
			return nil
		}

		if err := m.TraceDataset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_dataset")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetDatasetConfiguration) contextValidateLogDataset(ctx context.Context, formats strfmt.Registry) error {

	if m.LogDataset != nil {

		if swag.IsZero(m.LogDataset) { // not required
			return nil
		}

		if err := m.LogDataset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_dataset")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatasetDatasetConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetDatasetConfiguration) UnmarshalBinary(b []byte) error {
	var res DatasetDatasetConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
