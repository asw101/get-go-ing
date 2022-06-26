# aca-go

## docker

```bash
export HELLO='aca-go'

docker build -t aca-go .

docker run --rm \
    --env HELLO \
    -it aca-go
```

## azure container app

```bash
HELLO='aca-go'

az containerapp create \
  --resource-group "my-container-apps" \
  --environment "my-environment" \
  --name aca-go \
  --image ghcr.io/asw101/aca-go:latest \
  --env-vars "HELLO=${HELLO}" \
  --min-replicas 1 \
  --max-replicas 2
```
