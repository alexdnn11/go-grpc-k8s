#!/usr/bin/env bash

echo "Rebuild the generated Go code"

protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pb