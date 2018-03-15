package pongclient

import (
	"log"

	"golang.org/x/net/context"

	pb "github.com/PonGoLan/game/communication-protocol"
	pong "github.com/PonGoLan/game/shared"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type PongClient struct {
	Conn          *grpc.ClientConn
	Context       context.Context
	PongerService pb.PongerClient

	handshake string
}

var (
	client *PongClient
)

func init() {
	client = new(PongClient)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client.Conn = conn
	// defer client.Conn.Close()
	client.PongerService = pb.NewPongerClient(conn)
	client.Context = context.Background()
}

// Get the singleton
func Get() *PongClient { return client }

// Identify the client
// we retrieve the Handshake used for further discussion with
// the server. We return the assigned player number (by the server)
func (pc *PongClient) Identify() int32 {
	identificationResponse, err := pc.PongerService.IdentifyPlayer(pc.Context, &pb.IdentifyPlayerRequest{})
	if err != nil {
		log.Fatalf("Failed to handshake the server: %v\n", err)
	}

	log.Printf("Player Number  : %d\n", identificationResponse.PlayerNumber)
	log.Printf("Handshake used : %s\n", identificationResponse.Handshake)

	pc.handshake = identificationResponse.GetHandshake()

	return identificationResponse.PlayerNumber
}

// SendPlayerPosition send the current position of the player passed
// as parameter to the server
func (pc *PongClient) SendPlayerPosition(player *pong.Player) (int, int, error) {
	reply, err := pc.PongerService.SetPlayerPosition(pc.Context, &pb.SetPlayerPositionRequest{
		Handshake:    pc.handshake,
		PlayerNumber: int32(player.Number),
		X:            int32(player.X),
		Y:            int32(player.Y),
	})
	if err != nil {
		log.Printf("[ERR] couldnt send position\n")
		return 0, 0, err
	}
	return int(reply.X), int(reply.Y), nil
}

// GetBallPosition retrieve the ball position from the server
// we expect then the client to correct its own state if needed
func (pc *PongClient) GetBallPosition() (int, int, error) {
	reply, err := pc.PongerService.GetBallPosition(pc.Context, &pb.GetBallPositionRequest{})
	if err != nil {
		log.Printf("[ERR] couldnt get ball position")
		return 0, 0, err
	}
	return int(reply.X), int(reply.Y), nil
}

// GetOpponent get opponent position
// client will update its state from this
func (pc *PongClient) GetOpponent(player *pong.Player) (int, int, int, error) {
	reply, err := pc.PongerService.GetOpponent(pc.Context, &pb.GetOpponentRequest{
		PlayerNumber: int32(player.Number),
		Handshake:    pc.handshake,
	})
	if err != nil {
		log.Printf("[ERR] couldnt get opponent position")
		return 0, 0, 0, err
	}
	return int(reply.PlayerNumber), int(reply.X), int(reply.Y), nil
}
