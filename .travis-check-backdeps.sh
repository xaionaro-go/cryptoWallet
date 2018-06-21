#!/bin/bash -xe

go get github.com/rfjakob/gocryptfs
cd "$(go env GOPATH)/src/github.com/rfjakob/gocryptfs"
./crossbuild.bash
