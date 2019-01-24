# logger.go
Golang logger

[![Build Status](https://travis-ci.org/tiny-libs/logger-go.svg?branch=master)](https://travis-ci.org/tiny-libs/logger-go)
[![License](https://img.shields.io/npm/l/tiny-require.svg)](http://opensource.org/licenses/MIT)

## Available methods:
```go
log.Dump(...interface{}) // Reference to spew: https://github.com/davecgh/go-spew
log.Log(...interface{})
log.Info(...interface{})
log.Warning(...interface{})
log.Error(...interface{})
```

## Example usage:
```go
package main


import (
	log "github.com/tiny-libs/logger-go"
)

func main() {
	log.Info("Hello world")
}
```