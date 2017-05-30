package plugins

import (
    "github.com/bwmarrin/discordgo"
    "regexp"
    "code.lukas.moe/x/karen/src/helpers"
)

type Enlarge struct{}

func (p *Enlarge) Commands() []string {
    return []string{
        "enlarge",
    }
}

func (p *Enlarge) Init(session *discordgo.Session) {

}

func (p *Enlarge) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    emojis := regexp.MustCompile(`.*?<:\w+:\d+>.*?`).FindAllString(msg.Content, -1)
    emojiID := emojis[0]
    foundEmoji, _ := regexp.MatchString(`.*?<:\w+:\d+>.*?`, msg.Content)
    if !foundEmoji {
        _, err := session.ChannelMessageSend(msg.ChannelID, "I wasn't able to find an emoji :sad:")
        helpers.Relax(err)
        return
    }
    _,err := session.ChannelMessageSend(msg.ChannelID, discordgo.EndpointEmoji(emojiID))
    helpers.Relax(err)
}
