# Build
docker build \
  --build-arg http_proxy=$http_proxy \
  --build-arg https_proxy=$https_proxy \
  --build-arg no_proxy="$no_proxy" \
  --build-arg HTTP_PROXY=$HTTP_PROXY \
  --build-arg HTTPS_PROXY=$HTTPS_PROXY \
  --build-arg NO_PROXY="$NO_PROXY" \
  -t golang-standard:1.25-alpine \
  -f docker/alpine/Dockerfile .