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

// StateunstableListMetricUsagesByLabelNameResponse stateunstable list metric usages by label name response
//
// swagger:model stateunstableListMetricUsagesByLabelNameResponse
type StateunstableListMetricUsagesByLabelNameResponse struct {

	// page
	Page *Configv1PageResult `json:"page,omitempty"`

	// usages
	Usages []*StateunstableMetricUsageByLabelName `json:"usages"`
}

// Validate validates this stateunstable list metric usages by label name response
func (m *StateunstableListMetricUsagesByLabelNameResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateunstableListMetricUsagesByLabelNameResponse) validatePage(formats strfmt.Registry) error {
	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if m.Page != nil {
		if err := m.Page.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *StateunstableListMetricUsagesByLabelNameResponse) validateUsages(formats strfmt.Registry) error {
	if swag.IsZero(m.Usages) { // not required
		return nil
	}

	for i := 0; i < len(m.Usages); i++ {
		if swag.IsZero(m.Usages[i]) { // not required
			continue
		}

		if m.Usages[i] != nil {
			if err := m.Usages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stateunstable list metric usages by label name response based on the context it is used
func (m *StateunstableListMetricUsagesByLabelNameResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateunstableListMetricUsagesByLabelNameResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	if m.Page != nil {

		if swag.IsZero(m.Page) { // not required
			return nil
		}

		if err := m.Page.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *StateunstableListMetricUsagesByLabelNameResponse) contextValidateUsages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Usages); i++ {

		if m.Usages[i] != nil {

			if swag.IsZero(m.Usages[i]) { // not required
				return nil
			}

			if err := m.Usages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateunstableListMetricUsagesByLabelNameResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateunstableListMetricUsagesByLabelNameResponse) UnmarshalBinary(b []byte) error {
	var res StateunstableListMetricUsagesByLabelNameResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}