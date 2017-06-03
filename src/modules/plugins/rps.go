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
    "regexp"
)

type RPS struct{}

func (r *RPS) Commands() []string {
    return []string{
        "rps",
    }
}

func (r *RPS) Init(session *discordgo.Session) {

}

func (r *RPS) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    switch {
    case regexp.MustCompile("(?i)rock").MatchString(content):
        session.ChannelMessageSend(msg.ChannelID, "I've chosen :newspaper:\nMy paper wraps your stone.\nI win :smiley:")
        return

    case regexp.MustCompile("(?i)paper").MatchString(content):
        session.ChannelMessageSend(msg.ChannelID, "I've chosen :scissors:\nMy scissors cuts your paper!\nI win :smiley:")
        return

    case regexp.MustCompile("(?i)scissors").MatchString(content):
        session.ChannelMessageSend(msg.ChannelID, "I've chosen :white_large_square:\nMy stone breaks your scissors.\nI win :smiley:")
        return
    }

    session.ChannelMessageSend(msg.ChannelID, "That's an odd or invalid choice for RPS :neutral_face:")
}
