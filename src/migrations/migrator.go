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
