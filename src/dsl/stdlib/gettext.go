package dsl_stdlib

import (
    "code.lukas.moe/x/karen/src/helpers"
)

func GetText(str string) string {
    return helpers.GetText(str)
}

func GetTextF(str string, args []interface{}) string {
    return helpers.GetTextF(str, args...)
}
