to check

redis:
- rdb
- aof

beanstalkd:


throughput
ops/per sec


## Redis RDB
- read from queue
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

- write to queue
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
