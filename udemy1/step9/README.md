# setup

```bash
## search 3rd pkg lib
https://pkg.go.dev/golang.org/x/tools/cmd/godoc

## e.g. talib
https://pkg.go.dev/github.com/markcheno/go-talib


cd step9
touch main.go

# activate Go Module import
export GO111MODULE="on"

# create go.mod file
go mod init myapp

# update packages
go mod tidy

# run the main.go
go run main.go
```