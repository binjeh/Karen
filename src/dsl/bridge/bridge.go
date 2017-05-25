package bridge

import "sync"

var (
    scripts     = map[string]Script{}
    scriptMutex = sync.RWMutex{}
)

func PushScript(k string, v Script) {
    scriptMutex.Lock()
    scripts[k] = v
    scriptMutex.Unlock()
}

func DropScript(k string) {
    scriptMutex.Lock()
    delete(scripts, k)
    scriptMutex.Unlock()
}

func GetScripts() (map[string]Script) {
    return scripts
}
