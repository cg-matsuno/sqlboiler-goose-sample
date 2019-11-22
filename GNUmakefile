.PHONY: \
  build \
  install \
  all \
  exec \
  build-test-db \
  run-test-db \
  clear-test-db \

PKGS = ./. ./db

all: build test

build:
	(cd db && go build -v -o ./bin/goose ./cmd/)

install:
	go install ./cmd/

exec:
	docker exec -it sqlboiler_test_db psql -U matsuno -d testdb

build-test-db:
	docker build -t matsuno/sqlboiler-test-db -f Dockerfile .

run-test-db:
	docker run -t -i -d --name sqlboiler_test_db -m 2g -p 5432:5432 matsuno/sqlboiler-test-db

clear-test-db:
	docker rm -f sqlboiler_test_db
