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

package helpers

import (
    "github.com/bwmarrin/discordgo"
    "sync"
)

// UNASSIGNED is an alias for a guild that is not connected to VC right now
const UNASSIGNED = "___UNASSIGNED___"

var (
    // connections maps guild ids to occupier-ids
    connections = map[string]string{}

    // mutex locks connections when writing
    mutex = &sync.Mutex{}
)

// VoiceIsOccupied checks if any plugin blocks further voice connections
func VoiceIsOccupied(guild string) bool {
    if _, ok := connections[guild]; !ok {
        mutex.Lock()
        connections[guild] = UNASSIGNED
        mutex.Unlock()
    }

    return connections[guild] != UNASSIGNED
}

// VoiceIsOccupied checks if a given plugin blocks further voice connections
func VoiceIsOccupiedBy(guild string, id string) bool {
    return connections[guild] == id
}

// VoiceIsOccupied checks if a given plugin blocks further voice connections or if the voice is unassigned
func VoiceIsFreeOrOccupiedBy(guild string, id string) bool {
    return !VoiceIsOccupied(guild) || VoiceIsOccupiedBy(guild, id)
}

// VoiceOccupy marks a guild as occupied. Returns true if occupation was successful. False otherwise.
// Example usage:
// lock := helpers.VoiceOccupy(guild.ID, "music")
// helpers.RelaxAssertEqual(lock, true, nil)
func VoiceOccupy(guild string, reason string) bool {
    if VoiceIsOccupied(guild) {
        return false
    }

    mutex.Lock()
    connections[guild] = reason
    mutex.Unlock()

    return true
}

// VoiceFree marks a guild as open for new voice connections
func VoiceFree(guild string) {
    if !VoiceIsOccupied(guild) {
        return
    }

    mutex.Lock()
    connections[guild] = UNASSIGNED
    mutex.Unlock()
}

func VoiceSendStatus(channel string, guild string, session *discordgo.Session) {
    session.ChannelMessageSend(
        channel,
        "Sorry but voice is currently occupied by `"+connections[guild]+"`.\nCall leave for that plugin first.",
    )
}
