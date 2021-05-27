# POCKET SUPPORTED BLOCKCHAINS STACK 


## Overview

Contains a docker-compose.yaml with:

- Single pocket blockchain node
- GETH Ethereum mainnet full node 
- Grafana/prometheus/loki/cadvisor/alertmanager for monitoring/metrics
- Reverse proxy and certbot auto certificate
- Basic auth for protecting URLs 


The purpose is to provide a closest production-ready stack with all the blockchains that pocket has support in order to serve traffic for your pocket-validator 

We packed basics of security and other configurations. Security and server hardening it's up to the responsability of all node runners 


## Requirements

- Your own domain
- docker (19.03.0+) and docker-compose(1.27.4+) installed on the system 

## Hardware requirements


### Pocket node

- CPU: 2 CPUs
- Memory: 4 GB RAM
- Disk: 60GB SSD 

For more information, see [Pocket node hardware requirements](https://docs.pokt.network/docs/before-you-dive-in#hardware-requirements)


### Ethereum full node  

- CPU:  4 or more cores
- RAM: 16 GB or more
- DISK: SSD with at least 500GB 
- 25+ MBit/s bandwidth



### All blockchains mentioned


In case you are planning to serve all mentioned blockchain nodes along with the monitoring/network stack:

- CPU:  8 
- RAM: 32 GB
- DISK: SSD with at least 640GB 


NOTE: We strongly suggest that you separate the disks for every blockchain if your infrastructure services lets you to do it. Right now, with pocket-blockchain and geth-mainnet you can work it out, but if you integrate more blockchain nodes, is better to use a separate disks for them to avoid IOPs overload 


## Customizing your configuration


This all-in-one docker-compose stack is created with customization in mind. In case you want to remove any blockchain, you only need to remove it from the service from the `docker-compose.yaml`, remove his route from `proxy/conf.d/https.conf.template` and skip their install instructions from the quick tutorial below.

For ex. if you only want to serve ethereum full node and remove pocket you only need to remove the service `pocket` from the `docker-compose.yaml` file, remove the route to pocket in `proxy/conf.d/https.conf.template` and skip the install instructions related to the pocket-node 


## Instructions 


### Network configuration


It's convenient for extra security that this server resides in a private network 

In case you are running in public network. Additionally to this security layer, this setup includes instructions and settings for basic auth and port allowance to only the IP of the pocket-validator stack


#### Port configuration 


Open the ports:

* SSH 	 				
    * Only accessed by your IP  (x.x.x.x/32)
* HTTP/HTTPS for the nginx proxy  	
	* Public access (0.0.0.0/0) needed for the grafana dashboard  
* TCP/26656 for pocket peers
	* Public access (0.0.0.0/0)
* TCP/30303 for ethereum peers  	
	* Public access (0.0.0.0/0)


Verify to block all the other ports and harden the SSH access to only connect using your machine IP and a keypair 


####  Create domain records 

Create the following domain A records pointing to your pocket-blockhains server IP:

- pocket-mainnet.${DOMAIN}
- eth-mainnet.${DOMAIN}
- monitoring.${DOMAIN} 


After you finish, wait 5-10 mins or the time required given by your DNS until the spread is over. 


You can verify if domain are correctly configured by checking with nslookup that the domains return your IP: 

```
 nslookup pocket-mainnet.${DOMAIN}

 nslookup eth-mainnet.${DOMAIN}

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

####  Create basic auth for proxy access


Basic auth will provide us a http security layer for providing access to only the pocket node validators using a user/password to connect to our blockchain nodes. 

```
htpasswd -b proxy/conf.d/.htpasswd youruser mypassword
```

Then, in your chains.json of your pocket-validator point to your pocket blockchain and geth-mainnet url with this user/password as follows:

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

generate more information about chains.json file, see:

https://docs.pokt.network/changelog/chainsjs


#### Uncomment proxy routes

After your certbot certificate is issued, you can stop uncomment all the '#' on the file `proxy/conf.d/https.conf.templates`


### Setting up proxy and monitoring systems 


We use nginx for the web proxy and (loki/grafana/prometheus) stack for the monitoring systems

You can see their configurations in the docker-compose volume mappings and their respective folders
 

#### Configure loki driver and grafana/prometheus permissions 


The following script will install the loki driver for sending blockchain node logs to loki and grant file permissions needed for grafana and prometheus 

```
sudo bash install.sh
```

#### Set your domain and configure grafana access 
 

Change the following settings according your setup:
    - The env variable `DOMAIN` and `EMAIL` in the file `.env` used in docker-compose web, certbot services. Which indicates to nginx proxy what domain to use and certbot for generating the certificates
        For more info about the proxy configuration, you can see the (conf.d/https.conf.template)[]
 
    - The env variable `GF_SECURITY_ADMIN_PASSWORD` in docker-compose grafana service. Which is the grafana login password. The default login user is admin 


### Customizing blockchains configuration 


#### Setting up pocket core node


Configure your pocket node by editing on `pocket/config.json` the following variables:
    -  "Moniker" - Use a custom Moniker name for your node, in this case can be `pocket-mainnet.${DOMAIN}`
    -  "ExternalAddress" - point to your node address and port, in this case `pocket-mainnet.${DOMAIN}:26656`


NOTE: in this case you need to manual replace ${DOMAIN} by your domain name


#### Running


Just do `docker-compose down && docker-compose up -d ` and you will be up & running! 



### References


- [Create a pocket node](https://docs.pokt.network/docs/create-pocket-node)


### Troubleshooting notes



#### My pocket node container is hang or doesn't stop/respond 


- It's very possible that when you stop your pocket container or while doing docker-compose down your container get's stuck stoppping. In this case you need to stop the daemon and run the stack again as follows:


```
 sudo systemctl restart docker && docker-compose down && docker-compose up -d
```


#### The cadvisor container is showing permission denied error

- This particular error is related to the volume permission and it can vary by OS, you can see a fix here https://github.com/google/cadvisor/issues/2387#issuecomment-600840479

