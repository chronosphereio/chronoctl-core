#!/bin/sh

export GIT_REVISION=$(git rev-parse --short HEAD)
export GIT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo unknown)
export BUILD_DATE=$(date '+%F-%T') # outputs something in this format 2017-08-21-18:58:45
export BASE_PACKAGE=github.com/chronosphereio/chronoctl-core/src/cmd/pkg/buildinfo

LD_FLAGS="-X ${BASE_PACKAGE}.SHA=${GIT_REVISION} \
-X ${BASE_PACKAGE}.Version=${GIT_VERSION}             \
-X ${BASE_PACKAGE}.Date=${BUILD_DATE}"

echo $LD_FLAGS
