package logger

import (
    "fmt"
    "time"
    "runtime"
    "strings"
)

var (
    DEBUG_MODE = false
)

func (c LogLevel) LF(f string, x ...interface{}) {
    c.cout(fmt.Sprintf(f, x...))
}

func (c LogLevel) L(msg string) {
    c.cout(msg)
}

func (c LogLevel) cout(msg string) {
    if c == DEBUG && !DEBUG_MODE {
        return
    }

    // Get filename and line of caller
    filename := "???"
    line := 0

    _, file, line, ok := runtime.Caller(2)
    if ok {
        fileParts := strings.Split(file, "/")
        filename = fileParts[len(fileParts)-1]
    }

    // Only print line number in debug mode
    var codePoint string
    if c == DEBUG && DEBUG_MODE {
        codePoint = fmt.Sprintf("%s#%d", filename, line)
    } else {
        codePoint = fmt.Sprintf("%s", filename)
    }

    // Log message
    fmt.Printf(
        colors[c].Color("[%s] (%-7s) [%s] %s\n"),
        time.Now().Format("15:04:05 02-01-2006"),
        nicenames[c],
        codePoint,
        msg,
    )
}
