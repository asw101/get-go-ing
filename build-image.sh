#!/bin/bash
set -eou pipefail

[[ -z "${GITHUB_REPOSITORY:-}" ]] && echo "GITHUB_REPOSITORY (e.g. asw101/get-go-ing) not set" && exit 1
[[ -z "${GITHUB_BRANCH:-}" ]] && echo "GITHUB_BRANCH (e.g. main) not set" && exit 1
[[ -z "${DOCKER_LABELS:-}" ]] && DOCKER_LABELS='org.opencontainers.image.description=local build'

[[ -z "${DIR:-}" ]] && echo "DIR (e.g. app) not set" && exit 1

# remove any trailing slash, if supplied
DIR=$(echo $DIR | sed 's#/##')

TAG="ghcr.io/${GITHUB_REPOSITORY}/${DIR}:${GITHUB_BRANCH}"
if [ "$GITHUB_BRANCH" == "main" ]; then
    TAG="ghcr.io/${GITHUB_REPOSITORY}/${DIR}:latest"
fi

docker build \
    -t "$TAG" \
    --label "$DOCKER_LABELS" \
    "$DIR"

docker push "$TAG"
