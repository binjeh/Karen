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
    "fmt"
    "code.lukas.moe/x/karen/src/helpers"
    "github.com/bwmarrin/discordgo"
    "math/rand"
    "net/url"
)

type Giphy struct{}

func (g *Giphy) Commands() []string {
    return []string{
        "giphy",
        "gif",
    }
}

func (g *Giphy) Init(session *discordgo.Session) {

}

func (g *Giphy) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    const ENDPOINT = "http://api.giphy.com/v1/gifs/search"
    const API_KEY = "dc6zaTOxFJmzC"
    const RATING = "pg-13"
    const LIMIT = 5

    session.ChannelTyping(msg.ChannelID)

    // Send request
    json := helpers.GetJSON(
        fmt.Sprintf(
            "%s?q=%s&api_key=%s&rating=%s&limit=%d",
            ENDPOINT,
            url.QueryEscape(content),
            API_KEY,
            RATING,
            LIMIT,
        ),
    )

    // Get gifs
    gifs, err := json.Path("data").Children()
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, "Error parsing Giphy's response :frowning:")
        return
    }

    // Chose a random one
    m := ""
    if len(gifs) > 0 {
        m = gifs[rand.Intn(len(gifs))].Path("bitly_url").Data().(string)
    } else {
        m = "No gifs found :frowning:"
    }

    // Send the result
    session.ChannelMessageSend(msg.ChannelID, m)
}
