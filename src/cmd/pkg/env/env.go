// Copyright 2024 Chronosphere Inc.
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

// Package env provides environment variable definitions for chronoctl.
package env

const (
	// ChronosphereOrgNameKey is the environment variable that specifies the Chronosphere customer organization
	ChronosphereOrgNameKey = "CHRONOSPHERE_ORG_NAME"
	// ChronosphereOrgKey is the environment variable that specifies the Chronosphere customer organization
	ChronosphereOrgKey = "CHRONOSPHERE_ORG" // fallback for CHRONOSPHERE_ORG_NAME
	// ChronosphereAPITokenKey is the environment variable that specifies the Chronosphere API token
	ChronosphereAPITokenKey = "CHRONOSPHERE_API_TOKEN"
	// ChronosphereEntityNamespace is the environment variable that specifies the Chronosphere entity namespace
	ChronosphereEntityNamespace = "CHRONOSPHERE_ENTITY_NAMESPACE"
)
