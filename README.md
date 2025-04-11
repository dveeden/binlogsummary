# Usage

```
$ ./binlogsummary /var/lib/mysql/binlog.0*
/var/lib/mysql/binlog.000001	896e7882-18fe-11ef-ab88-22222d34d411:1-2
/var/lib/mysql/binlog.000002	896e7882-18fe-11ef-ab88-22222d34d411:3-141
/var/lib/mysql/binlog.000003	<empty>
/var/lib/mysql/binlog.000004	<empty>
/var/lib/mysql/binlog.000005	<empty>
/var/lib/mysql/binlog.000006	896e7882-18fe-11ef-ab88-22222d34d411:142-11584
/var/lib/mysql/binlog.000007	896e7882-18fe-11ef-ab88-22222d34d411:11585-11758,f90f239e-16b2-11f0-a55e-b68cb2e1b8f5:1
```

# TODO

- Test with compressed binlogs
- Test with MariaDB
- Test with tagged GTIDs
- Test with anonymous GTID
