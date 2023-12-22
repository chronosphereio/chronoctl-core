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

// Package clitemplates contains the templates used to generate Chronoctl CLI code.
package clitemplates

import _ "embed"

// Entity is the template for a specific entity
//
//go:embed entity.go.tmpl
var Entity string

// Create is the template for a specific entity's create functionality
//
//go:embed create.go.tmpl
var Create string

// Read is the template for a specific entity's read functionality
//
//go:embed read.go.tmpl
var Read string

// Update  is the template for a specific entity's update functionality
//
//go:embed update.go.tmpl
var Update string

// Delete is the template for a specific entity's delete functionality
//
//go:embed delete.go.tmpl
var Delete string

// List is the template for a specific entity's list functionality
//
//go:embed list.go.tmpl
var List string

// Scaffold contains the template for the empty scaffold object.
//
//go:embed scaffold.go.tmpl
var Scaffold string

// CLI is the template for generating CLI code that uses the
// generated entity commands.
//
//go:embed cli.go.tmpl
var CLI string
