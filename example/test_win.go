//go:build windows
// +build windows

package main

import (
	"fmt"
	escape "github.com/stolen-programmer/ansi-escape"
)

func main() {
	escape.Wrap(func() {
		fmt.Print(escape.AlternateScreenBuffer)
		fmt.Print(escape.Attr("Hello", escape.Underline, escape.ForegroundExt,
			"173", "255", "47"))
		fmt.Print(escape.Attr(" world", escape.FRed, escape.BBlue))
		fmt.Print(escape.Normal)
		fmt.Scanln()
		fmt.Print(escape.MainScreenBuffer)
	})
}
