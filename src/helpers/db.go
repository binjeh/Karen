/*
 * Karen - A highly efficient, multipurpose Discord bot written in Golang
 *
 * Copyright (C) 2015-2017 Lukas Breuer
 * Copyright (C) 2017 Subliminal Apps
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
    "code.lukas.moe/x/karen/src/cache"
    Logger "code.lukas.moe/x/karen/src/logger"
    "code.lukas.moe/x/karen/src/models"
    "github.com/getsentry/raven-go"
    rethink "github.com/gorethink/gorethink"
    "sync"
    "time"
)

var (
    dbSession *rethink.Session

    guildSettingsCache map[string]models.Config
    cacheMutex         sync.RWMutex
)

// ConnectDB connects to rethink and stores the session
func ConnectDB(url string, db string) {
    Logger.INFO.L("Connecting to " + url)

    rethink.SetTags("rethink", "json")

    session, err := rethink.Connect(rethink.ConnectOpts{
        Address:  url,
        Database: db,
    })

    if err != nil {
        Logger.ERROR.L(err.Error())
        panic(err)
    }

    dbSession = session

    cacheMutex.Lock()
    guildSettingsCache = make(map[string]models.Config)
    cacheMutex.Unlock()

    Logger.INFO.L("Connected!")
}

// GetDB is a simple getter for the rethink session.
// Might receive some singleton-like lazy-creation later
func GetDB() *rethink.Session {
    return dbSession
}

// GuildSettingsSet writes all $config into the db
func GuildSettingsSet(guild string, config models.Config) error {
    // Check if an config object exists
    var settings models.Config

    cursor, err := rethink.Table("guild_configs").Filter(map[string]interface{}{"guild": guild}).Run(GetDB())
    defer cursor.Close()

    if err != nil {
        return err
    }

    err = cursor.One(&settings)

    switch err {
    // Insert
    case rethink.ErrEmptyResult:
        _, err = rethink.Table("guild_configs").Insert(config).RunWrite(GetDB())
        break

        // Update
    case nil:
        _, err = rethink.Table("guild_configs").Filter(
            map[string]interface{}{"guild": guild},
        ).Update(config).RunWrite(GetDB())
        break

    default:
        panic(err)
    }

    // Update cache
    cacheMutex.Lock()
    guildSettingsCache[guild] = config
    cacheMutex.Unlock()

    return err
}

// GuildSettingsGet returns all config values for the guild or a default object
func GuildSettingsGet(guild string) (models.Config, error) {
    var settings models.Config
    var cursor *rethink.Cursor
    var err error

    cursor, err = rethink.Table("guild_configs").Filter(map[string]interface{}{"guild": guild}).Run(GetDB())
    defer cursor.Close()

    if err != nil {
        return settings, err
    }

    err = cursor.One(&settings)

    switch err {
    case rethink.ErrEmptyResult:
        settings = models.Config{}.Default(guild)
        return settings, nil
    default:
        return settings, err
    }
}

func GuildSettingsGetCached(id string) (models.Config) {
    cacheMutex.RLock()
    defer cacheMutex.RUnlock()

    return guildSettingsCache[id]
}

// GetPrefixForServer gets the prefix for $guild
func GetPrefixForServer(guild string) (string) {
    return GuildSettingsGetCached(guild).Prefix
}

// SetPrefixForServer sets the prefix for $guild to $prefix
func SetPrefixForServer(guild string, prefix string) error {
    settings := GuildSettingsGetCached(guild)

    settings.Prefix = prefix

    return GuildSettingsSet(guild, settings)
}

func GuildSettingsUpdater() {
    for {
        for _, guild := range cache.GetSession().State.Guilds {
            settings, e := GuildSettingsGet(guild.ID)
            if e != nil {
                raven.CaptureError(e, map[string]string{})
                continue
            }

            cacheMutex.Lock()
            guildSettingsCache[guild.ID] = settings
            cacheMutex.Unlock()
        }

        time.Sleep(15 * time.Second)
    }
}
