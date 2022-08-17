# setup

```bash
cd step8

mkdir src
mkdir src/myProject
mkdir src/myProject/mylib
touch src/myProject/mylib/math.go
touch src/myProject/main.go
touch src/myProject/mylib/math_test.go

cd step8
export GOPATH=$(pwd)
go env -w GO111MODULE=off
go run src/myProject/main.go

go test ./src/myProject/mylib/...
go test ./...
go test -v ./...


## gofmt
cd src/myProject/mylib
gofmt math.go

gofmt -w math.go
```

## Ginkgo - Testing Framework for Go

- https://onsi.github.io/ginkgo/