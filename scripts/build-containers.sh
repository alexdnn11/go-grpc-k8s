#!/usr/bin/env bash

echo "Building api"

docker build -t local/api:latest -f Dockerfile.api .

echo "Building gcd"

docker build -t local/gcd:latest -f Dockerfile.gcd .

echo "Building node"

docker build -t local/api-node:latest -f Dockerfile.node .
