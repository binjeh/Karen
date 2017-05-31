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
