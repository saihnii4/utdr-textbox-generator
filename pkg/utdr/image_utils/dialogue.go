package imageutils

import (
	"errors"
	"strings"

	"golang.org/x/image/math/fixed"
)

type DialogueDrawOptions struct {
	// unimplemented
}

func (ctx *Context) checkDialogueFlags() error {
	if ctx.Flags.DialogueDrawComplete {
		return errors.New("cannot write to textbox when it's finalized")
	}

	if !ctx.Flags.HasTextBox {
		return errors.New("textbox must be drawn first")
	}

	return nil
}

func (ctx *Context) DrawDialogue(text string, opts *DialogueDrawOptions) error {
	lines := strings.SplitSeq(text, "\n")
	for line := range lines {
		ctx.drawLine(line)
		ctx.NewLine()
	}
	return nil
}

func (ctx *Context) drawFragment(fragment string, padLeft bool) bool {
	if padLeft {
		fragment = "  " + fragment
	}

	textWidth := ctx.FontDrawer.MeasureString(fragment)

	if ctx.FontDrawer.Dot.X+textWidth+fixed.I(verticalPadding) > fixed.I(textboxWidth) {
		return false
	}

	ctx.FontDrawer.DrawString(fragment)
	return true
}

func (ctx *Context) drawLine(line string) error {
	line = "* " + line
	err := ctx.checkDialogueFlags()
	if err != nil {
		return err
	}

	words := strings.SplitSeq(line, " ")
	for word := range words {
		ok := ctx.drawFragment(word, false)
		if !ok {
			ctx.NewLine()
			ctx.drawFragment(word, true)
		}
		ctx.drawFragment(" ", false)
	}
	return nil
}

func (ctx *Context) NewLine() error {
	err := ctx.checkDialogueFlags()
	if err != nil {
		return err
	}

	if !ctx.Flags.HasDialogue {
		ctx.Flags.HasDialogue = true
	}

	yOffset := fixed.Int26_6(ctx.Settings.FontSize * ctx.Settings.LineSpacing * 64)
	ctx.FontDrawer.Dot.X = ctx.originPoint.X
	ctx.FontDrawer.Dot.Y += ctx.FontFace.Metrics().Height + yOffset
	return nil
}

func (ctx *Context) Finalize() error {
	if ctx.Flags.DialogueDrawComplete {
		return errors.New("already finalized")
	}
	ctx.Flags.DialogueDrawComplete = true
	return nil
}
