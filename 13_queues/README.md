## Comparison of Redis and Beanstalkd

| operation |   metric   | redis rdb | redis aof | beanstalkd |
|-----------|------------|-----------|-----------|------------|
| read      |  ops/sec   | 1000      | 979       | 396        |
| write     |  ops/sec   | 110       | 100       | 101        |

## Siege results

### Redis RDB
- read from a queue
``` bash
Transactions:                   2820 hits
Availability:                 100.00 %
Elapsed time:                   2.82 secs
Data transferred:               0.07 MB
Response time:                  0.02 secs
Transaction rate:            1000.00 trans/sec
Throughput:                     0.02 MB/sec
Concurrency:                   16.17
Successful transactions:        2820
Failed transactions:               0
Longest transaction:            0.17
Shortest transaction:           0.00
```

- write to a queue
``` bash
Transactions:                    314 hits
Availability:                 100.00 %
Elapsed time:                   2.83 secs
Data transferred:               0.00 MB
Response time:                  0.01 secs
Transaction rate:             110.95 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                    0.64
Successful transactions:         314
Failed transactions:               0
Longest transaction:            0.02
Shortest transaction:           0.00
```

### Redis AOF
- read from a queue
``` bash
Transactions:                   2273 hits
Availability:                 100.00 %
Elapsed time:                   2.32 secs
Data transferred:               0.04 MB
Response time:                  0.02 secs
Transaction rate:             979.74 trans/sec
Throughput:                     0.02 MB/sec
Concurrency:                   17.96
Successful transactions:        2273
Failed transactions:               0
Longest transaction:            0.16
Shortest transaction:           0.00
```

- write to a queue
``` bash
Transactions:                    239 hits
Availability:                 100.00 %
Elapsed time:                   2.39 secs
Data transferred:               0.00 MB
Response time:                  0.00 secs
Transaction rate:             100.00 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                    0.45
Successful transactions:         239
Failed transactions:               0
Longest transaction:            0.01
Shortest transaction:           0.00
```

### Beanstalkd

- read from a queue
``` bash
Transactions:                    991 hits
Availability:                  48.96 %
Elapsed time:                   2.50 secs
Data transferred:               0.03 MB
Response time:                  0.04 secs
Transaction rate:             396.40 trans/sec
Throughput:                     0.01 MB/sec
Concurrency:                   15.18
Successful transactions:         991
Failed transactions:            1033
Longest transaction:            0.16
```

- write to a queue
``` bash
Transactions:                    244 hits
Availability:                 100.00 %
Elapsed time:                   2.41 secs
Data transferred:               0.00 MB
Response time:                  0.01 secs
Transaction rate:             101.24 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                    0.77
Successful transactions:         244
Failed transactions:               0
Longest transaction:            0.03
Shortest transaction:           0.00
```
## How to run
- Set up services
``` bash    
docker-compose up -d
```

- To run tests use scripts in `scripts.sh` file


⚠️ there are 2 redis services in the compose file - RDB & AOF
To select what redis to use, change `REDIS_ADDRESS` in `config.env` file.
For rdb use `REDIS_ADDRESS=redis-rdb:6379` and for aof use `REDIS_ADDRESSredis-aof:6379`