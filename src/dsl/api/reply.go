package dsl_api

import (
    "github.com/yuin/gopher-lua"
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/karen/src/helpers"
    "code.lukas.moe/x/karen/src/dsl/bridge"
)

func RegisterReply(L *lua.LState) int {
    // Script name
    name := L.CheckString(1)

    // Script listeners + conversion from table to slice
    listenerTable := L.CheckTable(2)
    listenerSlice := make([]string, listenerTable.Len())
    listenerTable.ForEach(func(k lua.LValue, v lua.LValue) {
        listenerSlice = append(listenerSlice, v.String())
    })

    // The reply to send
    reply := L.CheckString(3)

    // Construct a plugin for this script using ProtoScript
    script := (&dsl_bridge.ProtoScript{}).Attach(
        func() string {
            return name
        },
        func() []string {
            return listenerSlice
        },
        func(author *discordgo.User, caller string, content string) string {
            return reply
        },
    )

    // Push plugin to bridge
    dsl_bridge.PushScript(script)

    L.Push(lua.LBool(true))
    return 1
}

func RegisterComplexReply(L *lua.LState) int {
    // Script name
    name := L.CheckString(1)

    // Script listeners + conversion from table to slice
    listenerTable := L.CheckTable(2)
    listenerSlice := make([]string, listenerTable.Len())
    listenerTable.ForEach(func(k lua.LValue, v lua.LValue) {
        listenerSlice = append(listenerSlice, v.String())
    })

    // The reply to send
    reply := L.CheckFunction(3)

    // Construct a plugin for this script using ProtoScript
    script := (&dsl_bridge.ProtoScript{}).Attach(
        func() string {
            return name
        },
        func() []string {
            return listenerSlice
        },
        func(author *discordgo.User, caller string, content string) string {
            //authorTable := lua.LTable{}
            //v := reflect.ValueOf(author)

            // Construct VM call and execute
            err := L.CallByParam(lua.P{
                Fn: reply,
            }, lua.LString(caller), lua.LString(content))
            helpers.Relax(err)

            // Receive return value from stack
            ret := L.Get(-1)
            L.Pop(1)

            return ret.String()
        },
    )

    // Push plugin to bridge
    dsl_bridge.PushScript(script)

    L.Push(lua.LBool(true))
    return 1
}
