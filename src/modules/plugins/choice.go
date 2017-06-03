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
    "github.com/bwmarrin/discordgo"
    "math/rand"
    "strings"
    "code.lukas.moe/x/karen/src/helpers"
)

type Choice struct{}

func (c *Choice) Commands() []string {
    return []string{
        "choose",
        "choice",
    }
}

func (c *Choice) Init(session *discordgo.Session) {

}

func (c *Choice) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    if !strings.Contains(content, "|") {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.choice.wrong-syntax"))
        return
    }

    if len(msg.Mentions) > 0 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.choice.no-mentions"))
        return
    }

    choices := strings.Split(content, "|")
    choice := strings.Replace(choices[rand.Intn(len(choices))], "`", "", -1)

    if len(choice) == 0 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.choice.weird-input"))
        return
    }

    session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF(
        "plugins.choice.result",
        strings.TrimSpace(choice),
    ))
}
