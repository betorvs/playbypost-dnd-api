---
title: "Design"
date: 2021-08-25T17:56:38+02:00
draft: false
weight: 6
---

It was created using `golangspell` command. 
```bash
cd $GOPATH/src/github.com/betorvs/
mkdir playbypost-dnd 
cd playbypost-dnd
golangspell init github.com/betorvs/playbypost-dnd playbypost-dnd
```

## Dependency Management
The project is using [Go Modules](https://blog.golang.org/using-go-modules) for dependency management
Module: github.com/betorvs/playbypost-dnd

## Golang Spell
The project was initialized using [Golang Spell](https://github.com/golangspell/golangspell). A good post about it can be found [here.](https://medium.com/golangspell-go-fast-small-and-productive/go-fast-small-and-productive-with-golangspell-be193c65a382).

## Architectural Model
The Architectural Model adopted to structure the application is based on The Clean Architecture.
Further details can be found here: [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and in the Clean Architecture Book.

