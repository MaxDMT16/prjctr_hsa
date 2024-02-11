#!/bin/bash
sudo apt update -y &&
sudo apt install -y nginx
echo "Hello from Server 1" > /var/www/html/index.html

echo "server {
    listen 8080;

    location / {
        root /var/www/html;
        index index.html index.htm;
    }

}" >> /etc/nginx/conf.d/default.conf

sudo systemctl restart nginx
