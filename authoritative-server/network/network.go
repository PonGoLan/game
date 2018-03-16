package pongserver

import (
	"fmt"
	"math"

	instances "github.com/PonGoLan/game/authoritative-server/instances"
	pb "github.com/PonGoLan/game/communication-protocol"
	pong "github.com/PonGoLan/game/shared"
	pongutils "github.com/PonGoLan/game/utils"

	"golang.org/x/net/context"
)

type Server struct{}

var (
	playerHashcodes map[int]string
)

func init() {
	playerHashcodes = make(map[int]string)
}

func (s *Server) SetPlayerPosition(ctx context.Context, in *pb.SetPlayerPositionRequest) (*pb.SetPlayerPositionReply, error) {
	// fmt.Printf("(handshake) %s\n", in.Handshake)
	instance := instances.GetInstanceWithHash(in.Handshake)
	game := instance.GetGame()

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
func (s *Server) GetBallPosition(ctx context.Context, in *pb.GetBallPositionRequest) (*pb.GetBallPositionReply, error) {
	instance := instances.GetInstanceWithHash(in.Handshake)
	game := instance.GetGame()

	x, y := game.Ball.GetPosition()

	return &pb.GetBallPositionReply{
		X: int32(x),
		Y: int32(y),
	}, nil
}

func (s *Server) IdentifyPlayer(ctx context.Context, in *pb.IdentifyPlayerRequest) (*pb.IdentifyPlayerReply, error) {
	// create a random handshake for the player
	handshake := pongutils.CreateRandomHandshake()
	fmt.Printf("ROOM = %v\n", in.Room)
	instance := instances.GetInstance(in.Room)
	if instance == nil {
		instance = instances.Create(in.Room)
		instances.LinkHashToRoom(handshake, in.Room)
	}

	playerNumber := instance.AddPlayer()
	// playerHashcodes[playerNumber] = handshake

	return &pb.IdentifyPlayerReply{
		PlayerNumber: int32(playerNumber),
		Handshake:    handshake,
	}, nil
}

func (s *Server) GetOpponent(ctx context.Context, in *pb.GetOpponentRequest) (*pb.GetOpponentReply, error) {
	instance := instances.GetInstanceWithHash(in.Handshake)
	game := instance.GetGame()

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

func (s *Server) GetScore(ctx context.Context, in *pb.GetScoreRequest) (*pb.GetScoreReply, error) {
	instance := instances.GetInstanceWithHash(in.Handshake)
	game := instance.GetGame()

	return &pb.GetScoreReply{
		Score0: int32(game.Score[0]),
		Score1: int32(game.Score[1]),
	}, nil
}
