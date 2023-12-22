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
	"testing"

	"github.com/chronosphereio/chronoctl-core/src/generated/swagger/configv1/models"
	"github.com/stretchr/testify/require"
)

func TestConfigv1ChangelogMessage(t *testing.T) {
	tests := []struct {
		name     string
		change   *models.ResourceChange
		expected string
	}{
		{
			name: "resource with underscores",
			change: &models.ResourceChange{
				Resource: models.Configv1ResourceTypeNOTIFICATIONPOLICY,
				Action:   models.ResourceChangeActionDELETED,
				Slug:     "snail",
			},
			expected: `NotificationPolicy{slug="snail"} deleted`,
		},
		{
			name: "one word resource",
			change: &models.ResourceChange{
				Resource: models.Configv1ResourceTypeBUCKET,
				Action:   models.ResourceChangeActionDELETED,
				Slug:     "snail2",
			},
			expected: `Bucket{slug="snail2"} deleted`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, Configv1ChangelogMessage(tt.change), tt.expected)
		})
	}
}
