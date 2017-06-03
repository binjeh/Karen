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

// Except.go: Contains functions to make handling panics less PITA
package helpers

import (
    "fmt"
    "code.lukas.moe/x/karen/src/cache"
    "github.com/bwmarrin/discordgo"
    "github.com/getsentry/raven-go"
    "reflect"
    "runtime"
    "strconv"
    "github.com/davecgh/go-spew/spew"
)

// RecoverDiscord recover()s and sends a message to discord
func RecoverDiscord(msg *discordgo.Message) {
    err := recover()
    if err != nil {
        spew.Dump(err)
        SendError(msg, err)
    }
}

// Recover recover()s and prints the error to console
func Recover() {
    err := recover()
    if err != nil {
        fmt.Printf("%#v\n", err)
    }
}

// SoftRelax is a softer form of Relax()
// Calls a callback instead of panicking
func SoftRelax(err error, cb Callback) {
    if err != nil {
        cb()
    }
}

// Relax is a helper to reduce if-checks if panicking is allowed
// If $err is nil this is a no-op. Panics otherwise.
func Relax(err error) {
    if err != nil {
        panic(err)
    }
}

// RelaxAssertEqual panics if a is not b
func RelaxAssertEqual(a interface{}, b interface{}, err error) {
    if !reflect.DeepEqual(a, b) {
        Relax(err)
    }
}

// RelaxAssertUnequal panics if a is b
func RelaxAssertUnequal(a interface{}, b interface{}, err error) {
    if reflect.DeepEqual(a, b) {
        Relax(err)
    }
}

// SendError Takes an error and sends it to discord and sentry.io
func SendError(msg *discordgo.Message, err interface{}) {
    if DEBUG_MODE == true {
        buf := make([]byte, 1<<16)
        stackSize := runtime.Stack(buf, false)

        cache.GetSession().ChannelMessageSend(
            msg.ChannelID,
            "Error :frowning:\n0xFADED#3237 has been notified.\n```\n"+fmt.Sprintf("%#v\n", err)+fmt.Sprintf("%s\n", string(buf[0:stackSize]))+"\n```\nhttp://i.imgur.com/FcV2n4X.jpg",
        )
    } else {
        cache.GetSession().ChannelMessageSend(
            msg.ChannelID,
            "Error :frowning:\n0xFADED#3237 has been notified.\n```\n"+fmt.Sprintf("%#v", err)+"\n```\nhttp://i.imgur.com/FcV2n4X.jpg",
        )
    }

    raven.SetUserContext(&raven.User{
        ID:       msg.ID,
        Username: msg.Author.Username + "#" + msg.Author.Discriminator,
    })

    raven.CaptureError(fmt.Errorf("%#v", err), map[string]string{
        "ChannelID":       msg.ChannelID,
        "Content":         msg.Content,
        "Timestamp":       string(msg.Timestamp),
        "TTS":             strconv.FormatBool(msg.Tts),
        "MentionEveryone": strconv.FormatBool(msg.MentionEveryone),
        "IsBot":           strconv.FormatBool(msg.Author.Bot),
    })
}
