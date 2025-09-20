package main

import (
	"log"

	"github.com/saihnii4/utdr-video-creator/v2/cmd/context"
	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils/portraits/asriel"
)

var ctx *imageutils.Context

func init() {
	var err error
	ctx, err = context.Init()
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
	err = ctx.DrawDialogue("howdy!\nhowdy i'm flowey! flowey the flower!", nil)
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
