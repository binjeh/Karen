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

package dsl

import (
    "github.com/yuin/gopher-lua"
    "path/filepath"
    "os"
    "code.lukas.moe/x/karen/src/helpers"
    "strings"
    "code.lukas.moe/x/karen/src/logger"
)

func Load() {
    vm := lua.NewState(lua.Options{
        IncludeGoStackTrace: true,
    })

    vm.OpenLibs()
    applyLibMapping(vm)

    ex, err := helpers.FileExists("./scripts")
    if !ex || err != nil {
        return
    }

    filepath.Walk("./scripts/lib", func(path string, info os.FileInfo, err error) error {
        if info.IsDir() || !strings.Contains(path, ".lua") {
            return nil
        }

        logger.INFO.L("Loading library " + path)

        parts := strings.Split(path, "/")
        name := strings.Replace(parts[len(parts)-1], ".lua", "", -1)

        mod, err := vm.LoadFile(path)
        if err != nil {
            panic(err)
        }

        preload := vm.GetField(vm.GetField(vm.Get(lua.EnvironIndex), "package"), "preload")
        vm.SetField(preload, name, mod)

        return nil
    })

    filepath.Walk("./scripts", func(path string, info os.FileInfo, err error) error {
        if info.IsDir() || !strings.Contains(path, ".lua") {
            return nil
        }

        if strings.Contains(path, "lib") {
            return nil
        }

        logger.INFO.L("Loading script " + path)
        err = vm.DoFile(path)
        if err != nil {
            logger.ERROR.L("Error loading " + path)
            panic(err)
        }

        return nil
    })
}
