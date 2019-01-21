package logger

import (
	"io"
	"bytes"
	"os"
)

func captureStdout() <-chan string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

    outC := make(chan string)
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

	w.Close()
	os.Stdout = rescueStdout

	return outC
}