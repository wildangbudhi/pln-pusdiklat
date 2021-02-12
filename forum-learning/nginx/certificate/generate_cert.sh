#!/bin/bash

openssl req -nodes -config openssl.conf -days 356 -x509 -newkey rsa:4096 -out /etc/ssl/certs/nginx-selfsigned.crt
