# POCKET-BLOCKCHAIN STACK 


## Overview

Contains a docker-compose.yaml with:

- Single pocket node 
- Grafana/prometheus/loki/cadvisor/alertmanager for monitoring/metrics
- Reverse proxy and certbot auto certificate

### Hardware requirements

CPU: 4 CPU’s (or vCPU’s)
Memory: 8 GB RAM
Disk: 100GB is the minimum HDD requirement

For more information, see: 
https://docs.pokt.network/docs/before-you-dive-in#hardware-requirements


## Instructions 


1. For running, you need to first run the following script for installing the loki driver for sending pocket logs to loki and grant file permissions needed for grafana and prometheus 

>  bash install.sh
>  sudo bash install.sh


2. Change the following settings according your needs:
    - The env variable `DOMAIN` and `EMAIL` in the file `.env` used in docker-compose services web,certbot services. 
    - The env variable `GF_SECURITY_ADMIN_PASSWORD` in docker-compose grafana service. 


3. Configure your pocket node by editing on `node1/config.json` the following variables:
    -  "Moniker"
    -  "ExternalAddress"


4. Open the ports HTTP/HTTPS and TCP/26656 in your server


5. Create and Verify that your domains node1.${DOMAIN} and monitoring.${DOMAIN} are pointing to your server IP. You can verify this using:


> nslookup node1.${DOMAIN}

> nslookup monitoring.${DOMAIN}


6. Proceed to validate the ssl certificate by doing 


> docker-compose up web certbot 

NOTE: In case you are not able to get the certificate be sure to give more time to the IP change to propagate or best test adding `--staging` This will prevent you from getting timeout or ratelimit from using certbot. Once you get it working, remember to remove the `--staging` parameter and remove the test certificate found at proxy/certbot/conf/live/node1.${DOMAIN}. The command looks like:

>  certonly --webroot --webroot-path=/var/www/certbot --email ${EMAIL} --agree-tos --no-eff-email --staging -d node1.${DOMAIN} -d monitoring.${DOMAIN}


7. After your certificate is issued uncomment all the lines with '#' on `proxy/conf.d/https.conf.templates`


8. Just do `docker-compose down && docker-compose up -d ` and you will be up & running! 


### References

- [Create a pocket node](https://docs.pokt.network/docs/create-pocket-node)


### Troubleshooting notes



#### My pocket node container is hang or doesn't stop/respond 

- It's very possible that when you stop your pocket container or while doing docker-compose down your container get's stuck stoppping. In this case you need to stop the daemon and run the stack again as follows:

> sudo systemctl restart docker && docker-compoes down && docker-compose up -d


#### The cadvisor container is showing permission denied error

- This particular error is related to the volume permission and it can vary by OS, you can see a fix here https://github.com/google/cadvisor/issues/2387#issuecomment-600840479
