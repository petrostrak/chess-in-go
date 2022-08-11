package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/notnil/chess"
)

func createGrid(g *chess.Game) *fyne.Container {
	var cells []fyne.CanvasObject

	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.NRGBA{0xF4, 0xE2, 0xB6, 0xFF})
			effect := canvas.NewImageFromResource(resourceOverlay1Png)

			// alternate color to achieve black n white chess pattern
			if x%2 == y%2 {
				bg.FillColor = color.NRGBA{0x73, 0x50, 0x32, 0xFF}
				effect.Resource = resourceOverlay2Png
			}

			p := newPiece(g, chess.Square(x+y*8))

			// add all 64 rectangles to the grid
			cells = append(cells, container.NewMax(bg, effect, p))
		}
	}

	return container.New(&boardLayout{}, cells...)
}

func squareToOffset(sq chess.Square) int {
	x := sq % 8
	y := 7 - ((sq - x) / 8)

	return int(x + y*8)
}

func positionToSquare(pos fyne.Position) chess.Square {
	var offX, offY = -1, -1
	for x := float32(0); x <= pos.X; x += grid.Size().Width / 8 {
		offX++
	}
	for y := float32(0); y <= pos.Y; y += grid.Size().Height / 8 {
		offY++
	}

	return chess.Square((7-offY)*8 + offX)
}

func move(m *chess.Move, game *chess.Game, grid *fyne.Container, over *canvas.Image) {
	off := squareToOffset(m.S1())
	cell := grid.Objects[off].(*fyne.Container)
	img := cell.Objects[2].(*piece)

	over.Resource = resourceForPiece(game.Position().Board().Piece(m.S1()))
	over.Resize(img.Size())
	over.Refresh()

	img.Resource = nil
	img.Refresh()

	off = squareToOffset(m.S2())
	cell = grid.Objects[off].(*fyne.Container)
	pos2 := cell.Position()

	a := canvas.NewPositionAnimation(over.Position(), pos2, time.Millisecond*500, func(p fyne.Position) {
		over.Move(p)
		over.Refresh()
	})
	a.Start()
	time.Sleep(time.Millisecond * 500)

	game.Move(m)
	refreshGrid(grid, game.Position().Board())
	over.Hide()

	if game.Outcome() != chess.NoOutcome {
		result := "draw"
		switch game.Outcome().String() {
		case "1-0":
			result = "won"
		case "0-1":
			result = "lost"
		}
		dialog.ShowInformation("Game ended", "Game "+result+" because "+game.Method().String(), w)
	}
}

func refreshGrid(grid *fyne.Container, b *chess.Board) {
	y, x := 7, 0
	for _, cell := range grid.Objects {
		p := b.Piece(chess.Square(x + y*8))

		img := cell.(*fyne.Container).Objects[2].(*piece)
		img.Resource = resourceForPiece(p)
		img.Refresh()

		x++
		if x == 8 {
			x = 0
			y--
		}
	}
}
