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
)

type RandomCat struct{}

func (rc RandomCat) Commands() []string {
    return []string{
        "cat",
    }
}

func (rc RandomCat) Init(session *discordgo.Session) {

}

func (rc RandomCat) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    const ENDPOINT = "http://random.cat/meow"

    session.ChannelTyping(msg.ChannelID)

    json := helpers.GetJSON(ENDPOINT)
    session.ChannelMessageSend(
        msg.ChannelID,
        "MEOW! :smiley_cat: \n "+json.Path("file").Data().(string),
    )
}
