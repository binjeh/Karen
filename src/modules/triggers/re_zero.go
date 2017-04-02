package triggers

import "code.lukas.moe/x/karen/src/helpers"

type ReZero struct{}

func (r *ReZero) Triggers() []string {
    return []string{
        "rem",
        "re:zero",
        "rezero",
    }
}

func (r *ReZero) Response(trigger string, content string) string {
    return helpers.GetText("triggers.re_zero.link")
}
