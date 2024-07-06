package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/XecFardin/grpc-user-service/proto"
	"github.com/XecFardin/grpc-user-service/server"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server.UserServiceServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Channel to listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run server in a goroutine to allow for graceful shutdown
	go func() {
		log.Println("Server is listening on port 50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Block until a signal is received
	<-stop

	log.Println("Shutting down server...")

	// Create a deadline to wait for
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully stop the server
	s.GracefulStop()

	log.Println("Server gracefully stopped")
}
