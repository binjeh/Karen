/*
 *
 * Copyright (C) 2015-2017 Lukas Breuer. All rights reserved.
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
    "regexp"
)

type Minecraft struct{}

func (m *Minecraft) Commands() []string {
    return []string{
        "minecraft",
        "mc",
    }
}

func (m *Minecraft) Init(session *discordgo.Session) {

}

func (m *Minecraft) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Deferred error handler
    defer func() {
        err := recover()

        if err != nil {
            if regexp.MustCompile("(?i)expected status 200.*").Match([]byte(err.(string))) {
                session.ChannelMessageSend(msg.ChannelID, "Make sure that name is correct. \n I didn't find a thing :neutral_face:")
                return
            }
        }

        panic(err)
    }()

    // Set typing
    session.ChannelTyping(msg.ChannelID)

    // Request to catch server errors and 404's
    url := "https://minotar.net/body/" + content + "/300.png"
    helpers.NetGet(url)

    // If NetGet didn't panic send the url
    session.ChannelMessageSend(msg.ChannelID, "Here you go :smiley: \n "+url)

}
