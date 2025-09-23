package main

import (
	"image"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	imageutils "github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils/portraits/asriel"
	"golang.org/x/image/font"
)

var ctx *imageutils.Context

var (
	utFont       *truetype.Font
	fontSize     = 9.0
	lineSpacing  = 1.25
	fontFilePath = "dtm-mono.ttf"
	dpi          = 512.0
	ctxBuilder   *imageutils.ContextBuilder
)

func init() {
	var err error

	utFont, err = imageutils.ReadFontFile(fontFilePath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	color := color.RGBA{0, 255, 0, 255}

	ctxBuilder = &imageutils.ContextBuilder{
		Hinting:     font.HintingFull,
		Font:        utFont,
		FontSize:    fontSize,
		LineSpacing: lineSpacing,
		FontSource:  image.NewUniform(color),
		DPI:         dpi,
	}
}

func main() {
	var err error
	ctx = ctxBuilder.Init()

	err = asriel.Init()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ctx.DrawTextbox()
	ctx.DrawPortrait(asriel.Portrait)
	err = ctx.DrawDialogue("honk shoo mi mi gng", nil)
	if err != nil {
		panic(err)
	}
	ctx.Finalize()
	ctx.ExportAsFile("out.png")

	// ctx.DrawTextbox()
	// ctx.DrawPortrait(asriel.Portrait)
	// ctx.DrawText("howdy", &imageutils.NewSentenceOpts)
	// ctx.NewLine(&imageutils.DialogueNewLineOptions{
	// 	NextLineIsSentence: true,
	// })
	// ctx.DrawText("howdy", &imageutils.NewSentenceOpts)
	// ctx.Finalize()
	// ctx.ExportAsFile("out.png")

	log.Println("success")
}
