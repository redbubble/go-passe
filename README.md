# go-test-json-summary

[![GoDoc](https://godoc.org/github.com/redbubble/go-test-json-summary?status.svg)](https://godoc.org/github.com/redbubble/go-test-json-summary)
[![Go Report Card](https://goreportcard.com/badge/github.com/redbubble/go-test-json-summary)](https://goreportcard.com/report/github.com/redbubble/go-test-json-summary)
[![Build Status](https://travis-ci.org/redbubble/go-test-json-summary.svg?branch=master)](https://travis-ci.org/redbubble/go-test-json-summary)

Utility for neatly summarising Go JSON test output.

## Introduction

This allows the output from a `go test -json` command to be piped into it for
prettier formatting and more easily identified test failures.

Example usage:

```
# Install this utility
$ go get -u github.com/redbubble/go-test-json-summary

# Run some tests
$ go test -v -json ./... | go-test-json-summary 
```
