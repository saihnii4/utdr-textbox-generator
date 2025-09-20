package context

import (
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
	imageutils "github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
	"golang.org/x/image/font"
)

var (
	utFont       *truetype.Font
	fontSize     = 9.0
	lineSpacing  = 1.25
	fontFilePath = "dtm-mono.ttf"
	dpi          = 512.0
)

func Init() (*imageutils.Context, error) {
	var err error

	utFont, err = imageutils.ReadFontFile(fontFilePath)
	if err != nil {
		return nil, err
	}

	color := color.RGBA{0, 255, 0, 255}

	ctxBuilder := imageutils.ContextBuilder{
		Hinting:     font.HintingFull,
		Font:        utFont,
		FontSize:    fontSize,
		LineSpacing: lineSpacing,
		FontSource:  image.NewUniform(color),
		DPI:         dpi,
	}

	return ctxBuilder.Init(), nil
}
