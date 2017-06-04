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
    "net/url"
    "regexp"
)

type Enlarge struct{}

func (p *Enlarge) Commands() []string {
    return []string{
        "enlarge",
        "wumbo",
    }
}

func (p *Enlarge) Init(session *discordgo.Session) {

}

func (p *Enlarge) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    emojis := regexp.MustCompile(`.*?<:(\w+):(\d+)>.*?`).FindStringSubmatch(content)

    if len(emojis) == 0 {
        _, err := session.ChannelMessageSend(msg.ChannelID, "I wasn't able to find an emoji :frowning:")
        helpers.Relax(err)
        return
    }

    emojiID := emojis[2]
    uri := "https://cdn.discordapp.com/emojis/" + emojiID + ".png"

    _, urlErr := url.ParseRequestURI(uri)
    if urlErr != nil {
        _, err := session.ChannelMessageSend(msg.ChannelID, "Error resolving the URL")
        helpers.Relax(err)
        return
    }

    _, err := session.ChannelMessageSend(msg.ChannelID, uri)
    helpers.Relax(err)
}
