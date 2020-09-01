#!/bin/bash

curl -u $GIT_USER:$GIT_TOKEN https://api.github.com/user
git clone $GIT_REPO

cd pocket-core-deployments
git checkout $BRANCH_ARTIFACTS
sed -ie "s/^\(ARG IMAGE_NAME\)\(=\)\(.*\)$/\1\2${1}/" $SUBTITUTION_FILE

dappnodesdk publish minor --eth_provider $ETH_PROVIDER > transaction.txt

curl -X POST -H 'Content-type: application/json' --data "{\"blocks\": [{\"type\": \"section\",\"text\": {\"type\": \"mrkdwn\",\"text\": \"\`\`\` $(cat transaction.txt) \`\`\`\"}}]}" $SLACK_URL

git add .
git commit -m "Dappnode updated to version ${1} via Arti. [ci skip]"
git push origin HEAD


