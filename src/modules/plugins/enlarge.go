package plugins

import (
    "github.com/bwmarrin/discordgo"
    "regexp"
    "code.lukas.moe/x/karen/src/helpers"
    "strings"
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
    regex := regexp.MustCompile(`.*?<:\w+:\d+>.*?`)
    emojis := regex.FindAllString(msg.Content, -1)
    emojiID := emojis[0]
    emojiInformation := strings.Split(emojiID, ":")
    if len(emojis) == 0 {
        _, err := session.ChannelMessageSend(msg.ChannelID, "I wasn't able to find an emoji :frowning:")
        helpers.Relax(err)
        return
    }
    _,err := session.ChannelMessageSend(msg.ChannelID, "https://cdn.discordapp.com/emojis/" + emojiInformation[2] + ".png")
    helpers.Relax(err)
}
