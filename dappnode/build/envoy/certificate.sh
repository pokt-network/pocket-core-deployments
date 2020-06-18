#!/bin/bash

commonname=${COMMON_NAME}
country=${COUNTRY}
state=${STATE}
locality=${LOCALITY}
organization=${ORGANIZATION}
organizationalunit=${ORGANIZATIONAL_UNIT}
email=${EMAIL}


ipVerification=`echo $commonname | awk -F"\." ' $0 ~ /^([0-9]{1,3}\.){3}[0-9]{1,3}$/ && $1 <=255 && $2 <= 255 && $3 <= 255 && $4 <= 255 '`

if [ -n "$ipVerification" ] 
then
   echo "IP"
    openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout /etc/envoy/key.key -out /etc/envoy/certificate.crt -addext subjectAltName=IP:${COMMON_NAME} -subj "/C=$country/ST=$state/L=$locality/O=$organization/OU=$organizationalunit/CN=$commonname/emailAddress=$email"
else 
   echo "NO IP"
    openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout /etc/envoy/key.key -out /etc/envoy/certificate.crt -subj "/C=$country/ST=$state/L=$locality/O=$organization/OU=$organizationalunit/CN=$commonname/emailAddress=$email"
fi



exit