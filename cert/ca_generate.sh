#!/usr/bin/env bash

# CA private key
openssl genrsa -out ca.key 2048
# generate CA root credential request
openssl req -new -key ca.key -out ca.csr
# generate CA root credential
openssl x509 -req -in ca.csr -signkey ca.key -days 365 -out ca.crt