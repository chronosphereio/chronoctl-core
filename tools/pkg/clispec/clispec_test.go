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

package clispec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {
	tests := []struct {
		goType       string
		expectedFlag string
	}{
		{
			goType:       "[]string",
			expectedFlag: "StringSliceVar",
		},
		{
			goType:       "string",
			expectedFlag: "StringVar",
		},
		{
			goType:       "[]bool",
			expectedFlag: "BoolSliceVar",
		},
		{
			goType:       "bool",
			expectedFlag: "BoolVar",
		},
		{
			goType:       "int",
			expectedFlag: "IntVar",
		},
		{
			goType:       "float64",
			expectedFlag: "Float64Var",
		},
		{
			goType:       "SomeUnkownStruct",
			expectedFlag: "SomeUnkownStructVar",
		},
	}

	for _, tt := range tests {
		param := &Parameter{
			GoType: tt.goType,
		}
		assert.Equal(t, tt.expectedFlag, param.FlagType())
	}
}

func TestSwaggerName(t *testing.T) {
	tests := []struct {
		paramName string
		want      string
	}{
		{
			paramName: "user-json",
			want:      "UserJSON",
		},
		{
			paramName: "user_json",
			want:      "UserJSON",
		},
		{
			paramName: "pid",
			want:      "Pid",
		},
		{
			paramName: "pid-id",
			want:      "PidID", // only suffix should be replaced.
		},
		{
			paramName: "id-pid-type",
			want:      "IDPidType", // only prefix should be replaced.
		},
		{
			paramName: "dashboard_json",
			want:      "DashboardJSON",
		},
		{
			paramName: "notification_policy-id",
			want:      "NotificationPolicyID",
		},
		{
			paramName: "UiComponent",
			want:      "UIComponent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.paramName, func(t *testing.T) {
			p := Parameter{Name: tt.paramName}
			assert.Equal(t, tt.want, p.SwaggerName())
		})
	}
}
