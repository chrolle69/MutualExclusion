package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "MutualExclusion/grpc"
)

// we need only the port where we receive messages
// output port is decided automatically from operative system
type Peer struct {
	proto.UnimplementedMutualExlusionServiceServer
	name    string
	address string
	port    int
}

// peer states
const (
	Released int = 0
	Wanted       = 1
	Held         = 2
)

var (
	my_row       = flag.Int("row", 1, "Indicate the row of parameter file for this peer") // set with "-row <port>" in terminal
	name         = flag.String("name", "peer", "name of the peer")
	lamport_time = 0 // Lamport variable
	confFile     = "confFile.csv"
	// default values for address and port
	my_address = "127.0.0.1"
	my_port    = 50051
	// store tcp connection to others peers
	peers = make(map[string]proto.MutualExlusionServiceClient)
	state = Released
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// read from confFile.txt and set the peer values
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
			fmt.Printf("Your settings are : %s address, %s port\n", row[0], row[1])
			my_address = row[0]
			my_port, _ = strconv.Atoi(row[1])
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Row with parameters not founded\n")
		return
	}

	peer := &Peer{
		// edit this with new variables
		name:    *name,
		address: my_address,
		port:    my_port,
	}

	// open the port to new connection, when new one appen we have to create a connection also in the other way
	StartListen(peer)

	// Connect to the others client
	connectToOthersPeer(peer)

	doSomething()
}

func StartListen(peer *Peer) {
	// Create a new grpc server
	grpcPeer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", peer.address, strconv.Itoa(peer.port)))

	if err != nil {
		log.Fatalf("Could not create the peer %v", err)
	}
	log.Printf("Started peer receiving at address: %s and at port: %d\n", peer.address, peer.port)

	// Register the grpc service
	proto.RegisterMutualExlusionServiceServer(grpcPeer, peer)
	serveError := grpcPeer.Serve(listener)

	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

// Connect to others peer
func connectToOthersPeer(p *Peer) {
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

	// try to connect to other peers
	for index, row := range rows {
		if len(row) < 2 || (index == *my_row) {
			log.Printf("Skipped row: %v", row)
			continue
		}

		peerAddress := row[0]
		peerPort, _ := strconv.Atoi(row[1])
		peerRef := row[0] + ":" + row[1]
		// retrieve connection
		connection := connectToPeer(peerAddress, peerPort)
		// add to map
		peers[peerRef] = connection
	}
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

func (peer *Peer) AskPermission(ctx context.Context, in *proto.Question) (*proto.Answer, error) {
	// check if the peer requesting permission is not in the list of connected peers
	peerRef := in.ClientReference.ClientAddress + ":" + strconv.Itoa(int(in.ClientReference.ClientPort))
	found := false
	for index := range peers {
		if index == peerRef {
			found = true
			break
		}
	}
	// receive request from a not known peer
	if !found {
		connection := connectToPeer(in.ClientReference.ClientAddress, int(in.ClientReference.ClientPort))
		peers[peerRef] = connection
	}
	/* like slide 15, lecture 7
	if(state == HELD ||
		(state == WANTED &&
		(T,pme) < (Ti,pi)))
		then queue req
		else reply to req */
	if (state == Held) || (state == Wanted) {
		// check lamport time
		//in.Time <> time
		return &proto.Answer{
			Reply: false,
		}, nil
	} else {
		return &proto.Answer{
			Reply: true,
		}, nil
	}
}

func doSomething() {
	for {
		var text string
		log.Printf("Press enter to do mutual execution ('exit' to quit): ")
		fmt.Scanln(&text)

		increaseTime() // an event occurred
		if text == "exit" {
			break
		}

		// authorized to enter critical section
		authorizated := true
		for _, peer := range peers {
			peerRef := &proto.ClientReference{
				ClientAddress: my_address,
				ClientPort:    int32(my_port),
				ClientName:    *name,
			}
			response, err := peer.AskPermission(context.Background(),
				&proto.Question{
					ClientReference: peerRef,
					Time:            int32(lamport_time),
				})
			if err != nil {
				log.Printf("RPC failed, peer no more available: %v", err)
				continue
			}
			log.Printf("Response: %v", response)
			if response.Reply == false {
				authorizated = false
				break
			}
		}
		if authorizated {
			// do something
			log.Println("Starting critical section")
			time.Sleep(time.Duration(rand.Intn(4)+2) * time.Second)
			fmt.Println("Ending critical section")
		}
	}
}

func increaseTime() {
	lamport_time++
}

func setTime(received int) {
	max := math.Max(float64(received), float64(lamport_time))
	lamport_time = int(max + 1)
}
