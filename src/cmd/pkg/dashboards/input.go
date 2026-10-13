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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/chronosphereio/chronoctl-core/src/generated/cli/configv1"
	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/chronosphereio/chronoctl-core/src/thirdparty/yaml"
	"github.com/chronosphereio/chronoctl-core/src/types"
)

var (
	jsonStart      = regexp.MustCompile(`^\s*\{`)
	jsonArrayStart = regexp.MustCompile(`^\s*\[`)
)

// importFlags are the command-line values that feed an import request.
type importFlags struct {
	// filename only labels raw JSON input in the status line.
	filename          string
	collectionSlug    string
	name              string
	slug              string
	dryRun            bool
	updateIfExists    bool
	permissiveParsing bool
}

// importRequest is one ImportDashboardFromClassic call built from the input.
type importRequest struct {
	// source names the input in the status line, e.g. `ClassicDashboard "xyz"`
	// or `-f grafana.json`.
	source string
	body   *models.Configv1ImportDashboardFromClassicRequest
}

// newImportRequest builds the request from -f content, which is either raw
// Grafana dashboard JSON or exactly one ClassicDashboard or GrafanaDashboard
// manifest. All input errors surface here, before any API call.
func newImportRequest(content []byte, f importFlags) (*importRequest, error) {
	if len(bytes.TrimSpace(content)) == 0 {
		return nil, errors.New("input is empty")
	}

	// A raw Grafana dashboard is always a JSON object, so an array can only be
	// manifests.
	isManifest := jsonArrayStart.Match(content)
	if !isManifest {
		meta, err := peekTypeMeta(content)
		if err != nil {
			return nil, err
		}
		isManifest = meta.APIVersion != ""
	}

	var (
		req *importRequest
		err error
	)
	if isManifest {
		req, err = requestFromManifest(content, f)
	} else {
		req, err = requestFromRawJSON(content, f)
	}
	if err != nil {
		return nil, err
	}

	req.body.DryRun = f.dryRun
	req.body.UpdateIfExists = f.updateIfExists
	return req, nil
}

// peekTypeMeta decodes api_version and kind from the first document only,
// choosing the decoder the same way types.Decode does. Stopping at the first
// document sends JSON lines and YAML streams down the manifest path, where
// they fail as multi-document input. JSON that does not decode into a
// TypeMeta is still candidate raw dashboard JSON, so only YAML parse failures
// are errors.
func peekTypeMeta(content []byte) (types.TypeMeta, error) {
	var meta types.TypeMeta
	if jsonStart.Match(content) {
		_ = json.NewDecoder(bytes.NewReader(content)).Decode(&meta)
		return meta, nil
	}
	if err := yaml.NewDecoder(bytes.NewReader(content)).Decode(&meta); err != nil && !errors.Is(err, io.EOF) {
		return types.TypeMeta{}, fmt.Errorf("expected raw Grafana dashboard JSON or a ClassicDashboard or GrafanaDashboard manifest: could not parse input as YAML: %w", err)
	}
	return meta, nil
}

func requestFromRawJSON(content []byte, f importFlags) (*importRequest, error) {
	if !jsonStart.Match(content) || !json.Valid(content) {
		return nil, errors.New("expected raw Grafana dashboard JSON or a ClassicDashboard or GrafanaDashboard manifest with api_version and kind")
	}
	if f.collectionSlug == "" {
		return nil, errors.New("--collection-slug is required")
	}
	if f.updateIfExists && f.slug == "" {
		return nil, errors.New("--slug is required when --update-if-exists is set")
	}

	// The JSON is sent byte-for-byte: the API owns parsing and converting it.
	return &importRequest{
		source: "-f " + f.filename,
		body: &models.Configv1ImportDashboardFromClassicRequest{
			ClassicDashboardJSON: string(content),
			CollectionSlug:       f.collectionSlug,
			Name:                 f.name,
			Slug:                 f.slug,
		},
	}, nil
}

// requestFromManifest decodes the manifest the way `dashboards create -f`
// does and maps its spec onto the request. A flag that is set always wins,
// and the manifest field it replaces is not consulted at all.
func requestFromManifest(content []byte, f importFlags) (*importRequest, error) {
	objs, err := types.Decode(bytes.NewReader(content), f.permissiveParsing)
	if err != nil {
		return nil, err
	}
	if len(objs) == 0 {
		return nil, errors.New("input does not contain any items")
	}
	if len(objs) > 1 {
		return nil, errors.New(`input contains more than one item; import one dashboard per run, e.g. pipe one "classic-dashboards read <slug>" at a time rather than "list"`)
	}

	// ClassicDashboard and GrafanaDashboard are two kinds over the same spec.
	var spec *models.Configv1GrafanaDashboard
	switch obj := objs[0].(type) {
	case *configv1.ClassicDashboard:
		spec = obj.Spec
	case *configv1.GrafanaDashboard:
		spec = obj.Spec
	default:
		return nil, fmt.Errorf("unsupported kind %s; expected a ClassicDashboard or GrafanaDashboard manifest", obj.Type())
	}
	if spec == nil {
		spec = &models.Configv1GrafanaDashboard{}
	}

	label := objs[0].Type().Kind
	if spec.Slug != "" {
		label += fmt.Sprintf(" %q", spec.Slug)
	}

	if spec.DashboardJSON == "" {
		return nil, fmt.Errorf("%s has no spec.dashboard_json", label)
	}

	collectionSlug, err := manifestCollectionSlug(label, spec, f.collectionSlug)
	if err != nil {
		return nil, err
	}

	name := f.name
	if name == "" {
		name = spec.Name
	}
	slug := f.slug
	if slug == "" {
		slug = spec.Slug
	}
	if f.updateIfExists && slug == "" {
		return nil, fmt.Errorf("%s has no spec.slug; --slug is required when --update-if-exists is set", label)
	}

	return &importRequest{
		source: label,
		body: &models.Configv1ImportDashboardFromClassicRequest{
			ClassicDashboardJSON: spec.DashboardJSON,
			CollectionSlug:       collectionSlug,
			Name:                 name,
			Slug:                 slug,
		},
	}, nil
}

// manifestCollectionSlug resolves the request's collection_slug. The API takes
// a simple collection slug only, and dashboards have no buckets, so any other
// placement needs an explicit --collection-slug.
func manifestCollectionSlug(label string, spec *models.Configv1GrafanaDashboard, flag string) (string, error) {
	if flag != "" {
		return flag, nil
	}
	if spec.CollectionSlug != "" {
		return spec.CollectionSlug, nil
	}
	if c := spec.Collection; c != nil && c.Slug != "" {
		if c.Type == "" || c.Type == models.Configv1CollectionReferenceTypeSIMPLE {
			return c.Slug, nil
		}
		return "", fmt.Errorf("%s belongs to a %s collection %q; the import takes a simple collection slug, pass --collection-slug", label, c.Type, c.Slug)
	}
	if spec.BucketSlug != "" {
		return "", fmt.Errorf("%s belongs to bucket %q, and dashboards have no buckets; pass --collection-slug", label, spec.BucketSlug)
	}
	return "", fmt.Errorf("%s has no collection; pass --collection-slug", label)
}
