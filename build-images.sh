# echo "DOCKER_LABELS: ${DOCKER_LABELS}"


for d in $(ls -d */ | sed 's#/##'); do
    
    TAG="ghcr.io/${GITHUB_ACTION_REPOSITORY}/${d}:${GIT_BRANCH}"
    if [ "$GIT_BRANCH" == "main" ]; then
        TAG="ghcr.io/${GITHUB_ACTION_REPOSITORY}/${d}:latest"
    fi

    docker build \
        -t "$TAG" \
        --label "$DOCKER_LABELS" \
        "$d/"

    docker push "$TAG"
done
