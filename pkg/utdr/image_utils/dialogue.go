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
		words := strings.Split(line, " ")
		for i, word := range words {
			var err error
			if i == 0 {
				err = ctx.DrawText(word, &NewSentenceOpts)
			} else {
				err = ctx.DrawText(word, nil)
			}
			ctx.DrawText(" ", nil)

			if err != nil {
				if !errors.Is(err, DrawOverflowError) {
					return err
				}

				ctx.NewLine(&DialogueNewLineOptions{
					NextLineIsSentence: false,
				})
				ctx.DrawText(word, nil)
				ctx.DrawText(" ", nil)
			}
		}
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
