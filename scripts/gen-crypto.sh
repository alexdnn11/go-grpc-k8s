#!/usr/bin/env bash

echo "Generate server private key (.key)"
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
cd server

openssl ecparam -genkey -name secp384r1 -out server.key

echo "Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key) for server"

openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

echo "Generate client private key (.key)"
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
cd ../client

echo "Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key) for client"

openssl ecparam -genkey -name secp384r1 -out client.key

#Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
openssl req -new -x509 -sha256 -key client.key -out client.crt -days 3650
