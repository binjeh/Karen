package plugins

import (
    "github.com/bwmarrin/discordgo"
)

type Music struct{}

func (m *Music) Commands() []string {
    return []string{
        "join",
        "leave",

        "play",
        "pause",
        "stop",
        "skip",
        "next",
        "playing",
        "np",

        "add",
        "list",
        "playlist",
        "random",
        "rand",
        "search",
        "find",

        "mdev",
    }
}

func (m *Music) Init(session *discordgo.Session) {
}

func (m *Music) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelMessageSend(
        msg.ChannelID,
        "Sorry but the music-plugin is currently being revamped",
    )
}
