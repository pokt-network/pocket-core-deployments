# POCKET VALIDATOR STACK 

## Overview


This folder contains a docker-compose.yaml with:

- Single pocket node validator
- Grafana/prometheus/loki/cadvisor/alertmanager for monitoring logs/metrics
- Reverse nginx proxy with basic WAF and certbot certificate generation

The purpose is to provide a closest production-ready stack in order to serve traffic for your pocket-validator 

We packed basics of security and other configurations. Security and server hardening it's up to the responsability of all node runners 


## Requirements

- 15,200 POKT (15,100 for stake and 100 at least for the transactions operations like unjail, stake)
- Your own domain
- A pocket wallet with the quantity mentioned above, you can create one at https://wallet.pokt.network/ or using the cli (make sure you download the keyfile.json for later)
- docker (19.03.0+) and docker-compose(1.27.4+) installed on the system 

### Hardware requirements 


#### Minimum requirements 


- CPU: 2 CPUs
- Memory: 4 GB RAM
- Disk: 60GB SSD 

For more information, see [Pocket node hardware requirements](https://docs.pokt.network/docs/before-you-dive-in#hardware-requirements)


#### Recommended requirements 


- CPU: 4 CPUs
- Memory: 6 GB RAM
- Disk: 100GB SSD 

For more information, see [Pocket node hardware requirements](https://docs.pokt.network/docs/before-you-dive-in#hardware-requirements)


## Instructions 


### Network configuration


It's convenient for extra security that this server resides in a private network. And only expose the pocket peer port and HTTP/HTTPS to public via loadbalancer.

In case you are running in public network. we included instructions for blocking the ports and included basic security in NGINX. 


#### Port configuration 


Open the ports:

* SSH 	 				
    * Only accessed by your IP  (x.x.x.x/32)
* HTTP/HTTPS for the nginx proxy  	
	* Public access (0.0.0.0/0) needed for the grafana dashboard  
* TCP/26656 for pocket peers
	* Public access (0.0.0.0/0)

Verify to block all the other ports and harden the SSH access to only connect using your machine IP and a keypair 

####  Create domain records 

Create the following domain A records pointing to your pocket-validator server IP:

- node1.${DOMAIN}
- monitoring.${DOMAIN} 


After you finish, wait 5-10 mins or the time required given by your DNS until the spread is over. 


You can verify if domain are correctly configured by checking with nslookup that the domains return your IP: 

```
 nslookup node1.${DOMAIN}

 nslookup monitoring.${DOMAIN}
```


#### Install loki driver and set grafana/prometheus permissions 


The following script will install the loki driver for sending blockchain node logs to loki and grant file permissions needed for grafana and prometheus 

```
sudo bash install.sh
```


#### Create and validate your SSL certificate 


Edit the env variables `DOMAIN` and `EMAIL` in the file `.env` with the values for your setup. 
Those variables are used in docker-compose web, certbot services. Which indicates to nginx proxy what domain to use and certbot for generating the certificates

For more info about the proxy configuration, you can see the conf.d/https.conf.template


After setting up your domain A records, let's now generate our SSL certificates by doing: 


```
docker-compose up web certbot 
```


You should see a message from certbot about the HTTP challenge going on and then a "congratulations!" message for succesfully generating your SSL certificate

When you generate your SSL certificate successfully you can stop the web and certbot service and continue the installation procedure 


NOTE: In case you are not able to get the certificate be sure to give more time to the IP change to propagate or best test adding `--staging` This will prevent you from getting timeout or ratelimit from using certbot. Once you get it working, remember to remove the `--staging` parameter and remove the test certificate found at proxy/certbot/conf/live/node1.${DOMAIN}. The command looks like:

```
certonly --webroot --webroot-path=/var/www/certbot --email ${EMAIL} --agree-tos --no-eff-email --staging -d node1.${DOMAIN} -d monitoring.${DOMAIN}
```


#### Uncomment proxy routes

After your certbot certificate is issued, you can stop uncomment all the '#' on the file `proxy/conf.d/https.conf.templates`


### Setting up proxy and monitoring systems 


We use nginx for the web proxy and (loki/grafana/prometheus) stack for the monitoring systems

You can see their configurations in the docker-compose volume mappings and their respective folders if you wish to customize then for yourself
 

#### Set your domain and configure grafana access 
 

Change the following settings according your setup:
    - The env variable `DOMAIN` and `EMAIL` in the file `.env` used in docker-compose web, certbot services. Which indicates to nginx proxy what domain to use and certbot for generating the certificates
        For more info about the proxy configuration, you can see the (conf.d/https.conf.template)[]
 
    - The env variable `GF_SECURITY_ADMIN_PASSWORD` in docker-compose grafana service. Which is the grafana login password on monitoring.{DOMAIN}. The default login user is admin 


### Configuring your validator


For more info. See: 
https://docs.pokt.network/docs/create-validator-node#


#### Configure your chains.json file


Next step, configure your chains.json located in node1/chains.json serving your blockchains as follows:


```
[
    {
        "id": "0001",
        "url": "https://pocket-mainnet.yourdomain/",
    },
    {
        "id": "0021",
        "url": "https://eth-mainnet.yourdomain/",
    }
]
``` 

In case you have your URLS protected by basic_auth as shown in the README tutorial in pocket-supported-blockchains folder, you need:

```
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

For more information about chains.json file, see:

https://docs.pokt.network/changelog/chainsjs


#### Setting up pocket core node config.json 


Configure your pocket node by editing on `pocket/config.json` the following variables:
    -  "Moniker" - Use a custom Moniker name for your node, in this case can be `pocket-mainnet.${DOMAIN}`
    -  "ExternalAddress" - point to your node address and port, in this case `pocket-mainnet.${DOMAIN}:26656`


NOTE: in this case you need to manual replace ${DOMAIN} by your domain name


#### Obtaining your node_key.json and priv_val_key.json


If your have your `node_key.json` and `priv_val_key.json`, just move those files inside node1 and skip this step

In case you don't have the files mentioned.  Assuming you have your keyfile.json. Get your `node_key.json` and `priv_val_key.json` by doing:


```

> docker-compose up -d  # Run all services

> docker exec -it node1 sh # Enter your node1
 
> echo '${KEYFILE}' > keyfile.json # Or just copy and paste your keyfile content using nano or vim

> pocket accounts import-armored keyfile.json # Enter your passphrase

> pocket accounts set-validator {YOURADDRESS}
```


Now copy the content of `priv_val_key.json` and `node_key.json` from your node1 to node1 folder on your host


```
> cp /root/.pocket/priv_val_key.json  /home/app/.pocket 

> cp /root/.pocket/node_key.json  /home/app/.pocket 

```

#### Set your POCKET_CORE_PASSPHRASE for this node

Inside the `.env` file. Fill the env variable with your node passphrase (same used in the step while importing account before `pocket accounts import-armored keyfile.json` 


##### Sync/stake your node 


Restart your stack so it reflect the changes


```
docker-compose down && docker-compose up -d 
```


Wait until sync. validate if you are in the latest block and that your key is correcly configured with your node by doing 


```
> docker exec -it node1 sh

> pocket query node <youraddr> # It should show your addr and the domain of your node

> pocket query height # You should be in the latest height of the blockchain. See https://explorer.pokt.network/
```


While syncing, verify that your network proxy nginx and grafana stack is correctly configured by entering to the grafana login on:


> http://monitoring.${DOMAIN} 


Also, verify if your pocket node is correctly exposed by checking the pocket core ver on:


> http://node1.${DOMAIN}/v1


As last step, stake your node by doing the following commands inside your pocket node:


```
# Staking Command
pocket nodes stake  <fromAddr> <amount in uPOKT> <chains> <serviceURI w/ rpc port> <chainID> <fees in Upokt> 

# example with dummy values
pocket nodes stake 45D50DB64E90C0109C778DAAB7EF36676FC03866 1510000000 0001,0021 https://my-pocket-url:<port> mainnet 100000
``` 

For additional information you can see [Staking your node](https://docs.pokt.network/docs/create-validator-node#staking-your-node)


### References


- [FAQ for nodes](https://docs.pokt.network/docs/faq-for-nodes)

- [Create a pocket validator node](https://docs.pokt.network/docs/create-validator-node)


### Troubleshooting notes


#### My pocket node container is hang or doesn't stop/respond 

- It's very possible that when you stop your pocket container or while doing docker-compose down your container get's stuck stoppping. In this case you need to stop the daemon and run the stack again as follows:


```
> sudo systemctl restart docker && docker-compose down && docker-compose up -d
```

#### The cadvisor container is showing permission denied error


- This particular error is related to the volume permission and it can vary by OS, you can see a fix here https://github.com/google/cadvisor/issues/2387#issuecomment-600840479


#### Your node crashes with the err "priv_val_state.json: device or resource busy"

- For this case, be sure to have your `priv_val_state.json` empty and that you are not mapping that file in
