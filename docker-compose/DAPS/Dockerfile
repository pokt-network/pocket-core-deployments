FROM ubuntu:latest 
ARG VERSION
ENV CGO_ENABLED=0
RUN apt-get update && apt-get install -y unzip wget && wget https://github.com/DAPSCoin/DAPSCoin/releases/download/$VERSION/master_linux-v$VERSION.4.zip && unzip master_linux-v$VERSION.4.zip -d /usr/bin/ && chmod +x /usr/bin/dapscoin-cli && chmod +x /usr/bin/dapscoind  
ENTRYPOINT ["/usr/bin/dapscoind"]
