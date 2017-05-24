package config

import (
    "github.com/Jeffail/gabs"
    "sync"
)

var (
    mutex     = &sync.Mutex{}
    container *gabs.Container
)

// Return the bare config container.
// Use the returned mutex to block reads while you're working with the gabs container.
func GetRaw() (*gabs.Container, *sync.Mutex) {
    return container, mutex
}

// Used to retrieve a value using dot notation.
// Example: auth.discord.user.token
func Get(path string) interface{} {
    mutex.Lock()
    defer mutex.Unlock()

    if container == nil {
        var err error

        container, err = gabs.ParseJSONFile("config.json")
        if err != nil {
            panic(err)
        }
    }

    if !container.ExistsP(path) {
        return nil
    }

    return container.Path(path).Data()
}
