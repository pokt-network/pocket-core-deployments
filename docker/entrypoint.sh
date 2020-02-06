#!/bin/sh

set -e

if [ -n "${POCKET_CORE_GENESIS}" ] ; then
    echo "${POCKET_CORE_GENESIS}" > /home/app/.pocket/config/genesis.json
fi

if [ -n "${POCKET_CORE_CHAINS}" ] ; then
    echo "${POCKET_CORE_CHAINS}" > /home/app/.pocket/config/chains.json
fi

exec "$@"
