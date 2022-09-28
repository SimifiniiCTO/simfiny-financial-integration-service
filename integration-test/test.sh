#1 /usr/bin/env sh

set -e

# wait for service
# kubectl rollout status deployment/service --timeout=10m

# test service
helm test service
