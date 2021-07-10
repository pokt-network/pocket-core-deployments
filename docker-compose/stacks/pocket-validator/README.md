# POCKET VALIDATOR STACK 

## Overview

This folder contains a docker-compose.yaml with:

- Single pocket node validator
- Grafana/prometheus/loki/cadvisor/alertmanager for monitoring logs/metrics
- Reverse nginx proxy and certbot certificate generation
- Automatic certificate generation

The purpose is to provide a closest production-ready stack in order to serve traffic to your pocket-validator 

We packed basics of security and other configurations. Security and server hardening it's up to the responsability of all node runners 


## Requirements

- 15,200 POKT (15,100 for stake and 100 at least for the transactions operations like unjail, stake)
- Your own domain
- A pocket wallet with the quantity mentioned above, you can create one at https://wallet.pokt.network/ or using the cli (make sure you download the keyfile.json for later)
- docker (19.03.0+) and docker-compose(1.27.4+) installed on the system 

## Hardware requirements 


### Minimum requirements 

- CPU: 2 CPUs
- Memory: 4 GB RAM
- Disk: 80GB SSD 


### Recommended requirements 

- CPU: 4 CPUs
- Memory: 6 GB RAM
- Disk: 120GB SSD 


_For more information, see [Pocket node hardware requirements](https://docs.pokt.network/docs/before-you-dive-in#hardware-requirements)_


## Instructions 


### Network configuration


It's convenient for extra security that this server resides in a private network. And only expose the pocket peer port and HTTP/HTTPS to public via loadbalancer.

In case you are running in public network. we included instructions for blocking the ports and included basic security in NGINX. 


### Port configuration 


Open the ports:

* SSH (22/tcp) 	 				
    * Only accessed by your IP  (x.x.x.x/32)
* HTTP/HTTPS for the nginx proxy (80/TCP, 443/TCP)	
	* Public access (0.0.0.0/0) needed for the grafana dashboard  
* Pocket peers (TCP/26656)
	* Public access (0.0.0.0/0)

Verify to block all the other ports and harden the SSH access to only connect using your machine IP and a keypair 

### Create domain records 

Create the following domain A records pointing to your pocket-validator server IP:

- node1.${DOMAIN}
- monitoring.${DOMAIN} 


After you finish, wait 5-10 mins or the time required given by your DNS until the spread is over. 

You can verify if domain are correctly configured by checking with nslookup that the domains return your IP: 

```bash
 nslookup node1.${DOMAIN}

 nslookup monitoring.${DOMAIN}
```


### Installing Loki driver plugin and set grafana/prometheus permissions 


First, clone the repo and move to this directory in order to continue the tutorial

```bash
git clone https://github.com/pokt-network/pocket-core-deployments.git 
cd pocket-core-deployments/docker-compose/stacks/pocket-validator
```

The following script will install the loki driver for sending blockchain node logs to loki and grant file permissions needed for grafana and prometheus. Just run this command in the host machine:

```
sudo bash install.sh
```

#### Create and validate your SSL certificate 

After setting up your domain A records, let's now generate our SSL domain certificate for the entire domain (*.yourdomain.com) by doing: 


```bash
docker run --volume /etc/letsencrypt/:/etc/letsencrypt/  -it certbot/certbot:latest certonly --manual --agree-tos --no-eff-email --preferred-challenges=dns -d \*.yourdomain.com -d yourdomain.com
```

Certbot will tell you to create a DNS TXT record _acme-challenge.<yourdomain> with a provided TXT value, create it, wait 5-10m until it propagates and test with:


```bash
nslookup -type=txt _acme-challenge.yourdomain.com
```

Once this command shows you the TXT you entered for your domain, you can hit enter and proceed in the certbot window 

> Note: In case you cannot verify. Retry the command and when you set the value of the TXT subdomain, wait a little bit longer 

If you finished. you will have your certificate succesfully generated. which we will be used by the nginx proxy to server the web server and by the certbot-renew service to be renewed automatically

### Setting up proxy and monitoring systems 

We use nginx for the web proxy and (loki/grafana/prometheus) stack for the monitoring systems

You can see their configurations in the docker-compose volume mappings and their respective folders if you wish to customize then for yourself 

### Set your domain and configure grafana access 

Change the following settings according your setup:
    - The env variable `DOMAIN`  in the file `.env` used in docker-compose web services. Which indicates to nginx proxy what domain to use and certbot for generating the certificates
        For more info about the proxy configuration, you can see the (conf.d/https.conf.template)[]
    - The env variable `GRAFANA_ADMIN_PASS` in .env. Which is the grafana login password on monitoring.{DOMAIN}. The default login user is admin. 


## Configuring your validator

For more info. See: 
https://docs.pokt.network/docs/create-validator-node#


### Configure your chains.json file

Next step, configure your chains.json located in node1/chains.json serving your blockchains as follows:


```json
[
    {
        "id": "0001",
        "url": "https://pocket-mainnet.yourdomain/",
    },
    {
        "id": "0021",
        "url": "https://eth-mainnet.yourdomain/",
    }
``` 

In case you have your URLS protected by basic_auth as shown in the README tutorial in pocket-supported-blockchains folder, you need:

```json
[
    {
        "id": "0001",
        "url": "https://pocket-mainnet.yourdomain/",
	"basic_auth": {
		"username": "YOURUSERNAME",
		"password": "YOURSECUREDPASSWORD"
    	}
    },
    {
        "id": "0021",
        "url": "https://eth-mainnet.yourdomain/",
	"basic_auth": {
		"username": "YOURUSERNAME",
		"password": "YOURSECUREDPASSWORD"
    	}
    }
]
``` 

For more information about chains, please see:

- [New chains](https://forum.pokt.network/t/pip-6-2-settlers-of-new-chains/1027) 

### Obtaining your node_key.json and priv_val_key.json


If your have your `node_key.json` and `priv_val_key.json`, just move those files inside node1 and skip this step

In case you don't have the files mentioned. Get your `node_key.json` and `priv_val_key.json` by executing the following commands in the host machine:


```bash
docker-compose up -d  # Run all services

docker exec -it node1 sh # Enter to your node1 container


#### import account by using private key


```bash
pocket accounts import-raw <YOUR PRIVATE_KEY> # Enter your password
```

####  Or you can import by keyfile

```bash
echo '${KEYFILE}' > keyfile.json # Or just copy and paste your keyfile content using nano or vim

pocket accounts import-armored keyfile.json # Enter your password 

pocket accounts set-validator {YOURADDRESS}
```

Now, inside your container, copy the content of `priv_val_key.json` and `node_key.json` from your node1 to node1 folder on your host


```bash
cp /root/.pocket/priv_val_key.json  /home/app/.pocket 

cp /root/.pocket/node_key.json  /home/app/.pocket

exit
```


Restart you node for refreshing the changes:


```bash
docker stop node1 && docker rm node1 && docker-compose up -d
```

For verifying that your node is configured with your address you can do:

```bash
docker exec -it node1 sh

apk add curl

curl localhost:26657/status 
``` 

The output of the curl will get you the information of the node:

```
"validator_info": {
      "address": "<ADDRESS>", # Address of the node
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "<key>"
      },
      "voting_power": "0"

```

Verify that the address in this node matches your Address, in case not retry the step and verify that you are creating your files correctly where you need to pass the node_key.json and priv_val_key.json  


### Setting up pocket core node config.json 


Configure your pocket node by editing on `node1/config.json` the following variables:
    -  "Moniker" - Use a custom Moniker name for your node, in this case can be `node1-${DOMAIN}`
    -  "ExternalAddress" - point to your node address and port, in this case `tcp://node1.${DOMAIN}:26656`

NOTE: In this case you need to manual replace ${DOMAIN} by your domain name


### Set your POCKET_CORE_PASSPHRASE for this node

Inside the `.env` file. Fill the env variable with your node passphrase (same used in the step while importing account before `pocket accounts import-armored keyfile.json` 


## Sync/stake your node 

We'll provide a temporary backup in order to avoid syncing from scratch for the latest version. So let's stop our current stack and download/extract this backup


### gsutil backup (recommended, no need to untar)

Requires [gsutil installed](https://cloud.google.com/storage/docs/gsutil_install)


```bash
docker-compose down # For stopping the nodes
gsutil -m cp -r  gs://pocket-blockchain-backup-latest/* node1/data/
```

### tar backup (need 2x the blockchain space)

```bash
docker-compose down # For stopping the nodes
rm -rf node1/data/* # for deleting current datadir
wget https://storage.googleapis.com/blockchains-data/pocket-network-mainnet-latest.tar 
tar -xvf pocket-network-mainnet-latest.tar -C node1/ 
```

Restart your stack so it reflect the changes


```bash
docker-compose down && docker-compose up -d 
```

Please verify that your container node is up and it's syncing

While syncing, verify that your network proxy nginx and grafana stack is correctly configured by entering to the grafana login on:

> https://monitoring.${DOMAIN} 

Also, verify if your pocket node is correctly exposed by checking the pocket core ver on:

> https://node1.${DOMAIN}/v1

As last step, stake your node by doing the following commands inside your pocket node:


```bash
# Staking Command
pocket nodes stake  <fromAddr> <amount in uPOKT> <chains> <serviceURI w/443> <chainID> <fees in Upokt> <legacy flag> 

# example with dummy values
pocket nodes stake 45D50DB64E90C0109C778DAAB7EF36676FC03866 1510000000 0001,0021 https://my-pocket-url:443 mainnet 10000 true 
``` 

Wait one block (~15 mins) and your node should be ready to serve relays

In case you want to verify. Please do:

```bash
docker exec -it node1 sh

pocket query node <youraddr> # It should show your addr and the domain of your node

pocket query height # You should be in the latest height of the blockchain. See https://explorer.pokt.network/
```


For additional information you can see [Staking your node](https://docs.pokt.network/docs/create-validator-node#staking-your-node)


## References

- [FAQ for nodes](https://docs.pokt.network/docs/faq-for-nodes)

- [Create a pocket validator node](https://docs.pokt.network/docs/create-validator-node)


## IMPORTANT NOTE

If you used this tutorial before and you staked your node with the RPC port 8081. Meaning that if your service URL looks like:

"service_url": "https://node1.{DOMAIN}:8081"

You need to:

- Open the port 8081 in your firewall
- Add 8081:8081 in the [ports for the web(nginx) container](https://github.com/pokt-network/pocket-core-deployments/blob/staging/docker-compose/stacks/pocket-validator/docker-compose.yaml#L17)
- Change the listen address [from 443 ssl; to 8081 ssl; on the proxy https URLS](https://github.com/pokt-network/pocket-core-deployments/blob/staging/docker-compose/stacks/pocket-validator/proxy/conf.d/https.conf.template#L22)
- Restart the containers with ``` docker-compose down && docker-compose up -d ```

For verification you can check that your node is serving in this port by checking if https://node1.yourdomain:8081/v1 returns the pocket version
 
In any case, feel free to refer to our discord or create an issue for any questions

## Troubleshooting notes


### My pocket node container is hang or doesn't stop/respond 

- It's very possible that when you stop your pocket container or while doing docker-compose down your container get's stuck stoppping. In this case you need to stop the daemon and run the stack again as follows:


```bash
sudo systemctl restart docker && docker-compose down && docker-compose up -d
```

### The cadvisor container is showing permission denied error


- This particular error is related to the volume permission and it can vary by OS, you can see a fix here https://github.com/google/cadvisor/issues/2387#issuecomment-600840479


### Your node crashes with the err "priv_val_state.json: device or resource busy"

- For this case, be sure to have your `priv_val_state.json` empty and that you are not mapping that file in.

### You can't access the node1 to generate the `priv_val_key.json` and `node_key.json`
- For this case you can execute those commands in the host machine:
```bash
# From this command output take the IMAGE ID that appears on the image of poktnetwork/pocket-core
docker images
# Substitute the ${ID} variable below with the IMAGE ID you got from the last command
docker run -it --entrypoint bash --user root -w '/root' ${ID}
```

### You get a permission error for the `priv_val_key.json`, `config.json` or `node_key.json` file inside the container.
- To fix this case you can change the permission of the folder of node1 inside the host machine using:
```bash
## This will change the permission of the folder and its's content to be accessed by everyone.
chmod 777 -R node1/
```
