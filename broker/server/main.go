package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/thanders/quiz-startup/broker/proto"
	"google.golang.org/grpc"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 50051, "The server port")
)

var addr string = "0.0.0.0:60061"

type Server struct {
	pb.BrokerServiceServer
}

type brokerServiceServer struct {
	pb.UnimplementedBrokerServiceServer
	savedFeatures *pb.BrokerResponse // read-only after initialized
}

func newServer() *brokerServiceServer {
	s := &brokerServiceServer{}
	return s
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startBrokerListener() {
	lis, err := net.Listen("tcp", addr)

	handleError(err)

	s := grpc.NewServer()
	pb.RegisterBrokerServiceServer(s, &Server{})

	fmt.Printf("Broker Server - Listening for client RPC on %s\n", addr)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}

func main() {

	rabbitConn, err := connectToRabbit()
	handleError(err)

	defer rabbitConn.Close()

	// don't continue until etcd is ready
	// etcConn, err := connectToEtcd()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer etcConn.Close()

	app := Config{
		Rabbit: rabbitConn,
		//Etcd:   etcConn,
	}
	var testPayload = &Payload{
		ServiceName: "questions",
		Data:        "test",
	}
	app.HandleSubmission("log", testPayload)

	startBrokerListener()
}
