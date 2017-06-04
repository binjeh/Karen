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
    "fmt"
    "github.com/bwmarrin/discordgo"
)

type Stone struct{}

func (s *Stone) Commands() []string {
    return []string{
        "stone",
    }
}

func (s *Stone) Init(session *discordgo.Session) {

}

func (s *Stone) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    mentionCount := len(msg.Mentions)

    if mentionCount == 0 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("bot.mentions.too-few"))
        return
    }

    if mentionCount > 1 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("bot.mentions.too-many"))
        return
    }

    session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf(
        "<@%s> IS GOING TO DIE!!!\n"+"COME ON GUYS! THROW SOME STONES WITH MEE!!!\n"+":grimacing: :wavy_dash::anger::dizzy_face:",
        msg.Mentions[0].ID,
    ))
}
