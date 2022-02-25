#!/usr/bin/env bash

if [ -z "$1" ]
then
  echo "Please supply a subdomain to create a certificate for";
  echo "e.g. www.site.com"
  exit;
fi

if [ ! -f ca.crt ]; then
  echo 'Please run "ca_generate.sh" first, and try again!'
  exit;
fi
if [ ! -f v3.ext ]; then
  echo 'Please download the "v3.ext" file and try again!'
  exit;
fi

# Create a new private key if one doesnt exist, or use the existing one if it does
if [ -f $DOMAIN.key ]; then
  KEY_OPT="-key"
else
  KEY_OPT="-keyout"
fi

DOMAIN=$1
COMMON_NAME=${2:-*.$1}
SUBJECT="/C=CA/ST=None/L=NB/O=None/CN=$COMMON_NAME"
NUM_OF_DAYS=365
openssl req -new -newkey rsa:2048 -sha256 -nodes $KEY_OPT $DOMAIN.key -subj "$SUBJECT" -out $DOMAIN.csr
cat v3.ext | sed s/%%DOMAIN%%/"$COMMON_NAME"/g > /tmp/__v3.ext
openssl x509 -req -in $DOMAIN.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out $DOMAIN.crt -days $NUM_OF_DAYS -sha256 -extfile /tmp/__v3.ext

echo Done!
