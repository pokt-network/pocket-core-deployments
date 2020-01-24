#!/bin/sh

set -e

if ${POCKET_CORE_GENESIS:-false} ; then
    echo "${POCKET_CORE_GENESIS}" > /home/app/.pocket/genesis.json
fi

if ${POCKET_CORE_CHAINS:-false} ; then
    echo "${POCKET_CORE_CHAINS}" > /home/app/.pocket/chains.json
fi

exec "$@"
