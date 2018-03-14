package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Board struct {
	SizeX int
	SizeY int
}

func (b *Board) Draw(imd *imdraw.IMDraw) {
	scale := GetConfig().Scale

	imd.Color = pixel.RGB(1, 1, 1)
	imd.Push(pixel.V(0, 0))
	imd.Push(pixel.V(float64(b.SizeX)*scale, float64(b.SizeY)*scale))

	imd.Rectangle(1 * scale)

	imd.Push(pixel.V(float64(b.SizeX/2)*scale, 0))
	imd.Push(pixel.V(float64(b.SizeX/2)*scale, float64(b.SizeY)*scale))

	imd.Line(1 * scale / 2)
}

func NewBoard() *Board {
	board := new(Board)

	board.SizeX = 150
	board.SizeY = 100

	return board
}
