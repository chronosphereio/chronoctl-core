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

// ChronoConfigCollectionsConfig Configuration for when ContainerType == COLLECTIONS or DEFAULT is collections
//
// swagger:model ChronoConfigCollectionsConfig
type ChronoConfigCollectionsConfig struct {

	// Ordered map of prometheus group regexes to chronosphere team slug. This is used to
	// determine which team will own the collection created for each prometheus group..
	GroupTeams []*CollectionsConfigGroupTeam `json:"group_teams"`

	// The team that will own the generated notification policy
	NotificationPolicyTeamSlug string `json:"notification_policy_team_slug,omitempty"`
}

// Validate validates this chrono config collections config
func (m *ChronoConfigCollectionsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChronoConfigCollectionsConfig) validateGroupTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupTeams) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupTeams); i++ {
		if swag.IsZero(m.GroupTeams[i]) { // not required
			continue
		}

		if m.GroupTeams[i] != nil {
			if err := m.GroupTeams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("group_teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("group_teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this chrono config collections config based on the context it is used
func (m *ChronoConfigCollectionsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChronoConfigCollectionsConfig) contextValidateGroupTeams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupTeams); i++ {

		if m.GroupTeams[i] != nil {

			if swag.IsZero(m.GroupTeams[i]) { // not required
				return nil
			}

			if err := m.GroupTeams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("group_teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("group_teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChronoConfigCollectionsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChronoConfigCollectionsConfig) UnmarshalBinary(b []byte) error {
	var res ChronoConfigCollectionsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}