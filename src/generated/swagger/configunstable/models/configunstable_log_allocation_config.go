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
	"github.com/go-openapi/validate"
)

// ConfigunstableLogAllocationConfig LogAllocationConfig is a singleton configuration object that specifies the
// configuration for Log budget allocations.
//
// swagger:model configunstableLogAllocationConfig
type ConfigunstableLogAllocationConfig struct {

	// Timestamp of when the LogAllocationConfig was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the LogAllocationConfig was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// default dataset
	DefaultDataset *LogAllocationConfigDefaultDataset `json:"default_dataset,omitempty"`

	// Defines datasets and budget allocations. Datasets are evaluated in order.
	DatasetAllocations []*LogAllocationConfigDatasetAllocation `json:"dataset_allocations"`
}

// Validate validates this configunstable log allocation config
func (m *ConfigunstableLogAllocationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultDataset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatasetAllocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableLogAllocationConfig) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogAllocationConfig) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogAllocationConfig) validateDefaultDataset(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultDataset) { // not required
		return nil
	}

	if m.DefaultDataset != nil {
		if err := m.DefaultDataset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_dataset")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogAllocationConfig) validateDatasetAllocations(formats strfmt.Registry) error {
	if swag.IsZero(m.DatasetAllocations) { // not required
		return nil
	}

	for i := 0; i < len(m.DatasetAllocations); i++ {
		if swag.IsZero(m.DatasetAllocations[i]) { // not required
			continue
		}

		if m.DatasetAllocations[i] != nil {
			if err := m.DatasetAllocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataset_allocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataset_allocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this configunstable log allocation config based on the context it is used
func (m *ConfigunstableLogAllocationConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultDataset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatasetAllocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableLogAllocationConfig) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogAllocationConfig) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogAllocationConfig) contextValidateDefaultDataset(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultDataset != nil {

		if swag.IsZero(m.DefaultDataset) { // not required
			return nil
		}

		if err := m.DefaultDataset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_dataset")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogAllocationConfig) contextValidateDatasetAllocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DatasetAllocations); i++ {

		if m.DatasetAllocations[i] != nil {

			if swag.IsZero(m.DatasetAllocations[i]) { // not required
				return nil
			}

			if err := m.DatasetAllocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataset_allocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataset_allocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableLogAllocationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableLogAllocationConfig) UnmarshalBinary(b []byte) error {
	var res ConfigunstableLogAllocationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
