#!/bin/sh
case ${TLS_OPTION} in
"I have a certificate")
    cp /tmp/certificate.crt /tmp/key.key /etc/envoy/
    /home/app/command.sh & envoy -c /etc/envoy/envoy.yaml --component-log-level upstream:error,connection:error
    ;;
"Already uploaded certificate")
    /home/app/command.sh & envoy -c /etc/envoy/envoy.yaml --component-log-level upstream:error,connection:error
    ;;
"Do not use a certificate")
    /home/app/command.sh
    ;;
esac
