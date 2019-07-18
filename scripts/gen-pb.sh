#!/usr/bin/env bash

echo "Rebuild the generated Go code"

protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pb

echo "Rebuild the generated Node code"
cd pb
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:../client-node --grpc_out=../client-node --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` gcd.proto
