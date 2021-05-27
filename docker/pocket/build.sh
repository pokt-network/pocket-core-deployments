#!/bin/bash -eo pipefail

# Set error conditions
# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Variable declaration
GOLANG_VERSION="$1"
BRANCH_NAME="$2"
DOCKER_IMAGE_NAME="$3"
DOCKER_TAG="$4"
STAGING_BRANCH="staging"
MASTER_BRANCH="master"

# Parse parameters
if [ ! -n "$GOLANG_VERSION" ]
then
    GOLANG_VERSION="1.13"
fi

if [ ! -n "$BRANCH_NAME" ]
then
    BRANCH_NAME="staging"
fi

if [ ! -n "$DOCKER_IMAGE_NAME" ]
then
    DOCKER_IMAGE_NAME="poktnetwork/pocket-core"
fi

if [ ! -n "$DOCKER_TAG" ]
then
    # Resolve DOCKER_TAG using the branch name
    if [ "$BRANCH_NAME" = "$STAGING_BRANCH" ]
    then
        # Handle staging branch
        echo "It's devnet!"
        DOCKER_TAG="devnet-latest"
    elif echo "$BRANCH_NAME" | grep -q "RC-"*
    then
        # Handle master branch
        echo "It's stagenet!"
        DOCKER_TAG="stagenet-latest"
    elif echo "$BRANCH_NAME" | grep -q "Beta-"*
    then
        # Handle master branch
        echo "It's beta!"
        DOCKER_TAG="beta-latest"
    else
        # It's a tag!
        echo "It's a tag!"
        DOCKER_TAG="$BRANCH_NAME"
    fi
fi

# Check the DOCKER_TAG has been resolved succesfully before procceding with the build
if [ ! -n "$DOCKER_TAG" ]
then
    echo "$0 - Error \$DOCKER_TAG not set or NULL"
    exit 1
fi

# Echo all the params!
echo "Golang version: $GOLANG_VERSION"
echo "Branch name: $BRANCH_NAME"
echo "Docker image name: $DOCKER_IMAGE_NAME"
echo "Docker tag: $DOCKER_TAG"

COMMAND="docker build --no-cache --build-arg GOLANG_IMAGE_VERSION=golang:$GOLANG_VERSION-alpine --build-arg BRANCH_NAME=$BRANCH_NAME -t pocket-core-$DOCKER_TAG -f docker/Dockerfile docker/."

# Run docker build!
eval $COMMAND
