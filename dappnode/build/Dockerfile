# This Dockerfile attempts to install dependencies, run the tests and build the pocket-core binary
# The result of this Dockerfile will put the pocket-core executable in the $GOBIN/bin, which in turn
# is part of the $PATH

# Dynamically pull Go-lang version for the image
ARG GOLANG_IMAGE_VERSION=golang:1.13-alpine

# First build step to build the app binary
FROM ${GOLANG_IMAGE_VERSION} AS builder

# Install dependencies
RUN apk -v --update --no-cache add \
	curl \
	git \
	groff \
	less \
	mailcap \
	gcc \
	libc-dev \
	bash  \
	leveldb-dev && \
	rm /var/cache/apk/* || true

# Environment and system dependencies setup
ENV POCKET_PATH=/go/src/github.com/pokt-network/pocket-core/
ENV GO111MODULE="on"

# Create node root directory
RUN mkdir -p ${POCKET_PATH}
WORKDIR $POCKET_PATH

# Creating the BRANCH_NAME variable
ARG BRANCH_NAME="RC-0.4.1"

# Copying deps sh to tmp folder
COPY deps.sh /tmp/

# Clone the repository
RUN git clone --branch ${BRANCH_NAME} https://github.com/pokt-network/pocket-core.git ${POCKET_PATH}

# Install rest of source code
COPY . .

# Run tests
# As the tests were removed recently for the fact that most of them were broken, I commented this line,
# It should be uncommented as soon as the new tests are available.
#RUN go test ./tests/...

# Install project dependencies and builds the binary
RUN go build -tags cleveldb -o ${GOBIN}/bin/pocket ./app/cmd/pocket_core/main.go

# Second build step: reduce image size to only use app binary
FROM envoyproxy/envoy-alpine:v1.14.1

COPY --from=builder /bin/pocket /bin/pocket
COPY entrypoint.sh /tmp/
COPY command.sh /tmp/
COPY envoy/envoy.yaml /etc/envoy/envoy.yaml
COPY envoy/certificate.sh ./
COPY envoy/entrypoint.sh ./
RUN apk add --update --no-cache expect bash leveldb-dev openssl
# Create app user and add permissions
RUN addgroup --gid 1001 -S app \
	&& adduser --uid 1005 -S -G app app
RUN mv /tmp/entrypoint.sh /home/app/ && chown -R app /etc/envoy && mv /tmp/command.sh /home/app/command.sh && chown -R app /bin/pocket  && mkdir -p /home/app/.pocket/config && chown -R app /home/app/.pocket
USER app

# Setup the WORKDIR with app user
ENTRYPOINT ["/entrypoint.sh"]