package api

import (
	"log"
	"net/http"

	"github.com/saihnii4/utdr-video-creator/v2/cmd/ctx"
	imageutils "github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils/portraits/asriel"
)

func GenerateTextBox(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	content := q.Get("content")
	if content == "" {
		http.Error(w, "textbox content missing", http.StatusBadRequest)
		return
	}

	ctx, err := ctx.Init()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "couldn't init context", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")

	err = asriel.Init()
	ctx.DrawTextbox()
	ctx.DrawPortrait(asriel.Portrait)
	ctx.DrawDialogue(content, &imageutils.NewSentenceOpts)
	ctx.Finalize()
	ctx.WriteRequest(w)
}

func GenerateAnimatedTextBox(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://google.com", http.StatusPermanentRedirect)
	// TODO:
}
