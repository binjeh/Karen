/*
 * Karen - A highly efficient, multipurpose Discord bot written in Golang
 *
 * Copyright (C) 2015-2017 Lukas Breuer
 * Copyright (C) 2017 Subliminal Apps
 *
 * This file is a part of the Karen Discord-Bot Project ("Karen").
 *
 * Karen is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or (at your option) any later version.
 *
 * Karen is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 * See the GNU Affero General Public License for more details.
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package logger

import (
    "code.lukas.moe/x/karen/src/version"
    "fmt"
    "runtime"
    "strings"
    "time"
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
