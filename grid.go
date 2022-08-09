package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func createGrid() *fyne.Container {
	grid := container.NewGridWithColumns(8)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})

			// alternate color to achieve black n white chess pattern
			if x%2 == y%2 {
				bg.FillColor = color.Gray{0xE0}
			}

			img := canvas.NewImageFromResource(resourceForPiece())

			// image maintains its aspect ratio while in the canvas
			img.FillMode = canvas.ImageFillContain

			// add all 64 rectangles to the grid
			grid.Add(container.NewMax(bg, img))
		}
	}

	return grid
}
