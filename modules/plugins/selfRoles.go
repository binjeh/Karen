package plugins

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "git.lukas.moe/sn0w/Karen/helpers"
    "git.lukas.moe/sn0w/Karen/models"
    "fmt"
)

// SelfRoles command
type SelfRoles struct {}

// Commands triggers for SelfRoles
func (s *SelfRoles) Commands() []string {
    return []string {
        // role [add, a|remove, rm, del] <role_name>
        "role",
        // roles
        "roles",
        // iam <role_name>
        "iam",
        // iamnot <role_name>
        "iamnot",
    }
}

// Init func
func (s *SelfRoles) Init(session *discordgo.Session) {}

// Action func
func (s *SelfRoles) Action(command, content string, msg *discordgo.Message, session *discordgo.Session) {
    switch command {
    case "role":
        s.role(content, msg, session)
    case "roles":
        s.roles(content, msg, session)
    case "iam":
        s.iam(content, msg, session)
    case "iamnot":
        s.iamnot(content, msg, session)
    }
}

func (s *SelfRoles) role(content string, msg *discordgo.Message, session *discordgo.Session) {
    // If not admin
    if !helpers.IsAdmin(msg) {
        // GTFO! :3 <3
        return
    }
    split := strings.Fields(content)
    if len(split) < 2 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.missing_param"))
        return
    }
    subcommand := split[0]
    roleName := split[1]
    switch subcommand {
    // Add role
    case "add", "a":
        s.addRole(roleName, msg, session)
    // Remove role
    case "remove", "rm", "del":
        s.removeRole(roleName, msg, session)
    // Wrong subcommand
    default:
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.self_roles.wrong_subcommand", subcommand))
        return
    }
}

func (s *SelfRoles) addRole(roleName string, msg *discordgo.Message, session *discordgo.Session) {
    // Get the channel
    channel, err := session.Channel(msg.ChannelID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_add.failure"))
        return
    }
    // Create the role
    role, err := session.GuildRoleCreate(channel.GuildID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_add.failure"))
        return
    }
    // Edit the role
    _, err = session.GuildRoleEdit(channel.GuildID, role.ID, roleName, role.Color, role.Hoist, 0, true)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_add.failure"))
        return
    }
    // Persist to db
    settings := helpers.GuildSettingsGetCached(channel.GuildID)
    settings.Roles = append(settings.Roles, models.Role{ID: role.ID, Name: roleName})
    err = helpers.GuildSettingsSet(channel.GuildID, settings)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_add.failure"))
        return
    }
    // Make senpai notice us
    session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
        Title:       "Added role",
        Description: helpers.GetTextF("plugins.self_roles.role_add.success", roleName),
        Color: 0x0FADED,
    })
}

func (s *SelfRoles) removeRole(roleName string, msg *discordgo.Message, session *discordgo.Session) {
    // Get the channel
    channel, err := session.Channel(msg.ChannelID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_remove.failure"))
        return
    }
    // Get all guild roles
    roles, err := session.GuildRoles(channel.GuildID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_remove.failure"))
        return
    }
    // Find our role
    for _, role := range roles {
        // We found it
        if role.Name == roleName {
            // Delete it
            err = session.GuildRoleDelete(channel.GuildID, role.ID)
            if err != nil {
                session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_remove.failure"))
                return
            }
            // Delete it from db
            settings := helpers.GuildSettingsGetCached(channel.GuildID)
            newRoles := make([]models.Role, len(settings.Roles)-1)
            for i, r := range settings.Roles {
                if r.Name != roleName {
                    newRoles[i] = r
                }
            }
            settings.Roles = newRoles
            err = helpers.GuildSettingsSet(channel.GuildID, settings)
            if err != nil {
                session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.role_remove.failure"))
                return
            }
            break
        }
    }
    // Notify senpai that we did it!!
    session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
        Title:       "Removed role",
        Description: helpers.GetTextF("plugins.self_roles.role_remove.success", roleName),
        Color: 0x0FADED,
    })
}

func (s *SelfRoles) roles(content string, msg *discordgo.Message, session *discordgo.Session) {
    channel, err := session.Channel(msg.ChannelID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.roles.failure"))
        return
    }
    settings := helpers.GuildSettingsGetCached(channel.GuildID)
    embed := &discordgo.MessageEmbed {
        Title: "List of available self-assignable roles",
        Color: 0x0FADED,
    }
    for _, role := range settings.Roles {
        embed.Description += fmt.Sprintf("%s\n", role.Name)
    }
    session.ChannelMessageSendEmbed(msg.ChannelID, embed)
}

func (s *SelfRoles) iam(content string, msg *discordgo.Message, session *discordgo.Session) {
    split := strings.Fields(content)
    if len(split) < 1 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.missing_param"))
        return
    }
    roleName := split[0]
    // Get the channel
    channel, err := session.Channel(msg.ChannelID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.self_roles.iam.failure", msg.Author.ID))
        return
    }
    // Get the role
    role, found := helpers.GuildRoleByName(channel.GuildID, roleName)
    if !found {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.wrong_subcommand"))
        return
    }
    // Add the role to the user
    err = session.GuildMemberRoleAdd(channel.GuildID, msg.Author.ID, role.ID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.self_roles.iam.failure", msg.Author.ID))
        return
    }
    // Profit :LeftShark:
    session.ChannelMessageSendEmbedWithMessage(msg.ChannelID, fmt.Sprintf("<@%s>", msg.Author.ID), &discordgo.MessageEmbed{
        Title: "Joined role",
        Description: helpers.GetTextF("plugins.self_roles.iam.success", role.Name),
        Color: 0x0FADED,
    })
}

func (s *SelfRoles) iamnot(content string, msg *discordgo.Message, session *discordgo.Session) {
    split := strings.Fields(content)
    if len(split) < 1 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.missing_param"))
        return
    }
    roleName := split[0]
    // Get the channel
    channel, err := session.Channel(msg.ChannelID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.self_roles.iamnot.failure", msg.Author.ID))
        return
    }
    // Get the role
    role, found := helpers.GuildRoleByName(channel.GuildID, roleName)
    if !found {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.self_roles.wrong_subcommand"))
        return
    }
    // Remove the role from the user
    err = session.GuildMemberRoleRemove(channel.GuildID, msg.Author.ID, role.ID)
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.self_roles.iamnot.failure", msg.Author.ID))
        return
    }
    // Profit :LeftShark:
    session.ChannelMessageSendEmbedWithMessage(msg.ChannelID, fmt.Sprintf("<@%s>", msg.Author.ID), &discordgo.MessageEmbed{
        Title: "Left role",
        Description: helpers.GetTextF("plugins.self_roles.iamnot.success", role.Name),
        Color: 0x0FADED,
    })
}
