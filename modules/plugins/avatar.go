package plugins

import (
    "git.lukas.moe/sn0w/Karen/helpers"
    "github.com/bwmarrin/discordgo"
)

// Avatar command
type Avatar struct{}

// Commands func
func (a *Avatar) Commands() []string {
    return []string{
        "avatar",
    }
}

// Init func
func (a *Avatar) Init(session *discordgo.Session) {

}

// Action func
func (a *Avatar) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    mentionCount := len(msg.Mentions)

    if mentionCount == 0 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("bot.mentions.too-few"))
        return
    }

    if mentionCount > 1 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("bot.mentions.too-many"))
        return
    }

    mention := msg.Mentions[0]

    embed := &discordgo.MessageEmbed{
        Title: "Avatar",
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: helpers.GetAvatarUrl(mention),
        },
        Fields: []*discordgo.MessageEmbedField{
            &discordgo.MessageEmbedField {
                Name: "Link",
                Value: helpers.GetAvatarUrl(mention),
            },
        },
        Color: 0x0FADED,
    }

    session.ChannelMessageSendEmbed(msg.ChannelID, embed)
}
