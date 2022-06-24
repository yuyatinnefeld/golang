```bash
# create initial fiels

mkdir mymodule && cd mymodule
vi main.go
mkdir mypackage && cd mypackage
mypackage.go

# build image
docker build -t my-golang-image .

# run container
docker run -it --rm --name my-golang-app my-golang-image

# run go in the container
ls
go run .
```