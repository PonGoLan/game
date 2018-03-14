package main

import (
	"log"
	"net"
	"time"

	pb "github.com/PonGoLan/game/communication-protocol"
	pong "github.com/PonGoLan/game/shared"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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
)

type server struct{}

func (s *server) SetPlayerPosition(ctx context.Context, in *pb.SetPlayerPositionRequest) (*pb.SetPlayerPositionReply, error) {
	oldX, oldY := game.Players[in.PlayerNumber].GetPosition()

	if oldX != int(in.X) || oldY != int(in.Y) {
		log.Printf("MOVE [%d] to (%d, %d)\n", in.PlayerNumber, in.X, in.Y)
	}
	game.Players[in.PlayerNumber].SetPosition(int(in.X), int(in.Y))

	return &pb.SetPlayerPositionReply{
		PlayerNumber: in.PlayerNumber,
		X:            in.X,
		Y:            in.Y,
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
	playerHashcodes[1] = "lol"

	return &pb.IdentifyPlayerReply{
		PlayerNumber: 1,
		Handshake:    "lol",
	}, nil
}

func init() {
	game = pong.NewGame()
	playerHashcodes = make(map[int]string)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "PonGoLAN - Server",
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
	player1 := pong.NewPlayer(1, board)
	player2 := pong.NewPlayer(2, board)

	game.Players[0] = player1
	game.Players[1] = player2

	ball := pong.NewBall(board)
	game.Ball = ball

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
