---
title: "Contributing"
date: 2021-08-26T09:59:31+02:00
draft: false
weight: 5
---

PlayByPost-DnD uses GitHub to manage reviews of pull requests:

- If you have a trivial fix or improvement, go ahead and create a pull request.
- If you plan to do something more involved, discuss your ideas on the relevant GitHub issue.

## Steps to contribute

For now, you need to add your fork as a remote on the original **\$GOPATH**/src/github.com/betorvs/playbypost-dnd clone, so:

```bash

$ go get github.com/betorvs/playbypost-dnd
$ cd $GOPATH/src/github.com/betorvs/playbypost-dnd # GOPATH is $HOME/go by default.

$ git remote add <FORK_NAME> <FORK_URL>
```

Notice: `go get` return `package github.com/betorvs/playbypost-dnd: no Go files in /go/src/github.com/betorvs/playbypost-dnd` is normal.


### Dependency management

We use [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) to manage dependencies on external packages.
This requires a working Go environment with version 1.15 or greater and git installed.

To add or update a new dependency, use the `go get` command:

```bash
# Pick the latest tagged release.
go get example.com/some/module/pkg

# Pick a specific version.
go get example.com/some/module/pkg@vX.Y.Z
```

Tidy up the `go.mod` and `go.sum` files:

```bash
go mod tidy
git add go.mod go.sum
git commit
```

You have to commit the changes to `go.mod` and `go.sum` before submitting the pull request.

## Coding Standards

### go imports
imports should follow `std libs`, `externals libs` and `local packages` format

Example
```
import (
        "context"
        "net/http"
        "os"
        "os/signal"
        "syscall"
        "time"

        "github.com/betorvs/playbypost-dnd/config"
        "github.com/betorvs/playbypost-dnd/controller"
        _ "github.com/betorvs/playbypost-dnd/gateway/customlog"
        _ "github.com/betorvs/playbypost-dnd/gateway/database"
        _ "github.com/betorvs/playbypost-dnd/gateway/diceroll"
        _ "github.com/betorvs/playbypost-dnd/gateway/mongodb"
        "github.com/labstack/echo/v4"
)
```