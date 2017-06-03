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

// Config is a struct describing all config options a guild may set
type Config struct {
    Id string `rethink:"id,omitempty"`
    // Guild contains the guild ID
    Guild string `rethink:"guild"`

    Prefix string `rethink:"prefix"`

    CleanupEnabled bool `rethink:"cleanup_enabled"`

    AnnouncementsEnabled bool `rethink:"announcements_enabled"`
    // AnnouncementsChannel stores the channel ID
    AnnouncementsChannel string `rethink:"announcements_channel"`

    JoinNotificationsEnabled bool   `rethink:"join_notifications_enabled"`
    JoinNotificationsChannel string `rethink:"join_notifications_channel"`
    JoinNotificationText     string `rethink:"join_notification_text"`

    LeaveNotificationsEnabled bool   `rethink:"leave_notifications_enabled"`
    LeaveNotificationsChannel string `rethink:"leave_notifications_channel"`
    LeaveNotificationText     string `rethink:"leave_notification_text"`

    // Roles contains the available self-assignable
    // roles on this guild
    Roles []Role `rethink:"roles"`

    // Polls contains all open polls for the guild,
    // closed polls are also stored but will be auto-deleted
    // one day after its state changes to closed
    Polls []Poll `rethink:"polls"`
}

// Role struct
type Role struct {
    ID   string `rethink:"id"`
    Name string `rethink:"name"`
}

// Default is a helper for generating default config values
func (c Config) Default(guild string) Config {
    return Config{
        Guild: guild,

        Prefix: "%",

        CleanupEnabled: false,

        AnnouncementsEnabled: false,
        AnnouncementsChannel: "",

        JoinNotificationsEnabled: false,
        JoinNotificationsChannel: "",
        JoinNotificationText:     "",

        LeaveNotificationsEnabled: false,
        LeaveNotificationsChannel: "",
        LeaveNotificationText:     "",
    }
}
