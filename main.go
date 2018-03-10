package main

import (
	"time"

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

	game := pong.NewGame()

	board := pong.NewBoard()
	player1 := pong.NewPlayer(1, board)
	player2 := pong.NewPlayer(2, board)

	ball := pong.NewBall(board)

	aTick := time.Tick(time.Second / 128)

	for !win.Closed() {
		pong.ApplyMatrixToWindow(win)

		board.Draw(imd)
		player1.Draw(imd)
		player2.Draw(imd)
		ball.Draw(imd)

		win.Clear(colornames.Black)
		imd.Draw(win)

		game.DrawScore(win)

		win.Update()
		imd.Clear()

		select {
		case <-aTick:
			pong.BallPlayerCollision(ball, player1)
			pong.BallPlayerCollision(ball, player2)
			ball.Move(game)
			player1.HandleWindowEvents(win)
			player2.HandleWindowEvents(win)
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
