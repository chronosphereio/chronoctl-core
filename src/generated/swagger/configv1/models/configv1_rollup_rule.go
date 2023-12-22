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

// Configv1RollupRule configv1 rollup rule
//
// swagger:model configv1RollupRule
type Configv1RollupRule struct {

	// Unique identifier of the RollupRule. If slug is not provided, one will be generated based of the name field. Cannot be modified after the RollupRule is created.
	Slug string `json:"slug,omitempty"`

	// Required name of the RollupRule. May be modified after the RollupRule is created.
	Name string `json:"name,omitempty"`

	// Timestamp of when the RollupRule was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Timestamp of when the RollupRule was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Required slug of the bucket the RollupRule belongs to.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// Filters that determine to which metrics to apply the rule.
	Filters []*Configv1LabelFilter `json:"filters"`

	// Name of the new metric created as a result of the rollup.
	MetricName string `json:"metric_name,omitempty"`

	// metric type
	MetricType RollupRuleMetricType `json:"metric_type,omitempty"`

	// aggregation
	Aggregation Configv1AggregationType `json:"aggregation,omitempty"`

	// storage policy
	StoragePolicy *Configv1RollupRuleStoragePolicy `json:"storage_policy,omitempty"`

	// Interval between aggregated data points, equivalent to the resolution
	// field in storage policy. If set, then the storage_policy field can't be
	// set.
	Interval string `json:"interval,omitempty"`

	// Enables expansive label matching behavior for the provided filters and
	// label_policy.keep (if set). By default (expansive_match=false), a series
	// matches and aggregates only if each label defined by filters and
	// label_policy.keep (respectively) exist in said series. Setting
	// expansive_match=true removes this restriction.
	ExpansiveMatch bool `json:"expansive_match,omitempty"`

	// Defines whether to add a `__rollup_type__` label in the new metric.
	AddMetricTypeLabel bool `json:"add_metric_type_label,omitempty"`

	// Defines whether to automatically generate drop rules for this rollup.
	DropRaw bool `json:"drop_raw,omitempty"`

	// label policy
	LabelPolicy *Configv1RollupRuleLabelPolicy `json:"label_policy,omitempty"`

	// Replaces or adds new labels to the rollup, optionally based on a matching
	// regex for tag values. This is only available on Graphite rollup rules.
	LabelReplace []*RollupRuleLabelReplace `json:"label_replace"`

	// mode
	Mode Configv1RollupRuleMode `json:"mode,omitempty"`
}

// Validate validates this configv1 rollup rule
func (m *Configv1RollupRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelReplace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1RollupRule) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1RollupRule) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1RollupRule) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1RollupRule) validateMetricType(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricType) { // not required
		return nil
	}

	if err := m.MetricType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metric_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("metric_type")
		}
		return err
	}

	return nil
}

func (m *Configv1RollupRule) validateAggregation(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregation) { // not required
		return nil
	}

	if err := m.Aggregation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aggregation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("aggregation")
		}
		return err
	}

	return nil
}

func (m *Configv1RollupRule) validateStoragePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.StoragePolicy) { // not required
		return nil
	}

	if m.StoragePolicy != nil {
		if err := m.StoragePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1RollupRule) validateLabelPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelPolicy) { // not required
		return nil
	}

	if m.LabelPolicy != nil {
		if err := m.LabelPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label_policy")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1RollupRule) validateLabelReplace(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelReplace) { // not required
		return nil
	}

	for i := 0; i < len(m.LabelReplace); i++ {
		if swag.IsZero(m.LabelReplace[i]) { // not required
			continue
		}

		if m.LabelReplace[i] != nil {
			if err := m.LabelReplace[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_replace" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_replace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1RollupRule) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this configv1 rollup rule based on the context it is used
func (m *Configv1RollupRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetricType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAggregation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelReplace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1RollupRule) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1RollupRule) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1RollupRule) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {

			if swag.IsZero(m.Filters[i]) { // not required
				return nil
			}

			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1RollupRule) contextValidateMetricType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.MetricType) { // not required
		return nil
	}

	if err := m.MetricType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metric_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("metric_type")
		}
		return err
	}

	return nil
}

func (m *Configv1RollupRule) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Aggregation) { // not required
		return nil
	}

	if err := m.Aggregation.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aggregation")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("aggregation")
		}
		return err
	}

	return nil
}

func (m *Configv1RollupRule) contextValidateStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.StoragePolicy != nil {

		if swag.IsZero(m.StoragePolicy) { // not required
			return nil
		}

		if err := m.StoragePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1RollupRule) contextValidateLabelPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelPolicy != nil {

		if swag.IsZero(m.LabelPolicy) { // not required
			return nil
		}

		if err := m.LabelPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label_policy")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1RollupRule) contextValidateLabelReplace(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelReplace); i++ {

		if m.LabelReplace[i] != nil {

			if swag.IsZero(m.LabelReplace[i]) { // not required
				return nil
			}

			if err := m.LabelReplace[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_replace" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_replace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1RollupRule) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1RollupRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1RollupRule) UnmarshalBinary(b []byte) error {
	var res Configv1RollupRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
