
```bash
# create a module > go.mod will be created
go mod init postgresql-app

# install all module relevant packages > go.sum will be created
go mod tidy

# build the project > server file will be created
go build server.go

# call the server
./server 

# or
go run server.go

curl -i http://localhost:3000/


```

# Info
- https://blog.logrocket.com/building-simple-app-go-postgresql/