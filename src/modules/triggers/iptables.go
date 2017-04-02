package triggers

import "code.lukas.moe/x/karen/src/helpers"

type IPTables struct{}

func (i *IPTables) Triggers() []string {
    return []string{
        "ipt",
        "iptables",
    }
}

func (i *IPTables) Response(trigger string, content string) string {
    return helpers.GetText("triggers.iptables")
}
