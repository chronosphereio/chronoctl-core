// Package converter is a library for converting prometheus/alertmanager configuration to Chronosphere.
package converter

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"
	"time"

	anyascii "github.com/anyascii/go"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/dispatch"
	"github.com/prometheus/alertmanager/pkg/labels"
	commoncfg "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/prometheus/prometheus/promql/parser"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

const (
	// allSeverities is a fake severity used only internally to functions in this package.
	// It represents notifiers to be used for all possible severities.
	allSeverities = "__all_severities"

	warningSiblingDifferentValue = "Ignoring continue=true on node since the following siblings have a different value for a non-receiver field and cannot be aggregated"

	warningTenOrMoreSiblingsFollowingRoute = "Ignoring continue=true on node with matchers since it has more than 10 following siblings; will import as continue=false"

	warningContinueTrueRouteHasChildren = "Ignoring continue=true on non-leaf node; will import as continue=false"

	alertNameLabelName      = "alertname"
	chronoSeverityLabelName = "severity"
	warnLabel               = "warn"
	criticalLabel           = "critical"

	separator = "-"
)

var (
	// templatedRegexp is used to detect if a label value is templated.
	templatedRegexp = regexp.MustCompile("{{.*}}")

	// neverMatchLabelValue is a placeholder alert label value.
	// In order to infer signal groupings we attempt to traverse the alert route tree
	// to notifier the alert routes to via its labels, and then take groupBys of that
	// route as the signal groupings. When alerts have templated values
	// we do not know what they will evaluate to so can not use them for routing.
	// In these cases we set the label value to neverMatchLabelName so that label the value
	// is not used to route the alert in the alertmanager route tree when inferring the
	// signals.
	neverMatchLabelValue = model.LabelValue("__chrono_never_match")

	// defaultRepeatInterval is the default repeat interval for notification routes.
	defaultRepeatInterval = model.Duration(1 * time.Hour)

	severities = []string{warnLabel, criticalLabel}

	validSlugChars    = regexp.MustCompile("[a-zA-Z0-9-_]+")
	validNonPreferred = regexp.MustCompile("[_]+")
)

// Notifier wraps a models.Configv1Notifier with additional fields that
// were not carried over to the config API.
type Notifier struct {
	*models.Configv1Notifier

	// Source is the Pagerduty source. Set only for Pagerduty notifiers.
	Source string
}

// ConvertedAlerts contains all the entities resulting from converting Prometheus-based alerts.
type ConvertedAlerts struct {
	EntityGroups       []*models.Configv1Collection
	Collections        []*models.Configv1Collection
	NotificationPolicy *models.Configv1NotificationPolicy
	Receivers          []*Notifier
	Monitors           []*models.Configv1Monitor
	RecordingRules     []*models.Configv1RecordingRule
}

// GroupTeamMatcher captures how to map evaluation groups to teams.
type GroupTeamMatcher struct {
	GroupRegex string
	TeamSlug   string
}

// Opts are options for converting prometheus/alertmanager configuration to Chronosphere entities.
type Opts struct {
	MonitorSignalPerSeries           bool
	SeverityLabelName                string
	SeverityMappings                 map[string]string
	InferMonitorSignals              bool
	ReceiverNamesToSlugs             map[string]string
	GroupTeamMatchers                []GroupTeamMatcher
	NotificationPolicyTeamSlug       string
	AddGroupByToNotificationPolicies bool

	// Legacy options
	DisableMonitorSlugAssignment bool
	ExistsOpNotSupported         bool
	DisableContinueTrueExpansion bool
	UseBuckets                   bool

	groupTeamMatcher groupTeamMatcher
}

func (o *Opts) severityLabelName() string {
	if o.SeverityLabelName != "" {
		return o.SeverityLabelName
	}
	return chronoSeverityLabelName
}

// severity returns the Chronosphere severity for a user-specified severity.
// If the user-specified severity doesn't map to a valid Chronosphere severity, critical is returned as a default.
func (c *Opts) severity(s string) (string, bool) {
	chronoSeverity := s
	if c.SeverityMappings != nil {
		if mappedSeverity, ok := c.SeverityMappings[s]; ok {
			chronoSeverity = mappedSeverity
		}
	}

	if !validSeverity(chronoSeverity) {
		return criticalLabel, false
	}

	return chronoSeverity, true
}

// matchersWithSeverity returns the matchers with an (optional) severity matcher excluded.
// The severity being matched is returned, or an empty string if no severity matcher is found.
func (c *Opts) matchersWithSeverity(matchers []*labels.Matcher) (string, []*labels.Matcher) {
	severityLabelName := c.severityLabelName()
	newMatchers := make([]*labels.Matcher, 0, len(matchers))
	var severityMatch string
	for _, matcher := range matchers {
		if matcher.Name == severityLabelName && matcher.Type == labels.MatchEqual {
			severityMatch = matcher.Value
			continue
		}
		newMatchers = append(newMatchers, matcher)
	}

	if severityMatch != "" {
		severityMatch, _ = c.severity(severityMatch)
	}

	return severityMatch, newMatchers
}

// ConvertPrometheus converts the provided prometheus/alertmanager configuration to Chronosphere entities.
func ConvertPrometheus(
	rulesYAML string,
	alertManagerYAML string,
	opts Opts,
) (*ConvertedAlerts, []string, error) {
	err := validateOpts(&opts, fieldPath("chrono_config"))
	if err != nil {
		return nil, nil, fmt.Errorf("validate options: %w", err)
	}

	var promRules rulefmt.RuleGroups
	if err := yaml.Unmarshal([]byte(rulesYAML), &promRules); err != nil {
		return nil, nil, fieldPath("rules_yaml").invalidReason(
			"Unable to unmarshal prometheus rules: %v", err)
	}

	// Check for nil, instead of empty to differentiate intentional empty groups vs field unset.
	if promRules.Groups == nil {
		return nil, nil, fieldPath("rules_yaml").invalidReason(
			"Prometheus input is missing top-level groups field, see " +
				"https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/")
	}

	promResult, err := convertPrometheusRuleGroups("rules_yaml", opts, &promRules)
	if err != nil {
		return nil, nil, fmt.Errorf("convert prometheus rule groups: %w", err)
	}

	var alertReceivers []*Notifier
	var notificationPolicy *models.Configv1NotificationPolicy
	var alertWarnings []string

	// Only apply alertmanager config if there's any monitor
	// A monitor is created from an alert rule
	if len(promResult.monitors) > 0 {
		var alertmgrConfig config.Config
		if err := yaml.Unmarshal([]byte(alertManagerYAML), &alertmgrConfig); err != nil {
			return nil, nil, fieldPath("alertmanager_yaml").invalidReason("Unable to unmarshal alertmanager rules: %v", err)
		}

		if numTemplates := len(alertmgrConfig.Templates); numTemplates > 0 {
			// nil out the list of template files since they are not supported.
			// We inline the template files into the notifier definitions (see ExpandAlertManagerTemplates)
			alertWarnings = append(alertWarnings, fmt.Sprintf("Ignoring and niling out alertmanager templates slice of size %d", numTemplates))
			alertmgrConfig.Templates = nil
		}

		// Compile alertmanager version of the route, this will make it simpler for us to walk the tree.
		rootRoute := *alertmgrConfig.Route
		if rootRoute.RepeatInterval == nil {
			rootRoute.RepeatInterval = &defaultRepeatInterval
		}
		route := dispatch.NewRoute(alertmgrConfig.Route, nil)

		policy, receivers, warnings, err := convertAlertManagerConfig("alertmanager_yaml", opts, newMultiReceiverRoute(route), &alertmgrConfig)
		if err != nil {
			return nil, nil, fmt.Errorf("convert alertmanager config: %w", err)
		}
		alertWarnings = append(alertWarnings, warnings...)

		// Update all monitors to use the generated notification policy, and possibly signal-per-series
		for _, monitor := range promResult.monitors {
			monitor.NotificationPolicySlug = policy.Slug
			if opts.InferMonitorSignals {
				if grouping, ok := inferSignal(monitor, route); ok {
					monitor.SignalGrouping = grouping
				} else {
					alertWarnings = append(alertWarnings, fmt.Sprintf(
						"Unable to determine a signal grouping for monitor with slug `%s`; will default to grouping by monitor and severity",
						monitor.Slug))
				}
			}
		}

		alertReceivers = receivers
		notificationPolicy = policy
	}

	alerts := &ConvertedAlerts{
		EntityGroups:       promResult.entityGroups,
		Collections:        promResult.collections,
		NotificationPolicy: notificationPolicy,
		Receivers:          alertReceivers,
		Monitors:           promResult.monitors,
		RecordingRules:     promResult.recordingRules,
	}

	if opts.UseBuckets {
		alerts.Collections = nil
	}

	return alerts, alertWarnings, nil
}

func validateOpts(opts *Opts, fp fieldPath) error {
	gtm, err := newGroupTeamMatcher(opts.GroupTeamMatchers, fp.field("group_teams"))
	if err != nil {
		return fmt.Errorf("invalid group to team mapping: %w", err)
	}
	opts.groupTeamMatcher = gtm
	if opts.InferMonitorSignals && opts.AddGroupByToNotificationPolicies {
		return fp.invalidReason("only one of add_group_by_to_notification_policies and infer_monitor_signals can be set")
	}
	for _, s := range opts.SeverityMappings {
		if !validSeverity(s) {
			return fp.invalidReason("mapped severity must be either `%s` or `%s`, got `%s`", warnLabel, criticalLabel, s)
		}
	}
	return nil
}

// groupTeam is a regex matcher for a prometheus group name and associated team.
type groupTeam struct {
	GroupMatcher *regexp.Regexp
	TeamSlug     string
}

type groupTeamMatcher []groupTeam

func newGroupTeamMatcher(
	groupTeams []GroupTeamMatcher,
	fp fieldPath,
) (groupTeamMatcher, error) {
	matcher := make(groupTeamMatcher, 0, len(groupTeams))
	for i, unparsedGroupToTeam := range groupTeams {
		if unparsedGroupToTeam.TeamSlug == "" {
			return nil, fp.index(i).field("team_slug").valueRequired()
		}
		if unparsedGroupToTeam.GroupRegex == "" {
			return nil, fp.index(i).field("group").valueRequired()
		}
		groupMatcher, err := regexp.Compile("^(?:" + unparsedGroupToTeam.GroupRegex + ")$")
		if err != nil {
			return nil, fp.index(i).field("group").invalidReason("invalid regex: %s", err.Error())
		}
		matcher = append(matcher, groupTeam{
			GroupMatcher: groupMatcher,
			TeamSlug:     unparsedGroupToTeam.TeamSlug,
		})
	}
	return matcher, nil
}

func (g groupTeamMatcher) Match(group string) string {
	for _, groupTeam := range g {
		if groupTeam.GroupMatcher.MatchString(group) {
			return groupTeam.TeamSlug
		}
	}
	return ""
}

func expandBinaryExpr(path fieldPath, binaryExpr *parser.BinaryExpr, allowExistsOp bool) (string, models.ConditionOp, float64, error) {
	// Extract the operator.
	// Save comparison ops for later.
	// Join and boolean expr ops are treated similarly to aggregation expressions: they are returned as-is and
	// evaluated as "expr > 0" (e.g. "x * y" evaluated as "x * y > 0").
	var op models.ConditionOp
	switch binaryExpr.Op {
	case parser.EQLC:
		op = models.ConditionOpEQ
	case parser.NEQ:
		op = models.ConditionOpNEQ
	case parser.LTE:
		op = models.ConditionOpLEQ
	case parser.LSS:
		op = models.ConditionOpLT
	case parser.GTE:
		op = models.ConditionOpGEQ
	case parser.GTR:
		op = models.ConditionOpGT
	case parser.ADD:
		fallthrough
	case parser.SUB:
		fallthrough
	case parser.MUL:
		fallthrough
	case parser.DIV:
		fallthrough
	case parser.MOD:
		fallthrough
	case parser.LAND:
		fallthrough
	case parser.LOR:
		fallthrough
	case parser.LUNLESS:
		if allowExistsOp {
			return binaryExpr.String(), models.ConditionOpEXISTS, 0, nil
		}
		return binaryExpr.String(), models.ConditionOpNEQ, 0, nil
	default:
		return "", "", 0, path.invalidReason("Unsupported operator `%v`", binaryExpr.Op)
	}

	// Simple case of RHS being a number ("x > 5")
	if threshold, ok := binaryExpr.RHS.(*parser.NumberLiteral); ok {
		return binaryExpr.LHS.String(), op, threshold.Val, nil
	}

	if allowExistsOp {
		return binaryExpr.String(), models.ConditionOpEXISTS, 0, nil
	}

	// RHS is a full expression, so transform to "(LHS - RHS) OP 0" ("x >= y" to "(x - y) >= 0").
	return (&parser.BinaryExpr{
		Op:         parser.SUB,
		LHS:        binaryExpr.LHS,
		RHS:        binaryExpr.RHS,
		ReturnBool: false,
	}).String(), op, 0, nil
}

// parseExpr parses a PromQL expression for use with a monitor query and condition.
// When allowExistsOp is true:
// If the form of the expression is "expr op number (e.g. "count(things) > 5") then the returned op and value
// will reflect "op" and "number" in the expression. Otherwise the returned op will be EXISTS.
//
// When allowExistsOp is false (legacy behavior):
// If the expression is not in the form of "expr op number" then it will be transformed to one with a numeric RHS.
// For example: "x > y" to "x - y > 0", or "time()" to "time != 0".
func parseExpr(path fieldPath, expr string, allowExistsOp bool) (string, models.ConditionOp, float64, error) {
	exprAST, err := parser.ParseExpr(expr)
	if err != nil {
		return "", "", 0, path.invalidReason("Unable to parse expression: %v", err)
	}

	// Unwrap parens
	for {
		if parenExpr, ok := exprAST.(*parser.ParenExpr); ok {
			exprAST = parenExpr.Expr
		} else {
			break
		}
	}

	// If the expression is a binary expression, use it unchanged.
	if binaryExpr, ok := exprAST.(*parser.BinaryExpr); ok {
		return expandBinaryExpr(path, binaryExpr, allowExistsOp)
	}

	// If the expression is an aggregate ("sum(up)"):
	// 1. If the EXISTS op is supported, keep it as is.
	// 2. Otherwise, keep it almost as-is but use != 0 as a threshold ("sum(up) != 0").
	if _, ok := exprAST.(*parser.AggregateExpr); ok {
		if allowExistsOp {
			return expr, models.ConditionOpEXISTS, 0, nil
		}

		return expr, models.ConditionOpNEQ, 0, nil
	}

	// If the expression is a call that returns a vector or a scalar ("time()" or "things{}"):
	// 1. If the EXISTS op is supported, keep it as is.
	// 2. Otherwise, keep it almost as-is but use != 0 as a threshold ("time() != 0" or "things{} != 0").
	if callExpr, ok := exprAST.(*parser.Call); ok {
		if allowExistsOp {
			return expr, models.ConditionOpEXISTS, 0, nil
		}

		switch callExpr.Func.ReturnType {
		case parser.ValueTypeVector:
			fallthrough
		case parser.ValueTypeScalar:
			return callExpr.String(), models.ConditionOpNEQ, 0, nil
		}
	}

	if vectorExpr, ok := exprAST.(*parser.VectorSelector); ok {
		if allowExistsOp {
			return expr, models.ConditionOpEXISTS, 0, nil
		}

		return vectorExpr.String(), models.ConditionOpNEQ, 0, nil
	}

	return "", "", 0, path.invalidReason(`Alert rule expression must evaluate to an instant vector or a number`)
}

// inferSignal attempts to determine a monitor's signal grouping based on a route tree.
// The inferred signal grouping, which many be nil, and a boolean indicating whether the signal grouping
// was inferred is returned.
func inferSignal(
	monitor *models.Configv1Monitor,
	route *dispatch.Route,
) (*models.MonitorSignalGrouping, bool) {
	// Build a set of known labels. We want to ignore templated labels as we do not know what they
	// will evaluate to at runtime. To ignore the templated variables we assign them a value
	// that can never be matched.
	knownMonitorLabels := model.LabelSet{}
	for k, v := range monitor.Labels {
		if templatedRegexp.MatchString(v) {
			knownMonitorLabels[model.LabelName(k)] = neverMatchLabelValue
			continue
		}
		knownMonitorLabels[model.LabelName(k)] = model.LabelValue(v)
	}

	// Find the matching node in the alertmanager route tree.
	routes := route.Match(knownMonitorLabels)
	if len(routes) == 0 {
		return nil, false
	}

	// groupBys merges GroupBy and GroupByAll into a new map
	groupBys := func(routeOpts dispatch.RouteOpts) map[model.LabelName]struct{} {
		newGroupBys := make(map[model.LabelName]struct{}, len(routeOpts.GroupBy)+1)
		for k, v := range routeOpts.GroupBy {
			newGroupBys[k] = v
		}
		if routeOpts.GroupByAll {
			newGroupBys["..."] = struct{}{}
		}

		return newGroupBys
	}

	// Get the common groupBys of the routes, these will be our signal groupings. We get the
	// common groupBys to handle the case where several routes are returned due to `continue: true`
	// and these routes have different groupBys.
	commonGroupBys := groupBys(routes[0].RouteOpts)
	for _, route := range routes {
		routeGroupBys := groupBys(route.RouteOpts)

		for groupBy := range commonGroupBys {
			if _, ok := routeGroupBys[groupBy]; !ok {
				delete(commonGroupBys, groupBy)
			}
		}
	}

	if _, ok := commonGroupBys["..."]; ok {
		return &models.MonitorSignalGrouping{
			SignalPerSeries: true,
		}, true
	}

	var signalLabels []string
	for groupBy := range commonGroupBys {
		groupByString := string(groupBy)
		// Drop __alertname__ and severity labels from group_by in converted policy to avoid confusion.
		// All notifications are grouped by severity and monitor automatically.
		if groupByString == alertNameLabelName || groupByString == chronoSeverityLabelName {
			continue
		}
		signalLabels = append(signalLabels, groupByString)
	}

	if len(signalLabels) > 0 {
		sort.Strings(signalLabels)
		return &models.MonitorSignalGrouping{
			LabelNames: signalLabels,
		}, true
	}

	return nil, true
}

func (a *continueConverter) addInheritedReceivers(r *multiReceiverRoute) *multiReceiverRoute {
	cp := copyRoute(r)
	cp.RouteOpts.Receivers = unionReceivers(cp.RouteOpts.Receivers, maps.Keys(a.inheritedContinueTrueReceivers))
	return cp
}

func copyRoute(r *multiReceiverRoute) *multiReceiverRoute {
	cp := &multiReceiverRoute{RouteOpts: r.RouteOpts}
	// create a copy as we don't want to modify the "receivers" for the elements within the r route
	tmpReceivers := make([]string, len(r.RouteOpts.Receivers))
	copy(tmpReceivers, r.RouteOpts.Receivers)
	cp.RouteOpts.Receivers = tmpReceivers
	// copy matchers and routes so that any changes to the "r".Matchers and "r".Routes doesn't affect "cp"
	for _, m := range r.Matchers {
		matcher := *m
		cp.Matchers = append(cp.Matchers, &matcher)
	}
	for _, r := range r.Routes {
		route := *r
		cp.Routes = append(cp.Routes, &route)
	}
	return cp
}

type routeSortedSet struct {
	sortedRoutes []*multiReceiverRoute
}

// Add adds a route, merging with existing routes as necessary.
func (s *routeSortedSet) add(r *multiReceiverRoute) {
	r = copyRoute(r)
	// sort and compact matchers for the incoming route to eliminate duplicate matchers
	slices.SortStableFunc(r.Matchers, func(i *labels.Matcher, j *labels.Matcher) int {
		if i.Name == j.Name {
			return strings.Compare(i.Value, j.Value)
		}
		return strings.Compare(i.Name, j.Name)
	})
	r.Matchers = slices.CompactFunc(r.Matchers, func(i *labels.Matcher, j *labels.Matcher) bool {
		return reflect.DeepEqual(i, j)
	})

	foundIdx := -1
	for idx, er := range s.sortedRoutes {
		if reflect.DeepEqual(r.Matchers, er.Matchers) {
			foundIdx = idx
			break
		}
	}
	// if there is an existing route with the same matchers as "r", merge the receivers for the two and store the result in the
	// existing route, otherwise add incoming route to "s.routes"
	if foundIdx != -1 {
		s.sortedRoutes[foundIdx].RouteOpts.Receivers = unionReceivers(s.sortedRoutes[foundIdx].RouteOpts.Receivers, r.RouteOpts.Receivers)
	} else {
		s.sortedRoutes = append(s.sortedRoutes, r)
	}
}

// routes returns the routes stored in routeSortedSet after sorting them
func (s *routeSortedSet) routes() []*multiReceiverRoute {
	// sort the result such that the routes with a higher number of matchers are matched first
	slices.SortStableFunc(s.sortedRoutes, func(r1 *multiReceiverRoute, r2 *multiReceiverRoute) int {
		return len(r2.Matchers) - len(r1.Matchers)
	})

	return s.sortedRoutes
}

type continueConverter struct {
	// aggregatedContinueTrueRoutes holds all the running combinations of continue: true routes. These are then
	// joined with the other continue:false routes to come up with the updatedRoutes. This is only valid for siblings i.e.
	// nodes at the same level. It is not inherited by child nodes.
	aggregatedContinueTrueRoutes []*multiReceiverRoute

	// inheritedContinueTrueReceivers represents the list of continue: true receivers inherited by this node. This is applicable
	// when a preceding sibling has continue=true routes. These need to be passed to the child nodes.
	// See "TestConvertAlertManagerConfig_ContinueEqualsTrue_WithNoChildrenOrMatchers/following_siblings_have_a_child_node_with_matchers.yaml" for an example
	inheritedContinueTrueReceivers map[string]struct{}
}

// addContinueTrueRoute adds a route where continue: true to aggregatedContinueTrueRoutes. Right now we only support "continue: true" routes with no
// matchers or child nodes.
// The method for now merges the routes, so that the result is a single route that has the receivers for all continue: true
// routes.
func (a *continueConverter) addContinueTrueRoute(trueRoute *multiReceiverRoute) {
	if len(a.aggregatedContinueTrueRoutes) == 0 {
		cp := a.addInheritedReceivers(trueRoute)
		cp.Continue = false
		a.aggregatedContinueTrueRoutes = []*multiReceiverRoute{cp}
		return
	}

	trueRoute = a.addInheritedReceivers(trueRoute)
	// sortedRoutes contains the result of combining the current "aggregatedContinueTrueRoutes" with the incoming trueRoute
	sortedRoutes := routeSortedSet{}
	for _, r := range a.aggregatedContinueTrueRoutes {
		if len(trueRoute.Matchers) == 0 && len(r.Matchers) == 0 {
			// merge the two into one route that has receivers for the existing route and the incoming "trueRoute"
			// e.g.
			// receiver a (continue=true)
			//          ↓
			// receiver b (continue=true)
			//
			// should get combined into
			//
			// receiver a, b (continue=true)
			r.RouteOpts.Receivers = unionReceivers(r.RouteOpts.Receivers, trueRoute.RouteOpts.Receivers)
			sortedRoutes.add(r)
		} else if len(r.Matchers) > 0 && len(trueRoute.Matchers) > 0 {
			// this is the case where the existing route and incoming route both have matchers.
			if !reflect.DeepEqual(r.Matchers, trueRoute.Matchers) {
				// When matchers are different, we generate three routes
				// 1. route corresponding to "r"
				// 2. route corresponding to "trueRoute"
				// 3. merge route that has matchers for "r" and "trueRoute"
				existingCopy := copyRoute(r)
				existingCopy.RouteOpts.Receivers = unionReceivers(existingCopy.RouteOpts.Receivers, trueRoute.RouteOpts.Receivers)
				existingCopy.Matchers = append(existingCopy.Matchers, trueRoute.Matchers...)
				sortedRoutes.add(existingCopy)
				sortedRoutes.add(r)
				sortedRoutes.add(trueRoute)
			} else {
				// matchers are same, we only need to generate one route that has receivers for both "r" and "trueRoute"
				r.RouteOpts.Receivers = unionReceivers(r.RouteOpts.Receivers, trueRoute.RouteOpts.Receivers)
				sortedRoutes.add(r)
			}
		} else if len(r.Matchers) > 0 {
			// existing "continue:true" route (r) has matchers but incoming route doesn't have matchers, this should expand into maximum two routes
			// 1. with matchers of existing route and receivers of both existing and the incoming trueRoute
			// 2. incoming trueRoute if there is no route within continueConverter without matchers. If there is a route within continueConverter
			//    without matchers then when that route is combined with this trueRoute, the resulting route would contain the trueRoute
			// e.g.
			// receiver a (continue=true)
			//          ↓
			// receiver b (continue=true) (foo=bar)
			//          ↓
			// receiver c (continue=true)
			// ---
			// Say after a and b have been combined, this is what aggregatedContinueTrueRoutes looks like
			// 1. receiver a, b (continue=true) (foo=bar)
			// 2. receiver a (continue=true)
			//
			// Now when c is being combined with 1 above, then we would generate the route that has a, b and c as receivers which is "existingRouteCopy" below
			// but we would not generate the route for c as 2 above only has "a" as receiver and that doesn't have any matchers. Later "a" would be combined
			// with "c". The final routes that we want here are
			//
			// receiver a, b, c (foo=bar) (continue=fale)
			//             ↓
			// receiver a, c (continue=false)
			r.RouteOpts.Receivers = unionReceivers(r.RouteOpts.Receivers, trueRoute.RouteOpts.Receivers)
			sortedRoutes.add(r)
			sortedRoutes.add(trueRoute)
		} else if len(trueRoute.Matchers) > 0 {
			// if the incoming route has matchers, then this should expand into two routes
			// 1. with the matchers, this would have receivers for the trueRoute and also the aggregatedContinueTrueRoutes
			// 2. without any matchers, this would only have receivers corresponding to the aggregatedContinueTrueRoutes
			// e.g.
			// receiver a (continue=true)
			//          ↓
			// receiver b (continue=true) (foo=bar)
			//
			// in this case above when trueCopy is route with receiver "b" and aggregatedContinueTrueRoutes has one route with receiver "a"
			// the result would be that aggregatedContinueTrueRoutes have two routes in them that look as follows
			//
			// receiver a, b (continue=true) (foo=bar) => trueRoute below
			//           ↓
			// receiver a (continue=true) => cp below
			trueRoute.RouteOpts.Receivers = unionReceivers(trueRoute.RouteOpts.Receivers, r.RouteOpts.Receivers)
			sortedRoutes.add(trueRoute)
			sortedRoutes.add(r)
		}
	}

	a.aggregatedContinueTrueRoutes = sortedRoutes.routes()
}

func (a *continueConverter) popAggregatedContinueTrueRoutes() []*multiReceiverRoute {
	routes := a.aggregatedContinueTrueRoutes
	a.aggregatedContinueTrueRoutes = nil
	return routes
}

func (a *continueConverter) removeContinueTrueRoutesWithSameMatchers(r *multiReceiverRoute) {
	// We can reset the aggregatedContinueTrueRoutes if this was a "continue: false" route with no matchers. This
	// is because a "continue: false" route with no matchers would match everything and the "continue: true" routes above
	// it should not carry over to the nodes beyond this node.
	if len(r.Matchers) == 0 {
		a.aggregatedContinueTrueRoutes = a.aggregatedContinueTrueRoutes[:0]
		return
	}

	var res []*multiReceiverRoute
	for _, ar := range a.aggregatedContinueTrueRoutes {
		if reflect.DeepEqual(r.Matchers, ar.Matchers) {
			// if the matchers for "r" match a particular route within "aggregatedContinueTrueRoutes" we can remove that from
			// the "aggregatedContinueTrueRoutes"
			continue
		}
		res = append(res, ar)
	}
	a.aggregatedContinueTrueRoutes = res
}

// transformContinueFalseRoute transforms a route with "continue: false" and with or without matchers with the
// currentContinueTrueRoutes and returns the newly created routes.
func (a *continueConverter) transformContinueFalseRoute(falseRoute *multiReceiverRoute) []*multiReceiverRoute {
	if len(a.aggregatedContinueTrueRoutes) == 0 {
		falseRouteCp := a.addInheritedReceivers(falseRoute)
		return []*multiReceiverRoute{falseRouteCp}
	}

	// sortedResultRoutes contains the routes that are created as a result of combining "aggregatedContinueTrueRoutes" with a "continue: false" route
	sortedResultRoutes := routeSortedSet{}
	for _, ar := range a.aggregatedContinueTrueRoutes {
		falseRoute = copyRoute(falseRoute)
		if len(ar.Matchers) > 0 && len(falseRoute.Matchers) > 0 {
			if !reflect.DeepEqual(ar.Matchers, falseRoute.Matchers) {
				// When matchers are different, we generate two routes
				// 1. route corresponding to "falseRoute"
				// 2. merge route that has matchers for "ar" and "falseRoute"
				existingCopy := copyRoute(ar)
				existingCopy.RouteOpts.Receivers = unionReceivers(existingCopy.RouteOpts.Receivers, falseRoute.RouteOpts.Receivers)
				existingCopy.Matchers = append(existingCopy.Matchers, falseRoute.Matchers...)
				sortedResultRoutes.add(existingCopy)
				sortedResultRoutes.add(falseRoute)
			} else {
				// matchers are same, we only need to generate one route that has receivers for both "ar" and "falseRoute"
				falseRoute.RouteOpts.Receivers = unionReceivers(ar.RouteOpts.Receivers, falseRoute.RouteOpts.Receivers)
				sortedResultRoutes.add(falseRoute)
			}
		} else if len(ar.Matchers) == 0 {
			// This covers two cases, 1. when falseRoute has matchers 2. when falseRoute doesn't have matchers.

			// copy over the receivers and the matchers for the falseRoute and add the receivers for "ar" route
			// e.g.
			// receiver a (continue=true)
			//          ↓
			// receiver b (continue=false) (foo=bar)
			//
			// should get combined into
			//
			// receiver a, b (continue=false) (foo=bar)
			falseRoute.RouteOpts.Receivers = unionReceivers(falseRoute.RouteOpts.Receivers, ar.RouteOpts.Receivers)
			sortedResultRoutes.add(falseRoute)
		} else if len(falseRoute.Matchers) == 0 {
			// this is the case where a "continue: true" (ar) node with matchers is followed by a "continue: false" (falseRoute) node
			// without matchers, this would expand to atmost two routes
			// 1. a route with receivers for both "ar" node and "falseRoute" node and matchers for "ar" node
			// 2. a route for the "falseRoute" node if the aggregate route tree doesn't contain a route without matchers
			// See these examples below
			//
			// receiver a (continue=true) (foo=bar)
			//          ↓
			// receiver b (continue=false)
			//
			// should get converted to
			// receiver a, b (continue=false) (foo=bar)
			//          ↓
			// receiver b (continue=false)
			//
			// ----
			//
			// whereas
			// receiver a, b (continue= true) (foo=bar)
			//          ↓
			// receiver b (continue=true)
			//          ↓
			// receiver c (continue=false)
			//
			// should get converted to below because aggregate route tree has a node without any matchers i.e. node with receiver b
			//
			// receiver a, b, c (continue=false) (foo=bar)
			//          ↓
			// receiver b, c (continue=false)

			ar.RouteOpts.Receivers = unionReceivers(ar.RouteOpts.Receivers, falseRoute.RouteOpts.Receivers)
			sortedResultRoutes.add(ar)
			sortedResultRoutes.add(falseRoute)
		}
	}

	return sortedResultRoutes.routes()
}

func unionReceivers(r1 []string, r2 []string) []string {
	receivers := make(map[string]bool)
	for _, r := range r1 {
		receivers[r] = true
	}
	for _, r := range r2 {
		receivers[r] = true
	}
	receiversSlice := maps.Keys(receivers)
	sort.Strings(receiversSlice)
	return receiversSlice
}

func shouldExpandContinueTrueRoute(trueRouteIdx int, siblings []*multiReceiverRoute) (string, bool) {
	// We do not support converting continue=true routes that have children or matchers right now so keep them
	// as they are.
	route := siblings[trueRouteIdx]
	if len(route.Routes) > 0 {
		return warningContinueTrueRouteHasChildren, false
	}

	// Verify that all the siblings following this route node have the same values for these fields. If they don't, then we would not expand
	// this route.
	// Note: Child nodes inherit the group_by (and other fields) from their parent if they do not set one themselves. So we wouldn't
	// convert a configuration like the following:
	// route:
	// group_by=service
	// routes:
	// - receiver="foo"
	//   continue=true
	//   - match
	//     - service="bar"
	// - receiver="bar"
	//   group_by=service
	trueRoute := *route
	trueRoute.RouteOpts.Receivers = nil
	for i := trueRouteIdx + 1; i < len(siblings); i++ {
		r := *(siblings[i])
		r.RouteOpts.Receivers = nil
		// comparing the complete route ensures that we still fail even if a new field is added to multiReceiverRoute
		if !reflect.DeepEqual(r.RouteOpts, trueRoute.RouteOpts) {
			return warningSiblingDifferentValue, false
		}
		if len(route.Matchers) > 0 && (i-trueRouteIdx > 10) {
			return warningTenOrMoreSiblingsFollowingRoute, false
		}
	}
	return "", true
}

// collectReceivers collects all the receivers from the aggregated routes
func collectReceivers(aggregatedRoutes []*multiReceiverRoute) map[string]struct{} {
	routes := make(map[string]struct{})
	for _, route := range aggregatedRoutes {
		for _, r := range route.RouteOpts.Receivers {
			routes[r] = struct{}{}
		}
	}
	return routes
}

// The function expandRoute operates by expanding routes within a route tree that contains nodes with the property "continue: true".
// Here's how the algorithm works:
// 1. As we traverse the nodes at a particular level, route nodes marked with "continue: true" are collected into an array named
// "aggregatedContinueTrueRoutes", which is a part of the continueConverter.
// 2. Whenever we encounter a node (r) marked with "continue: false", whether it has matchers or not, we merge the
// "aggregatedContinueTrueRoutes" array with r to generate an updated list of nodes. If this node has no matchers, we
// reset the "aggregatedContinueTrueRoutes"; otherwise, it remains available for expanding subsequent "continue: false" nodes.
// 3.The receivers for the preceding "continue: true" routes are also propagated to the children of the current node using "ancestorContinueTrueReceivers".

//		receiver: a (continue=true)
//		         ↓
//		receiver: b (continue=true)
//	             ↓
//	    receiver: c (continue=false) (foo=bar) has a child route
//	      => routes
//	        receiver: d (continue=false) (foo=bar, baz=qux)
//
// In this scenario, "a" and "b" would be merged and stored in "aggregatedContinueTrueRoutes".
// When encountering the route with "receiver: c", it would be merged with the "aggregatedContinueTrueRoutes".
// When encountering the route with "receiver: d", it would inherit "a" and "b" as receivers, which would have been
// passed to it as ancestorContinueTrueReceivers.
// The resultant route tree looks like the following
//
//			receiver: a, b, c (continue=false) (foo=bar)
//		      => routes
//		        receiver: a, b, d (continue=false) (foo=bar, baz=qux)
//	                ↓
//		    receiver: a, b (continue=false)
func expandRoute(path fieldPath, route *multiReceiverRoute, ancestorContinueTrueReceivers map[string]struct{}) []string {
	if len(route.Routes) == 0 {
		return nil
	}

	// if there are no routes with "continue: true" in this subtree and also if there are no ancestor routes then there is nothing to do here,
	// we can return early
	if len(ancestorContinueTrueReceivers) == 0 && !hasRouteWithContinueTrue(route) {
		return nil
	}

	a := &continueConverter{
		// the ancestor continue true receivers are inherited by this node
		inheritedContinueTrueReceivers: ancestorContinueTrueReceivers,
	}
	// updatedRoutes contains the results of the converting continue=true routes along with the original routes that were
	// not updated
	updatedRoutes := make([]*multiReceiverRoute, 0)

	var warnings []string
	for idx, r := range route.Routes {
		// if r.Continue is true, add this route to aggregated routes.
		if r.Continue {
			if warning, ok := shouldExpandContinueTrueRoute(idx, route.Routes); !ok {
				// if we don't expand the continue true route, then add the existing route to the updatedRoutes
				// and continue processing the following routes
				updatedRoutes = append(updatedRoutes, r)
				warnings = append(warnings, fmt.Sprintf("%s (%s)", warning, path.field("routes").index(idx)))
				continue
			}
			a.addContinueTrueRoute(r)
		} else {
			// Recursively call the function on child nodes when continue=false for this node
			ancestorContinueTrue := collectReceivers(a.aggregatedContinueTrueRoutes)
			warnings = append(warnings, expandRoute(path.field("routes").index(idx), r, ancestorContinueTrue)...)

			// expand current "continue: false" route along with the currentContinueTrueRoutes from preceding siblings
			res := a.transformContinueFalseRoute(r)
			updatedRoutes = append(updatedRoutes, res...)
			a.removeContinueTrueRoutesWithSameMatchers(r)
		}
	}

	// If there were any remaining aggregated continue true routes that were not matched by a terminal continue:false route without matchers,
	// add them to the route tree.
	if routes := a.popAggregatedContinueTrueRoutes(); len(routes) > 0 {
		updatedRoutes = append(updatedRoutes, routes...)
	}

	route.Routes = updatedRoutes
	return warnings
}

func hasRouteWithContinueTrue(r *multiReceiverRoute) bool {
	if r.Continue {
		return true
	}
	for _, cr := range r.Routes {
		if hasRouteWithContinueTrue(cr) {
			return true
		}
	}
	return false
}

// Identical to dispatch.Route but with references to multiReceiverRoute.
type multiReceiverRoute struct {
	RouteOpts multiReceiverRouteOpts
	Matchers  labels.Matchers
	Continue  bool
	Routes    []*multiReceiverRoute
}

type multiReceiverRouteOpts struct {
	Receivers         []string
	GroupBy           map[model.LabelName]struct{}
	GroupByAll        bool
	GroupWait         time.Duration
	GroupInterval     time.Duration
	RepeatInterval    time.Duration
	MuteTimeIntervals []string
}

func newMultiReceiverRoute(route *dispatch.Route) *multiReceiverRoute {
	childRoutes := make([]*multiReceiverRoute, 0, len(route.Routes))
	for _, r := range route.Routes {
		childRoutes = append(childRoutes, newMultiReceiverRoute(r))
	}
	var receivers []string
	if route.RouteOpts.Receiver != "" {
		receivers = append(receivers, route.RouteOpts.Receiver)
	}
	return &multiReceiverRoute{
		RouteOpts: multiReceiverRouteOpts{
			Receivers:         receivers,
			GroupBy:           route.RouteOpts.GroupBy,
			GroupByAll:        route.RouteOpts.GroupByAll,
			GroupWait:         route.RouteOpts.GroupWait,
			GroupInterval:     route.RouteOpts.GroupInterval,
			RepeatInterval:    route.RouteOpts.RepeatInterval,
			MuteTimeIntervals: route.RouteOpts.MuteTimeIntervals,
		},
		Matchers: route.Matchers,
		Continue: route.Continue,
		Routes:   childRoutes,
	}

}

// convertAlertManagerConfig creates notification policies and receivers from an alertmanager input.
// It attempts to create a notification policy that has overrides based on the matchers in the routes.
func convertAlertManagerConfig(
	path fieldPath,
	opts Opts,
	route *multiReceiverRoute,
	config *config.Config,
) (*models.Configv1NotificationPolicy, []*Notifier, []string, error) {
	var warnings []string
	if !opts.DisableContinueTrueExpansion {
		warnings = append(warnings, expandRoute(path, route, make(map[string]struct{}))...)
	}
	receivers, receiversByName, err := expandReceivers(config.Receivers, opts)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("expand receivers: %w", err)
	}

	// Recursively visit routes to find potential mappings of matchers to notifiers.
	// Default notifiers are included in this set: they will have no matchers defined.
	overrides, w, err := gatherPolicyOverrides(path.field("route"), opts, route, nil, receiversByName)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("gather policy overrides: %w", err)
	}
	warnings = append(warnings, w...)

	// The set of overrides may include duplicates with the same matchers and repeat intervals but different notifiers.
	// For each unique set of matchers, union all the notifiers and add to mergedOverrides.
	mergedOverrides := make([]*policyOverride, 0, len(overrides))
	for len(overrides) > 0 {
		override := overrides[0]
		newOverrides := make([]*policyOverride, 0, len(overrides))
		for i := 1; i < len(overrides); i++ {
			if merged, ok := mergeOverrides(override, overrides[i]); ok {
				override = merged
			} else {
				newOverrides = append(newOverrides, overrides[i])
				continue
			}
		}
		mergedOverrides = append(mergedOverrides, override)
		overrides = newOverrides
	}

	// For overrides with the special allSeverities severity, unset the allSeverities notifiers
	// and copy them to real severities that aren't set.
	// For example, if the override's severities are [warn, allSeverities] then this will set the "critical"
	// severity with the end result of the override having [warn, critical] severities set.
	for _, override := range mergedOverrides {
		if override.AllSeverities == nil {
			continue
		}

		for _, severity := range severities {
			if _, ok := notifierForSeverity(override, severity); !ok {
				setNotifierForSeverity(override, severity, override.AllSeverities)
			}
		}
		override.AllSeverities = nil
	}

	policy := &models.Configv1NotificationPolicy{
		Routes: &models.NotificationPolicyRoutes{
			Overrides: make([]*models.NotificationPolicyRoutesOverride, 0, len(mergedOverrides)),
		},
	}
	policy.Name = "Default"
	policy.Slug = "default"
	if !opts.UseBuckets && opts.NotificationPolicyTeamSlug != "" {
		policy.Name = opts.NotificationPolicyTeamSlug
		policy.Slug = opts.NotificationPolicyTeamSlug
		policy.TeamSlug = opts.NotificationPolicyTeamSlug
	}
	if opts.UseBuckets {
		policy.BucketSlug = "default"
	}

	// Populate the policy routes. Overrides with no matchers are treated as defaults.
	for _, override := range mergedOverrides {
		if len(override.AlertLabelMatchers) == 0 {
			policy.Routes.Defaults = override.Notifiers
		} else {
			policy.Routes.Overrides = append(policy.Routes.Overrides, override.NotificationPolicyRoutesOverride)
		}
	}

	if len(config.InhibitRules) > 0 {
		warnings = append(warnings, "Ignoring inhibit rules")
	}

	return policy, receivers, warnings, nil
}

func findNotifiers(
	overrides []*policyOverride,
	matchers []*models.Configv1LabelMatcher,
	severity string,
	defaultNotifiers *models.RoutesNotifierList,
) *models.RoutesNotifierList {
	for _, sibling := range overrides {
		if reflect.DeepEqual(sibling.AlertLabelMatchers, matchers) {
			if n, ok := notifierForSeverity(sibling, severity); ok {
				return n
			}
		}
	}
	return defaultNotifiers
}

type policyOverride struct {
	*models.NotificationPolicyRoutesOverride
	AllSeverities *models.RoutesNotifierList
}

// gatherPolicyOverrides extracts notification policy overrides from a route and its child route, doing a depth-first
// traversal in order to return overrides with more specific matchers first.
// It returns an ordered list of overrides along with warning messages.
// The list of returned overrides may have entries with duplicate matchers.
func gatherPolicyOverrides(
	path fieldPath,
	opts Opts,
	route *multiReceiverRoute,
	parentMatchers []*models.Configv1LabelMatcher,
	receiversByName map[string][]*Notifier,
) (
	[]*policyOverride,
	[]string,
	error,
) {
	var warnings []string

	if len(route.Routes) > 1 && routesAreLeaves(route.Routes) {
		childHasMatchers := false
		for _, child := range route.Routes {
			_, childMatchers := opts.matchersWithSeverity(child.Matchers)
			if len(childMatchers) > 0 {
				childHasMatchers = true
				break
			}
		}

		childHasContinue := false
		for i := 0; i < len(route.Routes)-1; i++ {
			child := route.Routes[i]
			childHasContinue = childHasContinue || child.Continue

			if !childHasMatchers && !childHasContinue {
				// We implicitly expect a route with multiple leaf-node children to have the first n-1 (at least) children have continue=true
				warnings = append(warnings,
					fmt.Sprintf("Ignoring continue=false on leaf node; will import as continue=true (%s)", path.field("routes").index(i)))
			}
		}
	} else if len(route.Routes) > 0 && route.Continue {
		// Non-leaf nodes should have continue=false
		warnings = append(warnings,
			fmt.Sprintf("%s (%s)", warningContinueTrueRouteHasChildren, path))
	}

	// If matching on a specific severity, remove the severity matcher
	severity, routeMatchers := opts.matchersWithSeverity(route.Matchers)
	if severity == "" {
		severity = allSeverities
	}

	// The matchers for this route is the union of the matchers defined on the route plus those of its ancestors
	matchers, err := convertMatchers(path.field("matchers"), routeMatchers)
	if err != nil {
		return nil, nil, fmt.Errorf("convert matchers: %w", err)
	}

	matchers = mergeMatchers(matchers, parentMatchers)

	// Recursively get the overrides of all child nodes.
	var childOverrides []*policyOverride
	for i, child := range route.Routes {
		overrides, childWarnings, err := gatherPolicyOverrides(path.field("routes").index(i), opts, child, matchers, receiversByName)
		if err != nil {
			return nil, nil, fmt.Errorf("gather policy overrides recursively: %w", err)
		}
		warnings = append(warnings, childWarnings...)

		childOverrides = append(childOverrides, overrides...)
	}

	notifiers := &models.RoutesNotifierList{}
	for _, receiverInRoute := range route.RouteOpts.Receivers {
		if receivers, ok := receiversByName[receiverInRoute]; !ok {
			notifiers.NotifierSlugs = append(notifiers.NotifierSlugs, receiverInRoute)

			warnings = append(warnings,
				fmt.Sprintf("Route references receiver `%s` not defined in config (%s)", receiverInRoute, path))
		} else {
			for _, receiver := range receivers {
				notifiers.NotifierSlugs = append(notifiers.NotifierSlugs, receiver.Slug)
			}
		}
	}

	if route.RouteOpts.RepeatInterval > 0 && route.RouteOpts.RepeatInterval != time.Duration(defaultRepeatInterval) {
		notifiers.RepeatIntervalSecs = int32(route.RouteOpts.RepeatInterval.Seconds())
	}

	if opts.AddGroupByToNotificationPolicies && len(route.RouteOpts.GroupBy) > 0 {
		notifiers.GroupBy = &models.NotificationPolicyRoutesGroupBy{}
		for l := range route.RouteOpts.GroupBy {
			notifiers.GroupBy.LabelNames = append(notifiers.GroupBy.LabelNames, string(l))
		}
		sort.Strings(notifiers.GroupBy.LabelNames)
	}

	// If any children have the same matchers as this route node, just return the children overrides
	// since they're more specific matches and should therefore take precedence.
	for idx, override := range childOverrides {
		if reflect.DeepEqual(override.AlertLabelMatchers, matchers) {
			if _, ok := notifierForSeverity(override, severity); !ok {
				// If this child node doesn't have a notifier corresponding to "severity", check following siblings
				// with the same matchers and "severity" for notifiers.
				// If none of the following siblings have a notifier corresponding to "severity" with the same matchers, then default to
				// notifiers for this route.
				// e.g.
				// receiver: blackhole
				// routes:
				//  - receiver: a
				//    continue: false
				//    match: severity: critical
				//  - receiver: b
				//    continue: false
				//
				// for this case when route node is blackhole, notifiers corresponding to "_allSeverities" should be set
				// to the "b" as its severity and matchers match those of "blackhole"
				setNotifierForSeverity(override, severity, findNotifiers(childOverrides[idx+1:], matchers, severity, notifiers))
			}
			return childOverrides, warnings, nil
		}
	}

	override := &policyOverride{
		NotificationPolicyRoutesOverride: &models.NotificationPolicyRoutesOverride{
			AlertLabelMatchers: matchers,
		},
	}
	setNotifierForSeverity(override, severity, notifiers)

	// This route's override is returned after those of its children, since child nodes are more specific matches.
	return append(childOverrides, override), warnings, nil
}

// routesAreLeaves returns true when each of the routes is a leaf route.
func routesAreLeaves(routes []*multiReceiverRoute) bool {
	for _, route := range routes {
		if len(route.Routes) > 0 {
			return false
		}
	}

	return true
}

// routeMatchers returns all matchers defined on an alertmanager route.
// The matchers include both exact and regex matchers, sorted by name and type.
func convertMatchers(path fieldPath, routeMatchers labels.Matchers) ([]*models.Configv1LabelMatcher, error) {
	matchers := make([]*models.Configv1LabelMatcher, 0, len(routeMatchers))
	for i, matcher := range routeMatchers {
		var matcherType models.Configv1LabelMatcherMatcherType
		switch matcher.Type {
		case labels.MatchEqual:
			matcherType = models.Configv1LabelMatcherMatcherTypeEXACT
		case labels.MatchRegexp:
			matcherType = models.Configv1LabelMatcherMatcherTypeREGEX
		default:
			return nil, path.index(i).invalidReason("Unsupported matcher type `%v`", matcher.Type)
		}

		matchers = append(matchers, &models.Configv1LabelMatcher{
			Type:  matcherType,
			Name:  matcher.Name,
			Value: matcher.Value,
		})
	}

	return sortMatchers(matchers), nil
}

// mergeMatchers returns a new matcher that is a deduped union of two matchers.
// The merged matchers are ordered by name and type.
func mergeMatchers(a, b []*models.Configv1LabelMatcher) []*models.Configv1LabelMatcher {
	matchers := make([]*models.Configv1LabelMatcher, len(a), len(a)+len(b))
	copy(matchers, a)

	for _, matcher := range b {
		dupe := false
		for _, m := range matchers {
			if m.Name == matcher.Name && m.Type == matcher.Type && m.Value == matcher.Value {
				dupe = true
				break
			}
		}
		if !dupe {
			matchers = append(matchers, matcher)
		}
	}

	return sortMatchers(matchers)
}

// sortMatchers sorts matchers by [name, type]
func sortMatchers(matchers []*models.Configv1LabelMatcher) []*models.Configv1LabelMatcher {
	sort.SliceStable(matchers, func(i, j int) bool {
		if matchers[i].Name < matchers[j].Name {
			return true
		}

		return matchers[i].Type < matchers[j].Type
	})

	return matchers
}

// mergeOverrides attempts to combine two overrides into one, which is allowed in the case where
// both overrides have identical matchers and the same repeat intervals.
// This returns either the merged override and true, or nil and false.
func mergeOverrides(a, b *policyOverride) (*policyOverride, bool) {
	if !reflect.DeepEqual(a.AlertLabelMatchers, b.AlertLabelMatchers) {
		return nil, false
	}

	override := &policyOverride{
		NotificationPolicyRoutesOverride: &models.NotificationPolicyRoutesOverride{
			AlertLabelMatchers: a.AlertLabelMatchers,
			Notifiers: &models.RoutesSeverityNotifiers{
				Warn:     copyNotifierList(a.Notifiers.Warn),
				Critical: copyNotifierList(a.Notifiers.Critical),
			},
		},
		AllSeverities: copyNotifierList(a.AllSeverities),
	}
	var ok bool
	override.Notifiers.Warn, ok = mergeNotifiers(override.Notifiers.Warn, b.Notifiers.Warn)
	if !ok {
		return nil, false
	}
	override.Notifiers.Critical, ok = mergeNotifiers(override.Notifiers.Critical, b.Notifiers.Critical)
	if !ok {
		return nil, false
	}
	override.AllSeverities, ok = mergeNotifiers(override.AllSeverities, b.AllSeverities)
	if !ok {
		return nil, false
	}

	return override, true
}

func mergeNotifiers(trunk, toMerge *models.RoutesNotifierList) (*models.RoutesNotifierList, bool) {
	trunk = copyNotifierList(trunk)
	if toMerge == nil {
		return trunk, true
	}
	if trunk == nil {
		trunk = &models.RoutesNotifierList{}
	} else if trunk.RepeatIntervalSecs != toMerge.RepeatIntervalSecs {
		return nil, false
	} else if trunk.GroupBy != toMerge.GroupBy {
		return nil, false
	}
	for _, n := range toMerge.NotifierSlugs {
		if !slices.Contains(trunk.NotifierSlugs, n) {
			trunk.NotifierSlugs = append(trunk.NotifierSlugs, n)
		}
	}
	return trunk, true
}

type convertPrometheusRulesResult struct {
	entityGroups   []*models.Configv1Collection
	collections    []*models.Configv1Collection
	monitors       []*models.Configv1Monitor
	recordingRules []*models.Configv1RecordingRule
}

func newSlugGenerators() slugGenerators {
	return slugGenerators{
		monitor:       newSlugGenerator(),
		recordingRule: newSlugGenerator(),
		collection:    newSlugGenerator(),
		entityGroup:   newSlugGenerator(),
	}
}

type slugGenerators struct {
	monitor       *slugGenerator
	recordingRule *slugGenerator
	collection    *slugGenerator
	entityGroup   *slugGenerator
}

// convertPrometheusRuleGroups converts Prometheus rule groups to entity groups, monitors, and recording groups.
// An entity group is created from each rule group. Each alert rule within a rule group is converted to
// a separate monitor.
// Each rule group with recording rules is converted to a recording group.
func convertPrometheusRuleGroups(
	path fieldPath,
	opts Opts,
	promRules *rulefmt.RuleGroups,
) (*convertPrometheusRulesResult, error) {
	slugs := newSlugGenerators()

	var monitors []*models.Configv1Monitor

	// number of times a group with the same name has been seen
	groupCounts := make(map[string]int)

	// used if opts.UseBuckets == true
	var entityGroups []*models.Configv1Collection

	// used if opts.UseBuckets == false
	var collections []*models.Configv1Collection
	var recRules []*models.Configv1RecordingRule
	for i, ruleGroup := range promRules.Groups {
		ruleGroupPath := path.field("groups").index(i)
		if ruleGroup.Name == "" {
			return nil, ruleGroupPath.field("name").valueRequired()
		}

		// If multiple groups have the same name, suffix all but the first group with a counter.
		groupName := ruleGroup.Name
		if seenCount := groupCounts[ruleGroup.Name]; seenCount > 0 {
			groupName = fmt.Sprintf("%s (%d)", ruleGroup.Name, seenCount)
		}

		if !opts.UseBuckets {
			// Create the slug first to pass into the convertRules function
			collectionSlug := slugs.collection.generateV1(ruleGroup.Name)

			// Convert the rules to recording rules and monitors
			groupRecRules, groupMonitors, err := convertRules(ruleGroup, ruleGroupPath, opts, slugs, collectionSlug)
			if err != nil {
				return nil, fmt.Errorf("convert rules, use_collections=true: %w", err)
			}

			// Create a collection for this rule group only if it contains monitors. Recording rules do not need collections.
			if len(groupMonitors) > 0 {
				collection := &models.Configv1Collection{
					Name:     groupName,
					TeamSlug: opts.groupTeamMatcher.Match(ruleGroup.Name),
					Slug:     collectionSlug,
				}
				collections = append(collections, collection)
				groupCounts[ruleGroup.Name]++
			}

			// If NoUseCollections == false, converting recording rules are returned as a list of recording rules
			recRules = append(recRules, groupRecRules...)
			monitors = append(monitors, uniquifyMonitorNames(groupMonitors)...)
		} else {
			// Create an entity group for this rule group
			entityGroup := &models.Configv1Collection{
				Name: groupName,
				Slug: slugs.entityGroup.generateV1(ruleGroup.Name),
			}
			entityGroups = append(entityGroups, entityGroup)
			groupCounts[ruleGroup.Name]++

			// Convert the rules to recording rules and monitors
			groupRecRules, groupMonitors, err := convertRules(ruleGroup, ruleGroupPath, opts, slugs, entityGroup.Slug)
			if err != nil {
				return nil, fmt.Errorf("convert rules, use_collections=false: %w", err)
			}
			recRules = append(recRules, groupRecRules...)

			monitors = append(monitors, uniquifyMonitorNames(groupMonitors)...)
		}

	}

	return &convertPrometheusRulesResult{
		entityGroups:   entityGroups,
		monitors:       monitors,
		recordingRules: recRules,
		collections:    collections,
	}, nil
}

func convertRules(
	ruleGroup rulefmt.RuleGroup,
	ruleGroupPath fieldPath,
	opts Opts,
	slugs slugGenerators,
	parentSlug string,
) (
	[]*models.Configv1RecordingRule,
	[]*models.Configv1Monitor,
	error,
) {
	var recRules []*models.Configv1RecordingRule
	var monitors []*models.Configv1Monitor
	for i, rule := range ruleGroup.Rules {
		rulePath := ruleGroupPath.field("rules").index(i)

		if rule.Alert == "" && rule.Record == "" {
			return nil, nil, rulePath.invalidReason("Either alert or record field must be set")
		}

		if rule.Record != "" {
			rr := convertRecordingRule(rule, ruleGroup)
			rr.Slug = slugs.recordingRule.generateV2(rr.Name)
			recRules = append(recRules, rr)
		} else {
			var monitorSlug string
			if !opts.DisableMonitorSlugAssignment {
				monitorSlug = slugs.monitor.generateV1(rule.Alert)
			}

			monitor, err := convertAlertingRule(rule, monitorSlug, parentSlug, ruleGroup, rulePath, opts)
			if err != nil {
				return nil, nil, fmt.Errorf("converting alerting rule: %w", err)
			}
			monitors = append(monitors, monitor)
		}
	}
	return recRules, monitors, nil
}

func convertAlertingRule(
	rule rulefmt.Rule,
	slug string,
	parentSlug string,
	ruleGroup rulefmt.RuleGroup,
	rulePath fieldPath,
	opts Opts,
) (*models.Configv1Monitor, error) {
	// Default severity to critical (if label is absent). This is an arbitrary decision.
	nonSeverityLabels := make(map[string]string, 0)
	severityLabel := criticalLabel
	severityLabelName := opts.severityLabelName()
	for k, v := range rule.Labels {
		if k == severityLabelName {
			severityLabel = v
			continue
		}
		nonSeverityLabels[k] = v
	}
	severity, ok := opts.severity(severityLabel)
	if !ok {
		return nil, rulePath.field("labels").invalidReason("invalid %s value `%s`", chronoSeverityLabelName, severityLabel)
	}

	expr, op, threshold, err := parseExpr(rulePath.field("expr"), rule.Expr, !opts.ExistsOpNotSupported)
	if err != nil {
		return nil, fmt.Errorf("parse expression: %w", err)
	}

	var signalGrouping *models.MonitorSignalGrouping
	if opts.MonitorSignalPerSeries {
		signalGrouping = &models.MonitorSignalGrouping{
			SignalPerSeries: true,
		}
	}

	seriesConditions := &models.MonitorSeriesConditions{
		Defaults: &models.SeriesConditionsSeverityConditions{},
	}
	condition := &models.SeriesConditionsConditions{
		Conditions: []*models.MonitorCondition{
			{
				Op:                 op,
				Value:              threshold,
				SustainSecs:        durationSeconds(rule.For),
				ResolveSustainSecs: durationSeconds(rule.KeepFiringFor),
			},
		},
	}

	if severity == warnLabel {
		seriesConditions.Defaults.Warn = condition
	} else {
		seriesConditions.Defaults.Critical = condition
	}

	monitor := &models.Configv1Monitor{
		Slug:             slug,
		Name:             rule.Alert,
		Labels:           nonSeverityLabels,
		Annotations:      rule.Annotations,
		PrometheusQuery:  strings.TrimSpace(expr),
		SignalGrouping:   signalGrouping,
		SeriesConditions: seriesConditions,
		IntervalSecs:     durationSeconds(ruleGroup.Interval),
	}

	if opts.UseBuckets {
		monitor.BucketSlug = parentSlug
	} else {
		monitor.CollectionSlug = parentSlug
	}

	return monitor, nil
}

func convertRecordingRule(
	rule rulefmt.Rule,
	ruleGroup rulefmt.RuleGroup,
) *models.Configv1RecordingRule {
	return &models.Configv1RecordingRule{
		Name:           rule.Record,
		MetricName:     rule.Record,
		PrometheusExpr: rule.Expr,
		IntervalSecs:   durationSeconds(ruleGroup.Interval),
		LabelPolicy: &models.Configv1RecordingRuleLabelPolicy{
			Add: rule.Labels,
		},
		ExecutionGroup: ruleGroup.Name,
	}
}

func uniquifyMonitorNames(monitors []*models.Configv1Monitor) []*models.Configv1Monitor {
	names := make([]string, 0, len(monitors))
	monsByName := make(map[string][]*models.Configv1Monitor)
	updatedMonitors := make([]*models.Configv1Monitor, 0, len(monitors))
	for _, mon := range monitors {
		if _, ok := monsByName[mon.Name]; !ok {
			names = append(names, mon.Name)
		}
		monsByName[mon.Name] = append(monsByName[mon.Name], mon)
	}

	for _, name := range names {
		mons := monsByName[name]
		if len(mons) > 1 {
			for i, mon := range mons {
				mon.Name = fmt.Sprintf("%s (%d)", mon.Name, i+1)
			}
		}
		updatedMonitors = append(updatedMonitors, mons...)
	}

	return updatedMonitors
}

func setReceiverSlugs(
	receivers []*Notifier,
	opts Opts,
) error {
	if len(receivers) == 0 {
		return nil
	}

	// For each receiver that doesn't already exist, attempt a best-effort slug assignment
	receiverSlugs := make(map[string]struct{})
	for _, slug := range opts.ReceiverNamesToSlugs {
		receiverSlugs[slug] = struct{}{}
	}

	for i, receiver := range receivers {
		if receiver.Name == "" {
			return fieldPath("notifiers").index(i).invalidReason("name must be set")
		}

		if receiver.Slug != "" {
			receiverSlugs[receiver.Slug] = struct{}{}
		}
	}

	receiverSlugGenerator := newSlugGeneratorWithUsed(receiverSlugs)
	for _, receiver := range receivers {
		if receiver.Slug != "" {
			continue
		}

		if slug, ok := opts.ReceiverNamesToSlugs[receiver.Name]; ok {
			receiver.Slug = slug
		} else {
			receiver.Slug = receiverSlugGenerator.generateV2(receiver.Name)
		}
	}

	return nil
}

// expandReceivers takes a set of receivers, some of which may define multiple targets, and returns a new set where
// each receiver has at most one target. Returned receivers also have slugs set.
// The returned map indexes the original receivers names the new receivers for that name.
func expandReceivers(
	receivers []config.Receiver,
	opts Opts,
) (
	[]*Notifier,
	map[string][]*Notifier,
	error,
) {
	var allReceivers []*Notifier
	recvByInputName := make(map[string][]*Notifier, len(receivers))

	for _, receiver := range receivers {
		expandedReceivers := expandReceiver(receiver)
		recvByInputName[receiver.Name] = append(recvByInputName[receiver.Name], expandedReceivers...)
		allReceivers = append(allReceivers, expandedReceivers...)
	}

	if err := setReceiverSlugs(allReceivers, opts); err != nil {
		return nil, nil, fmt.Errorf("set receiver slugs: %w", err)
	}

	return allReceivers, recvByInputName, nil
}

func expandReceiver(r config.Receiver) []*Notifier {
	receivers := make([]*Notifier, 0)
	for _, w := range r.WebhookConfigs {
		receivers = append(receivers, &Notifier{
			Configv1Notifier: &models.Configv1Notifier{
				Name:         r.Name,
				SkipResolved: !w.SendResolved(),
				Webhook:      webhookReceiver(w),
			},
		})
	}

	for _, e := range r.EmailConfigs {
		receivers = append(receivers, &Notifier{
			Configv1Notifier: &models.Configv1Notifier{
				Name:         r.Name,
				SkipResolved: !e.SendResolved(),
				Email:        emailReceiver(e),
			},
		})
	}

	for _, s := range r.SlackConfigs {
		receivers = append(receivers, &Notifier{
			Configv1Notifier: &models.Configv1Notifier{
				Name:         r.Name,
				SkipResolved: !s.SendResolved(),
				Slack:        slackReceiver(s),
			},
		})
	}

	for _, p := range r.PagerdutyConfigs {
		receivers = append(receivers, &Notifier{
			Configv1Notifier: &models.Configv1Notifier{
				Name:         r.Name,
				SkipResolved: !p.SendResolved(),
				Pagerduty:    pagerdutyReceiver(p),
			},
			Source: p.Source,
		})
	}

	for _, o := range r.OpsGenieConfigs {
		receivers = append(receivers, &Notifier{
			Configv1Notifier: &models.Configv1Notifier{
				Name:         r.Name,
				SkipResolved: !o.SendResolved(),
				OpsGenie:     opsGenieReceiver(o),
			},
		})
	}

	for _, v := range r.VictorOpsConfigs {
		receivers = append(receivers, &Notifier{
			Configv1Notifier: &models.Configv1Notifier{
				Name:         r.Name,
				SkipResolved: !v.SendResolved(),
				VictorOps:    victorOpsReceiver(v),
			},
		})
	}

	if len(receivers) == 0 {
		receivers = append(receivers, &Notifier{Configv1Notifier: &models.Configv1Notifier{Name: r.Name}})
	}

	if len(receivers) > 1 {
		for i := range receivers {
			receivers[i].Name = fmt.Sprintf("%s (%d)", receivers[i].Name, i)
		}
	}

	return receivers
}

func webhookReceiver(w *config.WebhookConfig) *models.NotifierWebhookConfig {
	return &models.NotifierWebhookConfig{
		URL:        w.URL.String(),
		HTTPConfig: fromHTTPClientConfig(w.HTTPConfig),
	}
}

func emailReceiver(e *config.EmailConfig) *models.NotifierEmailConfig {
	return &models.NotifierEmailConfig{
		To:   e.To,
		Text: e.Text,
		HTML: e.HTML,
	}
}

func slackReceiver(s *config.SlackConfig) *models.NotifierSlackConfig {
	apiURL := ""
	if s.APIURL != nil {
		apiURL = s.APIURL.String()
	}
	cfg := &models.NotifierSlackConfig{
		HTTPConfig:  fromHTTPClientConfig(s.HTTPConfig),
		APIURL:      apiURL,
		Channel:     s.Channel,
		Username:    s.Username,
		Color:       s.Color,
		Title:       s.Title,
		TitleLink:   s.TitleLink,
		Pretext:     s.Pretext,
		Text:        s.Text,
		Footer:      s.Footer,
		Fallback:    s.Fallback,
		CallbackID:  s.CallbackID,
		IconEmoji:   s.IconEmoji,
		IconURL:     s.IconURL,
		ImageURL:    s.ImageURL,
		ThumbURL:    s.ThumbURL,
		ShortFields: s.ShortFields,
		LinkNames:   s.LinkNames,
	}
	cfg.Fields = make([]*models.SlackConfigField, 0, len(s.Fields))
	for _, elem := range s.Fields {
		cfg.Fields = append(cfg.Fields, fromSlackField(elem))
	}
	cfg.Actions = make([]*models.NotifierSlackConfigAction, 0, len(s.Actions))
	for _, elem := range s.Actions {
		cfg.Actions = append(cfg.Actions, fromSlackAction(elem))
	}
	return cfg
}

func pagerdutyReceiver(p *config.PagerdutyConfig) *models.NotifierPagerdutyConfig {
	cfg := &models.NotifierPagerdutyConfig{
		HTTPConfig:  fromHTTPClientConfig(p.HTTPConfig),
		URL:         p.URL.String(),
		ServiceKey:  string(p.ServiceKey),
		RoutingKey:  string(p.RoutingKey),
		Client:      p.Client,
		ClientURL:   p.ClientURL,
		Description: p.Description,
		Severity:    p.Severity,
		Class:       p.Class,
		Component:   p.Component,
		Group:       p.Group,
		Details:     p.Details,
	}
	cfg.Links = make([]*models.PagerdutyConfigLink, 0, len(p.Links))
	for _, l := range p.Links {
		cfg.Links = append(cfg.Links, fromPagerdutyLink(l))
	}
	cfg.Images = make([]*models.PagerdutyConfigImage, 0, len(p.Images))
	for _, i := range p.Images {
		cfg.Images = append(cfg.Images, fromPagerdutyImage(i))
	}
	return cfg
}

func opsGenieReceiver(o *config.OpsGenieConfig) *models.NotifierOpsGenieConfig {
	cfg := &models.NotifierOpsGenieConfig{
		HTTPConfig:  fromHTTPClientConfig(o.HTTPConfig),
		APIKey:      string(o.APIKey),
		APIURL:      o.APIURL.String(),
		Message:     o.Message,
		Description: o.Description,
		Source:      o.Source,
		Details:     o.Details,
		Tags:        o.Tags,
		Note:        o.Note,
		Priority:    o.Priority,
	}

	cfg.Responders = make([]*models.OpsGenieConfigResponder, 0, len(o.Responders))
	for _, responder := range o.Responders {
		cfg.Responders = append(cfg.Responders, fromOpsGenieResponder(responder))
	}
	return cfg
}

func victorOpsReceiver(v *config.VictorOpsConfig) *models.NotifierVictorOpsConfig {
	return &models.NotifierVictorOpsConfig{
		HTTPConfig:        fromHTTPClientConfig(v.HTTPConfig),
		APIKey:            string(v.APIKey),
		APIURL:            v.APIURL.String(),
		RoutingKey:        v.RoutingKey,
		MessageType:       v.MessageType,
		EntityDisplayName: v.EntityDisplayName,
		StateMessage:      v.StateMessage,
		MonitoringTool:    v.MonitoringTool,
		CustomFields:      v.CustomFields,
	}
}

func fromHTTPClientConfig(
	value *commoncfg.HTTPClientConfig,
) *models.NotifierHTTPConfig {
	result := &models.NotifierHTTPConfig{}
	if value == nil {
		return result
	}

	if elem := value.BasicAuth; elem != nil {
		result.BasicAuth = &models.HTTPConfigBasicAuth{
			Username: elem.Username,
			Password: string(elem.Password),
		}
	}

	if elem := value.BearerToken; elem != "" {
		result.BearerToken = string(elem)
	}

	if elem := value.ProxyURL.URL; elem != nil {
		result.ProxyURL = elem.String()
	}

	result.TLSConfig = &models.HTTPConfigTLSConfig{
		InsecureSkipVerify: value.TLSConfig.InsecureSkipVerify,
	}

	return result
}

func fromSlackField(
	elem *config.SlackField,
) *models.SlackConfigField {
	short := false
	if elem.Short != nil {
		short = *elem.Short
	}
	return &models.SlackConfigField{
		Title: elem.Title,
		Value: elem.Value,
		Short: short,
	}
}

func fromSlackAction(
	elem *config.SlackAction,
) *models.NotifierSlackConfigAction {
	return &models.NotifierSlackConfigAction{
		Type:         elem.Type,
		Text:         elem.Text,
		URL:          elem.URL,
		Style:        elem.Style,
		Name:         elem.Name,
		Value:        elem.Value,
		ConfirmField: fromSlackConfirmationField(elem.ConfirmField),
	}
}

func fromPagerdutyImage(
	elem config.PagerdutyImage,
) *models.PagerdutyConfigImage {
	return &models.PagerdutyConfigImage{
		Src:  elem.Src,
		Alt:  elem.Alt,
		Href: elem.Href,
	}
}

func fromPagerdutyLink(
	elem config.PagerdutyLink,
) *models.PagerdutyConfigLink {
	return &models.PagerdutyConfigLink{
		Href: elem.Href,
		Text: elem.Text,
	}
}

func fromOpsGenieResponder(
	elem config.OpsGenieConfigResponder,
) *models.OpsGenieConfigResponder {
	var responderType models.ResponderResponderType
	switch elem.Type {
	case "team":
		responderType = models.ResponderResponderTypeTEAM
	case "user":
		responderType = models.ResponderResponderTypeUSER
	case "escalation":
		responderType = models.ResponderResponderTypeESCALATION
	case "schedule":
		responderType = models.ResponderResponderTypeSCHEDULE
	}

	return &models.OpsGenieConfigResponder{
		ID:            elem.ID,
		Name:          elem.Name,
		Username:      elem.Username,
		ResponderType: responderType,
	}
}

func fromSlackConfirmationField(
	elem *config.SlackConfirmationField,
) *models.SlackConfigConfirmationField {
	if elem == nil {
		return nil
	}
	return &models.SlackConfigConfirmationField{
		Text:        elem.Text,
		Title:       elem.Title,
		OkText:      elem.OkText,
		DismissText: elem.DismissText,
	}
}

type fieldPath string

// Field returns a path consisting of a field off the current path.
func (p fieldPath) field(fieldName string) fieldPath {
	if p == "" {
		return fieldPath(fieldName)
	}

	return fieldPath(fmt.Sprintf("%s.%s", p, fieldName))
}

func (p fieldPath) index(idx int) fieldPath {
	return fieldPath(fmt.Sprintf("%s[%d]", p, idx))
}

func (p fieldPath) valueRequired() error {
	return fmt.Errorf("Field `%s` must be set", string(p))
}

func (p fieldPath) invalidReason(format string, args ...interface{}) error {
	return fmt.Errorf("Invalid value for field `%s`: %s", string(p), fmt.Sprintf(format, args...))
}

func validSeverity(severity string) bool {
	return slices.Contains(severities, severity)
}

func newSlugGeneratorWithUsed(used map[string]struct{}) *slugGenerator {
	return &slugGenerator{used: used}
}

func newSlugGenerator() *slugGenerator {
	return newSlugGeneratorWithUsed(make(map[string]struct{}, 0))
}

type slugGenerator struct {
	used map[string]struct{}
}

func (g *slugGenerator) generateV1(name string) string {
	return g.generate(name, 36, false)
}

func (g *slugGenerator) generateV2(name string) string {
	return g.generate(name, 200, true)
}

func (g *slugGenerator) generate(name string, maxLength int, convertUnicodeToASCII bool) string {
	if convertUnicodeToASCII {
		name = anyascii.Transliterate(name)
	}

	baseSlug := strings.ToLower(name)
	// To ensure a value like -=-= gets collapsed into a single -, we replace it with
	// an invalid character that then becomes a separator.
	const unsupported = " "
	baseSlug = strings.ReplaceAll(baseSlug, separator, unsupported)
	baseSlug = validNonPreferred.ReplaceAllString(baseSlug, unsupported)

	// Ideally, we'd replace only non-matching characters which API doesn't support
	// so we instead find all contiguous matches, and put a single separator in between.
	baseSlug = strings.Join(validSlugChars.FindAllString(baseSlug, -1), separator)

	if len(baseSlug) > maxLength {
		baseSlug = baseSlug[:maxLength]
	}

	slug := baseSlug
	for index := 1; ; index++ {
		unique := g.uniqueCheck(slug)
		if slug == "" || !unique {
			slug = fmt.Sprintf("%v-%d", baseSlug, index)
			continue
		}
		break
	}
	return slug
}

func (g *slugGenerator) uniqueCheck(slug string) bool {
	if _, ok := g.used[slug]; ok {
		return false
	}
	g.used[slug] = struct{}{}
	return true
}

func notifierForSeverity(override *policyOverride, severity string) (*models.RoutesNotifierList, bool) {
	if override.Notifiers == nil {
		return nil, false
	}
	switch severity {
	case warnLabel:
		if override.Notifiers.Warn != nil {
			return override.Notifiers.Warn, true
		}
		return nil, false
	case criticalLabel:
		if override.Notifiers.Critical != nil {
			return override.Notifiers.Critical, true
		}
		return nil, false
	case allSeverities:
		if override.AllSeverities != nil {
			return override.AllSeverities, true
		}
		return nil, false
	default:
		return nil, false
	}
}

func setNotifierForSeverity(override *policyOverride, severity string, notifiers *models.RoutesNotifierList) {
	if override.Notifiers == nil {
		override.Notifiers = &models.RoutesSeverityNotifiers{}
	}
	switch severity {
	case warnLabel:
		override.Notifiers.Warn = notifiers
	case criticalLabel:
		override.Notifiers.Critical = notifiers
	case allSeverities:
		override.AllSeverities = notifiers
	}
}

func copyNotifierList(nl *models.RoutesNotifierList) *models.RoutesNotifierList {
	if nl == nil {
		return nil
	}
	return &models.RoutesNotifierList{
		NotifierSlugs:      slices.Clone(nl.NotifierSlugs),
		RepeatIntervalSecs: nl.RepeatIntervalSecs,
		GroupBy:            nl.GroupBy,
	}
}

func durationSeconds(d model.Duration) int32 {
	return int32(time.Duration(d).Seconds())
}
