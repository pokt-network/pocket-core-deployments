#!/bin/sh

# Set error conditions
# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Parse arguments
API_KEY="$1"
BRANCH_NAME="$2"
POCKET_CORE_DEPLOYMENTS_BRANCH="$3"

# Parse parameters
if [ ! -n "$POCKET_CORE_DEPLOYMENTS_BRANCH" ]
then
    POCKET_CORE_DEPLOYMENTS_BRANCH="master"
fi

if [ ! -n "$API_KEY"] || [ ! -n "$BRANCH_NAME"]
then
	exit 1
fi

TRIGGER_COMMAND="curl -u $API_KEY: -X POST -H \"Content-Type: application/json\" -d '{\"branch\": \"$POCKET_CORE_DEPLOYMENTS_BRANCH\",\"parameters\":{\"branch-name\":\"$BRANCH_NAME\"}}' https://circleci.com/api/v2/project/github/pokt-network/homebrew-pocket-core/pipeline"
eval $TRIGGER_COMMAND
