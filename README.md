# logger.go
Golang logger

## Available methods:
```go
logger.Dump // Reference to spew: https://github.com/davecgh/go-spew
logger.Log
logger.Info
logger.Warning
logger.Error
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