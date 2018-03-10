package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	X    int
	Y    int
	Size int

	Board *Board
}

func (p *Player) Draw(imd *imdraw.IMDraw) {
	scale := GetConfig().Scale

	imd.Color = pixel.RGB(1, 0, 1)
	imd.Push(pixel.V(float64(p.X)*scale, float64(p.Y)*scale))

	imd.Push(pixel.V(float64(p.X)*scale, float64(p.Y+p.Size)*scale))

	imd.Line(1 * scale)
}

func NewPlayer(b *Board) *Player {
	player := new(Player)

	player.X = 0
	player.Y = 0

	player.Size = 20

	player.Board = b

	return player
}

func (p *Player) Move(offsetX, offsetY int) {
	newX := p.X + offsetX
	newY := p.Y + offsetY
	if newX <= p.Board.SizeX && newX >= 0 {
		p.X = newX
	}
	if newY+p.Size <= p.Board.SizeY && newY >= 0 {
		p.Y = newY
	}
}

func (p *Player) HandleWindowEvents(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyUp) {
		p.Move(0, 1)
	}
	if win.Pressed(pixelgl.KeyDown) {
		p.Move(0, -1)
	}
	if win.Pressed(pixelgl.KeyLeft) {
		p.Move(-1, 0)
	}
	if win.Pressed(pixelgl.KeyRight) {
		p.Move(1, 0)
	}
}
