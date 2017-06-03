package plugins

import (
    "code.lukas.moe/x/karen/src/metrics"
    "github.com/bwmarrin/discordgo"
    "strconv"
    "time"
)

type Uptime struct{}

func (u *Uptime) Commands() []string {
    return []string{
        "uptime",
    }
}

func (u *Uptime) Init(session *discordgo.Session) {

}

func (u *Uptime) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Get uptime
    bootTime, err := strconv.ParseInt(metrics.Uptime.String(), 10, 64)
    if err != nil {
        bootTime = 0
    }

    uptime := time.Now().Sub(time.Unix(bootTime, 0)).String()

    session.ChannelMessageSend(msg.ChannelID, ":hourglass_flowing_sand: "+uptime)
}
