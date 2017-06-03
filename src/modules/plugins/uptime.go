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
    "code.lukas.moe/x/karen/src/metrics"
    "github.com/bwmarrin/discordgo"
    "strconv"
    "time"
)

type Uptime struct{}

func (u *Uptime) Commands() []string {
    return []string{
        "uptime",
    }
}

func (u *Uptime) Init(session *discordgo.Session) {

}

func (u *Uptime) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Get uptime
    bootTime, err := strconv.ParseInt(metrics.Uptime.String(), 10, 64)
    if err != nil {
        bootTime = 0
    }

    uptime := time.Now().Sub(time.Unix(bootTime, 0)).String()

    session.ChannelMessageSend(msg.ChannelID, ":hourglass_flowing_sand: "+uptime)
}
