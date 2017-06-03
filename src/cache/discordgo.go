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
