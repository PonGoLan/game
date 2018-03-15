package main

import (
	"log"
	"math"
	"net"
	"time"

	pb "github.com/PonGoLan/game/communication-protocol"
	pong "github.com/PonGoLan/game/shared"

	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/net/context"
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

type server struct{}

func (s *server) SetPlayerPosition(ctx context.Context, in *pb.SetPlayerPositionRequest) (*pb.SetPlayerPositionReply, error) {
	oldX, oldY := game.Players[in.PlayerNumber].GetPosition()

	if oldX != int(in.X) || oldY != int(in.Y) {
		// log.Printf("MOVE [%d] to (%d, %d)\n", in.PlayerNumber, in.X, in.Y)
		dx := math.Abs(float64(oldX - int(in.X)))
		dy := math.Abs(float64(oldY - int(in.Y)))
		if dx <= 1 && dy <= 1 {
			game.Players[in.PlayerNumber].SetPosition(int(in.X), int(in.Y))
		}
	}

	newX, newY := game.Players[in.PlayerNumber].GetPosition()

	return &pb.SetPlayerPositionReply{
		PlayerNumber: in.PlayerNumber,
		X:            int32(newX),
		Y:            int32(newY),
	}, nil
}
func (s *server) GetBallPosition(ctx context.Context, in *pb.GetBallPositionRequest) (*pb.GetBallPositionReply, error) {
	x, y := game.Ball.GetPosition()

	return &pb.GetBallPositionReply{
		X: int32(x),
		Y: int32(y),
	}, nil
}

func (s *server) IdentifyPlayer(ctx context.Context, in *pb.IdentifyPlayerRequest) (*pb.IdentifyPlayerReply, error) {
	nextPlayer++
	playerHashcodes[nextPlayer] = "lol"

	return &pb.IdentifyPlayerReply{
		PlayerNumber: int32(nextPlayer) - 1,
		Handshake:    "lol",
	}, nil
}

func (s *server) GetOpponent(ctx context.Context, in *pb.GetOpponentRequest) (*pb.GetOpponentReply, error) {
	var player *pong.Player

	if in.PlayerNumber == 0 {
		player = game.Players[1]
	} else if in.PlayerNumber == 1 {
		player = game.Players[0]
	}

	return &pb.GetOpponentReply{
		PlayerNumber: int32(player.Number),
		X:            int32(player.X),
		Y:            int32(player.Y),
	}, nil
}

func (s *server) GetScore(ctx context.Context, in *pb.GetScoreRequest) (*pb.GetScoreReply, error) {
	return &pb.GetScoreReply{
		Score0: int32(game.Score[0]),
		Score1: int32(game.Score[1]),
	}, nil
}

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

	board := pong.NewBoard()
	player1 := pong.NewPlayer(0, board)
	player2 := pong.NewPlayer(1, board)

	game.Players[0] = player1
	game.Players[1] = player2

	ball := pong.NewBall(board)
	game.Ball = ball

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
			pong.BallPlayerCollision(ball, player1)
			pong.BallPlayerCollision(ball, player2)
			ball.Move(game)
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
	pb.RegisterPongerServer(s, &server{})

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
