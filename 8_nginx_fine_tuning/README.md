# NGINX fine tuning

## How to run
```
docker-compose up -d
```

## Description
This solution contains a simple backend app that serves static files (`backend` folder) and NGINX server that proxies requests to the backend.
NGINX caches images that were requested more than twice. It was expected to provide ability to purge cache item, but it is not available in a free version of NGINX. As a workaround I've added ability to ommmit cache by adding `?cache=false` to the request.