# install private package lib


# setup 
```bash
go env -w GO111MODULE=off
export GOPATH=$(pwd)

mkdir src
mkdir src/mypkg
touch src/mypkg/bar.go
touch src/main.go

# run the main.go
go run src/main.go
```

## If not working open an another terminal and try!
