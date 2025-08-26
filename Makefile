SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
REPO_DIR := $(abspath $(SELF_DIR))

tools_bin_path            := $(abspath ./_tools/bin)
repo_package              := github.com/chronosphereio/chronoctl-core

BUILD                     := $(abspath ./bin)
GO_BUILD_LDFLAGS_CMD      := $(abspath ./build/go-build-ldflags.sh)
GO_BUILD_LDFLAGS          := $(shell $(GO_BUILD_LDFLAGS_CMD))
GO_BUILD_COMMON_ENV       := CGO_ENABLED=0

GO_RELEASER_WORKING_DIR   := /go/src/github.com/chronosphere/chronoctl
GO_RELEASER_RELEASE_ARGS  ?= --clean

UNSTABLE_ENTITIES := link-templates,saved-trace-searches,dashboards,trace-tail-sampling-rules,services,metric-usages-by-metric-name,metric-usages-by-label-name

.PHONY: clean-build
clean-build:
	@test -n "$(BUILD)"
	@rm -rf $(BUILD)/*

.PHONY: clean
clean: clean-build
	@rm -f *.html *.xml *.out *.test
	@rm -rf tmp/*

.PHONY: install-tools
install-tools: go-version-check
	@echo "--- Installing tools"
	cd tools && GOBIN=$(tools_bin_path) go install \
		github.com/golang/mock/mockgen \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/go-swagger/go-swagger/cmd/swagger \
		./cmd/chronogen

.PHONY: gen
gen: install-tools
	@echo "--- Generating swagger files and mocks for chronoctl"
	rm -fr src/generated/swagger/configunstable/client
	rm -fr src/generated/swagger/configunstable/models
	rm -fr src/generated/swagger/configv1/client
	rm -fr src/generated/swagger/configv1/models
	rm -fr src/generated/swagger/stateunstable/client
	rm -fr src/generated/swagger/stateunstable/models
	rm -fr src/generated/swagger/statev1/client
	rm -fr src/generated/swagger/statev1/models
	rm -fr src/generated/cli
	@REPO_ROOT=$(realpath ./) UNSTABLE_ENTITIES=${UNSTABLE_ENTITIES} go generate ./src

.PHONY: gen-cli
gen-cli: install-tools
	${tools_bin_path}/chronogen -spec=src/generated/swagger/configv1/spec.json -target=src/generated/cli/configv1 -pkg=configv1
	${tools_bin_path}/chronogen -spec=src/generated/swagger/configunstable/spec.json -target=src/generated/cli/configunstable -pkg=configunstable -allowed-entities=${UNSTABLE_ENTITIES}

.PHONY: update-swagger
update-swagger:
	@[ -n "${SWAGGER_PATH}" ] || (echo "SWAGGER_PATH must be set, please set it and rerun this command"; exit 1)
	cp ${SWAGGER_PATH}/unstable_config_swagger.json src/generated/swagger/configunstable/spec.json
	cp ${SWAGGER_PATH}/v1_config_swagger.json src/generated/swagger/configv1/spec.json
	cp ${SWAGGER_PATH}/unstable_state_swagger.json src/generated/swagger/stateunstable/spec.json
	cp ${SWAGGER_PATH}/v1_state_swagger.json src/generated/swagger/statev1/spec.json
	$(MAKE) gen gen-cli

# Tests that all currently generated types match their contents if they were regenerated
.PHONY: test-gen
test-gen: gen gen-cli
	@echo "--- :golang: go mod tidy"
	go mod tidy
	@echo "--- Testing generate code is up to date"
	@[ -z "$$(git status --porcelain)" ] || ((set -x; git status --porcelain; git diff); echo -e "^^^ +++\nCheck git status + diff above, there are changed or untracked files"; exit 1)

.PHONY: test
test:
	go test ./src/...

.PHONY: lint
lint: install-tools
	@echo "--- :golang: linting code"
	$(tools_bin_path)/golangci-lint run -D gosec

.PHONY: chronoctl
chronoctl:
	$(GO_BUILD_COMMON_ENV) go build -ldflags '$(GO_BUILD_LDFLAGS)'  -o $(BUILD)/chronoctl ./src/cmd/chronoctl/.

.PHONY: go-version-check
go-version-check:
	# make sure you're running the right version of Go, otherwise builds/codegen/tests
	# may have inconsistent results that are hard to debug.
	go version | grep go1.24 || (echo "Error: you must be running go1.24.x" && exit 1)

.PHONY: release
release:
	@echo "Releasing new version"
	GO_BUILD_LDFLAGS="$(GO_BUILD_LDFLAGS)" \
		GO_RELEASER_DOCKER_IMAGE=$(GO_RELEASER_DOCKER_IMAGE) \
		GO_RELEASER_RELEASE_ARGS="$(GO_RELEASER_RELEASE_ARGS)" \
		GO_RELEASER_WORKING_DIR=$(GO_RELEASER_WORKING_DIR) \
		SSH_AUTH_SOCK=$(SSH_AUTH_SOCK) \
		./scripts/run_goreleaser.sh ${GO_RELEASER_RELEASE_ARGS}

.PHONY: release-snapshot
release-snapshot:
	@echo Building binaries with goreleaser
	# --snapshot mode allows building artifacts w/o release tag present and w/ publishing mode disabled useful when we want to test whether we can build binaries, but not publish yet.
	make release GO_RELEASER_RELEASE_ARGS="--snapshot --clean --skip=publish"
