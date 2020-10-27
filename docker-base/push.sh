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
if [ ! -n "$GOLANG_VERSION" ]; then
    GOLANG_VERSION="1.13"
fi

if [ ! -n "$BRANCH_NAME" ]; then
    BRANCH_NAME="staging"
fi

if [ ! -n "$DOCKER_IMAGE_NAME" ]; then
    DOCKER_IMAGE_NAME="poktnetwork/pocket"
fi

if [ ! -n "$DOCKER_TAG" ]; then
    # Resolve DOCKER_TAG using the branch name
    case $BRANCH_NAME in
    staging)
        echo "It's devnet!"
        DOCKER_TAG="devnet-latest"
        ;;
    RC-*)
        echo "It's stagenet!"
        DOCKER_TAG="stagenet-latest"
        ;;
    Beta-*)
        echo "It's beta!"
        DOCKER_TAG="beta-latest"
        ;;
    *)
        if [[ $BRANCH_NAME =~ [0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9] ]]; then
            echo "It's prod!"
            DOCKER_TAG="$BRANCH_NAME"
        fi
        ;;
    esac


fi

# Check the DOCKER_TAG has been resolved succesfully before procceding with the build
if [ ! -n "$DOCKER_TAG" ]; then
    echo "$0 - Error \$DOCKER_TAG not set or NULL"
    exit 1
fi

# Echo all the params!
echo "Golang version: $GOLANG_VERSION"
echo "Branch name: $BRANCH_NAME"
echo "Docker tag: $DOCKER_TAG"
echo "Docker image name: $DOCKER_IMAGE_NAME"

# Run docker build!
BUILD_COMMAND="docker build --build-arg DOCKER_TAG=$DOCKER_TAG --build-arg BRANCH_NAME=$BRANCH_NAME -t pocket-core-$DOCKER_TAG -f docker-base/Dockerfile docker-base/."
eval $BUILD_COMMAND

TAG_COMMAND="docker tag pocket-core-$DOCKER_TAG:latest $DOCKER_IMAGE_NAME:$DOCKER_TAG"
eval $TAG_COMMAND

# Push image
case $BRANCH_NAME in
staging)
    eval "docker tag pocket-core-$DOCKER_TAG:latest $DOCKER_IMAGE_NAME:devnet-$CIRCLE_BUILD_NUM"
    eval "docker push $DOCKER_IMAGE_NAME:devnet-$CIRCLE_BUILD_NUM"
    echo staging
    ;;
RC-*)
    eval "docker tag pocket-core-$DOCKER_TAG:latest $DOCKER_IMAGE_NAME:$BRANCH_NAME"
    eval "docker push $DOCKER_IMAGE_NAME:$BRANCH_NAME"
    echo RC
    ;;
Beta-*)
    eval "docker tag pocket-core-$DOCKER_TAG:latest $DOCKER_IMAGE_NAME:$BRANCH_NAME"
    eval "docker push $DOCKER_IMAGE_NAME:$BRANCH_NAME"
    echo Beta
    ;;
# *)
#     if [[ $BRANCH_NAME =~ ^[0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9] ]]; then
#         IFS=. version=($BRANCH_NAME)
#         MAYOR=${version[0]}
#         MINOR="${version[0]}.${version[1]}"
#         for i in "$MAYOR" "$MINOR" "$BRANCH_NAME" "latest"
#         do
#             echo $i
#             eval "docker tag pocket-core-$DOCKER_TAG:latest $DOCKER_IMAGE_NAME:$i"
#             eval "docker push $DOCKER_IMAGE_NAME:$i"
#         done
#         exit 0
#     fi
#     ;;
esac


PUSH_COMMAND="docker push $DOCKER_IMAGE_NAME:$DOCKER_TAG"
echo "$PUSH_COMMAND"
eval $PUSH_COMMAND
