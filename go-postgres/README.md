# go-postgres

## Create database

```bash
createdb go-postgres

psql go-postgres < create.sql
```

## Run sample

```
export POSTGRES_URL='postgres://localhost:5432/go-postgres?sslmode=disable'

go run .
```
