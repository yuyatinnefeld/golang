```bash
# create initial fiels

# build image
docker build -t my-golang-image .

# run container
docker run -it --rm --name my-golang-app my-golang-image

# run go in the container
ls
go run .
go test
go test -v
```