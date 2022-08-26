# setup

```bash
cd step6_2

# create project
mkdir myProject && cd myProject

# create module
export GO111MODULE="on"
mkdir mymodule && cd mymodule
go mod init mymodule

# create pakcage
mkdir mypackage && nano hello.go

cd mymodule 
go run main.go
```