package main

// import (
// 	"image"
// 	"log"
//
// 	"github.com/golang/freetype/truetype"
// 	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
// 	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils/portraits/asriel"
// 	"golang.org/x/image/font"
// )

//
// func init() {
// 	var err error
//
// 	utFont, err = imageutils.ReadFontFile(fontFilePath)
// 	if err != nil {
// 		log.Fatal(err)
// 		panic(err)
// 	}
//
// 	ctxBuilder := imageutils.ContextBuilder{
// 		Hinting:     font.HintingFull,
// 		Font:        utFont,
// 		FontSize:    fontSize,
// 		LineSpacing: lineSpacing,
// 		FontSource:  image.White,
// 		DPI:         dpi,
// 	}
//
// 	ctx = ctxBuilder.Init()
// }

// func main() {
// 	var err error
//
// 	err = asriel.Init()
// 	if err != nil {
// 		log.Fatal(err)
// 		panic(err)
// 	}
//
// 	ctx.DrawTextbox()
// 	ctx.DrawPortrait(asriel.Portrait)
// 	ctx.DrawDialogue("howdy", &imageutils.NewSentenceOpts)
// 	ctx.NewLine(&imageutils.DialogueNewLineOptions{
// 		NextLineIsSentence: true,
// 	})
// 	ctx.DrawDialogue("howdy", &imageutils.NewSentenceOpts)
// 	ctx.Finalize()
// 	ctx.ExportAsFile("out.png")
//
// 	log.Println("success")
// }
