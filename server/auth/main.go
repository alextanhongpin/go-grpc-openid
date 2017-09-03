package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/alextanhongpin/grpc-openid/app/database"
	"github.com/alextanhongpin/grpc-openid/app/queue"
	"github.com/alextanhongpin/grpc-openid/model"
	pb "github.com/alextanhongpin/grpc-openid/proto/auth"
	"github.com/alextanhongpin/grpc-openid/utils/password"
)

type authserver struct{}

// Environment defines the application dependencies
type Environment struct {
	Database *database.Database
	Queue    *queue.Queue
}

var env Environment

func (s *authserver) Login(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthResponse, error) {
	hash, err := password.Generate(msg.Password)
	if err != nil {
		return &pb.AuthResponse{
			Error:            "Internal Server Error",
			ErrorDescription: "Error comparing password",
		}, err
	}
	user := model.User{
		Email:    msg.Email,
		Password: hash,
	}

	if err := env.Database.Ref.Where("email = ? AND password = ?", user.Email, user.Password).Find(&user).Error; err != nil {
		return &pb.AuthResponse{
			Error:            "Unauthorized",
			ErrorDescription: "Email or password is incorrect",
		}, nil
	}

	return &pb.AuthResponse{
		AccessToken: "123456",
		Id:          string(user.ID),
	}, nil
}

func (s *authserver) Register(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthResponse, error) {
	user := model.User{
		Email: msg.Email,
	}

	if err := env.Database.Ref.Where("email = ?", user.Email).Find(&user).Error; err == nil {
		return &pb.AuthResponse{
			Error:            "Unauthorized",
			ErrorDescription: "User already exists",
		}, nil
	}

	hash, err := password.Generate(msg.Password)
	if err != nil {
		return &pb.AuthResponse{
			Error:            "Internal Server Error",
			ErrorDescription: "Error generating password hash",
		}, err
	}

	user.Password = hash

	if err := env.Database.Ref.Create(&user).Error; err != nil {
		return &pb.AuthResponse{
			Error:            "Bad request",
			ErrorDescription: "Fail to create user",
		}, nil
	}

	// Publish webhook event, should do a feature toggle here
	env.Queue.Ref.Publish("NewUser", user)

	return &pb.AuthResponse{
		AccessToken: "1234567890",
		Id:          string(user.ID),
	}, nil
}

func main() {

	db := database.New()
	defer db.Ref.Close()

	q := queue.New()
	defer q.Ref.Close()

	env = Environment{
		Database: db,
		Queue:    q,
	}

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
