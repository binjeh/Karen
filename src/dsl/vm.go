package dsl

import (
    "github.com/yuin/gopher-lua"
    "path/filepath"
    "os"
    "code.lukas.moe/x/karen/src/helpers"
    "strings"
    "code.lukas.moe/x/karen/src/logger"
    "code.lukas.moe/x/karen/src/dsl/api"
    "code.lukas.moe/x/karen/src/dsl/stdlib"
    "layeh.com/gopher-luar"
)

func Load() {
    vm := lua.NewState(lua.Options{
        IncludeGoStackTrace: true,
    })

    vm.OpenLibs()
    registerGlobals(vm)

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

func registerGlobals(vm *lua.LState) {
    // API
    vm.SetGlobal("__KAREN_REGISTER_REPLY__", luar.New(vm, dsl_api.RegisterReply))
    vm.SetGlobal("__KAREN_REGISTER_COMPLEX__", luar.New(vm, dsl_api.RegisterComplexReply))

    // STDLIB
    vm.SetGlobal("__KAREN_GETTEXT__", luar.New(vm, dsl_stdlib.GetText))
    vm.SetGlobal("__KAREN_GETTEXT_F__", luar.New(vm, dsl_stdlib.GetTextF))
}
