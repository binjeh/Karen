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
    "code.lukas.moe/x/karen/src/except"
    "github.com/bwmarrin/discordgo"
    "strconv"
    "time"
)

type Ping struct{}

func (p *Ping) Commands() []string {
    return []string{
        "ping",
    }
}

func (p *Ping) Init(session *discordgo.Session) {

}

func (p *Ping) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    start := time.Now()

    m, err := session.ChannelMessageSend(msg.ChannelID, ":ping_pong: Pong! :grin:")
    except.Handle(err)

    end := time.Now()
    session.ChannelMessageEdit(
        msg.ChannelID,
        m.ID,
        m.Content+" ("+strconv.Itoa(int(end.Sub(start)/time.Millisecond)/2)+"ms)",
    )
}
