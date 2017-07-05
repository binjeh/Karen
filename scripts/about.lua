--
-- Karen - A highly efficient, multipurpose Discord bot written in Golang
--
-- Copyright (C) 2015-2017 Lukas Breuer
-- Copyright (C) 2017 Subliminal Apps
--
-- This file is a part of the Karen Discord-Bot Project ("Karen").
--
-- Karen is free software: you can redistribute it and/or modify
-- it under the terms of the GNU Affero General Public License as published by
-- the Free Software Foundation, either version 3 of the License,
-- or (at your option) any later version.
--
-- Karen is distributed in the hope that it will be useful,
-- but WITHOUT ANY WARRANTY; without even the implied warranty of
-- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
--
-- See the GNU Affero General Public License for more details.
-- You should have received a copy of the GNU Affero General Public License
-- along with this program. If not, see <http://www.gnu.org/licenses/>.
--

require("karen").registerReply(
    "about",
    {
        "about",
        "a",
        "info",
        "inf"
    },
    [[
Hi my name is Karen!
I'm a :robot: that will make this Discord Server a better place c:
Here is some information about me:
```
Karen Araragi (阿良々木 火憐, Araragi Karen) is the eldest of Koyomi Araragi's sisters and the older half of
the Tsuganoki 2nd Middle School Fire Sisters (栂の木二中のファイヤーシスターズ, Tsuganoki Ni-chuu no Faiya Shisutazu).

She is a self-proclaimed "hero of justice" who often imitates the personality and
quirks of various characters from tokusatsu series.
Despite this, she is completely uninvolved with the supernatural, until she becomes victim to a certain oddity.
She is the titular protagonist of two arcs: Karen Bee and Karen Ogre. She is also the narrator of Karen Ogre.
```
BTW: I'm :free:, open-source and built using the Go programming language.
Visit me at <http://karen.vc> or <https://git.lukas.moe/sn0w/Karen>
    ]]
)
