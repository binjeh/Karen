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
    "code.lukas.moe/x/karen/src/net"
    "github.com/bwmarrin/discordgo"
    "net/url"
    "strconv"
)

type UrbanDict struct{}

func (u *UrbanDict) Commands() []string {
    return []string{
        "urban",
        "ub",
    }
}

func (u *UrbanDict) Init(session *discordgo.Session) {

}

func (u *UrbanDict) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelTyping(msg.ChannelID)

    if content == "" {
        session.ChannelMessageSend(msg.ChannelID, "You should pass a word to define :thinking:")
        return
    }

    endpoint := "http://api.urbandictionary.com/v0/define?term=" + url.QueryEscape(content)

    json := net.GETJson(endpoint)

    res, e := json.Path("list").Children()
    except.Handle(e)

    if len(res) == 0 {
        session.ChannelMessageSend(msg.ChannelID, "No results :neutral_face:")
        return
    }

    object, e := res[0].ChildrenMap()
    except.Handle(e)

    children, e := json.Path("tags").Children()
    except.Handle(e)

    tags := ""
    for _, child := range children {
        tags += child.Data().(string) + ", "
    }

    session.ChannelMessageSendEmbed(
        msg.ChannelID,
        &discordgo.MessageEmbed{
            Color:       0x134FE6,
            Title:       object["word"].Data().(string),
            Description: object["definition"].Data().(string),
            URL:         object["permalink"].Data().(string),
            Fields: []*discordgo.MessageEmbedField{
                {Name: "Example(s)", Value: object["example"].Data().(string), Inline: false},
                {Name: "Tags", Value: tags, Inline: false},
                {Name: "Author", Value: object["author"].Data().(string), Inline: true},
                {
                    Name: "Votes",
                    Value: ":+1: " + strconv.FormatFloat(object["thumbs_up"].Data().(float64), 'f', 0, 64) +
                        " | :-1: " + strconv.FormatFloat(object["thumbs_down"].Data().(float64), 'f', 0, 64),
                    Inline: true,
                },
            },
            Footer: &discordgo.MessageEmbedFooter{
                Text: "powered by urbandictionary.com",
            },
        },
    )
}
