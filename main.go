package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/notnil/chess"
)

func main() {
	a := app.New()
	w := a.NewWindow("Chess")

	game := chess.NewGame()
	grid := createGrid(game.Position().Board())
	w.SetContent(grid)
	w.Resize(fyne.NewSize(480, 480))

	go func() {
		rand.Seed(time.Now().Unix())

		for game.Outcome() == chess.NoOutcome {
			time.Sleep(time.Millisecond * 500)
			valid := game.ValidMoves()
			m := valid[rand.Intn(len(valid))]

			move(m, game, grid)
		}

	}()

	w.ShowAndRun()
}
