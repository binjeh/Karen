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

package config

import (
    "github.com/pelletier/go-toml"
    "sync"
)

//noinspection GoNameStartsWithPackageName
const (
    CONFIG_FILE = "config.toml"
)

var (
    mutex = &sync.Mutex{}
    tree  *toml.Tree
)

func getTree() *toml.Tree {
    mutex.Lock()
    defer mutex.Unlock()

    if tree == nil {
        var err error

        tree, err = toml.LoadFile(CONFIG_FILE)
        if err != nil {
            panic(err)
        }
    }

    return tree
}

func Get(path string) interface{} {
    return getTree().Get(path)
}

func GetDefault(path string, def interface{}) interface{} {
    return getTree().GetDefault(path, def)
}
