package plugins

import (
    "github.com/bwmarrin/discordgo"
    "regexp"
    "image"
    "image/color/palette"
    "image/gif"
    "os"
    "time"
    "code.lukas.moe/x/karen/src/helpers"
    "strconv"
    "image/color"
    "image/draw"
    "github.com/golang/freetype"
    "code.lukas.moe/x/karen/src/assets"
    "golang.org/x/image/font"
    "strings"
)

const (
    GIF_WIDTH        = 640
    GIF_HEIGHT       = 480
    GIF_DPI          = 72
    GIF_FONT         = "assets/Helvetica.ttf"
    GIF_FONT_SIZE    = 42
    GIF_LINE_SPACING = 1.5
    GIF_FG           = 0xffff
    GIF_BG           = 0x0
)

type Spoiler struct{}

func (s *Spoiler) Commands() []string {
    return []string{
        // User-facing spoiler commands
        "s",
        "spoil",
        "spoiler",

        // Internal command for on-the-fly spoilers
        "stopic",

        // Admin command to mark spoilers
        "spoils",
    }
}

func (s *Spoiler) Init(session *discordgo.Session) {
    session.AddHandler(s.MessageInspector)
}

func (s *Spoiler) MessageInspector(session *discordgo.Session, e *discordgo.MessageCreate) {
    if regexp.MustCompile("^.*?:(s|spoil|spoiler):.*$").MatchString(e.Content) {
        s.Action("stopic", e.Content, e.Message, session)
    }
}

func (s *Spoiler) Action(command, content string, msg *discordgo.Message, session *discordgo.Session) {
    switch command {
    case "sflag":
        args := strings.Fields(content)
        s.MarkAndHide(
            msg.ChannelID,
            args[0],
            helpers.GetTextF("plugins.spoiler.flagged", msg.Author.Username),
            session,
        )
        break

    case "stopic":
        topic := strings.Split(content, ":s:")
        s.MarkAndHide(
            msg.ChannelID,
            msg.ID,
            helpers.GetTextF("plugins.spoiler.topic", msg.Author.Username, topic),
            session,
        )
        break

    default:
        s.MarkAndHide(
            msg.ChannelID,
            msg.ID,
            helpers.GetTextF("plugins.spoiler.topicless", msg.Author.Username),
            session,
        )
        break
    }
}

func (s *Spoiler) MarkAndHide(channelId string, messageId string, attachmentText string, session *discordgo.Session) {
    var e error

    // Store copy of image
    msg, e := session.ChannelMessage(channelId, messageId)
    helpers.Relax(e)

    // Create a new gif
    frames := []*image.Paletted{
        drawImage([]string{"Hover to reveal spoiler"}),
        drawImage([]string{msg.Content}),
    }
    delays := []int{0, 60000}

    filename := helpers.BtoA(strconv.Itoa(int(time.Now().Unix()))) + ".gif"

    fw, e := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
    helpers.Relax(e)

    e = gif.EncodeAll(fw, &gif.GIF{
        Image: frames,
        Delay: delays,
    })
    helpers.Relax(e)

    e = fw.Close()
    helpers.Relax(e)

    fr, e := os.OpenFile(filename, os.O_RDONLY, 0644)
    helpers.Relax(e)

    _, e = session.ChannelFileSendWithMessage(channelId, attachmentText, filename, fr)
    helpers.Relax(e)

    e = fr.Close()
    helpers.Relax(e)

    e = os.Remove(filename)
    helpers.Relax(e)

    // Delete the original message
    e = session.ChannelMessageDelete(channelId, messageId)
    helpers.Relax(e)
}

//noinspection GoStructInitializationWithoutFieldNames
func drawImage(text []string) (*image.Paletted) {
    ruler := color.RGBA{0x22, 0x22, 0x22, 0xff}
    fg := image.NewUniform(color.Gray16{GIF_FG})
    bg := image.NewUniform(color.Gray16{GIF_BG})

    fontBytes, err := assets.Asset(GIF_FONT)
    helpers.Relax(err)

    ttf, err := freetype.ParseFont(fontBytes)
    helpers.Relax(err)

    img := image.NewPaletted(image.Rect(0, 0, GIF_WIDTH, GIF_HEIGHT), palette.Plan9)
    draw.Draw(img, img.Bounds(), bg, image.ZP, draw.Src)

    c := freetype.NewContext()
    c.SetDPI(GIF_DPI)
    c.SetFont(ttf)
    c.SetFontSize(GIF_FONT_SIZE)
    c.SetClip(img.Bounds())
    c.SetDst(img)
    c.SetSrc(fg)
    c.SetHinting(font.HintingNone)

    for i := 0; i < 200; i++ {
        img.Set(10, 10+i, ruler)
        img.Set(10+i, 10, ruler)
    }

    pt := freetype.Pt(10, 10+int(c.PointToFixed(GIF_FONT_SIZE)>>6))
    for _, s := range text {
        _, err = c.DrawString(s, pt)
        helpers.Relax(err)
        pt.Y += c.PointToFixed(GIF_FONT_SIZE * GIF_LINE_SPACING)
    }

    return img
}
