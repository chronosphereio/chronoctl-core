package converter

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/dispatch"
	"github.com/prometheus/alertmanager/pkg/labels"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/sliceutil"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

func TestConvertPrometheus_RecordingRules(t *testing.T) {
	testRuleGroup := func(rules []rulefmt.RuleNode) rulefmt.RuleGroup {
		return rulefmt.RuleGroup{
			Name:  "test-group",
			Rules: rules,
		}
	}

	tests := []struct {
		name               string
		opts               Opts
		rules              []rulefmt.RuleNode
		wantRecordingRules []*models.Configv1RecordingRule
	}{
		{
			name: "collections",
			rules: []rulefmt.RuleNode{
				{
					Record: node("test-rule"),
					Expr:   node("sum(rate(errors[5m]))"),
					Labels: map[string]string{"zone": "us-east-1a"},
				},
			},
			wantRecordingRules: []*models.Configv1RecordingRule{
				{
					ExecutionGroup: "test-group",
					Name:           "test-rule",
					MetricName:     "test-rule",
					Slug:           "test-rule",
					PrometheusExpr: "sum(rate(errors[5m]))",
					LabelPolicy: &models.Configv1RecordingRuleLabelPolicy{
						Add: map[string]string{"zone": "us-east-1a"},
					},
				},
			},
		},
		{
			name: "buckets",
			opts: Opts{UseBuckets: true},
			rules: []rulefmt.RuleNode{
				{
					Record: node("test-rule"),
					Expr:   node("sum(rate(errors[5m]))"),
					Labels: map[string]string{"zone": "us-east-1a"},
				},
			},
			wantRecordingRules: []*models.Configv1RecordingRule{
				{
					Name:           "test-rule",
					MetricName:     "test-rule",
					Slug:           "test-rule",
					PrometheusExpr: "sum(rate(errors[5m]))",
					LabelPolicy: &models.Configv1RecordingRuleLabelPolicy{
						Add: map[string]string{"zone": "us-east-1a"},
					},
				},
			},
		},
		{
			name: "multiple recording rules, same name",
			rules: []rulefmt.RuleNode{
				{
					Record: node("test-rule"),
					Expr:   node("sum(rate(errors{zone='us-east-1a'}[5m]))"),
					Labels: map[string]string{"zone": "us-east-1a"},
				},
				{
					Record: node("test-rule"),
					Expr:   node("sum(rate(errors{zone='us-west-1a'}[5m]))"),
					Labels: map[string]string{"zone": "us-west-1a"},
				},
			},
			wantRecordingRules: []*models.Configv1RecordingRule{
				{
					ExecutionGroup: "test-group",
					Name:           "test-rule",
					MetricName:     "test-rule",
					Slug:           "test-rule",
					PrometheusExpr: "sum(rate(errors{zone='us-east-1a'}[5m]))",
					LabelPolicy: &models.Configv1RecordingRuleLabelPolicy{
						Add: map[string]string{"zone": "us-east-1a"},
					},
				},
				{
					ExecutionGroup: "test-group",
					Name:           "test-rule",
					MetricName:     "test-rule",
					Slug:           "test-rule-1",
					PrometheusExpr: "sum(rate(errors{zone='us-west-1a'}[5m]))",
					LabelPolicy: &models.Configv1RecordingRuleLabelPolicy{
						Add: map[string]string{"zone": "us-west-1a"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			promYAML, err := yaml.Marshal(rulefmt.RuleGroups{
				Groups: []rulefmt.RuleGroup{testRuleGroup(tt.rules)},
			})
			require.NoError(t, err)
			out, _, err := ConvertPrometheus(string(promYAML), "", tt.opts)
			require.NoError(t, err)
			require.Equal(t, tt.wantRecordingRules, out.RecordingRules)
		})
	}
}

func TestConvertPrometheus_SeverityOptions(t *testing.T) {
	t.Parallel()

	testRuleGroup := func(labels map[string]string) rulefmt.RuleGroup {
		return rulefmt.RuleGroup{
			Name: "test-group",
			Rules: []rulefmt.RuleNode{
				{
					Alert:  node("test-rule"),
					Expr:   node("sum(rate(errors[5m])) > 0"),
					Labels: labels,
				},
			},
		}
	}

	testMonitor := func(conditions *models.SeriesConditionsSeverityConditions) *models.Configv1Monitor {
		return &models.Configv1Monitor{
			Slug:                   "test-rule",
			Name:                   "test-rule",
			CollectionSlug:         "test-group",
			NotificationPolicySlug: "default",
			Labels:                 map[string]string{},
			PrometheusQuery:        "sum(rate(errors[5m]))",
			SeriesConditions: &models.MonitorSeriesConditions{
				Defaults: conditions,
			},
		}
	}

	testCondition := &models.MonitorCondition{Op: models.ConditionOpGT, Value: 0}

	warnMonitor := testMonitor(&models.SeriesConditionsSeverityConditions{
		Warn: conditions(testCondition),
	})
	criticalMonitor := testMonitor(&models.SeriesConditionsSeverityConditions{
		Critical: conditions(testCondition),
	})

	testAMConfig := func(m map[string]string) config.Config {
		matchers := make([]*labels.Matcher, 0, len(m))
		for k, v := range m {
			matchers = append(matchers, &labels.Matcher{Type: labels.MatchEqual, Name: k, Value: v})
		}

		return config.Config{
			Route: &config.Route{
				Receiver: "blackhole",
				Routes: []*config.Route{
					{
						Receiver: "pagerduty",
						Matchers: matchers,
					},
				},
			},
			Receivers: []config.Receiver{
				config.Receiver{Name: "blackhole"},
				config.Receiver{Name: "pagerduty"},
			},
		}
	}

	testPolicy := func(notifiers *models.RoutesSeverityNotifiers) *models.Configv1NotificationPolicy {
		return defaultPolicy(&models.NotificationPolicyRoutes{
			Defaults:  notifiers,
			Overrides: []*models.NotificationPolicyRoutesOverride{},
		})
	}

	warnPolicy := testPolicy(&models.RoutesSeverityNotifiers{
		Warn:     defaultNotifiers("pagerduty"),
		Critical: defaultNotifiers("blackhole"),
	})
	criticalPolicy := testPolicy(&models.RoutesSeverityNotifiers{
		Warn:     defaultNotifiers("blackhole"),
		Critical: defaultNotifiers("pagerduty"),
	})

	tests := []struct {
		name              string
		ruleLabels        map[string]string
		routeMatchers     map[string]string
		severityMappings  map[string]string
		severityLabelName string
		wantMonitor       *models.Configv1Monitor
		wantPolicy        *models.Configv1NotificationPolicy
		wantErr           string
	}{
		{
			name:             "severity mapping, critical",
			ruleLabels:       map[string]string{chronoSeverityLabelName: "defcon1"},
			routeMatchers:    map[string]string{chronoSeverityLabelName: "defcon1"},
			severityMappings: map[string]string{"defcon1": "critical"},
			wantMonitor:      criticalMonitor,
			wantPolicy:       criticalPolicy,
		},
		{
			name:             "severity mapping, warn",
			ruleLabels:       map[string]string{chronoSeverityLabelName: "defcon5"},
			routeMatchers:    map[string]string{chronoSeverityLabelName: "defcon5"},
			severityMappings: map[string]string{"defcon5": "warn"},
			wantMonitor:      warnMonitor,
			wantPolicy:       warnPolicy,
		},
		{
			name:             "standard severities still work with severity mappings set",
			ruleLabels:       map[string]string{chronoSeverityLabelName: warnLabel},
			routeMatchers:    map[string]string{chronoSeverityLabelName: warnLabel},
			severityMappings: map[string]string{"defcon1": "critical"},
			wantMonitor:      warnMonitor,
			wantPolicy:       warnPolicy,
		},
		{
			name:              "severity label name",
			ruleLabels:        map[string]string{"level": criticalLabel},
			routeMatchers:     map[string]string{"level": criticalLabel},
			severityLabelName: "level",
			wantMonitor:       criticalMonitor,
			wantPolicy:        criticalPolicy,
		},
		{
			name:              "severity label name and mapping",
			ruleLabels:        map[string]string{"level": "defcon1"},
			routeMatchers:     map[string]string{"level": "defcon1"},
			severityMappings:  map[string]string{"defcon1": "critical"},
			severityLabelName: "level",
			wantMonitor:       criticalMonitor,
			wantPolicy:        criticalPolicy,
		},
		{
			name:          "unmapped severity in prom config",
			ruleLabels:    map[string]string{chronoSeverityLabelName: "unknown"},
			routeMatchers: map[string]string{chronoSeverityLabelName: warnLabel},
			wantErr:       "Invalid value for field `rules_yaml.groups[0].rules[0].labels`: invalid severity value `unknown`",
		},
		{
			name:          "unmapped severity in alertmanager config",
			ruleLabels:    map[string]string{chronoSeverityLabelName: criticalLabel},
			routeMatchers: map[string]string{chronoSeverityLabelName: "unknown"},
			wantMonitor:   criticalMonitor,
			// Unmapped severities default to critical.
			wantPolicy: criticalPolicy,
		},
		{
			name:             "invalid severity mapping",
			ruleLabels:       map[string]string{chronoSeverityLabelName: criticalLabel},
			routeMatchers:    map[string]string{chronoSeverityLabelName: criticalLabel},
			severityMappings: map[string]string{"defcon1": "unknown"},
			wantErr:          "Invalid value for field `chrono_config`: mapped severity must be either `warn` or `critical`, got `unknown`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			promYAML, err := yaml.Marshal(rulefmt.RuleGroups{
				Groups: []rulefmt.RuleGroup{testRuleGroup(tt.ruleLabels)},
			})
			require.NoError(t, err)
			amYAML, err := yaml.Marshal(testAMConfig(tt.routeMatchers))
			require.NoError(t, err)
			out, _, err := ConvertPrometheus(
				string(promYAML),
				string(amYAML),
				Opts{
					SeverityMappings:  tt.severityMappings,
					SeverityLabelName: tt.severityLabelName,
				},
			)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, []*models.Configv1Monitor{tt.wantMonitor}, out.Monitors)
			require.Equal(t, tt.wantPolicy, out.NotificationPolicy)
		})
	}
}

func node(s string) yaml.Node {
	return yaml.Node{Kind: yaml.ScalarNode, Value: s}
}

func defaultPolicy(routes *models.NotificationPolicyRoutes) *models.Configv1NotificationPolicy {
	return &models.Configv1NotificationPolicy{
		Slug:   "default",
		Name:   "Default",
		Routes: routes,
	}
}

func defaultNotifiers(slugs ...string) *models.RoutesNotifierList {
	return &models.RoutesNotifierList{
		NotifierSlugs:      slugs,
		RepeatIntervalSecs: 14400,
	}
}

func conditions(c ...*models.MonitorCondition) *models.SeriesConditionsConditions {
	return &models.SeriesConditionsConditions{Conditions: c}
}

func TestParseExpr(t *testing.T) {
	t.Parallel()

	type output struct {
		expr      string
		op        models.ConditionOp
		threshold float64
		err       bool
	}

	for _, tt := range []struct {
		input            string
		output           output
		noExistsOpOutput output // Only set if different
	}{
		{
			input:  "foo > 5",
			output: output{expr: "foo", op: models.ConditionOpGT, threshold: 5},
		},
		{
			input:  "(foo >= 5)",
			output: output{expr: "foo", op: models.ConditionOpGEQ, threshold: 5},
		},
		{
			input:  `foo{x="y"} > 5`,
			output: output{expr: `foo{x="y"}`, op: models.ConditionOpGT, threshold: 5},
		},
		{
			input:            "foo < bar",
			output:           output{expr: "foo < bar", op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: "foo - bar", op: models.ConditionOpLT, threshold: 0},
		},
		{
			input:            "(foo > 5) and (bar < 3)",
			output:           output{expr: "(foo > 5) and (bar < 3)", op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: "(foo > 5) and (bar < 3)", op: models.ConditionOpNEQ, threshold: 0},
		},
		{
			input:  "sum(foo) by (bar) <= 123",
			output: output{expr: "sum by(bar) (foo)", op: models.ConditionOpLEQ, threshold: 123},
		},
		{
			input:            "sum(foo) by (bar)",
			output:           output{expr: "sum(foo) by (bar)", op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: "sum(foo) by (bar)", op: models.ConditionOpNEQ, threshold: 0},
		},
		{
			input:  "foo - bar == 100",
			output: output{expr: "foo - bar", op: models.ConditionOpEQ, threshold: 100},
		},
		{
			input:            `absent(something{other="this"})`,
			output:           output{expr: `absent(something{other="this"})`, op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: `absent(something{other="this"})`, op: models.ConditionOpNEQ, threshold: 0},
		},
		{
			input:            `time()`,
			output:           output{expr: `time()`, op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: `time()`, op: models.ConditionOpNEQ, threshold: 0},
		},
		{
			input:            `1 - time()`,
			output:           output{expr: `1 - time()`, op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: `1 - time()`, op: models.ConditionOpNEQ, threshold: 0},
		},
		{
			input:            `role:supervisor_fatal_process:count{role=~"minerva.*"}`,
			output:           output{expr: `role:supervisor_fatal_process:count{role=~"minerva.*"}`, op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: `role:supervisor_fatal_process:count{role=~"minerva.*"}`, op: models.ConditionOpNEQ, threshold: 0},
		},
		{
			input:            "poseidon_consolidator_config_reload_invalid",
			output:           output{expr: `poseidon_consolidator_config_reload_invalid`, op: models.ConditionOpEXISTS},
			noExistsOpOutput: output{expr: `poseidon_consolidator_config_reload_invalid`, op: models.ConditionOpNEQ, threshold: 0},
		},
		{
			input:  `foo[5m]`,
			output: output{err: true},
		},
		{
			input:  `foo[5m:10m]`,
			output: output{err: true},
		},
	} {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			for _, existsOp := range []bool{true, false} {
				existsOp := existsOp
				t.Run(fmt.Sprintf("existsOp=%t", existsOp), func(t *testing.T) {
					expected := tt.output
					if !existsOp && (tt.noExistsOpOutput != output{}) {
						expected = tt.noExistsOpOutput
					}

					expr, op, threshold, err := parseExpr("", tt.input, existsOp)
					if expected.err {
						assert.Error(t, err)
						return
					}

					require.NoError(t, err)
					assert.Equal(t, normalizeQuery(t, expected.expr), normalizeQuery(t, expr))
					assert.Equal(t, expected.op, op, "%v != %v", expected.op, op)
					assert.Equal(t, expected.threshold, threshold)
				})
			}
		})
	}
}

func normalizeQuery(t *testing.T, query string) string {
	t.Helper()
	expr, err := parser.ParseExpr(query)
	require.NoError(t, err)
	return expr.String()
}

func TestConvertPrometheus_ErrorWhenBothInferMonitorSignalsAndAddGroupByToNotificationPoliciesAreSet(t *testing.T) {
	t.Parallel()

	_, _, err := ConvertPrometheus("", "", Opts{
		InferMonitorSignals:              true,
		AddGroupByToNotificationPolicies: true,
	})
	require.ErrorContains(t, err,
		"only one of add_group_by_to_notification_policies and infer_monitor_signals can be set")
}

func TestInferSignal(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name           string
		route          *dispatch.Route
		monitor        *models.Configv1Monitor
		expectSignal   *models.MonitorSignalGrouping
		expectInferred bool
	}{
		{
			name: "no grouping, no labels",
			route: &dispatch.Route{
				RouteOpts: dispatch.RouteOpts{},
			},
			monitor:        &models.Configv1Monitor{},
			expectInferred: true,
		},
		{
			name: "default",
			route: &dispatch.Route{
				RouteOpts: dispatch.RouteOpts{
					GroupBy: map[model.LabelName]struct{}{
						alertNameLabelName: {},
					},
				},
			},
			monitor:        &models.Configv1Monitor{},
			expectInferred: true,
		},
		{
			name: "infer signal from root",
			route: &dispatch.Route{
				RouteOpts: dispatch.RouteOpts{
					GroupBy: map[model.LabelName]struct{}{
						alertNameLabelName: {},
						"team":             {},
					},
				},
			},
			monitor: &models.Configv1Monitor{},
			expectSignal: &models.MonitorSignalGrouping{
				LabelNames: []string{"team"},
			},
			expectInferred: true,
		},
		{
			name: "infer signal from matchers",
			route: &dispatch.Route{
				RouteOpts: dispatch.RouteOpts{
					GroupBy: map[model.LabelName]struct{}{
						alertNameLabelName: {},
					},
				},
				Routes: []*dispatch.Route{
					{
						RouteOpts: dispatch.RouteOpts{
							GroupBy: map[model.LabelName]struct{}{
								alertNameLabelName:      {},
								chronoSeverityLabelName: {},
								"service":               {},
								"environment":           {},
							},
						},
						Matchers: labels.Matchers{
							{Type: labels.MatchEqual, Name: "team", Value: "triage"},
						},
					},
					{
						RouteOpts: dispatch.RouteOpts{
							GroupBy: map[model.LabelName]struct{}{
								alertNameLabelName:      {},
								chronoSeverityLabelName: {},
								"flow":                  {},
							},
						},
						Matchers: labels.Matchers{
							{Type: labels.MatchEqual, Name: "team", Value: "rapid-response"},
						},
					},
				},
			},
			monitor: &models.Configv1Monitor{
				Labels: map[string]string{
					"team":       "triage",
					"importance": "high",
				},
			},
			expectSignal: &models.MonitorSignalGrouping{
				LabelNames: []string{"environment", "service"},
			},
			expectInferred: true,
		},
		{
			name: "infer most common grouping",
			route: &dispatch.Route{
				RouteOpts: dispatch.RouteOpts{
					GroupBy: map[model.LabelName]struct{}{
						alertNameLabelName: {},
					},
				},
				Routes: []*dispatch.Route{
					{
						RouteOpts: dispatch.RouteOpts{
							GroupBy: map[model.LabelName]struct{}{
								alertNameLabelName:      {},
								chronoSeverityLabelName: {},
								"service":               {},
								"environment":           {},
							},
						},
						Matchers: labels.Matchers{
							{Type: labels.MatchEqual, Name: "team", Value: "triage"},
							{Type: labels.MatchEqual, Name: "env", Value: "prod"},
						},
					},
					{
						RouteOpts: dispatch.RouteOpts{
							GroupBy: map[model.LabelName]struct{}{
								alertNameLabelName:      {},
								chronoSeverityLabelName: {},
								"service":               {},
							},
						},
						Matchers: labels.Matchers{
							{Type: labels.MatchEqual, Name: "team", Value: "triage"},
							{Type: labels.MatchEqual, Name: "env", Value: "prod"},
						},
					},
				},
			},
			monitor: &models.Configv1Monitor{
				Labels: map[string]string{
					"team":       "triage",
					"importance": "high",
				},
			},
			expectInferred: true,
		},
		{
			name: "signal per series",
			route: &dispatch.Route{
				RouteOpts: dispatch.RouteOpts{
					GroupByAll: true,
				},
			},
			monitor:        &models.Configv1Monitor{},
			expectSignal:   &models.MonitorSignalGrouping{SignalPerSeries: true},
			expectInferred: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			signal, inferred := inferSignal(tt.monitor, tt.route)
			assert.Equal(t, tt.expectSignal, signal)
			assert.Equal(t, tt.expectInferred, inferred)
		})
	}
}

func TestGroupMatcher(t *testing.T) {
	tests := []struct {
		name       string
		groupTeams []GroupTeamMatcher
		groups     []string
		expected   []string

		noMatch []string
		wantErr string
	}{
		{
			name: "simple groups",
			groupTeams: []GroupTeamMatcher{
				{GroupRegex: "group1", TeamSlug: "team1"},
				{GroupRegex: "group2", TeamSlug: "team2"},
			},
			groups:   []string{"group1", "group2"},
			expected: []string{"team1", "team2"},
		},
		{
			name: "early takes priority",
			groupTeams: []GroupTeamMatcher{
				{GroupRegex: "group1", TeamSlug: "team1"},
				{GroupRegex: "group2", TeamSlug: "team2"},
				{GroupRegex: "group1", TeamSlug: "team10"},
			},
			groups:   []string{"group1", "group2", "group1"},
			expected: []string{"team1", "team2", "team1"},
			noMatch:  []string{"group"},
		},
		{
			name: "anchoring",
			groupTeams: []GroupTeamMatcher{
				{GroupRegex: "group", TeamSlug: "team1"},
				{GroupRegex: "group1", TeamSlug: "team2"},
				{GroupRegex: "group20", TeamSlug: "team10"},
				{GroupRegex: "20group20", TeamSlug: "team10"},
			},
			groups:   []string{"group", "group1", "group20", "20group20"},
			expected: []string{"team1", "team2", "team10", "team10"},
			noMatch:  []string{"20"},
		},
		{
			name: "invalid regex",
			groupTeams: []GroupTeamMatcher{
				{GroupRegex: "group", TeamSlug: "team1"},
				{GroupRegex: "group1", TeamSlug: "team2"},
				{GroupRegex: "group20", TeamSlug: "team10"},
				{GroupRegex: "*\\d+]", TeamSlug: "team10"},
			},
			wantErr: "Invalid value for field `group_team[3].group`: invalid regex: error parsing regexp",
		},
		{
			name: "no group",
			groupTeams: []GroupTeamMatcher{
				{GroupRegex: "group", TeamSlug: "team1"},
				{GroupRegex: "group1", TeamSlug: "team2"},
				{GroupRegex: "", TeamSlug: "team10"},
			},
			wantErr: "Field `group_team[2].group` must be set",
		},
		{
			name: "no team",
			groupTeams: []GroupTeamMatcher{
				{GroupRegex: "group", TeamSlug: "team1"},
				{GroupRegex: "group1", TeamSlug: ""},
			},
			wantErr: "Field `group_team[1].team_slug` must be set",
		},
		{
			name: "regex default matcher",
			groupTeams: []GroupTeamMatcher{
				{GroupRegex: "group1", TeamSlug: "team1"},
				{GroupRegex: "group2", TeamSlug: "team2"},
				{GroupRegex: ".*", TeamSlug: "default"},
			},
			groups:   []string{"group1", "group2", "group3", "hello", "goodbye"},
			expected: []string{"team1", "team2", "default", "default", "default"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			matcher, err := newGroupTeamMatcher(tt.groupTeams, fieldPath("group_team"))
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
			for i, group := range tt.groups {
				team := matcher.Match(group)
				assert.Equal(t, tt.expected[i], team)
			}

			for _, group := range tt.noMatch {
				team := matcher.Match(group)
				assert.Equal(t, "", team)
			}
		})
	}
}

func TestConvertAlertManagerConfig(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name           string
		config         config.Config
		expectRoutes   *models.NotificationPolicyRoutes
		expectWarnings []string
		expectErr      string
	}{
		// Single route without children.
		// Its receiver is the default for all severities.
		{
			name: "simple - no matchers",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname", "severity"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{},
			},
			expectWarnings: []string{
				"Route references receiver `blackhole` not defined in config (route)",
			},
		},
		{
			name: "simple - receiver with multiple destinations",
			config: config.Config{
				Route: &config.Route{
					Receiver: "multiple",
					GroupBy:  []model.LabelName{"alertname", "severity"},
				},
				Receivers: []config.Receiver{
					{
						Name: "multiple",
						EmailConfigs: []*config.EmailConfig{
							{To: "e-1"},
							{To: "e-2"},
						},
						SlackConfigs: []*config.SlackConfig{
							{
								APIURL: &config.SecretURL{URL: &url.URL{Host: "s-1"}},
							},
						},
					},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"multiple-0", "multiple-1", "multiple-2"},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"multiple-0", "multiple-1", "multiple-2"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{},
			},
		},

		// Single route that matches a specific severity (critical).
		// There will be no notifier routed for other severities (warn).
		{
			name: "simple - match on critical only",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname"},
					Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{},
			},
		},

		// Route with two children. Each child matches the same severity (critical) but has different receivers.
		// The resulting policy has default receivers of:
		// warn     -> [blackhole] (from the root node defining defaults by not having matchers)
		// critical -> [pagerduty, slack] (these are from the children overriding the root's defaults by matching on a severity)
		{
			name: "simple - child route overrides default critical notifier",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname"},
					MatchRE:  nil,
					Routes: []*config.Route{
						{
							Receiver: "pagerduty",
							GroupBy:  []model.LabelName{"alertname"},
							Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
						},
						{
							Receiver: "slack",
							GroupBy:  []model.LabelName{"alertname"},
							Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
					{Name: "pagerduty"},
					{Name: "slack"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"pagerduty", "slack"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{},
			},
			expectWarnings: []string{
				"Ignoring continue=false on leaf node; will import as continue=true (route.routes[0])",
			},
		},

		// Route with two children, one of which has an arbitrary matcher.
		// Neither children match on a specific severity.
		// As a result, there are overrides for the first child route having specific matchers.
		// The second child, having no matchers, ends up overriding the root node to define the default notifiers.
		{
			name: "override based on matcher",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname", "severity"},
					Routes: []*config.Route{
						{
							Receiver: "slack_triage",
							GroupBy:  []model.LabelName{"alertname", "severity"},
							MatchRE:  map[string]config.Regexp{"service": {Regexp: regexp.MustCompile("gateway|ui")}},
							Continue: true,
						},
						{
							Receiver: "slack",
							GroupBy:  []model.LabelName{"alertname", "severity"},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
					{Name: "slack_triage"},
					{Name: "slack"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"slack"},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"slack"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{
								Type:  models.Configv1LabelMatcherMatcherTypeREGEX,
								Name:  "service",
								Value: "gateway|ui",
							},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"slack-triage"},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"slack-triage"},
							},
						},
					},
				},
			},
		},

		// Complicated subtree mapping to organizations with subteams:
		//                  ø(root)
		//              ----/  | \-----------------
		//             /       |                   \
		//      org=platform   org=control          ø(catchall)
		//         /        \
		//  team=triage   org=platform,team=foundation
		{
			name: "subtrees for organizations and teams within",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname", "severity"},
					Routes: []*config.Route{
						{
							Receiver: "blackhole",
							GroupBy:  []model.LabelName{"alertname", "severity"},
							Match:    map[string]string{"org": "platform"},
							Routes: []*config.Route{
								{
									Receiver: "triage_slack",
									GroupBy:  []model.LabelName{"alertname", "severity"},
									Match:    map[string]string{"team": "triage"},
									Routes: []*config.Route{
										{
											Receiver: "triage_pagerduty",
											GroupBy:  []model.LabelName{"alertname", "severity"},
											Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
											Continue: true,
										},
										{
											Receiver: "triage_email",
											GroupBy:  []model.LabelName{"alertname", "severity"},
											Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
										},
									},
								},
								{
									Receiver: "foundation_slack",
									GroupBy:  []model.LabelName{"alertname", "severity"},
									Match:    map[string]string{"org": "platform", "team": "foundation"},
								},
							},
						},
						{
							Receiver: "control_slack",
							GroupBy:  []model.LabelName{"alertname", "severity"},
							Match:    map[string]string{"org": "control"},
						},
						{
							Receiver: "catchall_pagerduty",
							GroupBy:  []model.LabelName{"alertname", "severity"},
							Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
					{Name: "triage_pagerduty"},
					{Name: "triage_email"},
					{Name: "triage_slack"},
					{Name: "foundation_slack"},
					{Name: "control_slack"},
					{Name: "catchall_pagerduty"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"catchall-pagerduty"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "org", Value: "platform"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "team", Value: "triage"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"triage-slack"},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"triage-pagerduty", "triage-email"},
							},
						},
					},
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "org", Value: "platform"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "team", Value: "foundation"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"foundation-slack"},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"foundation-slack"},
							},
						},
					},
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "org", Value: "platform"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"blackhole"},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"blackhole"},
							},
						},
					},
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "org", Value: "control"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"control-slack"},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"control-slack"},
							},
						},
					},
				},
			},
		},
		// subtree where only one of the child route matches on severity
		//                  ø(blackhole)
		//                    |
		// node with    (alert_type: foo) matcher
		//              ----/   \-----------------
		//             /                          \
		//      a,b (severity: critical)          a
		{
			name: "only one of the child node has a severity matcher",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname", "severity"},
					Routes: []*config.Route{
						{
							GroupBy: []model.LabelName{"alertname"},
							Match:   map[string]string{"alert_type": "foo"},
							Routes: []*config.Route{
								{
									Receiver: "a",
									Match:    map[string]string{"severity": "critical"},
								},
								{
									Receiver: "b",
								},
							},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
					{Name: "a"},
					{Name: "b"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "alert_type", Value: "foo"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"b"},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"a"},
							},
						},
					},
				},
			},
			expectWarnings: []string{
				"Ignoring continue=false on leaf node; will import as continue=true (route.routes[0].routes[0])",
			},
		},

		// subtree where only one of the child route matches on severity
		//                  ø(blackhole)
		//                    |
		// node with    (alert_type: foo) matcher
		//              ----/   \-----------------
		//             /                          \
		//      a,b (severity: warn)          a
		{
			name: "only one of the child node has a severity matcher",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname", "severity"},
					Routes: []*config.Route{
						{
							GroupBy: []model.LabelName{"alertname"},
							Match:   map[string]string{"alert_type": "foo"},
							Routes: []*config.Route{
								{
									Receiver: "a",
									Match:    map[string]string{"severity": "warn"},
								},
								{
									Receiver: "b",
								},
							},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
					{Name: "a"},
					{Name: "b"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "alert_type", Value: "foo"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"a"},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"b"},
							},
						},
					},
				},
			},
			expectWarnings: []string{
				"Ignoring continue=false on leaf node; will import as continue=true (route.routes[0].routes[0])",
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.config.Route.RepeatInterval == nil {
				tt.config.Route.RepeatInterval = &defaultRepeatInterval
			}

			route := newMultiReceiverRoute(dispatch.NewRoute(tt.config.Route, nil))

			policy, _, warnings, err := convertAlertManagerConfig(
				"", Opts{DisableContinueTrueExpansion: true}, route, &tt.config)
			if tt.expectErr != "" {
				require.ErrorContains(t, err, tt.expectErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, "Default", policy.Name)
			assert.Equal(t, "default", policy.Slug)
			assert.Equal(t, tt.expectRoutes, policy.Routes)
			assert.Equal(t, tt.expectWarnings, warnings)
		})
	}
}

func TestConvertAlertManagerConfig_GroupBy(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name           string
		config         config.Config
		expectRoutes   *models.NotificationPolicyRoutes
		expectWarnings []string
		expectErr      string
	}{
		{
			name: "route with group by",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname", "severity"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
						GroupBy: &models.NotificationPolicyRoutesGroupBy{
							LabelNames: []string{"alertname", "severity"},
						},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
						GroupBy: &models.NotificationPolicyRoutesGroupBy{
							LabelNames: []string{"alertname", "severity"},
						},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{},
			},
			expectWarnings: []string{
				"Route references receiver `blackhole` not defined in config (route)",
			},
		},
		{
			name: "route and child both have groupBy and use the same label names",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname"},
					MatchRE:  nil,
					Routes: []*config.Route{
						{
							Receiver: "pagerduty",
							GroupBy:  []model.LabelName{"alertname"},
							Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
					{Name: "pagerduty"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
						GroupBy: &models.NotificationPolicyRoutesGroupBy{
							LabelNames: []string{"alertname"},
						},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"pagerduty"},
						GroupBy: &models.NotificationPolicyRoutesGroupBy{
							LabelNames: []string{"alertname"},
						},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{},
			},
		},
		{
			name: "child should inherit groupBy from parent node",
			config: config.Config{
				Route: &config.Route{
					Receiver: "blackhole",
					GroupBy:  []model.LabelName{"alertname"},
					MatchRE:  nil,
					Routes: []*config.Route{
						{
							Receiver: "pagerduty",
							Match:    map[string]string{chronoSeverityLabelName: criticalLabel},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "blackhole"},
					{Name: "pagerduty"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Defaults: &models.RoutesSeverityNotifiers{
					Warn: &models.RoutesNotifierList{
						NotifierSlugs: []string{"blackhole"},
						GroupBy: &models.NotificationPolicyRoutesGroupBy{
							LabelNames: []string{"alertname"},
						},
					},
					Critical: &models.RoutesNotifierList{
						NotifierSlugs: []string{"pagerduty"},
						GroupBy: &models.NotificationPolicyRoutesGroupBy{
							LabelNames: []string{"alertname"},
						},
					},
				},
				Overrides: []*models.NotificationPolicyRoutesOverride{},
			},
		},
		{
			name: "child and grand-children should inherit groupBy",
			config: config.Config{
				Route: &config.Route{
					Receiver: "a",
					GroupBy: []model.LabelName{
						"foo", "bar", "baz",
					},
					Match: map[string]string{"type": "grpc"},
					Routes: []*config.Route{
						{
							Receiver: "b",
							Match:    map[string]string{"type": "grpc", "service": "auth"},
							Routes: []*config.Route{
								{
									Receiver: "c",
									Match:    map[string]string{"env": "prod"},
								},
							},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Overrides: []*models.NotificationPolicyRoutesOverride{
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "env", Value: "prod"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "service", Value: "auth"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "type", Value: "grpc"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"c"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"c"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
						},
					},
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "service", Value: "auth"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "type", Value: "grpc"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"b"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"b"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
						},
					},
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "type", Value: "grpc"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"a"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"a"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "child and grand-children have different values for groupBy",
			config: config.Config{
				Route: &config.Route{
					Receiver: "a",
					GroupBy: []model.LabelName{
						"foo", "bar", "baz",
					},
					Match: map[string]string{"type": "grpc"},
					Routes: []*config.Route{
						{
							Receiver: "b",
							Match:    map[string]string{"type": "grpc", "service": "auth"},
							GroupBy:  []model.LabelName{"foo"},
							Routes: []*config.Route{
								{
									Receiver: "c",
									Match:    map[string]string{"env": "prod"},
									GroupBy:  []model.LabelName{"bar", "baz"},
								},
							},
						},
					},
				},
				Receivers: []config.Receiver{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			expectRoutes: &models.NotificationPolicyRoutes{
				Overrides: []*models.NotificationPolicyRoutesOverride{
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "env", Value: "prod"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "service", Value: "auth"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "type", Value: "grpc"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"c"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz"},
								},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"c"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz"},
								},
							},
						},
					},
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "service", Value: "auth"},
							{Type: models.Configv1LabelMatcherMatcherTypeEXACT, Name: "type", Value: "grpc"},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"b"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"foo"},
								},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"b"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"foo"},
								},
							},
						},
					},
					{
						AlertLabelMatchers: []*models.Configv1LabelMatcher{
							{
								Type:  models.Configv1LabelMatcherMatcherTypeEXACT,
								Name:  "type",
								Value: "grpc",
							},
						},
						Notifiers: &models.RoutesSeverityNotifiers{
							Warn: &models.RoutesNotifierList{
								NotifierSlugs: []string{"a"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
							Critical: &models.RoutesNotifierList{
								NotifierSlugs: []string{"a"},
								GroupBy: &models.NotificationPolicyRoutesGroupBy{
									LabelNames: []string{"bar", "baz", "foo"},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.config.Route.RepeatInterval == nil {
				tt.config.Route.RepeatInterval = &defaultRepeatInterval
			}

			route := newMultiReceiverRoute(dispatch.NewRoute(tt.config.Route, nil))

			policy, _, warnings, err := convertAlertManagerConfig(
				"", Opts{AddGroupByToNotificationPolicies: true}, route, &tt.config)
			if tt.expectErr != "" {
				require.ErrorContains(t, err, tt.expectErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, "Default", policy.Name)
			assert.Equal(t, "default", policy.Slug)
			assert.Equal(t, tt.expectRoutes, policy.Routes)
			assert.Equal(t, tt.expectWarnings, warnings)
		})
	}
}

func TestExpandReceiver(t *testing.T) {
	t.Parallel()
	// Not expanded
	assert.Equal(t,
		notifiersFromV1([]*models.Configv1Notifier{
			{
				Name: "rname",
				Webhook: &models.NotifierWebhookConfig{
					HTTPConfig: &models.NotifierHTTPConfig{},
					URL:        "wh-one",
				},
			},
		}),
		expandReceiver(config.Receiver{
			Name: "rname",
			WebhookConfigs: []*config.WebhookConfig{
				{
					NotifierConfig: config.NotifierConfig{VSendResolved: true},
					URL:            &config.SecretURL{URL: mustParseURL(t, "wh-one")},
				},
			},
		}))

	// Multiple of same type
	assert.Equal(t,
		notifiersFromV1([]*models.Configv1Notifier{
			{
				Name:         "rname (0)",
				SkipResolved: true,
				Webhook: &models.NotifierWebhookConfig{
					HTTPConfig: &models.NotifierHTTPConfig{},
					URL:        "wh-one",
				},
			},
			{
				Name:         "rname (1)",
				SkipResolved: true,
				Webhook: &models.NotifierWebhookConfig{
					HTTPConfig: &models.NotifierHTTPConfig{},
					URL:        "wh-two",
				},
			},
		}),
		expandReceiver(config.Receiver{
			Name: "rname",
			WebhookConfigs: []*config.WebhookConfig{
				{URL: &config.SecretURL{URL: mustParseURL(t, "wh-one")}},
				{URL: &config.SecretURL{URL: mustParseURL(t, "wh-two")}},
			},
		}))

	// Multiple of multiple types
	assert.Equal(t,
		notifiersFromV1([]*models.Configv1Notifier{
			{
				Name:         "rname (0)",
				SkipResolved: true,
				Webhook: &models.NotifierWebhookConfig{
					HTTPConfig: &models.NotifierHTTPConfig{},
					URL:        "wh-one",
				},
			},
			{
				Name:         "rname (1)",
				SkipResolved: true,
				Webhook: &models.NotifierWebhookConfig{
					HTTPConfig: &models.NotifierHTTPConfig{},
					URL:        "wh-two",
				},
			},
			{
				Name:         "rname (2)",
				SkipResolved: true,
				Slack: &models.NotifierSlackConfig{
					HTTPConfig: &models.NotifierHTTPConfig{},
					Channel:    "s-one",
					Fields:     []*models.SlackConfigField{},
					Actions:    []*models.NotifierSlackConfigAction{},
				},
			},
		}),
		expandReceiver(config.Receiver{
			Name: "rname",
			WebhookConfigs: []*config.WebhookConfig{
				{URL: &config.SecretURL{URL: mustParseURL(t, "wh-one")}},
				{URL: &config.SecretURL{URL: mustParseURL(t, "wh-two")}},
			},
			SlackConfigs: []*config.SlackConfig{
				{Channel: "s-one"},
			},
		}))
}

type testConfigFile struct {
	Alertmanager       *config.Config                     `yaml:"alertmanager"`
	NotificationPolicy *models.Configv1NotificationPolicy `yaml:"notificationPolicy"`
}

// this function generates various different combinations of labelSets to feed to the route tree
// for e.g. if we have the following matchers
// foo = 1
// foo = 2
// bar = 3
// bar = 4
// It would generate the following combinations
// foo = 1, bar = 3
// foo = 1, bar = 4
// foo = 2, bar = 3
// foo = 2, bar = 4
func combine(t *testing.T, labelSetCombinations []model.LabelSet, matchers labels.Matchers) []model.LabelSet {
	result := []model.LabelSet{}
	for _, m := range matchers {
		var val string
		switch m.Type {
		case labels.MatchRegexp:
			// hardcoded value that matches the regex that we have in tests
			val = "foo_bar_baz"
		case labels.MatchEqual:
			val = m.Value
		default:
			t.Fatalf("unexpected matcher type: %v", m.Type)
		}

		// if there are no combinations then add this labelSet as the only combination
		if len(labelSetCombinations) == 0 {
			tmp := model.LabelSet{}
			tmp[model.LabelName(m.Name)] = model.LabelValue(val)
			result = append(result, tmp)
		}

		// otherwise combine this labelSet with other existing labelSetCombinations and add to result
		for _, ls := range labelSetCombinations {
			tmp := model.LabelSet{}
			tmp[model.LabelName(m.Name)] = model.LabelValue(val)
			for k, v := range ls {
				tmp[k] = v
				result = append(result, tmp)
			}
		}
	}
	return result
}

func findAllMatchersRecursively(r *multiReceiverRoute, uniqueMatchers map[labels.Matcher]bool) {
	for _, m := range r.Matchers {
		if _, ok := uniqueMatchers[*m]; !ok {
			uniqueMatchers[*m] = true
		}
	}

	for _, cr := range r.Routes {
		findAllMatchersRecursively(cr, uniqueMatchers)
	}
}

func validateReceiversMatchForDifferentLabelCombinations(t *testing.T, original, modified *multiReceiverRoute) {
	matcherByLabelName := map[string]labels.Matchers{}
	uniqueMatchers := make(map[labels.Matcher]bool)
	findAllMatchersRecursively(original, uniqueMatchers)
	for m := range uniqueMatchers {
		m := m
		matcherByLabelName[m.Name] = append(matcherByLabelName[m.Name], &m)
	}

	for name := range matcherByLabelName {
		// add the "" (empty) matcher for this label Name so that we generate matchers with "" string as well
		m, err := labels.NewMatcher(labels.MatchEqual, name, "")
		require.NoError(t, err)
		matcherByLabelName[name] = append(matcherByLabelName[name], m)
	}

	// generate different label combinations
	labelSetCombinations := make([]model.LabelSet, 0)
	for _, matchers := range matcherByLabelName {
		labelSetCombinations = combine(t, labelSetCombinations, matchers)
	}
	// if no labelSetCombinations were generate, lets generate a dummy one
	if len(labelSetCombinations) == 0 {
		labelSetCombinations = append(labelSetCombinations, model.LabelSet{"foo": "bar"})
	}

	numUniqueCombinations := 1
	for _, matchers := range matcherByLabelName {
		numUniqueCombinations *= len(matchers)
	}

	for _, combination := range labelSetCombinations {
		routesMatchingOriginal := match(original, combination)
		routesMatchingModified := match(modified, combination)

		var originalReceivers, modifiedReceivers []string
		for _, r := range routesMatchingOriginal {
			originalReceivers = append(originalReceivers, r.RouteOpts.Receivers...)
		}
		sort.Strings(originalReceivers)
		originalReceivers = slices.Compact(originalReceivers)
		for _, r := range routesMatchingModified {
			modifiedReceivers = append(modifiedReceivers, r.RouteOpts.Receivers...)
		}
		require.Equalf(t, originalReceivers, modifiedReceivers,
			"mismatch between receivers for combination: [%v], orig: [%v], modified: [%v]",
			combination, originalReceivers, modifiedReceivers)
	}
}

func TestConvertAlertManagerConfig_ContinueEqualsTrue_WithNoChildren(t *testing.T) {
	completeWarning := warningSiblingDifferentValue + " (routes[0])"

	t.Parallel()
	testCases := []struct {
		configFile string
		warnings   []string
	}{
		{
			configFile: "single_continue_true_route.yaml",
		},
		{
			configFile: "two_continue_true_routes.yaml",
		},
		{
			configFile: "continue_true_followed_by_continue_false.yaml",
		},
		{
			configFile: "continue_true_followed_by_continue_false_with_matcher.yaml",
		},
		{
			configFile: "continue_true_between_continue_false_routes.yaml",
		},
		{
			configFile: "multiple_continue_true_routes_mixed_with_continue_false.yaml",
		},
		{
			configFile: "continue_true_routes_nested_under_continue_false.yaml",
		},
		{
			configFile: "continue_true_followed_by_two_continue_false_with_matchers.yaml",
		},
		{
			configFile: "continue_true_followed_by_continue_false_with_and_without_matchers.yaml",
		},
		{
			configFile: "result_has_union_of_receivers.yaml",
		},
		{
			configFile: "continue_true_routes_nested_under_continue_false.yaml",
		},
		{
			configFile: "continue_true_routes_nested_under_continue_false_2.yaml",
		},
		{
			configFile: "following_siblings_have_children_including_a_child_with_continue_true.yaml",
		},
		{
			configFile: "following_siblings_have_a_child_node_with_matchers.yaml",
		},
		{
			configFile: "continue_true_has_two_siblings_that_have_children.yaml",
		},
		{
			configFile: "continue_true_with_siblings_that_have_grandchildren.yaml",
		},
		{
			configFile: "following_sibling_has_children.yaml",
		},
		// tests for continue true with matchers
		{
			configFile: "continue_true_with_matchers_followed_by_continue_false.yaml",
		},
		{
			configFile: "continue_true_without_matchers_followed_by_one_with_matchers.yaml",
		},
		{
			configFile: "continue_true_with_matchers_followed_by_one_without_matchers.yaml",
		},
		{
			configFile: "continue_true_with_matchers_is_present_in_sibling.yaml",
		},
		{
			configFile: "continue_true_with_matchers_separated_by_continue_false.yaml",
		},
		{
			configFile: "continue_true_with_matcher_followed_by_one_without_matcher.yaml",
		},
		// continue true with matcher followed by other routes with matchers
		{
			configFile: "matchers_are_disjoint_two_routes.yaml",
		},
		{
			configFile: "matchers_are_disjoint_two_routes_followed_by_continue_true_without_matchers.yaml",
		},
		{
			configFile: "matchers_are_disjoint_two_routes_followed_by_continue_false.yaml",
		},
		{
			configFile: "matchers_are_disjoint_two_routes_followed_by_continue_false_with_matchers.yaml",
		},
		{
			configFile: "matchers_are_disjoint_three_routes.yaml",
		},
		{
			configFile: "matchers_are_disjoint_and_have_routes_without_matchers_in_between.yaml",
		},
		{
			configFile: "matchers_are_disjoint_and_have_routes_without_matchers_in_between_followed_by_continue_false.yaml",
		},
		{
			configFile: "matchers_are_same_for_three_routes.yaml",
		},
		{
			configFile: "matchers_are_same_and_have_routes_without_matchers_in_between.yaml",
		},
		{
			configFile: "matchers_are_same_and_are_followed_by_continue_false.yaml",
		},
		{
			configFile: "matchers_are_common_three_routes.yaml",
		},
		{
			configFile: "matchers_are_common_and_have_routes_without_matchers_in_between.yaml",
		},
		{
			configFile: "matcher_keys_are_common_three_routes.yaml",
		},
		{
			configFile: "matcher_keys_are_common_and_have_routes_without_matchers_in_between.yaml",
		},
		{
			configFile: "matchers_of_children_are_disjoint_three_routes.yaml",
		},
		{
			configFile: "matchers_of_children_are_disjoint_two_routes_inherit_receiver.yaml",
		},
		{
			configFile: "matchers_are_disjoint_five_routes.yaml",
		},
		{
			configFile: "more_than_10_following_sibling_for_a_continue_true_route_with_matchers.yaml",
			warnings:   []string{warningTenOrMoreSiblingsFollowingRoute + " (routes[0])"},
		},
		{
			configFile: "following_sibling_have_different_value_of_group_by.yaml",
			warnings:   []string{completeWarning},
		},
		{
			configFile: "following_sibling_have_different_value_of_group_wait.yaml",
			warnings:   []string{completeWarning},
		},
		{
			configFile: "following_sibling_have_a_different_value_of_group_by_all.yaml",
			warnings:   []string{completeWarning},
		},
		{
			configFile: "following_sibling_have_different_value_of_repeat_interval.yaml",
			warnings:   []string{completeWarning},
		},
		{
			configFile: "continue_true_has_children_and_is_followed_by_other_valid_routes.yaml",
			warnings:   []string{warningContinueTrueRouteHasChildren + " (routes[0])"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.configFile, func(t *testing.T) {
			t.Parallel()

			b, err := os.ReadFile(filepath.Join("testdata", "continue_true", tt.configFile))
			require.NoError(t, err)

			var testData testConfigFile
			err = yaml.Unmarshal(b, &testData)
			require.NoError(t, err)

			// This route or is going to get modified by convertAlertManagerConfig
			or := newMultiReceiverRoute(dispatch.NewRoute(testData.Alertmanager.Route, nil))
			policy, _, warnings, err := convertAlertManagerConfig(
				"", Opts{}, or, testData.Alertmanager)
			require.NoError(t, err)

			for _, w := range tt.warnings {
				require.Truef(t, slices.Contains(warnings, w), "%v", warnings)
			}
			require.Equal(t, testData.NotificationPolicy, policy)
		})
	}
}

func generatePermutations(routes []*config.Route) [][]*config.Route {
	var result [][]*config.Route

	var helper func([]*config.Route, int)
	helper = func(currentPermutation []*config.Route, position int) {
		if position == len(routes) {
			permutationCopy := make([]*config.Route, len(currentPermutation))
			copy(permutationCopy, currentPermutation)
			result = append(result, permutationCopy)
			return
		}
		for _, route := range routes {
			helper(append(currentPermutation, route), position+1)
		}
	}

	helper([]*config.Route{}, 0)
	return result
}

func TestConvertAlertManagerConfig_ContinueEqualsTrue_FuzzyTesting(t *testing.T) {
	t.Parallel()

	continueTrueRoutes := []*config.Route{
		{
			Continue: true,
		},
		{
			Continue: true,
			Matchers: []*labels.Matcher{
				{
					Type:  labels.MatchEqual,
					Name:  "a",
					Value: "b",
				},
			},
		},
		{
			Continue: true,
			Matchers: []*labels.Matcher{
				{
					Type:  labels.MatchEqual,
					Name:  "w",
					Value: "x",
				},
			},
		},
		{
			Continue: true,
			Matchers: []*labels.Matcher{
				{
					Type:  labels.MatchEqual,
					Name:  "a",
					Value: "e",
				},
			},
		},
	}

	testCases := []struct {
		name               string
		continueFalseRoute *config.Route
	}{
		{
			"continue: false route with no matchers",
			&config.Route{
				Continue: false,
			},
		},
		{
			"continue: false route with same matcher as one of the continue: true routes",
			&config.Route{
				Continue: false,
				Matchers: []*labels.Matcher{
					{
						Type:  labels.MatchEqual,
						Name:  "a",
						Value: "b",
					},
				},
			},
		},
		{
			"continue: false route with different matcher to all of the continue: true routes",
			&config.Route{
				Continue: false,
				Matchers: []*labels.Matcher{
					{
						Type:  labels.MatchEqual,
						Name:  "g",
						Value: "h",
					},
				},
			},
		},
		{
			"continue: false route with same matcher name but different value to one of the continue: true routes",
			&config.Route{
				Continue: false,
				Matchers: []*labels.Matcher{
					{
						Type:  labels.MatchEqual,
						Name:  "a",
						Value: "z",
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// merge continueTrueRoutes and the continueFalseRoute specific to this test case into routes
			var routes []*config.Route
			routes = append(routes, continueTrueRoutes...)
			routes = append(routes, tt.continueFalseRoute)
			// Ensure that the length of routes is 5, so we only generate 5^5 permutations. The time complexity of this
			// test increases exponentially with the length of the routes. Currently, the test runs in 6-7 seconds in CI,
			// so it would be unwise to modify the test in a way that increases the length of the routes.
			require.Len(t, routes, 5)
			permutations := generatePermutations(routes)

			for _, permutation := range permutations {
				var updatedRoutes []*config.Route
				for jdx := range permutation {
					// copy the route and then add a receiver to the route. Copy is essential because until now all permutations
					// just hold references to the same initial routes. If we modify the permutation in place it would end up overwriting
					// other duplicate elements that are part of permutation as well
					r := *permutation[jdx]
					r.Receiver = strconv.Itoa(jdx)
					updatedRoutes = append(updatedRoutes, &r)
				}
				originalCfg := config.Config{
					Route: &config.Route{
						Receiver: "blackhole",
						Routes:   updatedRoutes,
					},
					Receivers: []config.Receiver{
						{Name: "blackhole"},
						{Name: "0"},
						{Name: "1"},
						{Name: "2"},
						{Name: "3"},
						{Name: "4"},
					},
				}

				b, err := yaml.Marshal(originalCfg)
				require.NoError(t, err)

				configToModify, err := config.Load(string(b))
				require.NoError(t, err)

				originalRoute := dispatch.NewRoute(originalCfg.Route, nil)
				routeToModify := newMultiReceiverRoute(dispatch.NewRoute(configToModify.Route, nil))

				_, _, _, err = convertAlertManagerConfig("", Opts{}, routeToModify, &originalCfg)
				require.NoError(t, err)

				// validate receivers match between the original and modified route for different label combinations
				validateReceiversMatchForDifferentLabelCombinations(t, newMultiReceiverRoute(originalRoute), routeToModify)
			}
			t.Logf("tested: %d permutations", len(permutations))
		})
	}
}

func TestConvertReceivers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		existingReceivers map[string]string
		inputReceivers    []*models.Configv1Notifier
		wantReceivers     []*models.Configv1Notifier
		wantErr           string
	}{
		{
			name: "no receivers provided",
		},
		{
			name:           "receiver with no name",
			inputReceivers: []*models.Configv1Notifier{{}},
			wantErr:        "name must be set",
		},
		{
			name:              "uses slug of existing receiver with same name",
			existingReceivers: map[string]string{"test-name": "test-slug"},
			inputReceivers: []*models.Configv1Notifier{
				{Name: "test-name"},
			},
			wantReceivers: []*models.Configv1Notifier{
				{Name: "test-name", Slug: "test-slug"},
			},
		},
		{
			name:              "no existing receiver with same name",
			existingReceivers: map[string]string{"test-name": "test-slug"},
			inputReceivers: []*models.Configv1Notifier{
				{Name: "different-name"},
			},
			wantReceivers: []*models.Configv1Notifier{
				{Name: "different-name", Slug: "different-name"},
			},
		},
		{
			name:              "no existing receiver with same name, long name",
			existingReceivers: map[string]string{"test-name": "test-slug"},
			inputReceivers: []*models.Configv1Notifier{
				{Name: strings.Repeat("a", 200)},
			},
			wantReceivers: []*models.Configv1Notifier{
				{Name: strings.Repeat("a", 200), Slug: strings.Repeat("a", 200)},
			},
		},
		{
			name:              "no existing receiver with same name, slug conflict",
			existingReceivers: map[string]string{"test-name": "conflict"},
			inputReceivers: []*models.Configv1Notifier{
				{Name: "conflict"},
			},
			wantReceivers: []*models.Configv1Notifier{
				{Name: "conflict", Slug: "conflict-1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := notifiersFromV1(tt.inputReceivers)
			err := setReceiverSlugs(
				input,
				Opts{ReceiverNamesToSlugs: tt.existingReceivers},
			)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, notifiersFromV1(tt.wantReceivers), input)
		})
	}
}

func notifiersFromV1(notifiers []*models.Configv1Notifier) []*Notifier {
	return sliceutil.Map(notifiers, func(n *models.Configv1Notifier) *Notifier {
		return &Notifier{Configv1Notifier: n}
	})
}

func TestConvertPrometheusOptions(t *testing.T) {
	t.Parallel()

	// Default options
	opts := &Opts{}

	assertSeverityMapped := func(opts *Opts, expected string, expectValid bool, input string) {
		mapped, valid := opts.severity(input)
		assert.Equal(t, valid, expectValid)
		assert.Equal(t, mapped, expected)
	}

	assertSeverityMapped(opts, criticalLabel, false, "unmapped")
	assertSeverityMapped(opts, criticalLabel, true, criticalLabel)
	assertSeverityMapped(opts, warnLabel, true, warnLabel)

	severity, matchers := opts.matchersWithSeverity(nil)
	assert.Equal(t, "", severity)
	assert.Empty(t, matchers)

	severity, matchers = opts.matchersWithSeverity([]*labels.Matcher{
		{Type: labels.MatchEqual, Name: "a", Value: "b"},
		{Type: labels.MatchEqual, Name: chronoSeverityLabelName, Value: "unmapped"},
	})
	assert.Equal(t, criticalLabel, severity)
	assert.Equal(t, []*labels.Matcher{{Type: labels.MatchEqual, Name: "a", Value: "b"}}, matchers)

	// Options with mapped severities
	opts = &Opts{
		SeverityMappings: map[string]string{
			"mapped": warnLabel,
		},
	}

	assertSeverityMapped(opts, criticalLabel, false, "unmapped")
	assertSeverityMapped(opts, warnLabel, true, "mapped")
	assertSeverityMapped(opts, criticalLabel, true, criticalLabel)
	assertSeverityMapped(opts, warnLabel, true, warnLabel)

	severity, matchers = opts.matchersWithSeverity([]*labels.Matcher{
		{Type: labels.MatchEqual, Name: "a", Value: "b"},
		{Type: labels.MatchEqual, Name: chronoSeverityLabelName, Value: "mapped"},
	})
	assert.Equal(t, warnLabel, severity)
	assert.Equal(t, []*labels.Matcher{{Type: labels.MatchEqual, Name: "a", Value: "b"}}, matchers)
}

func match(r *multiReceiverRoute, lset model.LabelSet) []*multiReceiverRoute {
	if !r.Matchers.Matches(lset) {
		return nil
	}
	var all []*multiReceiverRoute
	for _, cr := range r.Routes {
		matches := match(cr, lset)
		all = append(all, matches...)
		if matches != nil && !cr.Continue {
			break
		}
	}

	// If no child nodes were matches, the current node itself is a match.
	if len(all) == 0 {
		all = append(all, r)
	}

	return all
}

func mustParseURL(t *testing.T, rawURL string) *url.URL {
	t.Helper()
	u, err := url.Parse(rawURL)
	require.NoError(t, err)
	return u
}
