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
