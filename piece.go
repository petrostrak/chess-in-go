package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/notnil/chess"
)

var (
	moveStart    chess.Square = chess.NoSquare
	okColor                   = color.NRGBA{0, 0xff, 0, 0xff}
	okBGColor                 = color.NRGBA{0, 0xff, 0, 0x28}
	notOkColor                = color.NRGBA{0xff, 0, 0, 0xff}
	notOkBGColor              = color.NRGBA{0xff, 0, 0, 0x28}
)

type piece struct {
	widget.Icon
	game   *chess.Game
	square chess.Square
}

func newPiece(g *chess.Game, sq chess.Square) *piece {
	p := g.Position().Board().Piece(sq)

	ret := &piece{game: g, square: sq}
	ret.ExtendBaseWidget(ret)

	ret.Resource = resourceForPiece(p)
	return ret
}

func (p *piece) Tapped(ev *fyne.PointEvent) {
	if moveStart == p.square {
		moveStart = chess.NoSquare
		start.Hide()
		start.Refresh()
		return
	}

	if moveStart == chess.NoSquare {
		if m := isValidMove(p.square, chess.NoSquare, p.game); m != nil {
			moveStart = p.square
			start.FillColor = okBGColor
			start.StrokeColor = okColor
		} else {
			start.FillColor = notOkBGColor
			start.StrokeColor = notOkColor
		}

		off := squareToOffset(p.square)
		cell := grid.Objects[off].(*fyne.Container)

		start.Move(cell.Position())
		start.Resize(cell.Size())
		start.Refresh()
		start.Show()

		return
	}

	start.Hide()
	start.Refresh()

	if m := isValidMove(moveStart, p.square, p.game); m != nil {
		moveStart = chess.NoSquare
		move(m, p.game, grid, over)

		go func() {
			time.Sleep(time.Second)
			randomResponse(p.game)
		}()
		return
	}

	msg := fmt.Sprintf("Cannot move piece %d to square %v", p.game.Position().Board().Piece(moveStart), p.square)

	moveStart = chess.NoSquare
	dialog.ShowInformation("Invalid move", msg, w)
}

func randomResponse(game *chess.Game) {
	if game.Outcome() != chess.NoOutcome {
		dialog.ShowInformation("Game ended",
			"Game outcome "+game.Outcome().String()+" because "+game.Method().String(), w)
		return
	}
	rand.Seed(time.Now().Unix())
	valid := game.ValidMoves()
	m := valid[rand.Intn(len(valid))]
	move(m, game, grid, over)
}

func isValidMove(s1, s2 chess.Square, g *chess.Game) *chess.Move {
	valid := g.ValidMoves()
	for _, m := range valid {
		if m.S1() == s1 && (s2 == chess.NoSquare || m.S2() == s2) {
			return m
		}
	}

	return nil
}
