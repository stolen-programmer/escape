package escape

import (
	"fmt"
	"strings"
)

func Attr(v string, attrs ...string) (render string) {
	// ESC[attr;attr SGR v
	// ESC[ext;r;g;b SGR v rgb render

	attr := strings.Join(attrs, ";")

	return CSI + attr + SGR + v
}

func MovePos(y, x int) (render string) {
	// ESC[y;xH

	return fmt.Sprintf("%s%d;%d%s", CSI, y, x, Move)

}
