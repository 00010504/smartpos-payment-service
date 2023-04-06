package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Invan2/invan_payment_service/config"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()

	server := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.HttpPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
		return
	}
}
