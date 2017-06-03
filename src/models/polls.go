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

package models

import (
    "time"
)

// Poll struct
type Poll struct {
    // ID is the ID of the msg that created the current Poll
    ID string `rethink:"id"`
    // The ID of the channel where the poll was created
    ChannelID string `rethink:"channel_id"`
    // Title of the Poll
    Title string `rethink:"title"`
    // Fields of the Poll
    Fields []PollField `rethink:"fields"`
    // Open shows the current state for the Poll
    Open bool `rethink:"active"`
    // The users that already voted, a Participant can't
    // vote more than once nor change the field they voted for
    Participants []Participant `rethink:"participants"`
    // The same as calling len(Participants) but we
    // don't have to call it every time
    TotalParticipants int64 `rethink:"total_participants"`
    // The total number of votes across all fields
    TotalVotes int64 `rethink:"total_votes"`
    // The time when the Poll was created
    CreatedAt time.Time `rethink:"created_at"`
    // The time when the Poll state changed to inactive
    ClosedAt time.Time `rethink:"closed_at"`
    // CreatedBy contains the user ID that created the Poll.
    // This user will be the only one that will be able to
    // close this poll apart from the guild admins
    CreatedBy string `rethink:"created_by"`
}

// PollField is a field for a Poll
type PollField struct {
    ID    int    `rethink:"id"`
    Title string `rethink:"name"`
    Votes int64  `rethink:"votes"`
}

// Participant represents an user that already voted
type Participant struct {
    // ID is the user.ID
    ID string `rethink:"id"`
    // The ID of the field the current Participant voted for
    FieldID int `rethink:"field_id"`
}
