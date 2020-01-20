FROM ubuntu:latest
ARG SPECIFIC_VERSION
ARG MAJOR_VERSION
ENV CGO_ENABLED=0
## Setting up user for blockchain.
RUN useradd -d /home/DAPS_user -m -s /bin/bash DAPS_user
WORKDIR /home/DAPS_user
USER root

RUN apt-get update && apt-get install -y unzip wget && wget https://github.com/DAPSCoin/DAPSCoin/releases/download/$MAJOR_VERSION/dapscoin-v$SPECIFIC_VERSION-linux.zip && unzip dapscoin-v$SPECIFIC_VERSION-linux.zip -d /home/DAPS_user/ && chmod +x /home/DAPS_user/dapscoin-cli && chmod +x /home/DAPS_user/dapscoind && chown -R DAPS_user /home/DAPS_user/
COPY entrypoint.sh /home/DAPS_user/
USER DAPS_user 
ENTRYPOINT ["/home/DAPS_user/dapscoind"]
