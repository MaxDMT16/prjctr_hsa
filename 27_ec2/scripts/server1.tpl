#!/bin/bash
sudo apt update -y &&
sudo apt install -y nginx
echo "Hello from Server 1" > /var/www/html/index.html