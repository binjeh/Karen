package migrations

import (
    "code.lukas.moe/x/karen/src/config"
)

func m0_create_db() {
    CreateDBIfNotExists(
        config.Get("rethink.db").(string),
    )
}
