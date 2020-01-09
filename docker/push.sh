#!/bin/sh

# Set error conditions
# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Use the error status of the first failure, rather than that of the last item in a pipeline.
set -o pipefail

# Variable declaration
BRANCH_NAME="$1"
DOCKER_TAG="$2"

# Parse parameters
if [ ! -n "$BRANCH_NAME" ]
then
    GOLANG_VERSION="staging"
fi

if [ ! -n "$DOCKER_TAG" ]
then
    # Resolve DOCKER_TAG using the branch name
    if [ "$BRANCH_NAME" == "staging" ]
    then
        # Handle staging branch
        DOCKER_TAG="staging-latest"
    elif [ "$BRANCH_NAME" == "master" ]
    then
        # Handle master branch
        DOCKER_TAG="latest"
    else
        # It's a tag!
        DOCKER_TAG="$BRANCH_NAME"
    fi
fi

# Check the DOCKER_TAG has been resolved succesfully before procceding with the build
if [ ! -n "$DOCKER_TAG" ]
then
    echo "$0 - Error \$DOCKER_TAG not set or NULL"
    exit 1
fi

# Push the image
exec docker push pocket-core:"$DOCKER_TAG"
