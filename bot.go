package main

import (
    "fmt"
    Logger "github.com/sn0w/Karen/logger"
    "github.com/bwmarrin/discordgo"
    "math/rand"
    "time"
    "strings"
    "regexp"
    "github.com/getsentry/raven-go"
    "github.com/sn0w/Karen/plugins"
    "github.com/sn0w/Karen/utils"
)

func BotOnReady(session *discordgo.Session, event *discordgo.Ready) {
    Logger.INF("Connected to discord!")
    fmt.Printf(
        "\n To add me to your discord server visit https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot&permissions=%s\n\n",
        utils.GetConfig().Path("discord.id").Data().(string),
        utils.GetConfig().Path("discord.perms").Data().(string),
    )

    discordSession = session

    // Init plugins
    tmpl := "[PLUG] %s reacts to [ %s]"

    for _, plugin := range plugins.GetPlugins() {
        cmds := ""

        for cmd := range plugin.Commands() {
            cmds += cmd + " "
        }

        Logger.INF(fmt.Sprintf(tmpl, plugin.Name(), cmds))
        plugin.Init(session)
    }

    // Async stats
    go func() {
        time.Sleep(3 * time.Second)

        users := make(map[string]string)
        channels := 0
        guilds := session.State.Guilds

        for _, guild := range guilds {
            channels += len(guild.Channels)

            for _, u := range guild.Members {
                users[u.User.ID] = u.User.Username
            }
        }

        Logger.INF(fmt.Sprintf("Servers:%d | Channels:%d | Users:%d", len(guilds), channels, len(users)))

        // Change avatar/name
        if session.State.User.Username != utils.GetConfig().Path("bot.name").Data().(string) ||
            session.State.User.Avatar != utils.GetConfig().Path("bot.avatar").Data().(string) {
            session.UserUpdate(
                "",
                "",
                utils.GetConfig().Path("bot.name").Data().(string),
                utils.GetConfig().Path("bot.avatar").Data().(string),
                "",
            )
        }
    }()

    // Run async game-changer
    go changeGameInterval(session)
}

func BotOnMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
    // Ignore other bots and @everyone/@here
    if (!message.Author.Bot && !message.MentionEveryone) {
        // Get the channel
        // Ignore the event if we cannot resolve the cannel
        channel, err := session.Channel(message.ChannelID)
        if err != nil {
            go raven.CaptureError(err, map[string]string{})
            return
        }

        // We only do things in guilds.
        // Get a friend already and stop chatting with bots
        if (!channel.IsPrivate) {
            // Check if the message contains @mentions
            if (len(message.Mentions) >= 1) {
                // Check if someone is mentioning us
                if (message.Mentions[0].ID == session.State.User.ID) {
                    go utils.CCTV(session, message.Message)

                    // Prepare content for editing
                    msg := message.Content

                    /// Remove our @mention
                    msg = strings.Replace(msg, "<@" + session.State.User.ID + ">", "", -1)

                    // Trim message
                    msg = strings.Trim(msg, " ")

                    switch {
                    case regexp.MustCompile("^REFRESH CHAT SESSION$").Match([]byte(msg)):
                        utils.RequireAdmin(session, message.Message, func() {
                            // Refresh cleverbot session
                            utils.CleverbotRefreshSession(channel.ID)
                            discordSession.ChannelMessageSend(channel.ID, ":cyclone: Refreshed!")
                        })
                        return

                    case regexp.MustCompile("^SET PREFIX (.){0,10}$").Match([]byte(msg)):
                        utils.RequireAdmin(session, message.Message, func() {
                            // Set new prefix
                            err := utils.SetPrefixForServer(channel.GuildID, strings.Replace(msg, "SET PREFIX ", "", 1))

                            if err != nil {
                                utils.SendError(session, message.Message, err)
                            } else {
                                discordSession.ChannelMessageSend(channel.ID, ":white_check_mark: Saved!")
                            }
                        })
                        return

                    default:
                        // Send to cleverbot
                        session.ChannelTyping(message.ChannelID)
                        // Resolve other @mentions before sending the message
                        for _, user := range message.Mentions {
                            msg = strings.Replace(msg, "<@" + user.ID + ">", user.Username, -1)
                        }

                        // Remove smileys
                        msg = regexp.MustCompile(`:\w+:`).ReplaceAllString(msg, "")

                        // Send to cleverbot
                        utils.CleverbotSend(session, channel.ID, msg)
                        return
                    }
                }
            }

            // Only continue if a prefix is set
            prefix, err := utils.GetPrefixForServer(channel.GuildID)
            if err != nil {
                go raven.CaptureError(err, map[string]string{})
                return
            }

            // Split the message into parts
            parts := strings.Split(message.Content, " ")

            // Save a sanitized version of the command (no prefix)
            cmd := strings.Replace(parts[0], prefix, "", 1)

            // Check if the message is prefixed for us
            if (strings.HasPrefix(message.Content, prefix)) {
                // Check if the user calls for help
                if cmd == "h" || cmd == "help" {
                    // Find the longest plugin name and command
                    longestPlugin, longestCommand := 0, 0
                    for _, plugin := range plugins.GetPlugins() {
                        if len(plugin.Name()) > longestPlugin {
                            longestPlugin = len(plugin.Name())
                        }

                        for cmd := range plugin.Commands() {
                            if len(cmd) > longestCommand {
                                longestCommand = len(cmd)
                            }
                        }
                    }

                    // Print help of all plugins
                    msg := ""

                    msg += "Hi " + message.Author.Username + " :smiley:\n"
                    msg += "These are all usable commands:\n"
                    msg += "```\n"

                    for _, plugin := range plugins.GetPlugins() {
                        if plugin.HelpHidden() == false {
                            description := plugin.Description()

                            if description == "" {
                                description = "no description"
                            }

                            padding := (longestPlugin - len(plugin.Name())) + 8

                            msg += fmt.Sprintf(
                                "%s%s[%s]\n",
                                plugin.Name(),
                                strings.Repeat(" ", padding),
                                description,
                            )

                            for cmd, usage := range plugin.Commands() {
                                if usage == "" {
                                    usage = "(no usage information)"
                                }

                                cmdPadding := (longestCommand - len(cmd)) + 6

                                msg += fmt.Sprintf(
                                    "\t%s%s%s\n",
                                    prefix + cmd,
                                    strings.Repeat(" ", cmdPadding),
                                    usage,
                                )
                            }

                            msg += "\n"
                        }
                    }

                    msg += "\n```"

                    discordSession.ChannelMessageSend(
                        message.ChannelID,
                        fmt.Sprintf("<@%s> :mailbox_with_mail:", message.Author.ID),
                    )

                    uc, err := discordSession.UserChannelCreate(message.Author.ID)
                    if err != nil {
                        go raven.CaptureError(err, map[string]string{})
                        return
                    }

                    discordSession.ChannelMessageSend(uc.ID, msg)
                } else {
                    // Check if a module matches said command
                    // Do nothing otherwise
                    plugins.CallBotPlugin(
                        cmd,
                        strings.Replace(message.Content, prefix + cmd, "", -1),
                        message.Message,
                        discordSession,
                    )
                }
            }
        }
    }
}

func changeGameInterval(session *discordgo.Session) {
    for {
        err := session.UpdateStatus(0, games[rand.Intn(len(games))])
        if err != nil {
            raven.CaptureError(err, map[string]string{})
        }

        time.Sleep(10 * time.Second)
    }
}

var games = []string{
    // Random stuff
    "async is the future!",
    "down with OOP!",
    "spoopy stuff",
    "Planking",

    // Kaomoji
    "ʕ•ᴥ•ʔ",
    "༼ つ ◕_◕ ༽つ",
    "(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧",
    "( ͡° ͜ʖ ͡°)",
    "¯\\_(ツ)_/¯",
    "(ง ͠° ͟ل͜ ͡°)ง",
    "ಠ_ಠ",
    "(╯°□°)╯︵ ʞooqǝɔɐɟ",
    "♪~ ᕕ(ᐛ)ᕗ",
    "\\ (•◡•) /",
    "｡◕‿◕｡",

    // actual games
    "Hearthstone",
    "Overwatch",
    "HuniePop",
    "Candy Crush",
    "Hyperdimension Neptunia",
    "Final Fantasy MCMX",
    "CIV V",
    "Pokemon Go",
    "Simulation Simulator 2016",
    "Half Life 3",
    "Nekopara",

    // software
    "with FFMPEG",
    "with libav",
    "with gophers",
    "with python",
    "with reflections",

    // names
    "with Shinobu-Chan",
    "with Ako-Chan",
    "with Nadeko",
    "with Miku",
    "with you O_o",
    "with cats",
    "with JOHN CENA",
    "with senpai",
    "with Serraniel#8978",
    "with 0xFADED#3237",
    "with C0untLizzi",
    "with moot",
    "with your waifu",
    "with Trump",
}