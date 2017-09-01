package main

import (
	"log"
	"net"

	pb "github.com/alextanhongpin/grpc-openid/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type authserver struct{}

func (s *authserver) Login(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("at login route!")
	return &pb.AuthResponse{
		Msg: "hello",
		Err: "",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &authserver{})
	log.Println("listening to port *:9090")
	grpcServer.Serve(lis)
}
