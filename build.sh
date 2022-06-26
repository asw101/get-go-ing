TAG="ghcr.io/asw101/get-go-ing/aca-go:latest"
SOURCE="aca-go/"

echo "DOCKER_LABELS: ${DOCKER_LABELS}"

docker build \
    -t "$TAG" \
    --label "$DOCKER_LABELS" \
    "$SOURCE"

docker push "$TAG"
