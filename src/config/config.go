package config

import (
    "sync"
    "github.com/pelletier/go-toml"
)

//noinspection GoNameStartsWithPackageName
const (
    CONFIG_FILE = "config.toml"
)

var (
    mutex = &sync.Mutex{}
    tree  *toml.TomlTree
)

func getTree() *toml.TomlTree {
    mutex.Lock()
    defer mutex.Unlock()

    if tree == nil {
        var err error

        tree, err = toml.LoadFile(CONFIG_FILE)
        if err != nil {
            panic(err)
        }
    }

    return tree
}

func Get(path string) interface{} {
    return getTree().Get(path)
}

func GetDefault(path string, def interface{}) interface{} {
    return getTree().GetDefault(path, def)
}
