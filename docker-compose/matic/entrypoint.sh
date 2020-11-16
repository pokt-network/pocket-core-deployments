#!/bin/bash

set -e
set -m

service rabbitmq-server start & 
heimdalld start & 
heimdalld rest-server start &
heimdalld bridge start &
cd /root/node/bor/

bash "$@"
