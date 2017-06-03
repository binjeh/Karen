package main

import (
    "code.lukas.moe/x/karen/src/logger"
    "github.com/bwmarrin/discordgo"
    "time"
)

// Automatically leaves guilds that are not registered beta testers
func autoLeaver(session *discordgo.Session, betaGuilds []interface{}) {
    // Print enabled guilds
    for _, betaGuild := range betaGuilds {
        logger.INFO.L("[BETA] Enabled " + betaGuild.(string))
    }

    // Endless loop that checks for unpermitted usage
    for {
        for _, guild := range session.State.Guilds {
            match := false

            for _, betaGuild := range betaGuilds {
                if guild.ID == betaGuild.(string) {
                    match = true
                    break
                }
            }

            if !match {
                logger.WARNING.L("Leaving guild " + guild.ID + " (" + guild.Name + ") because it didn't apply for the beta")
                session.GuildLeave(guild.ID)
            }
        }

        time.Sleep(10 * time.Second)
    }
}
