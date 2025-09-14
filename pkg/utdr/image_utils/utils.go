package imageutils

import (
	"image"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func ReadFontFile(filepath string) (*truetype.Font, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	font, err := freetype.ParseFont(bytes)
	if err != nil {
		return nil, err
	}

	return font, nil
}

func LoadImageFile(filePath string) (image.Image, error) {
	fileRef, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(fileRef)
	if err != nil {
		return nil, err
	}

	return img, nil
}
