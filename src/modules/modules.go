package modules

import (
    //#ifdef(INCLUDE_PLUGINS)
    "code.lukas.moe/x/karen/src/modules/plugins"
    //#endif

    //#ifdef(INCLUDE_TRIGGERS)
    "code.lukas.moe/x/karen/src/modules/triggers"
    //#endif
)

var (
    //#ifdef(INCLUDE_PLUGINS)
    pluginCache  map[string]*Plugin
    //#endif

    //#ifdef(INCLUDE_TRIGGERS)
    triggerCache map[string]*TriggerPlugin
    //#endif

    //#ifdef(INCLUDE_PLUGINS)
    PluginList = []Plugin{
        &plugins.About{},
        &plugins.Announcement{},
        &plugins.Avatar{},
        &plugins.Calc{},
        &plugins.Changelog{},
        &plugins.Choice{},
        &plugins.FlipCoin{},
        &plugins.Giphy{},
        &plugins.Google{},
        &plugins.Headpat{},
        &plugins.Leet{},
        &plugins.ListenDotMoe{},
        &plugins.Minecraft{},
        &plugins.Music{},
        &plugins.Osu{},
        &plugins.Ping{},
        &plugins.Poll{},
        &plugins.RandomCat{},
        &plugins.Ratelimit{},
        &plugins.Reminders{},
        &plugins.Roll{},
        &plugins.RPS{},
        &plugins.SelfRoles{},
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

    //#ifdef(INCLUDE_TRIGGERS)
    TriggerPluginList = []TriggerPlugin{
        &triggers.CSS{},
        &triggers.Donate{},
        &triggers.Git{},
        &triggers.EightBall{},
        &triggers.Hi{},
        &triggers.HypeTrain{},
        &triggers.Invite{},
        &triggers.IPTables{},
        &triggers.Lenny{},
        &triggers.Nep{},
        &triggers.Kawaii{},
        &triggers.ReZero{},
        &triggers.Shrug{},
        &triggers.TableFlip{},
        &triggers.Triggered{},
    }
    //#endif
)
