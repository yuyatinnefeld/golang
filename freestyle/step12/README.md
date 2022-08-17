
```bash
# create initial fiels

# build image
docker build -t my-golang-image .

# run container
# terminal 1
docker run -it --rm --name my-golang-app my-golang-image
 go run tcpServer.go 1234

# terminal 2
docker ps
docker exec -it <container_id> bash
go run tcpClient.go 127.0.0.1:1234
>> Hello

# check the terminal 1 (server)

# stop (terminal 2)
STOP
```
