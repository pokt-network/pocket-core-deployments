#!/bin/bash

set -e

mkdir -p /home/app/.pocket/config

if [ -n "${POCKET_CORE_GENESIS}" ] ; then
    echo "${POCKET_CORE_GENESIS}" > /home/app/.pocket/config/genesis.json
fi

if [ -n "${POCKET_CORE_CHAINS}" ] ; then
    echo "${POCKET_CORE_CHAINS}" > /home/app/.pocket/config/chains.json
fi

if [ -n "${POCKET_CORE_CONFIG}" ] ; then
    echo "${POCKET_CORE_CONFIG}" > /home/app/.pocket/config/config.json
fi

# Allowing to distribute those files through a shared r/o volume for example
# (like Vault sidecar)

if [ -f "${POCKET_CORE_GENESIS_PATH}" ] ; then
    cp -f ${POCKET_CORE_GENESIS_PATH} /home/app/.pocket/config/genesis.json
fi

if [ -f "${POCKET_CORE_CHAINS_PATH}" ] ; then
    cp -f ${POCKET_CORE_CHAINS_PATH} /home/app/.pocket/config/chains.json
fi

if [ -f "${POCKET_CORE_CONFIG_PATH}" ] ; then
    cp -f ${POCKET_CORE_CONFIG_PATH} /home/app/.pocket/config/config.json
fi

if [ -f "${POCKET_CORE_KEY_FILE}" ] ; then
    export POCKET_CORE_KEY=$(cat ${POCKET_CORE_KEY_FILE} )
fi

if [ -f "${POCKET_CORE_PASSPHRASE_FILE}" ] ; then
    export POCKET_CORE_PASSPHRASE=$(cat ${POCKET_CORE_PASSPHRASE_FILE} )
fi


/usr/bin/expect /home/app/command.sh $@
