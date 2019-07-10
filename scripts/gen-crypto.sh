#!/usr/bin/env bash

echo "Generate private key (.key)"
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
mkdir -p certs

openssl ecparam -genkey -name secp384r1 -out certs/server.key

echo "Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)"

openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.crt -days 3650

echo "Copy certs to server"
cp -rf certs server/certs

echo "Copy certs to client"
mkdir -p client/certs
cp -rf certs/server.crt client/certs/server.crt

