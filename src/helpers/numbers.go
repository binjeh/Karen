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

// HumanizeNumber takes an int and turns it into a word
func HumanizeNumber(num int) string {
    switch num {
    case 1:
        return "one"

    case 2:
        return "two"

    case 3:
        return "three"

    case 4:
        return "four"

    case 5:
        return "five"

    case 6:
        return "six"

    case 7:
        return "seven"

    case 8:
        return "eight"

    case 9:
        return "nine"

    case 0:
        return "zero"

    default:
        return ""
    }
}
