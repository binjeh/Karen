package dsl

import (
    "code.lukas.moe/x/karen/src/dsl/api"
    "code.lukas.moe/x/karen/src/dsl/stdlib"
    "github.com/yuin/gopher-lua"
    "layeh.com/gopher-luar"
    "strings"
    "code.lukas.moe/x/karen/src/logger"
)

var libMapping = map[string]interface{}{
    "api.register_reply":         dsl_api.RegisterReply,
    "api.register_complex_reply": dsl_api.RegisterComplexReply,

    "utils.gettext":   dsl_stdlib.GetText,
    "utils.gettext_f": dsl_stdlib.GetTextF,
}

func applyLibMapping(L *lua.LState) {
    for k, v := range libMapping {
        name := "____KAREN_" + strings.Replace(strings.ToUpper(k), ".", "_", -1) + "____"

        logger.INFO.L("[LVM] Injecting global " + name)

        L.SetGlobal(
            name,
            luar.New(L, v),
        )
    }
}
