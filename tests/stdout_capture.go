package logger

import (
	"io"
	"bytes"
	"os"
)

var writePipe *os.File
var StdoutPointer *os.File

func startCapture() <-chan string {
	var read *os.File

	StdoutPointer = os.Stdout
	read, writePipe, _ = os.Pipe()
	os.Stdout = writePipe

    outC := make(chan string)
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, read)
        outC <- buf.String()
    }()

	return outC
}

func stopCapture() {
	writePipe.Close()
	os.Stdout = StdoutPointer
}