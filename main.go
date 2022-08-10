package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"
)

var (
	grid *fyne.Container
	over *canvas.Image
	w    fyne.Window
)

func main() {
	a := app.New()
	w = a.NewWindow("Chess")

	game := chess.NewGame()
	grid = createGrid(game)

	over = canvas.NewImageFromResource(nil)
	over.Hide()
	w.SetContent(container.NewMax(grid, container.NewWithoutLayout(over)))
	w.Resize(fyne.NewSize(480, 480))

	w.ShowAndRun()
}
