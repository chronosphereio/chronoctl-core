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

// NotifierOpsGenieConfig notifier ops genie config
//
// swagger:model NotifierOpsGenieConfig
type NotifierOpsGenieConfig struct {

	// http config
	HTTPConfig *NotifierHTTPConfig `json:"http_config,omitempty"`

	// Required OpsGenie API key.
	APIKey string `json:"api_key,omitempty"`

	// Required OpsGenie API URL to send requests to, e.g.
	// "https://api.opsgenie.com/".
	APIURL string `json:"api_url,omitempty"`

	// Alert text.
	Message string `json:"message,omitempty"`

	// Description of the alert.
	Description string `json:"description,omitempty"`

	// A backlink to the sender of the notification.
	Source string `json:"source,omitempty"`

	// A set of arbitrary key/value pairs that provide further detail about the
	// alert.
	Details map[string]string `json:"details,omitempty"`

	// List of responders responsible for notifications.
	Responders []*OpsGenieConfigResponder `json:"responders"`

	// Comma separated list of tags attached to the notifications.
	Tags string `json:"tags,omitempty"`

	// Additional alert note.
	Note string `json:"note,omitempty"`

	// Priority level of alert. Possible values are P1, P2, P3, P4, and P5.
	Priority string `json:"priority,omitempty"`
}

// Validate validates this notifier ops genie config
func (m *NotifierOpsGenieConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifierOpsGenieConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConfig) { // not required
		return nil
	}

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *NotifierOpsGenieConfig) validateResponders(formats strfmt.Registry) error {
	if swag.IsZero(m.Responders) { // not required
		return nil
	}

	for i := 0; i < len(m.Responders); i++ {
		if swag.IsZero(m.Responders[i]) { // not required
			continue
		}

		if m.Responders[i] != nil {
			if err := m.Responders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("responders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("responders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this notifier ops genie config based on the context it is used
func (m *NotifierOpsGenieConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifierOpsGenieConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPConfig != nil {

		if swag.IsZero(m.HTTPConfig) { // not required
			return nil
		}

		if err := m.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *NotifierOpsGenieConfig) contextValidateResponders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Responders); i++ {

		if m.Responders[i] != nil {

			if swag.IsZero(m.Responders[i]) { // not required
				return nil
			}

			if err := m.Responders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("responders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("responders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotifierOpsGenieConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifierOpsGenieConfig) UnmarshalBinary(b []byte) error {
	var res NotifierOpsGenieConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
