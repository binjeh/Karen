/*
 *
 * Copyright (C) 2015-2017 Lukas Breuer. All rights reserved.
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
