
<div align="center">
  <a href="https://www.pokt.network">
    <img src="https://user-images.githubusercontent.com/16605170/74199287-94f17680-4c18-11ea-9de2-b094fab91431.png" alt="Pocket Network logo" width="340"/>
  </a>
</div>

# Pocket Core Deployments

Deployment artifacts for [Pocket Network](https://pokt.network/).

## Table of Contents
1. [Overview](#Overview)
2. [Documentation](#Documentation)
    - [Docker](#docker)
        - [Building](#docker-build)
        - [Running](#docker-run)
            - [Storage](#docker-run-volume)
    - [Docker Compose](#compose)
        - [Building](#compose-build)
        - [Running](#compose-run)
    - [Kubernetes](#k8s)
        - [Deployment](#k8s-deployment)
        - [Statefulset](#k8s-statefulset)
    - [Hombrew](#homebrew)
    - [FAQ](#faq)
3. [Contributing](#contributing)
4. [Support & Contact](#support)



## Overview

This repository contains deployment artifacts to orchestrate [Pocket Core](https://github.com/pokt-network/pocket-core).

## <b>Documentation</b>

### <b>Docker<a name="docker"></a></b>
#### <b>Requirements</b>
- [Docker](https://docker.com/)

#### <b>Usage</b>
Pocket offers two images that could be used instead of building your own. These are [poktnetwork/pocket-core](https://hub.docker.com/r/poktnetwork/pocket-core/tags?page=1&ordering=last_updated) and [poktnetwork/pocket](https://hub.docker.com/r/poktnetwork/pocket/tags?page=1&ordering=last_updated). In order to use them you need to pull them from the docker registry using the following command:
```
docker pull poktnetwork/pocket-core:<tag>
```
or
```
docker pull poktnetwork/pocket:<tag>
```
You must replace `<tag>` by the version of pocket you want to run. For example `RC-0.6.3`.

The main difference between `poktnetwork/pocket-core` and `poktnetwork/pocket` is that the later does not contain any scripts to help you start your node.  

The `poktnetwork/pocket-core` image can be configured via environment variables. Here is the list of environment variables available:
```
POCKET_CORE_KEY            Raw private key             
POCKET_CORE_PASSPHRASE     Passphrase to add private key to keybase. Also used to set the validator 
POCKET_CORE_CONFIG         Pocket configuration file (json between single quotes)
POCKET_CORE_CHAINS         List of chains your node will server (json between single quotes)
```

You can start your node using the following command:  
`
docker run  poktnetwork/pocket-core:RC-0.6.3 pocket start --mainnet --datadir=/home/app/.pocket`

For instance, if you want to add extra configuration to the container, you won't be able to do so using the `poktnetwork/pocket-core`. In this case you want to use the `poktnetwork/pocket` image. It comes only with the pocket binary installed so you can add custom scripts and configurations.

## <b>Docker Compose<a name="compose"></b></a>

#### <b>Requirements</b>
- [Docker](https://docker.com/)
- [Docker compose](https://docs.docker.com/compose/)

Docker compose is a tool used to run and manage multiple containers. Use docker compose when running more than one pocket node or you want to run a pocket node plus one or multiple blockachain nodes. Note that we also provide docker compose files for various blockchains. It can be found under the `blockchains` directory.

To start your docker compose you must run `docker-compose up` from the directory containing your `docker-compose.yaml` file. This will start all services listed in your `docker-compose.yaml`. To restart the node run `docker-compose restart <service name>`. In case you want to stop the container use `docker-compose stop <service name>`, do not use `docker-compose down` unless you want to tear down your setup.

This is what a docker-compose file would look like for running a pocket node using the `poktnetwork/pocket` image:

```
version: '3'
services:
  pocket:
    image: poktnetwork/pocket:RC-0.6.3
    command: pocket start --mainnet --keybase=false --datadir=/path/to/datadir
    ports:
      - "8081:8081"
      - "26656:26656"
    networks:
      - pocket
    volumes:
      - pocket:/path/to/datadir
      - /path/to/config.json:/home/app/.pocket/config/config.json
      - /path/to/chains.json:/home/app/.pocket/config/chains.json
      - /path/to/node_key.json:/home/app/.pocket/node_key.json
      - /path/to/priv_val_key.json:/home/app/.pocket/priv_val_key.json
volumes:
  pocket:
networks:
  pocket:
    driver: bridge
```

#### Fields

```
image       Select which pocket image you want to run. As stated above, officially we provide two image: poktnetwork/pocket and poktnetwork/pocket-core. 
command     Which command the container will run when your node starts.
ports       Port mapping to your pocket node and your local machine. You can specify more than one.
volumes     Map volumes or bind local files to your node.
```

## <b>Kubernetes<a name="k8s"></b></a>
Coming soon.

## <b>Homebrew<a name="homebrew"></b></a>

[Homebrew documentation.](https://github.com/pokt-network/homebrew-pocket-core)

## <b>Stack<a name="stack"></b></a>
[Stack documentation.](/stacks/pocket-validator)

# FAQ<a name="faq"></a>
## I get "sh: -c requires an argument" when I start the container.
Make sure you pass a command to the container like: 
```
docker run poktnetwork/pocket-core:RC-0.6.3 pocket start --mainnet --datadir=/home/app/.pocket
```
Note that after the container image you have pass the pocket command to start running your node.

### When I try to run the container I get "canot open/create json file: open /dir/path/auth.json: no such file or directory"
Check if the datadir properties in your config.json are set to /home/app/.pocket or to the path specified when using the `--datadir` flag.

### I cannot ctrl+c to exit the container
Make sure when you started your container you added the `-ti` arguments. It will allow you to interact with the tty.



## Contributing<a name="contributing"></a>

Please read [CONTRIBUTING.md](https://github.com/pokt-network/pocket-core/blob/master/README.md) for details on contributions and the process of submitting pull requests.

## Support & Contact<a name="support"></a>

<div>
  <a  href="https://twitter.com/poktnetwork" ><img src="https://img.shields.io/twitter/url/http/shields.io.svg?style=social"></a>
  <a href="https://t.me/POKTnetwork"><img src="https://img.shields.io/badge/Telegram-blue.svg"></a>
  <a href="https://www.facebook.com/POKTnetwork" ><img src="https://img.shields.io/badge/Facebook-red.svg"></a>
  <a href="https://research.pokt.network"><img src="https://img.shields.io/discourse/https/research.pokt.network/posts.svg"></a>
</div>

## License

This project is licensed under the MIT License; see the [LICENSE.md](LICENSE.md) file for details.
