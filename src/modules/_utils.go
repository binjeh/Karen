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

package modules

import (
    "strconv"
    "os"
    "fmt"
    "code.lukas.moe/x/karen/src/cache"
    "code.lukas.moe/x/karen/src/helpers"
    "code.lukas.moe/x/karen/src/logger"
    "code.lukas.moe/x/karen/src/metrics"
    "code.lukas.moe/x/karen/src/ratelimits"
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/karen/src/dsl"
    "code.lukas.moe/x/karen/src/dsl/bridge"
    "github.com/davecgh/go-spew/spew"
)

// command - The command that triggered this execution
// content - The content without command
// msg     - The message object
// session - The discord session
func CallPlugin(command string, content string, msg *discordgo.Message) bool {
    //#ifeq EXCLUDE_PLUGINS 1
    //#warning modules#CallBotPlugin() will be a no-op in this build
    return false
    //#else
    defer helpers.RecoverDiscord(msg)

    if ref, ok := pluginCache[command]; ok {
        // Consume a key for this action
        ratelimits.Container.Drain(1, msg.Author.ID)

        // Track metrics
        metrics.CommandsExecuted.Add(1)

        // Call the module
        ref.Action(command, content, msg, cache.GetSession())

        return true
    }

    return false
    //#endif
}

func CallScript(caller string, content string, msg *discordgo.Message) bool {
    //#ifeq EXCLUDE_SCRIPTING 1
    //#warning modules#CallScript() will be a no-op in this build
    return false
    //#else
    defer helpers.RecoverDiscord(msg)

    if ref, ok := scriptCache[caller]; ok {
        ratelimits.Container.Drain(1, msg.Author.ID)

        metrics.CommandsExecuted.Add(1)

        cache.GetSession().ChannelMessageSend(
            msg.ChannelID,
            ref.Action(
                msg.Author,
                caller,
                content,
            ),
        )

        return true
    }

    return false
    //#endif
}

// Init warms the caches and initializes the plugins
func Init(session *discordgo.Session) {
    //#if EXCLUDE_PLUGINS==1 && EXCLUDE_SCRIPTING==1
    //#warning modules#Init() will only print a line of text in this build
    //#else
    checkDuplicateCommands()
    listeners := ""
    logTemplate := ""
    //#endif

    //#ifeq EXCLUDE_SCRIPTING 0
    dsl.Load()
    //#endif

    //#ifeq EXCLUDE_PLUGINS 0
    pluginCount := len(PluginList)
    pluginCache = make(map[string]Plugin)

    logTemplate = "[PLUG] %s reacts to [ %s]"

    for i := 0; i < pluginCount; i++ {
        ref := PluginList[i]

        for _, cmd := range ref.Commands() {
            pluginCache[cmd] = ref
            listeners += cmd + " "
        }

        logger.INFO.L(fmt.Sprintf(
            logTemplate,
            helpers.Typeof(ref),
            listeners,
        ))
        listeners = ""

        ref.Init(session)
    }
    //#endif

    //#ifeq EXCLUDE_SCRIPTING 0
    scriptCount := len(*dsl_bridge.GetScripts())
    scriptCache = make(map[string]dsl_bridge.Script)
    logTemplate = `[SCRI] Script "%s" reacts to [ %s]`

    for _, s := range *dsl_bridge.GetScripts() {
        for _, listener := range s.Listeners() {
            scriptCache[listener] = s
            listeners += listener + " "
        }

        logger.INFO.L(fmt.Sprintf(
            logTemplate,
            s.Name(),
            listeners,
        ))
        listeners = ""
    }
    //#endif

    var lenPlugins string
    var lenScripts string

    //#ifeq EXCLUDE_PLUGINS 0
    lenPlugins = strconv.Itoa(pluginCount) + " plugins"
    //#else
    lenPlugins = "no plugins (-DEXCLUDE_PLUGINS=1)"
    //#endif

    //#ifeq EXCLUDE_SCRIPTING 0
    lenScripts = strconv.Itoa(scriptCount) + " scripts"
    //#else
    lenScripts = "no scripts (-DEXCLUDE_SCRIPTING=1)"
    //#endif

    logger.INFO.L(
        "Initializer finished. Loaded " + lenPlugins + " and " + lenScripts,
    )
}

//#ifeq EXCLUDE_PLUGINS 1
//#warning modules#checkDuplicateCommands() will be stripped from this build
//#else
func checkDuplicateCommands() {
    cmds := make(map[string]string)

    for _, plug := range PluginList {
        for _, cmd := range plug.Commands() {
            t := helpers.Typeof(plug)

            if occupant, ok := cmds[cmd]; ok {
                logger.ERROR.L("Failed to load " + t + " because '" + cmd + "' was already registered by " + occupant)
                os.Exit(1)
            }

            cmds[cmd] = t
        }
    }
}

//#endif
