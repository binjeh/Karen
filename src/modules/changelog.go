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
    "code.lukas.moe/x/karen/src/logger"
    "code.lukas.moe/x/karen/src/net"
    "code.lukas.moe/x/karen/src/version"
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/karen/src/config"
    "net/http"
    "code.lukas.moe/x/karen/src/except"
    "bytes"
    "io"
    "github.com/Jeffail/gabs"
    "github.com/davecgh/go-spew/spew"
    "strings"
)

type Changelog struct {
    log map[string]string
}

func (c *Changelog) Commands() []string {
    return []string{
        "changelog",
        "changes",
        "updates",
    }
}

func (c *Changelog) Init(session *discordgo.Session) {
    logger.PLUGIN.L("Retrieving release information...")

    c.log = make(map[string]string)

    defer func() {
        err := recover()
        if err != nil {
            c.log = map[string]string{
                "number": version.BOT_VERSION,
                "date":   "-",
                "body":   "Sorry but i can't find a changelog for " + version.BOT_VERSION,
            }
            logger.PLUGIN.L("Network error. Applied fallback.")

            spew.Dump(err)
        }
    }()

    request, err := http.NewRequest("GET", config.Get("modules.changelog.url").(string), nil)
    except.Handle(err)

    request.Header.Set("User-Agent", net.USERAGENT)
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Accept", "application/vnd.github.v3+json")
    request.Header.Set("Authorization", "token "+config.Get("modules.changelog.key").(string))

    client := http.Client{}
    response, err := client.Do(request)
    except.Handle(err)

    defer response.Body.Close()

    buf := bytes.NewBuffer(nil)
    _, err = io.Copy(buf, response.Body)
    except.Handle(err)

    if response.StatusCode != 200 {
        panic("Unexpected HTTP Status")
    }

    releases, err := gabs.ParseJSON(buf.Bytes())
    except.Handle(err)

    // Map result
    release := releases.Data().([]interface{})[0].(map[string]interface{})

    c.log = map[string]string{
        "number": release["tag_name"].(string),
        "date":   release["published_at"].(string),
    }

    if body, ok := release["body"]; ok {
        c.log["body"] = body.(string)
    } else {
        c.log["body"] = "No changelog provided :("
    }

    c.log["body"] = strings.Replace(c.log["body"], "### New stuff", ":eight_spoked_asterisk: **NEW STUFF**", 1)
    c.log["body"] = strings.Replace(c.log["body"], "### Fixed stuff", ":wrench: **FIXED STUFF**", 1)
    c.log["body"] = strings.Replace(c.log["body"], "### Removed stuff", ":wastebasket: **REMOVED STUFF**", 1)
    c.log["body"] = strings.Replace(c.log["body"], "\n-", "\n•", -1)

    logger.PLUGIN.L("Done")
}

func (c *Changelog) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
        Color: 0x0FADED,
        Fields: []*discordgo.MessageEmbedField{
            {Name: "Version", Value: c.log["number"], Inline: true},
            {Name: "Date", Value: c.log["date"], Inline: true},
            {Name: "Changelog", Value: c.log["body"], Inline: false},
            {Name: "＿＿＿＿＿＿＿＿＿＿", Value: "Want to give feedback? Join the [Discord Server](https://discord.gg/wNPejct)!", Inline: false},
        },
    })
}
