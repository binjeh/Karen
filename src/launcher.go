package main

import (
    "code.lukas.moe/x/karen/src/helpers"
    Logger "code.lukas.moe/x/karen/src/logger"
    "code.lukas.moe/x/karen/src/metrics"
    "code.lukas.moe/x/karen/src/migrations"
    "code.lukas.moe/x/karen/src/version"
    "github.com/bwmarrin/discordgo"
    "github.com/getsentry/raven-go"
    "math/rand"
    "os"
    "os/signal"
    "time"
    "code.lukas.moe/x/karen/src/config"
)

// Entrypoint
func main() {
    Logger.BOOT.L("launcher", "Booting Karen...")

    // Read i18n
    helpers.LoadTranslations()

    // Show version
    version.DumpInfo()

    // Start metric server
    metrics.Init()

    // Make the randomness more random
    rand.Seed(time.Now().UTC().UnixNano())

    // Check if the bot is being debugged
    if config.Get("debug").(bool) {
        helpers.DEBUG_MODE = true
        Logger.DEBUG_MODE = true
    }

    // Print UA
    Logger.BOOT.L("launcher", "USERAGENT: '"+helpers.DEFAULT_UA+"'")

    // Call home
    Logger.BOOT.L("launcher", "[SENTRY] Calling home...")
    err := raven.SetDSN(config.Get("sentry").(string))
    if err != nil {
        panic(err)
    }
    Logger.BOOT.L("launcher", "[SENTRY] Someone picked up the phone \\^-^/")

    // Connect to DB
    Logger.BOOT.L("launcher", "Opening database connection...")
    helpers.ConnectDB(
        config.Get("rethink.url").(string),
        config.Get("rethink.db").(string),
    )

    // Close DB when main dies
    defer helpers.GetDB().Close()

    // Run migrations
    migrations.Run()

    // Connect and add event handlers
    Logger.BOOT.L("launcher", "Connecting to discord...")
    discord, err := discordgo.New("Bot " + config.Get("discord.token").(string))
    if err != nil {
        panic(err)
    }

    discord.Lock()
    discord.LogLevel = discordgo.LogError
    discord.Unlock()

    discord.AddHandler(BotOnReady)
    discord.AddHandler(BotOnMessageCreate)
    discord.AddHandler(BotOnReactionAdd)
    discord.AddHandler(BotOnGuildMemberJoin)
    discord.AddHandler(BotOnGuildMemberRemove)
    discord.AddHandler(metrics.OnReady)
    discord.AddHandler(metrics.OnMessageCreate)

    // Connect to discord
    err = discord.Open()
    if err != nil {
        raven.CaptureErrorAndWait(err, nil)
        panic(err)
    }

    // Make a channel that waits for a os signal
    channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt, os.Kill)

    // Wait until the os wants us to shutdown
    <-channel

    Logger.ERROR.L("launcher", "The OS is killing me :c")
    Logger.ERROR.L("launcher", "Disconnecting...")
    discord.Close()
}
