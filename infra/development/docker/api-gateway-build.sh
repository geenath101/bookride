#!/usr/bin/env bash
set -euo pipefail

echo $BASH_SOURCE[0]
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

REPO_ROOT="$(cd "${SCRIPT_DIR}/../../.." && pwd)"

cd "$REPO_ROOT"
mkdir -p build

# Build the whole package (not just main.go) so all gateway files are included.
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./services/api-gateway

echo "Built build/api-gateway"
