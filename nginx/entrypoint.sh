#!/bin/sh

envsubst < /etc/nginx/conf.d/nginx.conf.template > /etc/nginx/conf.d/nginx.conf

nginx -g 'daemon off;'
