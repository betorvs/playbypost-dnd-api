# playbypost-dnd

![Go Test](https://github.com/betorvs/playbypost-dnd/workflows/Go%20Test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/betorvs/playbypost-dnd/badge.svg?branch=main)](https://coveralls.io/github/betorvs/playbypost-dnd?branch=main)


## Environment variables

```sh
export PORT=8080

export APP_NAME=playbypost-dnd

export LOG_LEVEL=INFO
```

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=cover.out

go tool cover -html=cover.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```

## Docker Build

```sh
docker build .
```

### MongoDB Local

```
docker run -d -p 27017:27017 mongo
```

and export DB_CONNECTION_STRING
```
export DB_CONNECTION_STRING="mongodb://localhost:27017/playbypost"
```

## Refences

### Dependency Management
The project is using [Go Modules](https://blog.golang.org/using-go-modules) for dependency management
Module: github.com/betorvs/playbypost-dnd

### Golang Spell
The project was initialized using [Golang Spell](https://github.com/golangspell/golangspell). More details [here](1).

### Architectural Model
The Architectural Model adopted to structure the application is based on The Clean Architecture.
Further details can be found here: [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and in the Clean Architecture Book.

[1]: https://medium.com/golangspell-go-fast-small-and-productive/go-fast-small-and-productive-with-golangspell-be193c65a382