# ☸️ gotway 

[![CI](https://github.com/gotway/gotway/actions/workflows/ci.yml/badge.svg)](https://github.com/gotway/gotway/actions/workflows/ci.yml)
[![Release](https://github.com/gotway/gotway/actions/workflows/release.yml/badge.svg)](https://github.com/gotway/gotway/actions/workflows/release.yml)
[![Deploy](https://github.com/gotway/gotway/actions/workflows/deploy.yml/badge.svg)](https://github.com/gotway/gotway/actions/workflows/deploy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gotway/gotway)](https://goreportcard.com/report/github.com/gotway/gotway)
[![Go Reference](https://pkg.go.dev/badge/github.com/gotway/gotway.svg)](https://pkg.go.dev/github.com/gotway/gotway)
[![Artifact HUB](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/gotway)](https://artifacthub.io/packages/search?repo=gotway)

Cloud native API Gateway powered with in-redis cache.

- API composition: expose your services to the internet using a single endpoint
- Cloud native: configure routing and cache using [Kubernetes CRDs](./manifests/examples/catalog.yml)
- In-memory cache using redis 
- Cache invalidation using tags
- Health checking
- Management [REST API](#management-rest-api)
- ~6MB [Docker image](https://hub.docker.com/r/gotwaygateway/gotway/tags) available for multiple architectures
- [Helm chart](https://github.com/gotway/charts)

### Installation

```bash
helm repo add mmontes https://charts.mmontes-dev.duckdns.org
helm repo add gotway https://charts.gotway.duckdns.org
helm install redis mmontes/redis
helm install gotway gotway/gotway
```

### Quickstart

We will need microservices to route to in order to test gotway, you can deploy some by running:

```bash
helm upgrade --install gotway gotway/gotway --set examples.enabled=true
```

Let's register the [catalog](https://github.com/gotway/service-examples/tree/main/cmd/catalog) service into Gotway by creating an `IngressHTTP` CRD:

```bash
kubectl apply -f ./manifests/examples/catalog.yml 
``` 
```yaml
apiVersion: gotway.io/v1alpha1
kind: IngressHTTP
metadata:
  name: catalog
spec:
  match:
    host: catalog.gotway.duckdns.org:9111
  service:
    name: catalog
    url: http://gotway-catalog
    healthPath: /health
  cache:
    ttl: 30
    statuses:
      - 200
      - 404
    tags:
      - "catalog"
      - "products"

```

We are now able to route requests through Gotway:

```bash
curl -k --request GET 'https://catalog.gotway.duckdns.org:9111/products' | jq
```
```json
{
    "products": [
        {
            "id": 911902081,
            "name": "sneakers",
            "price": 69000,
            "color": "white",
            "size": "42"
        }
    ],
    "totalCount": 1
}
``` 

This response has a TTL of 30 seconds, let's invalidate the cache for the catalog service by providing one of its tags:

```bash
curl -k --request DELETE 'https://gotway.duckdns.org:9111/api/cache' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tags": ["catalog"]
}'
``` 
### Management REST API 

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/9776-3e976745-8b33-46c1-bfe6-d7211722d809?action=collection%2Ffork&collection-url=entityId%3D9776-3e976745-8b33-46c1-bfe6-d7211722d809%26entityType%3Dcollection%26workspaceId%3D10c73242-ad78-405e-b364-b37e56fbb5d3#?env%5BGotway%5D=W3sia2V5IjoidXJsIiwidmFsdWUiOiJodHRwczovL2dvdHdheS5kdWNrZG5zLm9yZzo5MTExIiwiZW5hYmxlZCI6dHJ1ZSwic2Vzc2lvblZhbHVlIjoiaHR0cHM6Ly9nb3R3YXkuZHVja2Rucy5vcmc6OTExMSIsInNlc3Npb25JbmRleCI6MH0seyJrZXkiOiJ1cmxDYXRhbG9nIiwidmFsdWUiOiJodHRwczovL2NhdGFsb2cuZ290d2F5LmR1Y2tkbnMub3JnOjkxMTEiLCJlbmFibGVkIjp0cnVlLCJzZXNzaW9uVmFsdWUiOiJodHRwczovL2NhdGFsb2cuZ290d2F5LmR1Y2tkbnMub3JnOjkxMTEiLCJzZXNzaW9uSW5kZXgiOjF9LHsia2V5IjoidXJsU3RvY2siLCJ2YWx1ZSI6Imh0dHBzOi8vc3RvY2suZ290d2F5LmR1Y2tkbnMub3JnOjQ0MzMiLCJlbmFibGVkIjp0cnVlLCJzZXNzaW9uVmFsdWUiOiJodHRwczovL3N0b2NrLmdvdHdheS5kdWNrZG5zLm9yZzo5MTExIiwic2Vzc2lvbkluZGV4IjoyfSx7ImtleSI6InByb2R1Y3RJZCIsInZhbHVlIjoiMTIzNCIsImVuYWJsZWQiOnRydWUsInNlc3Npb25WYWx1ZSI6IjEyMzQiLCJzZXNzaW9uSW5kZXgiOjN9LHsia2V5IjoicHJvZHVjdElkMiIsInZhbHVlIjoiNDU2IiwiZW5hYmxlZCI6dHJ1ZSwic2Vzc2lvblZhbHVlIjoiNDU2Iiwic2Vzc2lvbkluZGV4Ijo0fSx7ImtleSI6InByb2R1Y3RJZDMiLCJ2YWx1ZSI6Ijc4OSIsImVuYWJsZWQiOnRydWUsInNlc3Npb25WYWx1ZSI6Ijc4OSIsInNlc3Npb25JbmRleCI6NX1d)
