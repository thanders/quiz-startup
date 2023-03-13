package main

import (
	"context"
	"log"

	pb "github.com/thanders/quiz-startup/broker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:60061"

func doCreateGame(c pb.BrokerServiceClient) {
	res, err := c.CreateGame(context.Background(), &pb.BrokerRequest{
		NumberOfQuestions: "5",
		NumberOfPlayers:   "2",
	})

	if err != nil {
		log.Fatalf("Could not broker: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.GameId)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()
	c := pb.NewBrokerServiceClient(conn)

	doCreateGame(c)
}
