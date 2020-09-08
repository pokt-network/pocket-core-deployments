# This Dockerfile attempts to install dependencies, run the tests and build the pocket-core binary
# The result of this Dockerfile will put the pocket-core executable in the $GOBIN/bin, which in turn
# is part of the $PATH

# Dynamically pull Go-lang version for the image
ARG IMAGE_NAME=RC-0.5.1
# First build step to build the app binary
FROM poktnetwork/pocket-core:$IMAGE_NAME AS builder

# Second build step: reduce image size to only use app binary
FROM envoyproxy/envoy-alpine:v1.14.1

COPY --from=builder /bin/pocket /bin/pocket
COPY entrypoint.sh /tmp/
COPY account.sh /tmp/
COPY command.sh /tmp/
COPY config.json /tmp/
COPY envoy/envoy.yaml /etc/envoy/envoy.yaml
COPY envoy/certificate.sh ./
COPY envoy/entrypoint.sh ./
RUN apk add --update --no-cache expect bash leveldb-dev
# Create app user and add permissions
RUN addgroup --gid 1001 -S app \
	&& adduser --uid 1005 -S -G app app
RUN chown -R app /etc/envoy && chown -R app /bin/pocket  && mkdir -p /home/app/.pocket/config && chown -R app /home/app/.pocket
USER app
RUN cp /tmp/entrypoint.sh /home/app/ && cp /tmp/command.sh /home/app/command.sh && cp /tmp/config.json /home/app/config.json && cp /tmp/account.sh /home/app/account.sh && mkdir -p /home/app/.pocket/config


# Setup the WORKDIR with app user
ENTRYPOINT ["/entrypoint.sh"]
