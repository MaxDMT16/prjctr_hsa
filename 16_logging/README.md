# Logging

Created MySQL service with turned on slow query logs.
Configured filebeat input to listen to the MySQL logs and output the to graylog.

MySQL => Filebeat => Graylog

## How to test

1. Run services
```bash
docker-compose up -d
```

2. Setup input in Graylog:
    - System -> Inputs -> Beats -> Launch new input
    - Set port to 5044
    - Save

3. Run slow query in MySQL. Test query:
```sql
DO SLEEP(2);
```

4. Check logs in Graylog on the Search page.