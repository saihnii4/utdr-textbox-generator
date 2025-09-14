package imageutils

import (
	"image"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type ContextBuilder struct {
	Hinting     font.Hinting
	Font        *truetype.Font
	DPI         float64
	FontSize    float64
	LineSpacing float64
	FontSource  image.Image
}

func (c *ContextBuilder) Init() *Context {
	bg := image.NewRGBA(image.Rect(0, 0, textboxWidth, textboxHeight))

	fontFace := truetype.NewFace(c.Font, &truetype.Options{
		DPI:     c.DPI,
		Hinting: c.Hinting,
		Size:    c.FontSize,
	})

	drawer := font.Drawer{
		Face: fontFace,
		Dst:  bg,
		Src:  c.FontSource,
	}

	return &Context{
		Background:  bg,
		FontSource:  c.FontSource,
		Settings:    c,
		Flags:       &Flags{},
		FontFace:    fontFace,
		FontDrawer:  drawer,
		originPoint: fixed.P(0, 0),
	}
}
