package main

import (
	"github.com/PonGoLan/pong-go-lan/pong"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "PonGoLan",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	imd := imdraw.New(nil)

	board := pong.NewBoard()
	player := pong.NewPlayer(board)

	for !win.Closed() {
		win.Clear(colornames.Whitesmoke)

		board.Draw(imd)
		player.Draw(imd)

		player.HandleWindowEvents(win)

		imd.Draw(win)
		win.Update()
		imd.Clear()
	}
}

func main() {
	pixelgl.Run(run)
}