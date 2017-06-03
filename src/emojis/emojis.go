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

package emojis

import "strconv"

var list = map[string]string{
    "0":  `0⃣`,
    "1":  `1⃣`,
    "2":  `2⃣`,
    "3":  `3⃣`,
    "4":  `4⃣`,
    "5":  `5⃣`,
    "6":  `6⃣`,
    "7":  `7⃣`,
    "8":  `8⃣`,
    "9":  `9⃣`,
    "10": `🔟`,
}

// revlist is the reverse version of list
var revlist map[string]string

func init() {
    revlist = make(map[string]string, len(list))
    for k, v := range list {
        revlist[v] = k
    }
}

// From returns the unicode emoji code for the symbol
func From(symbol string) string {
    return list[symbol]
}

// To returns the symbol from the emoji
func To(symbol string) string {
    return revlist[symbol]
}

// ToNumber returns the number that corresponds to
// the emoji
func ToNumber(emoji string) int {
    v, err := strconv.Atoi(revlist[emoji])
    if err != nil {
        v = -1
    }
    return v
}
