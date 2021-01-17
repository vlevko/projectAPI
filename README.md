# projectAPI
Project management REST API

## Prerequisites
* PostgreSQL 12.4 running on `localhost:5432`, or using Docker `docker run -e POSTGRES_HOST_AUTH_METHOD=trust -it -p 5432:5432 -d postgres`
* DB connection set to environment variable `DATABASE_URL`, default is `postgres://postgres:@localhost:5432/postgres?sslmode=disable`
* Listening port set to environment variable `PORT`, default is `8000`
* Go 1.15.6

## Getting started
* Run locally
```bash
$ go build ./cmd/projectAPI
$ ./projectAPI --help
$ ./projectAPI -u
$ go test ./web
$ ./projectAPI
$ ./projectAPI -d
```
* Visit the [link](https://hidden-mountain-18927.herokuapp.com/)

## Endpoints
* Listening routes and implemented methods where {id} and {position} are integer numbers

  GET:
  - /projects/{id}/tasks
  
  GET, POST:
  - /projects
  - /projects/{id}/columns
  - /columns/{id}/tasks
  - /tasks/{id}/comments
  
  GET, PUT, DELETE:
  - /projects/{id} 
  - /columns/{id}
  - /tasks/{id}
  - /comments/{id}

  PUT:
  - /columns/{id}/{position}
  - /tasks/{id}/{position}
  - /tasks/{id}/columns/{id}

* Returning HTTP codes: 200, 201, 400, 404, 405, 500