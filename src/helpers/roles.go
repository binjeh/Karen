package helpers

import (
    "code.lukas.moe/x/karen/src/models"
)

func GuildRoleByName(guild, name string) (models.Role, bool) {
    settings := GuildSettingsGetCached(guild)
    for _, role := range settings.Roles {
        if role.Name == name {
            return role, true
        }
    }
    return models.Role{}, false
}
