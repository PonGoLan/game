package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Ball struct {
	X int
	Y int

	PredictedX int
	PredictedY int

	VectorX int
	VectorY int

	Board *Board
}

func (b *Ball) Draw(imd *imdraw.IMDraw) {
	scale := GetConfig().Scale

	imd.Color = colornames.Red
	imd.Push(pixel.V(float64(b.X)*scale, float64(b.Y)*scale))

	imd.Circle(1*scale, 0)

	// imd.Color = colornames.Red
	// imd.Push(pixel.V(float64(b.PredictedX)*scale, float64(b.PredictedY)*scale))
	// imd.Push(pixel.V(float64(b.PredictedX+5*b.VectorX)*scale, float64(b.PredictedY+5*b.VectorY)*scale))

	// imd.Line(2)
}

func (b *Ball) SetPosition(x, y int) {
	b.X = x
	b.Y = y
}
func (b *Ball) GetPosition() (int, int) {
	return b.X, b.Y
}

func NewBall(b *Board) *Ball {
	ball := new(Ball)

	ball.X = 2
	ball.Y = 5

	ball.VectorX = 1
	ball.VectorY = 1

	ball.PredictedX = ball.X + 1*ball.VectorX
	ball.PredictedY = ball.Y + 1*ball.VectorY

	ball.Board = b

	return ball
}

func (b *Ball) Move(game *Game) {
	newX := b.X + b.VectorX
	newY := b.Y + b.VectorY

	if newX == 0 {
		game.AddPoint(1)
	}
	if newX == b.Board.SizeX {
		game.AddPoint(2)
	}
	if newX == 0 || newX == b.Board.SizeX {
		b.VectorX *= -1
		return
	}
	if newY == 0 || newY == b.Board.SizeY {
		b.VectorY *= -1
		return
	}

	b.X = newX
	b.Y = newY

	b.PredictedX = b.X + 1*b.VectorX
	b.PredictedY = b.Y + 1*b.VectorY
}
