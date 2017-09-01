package main

import (
	"log"
	"net"

	"github.com/alextanhongpin/grpc-openid/app"
	pb "github.com/alextanhongpin/grpc-openid/auth"
	"github.com/alextanhongpin/grpc-openid/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type authserver struct{}

var env *app.Environment

func (s *authserver) Login(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("at login route!")

	userID := "random"
	user := model.User{
		UserID:   userID,
		Email:    "john.doe@mail.com",
		Password: "123456",
	}

	if err := env.Database.Ref.Where("email = ?", user.Email).Find(&user).Error; err != nil || user.UserID != "" {
		return &pb.AuthResponse{
			Error:            "Unauthorized",
			ErrorDescription: "User already exists",
		}, nil
	}

	if err := env.Database.Ref.Create(&user).Error; err != nil {
		return &pb.AuthResponse{
			Error:            "Bad request",
			ErrorDescription: "Fail to create user",
		}, nil
	}

	env.Queue.Ref.Publish("foo", []byte("hello world"))

	return &pb.AuthResponse{
		AccessToken: "1234567890",
		UserId:      userID,
	}, nil
}

func (s *authserver) Register(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("at login route!")
	return &pb.AuthResponse{
		Error:            "",
		ErrorDescription: "",
		AccessToken:      "",
		UserId:           "",
	}, nil
}

func main() {
	var err error
	env = app.New()
	// Remember to close the database
	defer env.Database.Ref.Close()

	// Automigrate
	env.Database.Ref.AutoMigrate(&model.User{})

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &authserver{})
	log.Println("listening to port *:9090")
	grpcServer.Serve(lis)
}
