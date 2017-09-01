package app

import (
	"github.com/alextanhongpin/grpc-openid/app/database"
	"github.com/alextanhongpin/grpc-openid/app/queue"
)

// Environment represents the dependencies used by the application
type Environment struct {
	Database *database.Database
	Queue    *queue.Queue
}

// New creates a new app with the default configuration
func New() *Environment {
	return &Environment{
		Database: database.New(),
		Queue:    queue.New(),
	}
}
