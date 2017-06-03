package cache

import "github.com/bwmarrin/discordgo"

func Guild(id string) (*discordgo.Guild, error) {
    ch, err := GetOrRequest(id, func(id string) (interface{}, error) {
        return GetSession().Guild(id)
    })

    return ch.(*discordgo.Guild), err
}
