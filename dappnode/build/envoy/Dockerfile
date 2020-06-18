FROM envoyproxy/envoy-alpine:v1.14.1
EXPOSE 8081

COPY ./envoy.yaml /etc/envoy/envoy.yaml
COPY ./certificate.sh ./
COPY ./entrypoint.sh ./
RUN apk add openssl