# logger.go
Golang logger

[![Build Status](https://travis-ci.org/tiny-libs/logger-go.svg?branch=master)](https://travis-ci.org/tiny-libs/logger-go)
[![License](https://img.shields.io/npm/l/tiny-require.svg)](http://opensource.org/licenses/MIT)

## Available methods:
```go
logger.Dump(...interface{}) // Reference to spew: https://github.com/davecgh/go-spew
logger.Log(...interface{})
logger.Info(...interface{})
logger.Warning(...interface{})
logger.Error(...interface{})
```

## Example usage:
```go
package main

var log *logger.Logger
func init() {
	log = logger.InitLogger()
	flag.Parse()
}

func main() {
	log.Info("Hello world")
}
```