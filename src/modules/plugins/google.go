package plugins

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "net/url"
    "git.lukas.moe/sn0w/Karen/x/helpers"
)

type Google struct{}

func (g *Google) Commands() []string {
    return []string{
        "google",
        "goog",
    }
}

func (g *Google) Init(session *discordgo.Session) {

}

func (g *Google) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    if len(content) < 2 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.google.no-term"))
        return
    }

    session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf(
        "<https://lmgtfy.com/?q=%s>",
        url.QueryEscape(content),
    ))
}
