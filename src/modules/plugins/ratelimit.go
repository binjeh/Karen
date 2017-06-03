package plugins

import (
    "code.lukas.moe/x/karen/src/ratelimits"
    "github.com/bwmarrin/discordgo"
    "strconv"
)

type Ratelimit struct{}

func (r *Ratelimit) Commands() []string {
    return []string{
        "limits",
    }
}

func (r *Ratelimit) Init(session *discordgo.Session) {

}

func (r *Ratelimit) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelMessageSend(
        msg.ChannelID,
        "You've still got "+strconv.Itoa(int(ratelimits.Container.Get(msg.Author.ID)))+" commands left",
    )
}
