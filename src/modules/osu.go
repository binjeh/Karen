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
    "code.lukas.moe/x/karen/src/config"
    "code.lukas.moe/x/karen/src/except"
    "code.lukas.moe/x/karen/src/net"
    "fmt"
    "github.com/bwmarrin/discordgo"
    "regexp"
    "strings"
)

type Osu struct{}

func (o *Osu) Commands() []string {
    return []string{
        "osu",
        "osu!mania",
        "osu!k",
        "osu!ctb",
        "osu!taiko",
    }
}

func (o *Osu) Init(session *discordgo.Session) {

}

func (o *Osu) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelTyping(msg.ChannelID)

    user := strings.TrimSpace(content)

    var mode string
    switch command {
    case "osu":
        mode = "0"
        break

    case "osu!taiko":
        mode = "1"
        break

    case "osu!ctb":
        mode = "2"
        break

    case "osu!mania", "osu!k":
        mode = "3"
        break
    }

    jsonc, err := net.GETJson(
        fmt.Sprintf(
            "https://osu.ppy.sh/api/get_user?k=%s&u=%s&type=u&m=%s",
            config.Get("modules.osu.key").(string),
            user,
            mode,
        ),
    ).Children()
    except.Handle(err)

    if len(jsonc) == 0 {
        session.ChannelMessageSend(msg.ChannelID, "User not found :frowning:")
        return
    }

    json := jsonc[0]
    html := string(net.GET("https://osu.ppy.sh/u/" + user))
    avatar := regexp.MustCompile(
        `"//a\.ppy\.sh/` + json.Path("user_id").Data().(string) + `_\d+\.\w{2,5}"`,
    ).FindString(html)

    if avatar == "" {
        avatar = "http://i.imgur.com/Ea1qmJX.png"
    } else {
        avatar = "https:" + avatar
    }

    avatar = strings.Replace(avatar, `"`, "", -1)

    if (!json.ExistsP("level")) || json.Path("level").Data() == nil {
        session.ChannelMessageSend(msg.ChannelID, "Seems like "+user+" didn't play this mode yet :thinking:")
        return
    }

    session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
        Color:       0xEF77AF,
        Description: "Showing stats for " + user,
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: avatar,
        },
        Fields: []*discordgo.MessageEmbedField{
            {Name: "Name", Value: json.Path("username").Data().(string), Inline: true},
            {Name: "Country", Value: json.Path("country").Data().(string), Inline: true},
            {Name: "Level", Value: json.Path("level").Data().(string), Inline: true},
            {Name: "Playcount", Value: json.Path("playcount").Data().(string), Inline: true},
            {Name: "Accuracy", Value: json.Path("accuracy").Data().(string) + "%", Inline: true},
            {Name: "Rank (Country)", Value: json.Path("pp_country_rank").Data().(string) + "th", Inline: true},
            {Name: "Rank (Global)", Value: json.Path("pp_rank").Data().(string) + "th", Inline: true},
        },
        Footer: &discordgo.MessageEmbedFooter{
            Text: "ppy powered :3",
        },
    })
}
