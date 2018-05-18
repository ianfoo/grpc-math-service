package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ianfoo/grpc-math-service/pb"
	"github.com/ianfoo/grpc-math-service/server"
	"google.golang.org/grpc"
)

const defaultAddr = ":3000"

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = defaultAddr
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting network listener: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMatherServer(grpcServer, &server.Server{})
	grpcServer.Serve(lis)
}
