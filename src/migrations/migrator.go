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

package migrations

import (
    "code.lukas.moe/x/karen/src/helpers"
    "code.lukas.moe/x/karen/src/logger"
    "reflect"
    "runtime"
)

var migrations = []helpers.Callback{
    m0_create_db,
    m1_create_table_guild_config,
    m2_create_table_reminders,
    m3_create_table_music,
}

// Run executes all registered migrations
func Run() {
    logger.BOOT.L("Running migrations...")
    for _, migration := range migrations {
        migrationName := runtime.FuncForPC(
            reflect.ValueOf(migration).Pointer(),
        ).Name()

        logger.BOOT.L("Running "+migrationName)
        migration()
    }

    logger.BOOT.L("Migrations finished!")
}
