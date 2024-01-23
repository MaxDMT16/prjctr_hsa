# Sharding

## How to run

### Postrges without sharding or with vertical sharding
```bash
docker compose up -d
```

### Citus cluster ([doc](https://github.com/citusdata/docker/blob/master/README.md#docker-compose))
```bash
docker compose up -d --scale worker=4
```

To distribute data across the cluster run:
```SQL
SELECT create_distributed_table('books', 'category_id');
```

### Run app
Move to `app` folder. Set `APP_DSN` var connection string to the postgres DB and run `main.go`

Sample:
```bash
export APP_DSN='host=localhost user=prjctr password=test dbname=prjctr port=5433 sslmode=disable' && go run main.go
```

## Benchmarks
|Type|Operation|Rows count|Time (s)|
|---|---|---|----|
|no sharding|write|1000000|307.124|
|no sharding|read|1000000|3.392|
|vertical sharding|write|1000000|198.555|
|vertical sharding|read|1000000|2.819|
|horizontal sharding|write|1000000|453.919|
|horizontal sharding|read|1000000|7.182|
