#!/usr/bin/env bash
set -euo pipefail

# kubectl patch prom -n prombench bench --type='json' -p='[{"op": "replace", "path": "/spec/shards", "value": 2}]'

# kubectl patch deploy -n prombench generate-timeseries --type='json' \
#   -p='[{"op": "replace", "path": "/spec/template/spec/containers/0/args/0", "value": "--timeseries=100000"}]'
# kubectl scale deploy -n prombench generate-timeseries --replicas 3

# sort_desc(irate(container_cpu_usage_seconds_total{namespace="prombench"}[5m]) / ignoring(cpu) (container_spec_cpu_quota / container_spec_cpu_period))
# sort_desc(container_memory_working_set_bytes{namespace="prombench"} / (container_spec_memory_limit_bytes > 0))
