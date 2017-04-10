package models

// Config is a struct describing all config options a guild may set
type Config struct {
    Id  string `rethink:"id,omitempty"`
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
