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

package output

import (
	"fmt"
	"strings"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/chronosphereio/chronoctl-core/src/types"
	"github.com/go-openapi/inflect"
)

// Configv1ChangelogMessage creates a user readable changelog message.
func Configv1ChangelogMessage(change *models.ResourceChange) string {
	description := types.Description(inflect.Camelize(strings.ToLower(string(change.Resource))), "slug", change.Slug)
	return fmt.Sprintf("%v %v", description, strings.ToLower(string(change.Action)))
}
