build: `go build .`

run: (replace `ssm-backup-rules` with your prometheus rule group)
```
kubectl get prometheusRules -o yaml ssm-backup-rules | yq '.spec' | awk '{gsub(/severity: warning/, "severity: warn"); print}' > rules.yaml; ./converter rules.yaml | jq 'del(.created_at, .updated_at, .labels.receiver, .series_conditions.overrides) | .collection_slug = "infra" | {dry_run: false, monitor: . }' | curl --request POST \
  --url https://plaid.chronosphere.io/api/v1/config/monitors \
  --header "API-Token: ${CHRONOSPHERE_API_TOKEN}" \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data "@-"
```
