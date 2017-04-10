package migrations

import "code.lukas.moe/x/karen/src/helpers"

func m0_create_db() {
    CreateDBIfNotExists(
        helpers.GetConfig().Path("rethink.db").Data().(string),
    )
}
