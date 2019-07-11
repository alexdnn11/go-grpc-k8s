#!/usr/bin/env bash

echo "Generate private key (.key)"
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
mkdir -p certs
mkdir -p certs/rsa

echo "### Generate a 2048 bit RSA key ###"
openssl genrsa -out certs/rsa/server.key 2048


echo "### Generate the certificate ###"
openssl req -new -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=gcd.example.com" -x509 -sha256 -key certs/rsa/server.key \
              -out certs/rsa/server.crt -days 3650

echo "### Generate a certificate signing request (.csr) using openssl ###"
openssl req -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=gcd.example.com" -new -sha256 -key certs/rsa/server.key -out certs/rsa/server.csr
openssl x509 -req -sha256 -in certs/rsa/server.csr -signkey certs/rsa/server.key \
               -out certs/rsa/server.crt -days 3650


