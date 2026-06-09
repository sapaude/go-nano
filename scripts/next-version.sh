#!/bin/bash
MODULE=${1:?usage: next-version.sh <module-dir>}
PREFIX="${MODULE}/v"
LATEST=$(git tag --list "${PREFIX}*" | sort -V | tail -1)
if [ -z "$LATEST" ]; then
    echo "1.0.0"
    exit 0
fi
VERSION=${LATEST#${PREFIX}}
MAJOR=$(echo "$VERSION" | cut -d. -f1)
MINOR=$(echo "$VERSION" | cut -d. -f2)
PATCH=$(echo "$VERSION" | cut -d. -f3)
echo "${MAJOR}.${MINOR}.$((PATCH + 1))"
