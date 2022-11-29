# PostgreSQL


## Environment
- macOS 13.0.1
- postgres (PostgreSQL) 14.3


## Get Started with psql

### Check your postgres version

```bash
postgres -V
```

### psql Config

For brew installed psql, use the following command to get the config file path.
```bash
echo "$(brew --prefix)/var/postgres/pg_hba.conf"
```

Alternative, user the following command
```bash
psql -d postgres -U username -qAt -c "show hba_file"
```

In addition, you can access and filter the config file content with **grep**
```bash
psql -d postgres -U username -qAt -c "show hba_file" | xargs grep -v -E '^[[:space:]]*#'
```

Restart the service after updating the config
```bash
brew services restart postgresql
```

### Connect to the Terminal Utility psql

Start **psql** with the automatically created database **"postgres"** and your current user
```bash
psql postgres
```

Available **psql** connection arguments
```text
-d, --dbname=DBNAME      database name to connect to
-h, --host=HOSTNAME      database server host or socket directory (default: "local socket")
-p, --port=PORT          database server port (default: "5432")
-U, --username=USERNAME  database user name (default: "current user")
```

### psql Connection Command Cheat Sheet

```text
\l[+]   [PATTERN]      list databases
\du[S+] [PATTERN]      list roles
\dt[S+] [PATTERN]      list tables
\d[S+]  NAME           describe table, view, sequence, or index
```

### CREATE ROLE with psql

Create a new user `username` and password `new_password`
```sql
CREATE ROLE username WITH PASSWORD 'new_password';
```

Alternative, you can update the password of user `username` with psql command.
```
\password username;
```

Update user `username` with different role attributes
```sql
ALTER ROLE username SUPERUSER CREATEROLE CREATEDB REPLICATION BYPASSRLS;
```

Rename a user from `oldUsername` to `newUsername`
```sql
ALTER ROLE oldUsername RENAME TO newUsername;
```

Drop a user `username`
```sql
DROP USER username;
```
