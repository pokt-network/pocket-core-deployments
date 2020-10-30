# SIMPLE POCKET STACK


## Overview

Contains a docker-compose.yaml with:

- Single pocket node validator
- Grafana/prometheus/loki/cadvisor/alertmanager for monitoring/metrics
- Reverse proxy and certbot auto certificate


## Instructions 


1. For running, you need to first run the following script for installing the loki driver for sending pocket logs to loki and grant file permissions needed for grafana and prometheus 

>  bash install.sh


2. Change the following settings according your needs:
    - The env variable `DOMAIN` and `EMAIL` in the file `.env` used in docker-compose services web,certbot services. 
    - The env variable `GF_SECURITY_ADMIN_PASSWORD` in docker-compose grafana service. 


3. Configure your pocket node by editing on `node1/config.json` the following variables:
    -  "Moniker"
    -  "ExternalAddress"

4. Open the ports HTTP/HTTPS and TCP/26656 in your server

5. Verify that your domains node1.${DOMAIN} and monitoring.${DOMAIN} are pointing to your server IP


6. Proceed to validate the ssl certificate by doing 

> docker-compose up web certbot 

NOTE: In case you are not able to get the certificate be sure to give more time to the IP change to propagate or best test adding `--staging` This will prevent you from getting timeout or ratelimit from using certbot. Once you get it working, remember to remove the `--staging` parameter and remove the test certificate found at proxy/certbot/conf/live/node1.${DOMAIN} 


7. After your certificate is issued you can uncomment the lines 21-78 on `proxy/conf.d/https.conf.templates`


8. Just do `docker-compose.yaml` and you will be up & running! 
