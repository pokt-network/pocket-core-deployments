#!/bin/bash

set -o errexit

DATA_DIR=/usr/src/app
DISPIP=dispatch.pokt.network
DISRPORT=443
PORT=${POCKET_CORE_SERVICE_PORT:-38081}
REQUEST_TIMEOUT=${POCKET_CORE_REQUEST_TIMEOUT:-0}

print_banner() {
    echo "#####################################################################################################"
    echo "## Pocket-core not configured to start!                                                            ##"
    echo "#####################################################################################################"
}

if [ ! -n "${POCKET_CORE_SERVICE_GID}" ]; then
    print_banner
    echo "##"
    echo "## You need to set a POCKET_CORE_SERVICE_GID to be able to start the package"
    echo "## here you can find information about how to get it:"
    echo "## https://docs.pokt.network/docs/how-to-participate#section-for-service-nodes"
    echo "##"
    echo "#####################################################################################################"
    while true; do sleep 5; done
fi 

if [ ! -n "${POCKET_CORE_SERVICE_IP}" ]; then
    print_banner
    echo "##"
    echo "## You need to set a POCKET_CORE_SERVICE_IP to be able to start the package"
    echo "## you can use your dappnode domain (XXXXXXXXXXXXXXX.dyndns.dappnode.io) or a static IP if you've one"
    echo "## here you can find more information about how to configure you dappnode:"
    echo "## https://docs.pokt.network/docs/service-node-dappnode-setup"
    echo "##"
    echo "#####################################################################################################"
    while true; do sleep 5; done
fi 

pocket-core \
    --datadirectory ${DATA_DIR} \
    --disip ${DISPIP} \
    --disrport ${DISRPORT} \
    --gid ${POCKET_CORE_SERVICE_GID} \
    --ip ${POCKET_CORE_SERVICE_IP} \
    --port ${PORT} \
    --requestTimeout ${REQUEST_TIMEOUT} 
