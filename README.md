# go-passe

[![GoDoc](https://godoc.org/github.com/redbubble/go-passe?status.svg)](https://godoc.org/github.com/redbubble/go-passe)
[![Go Report Card](https://goreportcard.com/badge/github.com/redbubble/go-passe)](https://goreportcard.com/report/github.com/redbubble/go-passe)
[![Build Status](https://travis-ci.org/redbubble/go-passe.svg?branch=master)](https://travis-ci.org/redbubble/go-passe)

Utility for neatly summarising Go JSON test output.

## Introduction

This allows the output from a `go test -json` command to be piped into it for
prettier formatting and more easily identified test failures.

Example usage:

```
# Install this utility
$ go get -u github.com/redbubble/go-passe

# Run some tests
$ go test -v -json ./... | go-passe
```
