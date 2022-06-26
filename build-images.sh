# echo "DOCKER_LABELS: ${DOCKER_LABELS}"


for d in $(ls -d */ | sed 's#/##'); do
    echo "$d"
done


# for d in */ ; do
#     TAG="ghcr.io/${GITHUB_ACTION_REPOSITORY}/${d}:latest"
    
#     docker build \
#         -t "$TAG" \
#         --label "$DOCKER_LABELS" \
#         "$d/"

#     docker push "$TAG"
# done
