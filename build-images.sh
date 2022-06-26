# echo "DOCKER_LABELS: ${DOCKER_LABELS}"

for DIR in */ ; do
    TAG="ghcr.io/${GITHUB_ACTION_REPOSITORY}/${DIR}:latest"
    
    docker build \
        -t "$TAG" \
        --label "$DOCKER_LABELS" \
        "$DIR/"

    docker push "$TAG"
done
