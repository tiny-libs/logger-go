package logger

import (
    "github.com/davecgh/go-spew/spew"
    "log"
    "os"
    "runtime"
    "strconv"
    "time"
    // "regexp"
    "path/filepath"
)

var workingDirectory string

func init() {
    workingDirectory, _ = os.Getwd()

    log.SetFlags(0)
}

func preparePrefix(data *interface{}, textPrefix string) {
    str, ok := (*data).(string)
    if (!ok) {
        return
    }
    logHandle := os.Stdout
    log.SetOutput(logHandle)

    _, filename, line, ok := runtime.Caller(2)
    if (!ok) {
        return
    }
    // functionObject := runtime.FuncForPC(pc)
    // extractFnName := regexp.MustCompile(`^.*\/.*?\.(.*)$`)
    // fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")

    rel, err := filepath.Rel(workingDirectory, filename)
    if (err != nil) {
        return
    }

    *data = time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + rel+":"+strconv.Itoa(line)+" " + textPrefix +"\x1b[0m " + str
}

func Log(data ...interface{}) {
    logHandle := os.Stdout
    log.SetOutput(logHandle)

    preparePrefix(&data[0], "\x1b[90m[LOG]")
    // data[0] = time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + rel+fnName+":"+strconv.Itoa(line)+" \x1b[93m[LOG]\x1b[0m " + str
    log.Println(data...)
}

func Info(data ...interface{}) {
    logHandle := os.Stdout
    log.SetOutput(logHandle)

    preparePrefix(&data[0], "\x1b[36m[INFO]")
    // data[0] = time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + rel+fnName+":"+strconv.Itoa(line)+" \x1b[93m[LOG]\x1b[0m " + str
    log.Println(data...)
}

func Warning(data ...interface{}) {
    logHandle := os.Stdout
    log.SetOutput(logHandle)

    preparePrefix(&data[0], "\x1b[93m[WARN]")
    // data[0] = time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + rel+fnName+":"+strconv.Itoa(line)+" \x1b[93m[LOG]\x1b[0m " + str
    log.Println(data...)
}

func Error(data ...interface{}) {
    logHandle := os.Stdout
    log.SetOutput(logHandle)

    preparePrefix(&data[0], "\x1b[31m[ERROR]")
    // data[0] = time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + rel+fnName+":"+strconv.Itoa(line)+" \x1b[93m[LOG]\x1b[0m " + str
    log.Println(data...)
}

func Dump(data ...interface{}) {
    logHandle := os.Stdout
    log.SetOutput(logHandle)

    _, filename, line, ok := runtime.Caller(2)
    if (!ok) {
        return
    }
    // functionObject := runtime.FuncForPC(pc)
    // extractFnName := regexp.MustCompile(`^.*\/.*?\.(.*)$`)
    // fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")

    rel, err := filepath.Rel(workingDirectory, filename)
    if (err != nil) {
        return
    }
    log.Println(time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + rel+":"+strconv.Itoa(line)+" \x1b[35m[DUMP]\x1b[0m ")
    spew.Dump(data...)
}