package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/alextanhongpin/grpc-openid/app/database"
	"github.com/alextanhongpin/grpc-openid/app/queue"
	pb "github.com/alextanhongpin/grpc-openid/auth"
	"github.com/alextanhongpin/grpc-openid/model"
)

type authserver struct{}

type Environment struct {
	Database *database.Database
	Queue    *queue.Queue
}

var env Environment

func (s *authserver) Login(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("at login route!", msg)

	user := model.User{
		Email:    msg.Email,
		Password: msg.Password,
	}

	if err := env.Database.Ref.Where("email = ?", user.Email).Find(&user).Error; err == nil {
		log.Println("found user", user, err)
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

	env.Queue.Ref.Publish("NewUser", user)

	return &pb.AuthResponse{
		AccessToken: "1234567890",
		Id:          string(user.ID),
	}, nil
}

func (s *authserver) Register(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println("at login route!")
	return &pb.AuthResponse{
		Error:            "",
		ErrorDescription: "",
		AccessToken:      "",
		Id:               "",
	}, nil
}

func main() {
	var err error
	env = Environment{
		Database: database.New(),
		Queue:    queue.New(),
	}
	// Remember to close the database
	defer env.Database.Ref.Close()

	defer env.Queue.Ref.Close()

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
