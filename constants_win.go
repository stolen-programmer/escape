// +build windows

package escape

// Cursor
const (
	VAP = "d" // current cursor vertical abstract position
)

// Alternate Screen Buffer ESC [?1049h
// Main Screen Buffer ESC [?1049l
const (
	AlternateScreenBuffer = CSI + "?1049h"
	MainScreenBuffer      = CSI + "?1049l"
)

