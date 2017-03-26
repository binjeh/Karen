package plugins

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"fmt"
	"git.lukas.moe/sn0w/Karen/helpers"
	"time"
)

// Timezone command
type Timezone struct {}

// Commands func
func (tz *Timezone) Commands() []string {
    return []string {
        "tz",
        "timezone",
    }
}

// Init func
func (tz *Timezone) Init(s *discordgo.Session) {}

// Action func
func (tz *Timezone) Action(command, content string, msg *discordgo.Message, session *discordgo.Session) {
    reminders := Reminders{}.getReminders(msg.Author.ID)
    if content == "" {
        if reminders.Timezone == "" {
            session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.time_zone.not_set"))
            return
        }
        session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Your timezone is: %q", reminders.Timezone))
        return
    }
    split := strings.Fields(content)
    if len(split) < 2 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.time_zone.missing_param"))
        return
    }
    switch split[0] {
    case "set", "s":
        _, err := time.LoadLocation(split[1])
        if err != nil {
            session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.time_zone.set.wrong_format", split[1]))
            return
        }
        reminders.Timezone = split[1]
        Reminders{}.setReminders(msg.Author.ID, reminders)
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.time_zone.set.success", split[1]))
    default:
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.time_zone.wrong_subcommand", split[0]))
    }
}