#!/bin/bash
set -e

FAIL=0

# R1: core must not have non-stdlib deps (except golang.org/x/sync)
echo "==> R1: checking core dependencies..."
CORE_DEPS=$(grep "^require" -A100 core/go.mod | grep -v "golang.org/x/sync" | grep -v "^require" | grep -v "^)" | grep -v "^$" | grep -E "^\s+\S" || true)
if [ -n "$CORE_DEPS" ]; then
    echo "FAIL R1: core has forbidden dependencies:"
    echo "$CORE_DEPS"
    FAIL=1
fi

# R2: component layers (log/db/mq/cache) must only depend on core
for mod in log db mq cache; do
    echo "==> R2: checking $mod dependencies..."
    DEPS=$(grep -E "require|^\s+github.com" ${mod}/go.mod | grep -v "golang.org/x" | grep -v "go-nano/core" | grep -v "^require" | grep -v "^)" | grep -v "^$" | grep -E "^\s+\S" || true)
    if [ -n "$DEPS" ]; then
        echo "FAIL R2: $mod has forbidden dependencies:"
        echo "$DEPS"
        FAIL=1
    fi
done

echo "==> Dependency check complete"
exit $FAIL
