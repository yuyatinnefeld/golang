```bash
# create initial fiels

vi main.go 
vi dockerfile

# build image
docker build -t my-golang-image .

# run container
docker run -it --rm --name my-golang-app my-golang-image

# run go in the container
ls
go run .
```