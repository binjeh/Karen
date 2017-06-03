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

package modules

import (
    //#ifeq EXCLUDE_PLUGINS 0
    "code.lukas.moe/x/karen/src/modules/plugins"
    //#endif

    //#ifeq EXCLUDE_SCRIPTING 0
    "code.lukas.moe/x/karen/src/dsl/bridge"
    //#endif
)

var (
    //#ifeq EXCLUDE_PLUGINS 0
    pluginCache map[string]Plugin
    //#endif

    //#ifeq EXCLUDE_SCRIPTING 0
    scriptCache map[string]dsl_bridge.Script
    //#endif

    //#ifeq EXCLUDE_PLUGINS 0
    PluginList = []Plugin{
        //&plugins.Announcement{},
        &plugins.Avatar{},
        &plugins.Calc{},
        &plugins.Changelog{},
        &plugins.Choice{},
        &plugins.Enlarge{},
        &plugins.FlipCoin{},
        &plugins.Giphy{},
        &plugins.Google{},
        &plugins.Headpat{},
        &plugins.Leet{},
        &plugins.Minecraft{},
        &plugins.Osu{},
        &plugins.Ping{},
        //&plugins.Poll{},
        &plugins.RandomCat{},
        &plugins.Ratelimit{},
        &plugins.Reminders{},
        &plugins.Roll{},
        &plugins.RPS{},
        &plugins.SelfRoles{},
        &plugins.Spoiler{},
        &plugins.Stats{},
        &plugins.Stone{},
        &plugins.Support{},
        &plugins.Toggle{},
        //&plugins.Translator{},
        &plugins.Timezone{},
        &plugins.Uptime{},
        &plugins.UrbanDict{},
        &plugins.Weather{},
        &plugins.WhoIs{},
        &plugins.XKCD{},
    }
    //#endif
)
