package model

import "github.com/jinzhu/gorm"

// User represents the user schema for the database
type User struct {
	gorm.Model
	UserID   string
	Email    string
	Password string
}
