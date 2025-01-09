##

```
mysqldump -h 127.0.0.1 -P 3306 -u riskcontrol -p123456 riskcontrol > riskcontrol_dump.sql
```

```
pg_dump --schema-only -U postgres -h localhost -p 10100 -d tfa -f tfa_dump.sql
```
