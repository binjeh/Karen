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
    Logger "code.lukas.moe/x/karen/src/logger"
    "github.com/bwmarrin/discordgo"
    "strings"
    "time"
)

// WhoIs command
type WhoIs struct{}

// Commands for WhoIs
func (w *WhoIs) Commands() []string {
    return []string{
        "whois",
    }
}

// Init func
func (w *WhoIs) Init(s *discordgo.Session) {}

// Action will return info about the first @user
func (w *WhoIs) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Check if the msg contains at least 1 mention
    if len(msg.Mentions) == 0 {
        session.ChannelMessageSend(msg.ChannelID, "you need to @mention someone")
        return
    }

    // Get channel info
    channel, err := session.Channel(msg.ChannelID)
    if err != nil {
        Logger.PLUGIN.L(err.Error())
        return
    }

    // Guild info
    guild, err := session.Guild(channel.GuildID)
    if err != nil {
        Logger.PLUGIN.L(err.Error())
        return
    }

    // Get the member object for the @user
    target, err := session.GuildMember(guild.ID, msg.Mentions[0].ID)
    if err != nil {
        Logger.PLUGIN.L(err.Error())
        return
    }

    // The roles name of the @user
    roles := []string{}
    for _, grole := range guild.Roles {
        for _, urole := range target.Roles {
            if urole == grole.ID {
                roles = append(roles, grole.Name)
            }
        }
    }

    joined, _ := time.Parse(time.RFC3339, target.JoinedAt)

    session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
        Title: "Information about " + target.User.Username + "#" + target.User.Discriminator,
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: helpers.GetAvatarUrl(target.User),
        },
        Color: 0x0FADED,
        Fields: []*discordgo.MessageEmbedField{
            {
                Name:   "Joined server",
                Value:  joined.Format(time.RFC1123),
                Inline: true,
            },
            {
                Name:   "Joined Discord",
                Value:  helpers.GetTimeFromSnowflake(target.User.ID).Format(time.RFC1123),
                Inline: true,
            },
            {
                Name:   "Avatar link",
                Value:  helpers.GetAvatarUrl(target.User),
                Inline: false,
            },
            {
                Name:   "Roles",
                Value:  strings.Join(roles, ","),
                Inline: true,
            },
        },
        Footer: &discordgo.MessageEmbedFooter{
            Text: "UserID: " + target.User.ID,
        },
    })
}
