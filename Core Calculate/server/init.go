package server

import (
	"Core/grpc-gen/core"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"Core/handler"
)

func Init() {
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	core.RegisterCoreServer(s, &handler.Handler{})
	err = s.Serve(lis)
	
	if err != nil {
		log.Fatalf("failed to init grpc on : %v", err)
	}
}