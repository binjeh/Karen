package cache

import (
    "github.com/bwmarrin/discordgo"
)

func Channel(id string) (*discordgo.Channel, error) {
    ch, err := GetOrRequest(id, func(id string) (interface{}, error) {
        return GetSession().Channel(id)
    })

    return ch.(*discordgo.Channel), err
}
