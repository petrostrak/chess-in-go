package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"
)

func main() {
	a := app.New()
	w := a.NewWindow("Chess")

	game := chess.NewGame()
	grid := createGrid(game.Position().Board())

	over := canvas.NewImageFromResource(nil)
	over.Hide()
	w.SetContent(container.NewMax(grid, container.NewWithoutLayout(over)))
	w.Resize(fyne.NewSize(480, 480))

	go func() {
		rand.Seed(time.Now().Unix())

		for game.Outcome() == chess.NoOutcome {
			time.Sleep(time.Millisecond * 500)
			valid := game.ValidMoves()
			m := valid[rand.Intn(len(valid))]

			move(m, game, grid, over)
		}

	}()

	w.ShowAndRun()
}
