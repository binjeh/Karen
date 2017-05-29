package logger

import (
    "fmt"
    "time"
    "runtime"
    "strings"
    "code.lukas.moe/x/karen/src/version"
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

func PrintLogo() {
    //
    // Added in honor of 551b2e1ef67b2321b83bdeb9f35b8bc5ec4149a9.
    // A time where the code was "rusty", old and in early stages.
    // We'll never forget you Karen V1.5 ;-;
    //

    logo := []string{
        "",
        "0xFADED proudly presents...",
        " ▄    ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄        ▄ ",
        "▐░▌  ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░▌      ▐░▌",
        "▐░▌ ▐░▌ ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░▌░▌     ▐░▌",
        "▐░▌▐░▌  ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌          ▐░▌▐░▌    ▐░▌",
        "▐░▌░▌   ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░▌ ▐░▌   ▐░▌",
        "▐░░▌    ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌  ▐░▌",
        "▐░▌░▌   ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀█░█▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌   ▐░▌ ▐░▌",
        "▐░▌▐░▌  ▐░▌       ▐░▌▐░▌     ▐░▌  ▐░▌          ▐░▌    ▐░▌▐░▌",
        "▐░▌ ▐░▌ ▐░▌       ▐░▌▐░▌      ▐░▌ ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌     ▐░▐░▌",
        "▐░▌  ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░▌      ▐░░▌",
        " ▀    ▀  ▀         ▀  ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀        ▀▀ ",
    }

    for _, line := range logo {
        fmt.Println(line)
    }

    DumpInfoV()
    fmt.Println()
}

func DumpInfo() {
    VERBOSE.L("VERSION: " + version.BOT_VERSION)
    VERBOSE.L("BUILD TIME: " + version.BUILD_TIME)
    VERBOSE.L("BUILD USER: " + version.BUILD_USER)
    VERBOSE.L("BUILD HOST: " + version.BUILD_HOST)
}

func DumpInfoV() {
    fmt.Println("VERSION: " + version.BOT_VERSION)
    fmt.Println("BUILD TIME: " + version.BUILD_TIME)
    fmt.Println("BUILD USER: " + version.BUILD_USER)
    fmt.Println("BUILD HOST: " + version.BUILD_HOST)
}
