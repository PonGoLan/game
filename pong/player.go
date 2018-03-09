package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Player struct {
	X    int
	Y    int
	Size int
}

func (p *Player) Draw(imd *imdraw.IMDraw) {
	scale := GetConfig().Scale

	imd.Color = pixel.RGB(1, 0, 1)
	imd.Push(pixel.V(float64(p.X)*scale, float64(p.Y)*scale))

	imd.Push(pixel.V(float64(p.X)*scale, float64(p.Y+p.Size)*scale))

	imd.Line(3)
}

func NewPlayer() *Player {
	player := new(Player)

	player.X = 0
	player.Y = 0

	player.Size = 10

	return player
}
