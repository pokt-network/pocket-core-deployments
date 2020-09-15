#!/bin/bash

default_seeds=b3d86cd8ab4aa0cb9861cb795d8d154e685a94cf@seed1.testnet.pokt.network:20656,17ca63e4ff7535a40512c550dd0267e519cafc1a@seed2.testnet.pokt.network:21656,f99386c6d7cd42a486c63ccd80f5fbea68759cd7@seed3.testnet.pokt.network:22656
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
"Ethereum Mainnet")
    cat > /home/app/.pocket/config/chains.json << EOF
    [{
        "id": "0021",
        "url": "http://my.ethchain.dnp.dappnode.eth:8545"
    }]
EOF

    ;;
"Ethereum Mainnet and Pocket Testnet")

    if [[ "${TLS_OPTION}" == "Do not use a certificate" ]]
    then
        cat > /home/app/.pocket/config/chains.json << EOF
        [{
            "id": "0002",
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
            "id": "0002",
            "url": "http://localhost:8082"
        }, {
            "id": "0021",
            "url": "http://my.ethchain.dnp.dappnode.eth:8545"
        }]
EOF
    fi
    ;;  
"Pocket Testnet")
    if [[ "${TLS_OPTION}" == "Do not use a certificate" ]]
    then
        cat > /home/app/.pocket/config/chains.json << EOF
        [{
            "id": "0002",
            "url": "http://localhost:8081"
        }]
EOF
        sed -ie s/\"rpc_port\"\.\*/\"rpc_port\"\:\ \"8081\",/ /home/app/.pocket/config/config.json
    else
        cat > /home/app/.pocket/config/chains.json << EOF
        [{
            "id": "0002",
            "url": "http://localhost:8082"
        }]
EOF
    fi
EOF

    ;;
"Rinkeby")
    cat > /home/app/.pocket/config/chains.json << EOF
    [{
        "id": "0022",
        "url": "http://my.rinkeby.dnp.dappnode.eth:8545"
    }]
EOF

    ;;
"Rinkeby and Pocket Testnet")

    if [[ "${TLS_OPTION}" == "Do not use a certificate" ]]
    then
        cat > /home/app/.pocket/config/chains.json << EOF
        [{
            "id": "0002",
            "url": "http://localhost:8081"
        }, {
            "id": "0022",
            "url": "http://my.rinkeby.dnp.dappnode.eth:8545"
        }]
EOF
        sed -ie s/\"rpc_port\"\.\*/\"rpc_port\"\:\ \"8081\",/ /home/app/.pocket/config/config.json
    else
        cat > /home/app/.pocket/config/chains.json << EOF
        [{
            "id": "0002",
            "url": "http://localhost:8082"
        }, {
            "id": "0022",
            "url": "http://my.rinkeby.dnp.dappnode.eth:8545"
        }]
EOF
    fi
    ;;
"Upload chains.json")
    cp /tmp/chains.json /home/app/.pocket/config/
    ;;



esac

pocket start --testnet


        "Pocket Testnet",
        "Rinkeby",
        "Ethereum Mainnet",
        "Rinkeby and Pocket Testnet",
        "Ethereum Mainnet and Pocket Testnet",
        "Upload chains.json"