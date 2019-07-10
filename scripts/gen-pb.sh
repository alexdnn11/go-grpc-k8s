#!/usr/bin/env bash

echo "Rebuild the generated Go code for Twirp"

protoc --proto_path=$GOPATH/src:. --twirp_out=./ --go_out=./ ./pb/*.proto