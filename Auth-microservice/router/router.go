package router

import (
	"Auth-microservice/config"
	"Auth-microservice/handlers"
	"Auth-microservice/proto"
	"log"
	"net"
	"google.golang.org/grpc"
)

func InitRouter() {
	lis, err := net.Listen("tcp", ":" + config.LoadEnv().PORT)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &handlers.AuthServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}