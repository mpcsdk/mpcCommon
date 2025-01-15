##

```
mysqldump -h 127.0.0.1 -P 3306 -u root -p123456 relayeradmin > relayeradmin_dump.sql
```

```
mysql -h localhost -P 3306 -u root -p < ./mpcdao/migratioin/relayeradmin/relayeradmin_dump.sql
```