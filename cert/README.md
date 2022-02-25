
prepare CA root crt

openssl genrsa -out ca.key 1024
openssl req -new -key ca.key -out ca.csr
###### self signed cert
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt


prepare company crt
1. Generate server private key  
   openssl genrsa -out server.key 1024
2. Generate crt request file  
   openssl req -new -key server.key -out server.csr
3. Use ca.crt ca.key, server.csr to generate server's crt  
   openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt

example:
    # create a root authority cert
    ./create_root_cert_and_key.sh
    
    # create a wildcard cert for mysite.com
    ./create_certificate_for_domain.sh mysite.com
    
    # or create a cert for www.mysite.com, no wildcards
    ./create_certificate_for_domain.sh www.mysite.com www.mysite.com

