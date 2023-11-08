package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"strconv"
	"sync"

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
	startListen(peer)

	// Connect to the others client
	connectToOthersPeer(peer)
}

func StartListen() {
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.address, strconv.Itoa(server.port)))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at address: %s and at port: %d\n", server.address, server.port)

	// Register the grpc service
	proto.RegisterChittyChatServiceServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)

	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func connect(peer *Peer) {

	// Connect to others peer

	serverConnection := connectToServer()

	client_reference := &proto.ClientReference{
		ClientAddress: client.address,
		ClientPort:    int32(client.portNumber),
		ClientName:    client.name,
	}

	// Get a stream to the server
	stream, err := serverConnection.SendMessage(context.Background())
	increaseTime() // an event occurred
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Printf("[Lamport Time: %d] Connected to server", time)
	}

	increaseTime() // an event occurred
	// Connect message from client to server
	connectMessage := &proto.Message{
		Text:            "connect message",
		Type:            int32(Connect),
		ClientReference: client_reference,
		Time:            int32(time),
	}

	if err := stream.Send(connectMessage); err != nil {
		log.Fatalf("Error while sending connection message: %v", err)
	} else {
		log.Printf("[Lamport Time: %d] Sent connect message to server", time)
	}

	// wait for go routine
	var wg sync.WaitGroup
	wg.Add(1)

	// Create go-routine for reading messages from server
	go func() {
		for {
			msg, err := stream.Recv()
			// update local time
			if msg.Time != 0 {
				setTime(int(msg.Time))
			}
			if err != nil {
				log.Fatalf("Error while receiving message: %v", err)
			}

			if msg.Type == Ack { // I receive the ack for the disconnected message that i sent
				defer wg.Done() // inform the main thread to terminate
				break
			} else {
				// R4: When a client receives a broadcasted message, it has to write the message and the current logical timestamp
				//senderReference := msg.ClientReference.ClientAddress + ":" + strconv.Itoa(int(msg.ClientReference.ClientPort))
				log.Printf("[Lamport Time: %d, Name: %s] %s\n", time, msg.ClientReference.ClientName, msg.Text)
				//log.Printf("[%s] sent %s serverTime %d  localTime %d \n", senderReference, msg.Text, msg.Time, time)
				log.Printf("Enter the content of the message ('exit' to quit): ") // just for better user experience
			}

		}
	}()

	for {
		var text string
		log.Printf("Enter the content of the message ('exit' to quit): ")
		fmt.Scanln(&text)

		if len(text) > 128 {
			log.Println("Message cannot be more than 128 characters")
			continue
		}

		// R7: Chat clients can drop out at any time
		if text == "exit" {
			increaseTime() // an event occurred
			// Disconnect message from client to server
			diconnectMessage := &proto.Message{
				Text:            "disconnect message",
				Type:            int32(Disconnect),
				ClientReference: client_reference,
				Time:            int32(time),
			}
			if err := stream.Send(diconnectMessage); err != nil {
				log.Fatalf("Error while sending disconnection message: %v", err)
			}
			break
		}

		increaseTime() // an event occurred
		// R2: Clients in Chitty-Chat can Publish a valid chat message at any time they wish
		err := stream.Send(&proto.Message{
			Text:            text,
			Type:            int32(Publish),
			ClientReference: client_reference,
			Time:            int32(time),
		})
		if err != nil {
			log.Fatalf("Error while sending message: %v", err)
		}
	}

	// wait for response to disconnect message
	wg.Wait()
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
