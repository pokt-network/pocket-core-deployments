#!/bin/bash

set -e

mkdir -p /home/app/.pocket/config

if [ -z "${POCKET_CORE_GENESIS}" ] ; then
    true
else
    echo "${POCKET_CORE_GENESIS}" > /home/app/.pocket/config/genesis.json
fi

if [ -z "${POCKET_CORE_CHAINS}" ] ; then
    true
else
    echo "${POCKET_CORE_CHAINS}" > /home/app/.pocket/config/chains.json
fi

/usr/bin/expect /home/app/command.sh $@
