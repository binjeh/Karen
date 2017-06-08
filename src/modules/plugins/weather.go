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
    "code.lukas.moe/x/karen/src/net"
    "github.com/bwmarrin/discordgo"
    "net/url"
    "regexp"
    "strings"
)

type Weather struct{}

func (w *Weather) Commands() []string {
    return []string{
        "weather",
        "wttr",
    }
}

func (w *Weather) Init(session *discordgo.Session) {

}

func (w *Weather) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelTyping(msg.ChannelID)

    if content == "" {
        session.ChannelMessageSend(msg.ChannelID, "You should pass a city :thinking:")
        return
    }

    text := string(net.UA_GET("http://wttr.in/"+url.QueryEscape(content), "curl/7.51.0"))
    if text == "ERROR" {
        session.ChannelMessageSend(msg.ChannelID, "Couldn't find that city :frowning:")
        return
    }

    lines := strings.Split(text, "\n")

    session.ChannelMessageSend(
        msg.ChannelID,
        "```\n"+regexp.MustCompile("\\[.*?m").ReplaceAllString(strings.Join(lines[0:7], "\n"), "")+"\n```",
    )
}
