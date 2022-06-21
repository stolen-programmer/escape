//go:build windows
// +build windows

package escape

import (
	escapes "github.com/snugfox/ansi-escapes"
	"os"
)

func Wrap(main func()) {
	escapes.EnableVirtualTerminal(os.Stdout.Fd())
	defer escapes.DisableVirtualTerminal(os.Stdout.Fd())
	main()
}
