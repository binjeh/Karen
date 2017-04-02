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
)

// command - The command that triggered this execution
// content - The content without command
// msg     - The message object
// session - The discord session
func CallBotPlugin(command string, content string, msg *discordgo.Message) {
    //#if(defined(EXCLUDE_PLUGINS))
        //#warning(modules#CallBotPlugin() will be a no-op in this build)
    //#else
        // Defer a recovery in case anything panics
        defer helpers.RecoverDiscord(msg)

        // Consume a key for this action
        ratelimits.Container.Drain(1, msg.Author.ID)

        // Track metrics
        metrics.CommandsExecuted.Add(1)

        // Call the module
        if ref, ok := pluginCache[command]; ok {
            (*ref).Action(command, content, msg, cache.GetSession())
        }
    //#endif
}

// msg     - The message that triggered the execution
// session - The discord session
func CallTriggerPlugin(trigger string, content string, msg *discordgo.Message) {
    //#if(defined(EXCLUDE_TRIGGERS))
        //#warning(modules#CallTriggerPlugin() will be a no-op in this build)
    //#else
        defer helpers.RecoverDiscord(msg)

        // Consume a key for this action
        ratelimits.Container.Drain(1, msg.Author.ID)

        // Redirect trigger
        if ref, ok := triggerCache[trigger]; ok {
            cache.GetSession().ChannelMessageSend(
                msg.ChannelID,
                (*ref).Response(trigger, content),
            )
        }
    //#endif
}

// Init warms the caches and initializes the plugins
func Init(session *discordgo.Session) {
    //#if(defined(EXCLUDE_PLUGINS) && defined(EXCLUDE_TRIGGERS))
        //#warning(modules#Init() will only print a line of text in this build)
    //#else
        checkDuplicateCommands()
        listeners := ""
        logTemplate := ""
    //#endif

    //#ifndef(EXCLUDE_PLUGINS)
    pluginCount := len(PluginList)
    pluginCache = make(map[string]*Plugin)

    logTemplate = "[PLUG] %s reacts to [ %s]"

    for i := 0; i < pluginCount; i++ {
        ref := &PluginList[i]

        for _, cmd := range (*ref).Commands() {
            pluginCache[cmd] = ref
            listeners += cmd + " "
        }

        logger.INFO.L("modules", fmt.Sprintf(
            logTemplate,
            helpers.Typeof(*ref),
            listeners,
        ))
        listeners = ""

        (*ref).Init(session)
    }
    //#endif

    //#ifndef(EXCLUDE_PLUGINS)
    triggerCount := len(TriggerPluginList)
    triggerCache = make(map[string]*TriggerPlugin)
    logTemplate = "[TRIG] %s gets triggered by [ %s]"

    for i := 0; i < triggerCount; i++ {
        ref := &TriggerPluginList[i]

        for _, trigger := range (*ref).Triggers() {
            triggerCache[trigger] = ref
            listeners += trigger + " "
        }

        logger.INFO.L("modules", fmt.Sprintf(
            logTemplate,
            helpers.Typeof(*ref),
            listeners,
        ))
        listeners = ""
    }
    //#endif

    var lenPlugins string
    var lenTriggers string

    //#ifndef(EXCLUDE_PLUGINS)
    lenPlugins = strconv.Itoa(len(PluginList)) + " plugins"
    //#else
    lenPlugins = "no plugins (-DEXCLUDE_PLUGINS)"
    //#endif

    //#ifndef(EXCLUDE_TRIGGERS)
    lenTriggers = strconv.Itoa(len(TriggerPluginList)) + " triggers"
    //#else
    lenTriggers = "no triggers (-DEXCLUDE_TRIGGERS)"
    //#endif

    logger.INFO.L(
        "modules",
        "Initializer finished. Loaded "+lenPlugins+" and "+lenTriggers,
    )
}

//#if(defined(EXCLUDE_PLUGINS) && defined(EXCLUDE_TRIGGERS))
//#warning(modules#checkDuplicateCommands() will be stripped from this build)
//#else
func checkDuplicateCommands() {
    cmds := make(map[string]string)

    //#ifndef(EXCLUDE_PLUGINS)
    for _, plug := range PluginList {
        for _, cmd := range plug.Commands() {
            t := helpers.Typeof(plug)

            if occupant, ok := cmds[cmd]; ok {
                logger.ERROR.L("modules", "Failed to load "+t+" because '"+cmd+"' was already registered by "+occupant)
                os.Exit(1)
            }

            cmds[cmd] = t
        }
    }
    //#endif

    //#ifndef(EXCLUDE_TRIGGERS)
    for _, trig := range TriggerPluginList {
        for _, cmd := range trig.Triggers() {
            t := helpers.Typeof(trig)

            if occupant, ok := cmds[cmd]; ok {
                logger.ERROR.L("modules", "Failed to load "+t+" because '"+cmd+"' was already registered by "+occupant)
                os.Exit(1)
            }

            cmds[cmd] = t
        }
    }
    //#endif
}

//#endif
