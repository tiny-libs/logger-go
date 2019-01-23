package logger

import (
	"fmt"
	check "gopkg.in/check.v1"
	"testing"
	logger "github.com/tiny-libs/logger-go"
	// "log"
)

var _ = check.Suite(new(Suite))

func Test(t *testing.T) {
	check.TestingT(t)
}

type Suite struct{}

func (s *Suite) TestLog(ch *check.C) {
	var captureChan <-chan string

	loggerInstance := logger.InitLogger()

	captureChan = startCapture()
	loggerInstance.Test("test")
	stopCapture()

	fmt.Printf("Captured: %s", <-captureChan)
	captureChan = startCapture()
	loggerInstance.Test("test2")
	stopCapture()

	fmt.Printf("Captured: %s", <-captureChan)
	captureChan = startCapture()
	loggerInstance.Test("test3")
	stopCapture()

	fmt.Printf("Captured: %s", <-captureChan)
	captureChan = startCapture()
	loggerInstance.Test("test4")
	stopCapture()

	fmt.Printf("Captured: %s", <-captureChan)
}