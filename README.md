
<div align="center">
  <a href="https://www.pokt.network">
    <img src="https://user-images.githubusercontent.com/16605170/74199287-94f17680-4c18-11ea-9de2-b094fab91431.png" alt="Pocket Network logo" width="340"/>
  </a>
</div>

# Pocket Core Deployments

Deployment artifacts for the [Pocket Network](https://pokt.network/).

## Overview

Dockerfiles and Compose YAMLs to orchestrate development and production-level environments of [Pocket Core](https://github.com/pokt-network/pocket-core).

## Documentation

The docker-compose folder contains all the different blockchains for running together with the Pocket Core service.

Every `docker-compose.yml` for every setup contains an additional `nginx` image optimized optimized with brotli and volumes for `persistent data` for the blockchains.

Every `docker-compose.yml` also contains the value `POCKET_CORE_START_DELAY` on the `pocket-service` container which is a variable that puts a delay while starting the pocket-service giving more time for the blockchains to start before receiving an erroneous response from them.

Before running, you need to define the three following environment variables:

#### POCKET_SERVICE_GID

This variable will be provided by the pocket team and then you need to export it in your environment.
```
> export POCKET_SERVICE_GID=XXX
```
#### ENV

This variable will define the pocket core version to run. By default it will run `mvp-staging-latest`
```
> export ENV=mvp-master-latest
```
#### POCKET_DISPATCH_IP

This variable contains the address to the pocket dispatcher. It defaults to `dispatch.pokt.network` which is our current production dispatch.

### Running

For running the `pocket-core` setup you only need to run `docker-compose up` inside the desired stack. For example in the case of aion, you only need to:
```
> `cd aion`
> `docker-compose up`
```
There's also a folder that contains all the deployments, if you want to run them all, do the same procedure as before.

_NOTE: You need at least 16GB of RAM and at least 2TB of hard drive if you wish to run all in a standalone setup._

## Contributing

Please read [CONTRIBUTING.md](https://github.com/pokt-network/pocket-core/blob/master/README.md) for details on contributions and the process of submitting pull requests.

## Support & Contact

<div>
  <a  href="https://twitter.com/poktnetwork" ><img src="https://img.shields.io/twitter/url/http/shields.io.svg?style=social"></a>
  <a href="https://t.me/POKTnetwork"><img src="https://img.shields.io/badge/Telegram-blue.svg"></a>
  <a href="https://www.facebook.com/POKTnetwork" ><img src="https://img.shields.io/badge/Facebook-red.svg"></a>
  <a href="https://research.pokt.network"><img src="https://img.shields.io/discourse/https/research.pokt.network/posts.svg"></a>
</div>

## License

This project is licensed under the MIT License; see the [LICENSE.md](LICENSE.md) file for details.
