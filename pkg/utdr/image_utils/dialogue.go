package imageutils

import (
	"errors"
	"strings"

	"golang.org/x/image/math/fixed"
)

type DialogueDrawOptions struct {
	IsSentence bool
}

type DialogueNewLineOptions struct {
	NextLineIsSentence bool
}

var NewSentenceOpts = DialogueDrawOptions{
	IsSentence: true,
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
		ctx.NewLine(&DialogueNewLineOptions{
			NextLineIsSentence: true,
		})
	}
	return nil
}

func (ctx *Context) DrawText(text string, opts *DialogueDrawOptions) error {
	err := ctx.checkDialogueFlags()
	if err != nil {
		return err
	}

	textWidth := ctx.FontDrawer.MeasureString(text)

	if ctx.FontDrawer.Dot.X+textWidth+verticalPadding > fixed.I(textboxWidth) {
		return DrawOverflowError
	}

	if opts != nil && opts.IsSentence {
		ctx.FontDrawer.DrawString("* ")
	}

	if !ctx.Flags.HasDialogue {
		ctx.Flags.HasDialogue = true
	}

	ctx.FontDrawer.DrawString(text)
	return nil
}

func (ctx *Context) drawFragment(fragment string) bool {
	textWidth := ctx.FontDrawer.MeasureString(fragment)

	if ctx.FontDrawer.Dot.X+textWidth+verticalPadding > fixed.I(textboxWidth) {
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
		ok := ctx.drawFragment(word)
		if !ok {
			ctx.NewLine(nil)
			ctx.drawFragment(word)
		}
		ctx.drawFragment(" ")
	}
	return nil
}

func (ctx *Context) NewLine(opts *DialogueNewLineOptions) error {
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
