//go:generate fyne bundle -o bundled.go pieces

package main

import "fyne.io/fyne/v2"

func resourceForPiece() fyne.Resource {
	return resourceWhitePawnSvg
}
