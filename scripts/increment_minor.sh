#!/bin/bash

# This script is used to automate incrementing Git tags. If the last tag is "v0.n.0" this script
# will create a new tag "v0.(n+1).0" pointing to the latest commit. This script is derived, with
# minor modifications, from https://stackoverflow.com/questions/3760086/automatic-tagging-of-releases.

set -e

# As a precaution, we only allow tags to be created on clean branches to ensure that any changes
# the user wants to be included in the tag have been committed.
if [ -n "$(git status --porcelain)" ]; then
  echo "Can only add a tag from a clean branch, check git status and try again."
  exit 1
fi

VERSION=$(git describe --abbrev=0 --tags)

# Remove leading "v" from version tag (for details, see the "Substring Extraction" section of
# https://www.tldp.org/LDP/abs/html/string-manipulation.html).
VERSION="${VERSION:1}"

# Replace "." with a space so we can split the version into an array.
read -r -a VERSION_BITS <<<"${VERSION//./ }"

MAJOR_VERSION=${VERSION_BITS[0]}
MINOR_VERSION=${VERSION_BITS[1]}
MINOR_VERSION=$((MINOR_VERSION + 1))

NEW_VERSION="v${MAJOR_VERSION}.${MINOR_VERSION}.0"

# Get the current hash and see if it already has a tag. We will only create the new tag if the
# current hash does not already have a tag.
GIT_COMMIT=$(git rev-parse HEAD)
NEEDS_TAG=$(git describe --contains "$GIT_COMMIT" 2>/dev/null || echo "")

if [ -n "$NEEDS_TAG" ]; then
  echo "Can only add a new tag to a commit which does not have one, but the latest commit already does."
  exit 1
fi

echo "Updating $VERSION to $NEW_VERSION"
git tag $NEW_VERSION
echo "Pushing new tag $NEW_VERSION"
git push origin $NEW_VERSION
