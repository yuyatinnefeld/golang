# install 3rd package lib

https://qiita.com/fetaro/items/31b02b940ce9ec579baf


# setup 
```bash

export GO111MODULE="on"

# install package
go install github.com/antonholmquist/jason@latest

# verify if the package is created
ls /Users/yuyatinnefeld/go/pkg/mod/github.com/antonholmquist

# create go.mod file
go mod init myapp

# update packages
go mod tidy

# run the main.go

go run main.go
```

