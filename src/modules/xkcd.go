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
    "code.lukas.moe/x/karen/src/except"
    "code.lukas.moe/x/karen/src/net"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/bwmarrin/discordgo"
    "math/rand"
    "regexp"
    "strconv"
    "strings"
)

type XKCD struct{}

func (x *XKCD) Commands() []string {
    return []string{
        "xkcd",
    }
}

func (x *XKCD) Init(session *discordgo.Session) {

}

func (x *XKCD) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelTyping(msg.ChannelID)

    var link string

    if regexp.MustCompile("^\\d+$").MatchString(content) {
        link = "https://xkcd.com/" + content + "/info.0.json"
    } else if strings.Contains(content, "rand") {
        // Get latest number
        doc, err := goquery.NewDocument("https://xkcd.com")
        except.Handle(err)

        var num string
        for _, attr := range doc.Find("ul.comicNav").Children().Get(1).FirstChild.Attr {
            if attr.Key == "href" {
                num = attr.Val
                break
            }
        }

        num = strings.Replace(num, "/", "", -1)

        max, err := strconv.ParseInt(num, 10, 32)
        if err != nil {
            session.ChannelMessageSend(msg.ChannelID, "Error getting latest comic. Try again later :frowning:")
            return
        }

        link = "https://xkcd.com/" + strconv.Itoa(rand.Intn(int(max-1))+1) + "/info.0.json"
    } else {
        link = "https://xkcd.com/info.0.json"
    }

    json := net.GETJson(link)
    session.ChannelMessageSend(
        msg.ChannelID,
        fmt.Sprintf(
            "#%d from %s/%s/%s\n%s\n%s",
            int(json.Path("num").Data().(float64)),
            json.Path("day").Data().(string),
            json.Path("month").Data().(string),
            json.Path("year").Data().(string),
            json.Path("title").Data().(string),
            json.Path("img").Data().(string),
        ),
    )
    session.ChannelMessageSend(msg.ChannelID, json.Path("alt").Data().(string))
}
