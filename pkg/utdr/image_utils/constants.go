package imageutils

import (
	"errors"
)

var DrawOverflowError = errors.New("overflow error while drawing")

var (
	textboxHeight         = 152 * 2
	textboxWidth          = 578 * 2
	textboxBorderWidth    = 10
	verticalOffset        = 36 // minimum spacing required to get text on screen
	verticalPadding       = 48
	horizontalPadding     = 48
	portraitHorizontalGap = 48
)
