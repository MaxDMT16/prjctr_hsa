## Results

- rows: 8353809
- create B-tree indext: 17.203 s
- crate hash index: 20.529 s

### Reading
|index|action|time|
|---|---|---|
|no index|read|486 ms|
|B-tree index|read|0.335 ms|
|hash index|read|3.667 ms|

### Writing

An analog of `innodb_flush_log_at_trx_commit` in MySQL is `fsync` in PostgreSQL (see [here](https://www.postgresql.org/docs/9.6/static/runtime-config-wal.html#GUC-FSYNC)).


|fsync|ops/sec|time (s)|inserted_rows|
|---|---|---|-----|
|on|307.91|5.69|1752|
|on|399.66|5.06|2022|
|off|352.18|5.27|1856|
|off|452.89|5.88|2663|

