package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	pb "github.com/PonGoLan/game/communication-protocol"
	pong "github.com/PonGoLan/game/shared"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func run() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPongerClient(conn)

	ctx := context.Background()
	// defer cancel()

	cfg := pixelgl.WindowConfig{
		Title:  "PonGoLan - Client",
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
	game.Players[0] = player1
	game.Players[1] = player2

	ball := pong.NewBall(board)

	aTick := time.Tick(time.Second / 128)

	identificationResponse, err := c.IdentifyPlayer(ctx, &pb.IdentifyPlayerRequest{})
	log.Printf("Player Number  : %d\n", identificationResponse.PlayerNumber)
	log.Printf("Handshake used : %s\n", identificationResponse.Handshake)

	player := game.Players[identificationResponse.PlayerNumber]

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
			// pong.BallPlayerCollision(ball, player1)
			// pong.BallPlayerCollision(ball, player2)
			// ball.Move(game)
			player.HandleWindowEvents(win)

			// Update current player position
			_, err := c.SetPlayerPosition(ctx, &pb.SetPlayerPositionRequest{
				Handshake:    identificationResponse.Handshake,
				PlayerNumber: identificationResponse.PlayerNumber,
				X:            int32(player.X),
				Y:            int32(player.Y),
			})
			if err != nil {
				log.Printf("could not set pos: %v", err)
			}

			// Update ball position
			r, err := c.GetBallPosition(ctx, &pb.GetBallPositionRequest{})
			if err != nil {
				log.Printf("could not get ball position: %v", err)
			}
			if r != nil {
				ball.SetPosition(int(r.X), int(r.Y))
			}

			// player2.HandleWindowEvents(win)
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
