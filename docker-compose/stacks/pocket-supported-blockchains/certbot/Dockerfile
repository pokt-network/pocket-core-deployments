FROM alpine:latest

# installs cron
RUN apk add --update certbot apk-cron docker openrc && rc-update add docker boot && rm -rf /var/cache/apk/*
# installs docker for restarting proxy
RUN mkdir -p /home/app
COPY entrypoint.sh /opt/certbot/entrypoint.sh
COPY certbot-renew /opt/certbot/certbot-renew
RUN chmod +x /opt/certbot/entrypoint.sh


ENTRYPOINT ["/opt/certbot/entrypoint.sh"]
