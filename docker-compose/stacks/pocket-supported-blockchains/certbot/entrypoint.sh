#!/bin/sh
crontab  /opt/certbot/certbot-renew
crontab -l
crond -f 
