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

package plugins

import (
    "code.lukas.moe/x/karen/src/helpers"
    "code.lukas.moe/x/karen/src/logger"
    "code.lukas.moe/x/karen/src/version"
    "github.com/bwmarrin/discordgo"
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
        }
    }()

    release := helpers.GetJSON(
        "https://git.lukas.moe/api/v3/projects/77/repository/tags/" + version.BOT_VERSION + "?private_token=9qvdMtLdxoC5amAmajN_",
    )

    c.log = map[string]string{
        "number": release.Path("name").Data().(string),
        "date":   release.Path("commit.committed_date").Data().(string),
    }

    if release.ExistsP("release.description") && release.Path("release.description").Data() != nil {
        c.log["body"] = release.Path("release.description").Data().(string)
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
