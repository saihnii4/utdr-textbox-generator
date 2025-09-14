package portraits

import (
	"encoding/json"
	"os"

	imageutils "github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/image_utils"
)

func ReadPortraitConfig(filepath string) (*imageutils.Portrait, error) {
	var portrait *imageutils.Portrait

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, portrait)
	if err != nil {
		return nil, err
	}

	return portrait, nil
}
