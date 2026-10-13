// Copyright 2023 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dashboards

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
)

const (
	testGrafanaJSON   = `{"title":"API Latency Overview","panels":[]}`
	dashboardJSONLine = `dashboard_json: '` + testGrafanaJSON + `'`
	multipleItemsErr  = `input contains more than one item; import one dashboard per run, e.g. pipe one "classic-dashboards read <slug>" at a time rather than "list"`
)

// manifestYAML builds a v1/config manifest of the given kind with one spec
// field per line.
func manifestYAML(kind string, spec ...string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "api_version: v1/config\nkind: %s\nspec:\n", kind)
	for _, line := range spec {
		b.WriteString("  " + line + "\n")
	}
	return b.String()
}

// manifestRequest is the request a manifest carrying testGrafanaJSON maps to.
func manifestRequest(source, collectionSlug, name, slug string) *importRequest {
	return &importRequest{
		source: source,
		body: &models.Configv1ImportDashboardFromClassicRequest{
			ClassicDashboardJSON: testGrafanaJSON,
			CollectionSlug:       collectionSlug,
			Name:                 name,
			Slug:                 slug,
		},
	}
}

func TestNewImportRequest(t *testing.T) {
	classicYAML := manifestYAML("ClassicDashboard",
		"slug: xyz",
		"name: API Latency Overview",
		`created_at: "2026-01-02T03:04:05.000Z"`,
		`updated_at: "2026-01-02T03:04:05.000Z"`,
		"collection_slug: team-a",
		dashboardJSONLine,
	)
	classicJSON := `{"api_version":"v1/config","kind":"ClassicDashboard","spec":{"slug":"xyz","name":"API Latency Overview","collection_slug":"team-a","dashboard_json":` + fmt.Sprintf("%q", testGrafanaJSON) + `}}`
	grafanaJSON := strings.Replace(classicJSON, `"kind":"ClassicDashboard"`, `"kind":"GrafanaDashboard"`, 1)

	tests := []struct {
		name    string
		content string
		flags   importFlags
		want    *importRequest
		wantErr string
	}{
		{
			name:    "raw JSON object",
			content: testGrafanaJSON,
			flags:   importFlags{filename: "grafana.json", collectionSlug: "team-a", dryRun: true},
			want: &importRequest{
				source: "-f grafana.json",
				body: &models.Configv1ImportDashboardFromClassicRequest{
					ClassicDashboardJSON: testGrafanaJSON,
					CollectionSlug:       "team-a",
					DryRun:               true,
				},
			},
		},
		{
			name:    "raw JSON with name, slug and update-if-exists",
			content: testGrafanaJSON,
			flags:   importFlags{filename: "-", collectionSlug: "team-a", name: "Latency", slug: "latency", updateIfExists: true},
			want: &importRequest{
				source: "-f -",
				body: &models.Configv1ImportDashboardFromClassicRequest{
					ClassicDashboardJSON: testGrafanaJSON,
					CollectionSlug:       "team-a",
					Name:                 "Latency",
					Slug:                 "latency",
					UpdateIfExists:       true,
				},
			},
		},
		{
			name:    "raw JSON bytes are sent verbatim",
			content: "\n  {\n\t\"title\" :  \"API Latency Overview\",\n  \"panels\": [ ]\n}\n\n",
			flags:   importFlags{filename: "grafana.json", collectionSlug: "team-a"},
			want: &importRequest{
				source: "-f grafana.json",
				body: &models.Configv1ImportDashboardFromClassicRequest{
					ClassicDashboardJSON: "\n  {\n\t\"title\" :  \"API Latency Overview\",\n  \"panels\": [ ]\n}\n\n",
					CollectionSlug:       "team-a",
				},
			},
		},
		{
			name:    "raw JSON without collection slug",
			content: testGrafanaJSON,
			flags:   importFlags{filename: "grafana.json"},
			wantErr: "--collection-slug is required",
		},
		{
			name:    "raw JSON with update-if-exists and no slug",
			content: testGrafanaJSON,
			flags:   importFlags{filename: "grafana.json", collectionSlug: "team-a", updateIfExists: true},
			wantErr: "--slug is required when --update-if-exists is set",
		},
		{
			name:    "invalid JSON object",
			content: `{"title": "API Latency Overview",`,
			flags:   importFlags{filename: "grafana.json", collectionSlug: "team-a"},
			wantErr: "expected raw Grafana dashboard JSON or a ClassicDashboard or GrafanaDashboard manifest with api_version and kind",
		},
		{
			name:    "empty input",
			content: "",
			wantErr: "input is empty",
		},
		{
			name:    "whitespace-only input",
			content: " \n\t\n",
			wantErr: "input is empty",
		},
		{
			name:    "ClassicDashboard YAML",
			content: classicYAML,
			want:    manifestRequest(`ClassicDashboard "xyz"`, "team-a", "API Latency Overview", "xyz"),
		},
		{
			name:    "ClassicDashboard JSON",
			content: classicJSON,
			want:    manifestRequest(`ClassicDashboard "xyz"`, "team-a", "API Latency Overview", "xyz"),
		},
		{
			name:    "GrafanaDashboard YAML",
			content: strings.Replace(classicYAML, "kind: ClassicDashboard", "kind: GrafanaDashboard", 1),
			want:    manifestRequest(`GrafanaDashboard "xyz"`, "team-a", "API Latency Overview", "xyz"),
		},
		{
			name:    "GrafanaDashboard JSON",
			content: grafanaJSON,
			want:    manifestRequest(`GrafanaDashboard "xyz"`, "team-a", "API Latency Overview", "xyz"),
		},
		{
			name:    "dry run and update-if-exists come from flags",
			content: classicYAML,
			flags:   importFlags{dryRun: true, updateIfExists: true},
			want: &importRequest{
				source: `ClassicDashboard "xyz"`,
				body: &models.Configv1ImportDashboardFromClassicRequest{
					ClassicDashboardJSON: testGrafanaJSON,
					CollectionSlug:       "team-a",
					Name:                 "API Latency Overview",
					Slug:                 "xyz",
					DryRun:               true,
					UpdateIfExists:       true,
				},
			},
		},
		{
			name:    "one-element JSON array",
			content: "[" + classicJSON + "]",
			want:    manifestRequest(`ClassicDashboard "xyz"`, "team-a", "API Latency Overview", "xyz"),
		},
		{
			name:    "empty JSON array",
			content: "[]",
			wantErr: "input does not contain any items",
		},
		{
			name:    "YAML without api_version",
			content: "title: API Latency Overview\npanels: []\n",
			flags:   importFlags{collectionSlug: "team-a"},
			wantErr: "expected raw Grafana dashboard JSON or a ClassicDashboard or GrafanaDashboard manifest with api_version and kind",
		},
		{
			name:    "unparseable YAML",
			content: "title: [unclosed\n",
			wantErr: "expected raw Grafana dashboard JSON or a ClassicDashboard or GrafanaDashboard manifest: could not parse input as YAML: ",
		},
		{
			name:    "YAML stream of two",
			content: classicYAML + "---\n" + classicYAML,
			wantErr: multipleItemsErr,
		},
		{
			name:    "two-element JSON array",
			content: "[" + classicJSON + "," + grafanaJSON + "]",
			wantErr: multipleItemsErr,
		},
		{
			name:    "two JSON lines",
			content: classicJSON + "\n" + grafanaJSON + "\n",
			wantErr: multipleItemsErr,
		},
		{
			name: "collection_slug wins over collection and bucket_slug",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine,
				"collection_slug: from-collection-slug",
				"collection: {type: SIMPLE, slug: from-collection}",
				"bucket_slug: from-bucket",
			),
			want: manifestRequest(`ClassicDashboard "xyz"`, "from-collection-slug", "", "xyz"),
		},
		{
			name: "SIMPLE collection reference",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine,
				"collection: {type: SIMPLE, slug: from-collection}",
				"bucket_slug: from-bucket",
			),
			want: manifestRequest(`ClassicDashboard "xyz"`, "from-collection", "", "xyz"),
		},
		{
			name: "collection reference with no type",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine,
				"collection: {slug: from-collection}",
			),
			want: manifestRequest(`ClassicDashboard "xyz"`, "from-collection", "", "xyz"),
		},
		{
			name: "SERVICE collection reference without --collection-slug",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine,
				"collection: {type: SERVICE, slug: checkout}",
			),
			wantErr: `ClassicDashboard "xyz" belongs to a SERVICE collection "checkout"; the import takes a simple collection slug, pass --collection-slug`,
		},
		{
			name: "SERVICE collection reference with --collection-slug",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine,
				"collection: {type: SERVICE, slug: checkout}",
			),
			flags: importFlags{collectionSlug: "team-a"},
			want:  manifestRequest(`ClassicDashboard "xyz"`, "team-a", "", "xyz"),
		},
		{
			name: "USER collection reference without --collection-slug",
			content: manifestYAML("GrafanaDashboard", "slug: xyz", dashboardJSONLine,
				"collection: {type: USER, slug: someone}",
			),
			wantErr: `GrafanaDashboard "xyz" belongs to a USER collection "someone"; the import takes a simple collection slug, pass --collection-slug`,
		},
		{
			name: "bucket_slug only without --collection-slug",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine,
				"bucket_slug: legacy",
			),
			wantErr: `ClassicDashboard "xyz" belongs to bucket "legacy", and dashboards have no buckets; pass --collection-slug`,
		},
		{
			name: "bucket_slug only with --collection-slug",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine,
				"bucket_slug: legacy",
			),
			flags: importFlags{collectionSlug: "team-a"},
			want:  manifestRequest(`ClassicDashboard "xyz"`, "team-a", "", "xyz"),
		},
		{
			name:    "no collection",
			content: manifestYAML("ClassicDashboard", "slug: xyz", dashboardJSONLine),
			wantErr: `ClassicDashboard "xyz" has no collection; pass --collection-slug`,
		},
		{
			name:    "--collection-slug overrides spec.collection_slug",
			content: classicYAML,
			flags:   importFlags{collectionSlug: "migrated"},
			want:    manifestRequest(`ClassicDashboard "xyz"`, "migrated", "API Latency Overview", "xyz"),
		},
		{
			name:    "--name overrides spec.name",
			content: classicYAML,
			flags:   importFlags{name: "Latency v2"},
			want:    manifestRequest(`ClassicDashboard "xyz"`, "team-a", "Latency v2", "xyz"),
		},
		{
			name:    "--slug overrides spec.slug",
			content: classicYAML,
			flags:   importFlags{slug: "xyz-native"},
			want:    manifestRequest(`ClassicDashboard "xyz"`, "team-a", "API Latency Overview", "xyz-native"),
		},
		{
			name:    "empty spec.dashboard_json",
			content: manifestYAML("ClassicDashboard", "slug: xyz", "collection_slug: team-a"),
			wantErr: `ClassicDashboard "xyz" has no spec.dashboard_json`,
		},
		{
			name:    "manifest without spec",
			content: "api_version: v1/config\nkind: GrafanaDashboard\n",
			wantErr: "GrafanaDashboard has no spec.dashboard_json",
		},
		{
			name:    "Dashboard rejected by kind",
			content: manifestYAML("Dashboard", "slug: xyz", "collection_slug: team-a", dashboardJSONLine),
			wantErr: "unsupported kind v1/config/Dashboard; expected a ClassicDashboard or GrafanaDashboard manifest",
		},
		{
			name:    "update-if-exists with manifest without slug",
			content: manifestYAML("GrafanaDashboard", "collection_slug: team-a", dashboardJSONLine),
			flags:   importFlags{updateIfExists: true},
			wantErr: "GrafanaDashboard has no spec.slug; --slug is required when --update-if-exists is set",
		},
		{
			name:    "update-if-exists with manifest without slug and --slug",
			content: manifestYAML("GrafanaDashboard", "collection_slug: team-a", dashboardJSONLine),
			flags:   importFlags{slug: "latency", updateIfExists: true},
			want: &importRequest{
				source: "GrafanaDashboard",
				body: &models.Configv1ImportDashboardFromClassicRequest{
					ClassicDashboardJSON: testGrafanaJSON,
					CollectionSlug:       "team-a",
					Slug:                 "latency",
					UpdateIfExists:       true,
				},
			},
		},
		{
			name:    "unknown field rejected when strict",
			content: classicYAML + "  favorite: true\n",
			wantErr: "field favorite not found",
		},
		{
			name:    "unknown field allowed with permissive parsing",
			content: classicYAML + "  favorite: true\n",
			flags:   importFlags{permissiveParsing: true},
			want:    manifestRequest(`ClassicDashboard "xyz"`, "team-a", "API Latency Overview", "xyz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newImportRequest([]byte(tt.content), tt.flags)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
