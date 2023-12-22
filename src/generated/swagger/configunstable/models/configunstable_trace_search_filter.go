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

// ConfigunstableTraceSearchFilter configunstable trace search filter
//
// swagger:model configunstableTraceSearchFilter
type ConfigunstableTraceSearchFilter struct {

	// trace
	Trace *TraceSearchFilterTraceFilter `json:"trace,omitempty"`

	// Each SpanFilter object represents all conditions that need to be true on
	// the same span for the span to be considered matching the SpanFilter. If
	// `span_count` is used, the number of spans within the trace that match the
	// SpanFilter needs to be within [min, max]. Multiple SpanFilters can be used,
	// and each can be satisfied by any number of spans within the trace.
	Span []*TraceSearchFilterSpanFilter `json:"span"`
}

// Validate validates this configunstable trace search filter
func (m *ConfigunstableTraceSearchFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableTraceSearchFilter) validateTrace(formats strfmt.Registry) error {
	if swag.IsZero(m.Trace) { // not required
		return nil
	}

	if m.Trace != nil {
		if err := m.Trace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableTraceSearchFilter) validateSpan(formats strfmt.Registry) error {
	if swag.IsZero(m.Span) { // not required
		return nil
	}

	for i := 0; i < len(m.Span); i++ {
		if swag.IsZero(m.Span[i]) { // not required
			continue
		}

		if m.Span[i] != nil {
			if err := m.Span[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("span" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("span" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this configunstable trace search filter based on the context it is used
func (m *ConfigunstableTraceSearchFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableTraceSearchFilter) contextValidateTrace(ctx context.Context, formats strfmt.Registry) error {

	if m.Trace != nil {

		if swag.IsZero(m.Trace) { // not required
			return nil
		}

		if err := m.Trace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableTraceSearchFilter) contextValidateSpan(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Span); i++ {

		if m.Span[i] != nil {

			if swag.IsZero(m.Span[i]) { // not required
				return nil
			}

			if err := m.Span[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("span" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("span" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableTraceSearchFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableTraceSearchFilter) UnmarshalBinary(b []byte) error {
	var res ConfigunstableTraceSearchFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
