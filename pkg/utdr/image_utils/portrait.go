package imageutils

import (
	"errors"
	"image"
	"image/draw"
	"log/slog"

	"github.com/nfnt/resize"
	"golang.org/x/image/math/fixed"
)

type Portrait struct {
	Image image.Image
}

type Face struct {
	Image *image.Image
	Loc   string
}

func centerRectOn(src, dst image.Rectangle) image.Rectangle {
	dstCenterOfMass := dst.Min.Add(dst.Max).Div(2)
	srcCenterOfMass := src.Min.Add(src.Max).Div(2)
	translationVector := dstCenterOfMass.Sub(srcCenterOfMass)
	return src.Add(translationVector)
}

// resizes img to fit onto dst while preserving aspect ratio
func clampImage(img image.Image, dst image.Rectangle) image.Image {
	// normalize the boundaries so Min is at the zero point
	normalizedBox := image.Rect(0, 0, dst.Dx(), dst.Dy())

	slog.Info("clamping image")
	if img.Bounds().In(normalizedBox) {
		slog.Info("image bounds already confined within dst; returning source")
		return img
	}

	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	slog.Info("image dimensions", "height", imgHeight, "width", imgWidth)
	if imgHeight > imgWidth {
		clamped := resize.Resize(0, uint(dst.Dy()), img, resize.Lanczos3)
		slog.Info(
			"clamped img dimensions",
			"height",
			clamped.Bounds().Dy(),
			"width",
			clamped.Bounds().Dx(),
		)
		return clamped
	} else {
		clamped := resize.Resize(uint(dst.Dx()), 0, img, resize.Lanczos3)
		slog.Info("clamped img dimensions", "height", clamped.Bounds().Dy(), "width", clamped.Bounds().Dx())
		return clamped
	}
}

func (ctx *Context) DrawPortrait(portrait *Portrait) error {
	// portrait will overlap over dialogue
	if ctx.Flags.HasDialogue {
		return errors.New("dialogue already drawn")
	}

	// arbitrary
	startingX := textboxBorderWidth + 48
	startingY := textboxBorderWidth + 36

	// the boundaries of the portrait overlay
	minPoint := image.Point{startingX, startingY}
	maxPoint := image.Point{startingX + 200, textboxHeight - startingY}
	textboxPasteRect := image.Rectangle{minPoint, maxPoint}

	// use the normalized boundaries to resize the image to fit within the box
	clampedImg := clampImage(portrait.Image, textboxPasteRect)

	pasteRect := centerRectOn(
		clampedImg.Bounds(),
		textboxPasteRect,
	) // textboxPasteRect is not normalized

	draw.Draw(ctx.Background, pasteRect.Bounds(), clampedImg, clampedImg.Bounds().Min, draw.Src)

	ctx.FontDrawer.Dot.X += fixed.I(textboxPasteRect.Dx() + portraitHorizontalGap)
	ctx.originPoint.X += fixed.I(textboxPasteRect.Dx() + portraitHorizontalGap)
	ctx.Flags.HasPortrait = true

	return nil
}
