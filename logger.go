package logger

import (
    "github.com/davecgh/go-spew/spew"
    "log"
    "os"
)

type Logger struct {
    Info    func(...interface{})
    Warning func(...interface{})
    Error   func(...interface{})
    Dump    func(...interface{})
    Log    func(...interface{})
}

func InitLogger() *Logger {
    logger := &Logger{}

    logHandle := os.Stdout
    infoHandle := os.Stdout
    warningHandle := os.Stdout
    errorHandle := os.Stderr

    logger.Dump = func(data ...interface{}) {
        // Clear coloring
        logger.Log("\x1b[93mDUMP: \x1b[0m")
        spew.Dump(data...)
    }

    logger.Log = log.New(logHandle,
        "\x1b[0m",
        log.Ldate|log.Ltime|log.Lshortfile).Println

    logger.Info = log.New(infoHandle,
        "\x1b[96mINFO: ",
        log.Ldate|log.Ltime|log.Lshortfile).Println

    logger.Warning = log.New(warningHandle,
        "\x1b[33;1m WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile).Println

    logger.Error = log.New(errorHandle,
        "\x1b[31;1m  ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile).Println
    return logger
}