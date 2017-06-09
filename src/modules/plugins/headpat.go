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
    "code.lukas.moe/x/karen/src/i18n"
    "github.com/bwmarrin/discordgo"
    "strings"
)

type Headpat struct{}

func (h *Headpat) Commands() []string {
    return []string{
        "headpat",
        "pat",
    }
}

func (h *Headpat) Init(session *discordgo.Session) {

}

func (h *Headpat) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Check mentions in the message
    mentionUsers := len(msg.Mentions)

    // Delete spaces from params
    params := strings.TrimSpace(content)

    // Case 1: pat yourself
    if params == "me" || mentionUsers == 1 && (msg.Author.ID == msg.Mentions[0].ID) {
        session.ChannelMessageSend(msg.ChannelID,
            i18n.GetText("bot.mentions.pat-yourself")+"\n"+"https://media.giphy.com/media/wUArrd4mE3pyU/giphy.gif",
        )
        return
    }

    // Case 2: pat @User#1234
    if mentionUsers == 1 {
        session.ChannelMessageSend(msg.ChannelID,
            i18n.GetTextF(
                "triggers.headpat.msg",
                msg.Author.ID,
                msg.Mentions[0].ID,
            )+ "\n"+ i18n.GetText("triggers.headpat.link"),
        )
        return
    }

    // Case 3: pat multiple users
    if msg.MentionEveryone || mentionUsers > 1 {
        session.ChannelMessageSend(msg.ChannelID, i18n.GetText("bot.mentions.pat-group"))
        return
    }

    // Case 4: no params || wrong params
    session.ChannelMessageSend(msg.ChannelID, i18n.GetText("bot.mentions.who-to-pat"))
}
