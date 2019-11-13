#!/bin/sh

# This script verifies if either TAG_NAME or BRANCH_NAME was provided, in case TAG_NAME is provided, it proceeds to clone the repo based on the TAG_NAME
# If BRANCH_NAME is provided, clones the repo branch with that name. 

if [ -n "$TAG_NAME" ] && [ -n "$BRANCH_NAME" ]
then
    echo "Can only send TAG_NAME or BRANCH_NAME variable, not both."
else
    if [ -n "$TAG_NAME" ]
    then
        git clone -b staging "https://github.com/pokt-network/pocket-core.git" $POCKET_PATH
        git checkout tags/$TAG_NAME
    else 
        if [ -n "$BRANCH_NAME" ]
        then
            git clone -b $BRANCH_NAME "https://github.com/pokt-network/pocket-core.git" $POCKET_PATH
        fi
    fi
fi
