# Backend Web Application with Golang

## Tech Stacks 
- Golang
- Postgres
- gRPC
- Kubernetes
- AWS

## Installation / Software

### 01 - golang
- https://go.dev/doc/install
### 02 - db migration for golang
- https://github.com/golang-migrate/migrate

### 03 - db table definition for postgresql
- https://dbdiagram.io/d

### 04 - SQL compiler
- https://docs.sqlc.dev/en/latest/overview/install.html


## Steps

### setup postgres server

```bash
# create golang migration file
mkdir simplebank
cd simplebank
export MG_DIR="db/migration"
export MG_FILE_NAME="init_schema"
migrate create -ext sql -dir $MG_DIR -seq $MG_FILE_NAME

# create makefile
touch Makefile

# set env variables
. ./set-env.sh

# run docker container and setup db
make postgres
make createdb

# verify
docker exec -it postgres14 /bin/sh
psql simple_bank

# migrate db  
make migrateup

# verify the tables
docker exec -it postgres14 /bin/sh 
psql simple_bank
\dt

# migrate down
make migratedown

# verify the tables
docker exec -it postgres14 /bin/sh
psql simple_bank
\dt
```

### connect db and golang application
```bash
# check if the sqlc installed
sqlc version

# create an empty sqlc.yaml settings file
cd simplebank
sqlc init

# update the setting file
vi sqlc.yaml

# run sql query
vi query/account.sql
make sqlc

# verify the new files
ls db/sqlc


# create go modules
mkdir db
mkdir util
```
