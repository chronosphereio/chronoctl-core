package converter

import (
	"errors"
	"fmt"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/sliceutil"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	tfmodels "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

// TerraformOutput are outputs of prometheus/alertmanager conversion that can be marshalled to HCL.
type TerraformOutput struct {
	NotificationPolicy *intschema.NotificationPolicy
	Monitors           []*intschema.Monitor
	Collections        []*intschema.Collection
	RecordingRules     []*intschema.RecordingRule
	PagerdutyNotifiers []*intschema.PagerdutyAlertNotifier
	SlackNotifiers     []*intschema.SlackAlertNotifier
	WebhookNotifiers   []*intschema.WebhookAlertNotifier
	EmailNotifiers     []*intschema.EmailAlertNotifier
	OpsGenieNotifiers  []*intschema.OpsgenieAlertNotifier
	VictorOpsNotifiers []*intschema.VictoropsAlertNotifier
}

// ToTerraform converts the configv1 output of prometheus/alertmanager conversion into Terraform types that can be
// marshalled to HCL. All resource references are resolved as <resource_type>.<resource_name>.<resource_attribute>
// except for references to teams, which are resolved as <team_slug>. Teams are the exception because their Terraform
// definitions are not generated via prom conversion and may live in separate modules.
func ToTerraform(alerts *ConvertedAlerts) (*TerraformOutput, error) {
	tfOut := &TerraformOutput{}

	notifierTypes := make(map[string]string, len(alerts.Receivers))
	for _, r := range alerts.Receivers {
		notifier, err := tfNotifier(r.Configv1Notifier)
		if err != nil {
			return nil, fmt.Errorf("error converting notifier to Terraform: %v", err)
		}
		switch {
		case notifier.Pagerduty != nil:
			tf, err := chronosphere.PagerdutyAlertNotifierFromModel(&notifier)
			if err != nil {
				return nil, fmt.Errorf("error converting pagerduty notifier to Terraform: %v", err)
			}
			tf.HCLID = tfid.SafeID(notifier.Slug)
			tfOut.PagerdutyNotifiers = append(tfOut.PagerdutyNotifiers, tf)
			notifierTypes[notifier.Slug] = "chronosphere_pagerduty_alert_notifier"
		case notifier.Email != nil:
			tf, err := chronosphere.EmailAlertNotifierFromModel(&notifier)
			if err != nil {
				return nil, fmt.Errorf("error converting email notifier to Terraform: %v", err)
			}
			tf.HCLID = tfid.SafeID(notifier.Slug)
			tfOut.EmailNotifiers = append(tfOut.EmailNotifiers, tf)
			notifierTypes[notifier.Slug] = "chronosphere_email_alert_notifier"
		case notifier.Slack != nil:
			tf, err := chronosphere.SlackAlertNotifierFromModel(&notifier)
			if err != nil {
				return nil, fmt.Errorf("error converting slack notifier to Terraform: %v", err)
			}
			tf.HCLID = tfid.SafeID(notifier.Slug)
			tfOut.SlackNotifiers = append(tfOut.SlackNotifiers, tf)
			notifierTypes[notifier.Slug] = "chronosphere_slack_alert_notifier"
		case notifier.OpsGenie != nil:
			tf, err := chronosphere.OpsgenieAlertNotifierFromModel(&notifier)
			if err != nil {
				return nil, fmt.Errorf("error converting ops genie notifier to Terraform: %v", err)
			}
			tf.HCLID = tfid.SafeID(notifier.Slug)
			tfOut.OpsGenieNotifiers = append(tfOut.OpsGenieNotifiers, tf)
			notifierTypes[notifier.Slug] = "chronosphere_opsgenie_alert_notifier"
		case notifier.VictorOps != nil:
			tf, err := chronosphere.VictoropsAlertNotifierFromModel(&notifier)
			if err != nil {
				return nil, fmt.Errorf("error converting victor ops notifier to Terraform: %v", err)
			}
			tf.HCLID = tfid.SafeID(notifier.Slug)
			tfOut.VictorOpsNotifiers = append(tfOut.VictorOpsNotifiers, tf)
			notifierTypes[notifier.Slug] = "chronosphere_victorops_alert_notifier"
		case notifier.Webhook != nil:
			tf, err := chronosphere.WebhookAlertNotifierFromModel(&notifier)
			if err != nil {
				return nil, fmt.Errorf("error converting webhook notifier to Terraform: %v", err)
			}
			tf.HCLID = tfid.SafeID(notifier.Slug)
			tfOut.WebhookNotifiers = append(tfOut.WebhookNotifiers, tf)
			notifierTypes[notifier.Slug] = "chronosphere_webhook_alert_notifier"
		default:
			return nil, errors.New("notifier must set a target")
		}
	}

	var err error
	tfOut.NotificationPolicy, err = tfPolicy(alerts.NotificationPolicy, newNotifierTFIDGenerator(notifierTypes))
	if err != nil {
		return nil, fmt.Errorf("error converting notification policy to Terraform: %v", err)
	}
	tfOut.Monitors, err = sliceutil.MapErr(alerts.Monitors, tfMonitor)
	if err != nil {
		return nil, fmt.Errorf("error converting monitor to Terraform: %v", err)
	}
	tfOut.Collections, err = sliceutil.MapErr(alerts.Collections, tfCollection)
	if err != nil {
		return nil, fmt.Errorf("error converting collection to Terraform: %v", err)
	}
	tfOut.RecordingRules, err = sliceutil.MapErr(alerts.RecordingRules, tfRecordingRule)
	if err != nil {
		return nil, fmt.Errorf("error converting recording rule to Terraform: %v", err)
	}
	return tfOut, nil
}

func tfMonitor(in *models.Configv1Monitor) (*intschema.Monitor, error) {
	b, err := in.MarshalBinary()
	if err != nil {
		return nil, err
	}
	var tf tfmodels.Configv1Monitor
	err = tf.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	out, err := chronosphere.MonitorFromModel(&tf)
	if err != nil {
		return nil, err
	}
	out.HCLID = tfid.SafeID(tf.Slug)
	// Convert tfid.ID fields to local references.
	out.NotificationPolicyId = policyTFID(out.NotificationPolicyId)
	out.CollectionId = collectionTFID(out.CollectionId)
	return out, nil
}

func tfCollection(in *models.Configv1Collection) (*intschema.Collection, error) {
	b, err := in.MarshalBinary()
	if err != nil {
		return nil, err
	}
	var tf tfmodels.Configv1Collection
	err = tf.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	out, err := chronosphere.CollectionFromModel(&tf)
	if err != nil {
		return nil, err
	}
	out.HCLID = tfid.SafeID(tf.Slug)
	out.NotificationPolicyId = policyTFID(out.NotificationPolicyId)
	return out, nil
}

func tfRecordingRule(in *models.Configv1RecordingRule) (*intschema.RecordingRule, error) {
	b, err := in.MarshalBinary()
	if err != nil {
		return nil, err
	}
	var tf tfmodels.Configv1RecordingRule
	err = tf.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	out, err := chronosphere.RecordingRuleFromModel(&tf)
	if err != nil {
		return nil, err
	}
	out.HCLID = tfid.SafeID(tf.Slug)
	out.ExecutionGroup = collectionTFID(out.ExecutionGroup)
	return out, nil
}

func tfPolicy(in *models.Configv1NotificationPolicy, notifierIDs *notifierTFIDGenerator) (*intschema.NotificationPolicy, error) {
	b, err := in.MarshalBinary()
	if err != nil {
		return nil, err
	}
	var tf tfmodels.Configv1NotificationPolicy
	err = tf.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	out, err := chronosphere.NotificationPolicyFromModel(&tf)
	if err != nil {
		return nil, err
	}
	out.HCLID = tfid.SafeID(tf.Slug)
	for _, r := range out.Route {
		for i, n := range r.Notifiers {
			r.Notifiers[i] = notifierIDs.generate(n)
		}
	}
	for _, o := range out.Override {
		for _, r := range o.Route {
			for i, n := range r.Notifiers {
				r.Notifiers[i] = notifierIDs.generate(n)
			}
		}
	}
	return out, nil
}

func tfNotifier(n *models.Configv1Notifier) (tfmodels.Configv1Notifier, error) {
	b, err := n.MarshalBinary()
	if err != nil {
		return tfmodels.Configv1Notifier{}, err
	}
	var tf tfmodels.Configv1Notifier
	err = tf.UnmarshalBinary(b)
	if err != nil {
		return tfmodels.Configv1Notifier{}, err
	}
	return tf, nil
}

func policyTFID(slug tfid.ID) tfid.ID {
	if slug.Slug() == "" {
		return slug
	}
	return tfid.LocalRef(tfid.Ref{
		Type: "chronosphere_notification_policy",
		ID:   slug.Slug(),
	})
}

func collectionTFID(slug tfid.ID) tfid.ID {
	if slug.Slug() == "" {
		return slug
	}
	return tfid.LocalRef(tfid.Ref{
		Type: "chronosphere_collection",
		ID:   slug.Slug(),
	})
}

func newNotifierTFIDGenerator(notifierTypes map[string]string) *notifierTFIDGenerator {
	return &notifierTFIDGenerator{notifierTypes: notifierTypes}
}

type notifierTFIDGenerator struct {
	notifierTypes map[string]string
}

func (n *notifierTFIDGenerator) generate(slug tfid.ID) tfid.ID {
	if slug.Slug() == "" {
		return slug
	}
	return tfid.LocalRef(tfid.Ref{
		Type: n.notifierTypes[slug.Slug()],
		ID:   slug.Slug(),
	})
}
