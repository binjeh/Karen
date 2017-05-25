package api

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
    script := (&bridge.ProtoScript{}).Attach(
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
    for _, v := range listenerSlice {
        bridge.PushScript(v, script)
    }

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
    script := (&bridge.ProtoScript{}).Attach(
        func() string {
            return name
        },
        func() []string {
            return listenerSlice
        },
        func(author *discordgo.User, caller string, content string) string {
            // Construct VM call and execute
            err := L.CallByParam(lua.P{
                Fn: reply,
            })
            helpers.Relax(err)

            // Receive return value from stack
            ret := L.CheckString(-1)
            L.Pop(1)

            return ret
        },
    )

    // Push plugin to bridge
    for _, v := range listenerSlice {
        bridge.PushScript(v, script)
    }

    L.Push(lua.LBool(true))
    return 1
}
