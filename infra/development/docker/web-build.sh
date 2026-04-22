#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../../.." && pwd)"

cd "$REPO_ROOT/web"

if [[ -f package-lock.json ]]; then
  npm ci
else
  npm install
fi

npm run build

echo "Built web app in web/.next"
