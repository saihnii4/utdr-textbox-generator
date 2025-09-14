package imageutils

import (
	"golang.org/x/image/math/fixed"
)

var (
	textboxHeight         = 152 * 2
	textboxWidth          = 578 * 2
	textboxBorderWidth    = 10
	verticalOffset        = fixed.I(36) // minimum spacing required to get text on screen
	verticalPadding       = fixed.I(48)
	horizontalPadding     = fixed.I(48)
	portraitHorizontalGap = fixed.I(48)
)
