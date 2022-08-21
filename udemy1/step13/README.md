# setup the golang package env 
```bash
# (TERMINAL 1)

cd step13

export GO111MODULE="on"

# create go.mod file
go mod init myapp

# update packages
go mod tidy

# create table and insert data into the table
go run main.go
```

# install sqlite
```bash
# (TERMINAL 2)

brew install sqlite

# go into sqlite
sqlite3

# exit
.exit

# go into the db and check the table
sqlite3 example.sql
.table

# check the data in the table (terminal2)
select * from person;

# get the saved data from the db table
go run main1.go

# delete the data from the table
go run main2.go

# call the table with the dynamic val
go run main3.go


```