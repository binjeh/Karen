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

package net

import (
    "bytes"
    "code.lukas.moe/x/karen/src/version"
    "fmt"
    "io"
    "net/http"
)

var (
    USERAGENT = "Karen/" + version.BOT_VERSION + " (https://github.com/SubliminalHQ/karen)"
)

func newRequest(method string, url string) *http.Request {
    return newUARequest(method, url, USERAGENT)
}

func newUARequest(method string, url string, ua string) *http.Request {
    request, err := http.NewRequest(method, url, nil)
    if err != nil {
        panic(err)
    }

    request.Header.Set("User-Agent", ua)

    return request
}

func newRequestWithBody(method string, url string, body string) *http.Request {
    return newUARequestWithBody(method, url, body, USERAGENT)
}

func newUARequestWithBody(method string, url string, body string, ua string) *http.Request {
    request, err := http.NewRequest(method, url, bytes.NewBufferString(body))
    if err != nil {
        panic(err)
    }

    request.Header.Set("User-Agent", ua)
    request.Header.Set("Content-Type", "application/json")

    return request
}

func executeRequest(request *http.Request, expectedStatus int) []byte {
    client := http.Client{}

    response, err := client.Do(request)
    if err != nil {
        panic(err)
    }

    defer response.Body.Close()

    buf := bytes.NewBuffer(nil)
    _, err = io.Copy(buf, response.Body)
    if err != nil {
        panic(err)
    }

    if response.StatusCode != expectedStatus {
        panic(fmt.Errorf(
            "Expected status %d; Got %d \nResponse: %#v",
            expectedStatus,
            response.StatusCode,
            buf.String(),
        ))
    }

    return buf.Bytes()
}
