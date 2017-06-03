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

package cleverbot

import (
    "github.com/bwmarrin/discordgo"
    "github.com/ugjka/cleverbot-go"
    "code.lukas.moe/x/karen/src/config"
)

// cleverbotSessions stores all cleverbot connections
var cleverbotSessions map[string]*cleverbot.Session

// CleverbotSend sends a message to cleverbot and responds with it's answer.
func Send(session *discordgo.Session, channel string, message string) {
    var msg string

    if _, e := cleverbotSessions[channel]; !e {
        if len(cleverbotSessions) == 0 {
            cleverbotSessions = make(map[string]*cleverbot.Session)
        }

        RefreshSession(channel)
    }

    response, err := cleverbotSessions[channel].Ask(message)
    if err != nil {
        if err == cleverbot.ErrTooManyRequests {
            msg = "I cannot talk to you right now. :speak_no_evil:\n" +
                "CleverBot costs money, and the plan I'm currently on has no requests left.\n" +
                "If you want to help 0xFADED buying larger plans, cosider making a donation on his patreon. :innocent:\n" +
                "Link: <https://www.patreon.com/sn0w>"
        } else {
            msg = "Error :frowning:\n```\n" + err.Error() + "\n```"
        }
    } else {
        msg = response
    }

    session.ChannelMessageSend(channel, msg)
}

// CleverbotRefreshSession refreshes the cleverbot session for said channel
func RefreshSession(channel string) {
    cleverbotSessions[channel] = cleverbot.New(
        config.Get("modules.cleverbot.key").(string),
    )
}
