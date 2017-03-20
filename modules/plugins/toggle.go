package plugins

import (
    "strings"

    "git.lukas.moe/sn0w/Karen/helpers"
    "github.com/bwmarrin/discordgo"
)

// Toggle command
type Toggle struct{}

// Commands func
func (e *Toggle) Commands() []string {
    return []string{
        "toggle",
    }
}

// Init func
func (e *Toggle) Init(s *discordgo.Session) {}

// Action func
func (e *Toggle) Action(command, content string, msg *discordgo.Message, session *discordgo.Session) {
	if !helpers.IsAdmin(msg) {
		session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.toggle.unauthorized", msg.Author.ID))
	}
    split := strings.Fields(content)
    if len(split) < 1 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.empty_content"))
        return
    }
    module := split[0]
    channel, err := session.Channel(msg.ChannelID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.failure"))
        return
    }
    switch module {
    case "joins":
        action := "disabled"
        settings := helpers.GuildSettingsGetCached(channel.GuildID)
        if settings.JoinNotificationsEnabled {
            settings.JoinNotificationsEnabled = false
            settings.JoinNotificationsChannel = ""
        } else {
            action = "enabled"
            settings.JoinNotificationsEnabled = true
            settings.JoinNotificationsChannel = msg.ChannelID
        }
        err := helpers.GuildSettingsSet(channel.GuildID, settings)
        if err != nil {
            session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.failure"))
            return
        }
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.joins."+action))
    case "leaves":
        action := "disabled"
        settings := helpers.GuildSettingsGetCached(channel.GuildID)
        if settings.LeaveNotificationsEnabled {
            settings.LeaveNotificationsEnabled = false
            settings.LeaveNotificationsChannel = ""
        } else {
            action = "enabled"
            settings.LeaveNotificationsEnabled = true
            settings.LeaveNotificationsChannel = msg.ChannelID
        }
        err := helpers.GuildSettingsSet(channel.GuildID, settings)
        if err != nil {
            session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.failure"))
            return
        }
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.leaves."+action))
    case "announcements":
        action := "disabled"
        settings := helpers.GuildSettingsGetCached(channel.GuildID)
        if settings.AnnouncementsEnabled {
            settings.AnnouncementsEnabled = false
            settings.AnnouncementsChannel = ""
        } else {
            action = "enabled"
            settings.AnnouncementsEnabled = true
            settings.AnnouncementsChannel = msg.ChannelID
        }
        err := helpers.GuildSettingsSet(channel.GuildID, settings)
        if err != nil {
            session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.failure"))
            return
        }
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.toggle.announcements."+action))
	default:
		session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.toggle.not_match", module))
    }
}
