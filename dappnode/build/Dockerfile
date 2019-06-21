FROM poktnetwork/pocket-core:mvp-1.7
ENTRYPOINT pocket-core --datadirectory datadir --disip dispatch.pokt.network --gid ${POCKET_CORE_SERVICE_GID:-MY_GID} --ip ${POCKET_CORE_SERVICE_IP:-127.0.0.1} --disrport 443 --port ${POCKET_CORE_SERVICE_PORT:-8081} --requestTimeout ${POCKET_CORE_REQUEST_TIMEOUT:-0} 
