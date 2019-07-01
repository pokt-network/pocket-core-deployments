# Pocket-core docker-compose


This folder contains all the different blockchains for running together with the pocket-service

Every `docker-compose.yml` for every setup contains an additional `nginx` image optimized optimized with brotli and volumes for `persitent data` for the blockchains

Every `docker-compose.yml` also contains the value `POCKET_CORE_START_DELAY` on the `pocket-service` container which is a variable that puts a delay while starting the pocket-service giving more time for the blockchains to start before receiving an erroneous response from them

Before running, you need to define the two following environment variables 


### POCKET_SERVICE_GID

This variable will be provided by the pocket team and then you need to export it in your environment

> export POCKET_SERVICE_GID=XXX


### ENV

This variable will define the pocket core version to run. By default it will run `mvp-staging-latest` 

> export ENV=mvp-master-latest

### POCKET_DISPATCH_IP

This variable contains the address to the pocket dispatcher. It defaults to `dispatch.pokt.network` which is our current production dispatch 


### Running

For running the `pocket-core` setup you only need to run `docker-compose up` inside the desired stack. For example in the case of aion, you only need to:


> `cd aion`

> `docker-compose up`

There's also a folder that contains all the deployments, if you want to run them all, do the same procedure as before

_NOTE: You need at least 16GB of RAM and at least 2TB of hard drive if you wish to run all in a standalone setup_
