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

func GET(url string) []byte {
    return SafeGET(url, 200)
}

func POST(url string, json string) []byte {
    return SafePOST(url, json, 200)
}

func PUT(url string) []byte {
    return SafePUT(url, 200)
}

func DELETE(url string) []byte {
    return SafeDELETE(url, 200)
}

func SafeGET(url string, expectedStatus int) []byte {
    return executeRequest(
        newRequest("GET", url),
        expectedStatus,
    )
}

func SafePOST(url string, json string, expectedStatus int) []byte {
    return executeRequest(
        newRequestWithBody("POST", url, json),
        expectedStatus,
    )
}

func SafePUT(url string, expectedStatus int) []byte {
    return executeRequest(
        newRequest("PUT", url),
        expectedStatus,
    )
}

func SafeDELETE(url string, expectedStatus int) []byte {
    return executeRequest(
        newRequest("DELETE", url),
        expectedStatus,
    )
}
