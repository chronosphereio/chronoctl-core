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

// Configv1ResourcePools configv1 resource pools
//
// swagger:model configv1ResourcePools
type Configv1ResourcePools struct {

	// Timestamp of when the ResourcePools was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the ResourcePools was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// default pool
	DefaultPool *ResourcePoolsDefaultPool `json:"default_pool,omitempty"`

	// Optional pools.
	Pools []*ResourcePoolsPool `json:"pools"`
}

// Validate validates this configv1 resource pools
func (m *Configv1ResourcePools) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ResourcePools) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ResourcePools) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ResourcePools) validateDefaultPool(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultPool) { // not required
		return nil
	}

	if m.DefaultPool != nil {
		if err := m.DefaultPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_pool")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1ResourcePools) validatePools(formats strfmt.Registry) error {
	if swag.IsZero(m.Pools) { // not required
		return nil
	}

	for i := 0; i < len(m.Pools); i++ {
		if swag.IsZero(m.Pools[i]) { // not required
			continue
		}

		if m.Pools[i] != nil {
			if err := m.Pools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this configv1 resource pools based on the context it is used
func (m *Configv1ResourcePools) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultPool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ResourcePools) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ResourcePools) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ResourcePools) contextValidateDefaultPool(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultPool != nil {

		if swag.IsZero(m.DefaultPool) { // not required
			return nil
		}

		if err := m.DefaultPool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_pool")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1ResourcePools) contextValidatePools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pools); i++ {

		if m.Pools[i] != nil {

			if swag.IsZero(m.Pools[i]) { // not required
				return nil
			}

			if err := m.Pools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1ResourcePools) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1ResourcePools) UnmarshalBinary(b []byte) error {
	var res Configv1ResourcePools
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
