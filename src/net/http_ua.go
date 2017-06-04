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

func UA_GET(url string, ua string) []byte {
    return UA_SafeGET(url, 200, ua)
}

func UA_POST(url string, json string, ua string) []byte {
    return UA_SafePOST(url, json, 200, ua)
}

func UA_PUT(url string, ua string) []byte {
    return UA_SafePUT(url, 200, ua)
}

func UA_DELETE(url string, ua string) []byte {
    return UA_SafeDELETE(url, 200, ua)
}

func UA_SafeGET(url string, expectedStatus int, ua string) []byte {
    return executeRequest(
        newUARequest("GET", url, ua),
        expectedStatus,
    )
}

func UA_SafePOST(url string, json string, expectedStatus int, ua string) []byte {
    return executeRequest(
        newUARequestWithBody("POST", url, json, ua),
        expectedStatus,
    )
}

func UA_SafePUT(url string, expectedStatus int, ua string) []byte {
    return executeRequest(
        newUARequest("PUT", url, ua),
        expectedStatus,
    )
}

func UA_SafeDELETE(url string, expectedStatus int, ua string) []byte {
    return executeRequest(
        newUARequest("DELETE", url, ua),
        expectedStatus,
    )
}
