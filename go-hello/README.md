# go-hello

go-hello is a simple http web server with multiple endpoints for testing purposes.

```
/
/echo
/env
/host
/healthz
/wait?ms=400
```

It is built using a [Distroless](https://github.com/GoogleContainerTools/distroless#distroless-container-images) container image, via GitHub Actions, and made available publicly via GitHub Container Registry at `ghcr.io/asw101/get-go-ing/go-hello:latest`.

Kubernetes manifests for deployment to Kubernetes are available under [deploy/](./deploy/).
