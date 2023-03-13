package main

import (
	"context"
	"log"

	pb "github.com/thanders/quiz-startup/broker/proto"
)

func (s *Server) CreateGame(ctx context.Context, in *pb.BrokerRequest) (*pb.BrokerResponse, error) {
	log.Printf("New game request - returning to client %v\n", in)
	return &pb.BrokerResponse{
		GameId: "Hello" + in.NumberOfQuestions,
	}, nil
}
