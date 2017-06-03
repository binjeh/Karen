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

package triggers

import "code.lukas.moe/x/karen/src/helpers"

type EightBall struct{}

func (e *EightBall) Triggers() []string {
    return []string{
        "8ball",
        "8",
    }
}

func (e *EightBall) Response(trigger string, content string) string {
    if len(content) < 3 {
        return helpers.GetText("triggers.8ball.ask_a_question")
    }

    return ":8ball: " + helpers.GetText("triggers.8ball")
}
