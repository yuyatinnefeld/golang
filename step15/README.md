
```bash
# create a module > go.mod will be created
go mod init simple-app

# install all module relevant packages > go.sum will be created
go mod tidy

# build the project > server file will be created
go build server.go

# call the server
./server 

# or
go run server.go

curl -i http://localhost:3000/api/list
curl -i http://localhost:3000/api/register
curl -i http://localhost:3000/api/param/1234

```

# Info
- https://docs.gofiber.io/api/app