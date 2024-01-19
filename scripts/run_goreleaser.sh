#!/usr/bin/env bash

set -euo pipefail
set -x
[[ -z ${DEBUG:-} ]] || set -o xtrace

DIR="$(cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")" && cd .. && pwd -P)"
GITCONFIG_VOLUME=${GIT_CONFIG:-"${HOME}/.gitconfig"}
GCLOUD_CONFIG_DIR="${CLOUDSDK_CONFIG:-${HOME}/.config/gcloud}"
SSH_CONFIG_DIR_VOLUME="${HOME}/.ssh"
SSH_AUTH_SOCK_ENV_VAR="/ssh-agent"
SSH_AUTH_SOCK_VOLUME="${SSH_AUTH_SOCK}:/ssh-agent"

if [[ "${BUILDKITE:-}" == "true" ]]; then
    GITCONFIG_VOLUME="/var/lib/buildkite-agent/.gitconfig"
    SSH_CONFIG_DIR_VOLUME="/var/lib/buildkite-agent/.ssh"
fi


DOCKER_OPTS=(
    -w "${GO_RELEASER_WORKING_DIR}"
    -e "GITHUB_TOKEN" # Set by CI
    -e "GO_BUILD_LDFLAGS"
    -e "CGO_ENABLED=0"
    -e "SSH_AUTH_SOCK=${SSH_AUTH_SOCK_ENV_VAR}"
    -v "${SSH_AUTH_SOCK_VOLUME}"
    -v "${SSH_CONFIG_DIR_VOLUME}:/root/.ssh"
    -v "${GITCONFIG_VOLUME}:/root/.gitconfig"
    -v "${DIR}:${GO_RELEASER_WORKING_DIR}"
)

if [[ "${BUILDKITE:-}" != "true" ]]; then
    DOCKER_OPTS+=(-v "${GCLOUD_CONFIG_DIR}:/root/.config/gcloud")
fi

# Run go mod vendor outside of container to utilize local cache
go mod vendor

# N.B. The GO_RELEASER_DOCKER_IMAGE is expected to be set by CI.
docker run "${DOCKER_OPTS[@]}" "${GO_RELEASER_DOCKER_IMAGE}" release "$@"
