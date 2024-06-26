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

// ConfigUnstableUpdateSavedTraceSearchBody config unstable update saved trace search body
//
// swagger:model ConfigUnstableUpdateSavedTraceSearchBody
type ConfigUnstableUpdateSavedTraceSearchBody struct {

	// saved trace search
	SavedTraceSearch *ConfigunstableSavedTraceSearch `json:"saved_trace_search,omitempty"`

	// If true, the SavedTraceSearch will be created if it does not already exist, identified by slug. If false, an error will be returned if the SavedTraceSearch does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`
}

// Validate validates this config unstable update saved trace search body
func (m *ConfigUnstableUpdateSavedTraceSearchBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSavedTraceSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigUnstableUpdateSavedTraceSearchBody) validateSavedTraceSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.SavedTraceSearch) { // not required
		return nil
	}

	if m.SavedTraceSearch != nil {
		if err := m.SavedTraceSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saved_trace_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saved_trace_search")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config unstable update saved trace search body based on the context it is used
func (m *ConfigUnstableUpdateSavedTraceSearchBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSavedTraceSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigUnstableUpdateSavedTraceSearchBody) contextValidateSavedTraceSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.SavedTraceSearch != nil {

		if swag.IsZero(m.SavedTraceSearch) { // not required
			return nil
		}

		if err := m.SavedTraceSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saved_trace_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saved_trace_search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigUnstableUpdateSavedTraceSearchBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigUnstableUpdateSavedTraceSearchBody) UnmarshalBinary(b []byte) error {
	var res ConfigUnstableUpdateSavedTraceSearchBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
