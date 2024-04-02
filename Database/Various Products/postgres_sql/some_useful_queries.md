Sure, here are some of the top useful queries for PostgreSQL:

* **Show running queries:** This query will show you all of the queries that are currently running in your database. This can be useful for troubleshooting performance problems.
```sql
SELECT procpid, age(clock_timestamp(), query_start), usename, current_query
FROM pg_stat_activity
WHERE current_query != '<IDLE>' AND current_query NOT ILIKE '%pg_stat_activity%'
ORDER BY query_start DESC;
```

* **Kill running queries:** This query will kill a specific running query. This can be useful if a query is causing a performance problem.
```sql
SELECT pg_cancel_backend(pid);
```

* **Vacuum:** This query will remove unused data from your database. This can improve the performance of your database.
```sql
VACUUM (VERBOSE, ANALYZE);
```

* **Show all databases:** This query will show you all of the databases that exist in your PostgreSQL installation.
```sql
SELECT datname
FROM pg_database;
```

* **Show all tables:** This query will show you all of the tables that exist in a specific database.
```sql
SELECT tablename
FROM pg_tables
WHERE schemaname = 'public';
```

* **Show all columns:** This query will show you all of the columns that exist in a specific table.
```sql
SELECT columnname
FROM pg_columns
WHERE tablename = 'customers';
```

* **Show all indexes:** This query will show you all of the indexes that exist in a specific table.
```sql
SELECT indexname
FROM pg_indexes
WHERE tablename = 'customers';
```

* **Show all database connections:** This query will show you all of the active connections to your database. This can be useful for troubleshooting performance problems or identifying connections that are using a lot of resources.

```sql
SELECT * FROM pg_stat_activity;
```

* **Show the size of all tables:** This query will show you the size of all of the tables in your database. This can be useful for managing your database storage and troubleshooting performance problems.

```sql
SELECT tablename, pg_size_pretty(pg_table_size(tablename))
FROM pg_tables;
```

* **Show the most active tables:** This query will show you the tables that are being accessed the most often. This can be useful for identifying tables that are causing performance problems or that need to be indexed.

```sql
SELECT tablename, n_live_tup rows_in_table
FROM pg_stat_user_tables
ORDER BY n_live_tup DESC;
```

* **Show the most active indexes:** This query will show you the indexes that are being used the most often. This can be useful for identifying indexes that are not being used and that can be dropped.

```sql
SELECT indexname, n_tup_live rows_in_index
FROM pg_stat_user_indexes
ORDER BY n_tup_live DESC;
```

* **Show the most expensive queries:** This query will show you the queries that are taking the longest to execute. This can be useful for identifying queries that need to be optimized.

```sql
SELECT query, total_time
FROM pg_stat_statements
ORDER BY total_time DESC;
```



Sources:
stackoverflow.com/questions/62639240/flyway-stuck-when-cleaning-database-postgres-11
github.com/asishraz/Code-Godown