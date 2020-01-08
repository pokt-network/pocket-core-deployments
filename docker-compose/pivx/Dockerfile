FROM ubuntu:latest
ARG VERSION
LABEL maintainer="lowell@pokt.network"
RUN useradd -d /home/pivx_user -m -s /bin/bash pivx_user
ENV CGO_ENABLED=0
RUN apt-get update && apt-get install -y unzip wget && wget https://github.com/PIVX-Project/PIVX/releases/download/v$VERSION/pivx-$VERSION-x86_64-linux-gnu.tar.gz && tar -xvf pivx-$VERSION-x86_64-linux-gnu.tar.gz && chmod +x /pivx-$VERSION/bin/* && ln -s /pivx-$VERSION/bin/pivxd /usr/bin/pivxd && ln -s /pivx-$VERSION/bin/pivx-cli /usr/bin/pivx-cli
USER pivx_user
ENTRYPOINT ["/usr/bin/pivxd"]
