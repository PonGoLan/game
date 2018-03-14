package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	Number int

	X    int
	Y    int
	Size int

	ButtonUp    pixelgl.Button
	ButtonDown  pixelgl.Button
	ButtonLeft  pixelgl.Button
	ButtonRight pixelgl.Button

	Board *Board
}

func (p *Player) Draw(imd *imdraw.IMDraw) {
	scale := GetConfig().Scale

	imd.Color = pixel.RGB(1, 0, 1)
	imd.Push(pixel.V(float64(p.X)*scale, float64(p.Y)*scale))

	imd.Push(pixel.V(float64(p.X)*scale, float64(p.Y+p.Size)*scale))

	imd.Line(1 * scale)
}

func NewPlayer(playerNumber int, b *Board) *Player {
	player := new(Player)

	player.Number = playerNumber

	player.Size = 20

	if playerNumber == 1 {
		player.ButtonUp = pixelgl.KeyW
		player.ButtonDown = pixelgl.KeyS
		player.ButtonLeft = pixelgl.KeyA
		player.ButtonRight = pixelgl.KeyD

		player.X = 1
		player.Y = 1
	} else if playerNumber == 2 {
		player.ButtonUp = pixelgl.KeyUp
		player.ButtonDown = pixelgl.KeyDown
		player.ButtonLeft = pixelgl.KeyLeft
		player.ButtonRight = pixelgl.KeyRight

		player.X = 150 - 1
		player.Y = 100 - player.Size - 1
	}

	player.Board = b

	return player
}

func (p *Player) SetPosition(x, y int) {
	p.X = x
	p.Y = y
}
func (p *Player) GetPosition() (int, int) {
	return p.X, p.Y
}

func (p *Player) Move(offsetX, offsetY int) {
	newX := p.X + offsetX
	newY := p.Y + offsetY
	if p.Number == 1 {
		if newX < p.Board.SizeX/2 && newX > 0 {
			p.X = newX
		}
	}
	if p.Number == 2 {
		if newX > p.Board.SizeX/2 && newX < p.Board.SizeX {
			p.X = newX
		}
	}

	if newY+p.Size <= p.Board.SizeY && newY >= 0 {
		p.Y = newY
	}
}

func (p *Player) HandleWindowEvents(win *pixelgl.Window) {
	if win.Pressed(p.ButtonUp) {
		p.Move(0, 1)
	}
	if win.Pressed(p.ButtonDown) {
		p.Move(0, -1)
	}
	if win.Pressed(p.ButtonLeft) {
		p.Move(-1, 0)
	}
	if win.Pressed(p.ButtonRight) {
		p.Move(1, 0)
	}
}
