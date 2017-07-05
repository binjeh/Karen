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

package helpers

import (
    "code.lukas.moe/x/karen/src/cache"
    "code.lukas.moe/x/karen/src/db"
    "code.lukas.moe/x/karen/src/emojis"
    "code.lukas.moe/x/karen/src/models"
    "errors"
    "fmt"
    "github.com/bwmarrin/discordgo"
    "strconv"
    "time"
)

// NewPoll creates a pool for the guild
func NewPoll(title, guild, pollID, channelID, author string, fields ...models.PollField) bool {
    settings := db.GuildSettingsGetCached(guild)
    settings.Polls = append(settings.Polls, models.Poll{
        ID:        pollID,
        ChannelID: channelID,
        Title:     title,
        Fields:    fields,
        Open:      true,
        CreatedAt: time.Now(),
        CreatedBy: author,
    })
    return db.GuildSettingsSet(guild, settings) == nil
}

//NewPollFieldGenerator returns a new PollField
func NewPollFieldGenerator() func(title string) models.PollField {
    id := 0
    return func(title string) models.PollField {
        id++
        return models.PollField{
            ID:    id,
            Title: title,
            Votes: 0,
        }
    }
}

// RemovePoll removes a pool from the guild
func RemovePoll(guild, pollID string, msg *discordgo.Message) bool {
    settings := db.GuildSettingsGetCached(guild)
    polls := []models.Poll{}
    removed := false
    for _, p := range settings.Polls {
        if p.ID == pollID {
            if msg.Author.ID == p.CreatedBy || IsAdmin(msg) {
                removed = true
                continue
            } else {
                return false
            }
        }
        polls = append(polls, p)
    }
    settings.Polls = polls
    if !removed {
        return false
    }
    return db.GuildSettingsSet(guild, settings) == nil
}

// GetPoll returns a Poll
func GetPoll(guild, pollID string) (models.Poll, error) {
    settings := db.GuildSettingsGetCached(guild)
    for _, p := range settings.Polls {
        if p.ID == pollID {
            return p, nil
        }
    }
    return models.Poll{}, errors.New("Poll not found")
}

// UpdatePollMsg updates the poll msg
func UpdatePollMsg(guild, pollID string) bool {
    settings := db.GuildSettingsGetCached(guild)
    for _, p := range settings.Polls {
        if p.ID == pollID {
            session := cache.GetSession()
            msg := GetEmbedMsgFromPoll(p)
            _, err := session.ChannelMessageEditEmbed(p.ChannelID, p.ID, msg)
            return err == nil
        }
    }
    return false
}

// GetEmbedMsgFromPoll creates an Embed Msg from p
func GetEmbedMsgFromPoll(p models.Poll) *discordgo.MessageEmbed {
    mfields := []*discordgo.MessageEmbedField{}
    for _, field := range p.Fields {
        mfields = append(mfields, &discordgo.MessageEmbedField{
            Name:  fmt.Sprintf("%v) %v", field.ID, field.Title),
            Value: fmt.Sprintf("%v votes", field.Votes),
        })
    }
    Status := "CLOSED"
    if p.Open {
        Status = "OPEN"
    }
    Data := fmt.Sprintf("Participants: %v | Total votes: %v", p.TotalParticipants, p.TotalVotes)
    embed := &discordgo.MessageEmbed{
        Title:       p.Title,
        Description: Data,
        Fields:      mfields,
        Color:       0x0FADED,
        Footer: &discordgo.MessageEmbedFooter{
            Text: fmt.Sprintf("Poll ID: %v | Status: %s", p.ID, Status),
        },
    }
    return embed
}

// AddPollField adds a poll field
func AddPollField(guild, pollID, title string, msg *discordgo.Message) bool {
    settings := db.GuildSettingsGetCached(guild)
    for i, p := range settings.Polls {
        // If its the poll we're looking for
        if p.ID == pollID {
            // If this user is allowed to do this
            if msg.Author.ID == p.CreatedBy || IsAdmin(msg) {
                if len(p.Fields) == 10 {
                    return false
                }
                ID := p.Fields[len(p.Fields)-1].ID + 1
                settings.Polls[i].Fields = append(settings.Polls[i].Fields, models.PollField{
                    ID:    ID,
                    Title: title,
                    Votes: 0,
                })
                session := cache.GetSession()
                session.MessageReactionAdd(msg.ChannelID, pollID, emojis.From(strconv.Itoa(ID)))
                return db.GuildSettingsSet(guild, settings) == nil
            }
        }
    }
    return false
}

// RemovePollField removes the poll field with ID = fieldID
func RemovePollField(guild, pollID string, fieldID int, msg *discordgo.Message) bool {
    settings := db.GuildSettingsGetCached(guild)
    removed := false
    for i, p := range settings.Polls {
        // If its the poll we're looking for
        if p.ID == pollID {
            // If this user is allowed to do this
            if msg.Author.ID == p.CreatedBy || IsAdmin(msg) {
                if len(p.Fields) == 2 {
                    return false
                }
                fields := []models.PollField{}
                for _, pf := range p.Fields {
                    if pf.ID == fieldID {
                        removed = true
                        continue
                    }
                    fields = append(fields, pf)
                }
                settings.Polls[i].Fields = fields
                break
            }
        }
    }
    if !removed {
        return false
    }
    return db.GuildSettingsSet(guild, settings) == nil
}

// OpenPoll changes the state of a poll to open
func OpenPoll(guild, pollID string, msg *discordgo.Message) bool {
    settings := db.GuildSettingsGetCached(guild)
    opened := false
    for i, p := range settings.Polls {
        // If its the poll we're looking for
        if p.ID == pollID {
            // If this user is allowed to do this
            if msg.Author.ID == p.CreatedBy || IsAdmin(msg) {
                if p.Open {
                    return false
                }
                settings.Polls[i].Open = true
                settings.Polls[i].ClosedAt = time.Time{}
                opened = true
            }
        }
    }
    if !opened {
        return false
    }
    return db.GuildSettingsSet(guild, settings) == nil
}

// ClosePoll changes the state of a poll to closed
func ClosePoll(guild, pollID string, msg *discordgo.Message) bool {
    settings := db.GuildSettingsGetCached(guild)
    closed := false
    for i, p := range settings.Polls {
        // If its the poll we're looking for
        if p.ID == pollID {
            // If this user is allowed to do this
            if msg.Author.ID == p.CreatedBy || IsAdmin(msg) {
                if !p.Open {
                    return false
                }
                settings.Polls[i].Open = false
                settings.Polls[i].ClosedAt = time.Now()
                closed = true
            }
        }
    }
    if !closed {
        return false
    }
    return db.GuildSettingsSet(guild, settings) == nil
}

// VotePollIfItsOne func
func VotePollIfItsOne(guild string, r *discordgo.MessageReaction) bool {
    settings := db.GuildSettingsGetCached(guild)
    voted := false
    // See if msg is a poll
    for i, p := range settings.Polls {
        // Skip non matching polls
        if p.ID != r.MessageID {
            continue
        }
        // Can't vote if its closed
        if !p.Open {
            return false
        }
        fieldID := emojis.ToNumber(r.Emoji.APIName())
        emojiAllowed := false
        for _, pf := range p.Fields {
            if fieldID == pf.ID {
                emojiAllowed = true
                break
            }
        }
        // Bad emoji
        if !emojiAllowed {
            session := cache.GetSession()
            session.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), r.UserID)
            return false
        }
        // Check if user voted
        for _, participant := range p.Participants {
            if participant.ID == r.UserID {
                session := cache.GetSession()
                session.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), r.UserID)
                return false
            }
        }
        // Search for the field
        for j, f := range p.Fields {
            // Found the field
            if f.ID == fieldID {
                settings.Polls[i].Participants = append(settings.Polls[i].Participants, models.Participant{
                    ID:      r.UserID,
                    FieldID: fieldID,
                })
                settings.Polls[i].TotalParticipants++
                settings.Polls[i].TotalVotes++
                settings.Polls[i].Fields[j].Votes++
                voted = true
                break
            }
        }
        break
    }
    if !voted {
        return false
    }
    return db.GuildSettingsSet(guild, settings) == nil
}

// GetListEmbedMsg returns an embed msg with the first 5 polls
func GetListEmbedMsg(guild string) *discordgo.MessageEmbed {
    settings := db.GuildSettingsGetCached(guild)
    notShowing := 0
    lPolls := len(settings.Polls)
    if lPolls > 5 {
        notShowing = lPolls - 5
    }
    mfields := []*discordgo.MessageEmbedField{}
    for _, p := range settings.Polls {
        Status := "CLOSED"
        if p.Open {
            Status = "OPEN"
        }
        mfields = append(mfields, &discordgo.MessageEmbedField{
            Name:   p.Title,
            Value:  fmt.Sprintf("Poll ID: %v | Status: %v", p.ID, Status),
            Inline: false,
        },
            &discordgo.MessageEmbedField{
                Name:   "Votes",
                Value:  fmt.Sprintf("%v", p.TotalVotes),
                Inline: true,
            },
            &discordgo.MessageEmbedField{
                Name:   "Participants",
                Value:  fmt.Sprintf("%v", p.TotalParticipants),
                Inline: true,
            })
    }
    if lPolls > 5 {
        mfields = mfields[:15]
    }
    if lPolls == 0 {
        mfields = append(mfields, &discordgo.MessageEmbedField{
            Name: "No polls created",
        })
    }

    embed := &discordgo.MessageEmbed{
        Title:  "Poll list",
        Fields: mfields,
        Color:  0x0FADED,
    }
    if notShowing != 0 {
        embed.Footer = &discordgo.MessageEmbedFooter{
            Text: fmt.Sprintf("Not showing: %v", notShowing),
        }
    }
    return embed

}

// PollTotalVotes returns the total votes for a poll
func PollTotalVotes(guild, pollID string) int64 {
    settings := db.GuildSettingsGetCached(guild)
    for _, p := range settings.Polls {
        if p.ID == pollID {
            return p.TotalVotes
        }
    }
    return 0
}

// PollTotalParticipants returns the total participants for a poll
func PollTotalParticipants(guild, pollID string) int64 {
    settings := db.GuildSettingsGetCached(guild)
    for _, p := range settings.Polls {
        if p.ID == pollID {
            return p.TotalParticipants
        }
    }
    return 0
}

// PollCount returns the number of polls currently has
func PollCount(guild string) int64 {
    return int64(len(db.GuildSettingsGetCached(guild).Polls))
}
