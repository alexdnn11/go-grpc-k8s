#!/usr/bin/env bash

echo "Building api"

docker build -t local/api:latest -f Dockerfile.api .

echo "Building gcd"

docker build -t local/ms-s:dev -f Dockerfile.ms .

echo "Building node"

docker build -t local/api-node:latest -f Dockerfile.node .
