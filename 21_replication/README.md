Create 3 docker containers: mysql-m, mysql-s1, mysql-s2
Setup master slave replication (Master: mysql-m, Slave: mysql-s1, mysql-s2)
Write script that will frequently write data to database
Ensure, that replication is working
Try to turn off mysql-s1 (stop slave), 
Try to remove a column in  database on slave node (try to delete last column and column from the middle)
Write conclusion in readme.md


# DB replication

## How to run

```bash
./bash.sh
```

If `bash.sh` is not executable, run `sudo chmod u+x build.sh` first.

## HA questions
1. Try to turn off mysql-s1 (stop slave)

After stopping slave on one of the slave-nodes, sync with master doesn't work - new data from master doesn't appear on slave node.
When slave is started again, it syncs with master - load diff between master and slave and new data appears on slave node.

2. Try to remove a column in  database on slave node

After removing the last column from table on slave node, it is present on master and another slave node. When a new data appears on master, it is replicated to slave node without deleted column. And `SHOW SLAVE STATUS` shows empty value for `Last_Error` column.

When a column from the middle of table is removed on slave node, it is present on master and another slave node. When a new data appears on master, it is *not* replicated to slave node. And `SHOW SLAVE STATUS` shows appropriate error.
E.g. 
```
Column 2 of table 'mydb.users' cannot be converted from type 'int' to type 'varchar(100(bytes) latin1)'
```
