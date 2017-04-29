package main

import (
    "code.lukas.moe/x/karen/src/logger"
    "github.com/bwmarrin/discordgo"
    "strings"
    "time"
)

// Automatically leaves guilds that are not registered beta testers
func autoLeaver(session *discordgo.Session, betaGuildsContainer []interface{}) {
    // Converts []interface to [][]string
    betaGuilds := [][]string{}
    for _, c := range betaGuildsContainer {
        betaGuilds = append(betaGuilds, strings.Split(c.(string), "|"))
    }

    // Print enabled guilds
    for _, betaGuild := range betaGuilds {
        logger.INFO.L("beta", "[BETA] Enabled "+betaGuild[0]+" ("+betaGuild[1]+") by "+betaGuild[2])
    }

    // Endless loop that checks for unpermitted usage
    for {
        for _, guild := range session.State.Guilds {
            match := false

            for _, betaGuild := range betaGuilds {
                if guild.ID == betaGuild[0] {
                    match = true
                    break
                }
            }

            if !match {
                logger.WARNING.L("beta", "Leaving guild "+guild.ID+" ("+guild.Name+") because it didn't apply for the beta")
                session.GuildLeave(guild.ID)
            }
        }

        time.Sleep(10 * time.Second)
    }
}
