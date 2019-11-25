FROM ubuntu
RUN apt-get update && apt-get install -y wget
RUN wget https://github.com/particl/particl-core/releases/download/v0.18.1.6/particl-0.18.1.6-x86_64-linux-gnu.tar.gz
RUN tar -xvf particl-0.18.1.6-x86_64-linux-gnu.tar.gz
WORKDIR particl-0.18.1.6/bin

ENTRYPOINT ["./particld"]
