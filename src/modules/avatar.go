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
    "code.lukas.moe/x/karen/src/i18n"
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
        session.ChannelMessageSend(msg.ChannelID, i18n.GetText("bot.mentions.too-few"))
        return
    }

    if mentionCount > 1 {
        session.ChannelMessageSend(msg.ChannelID, i18n.GetText("bot.mentions.too-many"))
        return
    }

    mention := msg.Mentions[0]

    embed := &discordgo.MessageEmbed{
        Title: "Avatar",
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: helpers.GetAvatarUrl(mention),
        },
        Fields: []*discordgo.MessageEmbedField{
            &discordgo.MessageEmbedField{
                Name:  "Link",
                Value: helpers.GetAvatarUrl(mention),
            },
        },
        Color: 0x0FADED,
    }

    session.ChannelMessageSendEmbed(msg.ChannelID, embed)
}
