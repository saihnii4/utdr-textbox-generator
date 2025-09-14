package imageutils

import (
	"bufio"
	"errors"
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Flags struct {
	HasPortrait          bool
	HasDialogue          bool
	HasBorder            bool
	HasTextBox           bool
	DialogueDrawComplete bool
}

type Context struct {
	Background  draw.Image
	FontSource  image.Image
	Flags       *Flags
	Settings    *ContextBuilder
	FontFace    font.Face
	FontDrawer  font.Drawer
	originPoint fixed.Point26_6
}

// returns a rectangle and a point where the rectangle should be drawn from
func borderRect(borderWidth int) image.Rectangle {
	return image.Rect(borderWidth, borderWidth, textboxWidth-borderWidth, textboxHeight-borderWidth)
}

func (ctx *Context) Fill(src image.Image) {
	draw.Draw(ctx.Background, ctx.Background.Bounds(), src, image.Pt(0, 0), draw.Src)
}

func (ctx *Context) DrawTextbox() error {
	if ctx.Flags.HasTextBox {
		return errors.New("cannot redraw textbox")
	}
	ctx.Fill(image.White)
	box := borderRect(textboxBorderWidth)
	draw.Draw(ctx.Background, box, image.Black, box.Min, draw.Src)
	ctx.FontDrawer.Dot = ctx.FontDrawer.Dot.Add(fixed.Point26_6{
		X: horizontalPadding + fixed.I(textboxBorderWidth),
		Y: verticalOffset + verticalPadding + fixed.I(textboxBorderWidth),
	})
	ctx.originPoint = ctx.FontDrawer.Dot
	ctx.Flags.HasTextBox = true
	return nil
}

func (ctx *Context) ExportAsFile(filepath string) error {
	var err error

	outRef, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer outRef.Close()

	buf := bufio.NewWriter(outRef)
	err = png.Encode(buf, ctx.Background)
	if err != nil {
		return err
	}

	err = buf.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (ctx *Context) WriteRequest(w http.ResponseWriter) error {
	err := png.Encode(w, ctx.Background)
	if err != nil {
		return err
	}
	return nil
}
