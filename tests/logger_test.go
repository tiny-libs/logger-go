package logger

import (
	"io"
	"bytes"
	"os"
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

func (s *Suite) TestSub(ch *check.C) {
	loggerInstance := logger.InitLogger()
	var val string
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	val = "world"
	ch.Check(val, check.Equals, "world")

	// loggerInstance.Log("test")
	loggerInstance.Test("test")

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

	w.Close()
	os.Stdout = rescueStdout
	out := <-outC

	fmt.Printf("Captured: %s", out)

	out = <-captureStdout()
	fmt.Printf("Captured: %s", out)

	ch.Check(val, check.Equals, "world")
}