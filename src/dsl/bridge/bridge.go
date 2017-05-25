package dsl_bridge

import (
    "sync"
)

var (
    scripts     = []Script{}
    scriptMutex = sync.RWMutex{}
)

func PushScript(v Script) {
    scriptMutex.Lock()
    defer scriptMutex.Unlock()

    scripts = append(scripts, v)
}

func GetScripts() (*[]Script) {
    scriptMutex.RLock()
    defer scriptMutex.RUnlock()

    return &scripts
}
