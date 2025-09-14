package asriel

import (
	"bytes"
	"image/png"
	"os"

	// "github.com/nfnt/resize"
	"github.com/nfnt/resize"
	imageutils "github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
)

var (
	Portrait    *imageutils.Portrait
	portraitLoc = "asriel.png"
	// portraitLoc = "spr_face_asriel0_0.png"
)

func Init() error {
	portraitBytes, err := os.ReadFile(portraitLoc)
	if err != nil {
		return err
	}

	portraitReader := bytes.NewReader(portraitBytes)

	portrait, err := png.Decode(portraitReader)
	if err != nil {
		return err
	}

	img := resize.Resize(125, 125, portrait, resize.Lanczos3)

	Portrait = &imageutils.Portrait{
		Image: img,
	}

	return nil
}
