# logger.go
Golang logger

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