package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"
)

func createGrid(b *chess.Board) *fyne.Container {
	grid := container.NewGridWithColumns(8)

	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.Gray{0xE0})

			// alternate color to achieve black n white chess pattern
			if x%2 == y%2 {
				bg.FillColor = color.Gray{0x30}
			}

			// revert fyne x and y axis to match chess's library
			piece := b.Piece(chess.Square(x + y*8))

			img := canvas.NewImageFromResource(resourceForPiece(piece))
			// image maintains its aspect ratio while in the canvas
			img.FillMode = canvas.ImageFillContain

			// add all 64 rectangles to the grid
			grid.Add(container.NewMax(bg, img))
		}
	}

	return grid
}

func move(m *chess.Move, game *chess.Game, grid *fyne.Container) {
	game.Move(m)
	refreshGrid(grid, game.Position().Board())
}

func refreshGrid(grid *fyne.Container, b *chess.Board) {
	y, x := 7, 0
	for _, cell := range grid.Objects {
		piece := b.Piece(chess.Square(x + y*8))

		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		img.Resource = resourceForPiece(piece)
		img.Refresh()

		x++
		if x == 8 {
			x = 0
			y--
		}
	}
}
