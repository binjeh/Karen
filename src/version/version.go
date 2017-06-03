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

package version

// Version related vars
// Set by compiler
var (
    // BOT_VERSION example: 0.5.2-4-g205bbb8
    BOT_VERSION string = "DEV_SNAPSHOT"

    // BUILD_TIME example: Fri Jan  6 00:45:46 CET 2017
    BUILD_TIME string = "UNSET"

    // BUILD_USER example: sn0w
    BUILD_USER string = "UNSET"

    // BUILD_HOST example: nepgear
    BUILD_HOST string = "UNSET"
)
