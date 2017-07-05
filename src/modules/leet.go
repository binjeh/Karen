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
    "github.com/bwmarrin/discordgo"
    "strings"
)

type Leet struct{}

var leetReplacements = map[string]string{
    "A": "@", "B": "8", "C": "C",
    "D": "D", "E": "3", "F": "ƒ",
    "G": "6", "H": "H", "I": "1",
    "J": "J", "K": "|<", "L": "L",
    "M": "/\\/\\", "N": "|/|", "O": "0",
    "P": "¶", "Q": "9", "R": "R",
    "S": "5", "T": "T", "U": "µ",
    "V": "\\//", "W": "\\/\\/", "X": "%",
    "Y": "¥", "Z": "Z",
}

func (l *Leet) Commands() []string {
    return []string{
        "leet",
        "l33t",
    }
}

func (l *Leet) Init(session *discordgo.Session) {

}

func (l *Leet) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    content = strings.ToUpper(content)

    for ascii, leet := range leetReplacements {
        content = strings.Replace(content, ascii, leet, -1)
    }

    session.ChannelMessageSend(msg.ChannelID, "```\n"+content+"\n```")
}
