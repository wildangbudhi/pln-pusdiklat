#!/bin/sh

envsubst < /etc/nginx/templates/prod.conf.template > /etc/nginx/sites-available/default
rm /etc/nginx/sites-enabled/default
ln -s /etc/nginx/sites-available/default /etc/nginx/sites-enabled/

nginx -g ndaemon off;