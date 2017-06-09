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

package modules

import (
    "code.lukas.moe/x/karen/src/helpers"
    "github.com/bwmarrin/discordgo"
    "strings"
    "code.lukas.moe/x/karen/src/db"
)

// Announcement such as updates, downtimes...
type Announcement struct{}

// Commands that are availble to trigger an announcement
func (a *Announcement) Commands() []string {
    return []string{
        "announce",
    }
}

// Init func
func (a *Announcement) Init(s *discordgo.Session) {}

// Action of the announcement
func (a *Announcement) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    if !helpers.IsBotAdmin(msg.Author.ID) {
        return
    }

    title := ""
    contentSplit := strings.Fields(content)
    if len(contentSplit) < 1 {
        return
    }
    subcommand := contentSplit[0]
    text := content[len(subcommand):]

    switch subcommand {
    case "update":
        title = ":loudspeaker: **UPDATE**"
    case "downtime":
        title = ":warning: **DOWNTIME**"
    case "maintenance":
        title = ":clock5: **MAINTENANCE**"
    default:
        return
    }
    // Iterate through all joined guilds
    for _, guild := range session.State.Guilds {
        settings := db.GuildSettingsGetCached(guild.ID)
        // Check if we have an announcement channel set for this guild
        if settings.AnnouncementsEnabled {
            // Get the announcement channel id
            channelID := settings.AnnouncementsChannel
            // Send the announce to the channel
            session.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
                Title:       title,
                Description: text,
                Color:       0x0FADED,
            })
        }
    }
}
