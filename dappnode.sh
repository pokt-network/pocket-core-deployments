#!/bin/bash

# Version comparison function

function splitVersion() {
    IFS=. old_version=($1)
    old_mayor=${old_version[0]}
    old_minor=${old_version[1]}
    old_patch=${old_version[2]}

    IFS=. new_version=($2)
    new_mayor=${new_version[0]}
    new_minor=${new_version[1]}
    new_patch=${new_version[2]}
}

function compare() {
    if [[ $1 =~ ^[0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9]$ ]] && [[ $2 =~ ^[0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9]$ ]]; then
        splitVersion $1 $2

    elif [[ $1 =~ ^RC-[0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9]$ ]] && [[ $2 =~ ^RC-[0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9]$ ]]; then
        oldTag=${1:3}
        newTag=${2:3}
        splitVersion $oldTag $newTag

    elif [[ $1 =~ ^RC-[0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9]$ ]] && [[ $2 =~ ^[0-9]\.[0-9]|[0-9][0-9]\.[0-9]|[0-9][0-9]$ ]]; then
        oldTag=${1:3}
        newTag=$2
        splitVersion $oldTag $newTag

    fi

    if ((old_mayor < new_mayor)); then
        updateType="major"
        return 0
    elif ((old_minor < new_minor)); then
        updateType="minor"
        return 0
    elif ((old_patch < new_patch)); then
        updateType="patch"
        return 0
    else
        return 1
    fi

}

# Publishes dappnode package, send slack notification
# Usage: publish <package-directory> <network>
function publish() {
    sed -ie "s/^\(ARG IMAGE_NAME\)\(=\)\(.*\)$/\1\2$NEW_VERSION/" $SUBTITUTION_FILE
    # Create transaction to publish dappnode package
    update=$(dappnodesdk publish $updateType --directory $1)

    # Leave only trasaction on log
    sed '0,/Generate\ transaction \[completed\]/d' $update

    # Send dappnode transaction to Slack
    payload=$(jq --arg transaction "\`\`\` DAppNode $2 transacction\n $(cat transaction.log)\`\`\`" -n '{"blocks": [{"type":"header","text":{"type": "plain_text","text": "DAppNode Package","emoji": true}},{"type":"divider"},{"type": "section","text": {"type": "mrkdwn","text": $transaction}}]}')
    curl -X POST -H 'Content-type: application/json' --data "$payload" $SLACK_URL
}

# Clone pocket-core-deployments
git clone $GIT_REPO
cd $GIT_REPO_DIR_DAPPNODE
git checkout $BRANCH_ARTIFACTS

compare $OLD_VERSION $NEW_VERSION

publish dappnode mainnet
publish dappnode-testnet testnet


expect expect.sh $NEW_VERSION DAppNode $GIT_REPO_DIR_DAPPNODE $GIT_BRANCH_DAPPNODE

exit
