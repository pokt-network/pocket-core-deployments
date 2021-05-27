#!/bin/bash
go get -u github.com/goware/modvendor
go get github.com/pokt-network/posmint@ea06d11007f9081929553c50153e9d5
go mod vendor
go mod download
modvendor -copy="**/*.c **/*.h" -v
