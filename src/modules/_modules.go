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
        //#ifeq EXCLUDE_MUSIC 0
        &plugins.Music{},
        //#endif

        //#ifeq EXCLUDE_RADIO 0
        &plugins.ListenDotMoe{},
        //#endif

        //&plugins.Announcement{},
        &plugins.Avatar{},
        &plugins.Calc{},
        &plugins.Changelog{},
        &plugins.Choice{},
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
