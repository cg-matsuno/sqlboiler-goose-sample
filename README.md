# sqlboiler-goose-sample

## Install Dependency

- psql 11 client

```
brew install postgresql
```

```
psql --version
psql (PostgreSQL) 11.5
```

## GOOSE

- install [goose](https://github.com/pressly/goose)

```
GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose
```

- build postgres image

```
make build-test-db
```

- create postgres server

```
make run-test-db
```

- migration

Apply all available migrations.

```
cd db
./goose up
```

- roll back

```
cd db
./goose down
```


## SQLBOILER

- generate

```
cd db
./sqlboiler
```

- go test

```
cd db
go test ./dbmodels
```
