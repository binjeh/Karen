package main

import (
    "code.lukas.moe/x/karen/src/helpers"
    Logger "code.lukas.moe/x/karen/src/logger"
    "code.lukas.moe/x/karen/src/metrics"
    "code.lukas.moe/x/karen/src/migrations"
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
    Logger.PrintLogo()

    Logger.BOOT.L("Booting Karen...")

    // Read i18n
    helpers.LoadTranslations()

    // Start metric server
    metrics.Init()

    // Make the randomness more random
    rand.Seed(time.Now().UTC().UnixNano())

    // Check if the bot is being debugged
    if config.GetDefault("core.debugMode", false).(bool) {
        helpers.DEBUG_MODE = true
        Logger.DEBUG_MODE = true
    }

    // Print UA
    Logger.BOOT.L("USERAGENT: '" + helpers.DEFAULT_UA + "'")

    // Call home
    Logger.BOOT.L("[SENTRY] Calling home...")
    err := raven.SetDSN(config.Get("core.sentry.dsn").(string))
    if err != nil {
        panic(err)
    }
    Logger.BOOT.L("[SENTRY] Someone picked up the phone \\^-^/")

    // Connect to DB
    Logger.BOOT.L("Opening database connection...")
    helpers.ConnectDB(
        config.Get("core.db.ip").(string)+":"+config.Get("core.db.port").(string),
        config.Get("core.db.name").(string),
    )

    // Close DB when main dies
    defer helpers.GetDB().Close()

    // Run migrations
    migrations.Run()

    // Connect and add event handlers
    Logger.BOOT.L("Connecting to discord...")
    discord, err := discordgo.New("Bot " + config.Get("core.discord.token").(string))
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

    Logger.ERROR.L("The OS is killing me :c")
    Logger.ERROR.L("Disconnecting...")
    discord.Close()
}
