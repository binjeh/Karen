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
    "code.lukas.moe/x/karen/src/i18n"
    "fmt"
    "github.com/bwmarrin/discordgo"
    "strings"
    "time"
)

// Timezone command
type Timezone struct{}

// Commands func
func (tz *Timezone) Commands() []string {
    return []string{
        "tz",
        "timezone",
    }
}

// Init func
func (tz *Timezone) Init(s *discordgo.Session) {}

// Action func
func (tz *Timezone) Action(command, content string, msg *discordgo.Message, session *discordgo.Session) {
    reminders := Reminders{}.getReminders(msg.Author.ID)
    if content == "" {
        if reminders.Timezone == "" {
            session.ChannelMessageSend(msg.ChannelID, i18n.GetText("plugins.time_zone.not_set"))
            return
        }
        session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Your timezone is: %q", reminders.Timezone))
        return
    }
    split := strings.Fields(content)
    if len(split) < 2 {
        session.ChannelMessageSend(msg.ChannelID, i18n.GetText("plugins.time_zone.missing_param"))
        return
    }
    switch split[0] {
    case "set", "s":
        _, err := time.LoadLocation(split[1])
        if err != nil {
            session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("plugins.time_zone.set.wrong_format", split[1]))
            return
        }
        reminders.Timezone = split[1]
        Reminders{}.setReminders(msg.Author.ID, reminders)
        session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("plugins.time_zone.set.success", split[1]))
    default:
        session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("plugins.time_zone.wrong_subcommand", split[0]))
    }
}
