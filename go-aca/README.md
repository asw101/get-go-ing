# go-aca

## docker

```bash
export HELLO='go-aca'

docker build -t go-aca .

docker run --rm \
    --env HELLO \
    -it go-aca
```

## azure container app

```bash
HELLO='go-aca'

az containerapp create \
  --resource-group "my-container-apps" \
  --environment "my-environment" \
  --name go-aca \
  --image ghcr.io/asw101/get-go-ing/go-aca:latest \
  --env-vars "HELLO=${HELLO}" \
  --min-replicas 1 \
  --max-replicas 2
```
