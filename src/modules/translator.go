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
    "cloud.google.com/go/translate"
    "code.lukas.moe/x/karen/src/config"
    "code.lukas.moe/x/karen/src/except"
    "code.lukas.moe/x/karen/src/i18n"
    "context"
    "github.com/bwmarrin/discordgo"
    "golang.org/x/text/language"
    "google.golang.org/api/option"
    "strings"
)

type Translator struct {
    ctx    context.Context
    client *translate.Client
}

func (t *Translator) Commands() []string {
    return []string{
        "translator",
        "translate",
        "t",
    }
}

func (t *Translator) Init(session *discordgo.Session) {
    t.ctx = context.Background()

    client, err := translate.NewClient(
        t.ctx,
        option.WithAPIKey(config.Get("modules.google.translator.key").(string)),
    )
    except.Handle(err)
    t.client = client
}

func (t *Translator) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Assumed format: <lang_in> <lang_out> <text>
    parts := strings.Split(content, " ")

    if len(parts) < 3 {
        session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("plugins.translator.check_format"))
        return
    }

    source, err := language.Parse(parts[0])
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("plugins.translator.unknown_lang", parts[0]))
        return
    }

    target, err := language.Parse(parts[1])
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, i18n.GetTextF("plugins.translator.unknown_lang", parts[1]))
        return
    }

    translations, err := t.client.Translate(
        t.ctx,
        parts[2:],
        target,
        &translate.Options{
            Format: translate.Text,
            Source: source,
        },
    )

    if err != nil {
        //session.ChannelMessageSend(msg.ChannelID, i18n.GetText("plugins.translator.error"))
        except.SendError(msg, err)
        return
    }

    m := ""
    for _, piece := range translations {
        m += piece.Text + " "
    }

    session.ChannelMessageSend(
        msg.ChannelID,
        ":earth_africa: "+strings.ToUpper(source.String())+" => "+strings.ToUpper(target.String())+"\n```\n"+m+"\n```",
    )
}
