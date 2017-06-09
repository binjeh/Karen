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

package plugins

import (
    "code.lukas.moe/x/karen/src/assets"
    "code.lukas.moe/x/karen/src/cache"
    "code.lukas.moe/x/karen/src/helpers"
    "github.com/bwmarrin/discordgo"
    "github.com/golang/freetype"
    "golang.org/x/image/font"
    "image"
    "image/color"
    "image/color/palette"
    "image/draw"
    "image/gif"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
    "code.lukas.moe/x/karen/src/except"
    "code.lukas.moe/x/karen/src/i18n"
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
    defer except.RecoverDiscord(e.Message)

    msg := strings.Replace(e.Content, "\n", "{{{NEWLINE}}}", -1)
    regex := regexp.MustCompile("(?i)^(.*?)(:s:|:spoil:|:spoiler:)(.*)$")

    if regex.MatchString(msg) {
        matches := regex.FindStringSubmatch(msg)

        s.MarkAndHide(
            e.Message.ChannelID,
            e.Message.ID,
            strings.Replace(matches[3], "{{{NEWLINE}}}", "\n", -1),
            i18n.GetTextF("plugins.spoiler.topic", e.Author.Username, matches[1]),
        )
    }
}

func (s *Spoiler) Action(command, content string, msg *discordgo.Message, session *discordgo.Session) {
    switch command {
    case "sflag":
        helpers.RequireAdmin(msg, func() {
            args := strings.Fields(content)
            flagged, e := session.ChannelMessage(msg.ChannelID, args[0])
            except.Handle(e)

            s.MarkAndHide(
                msg.ChannelID,
                flagged.ID,
                flagged.Content,
                i18n.GetTextF("plugins.spoiler.flagged", flagged.Author.Username, msg.Author.Username),
            )
        })
        break

    default:
        s.MarkAndHide(
            msg.ChannelID,
            msg.ID,
            content,
            i18n.GetTextF("plugins.spoiler.topicless", msg.Author.Username),
        )
        break
    }
}

func (s *Spoiler) MarkAndHide(channelId string, messageId string, spoilerText string, attachmentText string) {
    var e error

    // Create a new gif
    frames := []*image.Paletted{
        drawImage([]string{"Hover to reveal spoiler"}),
        drawImage(strings.Split(wordWrap(spoilerText, GIF_MAX_LINE_CHARS), "\n")),
    }
    delays := []int{0, 60000}

    filename := helpers.BtoA(strconv.Itoa(int(time.Now().Unix()))) + ".gif"

    fw, e := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
    except.Handle(e)

    e = gif.EncodeAll(fw, &gif.GIF{
        Image: frames,
        Delay: delays,
    })
    except.Handle(e)

    e = fw.Close()
    except.Handle(e)

    fr, e := os.OpenFile(filename, os.O_RDONLY, 0644)
    except.Handle(e)

    // Cleanup and close handles when this method dies
    defer func() {
        e = fr.Close()
        except.Handle(e)

        e = os.Remove(filename)
        except.Handle(e)
    }()

    // Delete the original message
    e = cache.GetSession().ChannelMessageDelete(channelId, messageId)
    if e != nil && strings.Contains(e.Error(), "403") {
        cache.GetSession().ChannelMessageSend(channelId, "I have no permissions to delete the spoiler :frowning:")
        return
    }

    _, e = cache.GetSession().ChannelFileSendWithMessage(channelId, attachmentText, filename, fr)
    except.Handle(e)
}

//noinspection GoStructInitializationWithoutFieldNames
func drawImage(text []string) (*image.Paletted) {
    fg := image.NewUniform(color.RGBA{255, 255, 255, 255})
    bg := image.NewUniform(color.RGBA{60, 63, 68, 255})

    fontBytes, err := assets.Asset(GIF_FONT)
    except.Handle(err)

    ttf, err := freetype.ParseFont(fontBytes)
    except.Handle(err)

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
        except.Handle(err)
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
