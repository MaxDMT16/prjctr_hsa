## Desscription
A solution with a wrapper for Redis Client that implement probabilistic cache clearing

## How to run
```bash
docker-compose up
```

## How to test
- open http://localhost:8989/data/key2
- in the logs of `app` service you will see that that key was not found. So, the value was taken from in memory DB
- open http://localhost:8989/data/key2 again
- in the logs of `app` service you will see that that key was found. So, the value was taken from Redis
- if query http://localhost:8989/data/key2 again after ~90 secods since the first request, in `app` service you will see that the TTL of the `key2` was refreshred in advance (10s till expiration or less)