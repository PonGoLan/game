package main

import (
	"time"

	client "github.com/PonGoLan/game/client/network"
	pong "github.com/PonGoLan/game/shared"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	frames = 0
	second = time.Tick(time.Second)
)

func run() {
	win, err := pixelgl.NewWindow(pong.GetConfig().WindoWConfig)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	imd := imdraw.New(nil)

	game := pong.NewGame()
	board := pong.NewBoard()
	game.Players[0] = pong.NewPlayer(0, board)
	game.Players[1] = pong.NewPlayer(1, board)

	ball := pong.NewBall(board)

	aTick := time.Tick(time.Second / 128)

	// [network] identify current client
	playerNumber := client.Get().Identify()
	player := game.Players[playerNumber]

	for !win.Closed() {
		pong.ApplyMatrixToWindow(win)

		board.Draw(imd)
		game.DrawPlayers(imd)
		ball.Draw(imd)

		win.Clear(colornames.Black)
		imd.Draw(win)

		game.DrawScore(win)

		win.Update()
		imd.Clear()

		frames++
		select {
		case <-second:
			win.SetTitle(pong.GetConfig().GetTitle(frames))
			client.Get().GetScore(game)
			frames = 0
		case <-aTick:
			// pong.BallPlayerCollision(ball, player1)
			// pong.BallPlayerCollision(ball, player2)
			// ball.Move(game)
			player.HandleWindowEvents(win)

			// Update current player position
			pX, pY, err := client.Get().SendPlayerPosition(player)
			if err == nil {
				player.SetPosition(pX, pY)
			}

			// Update ball position
			ballX, ballY, err := client.Get().GetBallPosition()
			if err == nil {
				ball.SetPosition(ballX, ballY)
			}

			// update opponent position
			oppNumber, oppX, oppY, err := client.Get().GetOpponent(player)
			if err == nil {
				game.Players[oppNumber].SetPosition(oppX, oppY)
			}
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
