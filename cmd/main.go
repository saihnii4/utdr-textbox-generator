package main

import (
	"image"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/saihnii4/utdr-video-creator/v2/cmd/ctx"
	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils/portraits/asriel"
	"golang.org/x/image/font"
)

var ctx *imageutils.Context

func init() {
	var err error
	ctx, err := ctx.Init()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func main() {
	var err error

	err = asriel.Init()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ctx.DrawTextbox()
	ctx.DrawPortrait(asriel.Portrait)
	ctx.DrawDialogue("howdy", &imageutils.NewSentenceOpts)
	ctx.NewLine(&imageutils.DialogueNewLineOptions{
		NextLineIsSentence: true,
	})
	ctx.DrawDialogue("howdy", &imageutils.NewSentenceOpts)
	ctx.Finalize()
	ctx.ExportAsFile("out.png")

	log.Println("success")
}
