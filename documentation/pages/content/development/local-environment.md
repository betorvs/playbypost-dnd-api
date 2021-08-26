---
title: "Local Environment"
date: 2021-08-26T10:17:35+02:00
draft: false
weight: 8
---

## Environment variables

```sh
export PORT=8080

export APP_NAME=playbypost-dnd

export LOG_LEVEL=INFO
```

## Local Run

```bash
go build
./playbypost-dnd
```

## Test and coverage and lint

Run the tests and coverage

```sh 
TESTRUN=true go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```bash
golangci-lint run
```

### MongoDB Local

```bash
docker run -d -p 27017:27017 mongo
```

and export DB_CONNECTION_STRING
```bash
export DB_CONNECTION_STRING="mongodb://localhost:27017/playbypost"
```

## Docker Build

```bash
docker build .
```

