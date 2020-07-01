#!/bin/sh
case ${TLS_OPTION} in
"I have a certificate")
    if test -f "/etc/envoy/certificate.crt" && test -f "/etc/envoy/key.key"; then
        /home/app/entrypoint.sh & envoy -c /etc/envoy/envoy.yaml --component-log-level upstream:debug,connection:trace
    else
        cp /tmp/certificate.crt /tmp/key.key /etc/envoy/
        /home/app/entrypoint.sh & envoy -c /etc/envoy/envoy.yaml --component-log-level upstream:debug,connection:trace
    fi
    ;;

"Do not use a certificate")
    /home/app/entrypoint.sh
    ;;
"Provide a selfsigned certificate for me")

    if test -f "/etc/envoy/certificate.crt" && test -f "/etc/envoy/key.key"; then
        /home/app/entrypoint.sh
    else
        sh certificate.sh
        wait 5
        /home/app/entrypoint.sh & envoy -c /etc/envoy/envoy.yaml --component-log-level upstream:debug,connection:trace
    fi
    ;;
esac
