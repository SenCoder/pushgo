#!/bin/sh

# Generate self signed root CA cert
openssl req -nodes -x509 -newkey rsa:2048 -keyout ca.key -out ca.crt -subj "/C=ZH/ST=Guangdong/L=Shenzhen/O=PaTech/OU=root/CN=localhost/emailAddress=peter_office@163.com"


# Generate server cert to be signed
openssl req -nodes -newkey rsa:2048 -keyout server.key -out server.csr -subj "/C=ZH/ST=Guangdong/L=Shenzhen/O=PaTech/OU=server/CN=localhost/emailAddress=peter_office@163.com"

# Sign the server cert
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt

# Combine server.crt + ca.crt
cat server.crt ca.crt > server1.crt


# Generate client cert to be signed
openssl req -nodes -newkey rsa:2048 -keyout client.key -out client.csr -subj "/C=ZH/ST=Guangdong/L=Shenzhen/O=PaTech/OU=client/CN=localhost/emailAddress=peter_office@163.com"

# Sign the client cert
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAserial ca.srl -out client.crt

# Combine client.crt + ca.crt
cat client.crt ca.crt > client1.crt