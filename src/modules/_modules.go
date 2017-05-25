package modules

import (
    //#ifndef EXCLUDE_PLUGINS
    "code.lukas.moe/x/karen/src/modules/plugins"
    //#endif

    //#ifndef EXCLUDE_SCRIPTING
    "code.lukas.moe/x/karen/src/dsl/bridge"
    //#endif
)

var (
    //#ifndef EXCLUDE_PLUGINS
    pluginCache map[string]Plugin
    //#endif

    //#ifndef EXCLUDE_SCRIPTING
    scriptCache map[string]dsl_bridge.Script
    //#endif

    //#ifndef EXCLUDE_PLUGINS
    PluginList = []Plugin{
        //#ifndef EXCLUDE_MUSIC
        &plugins.Music{},
        //#endif

        //#ifndef EXCLUDE_RADIO
        &plugins.ListenDotMoe{},
        //#endif

        &plugins.About{},
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
