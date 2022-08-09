package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"
)

func createGrid(b *chess.Board) *fyne.Container {
	var cells []fyne.CanvasObject

	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.Gray{0xD0})

			// alternate color to achieve black n white chess pattern
			if x%2 == y%2 {
				bg.FillColor = color.Gray{0x40}
			}

			// revert fyne x and y axis to match chess's library
			piece := b.Piece(chess.Square(x + y*8))

			img := canvas.NewImageFromResource(resourceForPiece(piece))
			// image maintains its aspect ratio while in the canvas
			img.FillMode = canvas.ImageFillContain

			// add all 64 rectangles to the grid
			cells = append(cells, container.NewMax(bg, img))
		}
	}

	return container.New(&boardLayout{}, cells...)
}

func squareToOffset(sq chess.Square) int {
	x := sq % 8
	y := 7 - ((sq - x) / 8)

	return int(x + y*8)
}

func move(m *chess.Move, game *chess.Game, grid *fyne.Container, over *canvas.Image) {
	off1 := squareToOffset(m.S1())
	cell := grid.Objects[off1].(*fyne.Container)
	img1 := cell.Objects[1].(*canvas.Image)
	pos1 := cell.Position()

	over.Resource = img1.Resource
	over.Move(pos1)
	over.Resize(img1.Size())

	img1.Resource = nil
	img1.Refresh()
	over.Show()

	off2 := squareToOffset(m.S2())
	cell = grid.Objects[off2].(*fyne.Container)
	pos2 := cell.Position()

	a := canvas.NewPositionAnimation(pos1, pos2, time.Millisecond*500, func(p fyne.Position) {
		over.Move(p)
		over.Refresh()
	})
	a.Start()
	time.Sleep(time.Millisecond * 500)

	game.Move(m)
	over.Hide()
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
