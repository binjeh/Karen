package plugins

import (
    "cloud.google.com/go/translate"
    "context"
    "code.lukas.moe/x/karen/src/helpers"
    "github.com/bwmarrin/discordgo"
    "golang.org/x/text/language"
    "google.golang.org/api/option"
    "strings"
    "code.lukas.moe/x/karen/src/config"
)

type Translator struct {
    ctx    context.Context
    client *translate.Client
}

func (t *Translator) Commands() []string {
    return []string{
        "translator",
        "translate",
        "t",
    }
}

func (t *Translator) Init(session *discordgo.Session) {
    t.ctx = context.Background()

    client, err := translate.NewClient(
        t.ctx,
        option.WithAPIKey(config.Get("modules.google.translator.key").(string)),
    )
    helpers.Relax(err)
    t.client = client
}

func (t *Translator) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
    // Assumed format: <lang_in> <lang_out> <text>
    parts := strings.Split(content, " ")

    if len(parts) < 3 {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.translator.check_format"))
        return
    }

    source, err := language.Parse(parts[0])
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.translator.unknown_lang", parts[0]))
        return
    }

    target, err := language.Parse(parts[1])
    if err != nil {
        session.ChannelMessageSend(msg.ChannelID, helpers.GetTextF("plugins.translator.unknown_lang", parts[1]))
        return
    }

    translations, err := t.client.Translate(
        t.ctx,
        parts[2:],
        target,
        &translate.Options{
            Format: translate.Text,
            Source: source,
        },
    )

    if err != nil {
        //session.ChannelMessageSend(msg.ChannelID, helpers.GetText("plugins.translator.error"))
        helpers.SendError(msg, err)
        return
    }

    m := ""
    for _, piece := range translations {
        m += piece.Text + " "
    }

    session.ChannelMessageSend(
        msg.ChannelID,
        ":earth_africa: "+strings.ToUpper(source.String())+" => "+strings.ToUpper(target.String())+"\n```\n"+m+"\n```",
    )
}
