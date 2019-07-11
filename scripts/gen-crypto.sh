#!/usr/bin/env bash

echo "Generate private key (.key)"
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
mkdir -p certs
mkdir -p certs/server
mkdir -p certs/client

openssl ecparam -genkey -name secp384r1 -out certs/server/server.key
openssl ecparam -genkey -name secp384r1 -out certs/client/server.key

echo "Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)"

openssl req -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=example.com" -new -x509 -sha256 -key certs/server/server.key -out certs/server/server.crt -days 3650
openssl req -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=example.com" -new -x509 -sha256 -key certs/client/server.key -out certs/client/server.crt -days 3650

