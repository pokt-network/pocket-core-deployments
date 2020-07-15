# This Dockerfile attempts to install dependencies, run the tests and build the pocket-core binary
# The result of this Dockerfile will put the pocket-core executable in the $GOBIN/bin, which in turn
# is part of the $PATH

# Dynamically pull Go-lang version for the image
ARG GOLANG_IMAGE_VERSION=golang:1.13-alpine

# First build step to build the app binary
FROM ${GOLANG_IMAGE_VERSION} AS builder

# Install dependencies
RUN apk -v --update --no-cache add \
	git \
	groff \
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
ARG BRANCH_NAME="staging"

# Copying deps sh to tmp folder
COPY deps.sh /tmp/

# Clone the repository
RUN git clone --branch ${BRANCH_NAME} https://github.com/pokt-network/pocket-core.git ${POCKET_PATH}

# Install rest of source code

# Run tests
# As the tests were removed recently for the fact that most of them were broken, I commented this line,
# It should be uncommented as soon as the new tests are available.
#RUN go test ./tests/...

# Install project dependencies and builds the binary
RUN go build -tags cleveldb -o ${GOBIN}/bin/pocket ./app/cmd/pocket_core/main.go
COPY ./entrypoint.sh /entrypoint.sh
WORKDIR /
# Setup the WORKDIR with app user
ENTRYPOINT ["./entrypoint.sh"]
