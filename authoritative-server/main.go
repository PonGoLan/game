package main

import (
	"log"
	"net"
	"time"

	serverdisplay "github.com/PonGoLan/game/authoritative-server/display"
	"github.com/PonGoLan/game/authoritative-server/instances"
	pongserver "github.com/PonGoLan/game/authoritative-server/network"
	pb "github.com/PonGoLan/game/communication-protocol"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func handleNetworkConnections() {
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

func serverHandling() {
	t := time.Tick(time.Second / 2)
	instancesManager := instances.Get()

	for 42 == 42 {
		select {
		case <-t:
			{
				instancesManager.RemoveStoppedInstance()
				// refresh the display
				serverdisplay.Print()
			}
		default:
		}
	}
}

func main() {
	go handleNetworkConnections()
	serverHandling()
}
