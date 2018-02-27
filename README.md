# go-passe

[![GoDoc](https://godoc.org/github.com/redbubble/go-passe?status.svg)](https://godoc.org/github.com/redbubble/go-passe)
[![Go Report Card](https://goreportcard.com/badge/github.com/redbubble/go-passe)](https://goreportcard.com/report/github.com/redbubble/go-passe)
[![Build Status](https://travis-ci.org/redbubble/go-passe.svg?branch=master)](https://travis-ci.org/redbubble/go-passe)

Utility for neatly summarising Go JSON test output.

## Introduction

This allows the output from a `go test -json` command (the `-json` flag is
introduced in Go 1.10) to be piped into it for prettier formatting and more
easily identified test failures.

Example usage:

```
# Install this utility
$ go get -u github.com/redbubble/go-passe

# Run some tests
$ go test -v -json ./... | go-passe
```

### What it looks like

Without go-passe: `$ go test -v ./...`

![Screenshot of tests passing without go-passe](screenshots/screenshot-pass-without.png)

With go-passe: `$ go test -v -json ./... | go-passe`

![Screenshot of tests passing with go-passe](screenshots/screenshot-pass-with.png)

Test failures are summarised at the end for readability.

Without go-passe: `$ go test -v ./...`

![Screenshot of tests failing without go-passe](screenshots/screenshot-fail-without.png)

With go-passe: `$ go test -v -json ./... | go-passe`

![Screenshot of tests failing with go-passe](screenshots/screenshot-fail-with.png)


## Developing

`go-passe` uses `dep` for dependency management.

```
# Install dep if necessary
$ go get -u github.com/golang/dep/cmd/dep

# Ensure all dependencies are available/fetched
$ dep ensure
```

## License

go-passe, Copyright © 2018 Redbubble

This software is made available under an [MIT license](LICENSE).
