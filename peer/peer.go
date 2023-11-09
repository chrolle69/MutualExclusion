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
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "MutualExclusion/grpc"
)

// link for Ricart & Agrawala algorithm https://www.geeksforgeeks.org/ricart-agrawala-algorithm-in-mutual-exclusion-in-distributed-system/
// we need only the port where we receive messages
// output port is decided automatically randomly by operating system

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
	my_row = flag.Int("row", 1, "Indicate the row of parameter file for this peer") // set with "-row <port>" in terminal
	name   = flag.String("name", "peer", "name of the peer")
	// Lamport variable
	lamport_time = 0
	confFile     = "confFile.csv"
	// default values for address and port
	my_address = "127.0.0.1"
	my_port    = 50050
	// store tcp connection to others peers
	peers = make(map[string]proto.MutualExlusionServiceClient)
	// state of the distributed mutex
	state = Released
	// wait for listen before try to connect to other peers
	wg sync.WaitGroup
)

func main() {
	flag.Parse()

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
		name:    *name,
		address: my_address,
		port:    my_port,
	}
	// wait for opening port to listen
	wg.Add(1)

	// open the port to new connections
	go StartListen(peer)

	wg.Wait()
	// Preparate tcp connection to the others client
	connectToOthersPeer(peer)

	// user interface menu
	doSomething()
}

func StartListen(peer *Peer) {
	// Create a new grpc server
	grpcPeer := grpc.NewServer()

	// Make the peer listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", peer.address, strconv.Itoa(peer.port)))

	if err != nil {
		log.Fatalf("Could not create the peer %v", err)
	}
	log.Printf("Started peer receiving at address: %s and at port: %d\n", peer.address, peer.port)
	wg.Done()

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
			// ignore corrupted rows and me
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
	// Dial doesn't check if the peer at that address:host is effectivly on (simply prepare TCP connection)
	conn, err := grpc.Dial(address+":"+strconv.Itoa(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect to peer %s at port %d", address, port)
	} else {
		log.Printf("Created TCP connection to the %s address at port %d\n", address, port)
	}
	return proto.NewMutualExlusionServiceClient(conn)
}

func (peer *Peer) AskPermission(ctx context.Context, in *proto.Question) (*proto.Answer, error) {
	// check if the peer requesting permission is not in the list of connected peers
	// it can be a reconnected peer or one not present in the configuration file
	peerRef := in.ClientReference.ClientAddress + ":" + strconv.Itoa(int(in.ClientReference.ClientPort))
	log.Printf("[Lamport Time: %d, Received Time: %d] Peer [%s] asked for a mutual exection", lamport_time, in.Time, peerRef)
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
	// Ricart–Agrawala Algorithm
	if (state == Held) || (state == Wanted && (in.Time > int32(lamport_time))) {
		// queue the reply (just wait unline i'm done)
		for state == Held || state == Wanted {
			time.Sleep(500 * time.Millisecond)
		}
	}
	log.Printf("Peer [%s] authorized to do mutual exection", peerRef)
	// update lamport time
	setTime(int(in.Time))
	return &proto.Answer{
		Reply: true,
	}, nil

}

func doSomething() {
	for {
		var text string
		log.Printf("Insert 'mutual' to do mutual execution or 'exit' to quit or anything else "+
			"to increment time [Actual Lamport Time: %d] ", lamport_time)
		fmt.Scanln(&text)

		increaseTime() // an event occurred
		if text == "exit" {
			break
		}

		if text != "mutual" {
			continue
		}

		state = Wanted
		/* before starting the phase to grant the permission to do the critical section
		there will be a delay of 5 second to allow to launch simultaneous request on different peer (only for testing)*/
		time.Sleep(5 * time.Second)

		// Peers enters the critical section if it has received the REPLY message from all other sites.
		peerRef := &proto.ClientReference{
			ClientAddress: my_address,
			ClientPort:    int32(my_port),
			ClientName:    *name,
		}
		for index, peer := range peers {
			_, err := peer.AskPermission(context.Background(),
				&proto.Question{
					ClientReference: peerRef,
					Time:            int32(lamport_time),
				})
			if err != nil {
				log.Printf("Peer [%s] no more available, removed from connected peers", index)
				delete(peers, index)
				continue
			}
		}
		// do critical section
		criticalSection()
	}
}

func criticalSection() {
	state = Held
	log.Println("Starting critical section")
	time.Sleep(time.Duration(rand.Intn(4)+10) * time.Second)
	log.Println("Ending critical section")
	state = Released
}

func increaseTime() {
	lamport_time++
}

func setTime(received int) {
	max := math.Max(float64(received), float64(lamport_time))
	lamport_time = int(max + 1)
}
