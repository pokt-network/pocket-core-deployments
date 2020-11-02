# Create docker dependencies
docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions


# Grant file permissions
chown 104 -R ./monitoring/grafana
chown 1000 -R ./monitoring/prometheus
