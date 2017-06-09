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

package migrations

import (
    "code.lukas.moe/x/karen/src/db"
    "code.lukas.moe/x/karen/src/except"
    rethink "github.com/gorethink/gorethink"
)

// CreateTableIfNotExists (works like the mysql call)
func CreateTableIfNotExists(tableName string) {
    cursor, err := rethink.TableList().Run(db.GetSession())
    except.Handle(err)
    defer cursor.Close()

    tableExists := false

    var row string
    for cursor.Next(&row) {
        if row == tableName {
            tableExists = true
            break
        }
    }

    if !tableExists {
        _, err := rethink.TableCreate(tableName).Run(db.GetSession())
        except.Handle(err)
    }
}

// CreateDBIfNotExists (works like the mysql call)
func CreateDBIfNotExists(dbName string) {
    cursor, err := rethink.DBList().Run(db.GetSession())
    except.Handle(err)
    defer cursor.Close()

    dbExists := false

    var row string
    for cursor.Next(&row) {
        if row == dbName {
            dbExists = true
            break
        }
    }

    if !dbExists {
        _, err := rethink.DBCreate(dbName).Run(db.GetSession())
        except.Handle(err)
    }
}
