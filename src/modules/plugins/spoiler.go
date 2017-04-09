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
    GIF_WIDTH          = 640
    GIF_HEIGHT         = 480
    GIF_DPI            = 72
    GIF_FONT           = "assets/Helvetica.ttf"
    GIF_FONT_SIZE      = 24
    GIF_LINE_SPACING   = 1.5
    GIF_MAX_LINE_CHARS = 55
)

type Spoiler struct{}

func (s *Spoiler) Commands() []string {
    return []string{
        // User-facing spoiler commands
        "s",
        "spoil",
        "spoiler",

        // Admin command to mark spoilers
        "sflag",
    }
}

func (s *Spoiler) Init(session *discordgo.Session) {
    session.AddHandler(s.MessageInspector)
}

func (s *Spoiler) MessageInspector(session *discordgo.Session, e *discordgo.MessageCreate) {
    msg := strings.Replace(e.Content, "\n", "{{{NEWLINE}}}", -1)
    regex := regexp.MustCompile("(?i)^(.*?)(:s:|:spoil:|:spoiler:)(.*)$")

    if regex.MatchString(msg) {
        matches := regex.FindStringSubmatch(msg)

        s.MarkAndHide(
            e.Message.ChannelID,
            e.Message.ID,
            strings.Replace(matches[3], "{{{NEWLINE}}}", "\n", -1),
            helpers.GetTextF("plugins.spoiler.topic", e.Author.Username, matches[1]),
            session,
        )
    }
}

func (s *Spoiler) Action(command, content string, msg *discordgo.Message, session *discordgo.Session) {
    switch command {
    case "sflag":
        args := strings.Fields(content)
        flagged, e := session.ChannelMessage(msg.ChannelID, args[0])
        helpers.Relax(e)

        s.MarkAndHide(
            msg.ChannelID,
            flagged.ID,
            flagged.Content,
            helpers.GetTextF("plugins.spoiler.flagged", flagged.Author.Username, msg.Author.Username),
            session,
        )
        break

    default:
        s.MarkAndHide(
            msg.ChannelID,
            msg.ID,
            content,
            helpers.GetTextF("plugins.spoiler.topicless", msg.Author.Username),
            session,
        )
        break
    }
}

func (s *Spoiler) MarkAndHide(channelId string, messageId string, spoilerText string, attachmentText string, session *discordgo.Session) {
    var e error

    // Create a new gif
    frames := []*image.Paletted{
        drawImage([]string{"Hover to reveal spoiler"}),
        drawImage(strings.Split(wordWrap(spoilerText, GIF_MAX_LINE_CHARS), "\n")),
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
    fg := image.NewUniform(color.RGBA{255, 255, 255, 255})
    bg := image.NewUniform(color.RGBA{60, 63, 68, 255})

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

    pt := freetype.Pt(10, 10+int(c.PointToFixed(GIF_FONT_SIZE)>>6))
    for _, s := range text {
        _, err = c.DrawString(s, pt)
        helpers.Relax(err)
        pt.Y += c.PointToFixed(GIF_FONT_SIZE * GIF_LINE_SPACING)
    }

    return img
}

func wordWrap(text string, lineWidth int) string {
    words := strings.Fields(strings.TrimSpace(text))

    if len(words) == 0 {
        return text
    }

    wrapped := words[0]
    spaceLeft := lineWidth - len(wrapped)

    for _, word := range words[1:] {
        if len(word)+1 > spaceLeft {
            wrapped += "\n" + word
            spaceLeft = lineWidth - len(word)
        } else {
            wrapped += " " + word
            spaceLeft -= 1 + len(word)
        }
    }

    return wrapped
}
