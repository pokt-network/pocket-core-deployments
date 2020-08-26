#!/bin/bash

default_seeds=03b74fa3c68356bb40d58ecc10129479b159a145@seed1.mainnet.pokt.network:20656,64c91701ea98440bc3674fdb9a99311461cdfd6f@seed2.mainnet.pokt.network:21656
expect /home/app/account.sh


if [[ "${CONFIG}" == "Default" ]]
    then
        cp /home/app/config.json /home/app/.pocket/config/config.json
        sed -ie s/\"Moniker\"\.\*/\"Moniker\"\:\ \"$RANDOM\",/ /home/app/.pocket/config/config.json
    else
        cp /tmp/config.json /home/app/.pocket/config/config.json
fi


if [[ "${SEEDS}" == "Enter seeds" ]]
    then
        sed -ie s/\"Seeds\"\.\*/\"Seeds\"\:\ \"$CUSTOM_SEEDS\",/ /home/app/.pocket/config/config.json
elif [[ "${SEEDS}" == "Default" ]]
    then
        sed -ie s/\"Seeds\"\.\*/\"Seeds\"\:\ \"$default_seeds\",/ /home/app/.pocket/config/config.json
fi


case ${CHAINS} in
"Only Ethereum Mainnet")
    cat > /home/app/.pocket/config/chains.json << EOF
    [{
        "id": "0021",
        "url": "http://my.ethchain.dnp.dappnode.eth:8545"
    }]
EOF

    ;;
"Ethereum mainnet and Pocket mainnet")

    if [[ "${TLS_OPTION}" == "Do not use a certificate" ]]
    then
        cat > /home/app/.pocket/config/chains.json << EOF
        [{
            "id": "0001",
            "url": "http://localhost:8081"
        }, {
            "id": "0021",
            "url": "http://my.ethchain.dnp.dappnode.eth:8545"
        }]
EOF
        sed -ie s/\"rpc_port\"\.\*/\"rpc_port\"\:\ \"8081\",/ /home/app/.pocket/config/config.json
    else
        cat > /home/app/.pocket/config/chains.json << EOF
        [{
            "id": "0001",
            "url": "http://localhost:8082"
        }, {
            "id": "0021",
            "url": "http://my.ethchain.dnp.dappnode.eth:8545"
        }]
EOF
    fi





    ;;
esac

pocket start --mainnet