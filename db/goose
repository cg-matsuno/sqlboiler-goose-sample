#!/bin/bash

GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose
go build -o $DIR/bin/goose $DIR/cmd/main.go
DIR=$(dirname "$0")
$DIR/bin/goose -dir $DIR/migrations postgres "user=matsuno dbname=testdb password=matsuno sslmode=disable" $@
