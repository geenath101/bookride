#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

"${SCRIPT_DIR}/api-gateway-build.sh"
"${SCRIPT_DIR}/trip-build.sh"

if [[ "${SKIP_WEB_BUILD:-0}" != "1" ]]; then
	"${SCRIPT_DIR}/web-build.sh"
fi

echo "Build flow finished."
