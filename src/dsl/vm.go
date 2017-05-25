package dsl

import (
    "github.com/yuin/gopher-lua"
    "path/filepath"
    "os"
    "code.lukas.moe/x/karen/src/helpers"
    "strings"
    "code.lukas.moe/x/karen/src/logger"
    "fmt"
    "code.lukas.moe/x/karen/src/dsl/api"
    "code.lukas.moe/x/karen/src/dsl/stdlib"
)

var vm *lua.LState

func Load() {
    vm = lua.NewState(lua.Options{
        IncludeGoStackTrace: true,
    })

    registerGlobals(vm)

    ex, err := helpers.FileExists("./_scripts")
    if !ex || err != nil {
        return
    }

    filepath.Walk("./_scripts", walkHandler)
}

func walkHandler(path string, info os.FileInfo, err error) error {
    if !info.IsDir() && strings.Contains(path, ".lua") {
        logger.INFO.L("Loading " + path)

        err = vm.DoFile(path)
        if err != nil {
            logger.ERROR.L("Error loading " + path)
            fmt.Printf("%#v", err)
            return nil
        }
    }

    return nil
}

func registerGlobals(vm *lua.LState) {
    vm.SetGlobal("RegisterReply", vm.NewFunction(api.RegisterReply))
    vm.SetGlobal("RegisterComplexReply", vm.NewFunction(api.RegisterComplexReply))
    vm.SetGlobal("__", vm.NewFunction(stdlib.GetText))
    vm.SetGlobal("_f", vm.NewFunction(stdlib.GetTextF))
}
