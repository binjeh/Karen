package plugins

import (
    "github.com/bwmarrin/discordgo"
    "math/rand"
    "strings"
    "git.lukas.moe/sn0w/Karen/helpers"
)

type Choice struct{}

func (c *Choice) Commands() []string {
    return []string{
        "choose",
        "choice",
    }
}

func (c *Choice) Init(session *discordgo.Session) {

}

func (c *Choice) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    if !strings.Contains(content, "|") {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.choice.wrong-syntax"))
        return
    }

    if len(msg.Mentions) > 0 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.choice.no-mentions"))
        return
    }

    choices := strings.Split(content, "|")
    choice := strings.Replace(choices[rand.Intn(len(choices))], "`", "", -1)

    if len(choice) == 0 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.choice.weird-input"))
        return
    }

    session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF(
        "plugins.choice.result",
        strings.TrimSpace(choice),
    ))
}
