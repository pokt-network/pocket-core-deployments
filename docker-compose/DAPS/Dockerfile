FROM ubuntu:latest
ARG VERSION
ENV CGO_ENABLED=0
## Setting up user for blockchain.
RUN useradd -d /home/DAPS_user -m -s /bin/bash DAPS_user
WORKDIR /home/DAPS_user
USER root
RUN apt-get update && apt-get install -y unzip wget && wget https://github.com/DAPSCoin/DAPSCoin/releases/download/$VERSION/master_linux-v$VERSION.4.zip && unzip master_linux-v$VERSION.4.zip -d /home/DAPS_user/ && chmod +x /home/DAPS_user/dapscoin-cli && chmod +x /home/DAPS_user/dapscoind
USER DAPS_user
ENTRYPOINT ["/home/DAPS_user/dapscoind"]
