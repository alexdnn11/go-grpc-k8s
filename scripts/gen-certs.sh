#!/usr/bin/env bash

echo "Generate private key (.key)"
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
mkdir -p certs
mkdir -p certs/rsa
mkdir -p certs/rsa/server
mkdir -p certs/rsa/client

echo "### Generate a 2048 bit RSA key ###"
echo "### RSA key for server ###"
openssl genrsa -out certs/rsa/server/server.key 2048
echo "### RSA key for client ###"
openssl genrsa -out certs/rsa/client/server.key 2048


echo "### Generate the certificate ###"
echo "### Certificate for server ###"
openssl req -new -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=gcd.example.com" -x509 -sha256 -key certs/rsa/server/server.key \
              -out certs/rsa/server/server.crt -days 3650
echo "### Certificate for client ###"
openssl req -new -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=api.example.com" -x509 -sha256 -key certs/rsa/client/server.key \
              -out certs/rsa/client/server.crt -days 3650

echo "### Generate a certificate signing request (.csr) using openssl ###"
echo "### CSR for server ###"
openssl req -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=gcd.example.com" -new -sha256 -key certs/rsa/server/server.key -out certs/rsa/server/server.csr
echo "### CSR for client ###"
openssl req -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=api.example.com" -new -sha256 -key certs/rsa/client/server.key -out certs/rsa/client/server.csr

echo "### Sign CRT for server ###"
openssl x509 -req -sha256 -in certs/rsa/server/server.csr -signkey certs/rsa/server/server.key \
               -out certs/rsa/server/server.crt -days 3650
echo "### Sign CRT for client ###"
openssl x509 -req -sha256 -in certs/rsa/client/server.csr -signkey certs/rsa/client/server.key \
               -out certs/rsa/client/server.crt -days 3650


