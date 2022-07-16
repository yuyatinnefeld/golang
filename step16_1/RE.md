
```bash
# start postgresql server
docker run --name basic-postgres --rm \
    -e POSTGRES_USER=bucketeer \
    -e POSTGRES_PASSWORD=bucketeer_pass \
    -e POSTGRES_DB=bucketeer_db \
    -e PGDATA=/var/lib/postgresql/data/pgdata -v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:14.1-alpine

# install module
go mod init gitlab.com/idoko/bucketeer

# install build
go get github.com/go-chi/chi github.com/go-chi/render github.com/lib/pq

# create Dockerfile, docker-composer.yml, .env, etc

# install golang-migrate pkg and setting up the db
brew install golang-migrate
migrate create -ext sql -dir db/migrations -seq create_items_table
export POSTGRESQL_URL="postgres://bucketeer:bucketeer_pass@localhost:5432/bucketeer_db?sslmode=disable"
migrate -database ${POSTGRESQL_URL} -path db/migrations up

# stop the postgres server
docker stop basic-postgres

# run the containers
docker-compose up --build

# verify
curl -X POST http://localhost:8080/items -H "Content-type: application/json" -d '{ "name": "swim across the River Benue", "description": "ho ho ho"}'

```

# Info
- https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/