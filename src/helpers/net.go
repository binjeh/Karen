/*
 *
 * Copyright (C) 2015-2017 Lukas Breuer. All rights reserved.
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

package helpers

import (
    "bytes"
    "errors"
    "code.lukas.moe/x/karen/src/version"
    "github.com/Jeffail/gabs"
    "io"
    "net/http"
    "strconv"
)

var DEFAULT_UA = "Karen/" + version.BOT_VERSION + " (https://git.lukas.moe/sn0w/karen)"

// NetGet executes a GET request to url with the Karen/Discord-Bot user-agent
func NetGet(url string) []byte {
    return NetGetUA(url, DEFAULT_UA)
}

// NetGetUA performs a GET request with a custom user-agent
func NetGetUA(url string, useragent string) []byte {
    // Allocate client
    client := &http.Client{}

    // Prepare request
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        panic(err)
    }

    // Set custom UA
    request.Header.Set("User-Agent", useragent)

    // Do request
    response, err := client.Do(request)
    Relax(err)

    // Only continue if code was 200
    if response.StatusCode != 200 {
        panic(errors.New("Expected status 200; Got " + strconv.Itoa(response.StatusCode)))
    } else {
        // Read body
        defer response.Body.Close()

        buf := bytes.NewBuffer(nil)
        _, err := io.Copy(buf, response.Body)
        Relax(err)

        return buf.Bytes()
    }
}

func NetPost(url string, data string) []byte {
    return NetPostUA(url, data, DEFAULT_UA)
}

func NetPostUA(url string, data string, useragent string) []byte {
    // Allocate client
    client := &http.Client{}

    // Prepare request
    request, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
    if err != nil {
        panic(err)
    }

    request.Header.Set("User-Agent", useragent)
    request.Header.Set("Content-Type", "application/json")

    // Do request
    response, err := client.Do(request)
    Relax(err)

    // Only continue if code was 200
    if response.StatusCode != 200 {
        panic(errors.New("Expected status 200; Got " + strconv.Itoa(response.StatusCode)))
    } else {
        // Read body
        defer response.Body.Close()

        buf := bytes.NewBuffer(nil)
        _, err := io.Copy(buf, response.Body)
        Relax(err)

        return buf.Bytes()
    }
}

// GetJSON sends a GET request to $url, parses it and returns the JSON
func GetJSON(url string) *gabs.Container {
    // Parse json
    json, err := gabs.ParseJSON(NetGet(url))
    Relax(err)

    return json
}
