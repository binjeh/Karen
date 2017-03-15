package plugins

import (
    "fmt"
    "git.lukas.moe/sn0w/Karen/helpers"
    "github.com/bwmarrin/discordgo"
    "math/rand"
    "net/url"
    "git.lukas.moe/sn0w/Karen/cache"
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

    // Disable GIFs for DOB because of abusive behaviour
    ch, err := cache.Channel(msg.ChannelID)
    if err != nil {
        return
    }
    if ch.ID == "265924970469654528" {
        session.ChannelMessageSend(msg.ChannelID, "This plugin has been temporarily disabled for DOB.")
        return
    }

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
