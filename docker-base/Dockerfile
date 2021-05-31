# This Dockerfile attempts to install dependencies, run the tests and build the pocket-core binary
# The result of this Dockerfile will put the pocket-core executable in the $GOBIN/bin, which in turn
# is part of the $PATH

# Dynamically pull Go-lang version for the image
ARG DOCKER_TAG

# First build step to build the app binary
FROM poktnetwork/pocket-core:$DOCKER_TAG AS builder

# Second build step: reduce image size to only use app binary
FROM alpine:3.13

COPY --from=builder /bin/pocket /bin/pocket
COPY entrypoint.sh entrypoint.sh
RUN apk add --update --no-cache expect leveldb-dev
# Create app user and add permissions
RUN apk add --no-cache tzdata && cp /usr/share/zoneinfo/America/New_York  /etc/localtime

# Setup the WORKDIR with app user
ENTRYPOINT ["pocket", "start"]
