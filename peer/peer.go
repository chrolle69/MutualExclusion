package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"os"
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
	my_row   = flag.Int("row", 1, "Indicate the parameters for this peer") // set with "-row <port>" in terminal
	time     = 0                                                           // Lamport variable
	confFile = "confFile.csv"
)

func main() {

	flag.Parse()
	// we will use a param passed through command line
	// read from confFile.txt and set the values

	// Apri il file CSV per la lettura
	csvFile, err := os.Open(confFile)
	if err != nil {
		fmt.Printf("Error while opening CSV file: %v\n", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error in reading CSV file: %v\n", err)
		return
	}

	found := false

	for index, row := range rows {
		if index == *my_row {
			fmt.Printf("Row founded : %s, %s , %s\n", row[0], row[1], row[2])
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Riga not founded\n")
	}
	/*

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
		connectToOthersPeer()*/
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

// Connect to others peer
func connectToOthersPeer() {
	// read csv file
	file, err := os.Open(confFile)
	if err != nil {
		log.Fatalf("Failed to open configuration file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read file data: %v", err)
	}

	// array of peers
	var peers []proto.MutualExlusionServiceClient

	// foreach row in the
	for index, row := range rows {
		if len(row) < 2 || (index == *my_row) {
			log.Printf("Skipped row: %v", row)
			continue
		}

		peerAddress := row[0]
		peerPort, _ := strconv.Atoi(row[1])

		client := connectToPeer(peerAddress, peerPort)
		peers = append(peers, client)
	}

	for {
		var text string
		log.Printf("Press enter to do mutual execution: ")
		fmt.Scanln(&text)

		increaseTime() // an event occurred
		/* Use array
		for index, client := range clients {
			response, err := client.AskPermission(context.Background(), &proto.Question{})
			if err != nil {
				log.Printf("RPC failed: %v", err)
				continue
			}
			log.Printf("Response: %v", response)
		}*/
	}
}

func (peer *Peer) AskPermission(ctx context.Context, in *proto.Question) (*proto.Answer, error) {
	// to do implement
	return &proto.Answer{
		Reply: false,
	}, nil
}

func connectToPeer(address string, port int) proto.MutualExlusionServiceClient {
	// Dial the server at the specified port.
	conn, err := grpc.Dial(address+":"+strconv.Itoa(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect to peer %s at port %d", address, port)
	} else {
		log.Printf("Connected to the peer %s at port %d\n", address, port)
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
