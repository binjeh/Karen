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
    "code.lukas.moe/x/karen/src/metrics"
    "code.lukas.moe/x/karen/src/version"
    "fmt"
    "github.com/bwmarrin/discordgo"
    "github.com/dustin/go-humanize"
    "runtime"
    "strconv"
    "time"
)

type Stats struct{}

func (s *Stats) Commands() []string {
    return []string{
        "stats",
        "sys",
    }
}

func (s *Stats) Init(session *discordgo.Session) {

}

func (s *Stats) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Count guilds, channels and users
    users := make(map[string]string)
    channels := 0
    guilds := session.State.Guilds

    for _, guild := range guilds {
        channels += len(guild.Channels)

        for _, u := range guild.Members {
            users[u.User.ID] = u.User.Username
        }
    }

    // Get RAM stats
    var ram runtime.MemStats
    runtime.ReadMemStats(&ram)

    // Get uptime
    bootTime, err := strconv.ParseInt(metrics.Uptime.String(), 10, 64)
    if err != nil {
        bootTime = 0
    }

    uptime := time.Now().Sub(time.Unix(bootTime, 0)).String()

    session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
        Color: 0x0FADED,
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: fmt.Sprintf(
                "https://cdn.discordapp.com/avatars/%s/%s.jpg",
                session.State.User.ID,
                session.State.User.Avatar,
            ),
        },
        Fields: []*discordgo.MessageEmbedField{
            // Build
            {Name: "Build Time", Value: version.BUILD_TIME, Inline: false},
            {Name: "Build System", Value: version.BUILD_USER + "@" + version.BUILD_HOST, Inline: false},

            // System
            {Name: "Bot Uptime", Value: uptime, Inline: true},
            {Name: "Bot Version", Value: version.BOT_VERSION, Inline: true},
            {Name: "GO Version", Value: runtime.Version(), Inline: true},

            // Bot
            {Name: "Used RAM", Value: humanize.Bytes(ram.Alloc) + "/" + humanize.Bytes(ram.Sys), Inline: true},
            {Name: "Collected garbage", Value: humanize.Bytes(ram.TotalAlloc), Inline: true},
            {Name: "Running coroutines", Value: strconv.Itoa(runtime.NumGoroutine()), Inline: true},

            // Discord
            {Name: "Connected servers", Value: strconv.Itoa(len(guilds)), Inline: true},
            {Name: "Watching channels", Value: strconv.Itoa(channels), Inline: true},
            {Name: "Users with access to me", Value: strconv.Itoa(len(users)), Inline: true},

            // Link
            {Name: "Want more stats and awesome graphs?", Value: "Visit [stats.karen.vc](https://stats.karen.vc)", Inline: false},
        },
    })
}
