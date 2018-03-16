package main

import (
	"log"
	"net"
	"time"

	pongserver "github.com/PonGoLan/game/authoritative-server/network"
	pb "github.com/PonGoLan/game/communication-protocol"
	pong "github.com/PonGoLan/game/shared"

	"github.com/faiface/pixel/pixelgl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

var (
	game *pong.Game

	playerHashcodes map[int]string
	nextPlayer      int

	ticks  = 0
	second = time.Tick(time.Second)
)

func init() {
	game = pong.NewGame()
	playerHashcodes = make(map[int]string)
}

func run() {
	// cfg := pixelgl.WindowConfig{
	// 	Title:  "PonGoLAN - Server",
	// 	Bounds: pixel.R(0, 0, 1024, 768),
	// 	VSync:  true,
	// }
	// win, err := pixelgl.NewWindow(cfg)
	// if err != nil {
	// 	panic(err)
	// }

	// win.Clear(colornames.Skyblue)

	// imd := imdraw.New(nil)

	// board := pong.NewBoard()
	// player1 := pong.NewPlayer(0, board)
	// player2 := pong.NewPlayer(1, board)
	//
	// game.Players[0] = player1
	// game.Players[1] = player2
	//
	// ball := pong.NewBall(board)
	// game.Ball = ball

	aTick := time.Tick(time.Second / 64)

	for 1 == 1 {
		// pong.ApplyMatrixToWindow(win)

		// board.Draw(imd)
		// player1.Draw(imd)
		// player2.Draw(imd)
		// ball.Draw(imd)

		// win.Clear(colornames.Black)
		// imd.Draw(win)

		// game.DrawScore(win)

		// win.Update()
		// imd.Clear()

		select {
		case <-second:
			log.Printf("[TICKS/s]: %d\n", ticks)
			ticks = 0
		case <-aTick:
			ticks++
			// pong.BallPlayerCollision(ball, player1)
			// pong.BallPlayerCollision(ball, player2)
			// ball.Move(game)
		default:
		}
	}
}

func stuff() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPongerServer(s, &pongserver.Server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	go stuff()

	pixelgl.Run(run)
}
