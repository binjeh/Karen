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

package dsl_api

import (
    "code.lukas.moe/x/karen/src/dsl/bridge"
    "github.com/bwmarrin/discordgo"
)

type ComplexReplyCallback func(author *discordgo.User, caller string, content string) string

func RegisterReply(name string, listeners []string, reply string) {
    // Construct a plugin for this script using ProtoScript
    script := (&dsl_bridge.ProtoScript{}).Attach(
        func() string {
            return name
        },
        func() []string {
            return listeners
        },
        func(author *discordgo.User, caller string, content string) string {
            return reply
        },
    )

    // Push plugin to bridge
    dsl_bridge.PushScript(script)
}

func RegisterComplexReply(name string, listeners []string, cb ComplexReplyCallback) {
    // Construct a plugin for this script using ProtoScript
    script := (&dsl_bridge.ProtoScript{}).Attach(
        func() string {
            return name
        },
        func() []string {
            return listeners
        },
        func(author *discordgo.User, caller string, content string) string {
            return cb(author, caller, content)
        },
    )

    // Push plugin to bridge
    dsl_bridge.PushScript(script)
}
