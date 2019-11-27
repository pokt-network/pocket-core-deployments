FROM golang:1.10-stretch

LABEL maintainer="lowell@pokt.network"

RUN apt update
RUN apt install git rsync wget -y && \
    wget "https://s3.eu-central-1.amazonaws.com/lightstreams-public/lightchain/latest/lightchain-linux-amd64" -O "/usr/local/bin/lightchain" && \
	wget "https://s3.eu-central-1.amazonaws.com/lightstreams-public/leth/latest/leth-linux-amd64" -O /usr/local/bin/leth && \
	chmod a+x /usr/local/bin/lightchain  && \
	chmod a+x /usr/local/bin/leth  && \
    lightchain version
    #git clone git@github.com:lightstreams-network/lightchain.git && \
    #mv lightchain/database /srv/lightchain/ && \
	

COPY entrypoint.sh /root/entrypoint.sh
RUN chmod a+x /root/entrypoint.sh &&   mkdir ${HOME}/.lightchain 

ENTRYPOINT ["/root/entrypoint.sh"]
CMD ["run", "--datadir=${DATADIR}", "--rpc", "--rpcaddr=0.0.0.0", "--rpcport=8545", "--rpcapi=eth,net,web3,personal,debug"]

EXPOSE 8545 26657 26656
