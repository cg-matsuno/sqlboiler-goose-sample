# sqlboiler-goose-sample

## Install Dependency

- psql client

```
brew install postgresql
```

## GOOSE

- install [goose](https://github.com/pressly/goose)

```
$ GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose
```

- build postgres image

```
$ make build-test-db
```

- create postgres server

```
$ make run-test-db
```

- migration

Apply all available migrations.

```
$ ./goose up
```

- roll back

```
$ ./goose down
```


## SQLBOILER

- generate

```
$ ./sqlboiler
```

- go test

```
$ go test ./dbmodels
```
