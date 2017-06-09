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

/*
 * Full credit to @Seklfreak and his original implementation for Robyul2
 * See https://github.com/Seklfreak/Robyul2 for more information
 */

package plugins

import (
    "code.lukas.moe/x/karen/src/except"
    "code.lukas.moe/x/karen/src/i18n"
    "fmt"
    "github.com/bwmarrin/discordgo"
    "github.com/miekg/dns"
    "net"
    "strings"
)

type Dig struct{}

func (d *Dig) Commands() []string {
    return []string{
        "dig",
        "drill",
    }
}

func (d *Dig) Init(session *discordgo.Session) {
}

func (d *Dig) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    session.ChannelTyping(msg.ChannelID)

    args := strings.Fields(content)

    if len(args) < 2 {
        session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("bot.arguments.invalid"))
        return
    }

    dnsIp := "8.8.8.8"
    if len(args) >= 3 {
        dnsIp = strings.Replace(args[2], "@", "", 1)
    }

    var lookupType uint16
    if k, ok := dns.StringToType[strings.ToUpper(args[1])]; ok {
        lookupType = k
    }

    if k, ok := dns.StringToClass[strings.ToUpper(args[1])]; ok {
        lookupType = k
    }

    if lookupType == 0 {
        lookupType = dns.TypeA
    }

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(args[0]), lookupType)

    in, err := dns.Exchange(m, dnsIp+":53")
    if err != nil {
        if err, ok := err.(*net.OpError); ok {
            session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("bot.errors.general", err.Err.Error()))
            return
        } else {
            except.Handle(err)
        }
    }

    questionText := ""
    for _, question := range in.Question {
        questionText += question.String() + "\n"
    }

    if questionText == "" {
        questionText = "N/A"
    }

    answerText := ""
    for _, answer := range in.Answer {
        answerText += "`" + answer.String() + "`\n"
    }

    if answerText == "" {
        answerText = "N/A"
    }

    resultEmbed := &discordgo.MessageEmbed{
        Title:       fmt.Sprintf("Dig `%s`:", questionText),
        Description: answerText,
        Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Server: %s", dnsIp)},
    }

    _, err = session.ChannelMessageSendEmbed(msg.ChannelID, resultEmbed)
    except.Handle(err)
}
