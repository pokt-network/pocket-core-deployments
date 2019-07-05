#!/usr/bin/env bash

set -e

cmd="$@"
DATADIR="/srv/lightchain"


if [ -z "${NETWORK}" ]; then
	echo "Using default network sirius"
	NETWORK="sirius"
fi

if [ ! -d "${DATADIR}/database" ]; then
	echo "Initialize lightchain node in ${DATADIR}"
	lightchain init --datadir=${DATADIR} --${NETWORK}
fi
echo "${HOME}"
exec lightchain $cmd
