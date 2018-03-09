package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Ball struct {
	X     int
	Y     int
	Speed int

	VectorX int
	VectorY int

	Board *Board
}

func (b *Ball) Draw(imd *imdraw.IMDraw) {
	scale := GetConfig().Scale

	imd.Color = colornames.Red
	imd.Push(pixel.V(float64(b.X)*scale, float64(b.Y)*scale))

	imd.Circle(1*scale, 0)
}

func NewBall(b *Board) *Ball {
	ball := new(Ball)

	ball.X = 2
	ball.Y = 5

	ball.VectorX = 1
	ball.VectorY = 1

	ball.Board = b

	return ball
}

func (b *Ball) Move() {
	newX := b.X + 1*b.VectorX
	newY := b.Y + 1*b.VectorY

	if newX == 0 || newX == b.Board.SizeX {
		b.VectorX *= -1
	}
	if newY == 0 || newY == b.Board.SizeY {
		b.VectorY *= -1
	}

	b.X = newX
	b.Y = newY
}
