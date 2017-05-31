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

package metrics

import (
    "expvar"
    "code.lukas.moe/x/karen/src/logger"
    "github.com/bwmarrin/discordgo"
    "net/http"
    "runtime"
    "time"
    "code.lukas.moe/x/karen/src/config"
)

var (
    // MessagesReceived counts all ever received messages
    MessagesReceived = expvar.NewInt("messages_received")

    // UserCount counts all logged-in users
    UserCount = expvar.NewInt("user_count")

    // ChannelCount counts all watching channels
    ChannelCount = expvar.NewInt("channel_count")

    // GuildCount counts all joined guilds
    GuildCount = expvar.NewInt("guild_count")

    // CommandsExecuted increases after each command execution
    CommandsExecuted = expvar.NewInt("commands_executed")

    // PollsCreated increases everytime a new pool is created
    PollsCreated = expvar.NewInt("polls_created")

    // CleverbotRequests increases after each request to cleverbot.com
    CleverbotRequests = expvar.NewInt("cleverbot_requests")

    // CoroutineCount counts all running coroutines
    CoroutineCount = expvar.NewInt("coroutine_count")

    // Uptime stores the timestamp of the bot's boot
    Uptime = expvar.NewInt("uptime")
)

func Init() {
    ip := config.Get("core.metrics.ip").(string)
    port := config.Get("core.metrics.port").(string)

    logger.INFO.L("Listening on " + ip + " [TCP/" + port + "]")
    Uptime.Set(time.Now().Unix())
    go http.ListenAndServe(
        ip+":"+port,
        nil,
    )
}

// OnReady listens for said discord event
func OnReady(session *discordgo.Session, event *discordgo.Ready) {
    go CollectDiscordMetrics(session)
    go CollectRuntimeMetrics()
}

// OnMessageCreate listens for said discord event
func OnMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
    MessagesReceived.Add(1)
}

// CollectDiscordMetrics counts Guilds, Channels and Users
func CollectDiscordMetrics(session *discordgo.Session) {
    for {
        time.Sleep(15 * time.Second)

        users := make(map[string]string)
        channels := 0
        guilds := session.State.Guilds

        for _, guild := range guilds {
            channels += len(guild.Channels)
            for _, u := range guild.Members {
                users[u.User.ID] = u.User.Username
            }
        }

        UserCount.Set(int64(len(users)))
        ChannelCount.Set(int64(channels))
        GuildCount.Set(int64(len(guilds)))
    }
}

// CollectRuntimeMetrics counts all running coroutines
func CollectRuntimeMetrics() {
    for {
        time.Sleep(15 * time.Second)
        CoroutineCount.Set(int64(runtime.NumGoroutine()))
    }
}
