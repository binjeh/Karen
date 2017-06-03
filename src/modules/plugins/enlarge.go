package plugins

import (
    "github.com/bwmarrin/discordgo"
    "regexp"
    "code.lukas.moe/x/karen/src/helpers"
    "net/url"
)

type Enlarge struct{}

func (p *Enlarge) Commands() []string {
    return []string{
        "enlarge",
        "wumbo",
    }
}

func (p *Enlarge) Init(session *discordgo.Session) {

}

func (p *Enlarge) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    emojis := regexp.MustCompile(`.*?<:(\w+):(\d+)>.*?`).FindStringSubmatch(content)

    if len(emojis) == 0 {
        _, err := session.ChannelMessageSend(msg.ChannelID, "I wasn't able to find an emoji :frowning:")
        helpers.Relax(err)
        return
    }

    emojiID := emojis[2]
    uri := "https://cdn.discordapp.com/emojis/" + emojiID + ".png"

    _, urlErr := url.ParseRequestURI(uri)
    if urlErr != nil {
        _, err := session.ChannelMessageSend(msg.ChannelID, "Error resolving the URL")
        helpers.Relax(err)
        return
    }

    _, err := session.ChannelMessageSend(msg.ChannelID, uri)
    helpers.Relax(err)
}
