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

package helpers

import "fmt"

// SecondsToDuration turns an int (seconds) into HH:MM:SS
func SecondsToDuration(input int) string {
    hours := 0
    minutes := 0
    seconds := input

    if seconds%60 != seconds {
        minutes = seconds / 60
        seconds %= 60
    }

    if minutes%60 != minutes {
        hours = minutes / 60
        minutes %= 60
    }

    return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
