#!/bin/bash -xe

go get github.com/rfjakob/gocryptfs
cd "$(go env GOPATH)/src/github.com/rfjakob/gocryptfs"
./crossbuild.bash

go get github.com/xaionaro-go/trezorLuks
cd "$(go env GOPATH)/src/github.com/xaionaro-go/trezorLuks"
go build
