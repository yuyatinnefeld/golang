# Backend Web Application with Golang

# Tech Stacks 
Golang + Postgres + gRPC + Kubernetes + AWS


## Steps

### create DB / Tables
- https://dbdiagram.io/d
- result: https://dbdiagram.io/d/6317979e0911f91ba54bb7c3
- resouces/create_tabels.SQL

### setup postgres server

```bash
# setup
./setup_postgres.sh

# check date
root=# select now();

# check tables
\dt *.*

# create tables
create_tables.SQL

# check the new tables
\dt *.*
```


### db migration for golang
- https://github.com/golang-migrate/migrate
```bash
# install the golang-migrate CLI
brew install golang-migrate

# verify
migrate -version
migrate -help

# create first migration file
export MG_DIR="db/migrate"
export MG_FILE_NAME="init_schema"
migrate create -ext sql -dir $MG_DIR -seq $MG_FILE_NAME

# write down and up schema sql scripts

# run migrate db scripts

```