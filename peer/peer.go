package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "MutualExclusion/grpc"
)

type Peer struct {
	proto.UnimplementedMutualExlusionServiceServer
	name    string
	address string
	portIn  int
	portOut int
}

var (
	time     = 0 // Lamport variable
	confFile = "confFile.txt"
)

func main() {

	// read from confFile.txt and set the values

	peer := &Peer{
		// edit this with new variables
		name:    "peer",
		address: "127.0.0.1",
		portIn:  5000,
		portOut: 50001,
	}

	// open the port to new connection, when new one appen we have to create a connection also in the other way
	StartListen(peer)

	// Connect to the others client
	connectToOthersPeer(peer)
}

func StartListen(peer *Peer) {
	// Create a new grpc server
	grpcPeer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", peer.address, strconv.Itoa(peer.portIn)))

	if err != nil {
		log.Fatalf("Could not create the peer %v", err)
	}
	log.Printf("Started peer receiving at address: %s and at port: %d\n", peer.address, peer.portIn)

	// Register the grpc service
	proto.RegisterMutualExlusionServiceServer(grpcPeer, peer)
	serveError := grpcPeer.Serve(listener)

	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func connectToOthersPeer(peer *Peer) {

	// Connect to others peer

	serverConnection := connectToServer()

	for {
		var text string
		log.Printf("Press enter to do mutual execution: ")
		fmt.Scanln(&text)

		increaseTime() // an event occurred
	}
}

func (c *Peer) AskPermission(ctx context.Context, in *proto.Question) (*proto.Answer, error) {
	// to do implement
	return &proto.Answer{
		Reply: false,
	}, nil
}

func connectToPeer() proto.MutualExlusionServiceClient {
	// Dial the server at the specified port.
	conn, err := grpc.Dial(*serverAddr+":"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect to peer at port %d", *serverPort)
	} else {
		log.Printf("Connected to the peer at port %d\n", *serverPort)
	}
	return proto.NewMutualExlusionServiceClient(conn)
}

func increaseTime() {
	time++
}

func setTime(received int) {
	max := math.Max(float64(received), float64(time))
	time = int(max + 1)
}
