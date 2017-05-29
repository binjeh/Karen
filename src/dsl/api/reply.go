package dsl_api

import (
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/karen/src/dsl/bridge"
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
