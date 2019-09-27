FROM poktnetwork/pocket-core:mvp-1.7 AS binary

FROM alpine:3.10

WORKDIR /usr/src/app
RUN apk add --update bash ca-certificates curl

COPY --from=binary /bin/pocket-core /usr/local/bin/

# Create app user and add permissions
RUN addgroup -S app \
	&& adduser -S -G app app

ADD default_chains.json /usr/src/app/chains.json
ADD default_service_whitelist.json /usr/src/app/service_whitelist.json
RUN chown -R app:app /usr/src/app
USER app

ADD start_poket-core.sh /usr/local/bin/start_poket-core

ENTRYPOINT [ "start_poket-core" ]