package imageutils

import (
	"errors"

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
	err := ctx.checkDialogueFlags()
	if err != nil {
		return err
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
