# setup

```bash
cd step6_1

mkdir src
mkdir src/myProject
mkdir src/myProject/mylib
touch src/myProject/mylib/math.go
touch src/myProject/main.go

cd step6_1
export GOPATH=$(pwd)
go env -w GO111MODULE=off
go run src/myProject/main.go
```