package helpers

import (
    "fmt"
    "github.com/Jeffail/gabs"
    "math/rand"
    "strings"
    "code.lukas.moe/x/karen/src/assets"
)

var translations *gabs.Container

func LoadTranslations() {
    jsonFile, err := assets.Asset("assets/i18n.json")
    Relax(err)

    json, err := gabs.ParseJSON(jsonFile)
    Relax(err)

    translations = json
}

func GetText(id string) string {
    if !translations.ExistsP(id) {
        return id
    }

    item := translations.Path(id)

    // If this is an object return __
    if strings.Contains(item.String(), "{") {
        item = item.Path("__")
    }

    // If this is an array return a random item
    if strings.Contains(item.String(), "[") {
        arr := item.Data().([]interface{})
        return arr[rand.Intn(len(arr))].(string)
    }

    return item.Data().(string)
}

func GetTextF(id string, replacements ...interface{}) string {
    return fmt.Sprintf(GetText(id), replacements...)
}
