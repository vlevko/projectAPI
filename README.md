# projectAPI
Project management REST API

## Prerequisites
* PostgreSQL 12.4 running on `localhost:5432`, or using Docker `docker run -e POSTGRES_HOST_AUTH_METHOD=trust -it -p 5432:5432 -d postgres`
* DB connection set to environment variable `DATABASE_URL`, default is `postgres://postgres:@localhost:5432/postgres?sslmode=disable`
* Go 1.15.6

## Getting started
* Run locally
```bash
$ go build ./cmd/projectAPI
$ ./projectAPI --help
$ ./projectAPI -u
$ go test ./cmd/projectAPI
$ ./projectAPI
$ ./projectAPI -d
```
* Visit the [link](https://hidden-mountain-18927.herokuapp.com/)