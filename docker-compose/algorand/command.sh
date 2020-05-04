#!/usr/bin/bash
export TERM=linux
./goal node start -d data
watch -n 2 ./goal node status -d data
