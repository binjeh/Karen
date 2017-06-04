/*
 * Karen - A highly efficient, multipurpose Discord bot written in Golang
 *
 * Copyright (C) 2015-2017 Lukas Breuer
 * Copyright (C) 2017 Subliminal Apps
 *
 * This file is a part of the Karen Discord-Bot Project ("Karen").
 *
 * Karen is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or (at your option) any later version.
 *
 * Karen is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 * See the GNU Affero General Public License for more details.
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package plugins

import (
    "code.lukas.moe/x/karen/src/helpers"
    "github.com/bwmarrin/discordgo"
    "strings"
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
