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

// Package src is the root of the chronoctl source code
package src

//go:generate ${REPO_ROOT}/_tools/bin/swagger generate client --spec=./generated/swagger/configunstable/spec.json --target=./generated/swagger/configunstable --keep-spec-order --skip-tag-packages
//go:generate ${REPO_ROOT}/_tools/bin/swagger generate client --spec=./generated/swagger/configv1/spec.json --target=./generated/swagger/configv1 --keep-spec-order --skip-tag-packages
//go:generate ${REPO_ROOT}/_tools/bin/swagger generate client --spec=./generated/swagger/stateunstable/spec.json --target=./generated/swagger/stateunstable --keep-spec-order --skip-tag-packages
//go:generate ${REPO_ROOT}/_tools/bin/swagger generate client --spec=./generated/swagger/statev1/spec.json --target=./generated/swagger/statev1 --keep-spec-order --skip-tag-packages

// Generate mocks for the above swagger clients
//go:generate ${REPO_ROOT}/_tools/bin/mockgen -source=./generated/swagger/configunstable/client/operations/operations_client.go -destination=./generated/swagger/configunstable/mocks/client_mock.go -package=mocks
//go:generate ${REPO_ROOT}/_tools/bin/mockgen -source=./generated/swagger/configv1/client/operations/operations_client.go -destination=./generated/swagger/configv1/mocks/client_mock.go -package=mocks
//go:generate ${REPO_ROOT}/_tools/bin/mockgen -source=./generated/swagger/stateunstable/client/operations/operations_client.go -destination=./generated/swagger/stateunstable/mocks/client_mock.go -package=mocks
//go:generate ${REPO_ROOT}/_tools/bin/mockgen -source=./generated/swagger/statev1/client/operations/operations_client.go -destination=./generated/swagger/statev1/mocks/client_mock.go -package=mocks

// Generate cli
//go:generate ${REPO_ROOT}/_tools/bin/chronogen -spec=./generated/swagger/configv1/spec.json -target=./generated/cli/configv1 -pkg=configv1
//go:generate ${REPO_ROOT}/_tools/bin/chronogen -spec=./generated/swagger/configunstable/spec.json -target=./generated/cli/configunstable -pkg=configunstable -allowed-entities=${UNSTABLE_ENTITIES}
