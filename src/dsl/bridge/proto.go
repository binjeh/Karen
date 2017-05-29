package dsl_bridge

import "github.com/bwmarrin/discordgo"

type Script interface {
    Name() string
    Listeners() []string
    Action(author *discordgo.User, caller string, content string) string
}

type ProtoScript struct {
    nameFunc      func() (string)
    listenersFunc func() ([]string)
    actionFunc    func(author *discordgo.User, caller string, content string) (string)
}

func (p *ProtoScript) Name() (string) {
    return p.nameFunc()
}

func (p *ProtoScript) Listeners() ([]string) {
    return p.listenersFunc()
}

func (p *ProtoScript) Action(author *discordgo.User, caller string, content string) (string) {
    return p.actionFunc(author, caller, content)
}

func (p *ProtoScript) Attach(
    nameFunc func() (string),
    listenersFunc func() ([]string),
    actionFunc func(author *discordgo.User, caller string, content string) (string),
) (*ProtoScript) {
    p.nameFunc = nameFunc
    p.listenersFunc = listenersFunc
    p.actionFunc = actionFunc
    return p
}
