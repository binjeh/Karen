package stdlib

import (
    "github.com/yuin/gopher-lua"
    "code.lukas.moe/x/karen/src/helpers"
)

func GetText(L *lua.LState) int {
    str := L.CheckString(1)
    str = helpers.GetText(str)
    L.Push(lua.LString(str))

    return 1
}

func GetTextF(L *lua.LState) int {
    str := L.CheckString(1)

    argTable := L.CheckTable(2)
    argSlice := make([]interface{}, argTable.Len())

    argTable.ForEach(func(k lua.LValue, v lua.LValue){
        argSlice = append(argSlice, v.String())
    })

    str = helpers.GetTextF(str, argSlice...)

    L.Push(lua.LString(str))
    return 1
}
