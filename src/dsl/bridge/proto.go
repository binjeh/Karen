/*
 * Karen - A highly efficient, multipurpose Discord bot written in Golang
 *
 * Copyright (C) 2015-2017 Lukas Breuer
 * Copyright (C) 2017 Subliminal Apps
 *
 * This file is a part of the Karen Discord-Bot Project ("Karen").
 *
 * Karen is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or (at your option) any later version.
 *
 * Karen is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 * See the GNU Affero General Public License for more details.
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

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
