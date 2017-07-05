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
    "code.lukas.moe/x/karen/src/cache"
    "code.lukas.moe/x/karen/src/db"
    "code.lukas.moe/x/karen/src/except"
    "code.lukas.moe/x/karen/src/i18n"
    "code.lukas.moe/x/karen/src/logger"
    "fmt"
    "github.com/bwmarrin/discordgo"
    rethink "github.com/gorethink/gorethink"
    "github.com/olebedev/when"
    "github.com/olebedev/when/rules/common"
    "github.com/olebedev/when/rules/en"
    "strings"
    "time"
)

// Reminders command
type Reminders struct {
    parser *when.Parser
}

// DB_Reminders struct
type DB_Reminders struct {
    Id     string        `gorethink:"id,omitempty"`
    UserID string        `gorethink:"userid"`
    // Timezone is stored in the format specified by
    // the IANA Timezone db, as well as
    // time.LoadLocation()/time.Time.In()
    Timezone  string        `gorethink:"timezone"`
    Reminders []DB_Reminder `gorethink:"reminders"`
}

// DB_Reminder struct
type DB_Reminder struct {
    Message   string `gorethink:"message"`
    ChannelID string `gorethink:"channelID"`
    GuildID   string `gorethink:"guildID"`
    Timestamp int64  `gorethink:"timestamp"`
}

// Commands that triggers reminders
func (r *Reminders) Commands() []string {
    return []string{
        "remind",
        "rm",
        "reminders",
        "rms",
    }
}

// Init the reminders loop
func (r *Reminders) Init(session *discordgo.Session) {
    r.parser = when.New(nil)
    r.parser.Add(en.All...)
    r.parser.Add(common.All...)

    go func() {
        defer except.Recover()

        for {
            var reminderBucket []DB_Reminders
            cursor, err := rethink.Table("reminders").Run(db.GetSession())
            except.Handle(err)

            err = cursor.All(&reminderBucket)
            except.Handle(err)

            for _, reminders := range reminderBucket {
                changes := false

                // Downward loop for in-loop element removal
                for idx := len(reminders.Reminders) - 1; idx >= 0; idx-- {
                    reminder := reminders.Reminders[idx]

                    loc, err := time.LoadLocation(reminders.Timezone)
                    if err != nil {
                        loc = time.UTC
                    }

                    if reminder.Timestamp <= time.Now().In(loc).Unix() {
                        user, err := session.User(reminders.UserID)
                        if err != nil {
                            continue
                        }
                        embed := &discordgo.MessageEmbed{
                            Title:       ":alarm_clock: Ring! Ring!",
                            Description: reminder.Message,
                            Color:       0x0FADED,
                            Footer: &discordgo.MessageEmbedFooter{
                                Text: fmt.Sprintf("Reminder for: %s", user.Username),
                            },
                        }

                        _, err = session.ChannelMessageSendComplex(reminder.ChannelID, &discordgo.MessageSend{
                            Tts:     false,
                            Content: fmt.Sprintf("<@%s>", reminders.UserID),
                            Embed:   embed,
                        })
                        if err != nil {
                            continue
                        }

                        reminders.Reminders = append(reminders.Reminders[:idx], reminders.Reminders[idx+1:]...)
                        changes = true
                    }
                }

                if changes {
                    r.setReminders(reminders.UserID, reminders)
                }
            }

            time.Sleep(10 * time.Second)
        }
    }()

    logger.PLUGIN.L("Started reminder loop (10s)")
}

// Action executes the reminders command
func (r *Reminders) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    switch command {
    case "rm", "remind":
        channel, err := cache.Channel(msg.ChannelID)
        except.Handle(err)

        parts := strings.Fields(content)

        if len(parts) < 3 {
            session.ChannelMessageSend(msg.ChannelID, ":x: Please check if the format is correct")
            return
        }

        reminders := r.getReminders(msg.Author.ID)

        loc, err := time.LoadLocation(reminders.Timezone)
        if err != nil {
            loc = time.UTC
        }

        result, err := r.parser.Parse(content, time.Now().In(loc))
        except.Handle(err)
        if result == nil {
            session.ChannelMessageSend(msg.ChannelID, ":x: Please check if the format is correct")
            return
        }

        reminders.Reminders = append(reminders.Reminders, DB_Reminder{
            Message:   strings.Replace(content, result.Text, "", 1),
            ChannelID: channel.ID,
            GuildID:   channel.GuildID,
            Timestamp: result.Time.Unix(),
        })
        reminders.UserID = msg.Author.ID
        r.setReminders(msg.Author.ID, reminders)

        if reminders.Timezone == "" {
            session.ChannelMessageSend(msg.ChannelID, i18n.GetText("plugins.reminders.empty_timezone"))
        } else {
            session.ChannelMessageSend(msg.ChannelID, "Ok I'll remind you :ok_hand:")
        }

    case "rms", "reminders":
        reminders := r.getReminders(msg.Author.ID)
        embedFields := []*discordgo.MessageEmbedField{}

        loc, err := time.LoadLocation(reminders.Timezone)
        if err != nil {
            loc = time.UTC
        }

        for _, reminder := range reminders.Reminders {
            ts := time.Unix(reminder.Timestamp, 0).In(loc)
            channel := "?"
            guild := "?"

            chanRef, err := session.Channel(reminder.ChannelID)
            if err == nil {
                channel = chanRef.Name
            }

            guildRef, err := session.Guild(reminder.GuildID)
            if err == nil {
                guild = guildRef.Name
            }

            embedFields = append(embedFields, &discordgo.MessageEmbedField{
                Inline: false,
                Name:   reminder.Message,
                Value:  fmt.Sprintf("At %s in #%s of %s", ts, channel, guild),
            })
        }

        if len(embedFields) == 0 {
            session.ChannelMessageSend(msg.ChannelID, i18n.GetText("plugins.reminders.empty"))
            return
        }

        session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
            Title:  "Pending reminders",
            Fields: embedFields,
            Color:  0x0FADED,
            Footer: &discordgo.MessageEmbedFooter{
                Text: fmt.Sprintf("Timezone: %s", reminders.Timezone),
            },
        })
    }
}

func (r Reminders) getReminders(uid string) DB_Reminders {
    var reminderBucket DB_Reminders
    listCursor, err := rethink.Table("reminders").Filter(
        rethink.Row.Field("userid").Eq(uid),
    ).Run(db.GetSession())
    defer listCursor.Close()
    err = listCursor.One(&reminderBucket)

    // If user has no DB entries create an empty document
    if err == rethink.ErrEmptyResult {
        _, e := rethink.Table("reminders").Insert(DB_Reminders{
            UserID:    uid,
            Reminders: make([]DB_Reminder, 0),
        }).RunWrite(db.GetSession())

        // If the creation was successful read the document
        if e != nil {
            panic(e)
        } else {
            return reminderBucket
        }
    } else if err != nil {
        panic(err)
    }

    return reminderBucket
}

func (r Reminders) setReminders(uid string, reminders DB_Reminders) {
    _, err := rethink.Table("reminders").Update(reminders).RunWrite(db.GetSession())
    except.Handle(err)
}
