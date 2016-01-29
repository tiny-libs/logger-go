# logger.go
Golang logger

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