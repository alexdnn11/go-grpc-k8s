#!/usr/bin/env bash

echo "### Make folders ###"

: ${DOMAIN:="example.com"}
: ${SERVER:="ms-s"}
: ${CLIENT:="api"}
: ${CA:="ca"}

mkdir -p certs
mkdir -p certs/server
mkdir -p certs/client
mkdir -p certs/ca

echo "### Make CA ###"

echo "### RSA key for CA ###"
openssl genrsa -out certs/ca/rootCA.key 2048

echo "### Certificate for CA ###"
openssl req -x509 -new -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=$CA.$DOMAIN" -key certs/ca/rootCA.key -days 10000 -out certs/ca/rootCA.crt


echo "### Generate a 2048 bit RSA key ###"

echo "### RSA key for server ###"
openssl genrsa -out certs/server/server.key 2048

echo "### RSA key for client ###"
openssl genrsa -out certs/client/client.key 2048

echo "### Generate a certificate signing request (.csr) using openssl ###"

echo "### CSR for server ###"
openssl req -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=$SERVER.$DOMAIN" -new -sha256 -key certs/server/server.key -out certs/server/server.csr

echo "### CSR for client ###"
openssl req -subj "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=$CLIENT.$DOMAIN" -new -sha256 -key certs/client/client.key -out certs/client/client.csr

echo "### Sign CRT ###"

echo "### Sign CRT for server ###"
openssl x509 -req -in certs/server/server.csr -CA certs/ca/rootCA.crt -CAkey certs/ca/rootCA.key -CAcreateserial -out certs/server/server.crt -days 5000
echo "### Sign CRT for client ###"
openssl x509 -req -in certs/client/client.csr -CA certs/ca/rootCA.crt -CAkey certs/ca/rootCA.key -CAcreateserial -out certs/client/client.crt -days 5000