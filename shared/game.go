package pong

import (
	"fmt"

	"github.com/faiface/pixel/imdraw"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	Score map[int]int

	Players [2]*Player
	Ball    *Ball
}

func NewGame() *Game {
	g := new(Game)

	g.Score = make(map[int]int)

	return g
}

func (g *Game) AddPoint(playerId int) {
	g.Score[playerId]++
}

func (g *Game) DrawPlayers(imd *imdraw.IMDraw) {
	for _, p := range g.Players {
		p.Draw(imd)
	}
}

func (g *Game) DrawScore(win *pixelgl.Window) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(770, 940), basicAtlas)

	fmt.Fprintf(txt, "%02d", g.Score[1])

	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 4))

	/// -----------

	txt = text.New(pixel.V(675, 940), basicAtlas)

	fmt.Fprintf(txt, "%02d", g.Score[2])

	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 4))
}
