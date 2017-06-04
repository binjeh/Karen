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

import (
    "code.lukas.moe/x/karen/src/assets"
    "fmt"
    "github.com/Jeffail/gabs"
    "math/rand"
    "strings"
)

var translations *gabs.Container

func LoadTranslations() {
    jsonFile, err := assets.Asset("assets/i18n.json")
    Relax(err)

    json, err := gabs.ParseJSON(jsonFile)
    Relax(err)

    translations = json
}

func GetText(id string) string {
    if !translations.ExistsP(id) {
        return id
    }

    item := translations.Path(id)

    // If this is an object return __
    if strings.Contains(item.String(), "{") {
        item = item.Path("__")
    }

    // If this is an array return a random item
    if strings.Contains(item.String(), "[") {
        arr := item.Data().([]interface{})
        return arr[rand.Intn(len(arr))].(string)
    }

    return item.Data().(string)
}

func GetTextF(id string, replacements ...interface{}) string {
    return fmt.Sprintf(GetText(id), replacements...)
}
