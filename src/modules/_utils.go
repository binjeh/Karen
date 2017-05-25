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
    //#ifdef EXCLUDE_PLUGINS
    //#warning modules#CallBotPlugin() will be a no-op in this build
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
    //#ifdef EXCLUDE_SCRIPTING
    //#warning modules#CallScript() will be a no-op in this build
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
    //#if defined(EXCLUDE_PLUGINS) && defined(EXCLUDE_SCRIPTING)
    //#warning modules#Init() will only print a line of text in this build
    //#else
    checkDuplicateCommands()
    listeners := ""
    logTemplate := ""
    //#endif

    //#ifndef EXCLUDE_SCRIPTING
    dsl.Load()
    //#endif

    //#ifndef EXCLUDE_PLUGINS
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

    //#ifndef EXCLUDE_SCRIPTING
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

    //#ifndef EXCLUDE_PLUGINS
    lenPlugins = strconv.Itoa(pluginCount) + " plugins"
    //#else
    lenPlugins = "no plugins (-DEXCLUDE_PLUGINS)"
    //#endif

    //#ifndef EXCLUDE_SCRIPTING
    lenScripts = strconv.Itoa(scriptCount) + " scripts"
    //#else
    lenScripts = "no scripts (-DEXCLUDE_SCRIPTING)"
    //#endif

    logger.INFO.L(
        "Initializer finished. Loaded " + lenPlugins + " and " + lenScripts,
    )
}

//#if defined(EXCLUDE_PLUGINS)
//#warning modules#checkDuplicateCommands() will be stripped from this build
//#else
func checkDuplicateCommands() {
    cmds := make(map[string]string)

    //#ifndef EXCLUDE_PLUGINS
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
    //#endif
}

//#endif
