#!/usr/bin/bash
export TERM=linux
chown root -R /root/node/.data

if [ "$NETWORK" == "testnet" ]; then
    cp -r genesisfiles/testnet/genesis.json /root/node/data
fi

cp -r /root/node/data/* /root/node/.data/

./goal node start -d .data
watch -n 2 ./goal node status -d .data
