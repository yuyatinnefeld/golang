# Docker with Golang and MongoDB

```bash
# set local env
export NETWORK_NAME="mongo-network"
export MONGO_DB_USERNAME="admin"
export MONGO_DB_PASSWOR="password"
export MONGO_CONTAINER_NAME="mongo"
export MONGO_EXPRESS_CONTAINER_NAME="mongo-express"

vi docker-compose.yml

# create a 2 containers, network is created automatically
docker-compose up
docker network ls


# setup in Mongo Express GUI
localhost:8081

# create a databese =  "db"  
# create a table/collection = "jobs"
# input data
python input_data.py

# TODO: WIP HERE
# main.go works but server.go NOT
# run go 
go mod init mongo-app
go mod tidy
go build main.go
go build server.go


curl -X GET http://localhost:9090/jobs
```

## Source:
https://www.melvinvivas.com/my-first-go-microservice