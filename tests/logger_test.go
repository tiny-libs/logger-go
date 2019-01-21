package logger_test

import (
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

func (s *Suite) TestSub(ch *check.C) {
	var val string

	loggerInstance := logger.InitLogger()
	val = "world"
	ch.Check(val, check.Equals, "world")
	loggerInstance.Log("test")
	loggerInstance.Test("test")
}