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

package cache

import (
    "time"
)

type ItemGetter func(id string) (interface{}, error)

func GetOrRequest(id string, cb ItemGetter) (item interface{}, e error) {
    dirty := false

    // Retrieve if not existant yet
    mutex.RLock()
    ok := false
    item, ok = objects[id]
    mutex.RUnlock()
    if !ok {
        dirty = true
        item, e = cb(id)
    }

    // Check if there is a timeout
    mutex.RLock()
    meta := objectMeta[id]
    mutex.RUnlock()
    if time.Now().Unix()-meta > timeout {
        dirty = true
        item, e = cb(id)
    }

    // Update the entry
    if dirty {
        mutex.Lock()
        objects[id] = item
        mutex.Unlock()
    }

    // Return data
    return
}
